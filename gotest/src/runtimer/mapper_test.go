package runtimer

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unsafe"
	"reflect"
)

const (
	// Maximum number of key/value pairs a bucket can hold.
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits

	// Maximum average load of a bucket that triggers growth is 6.5.
	// Represent as loadFactorNum/loadFactDen, to allow integer math.
	loadFactorNum = 13
	loadFactorDen = 2

	// Maximum key or value size to keep inline (instead of mallocing per element).
	// Must fit in a uint8.
	// Fast versions cannot handle big values - the cutoff size for
	// fast versions in ../../cmd/internal/gc/walk.go must be at most this value.
	maxKeySize   = 128
	maxValueSize = 128

	// data offset should be the size of the bmap struct, but needs to be
	// aligned correctly. For amd64p32 this means 64-bit alignment
	// even though pointers are 32 bit.
	dataOffset = unsafe.Offsetof(struct {
		b bmap
		v int64
	}{}.v)

	// Possible tophash values. We reserve a few possibilities for special marks.
	// Each bucket (including its overflow buckets, if any) will have either all or none of its
	// entries in the evacuated* states (except during the evacuate() method, which only happens
	// during map writes and thus no one else can observe the map during that time).
	empty          = 0 // cell is empty
	evacuatedEmpty = 1 // cell is empty, bucket is evacuated.
	evacuatedX     = 2 // key/value is valid.  Entry has been evacuated to first half of larger table.
	evacuatedY     = 3 // same as above, but evacuated to second half of larger table.
	minTopHash     = 4 // minimum tophash for a normal filled cell.

	// flags
	iterator     = 1 // there may be an iterator using buckets
	oldIterator  = 2 // there may be an iterator using oldbuckets
	hashWriting  = 4 // a goroutine is writing to the map
	sameSizeGrow = 8 // the current map growth is to a new map of the same size

	// sentinel bucket ID for iterator checks
	noCheck = 1<<(8*8) - 1
)

var (
	hans = []rune("赵客缦胡缨吴钩霜雪明银鞍照白马飒沓如流星" +
		"十步杀一人千里不留行事了拂衣去深藏身与名" +
		"闲过信陵饮脱剑膝前横将炙啖朱亥持觞劝侯嬴" +
		"三杯吐然诺五岳倒为轻眼花耳热后意气素霓生" +
		"救赵挥金槌邯郸先震惊千秋二壮士烜赫大梁城" +
		"纵死侠骨香不惭世上英谁能书阁下白首太玄经")
	rander = rand.New(rand.NewSource(time.Now().Unix()))
)

// runtime/hashmap.go
// A header for a Go map.
type hmap struct {
	// Note: the format of the Hmap is encoded in ../../cmd/internal/gc/reflect.go and
	// ../reflect/type.go. Don't change this structure without also changing that code!
	count     int // # live cells == size of map.  Must be first (used by len() builtin)
	flags     uint8
	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
	hash0     uint32 // hash seed

	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

	extra *mapextra // optional fields
}

// mapextra holds fields that are not present on all maps.
type mapextra struct {
	// If both key and value do not contain pointers and are inline, then we mark bucket
	// type as containing no pointers. This avoids scanning such maps.
	// However, bmap.overflow is a pointer. In order to keep overflow buckets
	// alive, we store pointers to all overflow buckets in hmap.overflow and h.map.oldoverflow.
	// overflow and oldoverflow are only used if key and value do not contain pointers.
	// overflow contains overflow buckets for hmap.buckets.
	// oldoverflow contains overflow buckets for hmap.oldbuckets.
	// The indirection allows to store a pointer to the slice in hiter.
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	// nextOverflow holds a pointer to a free overflow bucket.
	nextOverflow *bmap
}

// A bucket for a Go map.
type bmap struct {
	// tophash generally contains the top byte of the hash value
	// for each key in this bucket. If tophash[0] < minTopHash,
	// tophash[0] is a bucket evacuation state instead.
	tophash [bucketCnt]uint8
	// Followed by bucketCnt keys and then bucketCnt values.
	// NOTE: packing all the keys together and then all the values together makes the
	// code a bit more complicated than alternating key/value/key/value/... but it allows
	// us to eliminate padding which would be needed for, e.g., map[int64]int8.
	// Followed by an overflow pointer.
}

func TestHashMap(t *testing.T) {
	mapper := make(map[int64]string)
	num := 50
	maxnum := 200
	keys := make([]int64, 0)
	for i := 0; i < num; i++ {
		k := rander.Int63n(int64(maxnum))
		keys = append(keys, k)
		v := getRandomStr(7, 3)
		mapper[k] = v
	}
	fmt.Println("map:", mapper)
	fmt.Println("len(map):", len(mapper))
	fmt.Println()

	displayMapType(mapper)
	fmt.Println()

	// unsafe.Sizeof(map) is 8, indicates that 'map[k]v' is actually *hmap
	hp := *(**hmap)(unsafe.Pointer(&mapper))  // use **hmap to get *hmap
	fmt.Println("count:", hp.count)
	fmt.Println("flag:", hp.flags)
	fmt.Println("B:", hp.B)
	fmt.Println("noverflow:", hp.noverflow)
	fmt.Println("hash0:", hp.hash0)
	fmt.Println("bucket pointer:", uintptr(hp.buckets))
	fmt.Println("old bucket:", uintptr(hp.oldbuckets))
	fmt.Println("nevacuate:", hp.nevacuate)
	fmt.Println("extra:", uintptr(unsafe.Pointer(hp.extra)))
	fmt.Println()

	numBucket := uintptr(1 << hp.B)
	//mask := numBucket - 1
	_, _, bucketsize := getKeyValueBucketSize(mapper)

	// get key hash func
	//hasher := getKeyHashFunc(mapper)
	// todo show bucket
	for i:=uintptr(0); i<numBucket; i++ {
		fmt.Println("bucket:", i)
		b := (*bmap)(add(hp.buckets, i*bucketsize))
		for j:=uintptr(0); j<bucketsize/8; j++ {
			ip := (*int64)(add(unsafe.Pointer(b), j*8))
			fmt.Print(*ip, ", ")
		}
		fmt.Println()
	}
}

func getRandomStr(max, min int) string {
	length := int(rander.Int31n(int32(max-min))) + min
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = hans[rander.Int()%len(hans)]
	}
	return string(runes)
}

type maptype struct {
	typ        _type
	key        *_type
	elem       *_type
	bucket     *_type // internal type representing a hash bucket
	keysize    uint8  // size of key slot
	valuesize  uint8  // size of value slot
	bucketsize uint16 // size of bucket
	flags      uint32
}

func (mt *maptype) indirectkey() bool {   // store ptr to key instead of key itself
	return mt.flags&1 != 0
}
func (mt *maptype) indirectvalue() bool { // store ptr to value instead of value itself
	return mt.flags&2 != 0
}

func getKeyHashFunc(mapper interface{}) func(unsafe.Pointer, uintptr) uintptr {
	efacer := (*eface)(unsafe.Pointer(&mapper))
	mt := (*maptype)(unsafe.Pointer(efacer._type))   // *_type to *maptype
	return mt.key.alg.hash
}

func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

func getKeyValueBucketSize(mapper interface{}) (uintptr, uintptr, uintptr) {
	efacer := (*eface)(unsafe.Pointer(&mapper))
	typ := efacer._type
	mt := (*maptype)(unsafe.Pointer(typ))   // *_type to *maptype
	return mt.key.size, mt.elem.size, mt.bucket.size
}

func TestMapType(t *testing.T) {
	mapper := make(map[int64]string)
	var mapEface interface{} = mapper
	efacer := (*eface)(unsafe.Pointer(&mapEface))
	typ := efacer._type

	mt := (*maptype)(unsafe.Pointer(typ))   // *_type to *maptype
	fmt.Println("typ size      :", mt.typ.size)
	fmt.Println("typ ptrdata   :", mt.typ.ptrdata)
	fmt.Println("typ hash      :", mt.typ.hash)
	fmt.Println("typ align     :", mt.typ.align)
	fmt.Println("typ fieldalign:", mt.typ.fieldalign)
	fmt.Println("typ kind      :", mt.typ.kind, reflect.Kind(mt.typ.kind))
	fmt.Println("typ str       :", mt.typ.str)
	fmt.Println("typ ptrToThis :", mt.typ.ptrToThis)
	fmt.Println("indirectkey:", mt.indirectkey())
	fmt.Println("indirectvalue:", mt.indirectvalue())
	fmt.Println()

	// map[int64]string information
	fmt.Println("*key type:", uintptr(unsafe.Pointer(mt.key)))        // *rtype, map key type
	fmt.Println("*value type:", uintptr(unsafe.Pointer(mt.elem)))     // *rtype, map element type
	fmt.Println("*bucket typed:", uintptr(unsafe.Pointer(mt.bucket))) // *rtype, bucket structure
	fmt.Println("key size:", mt.keysize)
	fmt.Println("value size:", mt.valuesize)
	fmt.Println("bucket size:", mt.bucketsize)
	fmt.Println("flag:", mt.flags)
	fmt.Println()

	// int64 type
	fmt.Println("key type size      :", mt.key.size)        // 8 bytes
	fmt.Println("key type ptrdata   :", mt.key.ptrdata)
	fmt.Println("key type hash      :", mt.key.hash)
	fmt.Println("key type align     :", mt.key.align)
	fmt.Println("key type fieldalign:", mt.key.fieldalign)
	fmt.Println("key type kind      :", mt.key.kind, reflect.Kind(mt.key.kind))   // kind134
	fmt.Println("key type str       :", mt.key.str)
	fmt.Println("key type ptrToThis :", mt.key.ptrToThis)   // *int64 type
	fmt.Println()

	// string type
	fmt.Println("value type size      :", mt.elem.size)     // 16 bytes
	fmt.Println("value type ptrdata   :", mt.elem.ptrdata)
	fmt.Println("value type hash      :", mt.elem.hash)
	fmt.Println("value type align     :", mt.elem.align)
	fmt.Println("value type fieldalign:", mt.elem.fieldalign)
	fmt.Println("value type kind      :", mt.elem.kind, reflect.Kind(mt.elem.kind))  // string
	fmt.Println("value type str       :", mt.elem.str)
	fmt.Println("value type ptrToThis :", mt.elem.ptrToThis) // *string
	fmt.Println()

	// bucket type
	// 208 bytes: [8]uint8 tophash (8bytes), 8 int64 keys (8*8bytes),
	// 8 string values (8*16bytes), *overflow (8 bytes) = 8 + 64 + 128 + 8 =208 bytes
	fmt.Println("bucket type size      :", mt.bucket.size)
	fmt.Println("bucket type ptrdata   :", mt.bucket.ptrdata)
	fmt.Println("bucket type hash      :", mt.bucket.hash)
	fmt.Println("bucket type align     :", mt.bucket.align)
	fmt.Println("bucket type fieldalign:", mt.bucket.fieldalign)
	fmt.Println("bucket type kind      :", mt.bucket.kind, reflect.Kind(mt.bucket.kind))  // struct
	fmt.Println("bucket type str       :", mt.bucket.str)
	fmt.Println("bucket type ptrToThis :", mt.bucket.ptrToThis) // *bucket
	fmt.Println()
}

func displayMapType(mapper interface{}) {
	fmt.Println("reflect type str:", reflect.TypeOf(mapper).String())
	efacer := (*eface)(unsafe.Pointer(&mapper))
	typ := efacer._type
	mt := (*maptype)(unsafe.Pointer(typ))   // *_type to *maptype
	fmt.Println("key type size:", mt.key.size)
	fmt.Println("value type size:", mt.elem.size)
	fmt.Println("bucket type size:", mt.bucket.size)
	fmt.Println("indirectkey:", mt.indirectkey())
	fmt.Println("indirectvalue:", mt.indirectvalue())
}
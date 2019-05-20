package runtimer

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

const (
	// 8 key/value pairs a bucket can hold.
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits

	// load factor 6.5
	loadFactorNum = 13
	loadFactorDen = 2

	// Maximum key or value size (in bytes) to keep inline (instead of mallocing per element).
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
	emptyRest      = 0 // this cell is empty, and there are no more non-empty cells at higher indexes or overflows.
	emptyOne       = 1 // this cell is empty
	evacuatedX     = 2 // key/value is valid.  Entry has been evacuated to first half of larger table.
	evacuatedY     = 3 // same as above, but evacuated to second half of larger table.
	evacuatedEmpty = 4 // cell is empty, bucket is evacuated.
	minTopHash     = 5 // minimum tophash for a normal filled cell.

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

func (mt *maptype) indirectkey() bool { // store ptr to key instead of key itself
	return mt.flags&1 != 0
}
func (mt *maptype) indirectvalue() bool { // store ptr to value instead of value itself
	return mt.flags&2 != 0
}

func getKeyHashFunc(mapper interface{}) func(unsafe.Pointer, uintptr) uintptr {
	efacer := (*eface)(unsafe.Pointer(&mapper))
	mt := (*maptype)(unsafe.Pointer(efacer._type)) // *_type to *maptype
	return mt.key.alg.hash
}

func add(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

func getKeyValueBucketSize(mapper interface{}) (uintptr, uintptr, uintptr) {
	efacer := (*eface)(unsafe.Pointer(&mapper))
	typ := efacer._type
	mt := (*maptype)(unsafe.Pointer(typ)) // *_type to *maptype
	return mt.key.size, mt.elem.size, mt.bucket.size
}

func TestMapType(t *testing.T) {
	mapper := make(map[int64]string)
	var mapEface interface{} = mapper
	efacer := (*eface)(unsafe.Pointer(&mapEface))
	typ := efacer._type

	mt := (*maptype)(unsafe.Pointer(typ)) // *_type to *maptype
	fmt.Println("typ size      :", mt.typ.size)
	fmt.Println("typ kind      :", mt.typ.kind, reflect.Kind(mt.typ.kind))
	fmt.Println("typ ptrToThis :", mt.typ.ptrToThis)
	fmt.Println("indirectkey:", mt.indirectkey())
	fmt.Println("indirectvalue:", mt.indirectvalue())
	fmt.Println()

	// map[int64]string information
	fmt.Println("key size:", mt.keysize)
	fmt.Println("value size:", mt.valuesize)
	fmt.Println("bucket size:", mt.bucketsize)
	fmt.Println("flag:", mt.flags)
	fmt.Println()

	// int64 type
	fmt.Println("key type size      :", mt.key.size)                            // 8 bytes
	fmt.Println("key type kind      :", mt.key.kind, reflect.Kind(mt.key.kind)) // kind134
	fmt.Println("key type ptrToThis :", mt.key.ptrToThis)                       // *int64 type
	fmt.Println()

	// string type
	fmt.Println("value type size      :", mt.elem.size)                             // 16 bytes
	fmt.Println("value type kind      :", mt.elem.kind, reflect.Kind(mt.elem.kind)) // string
	fmt.Println("value type ptrToThis :", mt.elem.ptrToThis)                        // *string
	fmt.Println()

	// bucket type
	// 208 bytes: [8]uint8 tophash (8bytes), 8 int64 keys (8×8bytes),
	// 8 string values (8×16bytes), *overflow (8bytes) = 8 + 64 + 128 + 8 =208 bytes
	fmt.Println("bucket type size      :", mt.bucket.size)
	fmt.Println("bucket type kind      :", mt.bucket.kind, reflect.Kind(mt.bucket.kind)) // struct
	fmt.Println("bucket type ptrToThis :", mt.bucket.ptrToThis)                          // *bucket
	fmt.Println()
}

func displayMapType(mapper interface{}) {
	fmt.Println("reflect type str:", reflect.TypeOf(mapper).String())
	efacer := (*eface)(unsafe.Pointer(&mapper))
	typ := efacer._type
	/*
	// we can cast a *_type tp *maptype, which indicates:
	// for a map type in golang, not only common '_type' information is recorded in memory,
	// but also the key, value and buckets information.
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
	*/
	mt := (*maptype)(unsafe.Pointer(typ))
	fmt.Println("key type size:", mt.key.size)
	fmt.Println("value type size:", mt.elem.size)
	fmt.Println("bucket type size:", mt.bucket.size)
	fmt.Println("indirectkey:", mt.indirectkey())
	fmt.Println("indirectvalue:", mt.indirectvalue())
}

// load factor is 6.5, which means when hmap.count > 6.5×(1<<hmap.B), need to expand the map
// e.g., h.B=2, there are 4 buckets, when count > 6.5×4=26, double the buckets to 8
//       h.B=3, there are 8 buckets, when count > 6.5×8=52, double the buckets to 16
// expand buckets need to split oldBucket[i] to newBucket[i] and newBucket[i+oldsize].
// however split the buckets are not all done when expanded, but doing gradually when next add or delete.
// when expand, make a new array, and make the original bucket array as oldBuckets.
// when nex add or delete happens, we pick one or two buckets in the oldBuckets, and split it into new buckets.
// then finish splitting buckets, it is marked as 'evacuated'. the oldBuckets and new buckets are together
// responsible for read. when oldBuckets got empty, set it to zero.
func TestHashMapExpand(t *testing.T) {
	mapper := make(map[int64]string)
	displayMapType(mapper)
	fmt.Println()
	num := 60
	maxnum := 1000
	keys := make([]int64, 0)
	for i := 1; i <= num; i++ {
		key := rander.Int63n(int64(maxnum)) + 1
		keys = append(keys, key)
		value := getRandomStr(5, 3)
		mapper[key] = value
		if (i>=26 && i<=30) || (i>=52 && i<=58) {
			printBuckets(mapper)
		}
		if i==54 || i==55 {
			ki := rander.Int63n(int64(len(keys)))
			if value, ok := mapper[keys[ki]]; ok {
				fmt.Println("======key to delete:", keys[ki], "=", value)
				delete(mapper, keys[ki])
				printBuckets(mapper)
			}
		}
	}
}

func TestIndirectKeyValueSize(t *testing.T) {
	// the results indicate: the max direct kv size is 128 bytes,
	// when key or value size is larger than 128 bytes, use pointer instead
	mapper1 := make(map[[128]byte][128]byte)
	mapper2 := make(map[[129]byte][128]byte)
	mapper3 := make(map[[128]byte][129]byte)
	mapper4 := make(map[[129]byte][129]byte)
	displayMapType(mapper1)    // direct key, value
	fmt.Println()
	displayMapType(mapper2)    // indirect key, direct value
	fmt.Println()
	displayMapType(mapper3)    // direct key, indirect value
	fmt.Println()
	displayMapType(mapper4)    // indirect key, indirect value
}

// deleting a key would set the corresponding tophash to 1, indicates that this position is deleted and empty.
// when new key is add to this bucket, this position can hold the new kv pair.
func TestHashMapDelete(t *testing.T) {
	mapper := make(map[int64]string)
	displayMapType(mapper)
	fmt.Println()

	num := 50
	maxnum := 400
	keys := make([]int64, 0)
	for i := 0; i < num; i++ {
		k := rander.Int63n(int64(maxnum)) + 1
		keys = append(keys, k)
		v := getRandomStr(7, 3)
		mapper[k] = v
	}

	printBuckets(mapper)
	deleteNum := 10
	keysToDelete := make([]int64, 0, 10)
	for k:=0; k<deleteNum; k++ {
		ki := rander.Int63n(int64(len(keys)))
		if value, ok := mapper[keys[ki]]; ok {
			fmt.Println("======key to delete:", keys[ki], "=", value)
			keysToDelete = append(keysToDelete, keys[ki])
			delete(mapper, keys[ki])
			printBuckets(mapper)
		}
	}
	for _, key := range keysToDelete {
		value := getRandomStr(7, 3)
		fmt.Println("======put:", key, value)
		mapper[key] = value
		printBuckets(mapper)
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

func printBuckets(mapper map[int64]string) {
	hp := *(**hmap)(unsafe.Pointer(&mapper)) // use **hmap to get *hmap
	numBucket := uintptr(1 << hp.B)
	mask := numBucket - 1
	keysize, valuesize, bucketsize := getKeyValueBucketSize(mapper)
	hasher := getKeyHashFunc(mapper)
	fmt.Printf("====== B: %d, count: %d, noverflow: %d, nevacuate: %d\n",
		hp.B, hp.count, hp.noverflow, hp.nevacuate)
	for i := uintptr(0); i < numBucket; i++ {
		b := (*bmap)(add(hp.buckets, i*bucketsize))
		if *(*int64)(unsafe.Pointer(b)) == 0 {
			continue   // tophash are all empty, skip
		}
		fmt.Println("bucket:", i)
		for b != nil {
			tophash := *(*[bucketCnt]uint8)(unsafe.Pointer(b))
			fmt.Println("tophash:", tophash)
			for j := uintptr(0); j < bucketCnt; j++ {
				if tophash[j] == emptyRest {  // tophash[j]==0, rest positions are empty, break
					break
				}
				if tophash[j] == emptyOne {   // tophash[j]==1, this position are deleted, can be re-filled
					fmt.Print("[deleted], ")
					continue
				}
				curkeyp := (*int64)(add(unsafe.Pointer(b), bucketCnt+j*keysize))
				curvaluep := add(unsafe.Pointer(b), bucketCnt+bucketCnt*keysize+j*valuesize)
				khash := hasher(unsafe.Pointer(curkeyp), uintptr(hp.hash0))
				curtophash := khash >> (64 - bucketCnt)
				curindex := khash & mask
				fmt.Print("[", curtophash, "|", curindex, "|", *curkeyp, "=", *(*string)(curvaluep), "], ")
			}
			fmt.Println()
			b = *(**bmap)(add(unsafe.Pointer(b), bucketsize-bucketCnt)) // get next bucket
		}
	}
	fmt.Println()
	if uintptr(hp.oldbuckets)>0 {
		numBucket = numBucket/2
		mask = numBucket - 1
		for i := uintptr(0); i < numBucket; i++ {
			b := (*bmap)(add(hp.oldbuckets, i*bucketsize))
			if *(*int64)(unsafe.Pointer(b)) == 0 {
				continue   // tophash are all empty, skip
			}
			fmt.Println("old bucket:", i)
			for b != nil {
				tophash := *(*[bucketCnt]uint8)(unsafe.Pointer(b))
				fmt.Println("tophash:", tophash)
				for j := uintptr(0); j < bucketCnt; j++ {
					if tophash[j] == emptyRest {   // tophash[j]==0, rest positions are empty, break
						break
					}
					if tophash[j] == emptyOne {    // tophash[j]==1, this position are deleted, can be re-filled
						fmt.Print("[deleted], ")
						continue
					}
					if tophash[j] == evacuatedX {  // tophash[j]==2, evacuated to lower half of new buckets
						fmt.Print("[lower], ")
						continue
					}
					if tophash[j] == evacuatedY {  // tophash[j]==3, evacuated to higher half of new buckets
						fmt.Print("[higher], ")
						continue
					}
					if tophash[j] == evacuatedEmpty {  // tophash[j]==4, evacuation finished
						fmt.Print("[evacuated], ")
						continue
					}
					curkeyp := (*int64)(add(unsafe.Pointer(b), bucketCnt+j*keysize))
					curvaluep := add(unsafe.Pointer(b), bucketCnt+bucketCnt*keysize+j*valuesize)
					khash := hasher(unsafe.Pointer(curkeyp), uintptr(hp.hash0))
					curtophash := khash >> (64 - bucketCnt)
					curindex := khash & mask
					fmt.Print("[", curtophash, "|", curindex, "|", *curkeyp, "=", *(*string)(curvaluep), "], ")
				}
				fmt.Println()
				b = *(**bmap)(add(unsafe.Pointer(b), bucketsize-bucketCnt)) // get next bucket
			}
		}
	}
	fmt.Println()
}

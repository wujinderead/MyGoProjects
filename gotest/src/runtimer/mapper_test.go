package runtimer

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
	"unsafe"
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
	mask := 200
	for i := 0; i < num; i++ {
		mapper[rander.Int63n(int64(mask))] = getRandomStr(7, 3)
	}
	fmt.Println("map:", mapper)
	fmt.Println("len(map):", len(mapper))
	hp := (*hmap)(unsafe.Pointer(&mapper))
	fmt.Println("count:", hp.count)
	fmt.Println("flag:", hp.flags)
	fmt.Println("B:", hp.B)
	fmt.Println("noverflow:", hp.noverflow)
	fmt.Println("hash0:", hp.hash0)
	fmt.Println("bucket pointer:", uintptr(hp.buckets))
	fmt.Println("old bucket:", uintptr(hp.oldbuckets))
	fmt.Println("nevacuate:", hp.nevacuate)
	fmt.Println("extra:", uintptr(unsafe.Pointer(hp.extra)))
}

func getRandomStr(max, min int) string {
	length := int(rander.Int31n(int32(max-min))) + min
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = hans[rander.Int()%len(hans)]
	}
	return string(runes)
}

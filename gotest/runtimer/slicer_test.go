package runtimer

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

// slice in runtime/slice.go
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

// fixed size array has no slice header, so where are the length and pointer stored?
// fixed size array is determined in compiling, before runtime.
// so it's likely that the length and pointer are stored in something like a variable table
func TestFixedSizeArray(t *testing.T) {
	a := []int32{1, 2, 3}
	b := [...]int64{1, 2, 3, 4, 5, 6} // fixed-size array not specify length
	bp := &b
	c := [...]complex128{complex(float64(4.1), float64(4.1)), complex(float64(4.1), float64(4.1)),
		complex(float64(4.2), float64(4.2)), complex(float64(4.1), float64(4.1)),
		complex(float64(4.9), float64(4.9))} // 5 complex128

	_ = bp[1] // pointer of fixed size array can be indexed directly

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b),
		reflect.TypeOf(bp), reflect.TypeOf(c)) // []int, [3]int, *[3]int, [4]complex128

	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c)) // 24=3*8, 48=6*8, 80=5*16
	// 3 int32 should be 12 bytes, but unsafe.Sizeof(a) is 24, so it has been aligned

	// truncating fixed-size array makes a new slice header
	// the underlying array of the slice is the fixed-size array directly
	d := b[2:4]
	fmt.Println("d:", d, ", type:", reflect.TypeOf(d))
	slip := (*slice)(unsafe.Pointer(&d))
	fmt.Println("d slice header:", int(uintptr(slip.array)), slip.len, slip.cap)
	fmt.Println("d[0] addr:", uintptr(unsafe.Pointer(&d[0])),
		", b[2] addr:", uintptr(unsafe.Pointer(&b[2]))) // the same
	b[2], b[3] = 123456, 654321 // b change, d also change
	fmt.Println("d:", d)

	e := [...]int64{2: 3, 4: 5, 0: 1} // fixed-size array can specify index
	fmt.Println(e)                    // [1 0 3 0 5]
}

func TestSliceAppendReference(t *testing.T) {
	// slice append use the same underlying array, it may infer other slices that reference this underlying array
	/*
		slicer := make([]int, 3)        // slicer [0, 0, 0]
		second := append(slicer, 3)     // second [0, 0, 0, 3]
		slice[1] = 1                    // second also change: slicer [0, 1, 0], second [0, 1, 0, 3]
		third := append(slicer, 4, 5)   // second also change: slicer [0, 1, 0], second [0, 1, 0, 4], third [0, 1, 0, 4, 5]
		fourth := append(slicer, 7, 8, 9, 10)   // second, third no change, fourth [0 1 0 7 8 9 10], len 7 cap 12.
												// this is because that fourth use a new underlying slice.
			                                    // slice capacity is 6, fourth length will be 7 after append,
			                                    // so double the capacity and make a new slice to contain fourth
		trunc := fourth[2:5]    // trunc [0 7 8], len 3, cap 10
		                        // trunc and fourth use the same underlying array
		                        // since truncate start from fourth[2], trunc cap reduce 2,
		                        // trunc pointer is point to fourth[2]
	*/
	getSliceHeader := func(sli []int) []int {
		slip := (*slice)(unsafe.Pointer(&sli))
		return []int{int(uintptr(slip.array)), slip.len, slip.cap}
	}
	slicer := make([]int, 3, 6)
	fmt.Println("slicer[0] addr:", uintptr(unsafe.Pointer(&slicer[0])))
	fmt.Println("slicer head:", getSliceHeader(slicer))
	fmt.Println("slicer: ", slicer)
	fmt.Println()

	second := append(slicer, 3)
	fmt.Println("slicer head:", getSliceHeader(slicer))
	fmt.Println("slicer: ", slicer)
	fmt.Println("second head:", getSliceHeader(second))
	fmt.Println("second: ", second)
	fmt.Println()

	slicer[1] = 1
	fmt.Println("slicer head:", getSliceHeader(slicer))
	fmt.Println("slicer:", slicer)
	fmt.Println("second head:", getSliceHeader(second))
	fmt.Println("second:", second)
	fmt.Println()

	third := append(slicer, 4, 5)
	fmt.Println("slicer head:", getSliceHeader(slicer))
	fmt.Println("slicer: ", slicer)
	fmt.Println("second head:", getSliceHeader(second))
	fmt.Println("second:", second)
	fmt.Println("third head:", getSliceHeader(third))
	fmt.Println("third:", third)
	fmt.Println()

	fourth := append(slicer, 7, 8, 9, 10)
	fmt.Println("slicer head:", getSliceHeader(slicer)) // slicer head: [824633819856 3 6]
	fmt.Println("slicer: ", slicer)
	fmt.Println("second head:", getSliceHeader(second)) // second head: [824633819856 4 6]
	fmt.Println("second:", second)
	fmt.Println("third head:", getSliceHeader(third)) // third head: [824633819856 5 6]
	fmt.Println("third:", third)
	fmt.Println("fourth[0] addr:", uintptr(unsafe.Pointer(&fourth[0]))) // 824633836480
	fmt.Println("fourth head:", getSliceHeader(fourth))                 // fourth head: [824633836480 7 12], a new underlying slice
	fmt.Println("fourth:", fourth)
	fmt.Println()

	// truncate slice
	trunc := fourth[2:5] // truncate from fourth[2], so trunc cap reduce 2, trunc pointer is fourth[2]
	fmt.Println("trunc:", trunc)
	fmt.Println("trunc header:", getSliceHeader(trunc))                 // [824633836480 3 10]
	fmt.Println("trunc[0] addr:", uintptr(unsafe.Pointer(&trunc[0])))   // 824633836496
	fmt.Println("fourth[2] addr:", uintptr(unsafe.Pointer(&fourth[2]))) // 824633836496
}

func TestPointerIndexable(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}
	p := &a[0]                           // a pointer of an underlying array
	b := (*[4000]int)(unsafe.Pointer(p)) // the most convenient way to convert a pointer to something indexable
	fmt.Println(b[3])                    // that's ok
	// that's legal, but we should avoid, because the value has no meaning,
	// and it's dangerous to access arbitrary address during runtime.
	fmt.Println(b[3500])

	// pointer to array to slice
	bb := (*[4000]int)(unsafe.Pointer(p))[:]
	fmt.Println(bb[3], bb[3500], len(bb), cap(bb))

	_ = (*[]int)(unsafe.Pointer(p)) // that's not the right way to make a pointer indexable

	// make a slice header
	h := &struct {
		p   unsafe.Pointer
		len int
		cap int
	}{unsafe.Pointer(p), 6, 6}
	c := *(*[]int)(unsafe.Pointer(h))
	fmt.Println(c[0], c[5])
}

func TestDoubleColon(t *testing.T) {
	// double colon when truncating a slice is to specify the cap
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	b := a[3:6:8]                  // a[3,4,5] is elements of b, a[6,7] is as empty slots of b
	fmt.Println(b, len(b), cap(b)) // len 3, cap 5
	b = append(b, -1)              // a[6] change to -1, because a, b share the same underlying array
	fmt.Println(a, b)
	b = append(b, -2) // a[7] change to -2, because a, b share the same underlying array
	fmt.Println(a, b)
	b = append(b, -3) // a doesn't change, because b is out of capacity, and copied to a new underlying array
	fmt.Println(a, b)
}

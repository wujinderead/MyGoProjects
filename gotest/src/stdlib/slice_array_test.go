package stdlib

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"unsafe"
)

func TestAppend(t *testing.T) {
	x := []string{"start"}

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		y := append(x, "hello", "world")
		t.Log(cap(y), len(y))
	}()
	go func() {
		defer wg.Done()
		z := append(x, "goodbye", "bob")
		t.Log(cap(z), len(z))
	}()
	wg.Wait()

}

func TestAppend1(t *testing.T) {
	x := make([]string, 0, 6)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		y := append(x, "hello", "world")
		t.Log(len(y))
	}()
	go func() {
		defer wg.Done()
		z := append(x, "goodbye", "bob")
		t.Log(len(z))
	}()
	wg.Wait()
}

func TestFixedSizeArray(t *testing.T) {
	a := []int{1, 2, 3}
	b := [3]int{1, 2, 3}
	bp := &b
	c := [...]int{4, 5, 6} // fixed-size array not specify length
	_ = bp[1]              // pointer of fixed size array can be indexed directly
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(bp), reflect.TypeOf(c))
	// fixed size array has no header
}

func TestSliceAppendReference(t *testing.T) {
	// slice append use the same underlying array, if make infer other slice that references this underlying array
	/*
		slice := make([]int, 3)         // slice [0, 0, 0]
		second := append(slicer, 3)     // second [0, 0, 0, 3]
		slice[1] = 1                    // second also change: slice [0, 1, 0], second [0, 1, 0, 3]
		third := append(slicer, 4, 5)   // second also change: slice [0, 1, 0], second [0, 1, 0, 4], third [0, 1, 0, 4, 5]
		fourth := append(slicer, 7, 8, 9, 10)   // second, third no change, fourth [0 1 0 7 8 9 10]
		                                        // slice capacity is 6, fourth length will be 7 after append,
		                                        // so double the capacity and make a new slice to contain fourth
	*/

	// slice header has 3 ints, header[0] is the slice address, i.e. address of slice[0]
	// header[1] is slice length, header[2] is slice capacity
	getSliceHeader := func(sli []int) []int64 {
		p := &sli
		size := unsafe.Sizeof(sli)
		up := uintptr(unsafe.Pointer(p))
		ret := make([]int64, size/8)
		for i := 0; i < int(size/8); i++ {
			ret[i] = *(*int64)(unsafe.Pointer(up + uintptr(i<<3)))
		}
		return ret
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
}

package runtimer

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestEmptyStruct(t *testing.T) {
	fmt.Println(reflect.TypeOf(struct{}{}))
	fmt.Println(reflect.TypeOf(struct{}{}).Size())
	slicer := []struct{}{{}, {}, {}}
	inter := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(reflect.TypeOf(slicer), reflect.TypeOf(slicer).Size())
	fmt.Println(reflect.TypeOf(inter), reflect.TypeOf(inter).Size())
	fmt.Println(&slicer[0]) // &{}
	fmt.Println(&slicer[1])
	fmt.Println(&slicer[2])
	fmt.Println(&inter[0])
	fmt.Println(&inter[1])
	fmt.Println(&inter[2])
	fmt.Println(unsafe.Pointer(&slicer[0])) // 0x1213a70, all the same
	fmt.Println(unsafe.Pointer(&slicer[1])) // struct{}{} is in the constant pool
	fmt.Println(unsafe.Pointer(&slicer[2]))
	fmt.Println(unsafe.Pointer(&struct{}{}))
	fmt.Println(unsafe.Pointer(&struct{}{}))

	sliceStruct := (*slice)(unsafe.Pointer(&slicer))
	fmt.Println(sliceStruct.len)
	fmt.Println(sliceStruct.cap)
	fmt.Println(sliceStruct.array)
}

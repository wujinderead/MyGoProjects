package main

import (
	"fmt"
	"unsafe"
)

func main() {
	type Int struct {
		x int
	}
	inter := &Int{1}
	fmt.Println(uintptr(unsafe.Pointer(inter)))
	fmt.Println(uintptr(unsafe.Pointer(&inter.x)))
}
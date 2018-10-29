package main

import "fmt"

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	var x *int = nil
	if x == nil {
		fmt.Println("empty interface")
	}
	Foo(x)
	/*
		x is nil. however when apply to interface, the interface.data is nil, but interface is not nil.
	*/
}

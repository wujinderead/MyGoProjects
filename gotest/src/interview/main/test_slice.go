package main

import "fmt"

func main() {
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)
	/*
		list := new([]int)
		list = append(list, 1)
		cannot compile, new(X) return *X, which cannot apply to []Type (slice).
	*/
}

package main

import (
	"data_structure/heap"
	"fmt"
)

var test1 = []int{
	69, 59, 71, 39, 58, 12, 5, 29, 50, 78, 20, 67, 92, 52, 19, 54, 62,
	43, 1, 83, 45, 31, 23, 18, 46, 79, 47, 55, 14, 98, 49, 27, 81, 21,
	15, 63, 7, 60, 22, 87, 80, 70, 86, 41, 8, 84, 99, 76, 51, 93, 26,
	42, 25, 72, 40, 91, 36, 17, 77, 34, 90, 3, 9, 6, 44, 94, 33, 65,
	24, 73, 16, 37, 64, 2, 89, 53, 4, 11, 66, 61, 75, 74, 85, 10, 38,
	82, 95, 96, 30, 28, 32, 88, 13, 57, 35, 48, 68, 56, 97}

var test2 = []int{6, 1, 3, 5, 7, 2, 4}

func main() {
	//testBatch(test1)
	testBatch(test2)
	//testSingle(test1)
	testSingle(test2)
	heap.HeapAscendSort([]int{6, 1, 15, 2, 8, 18, 0, 13, 16, 7, 11,
		5, 4, 12, 19, 17, 14, 9, 10, 3})
}

func testBatch(test []int) {
	h := heap.NewHeapFromArray(test)
	fmt.Println(h)

	h.Add(0)
	fmt.Println(h)

	fmt.Println(h.Poll())
	fmt.Println(h)
	fmt.Println(h.Poll())
	fmt.Println(h)
	fmt.Println(h.Poll())
	fmt.Println(h)

	h.Add(1)
	fmt.Println(h)
	h.Add(0)
	fmt.Println(h)
	h.Add(2)
	fmt.Println(h)

	fmt.Println(h.Poll())
	fmt.Println(h)
	fmt.Println(h.Poll())
	fmt.Println(h)
	fmt.Println(h.Poll())
	fmt.Println(h)
}

func testSingle(test []int) {
	h := heap.NewHeap()
	for _, v := range test {
		h.Add(v)
		fmt.Println(h)
	}

	h.Add(0)
	fmt.Println(h)

	fmt.Println(h.Poll())
	fmt.Println(h)
	fmt.Println(h.Poll())
	fmt.Println(h)
	fmt.Println(h.Poll())
	fmt.Println(h)

	h.Add(1)
	fmt.Println(h)
	h.Add(0)
	fmt.Println(h)
	h.Add(2)
	fmt.Println(h)

	fmt.Println(h.Poll())
	fmt.Println(h)
	fmt.Println(h.Poll())
	fmt.Println(h)
	fmt.Println(h.Poll())
	fmt.Println(h)
}

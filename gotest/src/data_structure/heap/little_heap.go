package heap

import (
	"errors"
	"fmt"
)

type Heap struct {
	size  int
	array []int
}

func NewHeap() *Heap {
	heap := &Heap{}
	heap.array = make([]int, 0)
	heap.size = 0
	return heap
}

func NewHeapFromArray(ints []int) *Heap {
	heap := &Heap{}
	heap.array = make([]int, len(ints))
	copy(heap.array, ints)
	heap.size = len(ints)
	heap.heapify()
	return heap
}

func (heap *Heap) heapify() {
	for i := heap.size/2 - 1; i >= 0; i-- {
		heap.siftDown(i, heap.array[i])
	}
}

func (heap *Heap) siftUp(i, value int) {
	for i > 0 {
		parent := (i - 1) / 2
		if heap.array[parent] <= value {
			break
		}
		heap.array[i] = heap.array[parent]
		i = parent
	}
	heap.array[i] = value
}

func (heap *Heap) siftDown(i, value int) {
	half := heap.size / 2
	for i < half {
		left := i*2 + 1   // left must exists
		right := left + 1 // right not certain
		if right < heap.size && heap.array[right] < heap.array[left] {
			left = right
		}
		if value <= heap.array[left] {
			break
		}
		heap.array[i] = heap.array[left]
		i = left
	}
	heap.array[i] = value
	//fmt.Println(heap.array)
}

func (heap *Heap) String() string {
	return fmt.Sprint(heap.array)
}

func (heap *Heap) Add(ele int) {
	heap.array = append(heap.array, ele)
	heap.size++
	heap.siftUp(heap.size-1, ele)
}

func (heap *Heap) Peek() (int, error) {
	if heap.size == 0 {
		return 0, errors.New("no elements in queue")
	}
	return heap.array[0], nil
}

func (heap *Heap) Poll() (int, error) {
	if heap.size == 0 {
		return 0, errors.New("no elements in queue")
	}
	re := heap.array[0]
	last := heap.array[heap.size-1]
	heap.size--
	heap.array = heap.array[:heap.size]
	heap.siftDown(0, last)
	return re, nil
}

func HeapAscendSort(array []int) []int {
	if array == nil || len(array) == 0 {
		return array
	}
	siftDown := func(arr []int, size, i, value int) {
		for i < size/2 {
			left := 2*i + 1
			right := left + 1
			if right < size && arr[right] > arr[left] {
				left = right
			}
			if value >= array[left] {
				break
			}
			arr[i] = arr[left]
			i = left
		}
		arr[i] = value
	}
	// make a big end heap
	for i := len(array)/2 - 1; i >= 0; i-- {
		siftDown(array, len(array), i, array[i])
	}
	fmt.Println(array)
	// heap sort
	size := len(array)
	for size > 1 {
		array[0], array[size-1] = array[size-1], array[0]
		size--
		siftDown(array, size, 0, array[0])
		fmt.Println(array)
	}
	return array
}

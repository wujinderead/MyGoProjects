package main

import (
	"fmt"
	"sort"
	"math/rand"
	"container/heap"
)

// https://leetcode.com/problems/find-median-from-data-stream/

// Median is the middle value in an ordered integer list. If the size of the list
// is even, there is no middle value. So the median is the mean of the two middle value.
// For example, [2,3,4], the median is 3. [2,3], the median is (2 + 3) / 2 = 2.5.
// Design a data structure that supports the following two operations:
//   void addNum(int num) - Add a integer number from the data stream to the data structure.
//   double findMedian() - Return the median of all elements so far.
// Example:
//   addNum(1)
//   addNum(2)
//   findMedian() -> 1.5
//   addNum(3)
//   findMedian() -> 2
// Follow up:
//   If all integer numbers from the stream are between 0 and 100, how would you optimize it?
//   If 99% of all integer numbers from the stream are between 0 and 100, how would you optimize it?


// we need a max heap for the left part, and a min heap for the right part
type MedianFinder struct {
    maxheap *maxheap
    minheap *minheap
}

func Constructor() MedianFinder {
    return MedianFinder{
		maxheap: &maxheap{intheap{sort.IntSlice(make([]int, 0))}},
		minheap: &minheap{intheap{sort.IntSlice(make([]int, 0))}},
	}
}

func (this *MedianFinder) AddNum(num int)  {
    if this.maxheap.Len()==0 {
    	heap.Push(this.maxheap, num)
    	return
	}
	if num <= this.maxheap.IntSlice[0] {
		heap.Push(this.maxheap, num)
		if this.maxheap.Len()-this.minheap.Len()==2 {
			x := heap.Pop(this.maxheap).(int)
			heap.Push(this.minheap, x)
		}
	} else {
		heap.Push(this.minheap, num)
		if this.minheap.Len()-this.maxheap.Len()==2 {
			x := heap.Pop(this.minheap).(int)
			heap.Push(this.maxheap, x)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
    if this.minheap.Len()+this.maxheap.Len() == 0 {
    	return 0
	}
	if this.maxheap.Len() > this.minheap.Len() {
		return float64(this.maxheap.IntSlice[0])
	} else if this.maxheap.Len() < this.minheap.Len() {
		return float64(this.minheap.IntSlice[0])
	} else {
		return float64(this.maxheap.IntSlice[0]+this.minheap.IntSlice[0])/2
	}
}

type intheap struct {
	sort.IntSlice
}

func (h *intheap) Push(x interface{}) {
	h.IntSlice = append([]int(h.IntSlice), x.(int))
}

func (h *intheap) Pop() interface{} {
	x := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return x
}

type maxheap struct {
	intheap
}

type minheap struct {
	intheap
}

func (h maxheap) Less(i, j int) bool {
	return h.intheap.IntSlice[i] > h.intheap.IntSlice[j]
}

func main() {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	fmt.Println(m.FindMedian())
	m.AddNum(3)
	fmt.Println(m.FindMedian())
	m = Constructor()
	arr := rand.Perm(20)
	fmt.Println(arr)
	for i := range arr {
		m.AddNum(arr[i])
		mid := m.FindMedian()
		a := make([]int, i+1)
		copy(a, arr[:i+1])
		sort.Sort(sort.IntSlice(a))
		var median float64
		if len(a)%2 == 1 {
			median = float64(a[len(a)/2])
		} else {
			median = float64(a[len(a)/2-1]+a[len(a)/2])/2
		}
		fmt.Println(mid, median, a)
	}
}
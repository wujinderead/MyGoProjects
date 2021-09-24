package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/find-original-array-from-doubled-array/

// An integer array original is transformed into a doubled array changed by appending twice the value
// of every element in original, and then randomly shuffling the resulting array.
// Given an array changed, return original if changed is a doubled array. If changed is not a doubled
// array, return an empty array. The elements in original may be returned in any order.
// Example 1:
//   Input: changed = [1,3,4,2,6,8]
//   Output: [1,3,4]
//   Explanation: One possible original array could be [1,3,4]:
//     - Twice the value of 1 is 1 * 2 = 2.
//     - Twice the value of 3 is 3 * 2 = 6.
//     - Twice the value of 4 is 4 * 2 = 8.
//     Other original arrays could be [4,3,1] or [3,1,4].
// Example 2:
//   Input: changed = [6,3,0,1]
//   Output: []
//   Explanation: changed is not a doubled array.
// Example 3:
//   Input: changed = [1]
//   Output: []
//   Explanation: changed is not a doubled array.
// Constraints:
//   1 <= changed.length <= 10^5
//   0 <= changed[i] <= 10^5

// make a max-heap, each time, pop max_number, and also remove the max_number/2 from the heap
func findOriginalArray(changed []int) []int {
	if len(changed)%2 == 1 {
		return []int{}
	}
	count := make(map[int]int)
	for _, v := range changed {
		count[v] = count[v] + 1
	}
	hh := h(changed)
	heap.Init(&hh)
	original := make([]int, 0, len(changed)/2)
	for hh.Len() > 0 {
		max := heap.Pop(&hh).(int) // pop max value in heap
		if count[max] == 0 {
			continue
		}
		count[max] = count[max] - 1
		if max%2 != 0 || count[max/2] == 0 { // check if max/2 present, if present, delete it
			return []int{}
		}
		original = append(original, max/2)
		count[max/2] = count[max/2] - 1
	}
	return original
}

type h []int

func (h h) Len() int {
	return len(h)
}

func (h h) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h h) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h *h) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *h) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func main() {
	for _, v := range []struct {
		c   []int
		ans []int
	}{
		{[]int{1, 3, 4, 2, 6, 8}, []int{1, 3, 4}},
		{[]int{6, 3, 0, 1}, []int{}},
		{[]int{1}, []int{}},
		{[]int{1, 2, 5, 10, 10, 20}, []int{1, 5, 10}},
		{[]int{0, 3, 2, 4, 6, 0}, []int{0, 2, 3}},
	} {
		fmt.Println(findOriginalArray(v.c), v.ans)
	}
}

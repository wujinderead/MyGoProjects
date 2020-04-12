package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/construct-target-array-with-multiple-sums/

// Given an array of integers target. From a starting array, A consisting of all 1's,
// you may perform the following procedure:
//   let x be the sum of all elements currently in your array.
//   choose index i, such that 0 <= i < target.size and set the value of A at index i to x.
// You may repeat this procedure as many times as needed.
// Return True if it is possible to construct the target array from A otherwise return False.
// Example 1:
//   Input: target = [9,3,5]
//   Output: true
//   Explanation: Start with [1, 1, 1]
//     [1, 1, 1], sum = 3 choose index 1
//     [1, 3, 1], sum = 5 choose index 2
//     [1, 3, 5], sum = 9 choose index 0
//     [9, 3, 5] Done
// Example 2:
//   Input: target = [1,1,1,2]
//   Output: false
//   Explanation: Impossible to create target array from [1,1,1,1].
// Example 3:
//  Input: target = [8,5]
//  Output: true
// Constraints:
//   N == target.length
//   1 <= target.length <= 5 * 10^4
//   1 <= target[i] <= 10^9

// for numbers in target, the maximal number is got from the sum of previous target.
// so use a heap to track the maximal value
func isPossible(target []int) bool {
	if len(target) == 1 && target[0] > 1 {
		return false
	}
	allsum := 0
	h := heaper([]int{})
	for i := range target {
		allsum += target[i]
		heap.Push(&h, target[i])
	}
	for h[0] > 1 { // the heap head is 1, means the heap is all 1, we can stop
		max := heap.Pop(&h).(int)
		remain := allsum - max
		if max <= remain || (max%remain == 0 && remain != 1) {
			return false
		}
		newval := max - (max/remain)*remain
		allsum -= max - newval
		heap.Push(&h, newval)
	}
	return true
}

func main() {
	fmt.Println(isPossible([]int{9, 3, 5}))
	fmt.Println(isPossible([]int{9, 17, 5}))
	fmt.Println(isPossible([]int{1, 1, 1, 2}))
	fmt.Println(isPossible([]int{8, 5}))
	fmt.Println(isPossible([]int{5, 1, 9, 1, 49}))
	fmt.Println(isPossible([]int{2, 900000002}))
	fmt.Println(isPossible([]int{1}))
	fmt.Println(isPossible([]int{2}))
	fmt.Println(isPossible([]int{1, 2}))
	fmt.Println(isPossible([]int{1, 1, 2}))
	fmt.Println(isPossible([]int{1, 1000000}))
}

type heaper []int

func (a heaper) Len() int {
	return len(a)
}

func (a heaper) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a heaper) Less(i, j int) bool { // we need max heap
	return a[i] > a[j]
}

func (a *heaper) Push(x interface{}) {
	*a = append(*a, x.(int))
}

func (a *heaper) Pop() interface{} {
	x := (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	return x
}

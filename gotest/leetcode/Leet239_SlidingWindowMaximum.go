package main

import (
	"fmt"
	"container/list"
)

// https://leetcode.com/problems/sliding-window-maximum/

// Given an array nums, there is a sliding window of size k which is moving from
// the very left of the array to the very right. You can only see the k numbers in
// the window. Each time the sliding window moves right by one position. Return the
// max sliding window.
// Follow up: Could you solve it in linear time?
// Example:
//   Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
//   Output: [3,3,5,5,6,7]
//   Explanation:
//     Window position                Max
//     ---------------               -----
//     [1  3  -1] -3  5  3  6  7       3
//      1 [3  -1  -3] 5  3  6  7       3
//      1  3 [-1  -3  5] 3  6  7       5
//      1  3  -1 [-3  5  3] 6  7       5
//      1  3  -1  -3 [5  3  6] 7       6
//      1  3  -1  -3  5 [3  6  7]      7
// Constraints:
//   1 <= nums.length <= 10^5
//   -10^4 <= nums[i] <= 10^4
//   1 <= k <= nums.length

// monotonic queue method
func maxSlidingWindow(nums []int, k int) []int {
    l := list.New()
    result := make([]int, 0, len(nums)-k+1)
    for i:=0; i<k-1; i++ {
		//fmt.Println("add", nums[i])
    	pushMonoQueue(l, nums[i])
    	//printlist(l)
    	//fmt.Println()
	}
	for i:=k-1; i<len(nums); i++ {
		//fmt.Println("add", nums[i])
		pushMonoQueue(l, nums[i])
		//printlist(l)
		result = append(result, getMaxMonoQueue(l))
		//fmt.Println("pop", nums[i-k+1])
		popMonoQueue(l)
		//printlist(l)
		//fmt.Println()
	}
	return result
}

// the element of monotonic queue is a pair
//   first is the value
//   second is how many elements are deleted before current element will be deleted.
func pushMonoQueue(l *list.List, val int) {  // push a value to the queue
	count := 0
	for l.Len()>0 && l.Back().Value.([2]int)[0]<val {
		count += l.Back().Value.([2]int)[1]+1
		l.Remove(l.Back())
	}
	l.PushBack([2]int{val, count})
}

func getMaxMonoQueue(l *list.List) int {    // get current max for current window
	return l.Front().Value.([2]int)[0]
}

// pop a value, may not really delete from the queue, may just decease the size
func popMonoQueue(l *list.List) {
	pair := l.Front().Value.([2]int)
	if pair[1] > 0 {      // if front.count>0, virtually pop a value
		l.Front().Value = [2]int{pair[0], pair[1]-1}
		return
	}
	l.Remove(l.Front())   // real remove front
	return
}

func printlist(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.([2]int), " ")
	}
	fmt.Println()
}

func maxSlidingWindowDeque(nums []int, k int) []int {
	l := list.New()
	result := make([]int, 0, len(nums)-k+1)
	for i:=0; i<len(nums); i++ {
		if l.Len()>0 && i-k==l.Front().Value.(int) {  // i-k==l.front, pop front
			l.Remove(l.Front())
		}
		for l.Len()>0 && nums[l.Back().Value.(int)]<nums[i] {
			l.Remove(l.Back())
		}
		l.PushBack(i)
		if i>=k-1 {
			result = append(result, nums[l.Front().Value.(int)])
		}
	}
	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
	fmt.Println(maxSlidingWindowDeque([]int{1,3,-1,-3,5,3,6,7}, 3))
	//fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 1))
	//fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 2))
	fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 6))
	fmt.Println(maxSlidingWindowDeque([]int{1,3,-1,-3,5,3,6,7}, 6))
	//fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 7))
	//fmt.Println(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 8))
}
package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
// Given an array of integers nums and an integer limit, return the size of the longest
// continuous subarray such that the absolute difference between any two elements is less
// than or equal to limit.
// In case there is no subarray satisfying the given condition return 0.
// Example 1:
//   Input: nums = [8,2,4,7], limit = 4
//   Output: 2
//   Explanation: All subarrays are:
//     [8] with maximum absolute diff |8-8| = 0 <= 4.
//     [8,2] with maximum absolute diff |8-2| = 6 > 4.
//     [8,2,4] with maximum absolute diff |8-2| = 6 > 4.
//     [8,2,4,7] with maximum absolute diff |8-2| = 6 > 4.
//     [2] with maximum absolute diff |2-2| = 0 <= 4.
//     [2,4] with maximum absolute diff |2-4| = 2 <= 4.
//     [2,4,7] with maximum absolute diff |2-7| = 5 > 4.
//     [4] with maximum absolute diff |4-4| = 0 <= 4.
//     [4,7] with maximum absolute diff |4-7| = 3 <= 4.
//     [7] with maximum absolute diff |7-7| = 0 <= 4.
//     Therefore, the size of the longest subarray is 2.
// Example 2:
//   Input: nums = [10,1,2,4,7,2], limit = 5
//   Output: 4
//   Explanation: The subarray [2,4,7,2] is the longest since the maximum absolute
//     diff is |2-7| = 5 <= 5.
// Example 3:
//   Input: nums = [4,2,2,2,4,4,2,2], limit = 0
//   Output: 3
// Constraints:
//   1 <= nums.length <= 10^5
//   1 <= nums[i] <= 10^9
//   0 <= limit <= 10^9

// use deque, O(N); other method, use two heap or treemap, O(NlogN)
func longestSubarray(nums []int, limit int) int {
    // use deque to keep track of the max value and min value of a subarray
	allmax := 0
	left, right := 0, 0
	minq, maxq := list.New(), list.New()
	for right<len(nums) {
		// push to max deque
		for maxq.Len()>0 && nums[maxq.Back().Value.(int)] < nums[right] {
			maxq.Remove(maxq.Back())
		}
		maxq.PushBack(right)
		// push to min deque
		for minq.Len()>0 && nums[minq.Back().Value.(int)] > nums[right] {
			minq.Remove(minq.Back())
		}
		minq.PushBack(right)
		right++
		// get current max min
		curmax := nums[maxq.Front().Value.(int)]
		curmin := nums[minq.Front().Value.(int)]
		if curmax-curmin <= limit {      // if it's valid subarray, increment right index
			allmax = max(allmax, right-left)
			//fmt.Println("v1:", nums[left: right])
			continue
		}
		// if not valid, we need pop the queue (remove index=left), then increment left
		for {
			// remove index=left in queue
			if left == minq.Front().Value.(int) {
				minq.Remove(minq.Front())
			}
			if left == maxq.Front().Value.(int) {   // if index < queue.front, no need to pop actually
				maxq.Remove(maxq.Front())
			}
			left++
			curmax := nums[maxq.Front().Value.(int)]
			curmin := nums[minq.Front().Value.(int)]
			if curmax-curmin <= limit {   // if valid, stop increment left, continue to increment right
				//fmt.Println("v2:", nums[left: right])
				break
			}
		}
	}
	return allmax
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestSubarray([]int{45}, 1))
	fmt.Println(longestSubarray([]int{45}, 0))
	fmt.Println(longestSubarray([]int{8,2,4,7}, 4))
	fmt.Println(longestSubarray([]int{10,1,2,4,7,2}, 5))
	fmt.Println(longestSubarray([]int{4,2,2,2,4,4,2,2}, 0))
}
package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/

// Given an array nums, you are allowed to choose one element of nums and 
// change it by any value in one move. Return the minimum difference between 
// the largest and smallest value of nums after perfoming at most 3 moves.
// Example 1:
//   Input: nums = [5,3,2,4]
//   Output: 0
//   Explanation: Change the array [5,3,2,4] to [2,2,2,2].
//     The difference between the maximum and minimum is 2-2 = 0.
// Example 2:
//   Input: nums = [1,5,0,10,14]
//   Output: 1
//   Explanation: Change the array [1,5,0,10,14] to [1,1,0,1,1]. 
//     The difference between the maximum and minimum is 1-0 = 1.
// Example 3:
//   Input: nums = [6,6,0,1,1,4,6]
//   Output: 2
// Example 4:
//   Input: nums = [1,5,6,14,15]
//   Output: 1
// Constraints:
//   1 <= nums.length <= 10^5
//   -10^9 <= nums[i] <= 10^9

func minDifference(nums []int) int {
	if len(nums)<5 {
		return 0
	}
	// IMPROVEMENT: actually, no need to sort the whole array, 
	// we need only the top4 smallest and largest values
	sort.Sort(sort.IntSlice(nums))
	L := len(nums)
	mindiff := nums[L-1]-nums[0]
	// We have 4 plans:
    //   kill 3 biggest elements
    //   kill 2 biggest elements + 1 smallest elements
    //   kill 1 biggest elements + 2 smallest elements
    //   kill 3 smallest elements
	if mindiff > nums[L-4]-nums[0] {
		mindiff = nums[L-4]-nums[0]
	}
	if mindiff > nums[L-3]-nums[1] {
		mindiff = nums[L-3]-nums[1]
	}
	if mindiff > nums[L-2]-nums[2] {
		mindiff = nums[L-2]-nums[2]
	}
	if mindiff > nums[L-1]-nums[3] {
		mindiff = nums[L-1]-nums[3]
	}
	return mindiff
}

func main() {
	fmt.Println(minDifference([]int{5,3,2,4}))
	fmt.Println(minDifference([]int{1,5,0,10,14}))
	fmt.Println(minDifference([]int{6,6,0,1,1,4,6}))
	fmt.Println(minDifference([]int{1,5,6,14,15}))
}
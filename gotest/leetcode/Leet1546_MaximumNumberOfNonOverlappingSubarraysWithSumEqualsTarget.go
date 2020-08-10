package main

import "fmt"

// https://leetcode.com/problems/maximum-number-of-non-overlapping-subarrays-with-sum-equals-target

// Given an array nums and an integer target.
// Return the maximum number of non-empty non-overlapping subarrays such that the sum of values
// in each subarray is equal to target.
// Example 1:
//   Input: nums = [1,1,1,1,1], target = 2
//   Output: 2
//   Explanation: There are 2 non-overlapping subarrays [1,1,1,1,1] with sum equals
//     to target(2).
// Example 2:
//   Input: nums = [-1,3,5,1,4,2,-9], target = 6
//   Output: 2
//   Explanation: There are 3 subarrays with sum equal to 6.
//     ([5,1], [4,2], [3,5,1,4,2,-9]) but only the first 2 are non-overlapping.
// Example 3:
//   Input: nums = [-2,6,6,3,5,4,1,2,8], target = 10
//   Output: 3
// Example 4:
//   Input: nums = [0,0,0], target = 0
//   Output: 3
// Constraints:
//   1 <= nums.length <= 10^5
//   -10^4 <= nums[i] <= 10^4
//   0 <= target <= 10^6

func maxNonOverlapping(nums []int, target int) int {
	mapp := make(map[int]int)
	mapp[0] = -1
	maxend := -1
	count := 0
	sum := 0
	for i := range nums {
		sum += nums[i]
		if v, ok := mapp[sum-target]; ok {   // check if sum-target is in previous sums
			if v+1 > maxend {
				count++   // greedy, if we can find a non-overlapping target subarray, count it
				maxend = i
			}
		}
		mapp[sum] = i
	}
	return count
}

func main() {
	for _, v := range []struct{arr []int; target, ans int} {
		{[]int{1,1,1,1,1}, 2, 2},
		{[]int{-1,3,5,1,4,2,-9}, 6, 2},
		{[]int{-2,6,6,3,5,4,1,2,8}, 10, 3},
		{[]int{0,0,0}, 0, 3},
		{[]int{3,0,2,0,2,3,3,0,0,2,1,1,1,0,-1,-1,1,-1,1,0,2,0,0,3,0,0,3,1,0,2,0,-1,2,-1,1,1,3,0,2,3,3,0,0,2,-1,1}, 3, 12},
	} {
		fmt.Println(maxNonOverlapping(v.arr, v.target), v.ans)
	}
}
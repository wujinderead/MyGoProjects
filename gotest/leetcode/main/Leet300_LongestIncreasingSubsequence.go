package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/longest-increasing-subsequence/
// Given an unsorted array of integers, find the length of longest increasing subsequence.
// Example:
//   Input: [10,9,2,5,3,7,101,18]
//   Output: 4
//   Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
// Your algorithm should run in O(n^2) complexity.
// Follow up: Could you improve it to O(n log n) time complexity?

// e.g. input: [0, 8, 4, 12, 2], use binary search to insert or replace dp.
// dp: [0]              insert 0
// dp: [0, 8]           insert 8
// dp: [0, 4]           insert 4
// dp: [0, 4, 12]       insert 12
// dp: [0, 2, 12]       insert 2
// dp[0...i] is the minimal values for LIS length i
// e.g., [0] is minimal value for lis length 1; [0,2] are minimal values for lis length 2
func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	// dp O(nlogn) method
	dp := make([]int, 0, len(nums))
	for _, v := range nums {
		ind := sort.SearchInts(dp, v)
		if ind == len(dp) {
			dp = append(dp, v)
		} else {
			dp[ind] = v
		}
	}
	return len(dp)
}

func main() {
	fmt.Println(lengthOfLIS([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}))
	fmt.Println(lengthOfLIS([]int{0, 8, 4, 12, 2}))
}

package main

import (
	"fmt"
	"container/list"
)

// https://leetcode.com/problems/constrained-subsequence-sum/

// Given an integer array nums and an integer k, return the maximum sum of a non-empty
// subsequence of that array such that for every two consecutive integers in
// the subsequence, nums[i] and nums[j], where i<j, the condition j-i<=k is satisfied.
// A subsequence of an array is obtained by deleting some number of elements (can be zero)
// from the array, leaving the remaining elements in their original order.
// Example 1:
//   Input: nums = [10,2,-10,5,20], k = 2
//   Output: 37
//   Explanation: The subsequence is [10, 2, 5, 20].
// Example 2:
//   Input: nums = [-1,-2,-3], k = 1
//   Output: -1
//   Explanation: The subsequence must be non-empty, so we choose the largest number.
// Example 3:
//   Input: nums = [10,-2,-10,-5,20], k = 2
//   Output: 23
//   Explanation: The subsequence is [10, -2, -5, 20].
// Constraints:
//   1 <= k <= nums.length <= 10^5
//   -10^4 <= nums[i] <= 10^4

func constrainedSubsetSum(nums []int, k int) int {
    // we can always include positive number, for negative number,
    // we need dp to check the max value we can get
    dp := make([]int, len(nums))   // dp[i] is the max sum for nums[i:] with nums[i] included
    dp[len(nums)-1] = nums[len(nums)-1]
	allmax := nums[len(nums)-1]
	l := list.New()
	l.PushBack(len(nums)-1)

	// first get max for dp[i+1], ..., dp[i+k]
	// then check if we need pop the queue (l.front-i == k)
	// then add i to queue back, before add, delete all elements that less than dp[i]
    for i:=len(nums)-2; i>=0; i--{
    	// max(dp[i+1], ..., dp[i+k]) = dp[l.front]
		// dp[i]=nums[i]+max(dp[i+1], dp[i+2], ..., dp[i+k], 0)
		dp[i] = nums[i] + max(dp[l.Front().Value.(int)], 0)
    	if l.Len()>0 && l.Front().Value.(int)-i == k {      // l.front-i == k
    		l.Remove(l.Front())
		}
		for l.Len()>0 && dp[l.Back().Value.(int)] < dp[i] {
			l.Remove(l.Back())
		}
		l.PushBack(i)
		if dp[i]>allmax {
			allmax = dp[i]
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
	fmt.Println(constrainedSubsetSum([]int{10,2,-10,5,20}, 2))
	fmt.Println(constrainedSubsetSum([]int{-1,-2,-3}, 1))
	fmt.Println(constrainedSubsetSum([]int{-3,-1,-2}, 1))
	fmt.Println(constrainedSubsetSum([]int{10,-2,-10,-5,20}, 1))
	fmt.Println(constrainedSubsetSum([]int{10,-2,-10,-5,20}, 2))
}
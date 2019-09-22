package main

import "fmt"

// https://leetcode.com/problems/house-robber/

// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed, the only constraint
// stopping you from robbing each of them is that adjacent houses have
// security system connected and it will automatically contact the police
// if two adjacent houses were broken into on the same night.
// Given a list of non-negative integers representing the amount of money of each house,
// determine the maximum amount of money you can rob tonight without alerting the police.
// Example 1:
//   Input: [1,2,3,1]
//   Output: 4
//   Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//                Total amount you can rob = 1 + 3 = 4.
//
// Example 2:
//   Input: [2,7,9,3,1]
//   Output: 12
//   Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
//                Total amount you can rob = 2 + 9 + 1 = 12.

func rob(nums []int) int {
	// let sum[i] be the max amount to rob for house[0...i],
	// then sum[i] = max( sum[i-1], sum[i-2]+nums[i] )
	if len(nums) == 0 {
		return 0
	}
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	sum[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		sum[i] = max(sum[i-1], sum[i-2]+nums[i-1])
	}
	return sum[len(nums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{}))
	fmt.Println(rob([]int{2}))
	fmt.Println(rob([]int{1, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

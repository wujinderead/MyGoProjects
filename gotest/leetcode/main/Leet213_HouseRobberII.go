package main

import "fmt"

// https://leetcode.com/problems/house-robber-ii/

// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed. All houses at this place
// are arranged in a circle. That means the first house is the neighbor of the last one.
// Meanwhile, adjacent houses have security system connected and it will automatically
// contact the police if two adjacent houses were broken into on the same night.
// Given a list of non-negative integers representing the amount of money of each house,
// determine the maximum amount of money you can rob tonight without alerting the police.
// Example 1:
//   Input: [2,3,2]
//   Output: 3
//   Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
//                because they are adjacent houses.
// Example 2:
//   Input: [1,2,3,1]
//   Output: 4
//   Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//                Total amount you can rob = 1 + 3 = 4.

func rob(nums []int) int {
	// since the first and last house are adjacent, they can't be robbed together,
	// so the answer is the max( houseRobI(house[0...n-1]), houseRobI(house[1...n]) )
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	// get max for house[1...n]
	sum := make([]int, len(nums)+1) // the index for sum is 1 based, sum[i] relates to nums[i-1]
	sum[1] = 0
	sum[2] = nums[1]
	for i := 3; i <= len(nums); i++ {
		sum[i] = max(sum[i-1], sum[i-2]+nums[i-1])
	}
	// get max for house[0...n-1]
	sum[0] = 0
	sum[1] = nums[0]
	for i := 2; i <= len(nums)-1; i++ {
		sum[i] = max(sum[i-1], sum[i-2]+nums[i-1])
	}
	return max(sum[len(nums)], sum[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{}))
	fmt.Println(rob([]int{3}))
	fmt.Println(rob([]int{3, 5}))
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

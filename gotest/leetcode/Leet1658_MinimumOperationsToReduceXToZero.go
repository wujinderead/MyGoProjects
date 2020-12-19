package main

import "fmt"

// https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/

// You are given an integer array nums and an integer x. In one operation, you can either
// remove the leftmost or the rightmost element from the array nums and subtract its value from x.
// Note that this modifies the array for future operations.
// Return the minimum number of operations to reduce x to exactly 0 if it's possible, otherwise, return -1.
// Example 1:
//   Input: nums = [1,1,4,2,3], x = 5
//   Output: 2
//   Explanation: The optimal solution is to remove the last two elements to reduce x to zero.
// Example 2:
//   Input: nums = [5,6,7,8,9], x = 4
//   Output: -1
// Example 3:
//   Input: nums = [3,2,20,1,1,3], x = 10
//   Output: 5
//   Explanation: The optimal solution is to remove the last three elements and the
//     first two elements (5 operations in total) to reduce x to zero.
// Constraints:
//   1 <= nums.length <= 10^5
//   1 <= nums[i] <= 10^4
//   1 <= x <= 10^9

// subarray sum method
func minOperations(nums []int, x int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if x == sum {
		return len(nums)
	}
	target := sum - x // find the longest subarray that sum is target
	sum = 0
	mapp := make(map[int]int)
	mapp[0] = -1
	ans := -1
	for i := range nums {
		sum += nums[i]
		if v, ok := mapp[sum-target]; ok {
			if i-v > ans {
				ans = i - v
			}
		}
		mapp[sum] = i
	}
	if ans == -1 {
		return -1
	}
	return len(nums) - ans
}

// two pointer method: let two parts sum to x
func minOperations1(nums []int, x int) int {
	left, right := -1, len(nums)
	sum := 0
	min := len(nums) + 1
	for left+1 < len(nums) && sum+nums[left+1] <= x {
		sum += nums[left+1]
		left++
	}
	for left >= -1 {
		for right-1 >= 0 && sum+nums[right-1] <= x {
			sum += nums[right-1]
			right--
		}
		if sum == x && left+1+len(nums)-right < min {
			min = left + 1 + len(nums) - right
		}
		if left >= 0 {
			sum -= nums[left]
		}
		left--
	}
	if min == len(nums)+1 {
		return -1
	}
	return min
}

func main() {
	for _, v := range []struct {
		nums   []int
		x, ans int
	}{
		{[]int{1, 1, 4, 2, 3}, 5, 2},
		{[]int{5, 6, 7, 8, 9}, 4, -1},
		{[]int{3, 2, 20, 1, 1, 3}, 10, 5},
		{[]int{3, 1, 1, 1, 1, 5}, 5, 1},
		{[]int{5, 1, 1, 1, 1, 3}, 5, 1},
		{[]int{6, 1, 1, 1, 1, 3}, 5, 3},
		{[]int{1, 2, 3, 4, 5}, 5, 1},
		{[]int{1, 2, 3, 4, 5}, 14, 4},
		{[]int{1, 2, 3, 4, 5}, 15, 5},
		{[]int{1, 2, 3, 4, 5}, 16, -1},
	} {
		fmt.Println(minOperations(v.nums, v.x), minOperations1(v.nums, v.x), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/max-number-of-k-sum-pairs/

// You are given an integer array nums and an integer k.
// In one operation, you can pick two numbers from the array whose sum equals k and
// remove them from the array.
// Return the maximum number of operations you can perform on the array.
// Example 1:
//   Input: nums = [1,2,3,4], k = 5
//   Output: 2
//   Explanation: Starting with nums = [1,2,3,4]:
//     - Remove numbers 1 and 4, then nums = [2,3]
//     - Remove numbers 2 and 3, then nums = []
//     There are no more pairs that sum up to 5, hence a total of 2 operations.
// Example 2:
//   Input: nums = [3,1,3,4,3], k = 6
//   Output: 1
//   Explanation: Starting with nums = [3,1,3,4,3]:
//     - Remove the first two 3's, then nums = [1,4,3]
//     There are no more pairs that sum up to 6, hence a total of 1 operation.
// Constraints:
//   1 <= nums.length <= 10^5
//   1 <= nums[i] <= 10^9
//   1 <= k <= 10^9

func maxOperations(nums []int, k int) int {
	ans := 0
	mapp := make(map[int]int)
	for _, v := range nums {
		if va := mapp[k-v]; va > 0 {
			mapp[k-v] = va - 1
			ans++
		} else {
			mapp[v] = mapp[v] + 1
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		nums   []int
		k, ans int
	}{
		{[]int{1, 2, 3, 4}, 5, 2},
		{[]int{3, 1, 3, 4, 3}, 6, 1},
	} {
		fmt.Println(maxOperations(v.nums, v.k), v.ans)
	}
}

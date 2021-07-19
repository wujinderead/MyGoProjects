package main

import "fmt"

// https://leetcode.com/problems/maximum-alternating-subsequence-sum/

// The alternating sum of a 0-indexed array is defined as the sum of the elements at even indices
// minus the sum of the elements at odd indices.
// For example, the alternating sum of [4,2,5,3] is (4 + 5) - (2 + 3) = 4.
// Given an array nums, return the maximum alternating sum of any subsequence of
// nums (after reindexing the elements of the subsequence).
// A subsequence of an array is a new array generated from the original array by
// deleting some elements (possibly none) without changing the remaining elements'
// relative order. For example, [2,7,4] is a subsequence of [4,2,3,7,2,1,4]
// (the underlined elements), while [2,4,2] is not.
// Example 1:
//   Input: nums = [4,2,5,3]
//   Output: 7
//   Explanation: It is optimal to choose the subsequence [4,2,5] with alternating sum (4 + 5) - 2 = 7.
// Example 2:
//   Input: nums = [5,6,7,8]
//   Output: 8
//   Explanation: It is optimal to choose the subsequence [8] with alternating sum 8.
// Example 3:
//   Input: nums = [6,2,1,2,4,5]
//   Output: 10
//   Explanation: It is optimal to choose the subsequence [6,1,5] with alternating sum (6 + 5) - 1 = 10.
// Constraints:
//   1 <= nums.length <= 10^5
//   1 <= nums[i] <= 10^5

func maxAlternatingSum(nums []int) int64 {
	// for the first k elements, let Even and Odd be the max diff we can get,
	// if we use odd or even numbers of elements
	even, odd := 0, nums[0]
	var newe, newo int

	// for each num, check whether we want to include it; if we want include it:
	// if previous has odd nums, we have even = odd - nums[i]
	// if previous has event nums, we have odd = even + nums[i]
	for i := 1; i < len(nums); i++ {
		newo, newe = odd, even
		if odd-nums[i] > even {
			newe = odd - nums[i]
		}
		if even+nums[i] > odd {
			newo = even + nums[i]
		}
		odd, even = newo, newe
	}
	return int64(odd)
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int64
	}{
		{[]int{4, 2, 5, 3}, 7},
		{[]int{5, 6, 7, 8}, 8},
		{[]int{6, 2, 1, 2, 4, 5}, 10},
	} {
		fmt.Println(maxAlternatingSum(v.nums), v.ans)
	}
}

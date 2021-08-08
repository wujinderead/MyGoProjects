package main

import "fmt"

// https://leetcode.com/problems/count-number-of-special-subsequences/

// A sequence is special if it consists of a positive number of 0s, followed by a positive number
// of 1s, then a positive number of 2s.
// For example, [0,1,2] and [0,0,1,1,1,2] are special.
// In contrast, [2,1,0], [1], and [0,1,2,0] are not special.
// Given an array nums (consisting of only integers 0, 1, and 2), return the number of different
// subsequences that are special. Since the answer may be very large, return it modulo 10^9 + 7.
// A subsequence of an array is a sequence that can be derived from the array by deleting some
// or no elements without changing the order of the remaining elements. Two subsequences are
// different if the set of indices chosen are different.
// Example 1:
//   Input: nums = [0,1,2,2]
//   Output: 3
//   Explanation: The special subsequences are [0,1,2,2], [0,1,2,2], and [0,1,2,2].
// Example 2:
//   Input: nums = [2,2,0,0]
//   Output: 0
//   Explanation: There are no special subsequences in [2,2,0,0].
// Example 3:
//   Input: nums = [0,1,2,0,1,2]
//   Output: 7
//   Explanation: The special subsequences are:
//     - [0,1,2,0,1,2]
//     - [0,1,2,0,1,2]
//     - [0,1,2,0,1,2]
//     - [0,1,2,0,1,2]
//     - [0,1,2,0,1,2]
//     - [0,1,2,0,1,2]
//     - [0,1,2,0,1,2]
// Constraints:
//   1 <= nums.length <= 10^5
//   0 <= nums[i] <= 2

func countSpecialSubsequences(nums []int) int {
	var zero, one, two int
	const mod = int(1e9 + 7)
	if nums[0] == 0 {
		zero = 1
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			zero = (zero*2 + 1) % mod
		}
		// we have x '0...0' and y '0...1' before this '1', then,
		// append '1' to x '0...0', got x '0...1'
		// append '1' to y '0...1', got y '0...1'
		// no append, got y '0...1', finally got y*2+x '0...1'
		if nums[i] == 1 {
			one = (one*2 + zero) % mod
		}
		// we have x '0...1' and y '0.1.2' before this '2', then,
		// append '2' to x '0...1', got x '0...12'
		// append '2' to y '0.1.2', got y '0.1.22'
		// no append, got y '0.1.2', finally got y*2+x '0.1.2'
		if nums[i] == 2 {
			two = (two*2 + one) % mod
		}
	}
	return two
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{0, 1, 2, 2}, 3},
		{[]int{2, 2, 0, 0}, 0},
		{[]int{0, 1, 2, 0, 1, 2}, 7},
	} {
		fmt.Println(countSpecialSubsequences(v.nums), v.ans)
	}
}

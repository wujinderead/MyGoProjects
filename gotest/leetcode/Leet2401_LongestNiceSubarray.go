package main

import "fmt"

// https://leetcode.com/problems/longest-nice-subarray/

// You are given an array nums consisting of positive integers.
// We call a subarray of nums nice if the bitwise AND of every pair of elements that are in different
// positions in the subarray is equal to 0.
// Return the length of the longest nice subarray.
// A subarray is a contiguous part of an array.
// Note that subarrays of length 1 are always considered nice.
// Example 1:
// Input: nums = [1,3,8,48,10]
// Output: 3
// Explanation: The longest nice subarray is [3,8,48]. This subarray satisfies the conditions:
//   - 3 AND 8 = 0.
//   - 3 AND 48 = 0.
//   - 8 AND 48 = 0.
//   It can be proven that no longer nice subarray can be obtained, so we return 3.
// Example 2:
//   Input: nums = [3,1,5,11,13]
//   Output: 1
//   Explanation: The length of the longest nice subarray is 1. Any subarray of length 1 can be chosen.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁹

func longestNiceSubarray1(nums []int) int {
	occur := [30]int{}
	leng := [30]int{}
	for j := range occur { // prev[j]=i, nums[i]'s j-th bit is the last occurrence
		occur[j] = -1
	}
	max := 0

	for i := range nums {
		min := 1 << 30
		for j := range occur {
			if (1<<j)&nums[i] > 0 { //nums[i]'s j-th bit is 1
				leng[j] = i - occur[j]
				occur[j] = i
			} else {
				leng[j] += 1
			}
			if leng[j] < min {
				min = leng[j]
			}
		}
		if min > max {
			max = min
		}
	}
	return max
}

// sliding window
func longestNiceSubarray(nums []int) int {
	max := 0
	start := 0
	set := 0
	for i := 0; i < len(nums); i++ {
		for set&nums[i] > 0 {
			set = set ^ nums[start]
			start++
		}
		set |= nums[i]
		if i-start+1 > max {
			max = i - start + 1
		}
	}
	return max
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 3, 8, 48, 10}, 3},
		{[]int{3, 1, 5, 11, 13}, 1},
	} {
		fmt.Println(longestNiceSubarray(v.nums), longestNiceSubarray1(v.nums), v.ans)
	}
}

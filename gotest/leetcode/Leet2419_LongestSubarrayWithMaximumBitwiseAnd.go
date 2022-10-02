package main

import "fmt"

// https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/

// You are given an integer array nums of size n.
// Consider a non-empty subarray from nums that has the maximum possible bitwise AND.
// In other words, let k be the maximum value of the bitwise AND of any subarray
// of nums. Then, only subarrays with a bitwise AND equal to k should be considered.
// Return the length of the longest such subarray.
// The bitwise AND of an array is the bitwise AND of all the numbers in it.
// A subarray is a contiguous sequence of elements within an array.
// Example 1:
//   Input: nums = [1,2,3,3,2,2]
//   Output: 2
//   Explanation:
//     The maximum possible bitwise AND of a subarray is 3.
//     The longest subarray with that value is [3,3], so we return 2.
// Example 2:
//   Input: nums = [1,2,3,4]
//   Output: 1
//   Explanation:
//     The maximum possible bitwise AND of a subarray is 4.
//     The longest subarray with that value is [4], so we return 1.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁶

// maximum bitwise and is the max number
func longestSubarray(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	maxlen := 1
	curlen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == max {
			curlen++
		} else {
			curlen = 0
		}
		if curlen > maxlen {
			maxlen = curlen
		}
	}
	return maxlen
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{1, 2, 3, 3, 2, 2}, 2},
		{[]int{1, 2, 3, 4}, 1},
		{[]int{100, 5, 5}, 1},
	} {
		fmt.Println(longestSubarray(v.nums), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/number-of-subarrays-with-lcm-equal-to-k/

// Given an integer array nums and an integer k, return the number of subarrays of nums where
// the least common multiple of the subarray's elements is k.
// A subarray is a contiguous non-empty sequence of elements within an array.
// The least common multiple of an array is the smallest positive integer that is divisible by all
// the array elements.
// Example 1:
//   Input: nums = [3,6,2,7,1], k = 6
//   Output: 4
//    Explanation: The subarrays of nums where 6 is the least common multiple of all
//      the subarray's elements are:
//      - [3,6,2,7,1]
//      - [3,6,2,7,1]
//      - [3,6,2,7,1]
//      - [3,6,2,7,1]
// Example 2:
//   Input: nums = [3], k = 2
//   Output: 0
//   Explanation: There are no subarrays of nums where 2 is the least common
//     multiple of all the subarray's elements.
// Constraints:
//   1 <= nums.length <= 1000
//   1 <= nums[i], k <= 1000

func subarrayLCM(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		lcm := nums[i]
		for j := i; j < len(nums); j++ {
			g := gcd(lcm, nums[j])
			lcm = lcm * nums[j] / g
			if lcm == k {
				count++
			}
			if lcm > k {
				break
			}
		}
	}
	return count
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func main() {
	for _, v := range []struct {
		nums   []int
		k, ans int
	}{
		{[]int{3, 6, 2, 7, 1}, 6, 4},
		{[]int{3}, 2, 0},
	} {
		fmt.Println(subarrayLCM(v.nums, v.k), v.ans)
	}
}

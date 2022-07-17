package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-deletions-to-make-array-divisible/

// You are given two positive integer arrays nums and numsDivide. You can delete any number of
// elements from nums.
// Return the minimum number of deletions such that the smallest element in nums divides all the
// elements of numsDivide. If this is not possible, return -1.
// Note that an integer x divides y if y % x == 0.
// Example 1:
//   Input: nums = [2,3,2,4,3], numsDivide = [9,6,9,3,15]
//   Output: 2
//   Explanation:
//     The smallest element in [2,3,2,4,3] is 2, which does not divide all the elements of numsDivide.
//     We use 2 deletions to delete the elements in nums that are equal to 2 which makes nums = [3,4,3].
//     The smallest element in [3,4,3] is 3, which divides all the elements of numsDivide.
//     It can be shown that 2 is the minimum number of deletions needed.
// Example 2:
//   Input: nums = [4,3,6], numsDivide = [8,2,6,10]
//   Output: -1
//   Explanation:
//     We want the smallest element in nums to divide all the elements of numsDivide.
//     There is no way to delete elements from nums to allow this.
// Constraints:
//   1 <= nums.length, numsDivide.length <= 10⁵
//   1 <= nums[i], numsDivide[i] <= 10⁹

func minOperations(nums []int, numsDivide []int) int {
	// find the GCD of numsDivide
	g := numsDivide[0]
	for i := 1; i < len(numsDivide); i++ {
		g = gcd(g, numsDivide[i])
	}

	// find smallest nums[i] that divides GCD
	small := 1 << 31
	for i := range nums {
		if g%nums[i] == 0 && nums[i] < small {
			small = nums[i]
		}
	}

	// find how many number < small
	if small == 1<<31 {
		return -1
	}
	count := 0
	for i := range nums {
		if nums[i] < small {
			count++
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
		nums, numsDivide []int
		ans              int
	}{
		{[]int{2, 3, 2, 4, 3}, []int{9, 6, 9, 3, 15}, 2},
		{[]int{4, 3, 6}, []int{8, 2, 6, 10}, -1},
		{[]int{3, 2, 6, 2, 35, 5, 35, 2, 5, 8, 7, 3, 4}, []int{105, 70, 70, 175, 105, 105, 105}, 6},
	} {
		fmt.Println(minOperations(v.nums, v.numsDivide), v.ans)
	}
}

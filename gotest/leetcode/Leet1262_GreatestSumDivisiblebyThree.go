package main

import (
	"fmt"
	"math/rand"
)

// https://leetcode.com/problems/greatest-sum-divisible-by-three

// Given an array nums of integers, we need to find the maximum possible sum
// of elements of the array such that it is divisible by three.
// Example 1:
//   Input: nums = [3,6,5,1,8]
//   Output: 18
//   Explanation: Pick numbers 3, 6, 1 and 8 their sum is 18 (maximum sum divisible by 3).
// Example 2:
//   Input: nums = [4]
//   Output: 0
//   Explanation: Since 4 is not divisible by 3, do not pick any number.
// Example 3:
//   Input: nums = [1,2,3,4,4]
//   Output: 12
//   Explanation: Pick numbers 1, 3, 4 and 4 their sum is 12 (maximum sum divisible by 3).
// Constraints:
//   1 <= nums.length <= 4 * 10^4
//   1 <= nums[i] <= 10^4

func maxSumDivThree(nums []int) int {
	r := make([][]int, 3)
	r[0] = make([]int, len(nums)+1)
	r[1] = make([]int, len(nums)+1)
	r[2] = make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		r[0][i] = r[0][i-1] // record the remain divided by 3
		r[1][i] = r[1][i-1] // need only O(3) space since line by line
		r[2][i] = r[2][i-1]
		for j := 0; j < 3; j++ {
			sum := r[j][i-1] + nums[i-1]
			if r[sum%3][i] < sum {
				r[sum%3][i] = sum
			}
		}
	}
	//fmt.Println(r)
	return r[0][len(nums)]
}

func main() {
	fmt.Println(maxSumDivThree([]int{3, 6, 5, 1, 8}))
	fmt.Println(maxSumDivThree([]int{1, 2, 3, 4, 4}))
	fmt.Println(maxSumDivThree([]int{4}))
	ints := rand.Perm(100)[:30]
	fmt.Println(ints, maxSumDivThree(ints))
}

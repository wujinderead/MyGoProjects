package main

import "fmt"

// https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/

// You are given an integer array nums sorted in non-decreasing order.
// Build and return an integer array result with the same length as nums such that
// result[i] is equal to the summation of absolute differences between nums[i] and
// all the other elements in the array.
// In other words, result[i] is equal to sum(|nums[i]-nums[j]|) where 0 <= j < nums.length
// and j != i (0-indexed).
// Example 1:
//   Input: nums = [2,3,5]
//   Output: [4,3,5]
//   Explanation: Assuming the arrays are 0-indexed, then
//     result[0] = |2-2| + |2-3| + |2-5| = 0 + 1 + 3 = 4,
//     result[1] = |3-2| + |3-3| + |3-5| = 1 + 0 + 2 = 3,
//     result[2] = |5-2| + |5-3| + |5-5| = 3 + 2 + 0 = 5.
// Example 2:
//   Input: nums = [1,4,6,8,10]
//   Output: [24,15,13,15,21]
// Constraints:
//   2 <= nums.length <= 10^5
//   1 <= nums[i] <= nums[i + 1] <= 10^4

// for a sorted array as [x0,x1,x2,x3,x4], the differences between adjacent numbers are
// x1-x0=a, x2-x1=b, x3-x2=c, x4-x3=d. then the result array is
// r0 = 4a+3b+2c+1d
// r1 = 1, 3  2  1
// r2 = 1  2, 2  1
// r3 = 1  2  3, 1
// r4 = 1  2  3  4,  the coefficients only change one per time
func getSumAbsoluteDifferences(nums []int) []int {
	ans := make([]int, len(nums))
	cos := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		cos[i] = len(nums) - i
		ans[0] += (nums[i] - nums[i-1]) * cos[i]
	}
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] + (nums[i]-nums[i-1])*(cos[i-1]+1-cos[i])
		cos[i] = cos[i-1] + 1
	}
	return ans
}

func main() {
	for _, v := range []struct {
		nums, ans []int
	}{
		{[]int{2, 3, 5}, []int{4, 3, 5}},
		{[]int{1, 4, 6, 8, 10}, []int{24, 15, 13, 15, 21}},
	} {
		fmt.Println(getSumAbsoluteDifferences(v.nums), v.ans)
	}
}

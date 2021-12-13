package main

import "fmt"

// https://leetcode.com/problems/sum-of-subarray-ranges/

// You are given an integer array nums. The range of a subarray of nums is the difference between
// the largest and smallest element in the subarray.
// Return the sum of all subarray ranges of nums.
// A subarray is a contiguous non-empty sequence of elements within an array.
// Example 1:
//   Input: nums = [1,2,3]
//   Output: 4
//   Explanation: The 6 subarrays of nums are the following:
//     [1], range = largest - smallest = 1 - 1 = 0
//     [2], range = 2 - 2 = 0
//     [3], range = 3 - 3 = 0
//     [1,2], range = 2 - 1 = 1
//     [2,3], range = 3 - 2 = 1
//     [1,2,3], range = 3 - 1 = 2
//     So the sum of all ranges is 0 + 0 + 0 + 1 + 1 + 2 = 4.
// Example 2:
//   Input: nums = [1,3,3]
//   Output: 4
//   Explanation: The 6 subarrays of nums are the following:
//     [1], range = largest - smallest = 1 - 1 = 0
//     [3], range = 3 - 3 = 0
//     [3], range = 3 - 3 = 0
//     [1,3], range = 3 - 1 = 2
//     [3,3], range = 3 - 3 = 0
//     [1,3,3], range = 3 - 1 = 2
//     So the sum of all ranges is 0 + 0 + 0 + 2 + 0 + 2 = 4.
// Example 3:
//   Input: nums = [4,-2,-3,4,1]
//   Output: 59
//   Explanation: The sum of all subarray ranges of nums is 59.
// Constraints:
//   1 <= nums.length <= 1000
//   -10⁹ <= nums[i] <= 10⁹

func subArrayRanges(nums []int) int64 {
	sum := 0
	for i := 0; i < len(nums); i++ {
		max := nums[i]
		min := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > max {
				max = nums[j]
			}
			if nums[j] < min {
				min = nums[j]
			}
			sum += max - min
		}
	}
	return int64(sum)
}

func main() {
	for _, v := range []struct {
		n   []int
		ans int
	}{
		{[]int{1, 2, 3}, 4},
		{[]int{1, 3, 3}, 4},
		{[]int{4, -2, -3, 4, 1}, 59},
	} {
		fmt.Println(subArrayRanges(v.n), v.ans)
	}
}

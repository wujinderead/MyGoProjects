package main

import "fmt"

// https://leetcode.com/problems/next-greater-element-ii/

// Given a circular integer array nums (i.e., the next element of nums[nums.length - 1] is nums[0]),
// return the next greater number for every element in nums.
// The next greater number of a number x is the first greater number to its traversing-order next in
// the array, which means you could search circularly to find its next greater number. If it doesn't
// exist, return -1 for this number.
// Example 1:
//   Input: nums = [1,2,1]
//   Output: [2,-1,2]
//   Explanation: The first 1's next greater number is 2;
//     The number 2 can't find next greater number.
//     The second 1's next greater number needs to search circularly, which is also 2.
// Example 2:
//   Input: nums = [1,2,3,4,3]
//   Output: [2,3,4,-1,4]
// Constraints:
//   1 <= nums.length <= 10⁴
//   -10⁹ <= nums[i] <= 10⁹

// just loop twice
func nextGreaterElements(nums []int) []int {
	var stack []int
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	ans := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			ans[i] = stack[len(stack)-1]
		} else {
			ans[i] = -1
		}
		stack = append(stack, nums[i])
	}
	return ans
}

func main() {
	for _, v := range []struct {
		nums, ans []int
	}{
		{[]int{1, 2, 1}, []int{2, -1, 2}},
		{[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
	} {
		fmt.Println(nextGreaterElements(v.nums), v.ans)
	}
}

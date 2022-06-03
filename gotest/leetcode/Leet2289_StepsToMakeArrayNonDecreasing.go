package main

import "fmt"

// https://leetcode.com/problems/steps-to-make-array-non-decreasing/

// You are given a 0-indexed integer array nums. In one step, remove all elements nums[i] where
// nums[i - 1] > nums[i] for all 0 < i < nums.length.
// Return the number of steps performed until nums becomes a non-decreasing array.
// Example 1:
//   Input: nums = [5,3,4,4,7,3,6,11,8,5,11]
//   Output: 3
//   Explanation: The following are the steps performed:
//     - Step 1: [5,3,4,4,7,3,6,11,8,5,11] becomes [5,4,4,7,6,11,11]
//     - Step 2: [5,4,4,7,6,11,11] becomes [5,4,7,11,11]
//     - Step 3: [5,4,7,11,11] becomes [5,7,11,11]
//     [5,7,11,11] is a non-decreasing array. Therefore, we return 3.
// Example 2:
//   Input: nums = [4,5,7,7,13]
//   Output: 0
//   Explanation: nums is already a non-decreasing array. Therefore, we return 0.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁹

https://leetcode.com/problems/steps-to-make-array-non-decreasing/discuss/2085864/JavaC%2B%2BPython-Stack-%2B-DP-%2B-Explanation-%2B-Poem
func totalSteps(nums []int) int {
	ans := 0
	var stack []int
	dp := make([]int, len(nums))
	for i, v := range nums {
		c := 0
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= v {
			c = max(c, dp[stack[len(stack)-1]])
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			dp[i] = c + 1
			ans = max(ans, dp[i])
		}
		stack = append(stack, i)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{5, 3, 4, 4, 7, 3, 6, 11, 8, 5, 11}, 3},
		{[]int{4, 5, 7, 7, 13}, 0},
		{[]int{7, 1, 2, 3, 6, 5, 4, 3, 2, 3, 4, 5, 6, 11}, 5},
	} {
		fmt.Println(totalSteps(v.nums), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/

// You are given a 0-indexed array nums consisting of positive integers. You can choose two indices
// i and j, such that i != j, and the sum of digits of the number nums[i] is equal to that of nums[j].
// Return the maximum value of nums[i] + nums[j] that you can obtain over all possible indices i and
// j that satisfy the conditions.
// Example 1:
//   Input: nums = [18,43,36,13,7]
//   Output: 54
//   Explanation: The pairs (i, j) that satisfy the conditions are:
//     - (0, 2), both numbers have a sum of digits equal to 9, and their sum is 18 + 36 = 54.
//     - (1, 4), both numbers have a sum of digits equal to 7, and their sum is 43 + 7 = 50.
//     So the maximum sum that we can obtain is 54.
// Example 2:
//   Input: nums = [10,12,19,14]
//   Output: -1
//   Explanation: There are no two numbers that satisfy the conditions, so we return -1.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁹

func maximumSum(nums []int) int {
	mapp := make(map[int][2]int)
	for i, v := range nums {
		sum := 0 // sum digits
		for v > 0 {
			sum += v % 10
			v /= 10
		}
		x := mapp[sum] // for same sum digits, store max two numbers
		if nums[i] > x[0] {
			x[1] = x[0]
			x[0] = nums[i]
			mapp[sum] = x
		} else if nums[i] > x[1] {
			x[1] = nums[i]
			mapp[sum] = x
		}
	}
	ans := -1
	for _, x := range mapp {
		if x[1] > 0 && x[0] > 0 && x[1]+x[0] > ans {
			ans = x[1] + x[0]
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{18, 43, 36, 13, 7}, 54},
		{[]int{10, 12, 19, 14}, -1},
	} {
		fmt.Println(maximumSum(v.nums), v.ans)
	}
}

package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-limit-of-balls-in-a-bag/

// You are given an integer array nums where the ith bag contains nums[i] balls.
// You are also given an integer maxOperations.
// You can perform the following operation at most maxOperations times:
// Take any bag of balls and divide it into two new bags with a positive number of balls.
// For example, a bag of 5 balls can become two new bags of 1 and 4 balls,
// or two new bags of 2 and 3 balls.
// Your penalty is the maximum number of balls in a bag. You want to minimize your penalty
// after the operations.
// Return the minimum possible penalty after performing the operations.
// Example 1:
//   Input: nums = [9], maxOperations = 2
//   Output: 3
//   Explanation:
//     - Divide the bag with 9 balls into two bags of sizes 6 and 3. [9] -> [6,3].
//     - Divide the bag with 6 balls into two bags of sizes 3 and 3. [6,3] -> [3,3,3].
//     The bag with the most number of balls has 3 balls, so your penalty is 3 and you should return 3.
// Example 2:
//   Input: nums = [2,4,8,2], maxOperations = 4
//   Output: 2
//   Explanation:
//     - Divide the bag with 8 balls into two bags of sizes 4 and 4. [2,4,8,2] -> [2,4,4,4,2].
//     - Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,4,4,4,2] -> [2,2,2,4,4,2].
//     - Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,2,2,4,4,2] -> [2,2,2,2,2,4,2].
//     - Divide the bag with 4 balls into two bags of sizes 2 and 2. [2,2,2,2,2,4,2] -> [2,2,2,2,2,2,2,2].
//     The bag with the most number of balls has 2 balls, so your penalty is 2 an you should return 2.
// Example 3:
//   Input: nums = [7,17], maxOperations = 2
//   Output: 7
// Constraints:
//   1 <= nums.length <= 10^5
//   1 <= maxOperations, nums[i] <= 10^9

// x operations can divide n balls into ceiling(n/(x+1)).
// use binary search to find the final target.
func minimumSize(nums []int, maxOperations int) int {
	maxe := 0
	for _, v := range nums {
		if v > maxe {
			maxe = v
		}
	}
	l, r := 1, maxe
	for l < r {
		mid := (l + r) / 2
		all := 0
		// count how many ops we need to make the max number of balls as mid
		// to divide x balls into some parts where each part <= t,
		// we need ceiling(x/t)-1 = (x-1)/t times of divide.
		for _, v := range nums {
			all += (v - 1) / mid
		}
		if all > maxOperations {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func main() {
	for _, v := range []struct {
		nums   []int
		m, ans int
	}{
		{[]int{9}, 2, 3},
		{[]int{2, 4, 8, 2}, 4, 2},
		{[]int{7, 17}, 2, 7},
	} {
		fmt.Println(minimumSize(v.nums, v.m), v.ans)
	}
}

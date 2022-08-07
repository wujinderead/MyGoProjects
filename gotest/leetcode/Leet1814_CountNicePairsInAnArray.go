package main

import "fmt"

// https://leetcode.com/problems/count-nice-pairs-in-an-array/

// You are given an array nums that consists of non-negative integers. Let us define rev(x) as
// the reverse of the non-negative integer x. For example, rev(123) = 321, and rev(120) = 21.
// A pair of indices (i, j) is nice if it satisfies all of the following conditions:
//   0 <= i < j < nums.length
//   nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])
// Return the number of nice pairs of indices. Since that number can be too large, return it
// modulo 10⁹ + 7.
// Example 1:
//   Input: nums = [42,11,1,97]
//   Output: 2
//   Explanation: The two pairs are:
//     - (0,3) : 42 + rev(97) = 42 + 79 = 121, 97 + rev(42) = 97 + 24 = 121.
//     - (1,2) : 11 + rev(1) = 11 + 1 = 12, 1 + rev(11) = 1 + 11 = 12.
// Example 2:
//   Input: nums = [13,10,35,24,76]
//   Output: 4
// Constraints:
//   1 <= nums.length <= 10⁵
//   0 <= nums[i] <= 10⁹

// nums[i] + rev(nums[j]) == nums[j] + rev(nums[i]) equals to
// nums[i] - rev(nums[i]) == nums[j] - rev(nums[j])
func countNicePairs(nums []int) int {
	count := 0
	const p = int(1e9) + 7
	mapp := make(map[int]int)
	for i, v := range nums {
		rev := 0
		for v > 0 {
			rev = rev*10 + v%10
			v /= 10
		}
		x := nums[i] - rev
		count = (count + mapp[x]) % p
		mapp[x] += 1
	}
	return count
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{42, 11, 1, 97}, 2},
		{[]int{13, 10, 35, 24, 76}, 4},
	} {
		fmt.Println(countNicePairs(v.nums), v.ans)
	}
}

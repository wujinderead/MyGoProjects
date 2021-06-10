package main

import (
	"fmt"
)

// https://leetcode.com/problems/sum-of-floored-pairs/

// Given an integer array nums, return the sum of floor(nums[i] / nums[j]) for all pairs
// of indices 0 <= i, j < nums.length in the array. Since the answer may be too large,
// return it modulo 109 + 7.
// The floor() function returns the integer part of the division.
// Example 1:
//   Input: nums = [2,5,9]
//   Output: 10
//   Explanation:
//     floor(2 / 5) = floor(2 / 9) = floor(5 / 9) = 0
//     floor(2 / 2) = floor(5 / 5) = floor(9 / 9) = 1
//     floor(5 / 2) = 2
//     floor(9 / 2) = 4
//     floor(9 / 5) = 1
//     We calculate the floor of the division for every pair of indices in the array then sum them up.
// Example 2:
//   Input: nums = [7,7,7,7,7,7,7]
//   Output: 49
// Constraints:
//   1 <= nums.length <= 10^5
//   1 <= nums[i] <= 10^5

// calculate frequency of numbers
// for example nums=[2,3,5,6,9], frequency of numbers is
//   num   0 1 2 3 4 5 6 7 8 9
//   freq  0 0 1 1 0 1 1 0 0 1
// if the divisor is 3, then
//   0,1,2 contribute 0 to the final answer, ans += sum(freq[0:2])*0
//   3,4,5 contribute 1 to the final answer, ans += sum(freq[3:5])*1
//   6,7,8 contribute 2 to the final answer, ans += sum(freq[6:8])*2
//   9,10,11 contribute 3 to the final answer, ans += sum(freq[9:11])*3
// worst case: time is n + n/1 + n/2 + ...n/n ≈ nlogn
func sumOfFlooredPairs(nums []int) int {
	mod := int(1e9 + 7)
	max := int(10e5)
	freq := make([]int, max+1)
	for _, v := range nums {
		freq[v]++
	}
	pre := make([]int, max+1)
	for i := 1; i < max+1; i++ {
		pre[i] = pre[i-1] + freq[i]
	}
	sum := 0
	for i := 1; i < max+1; i++ {
		if freq[i] > 0 {
			for times := 1; times*i <= max; times++ {
				x := (times+1)*i - 1
				if x > max {
					x = max
				}
				sum += freq[i] * (pre[x] - pre[times*i-1]) * times
				sum = sum % mod
			}
		}
	}
	return sum
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{2, 5, 9}, 10},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 49},
	} {
		fmt.Println(sumOfFlooredPairs(v.nums), v.ans)
	}
	x := make([]int, 10000)
	for i := range x {
		x[i] = 1
	}
	x[9999] = 10000
	fmt.Println(sumOfFlooredPairs(x), 9999*9999+10000*9999+1)
}

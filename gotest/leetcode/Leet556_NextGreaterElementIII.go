package main

import (
	"fmt"
)

// https://leetcode.com/problems/next-greater-element-iii/

// Given a positive integer n, find the smallest integer which has exactly the same digits existing in
// the integer n and is greater in value than n. If no such positive integer exists, return -1.
// Note that the returned integer should fit in 32-bit integer, if there is a valid answer but it does
// not fit in 32-bit integer, return -1.
// Example 1:
//   Input: n = 12
//   Output: 21
// Example 2:
//   Input: n = 21
//   Output: -1
// Constraints:
//   1 <= n <= 2³¹ - 1

// same to LeetCode 31, next permutation
func nextGreaterElement(n int) int {
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}
	for i := 1; i < len(digits); i++ {
		if digits[i] < digits[i-1] {
			// in digits[0:i], find the minimal value that > digits[i], digits[0:i] is non-decreasing
			j := 0
			for j = 0; j < i; j++ {
				if digits[j] > digits[i] {
					break
				}
			}
			digits[i], digits[j] = digits[j], digits[i] // swap i, j
			// digits[:i] is still sorted, just reverse it
			for x, y := 0, i-1; x < y; x, y = x+1, y-1 {
				digits[x], digits[y] = digits[y], digits[x]
			}
			// make the number, check if in int32
			x := int64(0)
			for j := len(digits) - 1; j >= 0; j-- {
				x *= 10
				x += int64(digits[j])
			}
			if x <= (1<<31)-1 {
				return int(x)
			}
			return -1
		}
	}
	return -1
}

func main() {
	for _, v := range []struct {
		n, ans int
	}{
		{12, 21},
		{45732, 47235},
		{41732, 42137},
		{45762, 46257},
		{21, -1},
		{221, -1},
		{3, -1},
		{723, 732},
	} {
		fmt.Println(nextGreaterElement(v.n), v.ans)
	}
}

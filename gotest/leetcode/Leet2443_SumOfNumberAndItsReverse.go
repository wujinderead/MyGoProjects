package main

import "fmt"

// https://leetcode.com/problems/sum-of-number-and-its-reverse/

// Given a non-negative integer num, return true if num can be expressed as the sum of any
// non-negative integer and its reverse, or false otherwise.
// Example 1:
//   Input: num = 443
//   Output: true
//   Explanation: 172 + 271 = 443 so we return true.
// Example 2:
//   Input: num = 63
//   Output: false
//   Explanation: 63 cannot be expressed as the sum of a non-negative integer and
//     its reverse so we return false.
// Example 3:
//   Input: num = 181
//   Output: true
//   Explanation: 140 + 041 = 181 so we return true. Note that when a number is
//     reversed, there may be leading zeros.
// Constraints:
//   0 <= num <= 10⁵

// brute force, from n/2 to n
func sumOfNumberAndReverse(num int) bool {
	for i := num / 2; i <= num; i++ {
		v := i
		r := 0
		for v > 0 {
			r = r*10 + v%10
			v /= 10
		}

		if i+r == num {
			return true
		}
	}
	return false
}

func main() {
	for _, v := range []struct {
		nums int
		ans  bool
	}{
		{443, true},
		{63, false},
		{181, true},
		{1, false},
		{2, true},
	} {
		fmt.Println(sumOfNumberAndReverse(v.nums), v.ans)
	}
}

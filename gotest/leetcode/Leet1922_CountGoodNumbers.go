package main

import "fmt"

// https://leetcode.com/problems/count-good-numbers/

// A digit string is good if the digits (0-indexed) at even indices are even and
// the digits at odd indices are prime (2, 3, 5, or 7).
// For example, "2582" is good because the digits (2 and 8) at even positions are even and the
// digits (5 and 2) at odd positions are prime. However, "3245" is not good because 3 is at an
// even index but is not even.
// Given an integer n, return the total number of good digit strings of length n. Since the answer
// may be large, return it modulo 10^9 + 7.
// A digit string is a string consisting of digits 0 through 9 that may contain
// leading zeros.
// Example 1:
//   Input: n = 1
//   Output: 5
//   Explanation: The good numbers of length 1 are "0", "2", "4", "6", "8".
// Example 2:
//   Input: n = 4
//   Output: 400
// Example 3:
//   Input: n = 50
//   Output: 564908303
// Constraints:
//   1 <= n <= 10^15

func countGoodNumbers(n int64) int {
	ans := 1
	pow := 20
	mod := int(1e9 + 7)
	x := n / 2 // calculate 20^(n/2)
	for x > 0 {
		if x%2 == 1 {
			ans *= pow
			ans = ans % mod
		}
		x = x / 2
		pow = pow * pow
		pow = pow % mod
	}
	if n%2 == 1 {
		ans *= 5
		ans %= mod
	}
	return ans
}

func main() {
	for _, v := range []struct {
		n   int64
		ans int
	}{
		{1, 5},
		{4, 400},
		{5, 8000},
		{50, 564908303},
	} {
		fmt.Println(countGoodNumbers(v.n), v.ans)
	}
}

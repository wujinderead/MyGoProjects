package main

import "fmt"

// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/

// Given an integer n, return the decimal value of the binary string formed by concatenating
// the binary representations of 1 to n in order, modulo 10^9 + 7.
// Example 1:
//   Input: n = 1
//   Output: 1
//   Explanation: "1" in binary corresponds to the decimal value 1.
// Example 2:
//   Input: n = 3
//   Output: 27
//   Explanation: In binary, 1, 2, and 3 corresponds to "1", "10", and "11".
//     After concatenating them, we have "11011", which corresponds to the decimal value 27.
// Example 3:
//   Input: n = 12
//   Output: 505379714
//   Explanation: The concatenation results in "1101110010111011110001001101010111100".
//     The decimal value of that is 118505380540.
//     After modulo 109 + 7, the result is 505379714.
// Constraints:
//   1 <= n <= 10^5

// just calculate it, e.g for 4 <= i <= 7, ans = (ans << 3)+i
func concatenatedBinary(n int) int {
	ans := 1
	mod := int(1e9 + 7)
	var shift uint
	for shift = 1; (1 << shift) <= n; shift++ {
		i := 1 << shift
		for i <= n && i < 1<<(shift+1) {
			ans = (ans << (shift + 1)) + i
			ans = ans % mod
			i++
		}
		if i > n {
			break
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		n, ans int
	}{
		{1, 1},
		{2, 6},
		{3, 27},
		{4, 220},
		{5, 1765},
		{8, 1808248},
		{9, 28931977},
		{12, 505379714},
	} {
		fmt.Println(concatenatedBinary(v.n), v.ans)
	}
}

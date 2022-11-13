package main

import "fmt"

// https://leetcode.com/problems/count-ways-to-build-good-strings/

// Given the integers zero, one, low, and high, we can construct a string by
// starting with an empty string, and then at each step perform either of the following:
//   Append the character '0' zero times.
//   Append the character '1' one times.
// This can be performed any number of times.
// A good string is a string constructed by the above process having a length between low and
// high (inclusive).
// Return the number of different good strings that can be constructed satisfying these properties.
// Since the answer can be large, return it modulo 10⁹ + 7.
// Example 1:
//   Input: low = 3, high = 3, zero = 1, one = 1
//   Output: 8
//   Explanation:
//     One possible valid good string is "011".
//     It can be constructed as follows: "" -> "0" -> "01" -> "011".
//     All binary strings from "000" to "111" are good strings in this example.
// Example 2:
//   Input: low = 2, high = 3, zero = 1, one = 2
//   Output: 5
//   Explanation: The good strings are "00", "11", "000", "110", and "011".
// Constraints:
//   1 <= low <= high <= 10⁵
//   1 <= zero, one <= low

func countGoodStrings(low int, high int, zero int, one int) int {
	// dp[i] = dp[i-zero]+dp[i-one]
	dp := make([]int, high+1)
	dp[0] = 1
	const P = int(1e9 + 7)
	ans := 0
	for i := 1; i <= high; i++ {
		if i-zero >= 0 {
			dp[i] += dp[i-zero]
		}
		if i-one >= 0 {
			dp[i] += dp[i-one]
		}
		if i >= low {
			ans = (ans + dp[i]) % P
		}
		dp[i] %= P
	}
	return ans
}

func main() {
	for _, v := range []struct {
		low, high, zero, one, ans int
	}{
		{3, 3, 1, 1, 8},
		{2, 3, 1, 2, 5},
	} {
		fmt.Println(countGoodStrings(v.low, v.high, v.zero, v.one), v.ans)
	}
}

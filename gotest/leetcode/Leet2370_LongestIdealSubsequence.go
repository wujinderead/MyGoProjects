package main

import "fmt"

// https://leetcode.com/problems/longest-ideal-subsequence/

// You are given a string s consisting of lowercase letters and an integer k. We
// call a string t ideal if the following conditions are satisfied:
//   t is a subsequence of the string s.
//   The absolute difference in the alphabet order of every two adjacent letters in
//     t is less than or equal to k.
// Return the length of the longest ideal string.
// A subsequence is a string that can be derived from another string by deleting some or no
// characters without changing the order of the remaining characters.
// Note that the alphabet order is not cyclic. For example, the absolute difference in the
// alphabet order of 'a' and 'z' is 25, not 1.
// Example 1:
//   Input: s = "acfgbd", k = 2
//   Output: 4
//   Explanation: The longest ideal string is "acbd". The length of this string is 4, so 4 is returned.
//     Note that "acfgbd" is not ideal because 'c' and 'f' have a difference of 3 in alphabet order.
// Example 2:
//   Input: s = "abcd", k = 3
//   Output: 4
//   Explanation: The longest ideal string is "abcd". The length of this string is 4,
//     so 4 is returned.
// Constraints:
//   1 <= s.length <= 10⁵
//   0 <= k <= 25
//   s consists of lowercase English letters.

func longestIdealString(s string, k int) int {
	ss := []byte(s)
	dp := make([]int, 26) // dp[i] the length of LongestIdealString that ends with char[i]
	ans := 0
	for i := range ss {
		ci := int(ss[i] - 'a')
		m := 1
		// dp[i] = max(dp[i-k], ..., dp[i+k]) + 1
		for j := max(0, ci-k); j <= min(26, ci+k); j++ {
			m = max(m, dp[j]+1)
		}
		dp[ci] = m
		ans = max(ans, m)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		s      string
		k, ans int
	}{
		{"abcd", 3, 4},
		{"acfgbd", 2, 4},
	} {
		fmt.Println(longestIdealString(v.s, v.k), v.ans)
	}
}

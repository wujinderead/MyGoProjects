package main

import "fmt"

// https://leetcode.com/problems/maximum-number-of-non-overlapping-palindrome-substrings/

// You are given a string s and a positive integer k.
// Select a set of non-overlapping substrings from the string s that satisfy the following conditions:
//   The length of each substring is at least k.
//   Each substring is a palindrome.
// Return the maximum number of substrings in an optimal selection.
// A substring is a contiguous sequence of characters within a string.
// Example 1:
//   Input: s = "abaccdbbd", k = 3
//   Output: 2
//   Explanation: We can select the substrings underlined in s = "abaccdbbd". Both
//     "aba" and "dbbd" are palindromes and have a length of at least k = 3.
//     It can be shown that we cannot find a selection with more than two valid substrings.
// Example 2:
//   Input: s = "adbcda", k = 2
//   Output: 0
//   Explanation: There is no palindrome substring of length at least 2 in the string.
// Constraints:
//   1 <= k <= s.length <= 2000
//   s consists of lowercase English letters.

func maxPalindromes(s string, k int) int {
	if k == 1 {
		return len(s)
	}
	short := make([]int, len(s)) // short[i] the shortest palindrome length that ends at i and >= k
	for i := 0; i < len(s); i++ {
		for l := 1; i-l >= 0 && i+l < len(s); l++ {
			if s[i-l] != s[i+l] {
				break
			}
			if 2*l+1 >= k && (short[i+l] == 0 || 2*l+1 < short[i+l]) {
				short[i+l] = 2*l + 1
			}
		}
	}
	for i := 0; i < len(s)-1; i++ {
		for l := 0; i-l >= 0 && i+1+l < len(s); l++ {
			if s[i-l] != s[i+1+l] {
				break
			}
			if 2*(l+1) >= k && (short[i+l+1] == 0 || 2*(l+1) < short[i+l+1]) {
				short[i+l+1] = 2 * (l + 1)
			}
		}
	}
	dp := make([]int, len(s))
	getDp := func(i int) int { // dp[-1]=0
		if i >= 0 {
			return dp[i]
		}
		return 0
	}
	for i := 1; i < len(s); i++ {
		// s[(i-short[i]+1)...i] is palindrome, so dp[i-short[i]]+1 is a candidate
		if short[i] > 0 && getDp(i-short[i])+1 > dp[i-1] {
			dp[i] = getDp(i-short[i]) + 1
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(s)-1]
}

func main() {
	for _, v := range []struct {
		s      string
		k, ans int
	}{
		{"abaccdbbd", 3, 2},
		{"adbcda", 2, 0},
		{"zmmmmz", 2, 2},
		{"zqzogfurlfmrnlffuipuupidkfhkggkhdrzezghwziopoinnsdkwkymhygonbiizmmmmzjhmyczzlz", 2, 12},
	} {
		fmt.Println(maxPalindromes(v.s, v.k), v.ans)
	}
}

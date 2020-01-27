package main

import "fmt"

// https://leetcode.com/problems/longest-common-subsequence

// Given two strings text1 and text2, return the length of their longest common subsequence.
// A subsequence of a string is a new string generated from the original string with some characters
// (can be none) deleted without changing the relative order of the remaining characters.
// (eg, "ace" is a subsequence of "abcde" while "aec" is not). A common subsequence of two strings
// is a subsequence that is common to both strings.
// If there is no common subsequence, return 0.
// Example 1:
//   Input: text1 = "abcde", text2 = "ace"
//   Output: 3
//   Explanation: The longest common subsequence is "ace" and its length is 3.
// Example 2:
//   Input: text1 = "abc", text2 = "abc"
//   Output: 3
//   Explanation: The longest common subsequence is "abc" and its length is 3.
// Example 3:
//   Input: text1 = "abc", text2 = "def"
//   Output: 0
//   Explanation: There is no such common subsequence, so the result is 0.
// Constraints:
//   1 <= text1.length <= 1000
//   1 <= text2.length <= 1000
//   The input strings consist of lowercase English characters only.

func longestCommonSubsequence(a string, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	if len(a) < len(b) {
		a, b = b, a // let b be shorter
	}
	// let lcs(i, j) be the lcs of a[0...i] and b[0...j], then
	// if a[i]==b[j], lcs(i, j) = lcs(i-1, j-1); if not, lcs(i, j) = max(lcs(i-1, j), lcs(i, j-1))
	lcs := make([]int, len(b))

	// make line 0
	if a[0] == b[0] {
		lcs[0] = 1
	}
	for j := 1; j < len(b); j++ {
		lcs[j] = lcs[j-1]
		if a[0] == b[j] {
			lcs[j] = 1
		}
	}

	// dp
	for i := 1; i < len(a); i++ {
		tmp := lcs[0] // to store lcs(i-1, j-1)
		if a[i] == b[0] {
			lcs[0] = 1
		}
		for j := 1; j < len(b); j++ {
			tmp2 := lcs[j]
			if a[i] == b[j] {
				lcs[j] = tmp + 1
			} else {
				lcs[j] = max(lcs[j-1], lcs[j])
			}
			tmp = tmp2
		}
	}
	return lcs[len(b)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestCommonSubsequence("xabxac", "abcabxabcd"))
	fmt.Println(longestCommonSubsequence("xabxaabxa", "babxba"))
	fmt.Println(longestCommonSubsequence("abede", "eghie"))
	fmt.Println(longestCommonSubsequence("pqrst", "uvwxyz"))
	fmt.Println(longestCommonSubsequence("a", "bcde"))
	fmt.Println(longestCommonSubsequence("a", "bcade"))
	fmt.Println(longestCommonSubsequence("adsd", "a"))
	fmt.Println(longestCommonSubsequence("a", "ab"))
	fmt.Println(longestCommonSubsequence("a", ""))
}

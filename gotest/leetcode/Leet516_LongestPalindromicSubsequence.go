package main

import "fmt"

// https://leetcode.com/problems/longest-palindromic-subsequence/

// Given a string s, find the longest palindromic subsequence's length in s.
// You may assume that the maximum length of s is 1000.
// Example 1:
//   Input: "bbbab"
//   Output: 4
//   One possible longest palindromic subsequence is "bbbb".
// Example 2:
//   Input: "cbbd"
//   Output: 2
//   One possible longest palindromic subsequence is "bb".

func longestPalindromeSubseq(s string) int {
	// let lps(i, j) be the length of the longest palindromic subseq of s[i:j],
	// then if s[i]==s[j], lps(i, j) = lps(i+1, j-1)+2;
	// if s[i] != s[j], lps(i, j) = max( lps(i+1, j), lps(i, j-1))
	// base case lps(i, i)=1, lps(i, i+1)=2 if s[i]==s[i+1]
	lps := make(map[[2]int]int)
	for i := 0; i < len(s)-1; i++ {
		lps[[2]int{i, i}] = 1
		lps[[2]int{i, i + 1}] = 1
		if s[i] == s[i+1] {
			lps[[2]int{i, i + 1}] = 2
		}
	}
	lps[[2]int{len(s) - 1, len(s) - 1}] = 1
	for l := 2; l < len(s); l++ {
		for i := 0; i+l < len(s); i++ {
			if s[i] == s[i+l] {
				lps[[2]int{i, i + l}] = lps[[2]int{i + 1, i + l - 1}] + 2
			} else {
				lps[[2]int{i, i + l}] = max(lps[[2]int{i + 1, i + l}], lps[[2]int{i, i + l - 1}])
			}
		}
	}
	return lps[[2]int{0, len(s) - 1}]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestPalindromeSubseq(""))
	fmt.Println(longestPalindromeSubseq("a"))
	fmt.Println(longestPalindromeSubseq("ab"))
	fmt.Println(longestPalindromeSubseq("aba"))
	fmt.Println(longestPalindromeSubseq("abab"))
	fmt.Println(longestPalindromeSubseq("bbbab"))
	fmt.Println(longestPalindromeSubseq("cbbd"))
}

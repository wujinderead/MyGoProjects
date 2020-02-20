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

// n space
func longestCommonSubsequence(a string, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	if len(a) < len(b) {
		a, b = b, a // let b be shorter
	}
	// let pa[i] be the prefix of string a with length i,
	// i.e, pa[0]="", pa[1]=a[0...0], pa[2]=a[0...1], pa[i]=a[0...i-1]
	// let lcs(i, j) be the lcs of pa[i] and pb[j], then
	// if i*j==0, lcs(i, j)=0;
	// if a[i-1]==b[j-1], lcs(i, j)=lcs(i-1, j-1)+1;
	// if a[i-1]!=b[j-1], lcs(i, j)=max(lcs(i-1, j), lcs(i, j-1))
	lcs := make([]int, len(b)+1)

	// dp
	for i := 1; i <= len(a); i++ {
		tmp := lcs[0] // to store lcs(i-1, j-1)
		for j := 1; j <= len(b); j++ {
			tmp2 := lcs[j] // to store lcs(i-1, j)
			if a[i-1] == b[j-1] {
				lcs[j] = tmp + 1
			} else {
				lcs[j] = max(lcs[j-1], lcs[j])
			}
			tmp = tmp2
		}
	}
	return lcs[len(b)]
}

// 2n space
func longestCommonSubsequence1(a string, b string) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	if len(a) < len(b) {
		a, b = b, a // let b be shorter
	}
	// let pa[i] be the prefix of string a with length i,
	// i.e, pa[0]="", pa[1]=a[0...0], pa[2]=a[0...1], pa[i]=a[0...i-1]
	// let lcs(i, j) be the lcs of pa[i] and pb[j], then
	// if i*j==0, lcs(i, j)=0;
	// if a[i-1]==b[j-1], lcs(i, j)=lcs(i-1, j-1)+1;
	// if a[i-1]!=b[j-1], lcs(i, j)=max(lcs(i-1, j), lcs(i, j-1))
	prev := make([]int, len(b)+1)
	cur := make([]int, len(b)+1)

	// dp
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cur[j] = prev[j-1] + 1
			} else {
				cur[j] = max(prev[j], cur[j-1])
			}
		}
		prev, cur = cur, prev
	}
	return prev[len(b)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	f1 := longestCommonSubsequence
	f2 := longestCommonSubsequence1
	for _, v := range [][]string{
		{"xabxac", "abcabxabcd"},
		{"xabxaabxa", "babxba"},
		{"abede", "eghie"},
		{"pqrst", "uvwxyz"},
		{"a", "bcde"},
		{"a", "bcade"},
		{"adsd", "a"},
		{"a", "ab"},
		{"a", ""},
	} {
		fmt.Println(f1(v[0], v[1]), f2(v[0], v[1]))
	}

}

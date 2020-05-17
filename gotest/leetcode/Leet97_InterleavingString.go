package main

import (
	"fmt"
)

// https://leetcode.com/problems/interleaving-string/

// Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.
// Example 1:
//   Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//   Output: true
// Example 2:
//   Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//   Output: false

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	// let dp(i, j) be if s1[:i] and s2[:j] can interleave s3[:i+j], then
	// if s1[i]==s3_last_letter, dp(i, j) = dp(i-1, j)
	// if s2[j]==s3_last_letter, dp(i, j) = dp(i, j-1)
	// if s2[i]==s2[j]==s3_last_letter, dp(i, j) = dp(i-1, j) || dp(i, j-1)
	// if s2[i]!=s2[j]!=s3_last_letter, dp(i, j) = false
	// base case: dp(i, 0) = s1[:i]==s3[:i]
	old, new := make([]bool, len(s2)+1), make([]bool, len(s2)+1)
	old[0] = true
	for j:=1; j<=len(s2); j++ {
		old[j] = old[j-1] && s2[j-1]==s3[j-1]   // dp(0, j)
	}
	//fmt.Println(old)
	for i:=1; i<=len(s1); i++ {
		new[0] = old[0] && s1[i-1]==s3[i-1]
		for j:=1; j<=len(s2); j++ {
			if s1[i-1] == s2[j-1] && s1[i-1] == s3[i+j-1] {
				new[j] = old[j] || new[j-1]
				continue
			}
			if s1[i-1] == s3[i+j-1] {
				new[j] = old[j]
				continue
			}
			if s2[j-1] == s3[i+j-1] {
				new[j] = new[j-1]
				continue
			}
			new[j] = false
		}
		//fmt.Println(new)
		old, new = new, old
	}
	return old[len(s2)]
}

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	fmt.Println(isInterleave("abc", "", "abc"), true)
	fmt.Println(isInterleave("", "abc", "abc"), true)
	fmt.Println(isInterleave("", "abc", "ab"), false)
	fmt.Println(isInterleave("ab", "c", "ab"), false)
	fmt.Println(isInterleave("ab", "c", "abc"), true)
	fmt.Println(isInterleave("ac", "b", "abc"), true)
	fmt.Println(isInterleave("ac", "b", "bca"), false)
	fmt.Println(isInterleave("abc", "ade", "aabcdf"), false)
	fmt.Println(isInterleave("abc", "ade", "aabcde"), true)
	fmt.Println(isInterleave("abc", "ade", "abdace"), false)
	fmt.Println(isInterleave("", "a", "a"), true)
	fmt.Println(isInterleave("", "a", "c"), false)
	fmt.Println(isInterleave("", "", ""), true)
}
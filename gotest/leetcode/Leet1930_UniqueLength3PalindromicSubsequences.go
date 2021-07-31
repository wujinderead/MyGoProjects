package main

import "fmt"

// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/

// Given a string s, return the number of unique palindromes of length three that are a subsequence of s.
// Note that even if there are multiple ways to obtain the same subsequence, it is still only counted once.
// A palindrome is a string that reads the same forwards and backwards.
// A subsequence of a string is a new string generated from the original string with some characters
// (can be none) deleted without changing the relative order of the remaining characters. //
// For example, "ace" is a subsequence of "abcde".
// Example 1:
//   Input: s = "aabca"
//   Output: 3
//   Explanation: The 3 palindromic subsequences of length 3 are:
//     - "aba" (subsequence of "aabca")
//     - "aaa" (subsequence of "aabca")
//     - "aca" (subsequence of "aabca")
// Example 2:
//   Input: s = "adc"
//   Output: 0
//   Explanation: There are no palindromic subsequences of length 3 in "adc".
// Example 3:
//   Input: s = "bbcbaba"
//   Output: 4
//   Explanation: The 4 palindromic subsequences of length 3 are:
//     - "bbb" (subsequence of "bbcbaba")
//     - "bcb" (subsequence of "bbcbaba")
//     - "bab" (subsequence of "bbcbaba")
//     - "aba" (subsequence of "bbcbaba")
// Constraints:
//   3 <= s.length <= 10^5
//   s consists of only lowercase English letters.

func countPalindromicSubsequence(s string) int {
	count := 0
	var single [26]bool
	var double [26][26]bool
	var triple [26][26][26]bool
	for _, v := range s { // denote current char as C
		for i := 0; i < 26; i++ {
			if double[int(v-'a')][i] { // if CX exists in double, then CXC forms a palindrome
				if !triple[int(v-'a')][i][int(v-'a')] {
					triple[int(v-'a')][i][int(v-'a')] = true // add CXC to triple
					count++
				}
			}
		}
		for i := 0; i < 26; i++ { // if X exists, add XC to double
			if single[i] {
				double[i][int(v-'a')] = true
			}
		}
		single[int(v-'a')] = true // add C to single
	}
	return count
}

func main() {
	for _, v := range []struct {
		s   string
		ans int
	}{
		{"aabca", 3},
		{"abc", 0},
		{"bbcbaba", 4},
	} {
		fmt.Println(countPalindromicSubsequence(v.s), v.ans)
	}
}

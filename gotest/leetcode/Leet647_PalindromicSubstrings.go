package main

import "fmt"

// https://leetcode.com/problems/palindromic-substrings/

// Given a string, your task is to count how many palindromic substrings in this string.
// The substrings with different start indexes or end indexes are counted as different
// substrings even they consist of same characters.
// Example 1:
//   Input: "abc"
//   Output: 3
//   Explanation: Three palindromic strings: "a", "b", "c".
// Example 2:
//   Input: "aaa"
//   Output: 6
//   Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
// Note:
//   The input string length won't exceed 1000.

func countSubstrings(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	count := 0
	for i := 0; i < len(s); i++ {
		count++ // single character
		for k := 1; i+k < len(s) && i-k >= 0; k++ {
			if s[i+k] != s[i-k] {
				break
			}
			count++
		}
	}
	for i := 0; i < len(s)-1; i++ {
		for k := 0; i-k >= 0 && i+1+k < len(s); k++ {
			if s[i-k] != s[i+1+k] {
				break
			}
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}

package main

import "fmt"

// https://leetcode.com/problems/shortest-palindrome/

// Given a string s, you are allowed to convert it to a palindrome by adding characters
// in front of it. Find and return the shortest palindrome you can find by performing
// this transformation.
// Example 1:
//   Input: "aacecaaa"
//   Output: "aaacecaaa"
// Example 2:
//   Input: "abcd"
//   Output: "dcbabcd"

// IMPROVEMENT: kmp O(n) method: https://leetcode.com/problems/shortest-palindrome/solution/

func shortestPalindrome(s string) string {
	// find longest palindromic prefix
	var j int
outer:
	for j = len(s) - 1; j > 0; j-- {
		i := 0
		k := j
		for i < k {
			if s[i] != s[k] {
				continue outer
			}
			i++
			k--
		}
		break // found max j that s[i...j] is palindrome
	}
	// s[i...j] is palindrome
	buf := make([]byte, len(s)+len(s)-len(s[:j+1]))
	copy(buf[len(s)-len(s[:j+1]):], s)
	i := 0
	for k := len(s) - 1; k > j; k-- {
		buf[i] = s[k]
		i++
	}
	return string(buf)
}

func main() {
	fmt.Println(shortestPalindrome("aacecaaaf"))
	fmt.Println(shortestPalindrome("abcd"))
	fmt.Println(shortestPalindrome("a"))
	fmt.Println(shortestPalindrome("ab"))
	fmt.Println(shortestPalindrome("aa"))
	fmt.Println(shortestPalindrome("aab"))
}

package main

import "fmt"

// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/

// You are given two strings s and t consisting of only lowercase English letters.
// Return the minimum number of characters that need to be appended to the end of s so that t becomes
// a subsequence of s.
// A subsequence is a string that can be derived from another string by deleting some or no characters
// without changing the order of the remaining characters.
// Example 1:
//   Input: s = "coaching", t = "coding"
//   Output: 4
//   Explanation: Append the characters "ding" to the end of s so that s = "coachingding".
//     Now, t is a subsequence of s ("coachingding").
//     It can be shown that appending any 3 characters to the end of s will never make t a subsequence.
// Example 2:
//   Input: s = "abcde", t = "a"
//   Output: 0
//   Explanation: t is already a subsequence of s ("abcde").
// Example 3:
//   Input: s = "z", t = "abcde"
//   Output: 5
//   Explanation: Append the characters "abcde" to the end of s so that s = "zabcde".
//     Now, t is a subsequence of s ("zabcde").
//     It can be shown that appending any 4 characters to the end of s will never make t a subsequence.
// Constraints:
//   1 <= s.length, t.length <= 10⁵
//   s and t consist only of lowercase English letters.

// a short code:
/*
	i = j = 0
	while i < len(s) and j < len(t):  # find prefix subsequence
		j += s[i] == t[j]
		i += 1
	return len(t) - j
*/
func appendCharacters(s string, t string) int {
	j := 0
	i := 0
outer:
	for i < len(t) {
		for j < len(s) {
			if t[i] == s[j] {
				i++
				j++
				continue outer
			}
			j++
		}
		break
	}
	return len(t) - i
}

func main() {
	for _, v := range []struct {
		s, t string
		ans  int
	}{
		{"coaching", "coding", 4},
		{"z", "abcde", 5},
		{"abcde", "a", 0},
		{"abcde", "abcd", 0},
		{"abcde", "abcde", 0},
		{"abcde", "abcdef", 1},
	} {
		fmt.Println(appendCharacters(v.s, v.t), v.ans)
	}
}

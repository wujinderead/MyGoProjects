package main

import "fmt"

// https://leetcode.com/problems/maximize-palindrome-length-from-subsequences/

// You are given two strings, word1 and word2. You want to construct a string in
// the following manner:
//   Choose some non-empty subsequence subsequence1 from word1.
//   Choose some non-empty subsequence subsequence2 from word2.
//   Concatenate the subsequences: subsequence1 + subsequence2, to make the string.
// Return the length of the longest palindrome that can be constructed in the
// described manner. If no palindromes can be constructed, return 0.
// A subsequence of a string s is a string that can be made by deleting some (possibly
// none) characters from s without changing the order of the remaining characters.
// A palindrome is a string that reads the same forward as well as backward.
// Example 1:
//   Input: word1 = "cacb", word2 = "cbba"
//   Output: 5
//   Explanation: Choose "ab" from word1 and "cba" from word2 to make "abcba",
//     which is a palindrome.
// Example 2:
//   Input: word1 = "ab", word2 = "ab"
//   Output: 3
//   Explanation: Choose "ab" from word1 and "a" from word2 to make "aba",
//     which is a palindrome.
// Example 3:
//   Input: word1 = "aa", word2 = "bb"
//   Output: 0
//   Explanation: You cannot construct a palindrome from the described method, so return 0.
// Constraints:
//   1 <= word1.length, word2.length <= 1000
//   word1 and word2 consist of lowercase English letters.

// just concatenate word1 and word2, use the normal way to get the longest palindromic subsequence
// the only thing need concern is that both subsequence of two words are non-empty
func longestPalindrome(word1 string, word2 string) int {
	s := word1 + word2
	ans := 0

	// dp(j, i) is the lps of s[i...j]
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, i+1)
		dp[i][i] = 1
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			dp[i+1][i] = 2
			if i < len(word1) && i+1 >= len(word1) {
				ans = 2
			}
		} else {
			dp[i+1][i] = 1
		}
	}
	for diff := 2; diff < len(s); diff++ {
		for i := 0; i+diff < len(s); i++ {
			j := i + diff
			if s[i] == s[j] {
				dp[j][i] = dp[j-1][i+1] + 2
				// update max length only when the both side are in two words
				// i.e. s[i] in word1 and s[j] in word2
				if i < len(word1) && j >= len(word1) {
					ans = max(dp[j][i], ans)
				}
			} else {
				dp[j][i] = max(dp[j-1][i], dp[j][i+1])
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		w1, w2 string
		ans    int
	}{
		{"cacb", "cbba", 5},
		{"ab", "ab", 3},
		{"aa", "bb", 0},
		{"rhuzwqohquamvsz", "kvunbxje", 7},
		{"a", "abccb", 2},
	} {
		fmt.Println(longestPalindrome(v.w1, v.w2), v.ans)
	}
}

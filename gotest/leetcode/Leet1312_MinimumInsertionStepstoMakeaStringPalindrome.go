package main

import "fmt"

// https://leetcode.com/problems/minimum-insertion-steps-to-make-a-string-palindrome/

// Given a string s. In one step you can insert any character at any index of the string.
// Return the minimum number of steps to make s palindrome.
// A Palindrome String is one that reads the same backward as well as forward.
// Example 1:
//   Input: s = "zzazz"
//   Output: 0
//   Explanation: The string "zzazz" is already palindrome we don't need any insertions.
// Example 2:
//   Input: s = "mbadm"
//   Output: 2
//   Explanation: String can be "mbdadbm" or "mdbabdm".
// Example 3:
//   Input: s = "leetcode"
//   Output: 5
//   Explanation: Inserting 5 characters the string becomes "leetcodocteel".
// Example 4:
//   Input: s = "g"
//   Output: 0
// Example 5:
//   Input: s = "no"
//   Output: 1
// Constraints:
//   1 <= s.length <= 500
//   All characters of s are lower case English letters.

func minInsertions(s string) int {
	// find the longest palindromic sub-sequence, the result is len(s)-len(lps)
	// let lps(i,j) is the lps length of s[i...j],
	// then if s[i]==s[j], lsp(i, j)=lps(i+1,j-1)+2
	// base case lps(i, i)=1, lps(i,i+1)=2 if s[i]=s[i+1]
	lps := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		lps[i] = make([]int, i+1)
	}
	for i := 0; i < len(s)-1; i++ {
		lps[i][i] = 1
		lps[i+1][i] = 1
		if s[i] == s[i+1] {
			lps[i+1][i] = 2
		}
	}
	lps[len(s)-1][len(s)-1] = 1
	for l := 2; l < len(s); l++ {
		for j := 0; j+l < len(s); j++ {
			i := j + l
			if s[i] == s[j] {
				lps[i][j] = lps[i-1][j+1] + 2
			} else {
				lps[i][j] = max(lps[i-1][j], lps[i][j+1])
			}
		}
	}
	return len(s) - lps[len(s)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minInsertions("zzazz"))
	fmt.Println(minInsertions("mbadm"))
	fmt.Println(minInsertions("leetcode"))
	fmt.Println(minInsertions("g"))
	fmt.Println(minInsertions("no"))
}

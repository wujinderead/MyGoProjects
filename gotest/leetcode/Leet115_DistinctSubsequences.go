package main

import "fmt"

// https://leetcode.com/problems/distinct-subsequences/

// Given a string S and a string T, count the number of distinct subsequences of S which equals T.
// A subsequence of a string is a new string which is formed from the original string by
// deleting some (can be none) of the characters without disturbing the relative positions
// of the remaining characters. (ie, "ACE" is a subsequence of "ABCDE" while "AEC" is not).
// Example 1:
//   Input: S = "rabbbit", T = "rabbit"
//   Output: 3
//   Explanation:
//     As shown below, there are 3 ways you can generate "rabbit" from S.
//     (The caret symbol ^ means the chosen letters)
//
//     rabbbit
//     ^^^^ ^^
//     rabbbit
//     ^^ ^^^^
//     rabbbit
//     ^^^ ^^^
// Example 2:
//   Input: S = "babgbag", T = "bag"
//   Output: 5
//   Explanation:
//     As shown below, there are 5 ways you can generate "bag" from S.
//     (The caret symbol ^ means the chosen letters)
//
//     babgbag
//     ^^ ^
//     babgbag
//     ^^    ^
//     babgbag
//     ^    ^^
//     babgbag
//       ^  ^^
//     babgbag
//         ^^^

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	// time O(ST), space O(S)
	prev := make([]int, len(s)) // use len(s)+1 to make it more concise
	now := make([]int, len(s))
	if s[0] == t[0] {
		prev[0] = 1
	}
	for i := 1; i < len(s); i++ {
		prev[i] = prev[i-1]
		if s[i] == t[0] {
			prev[i]++
		}
	}
	//fmt.Println(prev)
	for j := 1; j < len(t); j++ {
		char := t[j]
		now[j-1] = 0
		for i := j; i < len(s); i++ {
			now[i] = now[i-1]
			if s[i] == char {
				now[i] += prev[i-1]
			}
		}
		prev, now = now, prev
		//fmt.Println(prev)
	}
	return prev[len(s)-1]
}

func main() {
	fmt.Println(numDistinct("tbabgbag", "bag"))
	fmt.Println(numDistinct("bag", "bag"))
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}

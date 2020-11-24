package main

import "fmt"

// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram/

// Given two equal-size strings s and t. In one step you can choose any character of t and
// replace it with another character.
// Return the minimum number of steps to make t an anagram of s.
// An Anagram of a string is a string that contains the same characters with a different
// (or the same) ordering.
// Example 1:
//   Input: s = "bab", t = "aba"
//   Output: 1
//   Explanation: Replace the first 'a' in t with b, t = "bba" which is anagram of s.
// Example 2:
//   Input: s = "leetcode", t = "practice"
//   Output: 5
//   Explanation: Replace 'p', 'r', 'a', 'i' and 'c' from t with proper characters to make t anagram of s.
// Example 3:
//   Input: s = "anagram", t = "mangaar"
//   Output: 0
//   Explanation: "anagram" and "mangaar" are anagrams.
// Example 4:
//   Input: s = "xxyyzz", t = "xxyyzz"
//   Output: 0
// Example 5:
//   Input: s = "friend", t = "family"
//   Output: 4
// Constraints:
//   1 <= s.length <= 50000
//   s.length == t.length
//   s and t contain lower-case English letters only.

// sum the differences of occurrence, the answer is sum/2
// e.g., for s="aa" t="ab", occurrence of letters is (a2, b0), (a1, b1)
// sum is |2-1| + |0-1| = 2, so answer is 1 (change a to b)
func minSteps(s string, t string) int {
	sc, tc := [26]int{}, [26]int{}
	for i := range s {
		sc[int(s[i]-'a')]++
		tc[int(t[i]-'a')]++
	}
	ans := 0
	for i := 0; i < 26; i++ {
		if sc[i] > tc[i] {
			ans += sc[i] - tc[i]
		} else {
			ans += tc[i] - sc[i]
		}
	}
	return ans / 2
}

func main() {
	for _, v := range []struct {
		s, t string
		ans  int
	}{
		{"aba", "bab", 1},
		{"leetcode", "practice", 5},
		{"anagram", "mangaar", 0},
		{"xxyyzz", "xxyyzz", 0},
		{"friend", "family", 4},
	} {
		fmt.Println(minSteps(v.s, v.t), v.ans)
	}
}

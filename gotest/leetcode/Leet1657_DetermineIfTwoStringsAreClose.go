package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/determine-if-two-strings-are-close/

// Two strings are considered close if you can attain one from the other using the following operations:
// Operation 1: Swap any two existing characters.
//   For example, abcde -> aecdb
// Operation 2: Transform every occurrence of one existing character into another existing character,
//   and do the same with the other character. For example, aacabb -> bbcbaa (all a's turn into b's,
//   and all b's turn into a's)
// You can use the operations on either string as many times as necessary.
// Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.
// Example 1:
//   Input: word1 = "abc", word2 = "bca"
//   Output: true
//   Explanation: You can attain word2 from word1 in 2 operations.
//     Apply Operation 1: "abc" -> "acb"
//     Apply Operation 1: "acb" -> "bca"
// Example 2:
//   Input: word1 = "a", word2 = "aa"
//   Output: false
//   Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.
// Example 3:
//   Input: word1 = "cabbba", word2 = "abbccc"
//   Output: true
//   Explanation: You can attain word2 from word1 in 3 operations.
//     Apply Operation 1: "cabbba" -> "caabbb"
//     Apply Operation 2: "caabbb" -> "baaccc"
//     Apply Operation 2: "baaccc" -> "abbccc"
// Example 4:
//   Input: word1 = "cabbba", word2 = "aabbss"
//   Output: false
//   Explanation: It is impossible to attain word2 from word1, or vice versa, in any amount of operations.
// Constraints:
//   1 <= word1.length, word2.length <= 10^5
//   word1 and word2 contain only lowercase English letters.

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	c1, c2 := make([]int, 26), make([]int, 26)
	o1, o2 := make([]byte, 26), make([]byte, 26)
	for i := range word1 {
		c1[int(word1[i]-'a')]++
		c2[int(word2[i]-'a')]++
		o1[int(word1[i]-'a')] = 1
		o2[int(word2[i]-'a')] = 1
	}
	sort.Ints(c1)
	sort.Ints(c2)
	for i := range c1 {
		if c1[i] != c2[i] {
			return false
		}
	}
	return string(o1) == string(o2)
}

func main() {
	for _, v := range []struct {
		w1, w2 string
		ans    bool
	}{
		{"abc", "bca", true},
		{"a", "aa", false},
		{"cabbba", "abbccc", true},
		{"sabbba", "abbccc", false},
		{"cabbba", "aabbss", false},
	} {
		fmt.Println(closeStrings(v.w1, v.w2), v.ans)
	}
}

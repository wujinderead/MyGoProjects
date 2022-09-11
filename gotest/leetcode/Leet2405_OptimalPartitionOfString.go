package main

import "fmt"

// https://leetcode.com/problems/optimal-partition-of-string/

// Given a string s, partition the string into one or more substrings such that the characters in each
// substring are unique. That is, no letter appears in a single substring more than once.
// Return the minimum number of substrings in such a partition.
// Note that each character should belong to exactly one substring in a partition.
// Example 1:
//   Input: s = "abacaba"
//   Output: 4
//   Explanation:
//     Two possible partitions are ("a","ba","cab","a") and ("ab","a","ca","ba").
//     It can be shown that 4 is the minimum number of substrings needed.
// Example 2:
//   Input: s = "ssssss"
//   Output: 6
//   Explanation:
//     The only valid partition is ("s","s","s","s","s","s").
// Constraints:
//   1 <= s.length <= 10⁵
//   s consists of only English lowercase letters.

// greedy
func partitionString(s string) int {
	occur := [26]bool{}
	count := 0
	for i := range s {
		if occur[int(s[i]-'a')] { // found a duplicated char, terminate current substring, add count
			occur = [26]bool{}
			count++
		}
		occur[int(s[i]-'a')] = true
	}
	return count + 1 // always add the last substring
}

func main() {
	for _, v := range []struct {
		s   string
		ans int
	}{
		{"ssssss", 6},
		{"abacaba", 4},
	} {
		fmt.Println(partitionString(v.s), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/partition-labels/

// A string S of lowercase letters is given. We want to partition this string into
// as many parts as possible so that each letter appears in at most one part,
// and return a list of integers representing the size of these parts.
// Example 1:
//   Input: S = "ababcbacadefegdehijhklij"
//   Output: [9,7,8]
//   Explanation:
//     The partition is "ababcbaca", "defegde", "hijhklij".
//     This is a partition so that each letter appears in at most one part.
//     A partition like "ababcbacadefegde", "hijhklij" is incorrect,
//     because it splits S into less parts.
// Note:
//   S will have length in range [1, 500].
//   S will consist of lowercase letters ('a' to 'z') only.

func partitionLabels(S string) []int {
	last := make([]int, 26)
	for i := range S {
		last[S[i]-'a'] = i // last occurrence of letter
	}
	ans := make([]int, 0)
	start, end := 0, 0 // start and end index f current partition
	for i := range S {
		end = max(end, last[S[i]-'a'])
		if i == end { // current partition should stop
			ans = append(ans, end-start+1)
			start = i + 1 // start a new partition
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
	fmt.Println(partitionLabels(""))
	fmt.Println(partitionLabels("a"))
	fmt.Println(partitionLabels("aba"))
	fmt.Println(partitionLabels("abcdefghijk"))
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
}

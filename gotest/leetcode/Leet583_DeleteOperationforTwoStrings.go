package main

import "fmt"

// https://leetcode.com/problems/delete-operation-for-two-strings

// Given two words word1 and word2, find the minimum number of steps
// required to make word1 and word2 the same, where in each step you
// can delete one character in either string.
// Example 1:
//   Input: "sea", "eat"
//   Output: 2
//   Explanation: You need one step to make "sea" to "ea" and another step to make "eat" to "ea".
// Note:
//   The length of given words won't exceed 500.
//   Characters in given words can only be lower-case letters.

func minDistance(word1 string, word2 string) int {
	if word1 == "" || word2 == "" {
		return len(word1) + len(word2)
	}
	// find the longest common sequence of two strings, the result is len(w1)-lcs+len(w2)-lcs
	// let lcs(i, j) be the longest common sequence between a[0...i] and b[0...j], then
	// if a[i]==b[j], lcs(i, j)=lcs(i-1, j-1)+1; if not, lcs(i, j)=max(lcs(i-1, j), lcs(i, j-1))
	// base case: if a[x]==b[0], lcs(x, 0)=1
	if len(word1) < len(word2) {
		word2, word1 = word1, word2
	}
	old, new := make([]int, len(word2)), make([]int, len(word2))
	if word1[0] == word2[0] {
		old[0] = 1
	}
	for j := 1; j < len(word2); j++ {
		if word1[0] == word2[j] {
			old[j] = 1
		} else {
			old[j] = old[j-1]
		}
	}
	fmt.Println(old)
	for i := 1; i < len(word1); i++ {
		if word1[i] == word2[0] {
			new[0] = 1
		} else {
			new[0] = old[0]
		}
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				new[j] = old[j-1] + 1
			} else {
				new[j] = max(old[j], new[j-1])
			}
		}
		fmt.Println(new)
		old, new = new, old
	}
	return len(word1) + len(word2) - old[len(word2)-1] - old[len(word2)-1]
	//return old[len(word2)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minDistance("xabxac", "abcabxabcd"))
	fmt.Println(minDistance("xabxaabxa", "babxba"))
	fmt.Println(minDistance("abede", "eghie"))
	fmt.Println(minDistance("pqrst", "uvwxyz"))
	fmt.Println(minDistance("a", "bcde"))
	fmt.Println(minDistance("a", "bcade"))
	fmt.Println(minDistance("adsd", "a"))
	fmt.Println(minDistance("a", "ab"))
}

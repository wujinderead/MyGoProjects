package main

import (
	"fmt"
)

// https://leetcode.com/problems/sum-of-prefix-scores-of-strings/

// You are given an array words of size n consisting of non-empty strings.
// We define the score of a string word as the number of strings words[i] such that word is a prefix
// of words[i].
// For example, if words = ["a", "ab", "abc", "cab"], then the score of "ab" is 2, since "ab" is a
// prefix of both "ab" and "abc".
// Return an array answer of size n where answer[i] is the sum of scores of every non-empty prefix
// of words[i].
// Note that a string is considered as a prefix of itself.
// Example 1:
//   Input: words = ["abc","ab","bc","b"]
//   Output: [5,4,3,2]
//   Explanation: The answer for each string is the following:
//     - "abc" has 3 prefixes: "a", "ab", and "abc".
//     - There are 2 strings with the prefix "a", 2 strings with the prefix "ab", and 1
//       string with the prefix "abc".
//       The total is answer[0] = 2 + 2 + 1 = 5.
//     - "ab" has 2 prefixes: "a" and "ab".
//     - There are 2 strings with the prefix "a", and 2 strings with the prefix "ab".
//       The total is answer[1] = 2 + 2 = 4.
//     - "bc" has 2 prefixes: "b" and "bc".
//     - There are 2 strings with the prefix "b", and 1 string with the prefix "bc".
//       The total is answer[2] = 2 + 1 = 3.
//     - "b" has 1 prefix: "b".
//     - There are 2 strings with the prefix "b".
//       The total is answer[3] = 2.
// Example 2:
//   Input: words = ["abcd"]
//   Output: [4]
//   Explanation:
//     "abcd" has 4 prefixes: "a", "ab", "abc", and "abcd".
//     Each prefix has a score of one, so the total is answer[0] = 1 + 1 + 1 + 1 = 4.
// Constraints:
//   1 <= words.length <= 1000
//   1 <= words[i].length <= 1000
//   words[i] consists of lowercase English letters.

func sumPrefixScores(words []string) []int {
	root := &trieNode{}
	// make a trie
	for _, w := range words {
		cur := root
		for _, c := range w {
			ci := int(c - 'a')
			if cur.children[ci] == nil {
				cur.children[ci] = &trieNode{
					count: 1, // count is how many occurrence of this prefix in the words
				}
			} else {
				cur.children[ci].count++
			}
			cur = cur.children[ci]
		}
	}
	// traverse words again
	ans := make([]int, len(words))
	for i, w := range words {
		cur := root
		for _, c := range w {
			ci := int(c - 'a')
			cur = cur.children[ci]
			ans[i] += cur.count
		}
	}
	return ans
}

type trieNode struct {
	children [26]*trieNode
	count    int
}

func main() {
	for _, v := range []struct {
		words []string
		ans   []int
	}{
		{[]string{"abc", "ab", "bc", "b"}, []int{5, 4, 3, 2}},
		{[]string{"abcd"}, []int{4}},
	} {
		fmt.Println(sumPrefixScores(v.words), v.ans)
	}
}

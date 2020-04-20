package main

import "fmt"

// https://leetcode.com/problems/palindrome-pairs/

// Given a list of unique words, find all pairs of distinct indices (i, j) in the given list,
// so that the concatenation of the two words, i.e. words[i] + words[j] is a palindrome.
// Example 1:
//   Input: ["abcd","dcba","lls","s","sssll"]
//   Output: [[0,1],[1,0],[3,2],[2,4]]
//   Explanation: The palindromes are ["dcbaabcd","abcddcba","slls","llssssll"]
// Example 2:
//   Input: ["bat","tab","cat"]
//   Output: [[0,1],[1,0]]
//   Explanation: The palindromes are ["battab","tabbat"]

func palindromePairs(words []string) [][]int {
	ans := make([][]int, 0)
	// we construct a trie. e.g string "abcdfff", "abcdyyh" in trie.
	// for string "dcba", we reverse it and match in trie, we got "abcd" is matched.
	// it means "dcba" is image of the prefix of "abcd....".
	// if "..." is palindromic, we can get palindromic "abcd...dcba".
	// "fff" is palindromic while "yyh" is not, thus we got palindromic "abcdfffdcba".
	root := &trieNode{end: -1}
	for i := range words {
		root.add(words[i], i)
	}
	stack := make([]*trieNode, 0)
outer:
	for i := range words {
		cur := root
		for j := len(words[i]) - 1; j >= 0; j-- {
			off := int(words[i][j] - 'a')
			if cur.end > -1 && check(words[i][:j+1]) { // e.g., "ab" in trie, we are matching "fgrba"; we need check "fgr"
				//fmt.Println("@@", words[cur.end], words[i])
				ans = append(ans, []int{cur.end, i})
			}
			if cur.children[off] == nil {
				continue outer // no match prefix, continue words
			}
			cur = cur.children[off]
		}
		// found a prefix match
		stack = append(stack, cur) // we got "abfff', "abert" in trie, we have matched "ba"; we need check "fff" and "ert"
		for len(stack) > 0 {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if cur.end > -1 && check(words[cur.end][len(words[i]):]) && cur.end != i {
				//fmt.Println("**", words[cur.end], words[i])
				if words[i] == "" {
					ans = append(ans, []int{i, cur.end})
				}
				ans = append(ans, []int{cur.end, i})
			}
			for j := range cur.children {
				if cur.children[j] != nil {
					stack = append(stack, cur.children[j])
				}
			}
		}
	}
	return ans
}

type trieNode struct {
	end      int
	children [26]*trieNode
}

func check(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func (t *trieNode) add(s string, ind int) {
	cur := t
	for i := range s {
		off := int(s[i] - 'a')
		if cur.children[off] == nil {
			cur.children[off] = &trieNode{end: -1}
			cur = cur.children[off]
		} else {
			cur = cur.children[off]
		}
		if i == len(s)-1 {
			cur.end = ind
		}
	}
}

func main() {
	fmt.Println(palindromePairs([]string{"abcd", "abcdfff", "dcba", "gggdcba", "lls", "s", "sssll", "ss", "sss"}))
	fmt.Println(palindromePairs([]string{"tab", "bat", "cat"}))
	fmt.Println(palindromePairs([]string{"a", ""}))
	fmt.Println(palindromePairs([]string{"a", "b", "c", "ab", "ac", "aa"}))
}

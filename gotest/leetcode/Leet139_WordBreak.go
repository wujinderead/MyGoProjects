package main

import "fmt"

// https://leetcode.com/problems/word-break/

// Given a non-empty string s and a dictionary wordDict containing a list of non-
// empty words, determine if s can be segmented into a space-separated sequence of
// one or more dictionary words.
// Note: The same word in the dictionary may be reused multiple times in the segmentation.
// You may assume the dictionary does not contain duplicate words.
// Example 1:
//   Input: s = "leetcode", wordDict = ["leet", "code"]
//   Output: true
//   Explanation: Return true because "leetcode" can be segmented as "leet code".
// Example 2:
//   Input: s = "applepenapple", wordDict = ["apple", "pen"]
//   Output: true
//   Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
//     Note that you are allowed to reuse a dictionary word.
// Example 3:
//   Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
//   Output: false

func wordBreak(s string, wordDict []string) bool {
	root := &trieNode{}
	for i:= range wordDict {
		addTrie(root, wordDict[i], i+1)
	}
	can := make([]bool, len(s)+1)
	can[len(s)] = true
	var cur *trieNode
	for i:=len(s)-1; i>=0; i-- {
		cur = root
		for j:=i; j<len(s); j++ {
			if cur.childs[int(s[j]-'a')]==nil {
				break
			}
			cur = cur.childs[int(s[j]-'a')]
			if cur.ind>0 && can[j+1] {
				can[i] = true
				break
			}
		}
	}
	return can[0]
}

type trieNode struct {
	childs [26]*trieNode
	ind int
}

func addTrie(root *trieNode, s string, ind int) {
	cur := root
	for i := range s {
		ord := int(s[i]-'a')
		if cur.childs[ord] == nil {
			cur.childs[ord] = &trieNode{}
		}
		cur = cur.childs[ord]
		if i==len(s)-1 {
			cur.ind = ind+1
		}
	}
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreak("a", []string{"a", "b"}))
	fmt.Println(wordBreak("ca", []string{"a", "b"}))
	fmt.Println(wordBreak("aaaabaaaa", []string{"a", "aa", "aaa"}))
}
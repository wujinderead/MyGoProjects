package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/word-ladder/

// Given two words (beginWord and endWord), and a dictionary's word list, find the
// length of shortest transformation sequence from beginWord to endWord, such that:
//   Only one letter can be changed at a time.
//   Each transformed word must exist in the word list. Note that beginWord is not
//     a transformed word.
// Note:
//   Return 0 if there is no such transformation sequence.
//   All words have the same length.
//   All words contain only lowercase alphabetic characters.
//   You may assume no duplicates in the word list.
//   You may assume beginWord and endWord are non-empty and are not the same.
// Example 1:
//   Input:
//     beginWord = "hit",
//     endWord = "cog",
//     wordList = ["hot","dot","dog","lot","log","cog"]
//   Output: 5
//   Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     return its length 5.
// Example 2:
//   Input:
//     beginWord = "hit"
//     endWord = "cog"
//     wordList = ["hot","dot","dog","lot","log"]
//   Output: 0
//   Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// check if end word are in the word list
	endInd := -1
	for i := range wordList {
		if wordList[i] == endWord {
			endInd = i
			break
		}
	}
	if endInd == -1 {
		return 0
	}

	// from beginInd 0 to endInd+1
	queue := list.New()
	queue.PushBack(0)
	step := make([]int, len(wordList)+1)
	visited := make([]bool, len(wordList))
	step[0] = 1
	var s string

	// bfs to find the minimal length
	for queue.Len() > 0 {
		ind := queue.Remove(queue.Front()).(int)
		if ind == 0 {
			s = beginWord
		} else {
			s = wordList[ind-1]
		}
		curstep := step[ind]
		for i := range wordList {
			if !visited[i] && check(s, wordList[i]) {
				visited[i] = true
				queue.PushBack(i + 1)
				step[i+1] = curstep + 1
				if i == endInd {
					return curstep + 1
				}
			}
		}
	}
	return 0
}

func check(s1, s2 string) bool {
	count := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	if count == 1 {
		return true
	}
	return false
}

func main() {
	fmt.Println(ladderLength("hit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("hit", "hot",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("hit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log"}))
}

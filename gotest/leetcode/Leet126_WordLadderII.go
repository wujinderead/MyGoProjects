package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/word-ladder-ii/

// Given two words (beginWord and endWord), and a dictionary's word list, find all
// shortest transformation sequence(s) from beginWord to endWord, such that:
//   Only one letter can be changed at a time
//   Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
// Note:
//   Return an empty list if there is no such transformation sequence.
//   All words have the same length.
//   All words contain only lowercase alphabetic characters.
//   You may assume no duplicates in the word list.
//   You may assume beginWord and endWord are non-empty and are not the same.
// Example 1:
//  Input:
//    beginWord = "hit",
//    endWord = "cog",
//    wordList = ["hot","dot","dog","lot","log","cog"]
//  Output:
//    [
//     ["hit","hot","dot","dog","cog"],
//     ["hit","hot","lot","log","cog"]
//    ]
// Example 2:
//   Input:
//     beginWord = "hit"
//     endWord = "cog"
//     wordList = ["hot","dot","dog","lot","log"]
//   Output: []
//   Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	paths := make([][]string, 0)
	endInd := -1
	for i, s := range wordList {
		if s == endWord {
			endInd = i
			break
		}
	}
	if endInd == -1 { // end word not in list
		return paths
	}

	// make a graph
	N := len(wordList)
	graph := make([][]int, N+1)
	for i := range graph {
		graph[i] = make([]int, 0, N/3)
	}
	var s string
	for i := 1; i <= N; i++ {
		if i == N {
			s = beginWord // index of begin word is N
		} else {
			s = wordList[i]
		}
		for j := 0; j < i; j++ {
			if check(s, wordList[j]) {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	// bfs to find the shortest path from beginInd=len(words) to endInd
	queue := list.New()
	queue.PushBack(len(wordList))
	tier := make([]int, N+1) // the tier for bfs
	visited := make([]bool, N+1)
	tier[N] = 0
	visited[N] = true
	shortest := 0
outer:
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(int)
		for _, v := range graph[cur] {
			if !visited[v] {
				tier[v] = tier[cur] + 1
				visited[v] = true
				if v == endInd { // find a path to endInd, we also find the shortest path
					shortest = tier[v]
					break outer
				}
				queue.PushBack(v)
			}
		}
	}

	// no path from beginInd to endInd
	if !visited[endInd] {
		return paths
	}

	// use backtracking to find all shortest path
	path := make([]string, shortest+1)
	path[0] = beginWord
	findNext(N, endInd, tier, graph, wordList, path, &paths)
	return paths
}

func findNext(curind, endind int, tier []int, graph [][]int, wordlist, path []string, paths *[][]string) {
	if curind == endind {
		pathcopy := make([]string, len(path))
		copy(pathcopy, path)
		*paths = append(*paths, pathcopy)
	}
	curtier := tier[curind]
	for _, v := range graph[curind] {
		if tier[v] == curtier+1 { // find shortest path tier by tier
			path[curtier+1] = wordlist[v]
			findNext(v, endind, tier, graph, wordlist, path, paths)
		}
	}
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
	return true
}

func main() {
	fmt.Println(findLadders("hit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(findLadders("hit", "hot",
		[]string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(findLadders("hit", "cog",
		[]string{"hot", "dot", "dog", "lot", "log"}))
}

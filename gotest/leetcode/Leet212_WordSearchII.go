package main

import "fmt"

// https://leetcode.com/problems/word-search-ii/

// Given a 2D board and a list of words from the dictionary, find all words in the board.
// Each word must be constructed from letters of sequentially adjacent cell, where "adjacent"
// cells are those horizontally or vertically neighboring. The same letter cell may not be
// used more than once in a word.
// Example:
//   Input:
//     board = [
//       ['o','a','a','n'],
//       ['e','t','a','e'],
//       ['i','h','k','r'],
//       ['i','f','l','v']
//     ]
//     words = ["oath","pea","eat","rain"]
//   Output: ["eat","oath"]
// Note:
//   All inputs are consist of lowercase letters a-z.
//   The values of words are distinct.

func findWords(board [][]byte, words []string) []string {
	// construct a trie to represent all words
	t := &trie{root: &trieNode{children: make([]*trieNode, 0)}}
	for i := range words {
		t.add(words[i])
	}
	visited := make([][]bool, len(board))
	for i := range board {
		visited[i] = make([]bool, len(board[0]))
	}
	occur := make(map[string]struct{})
	t.buf = make([]byte, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			ch := board[i][j]
			if has, ind := hasChild(t.root, ch); has {
				t.stack = append(t.stack, t.root.children[ind])
				check(board, i, j, t, occur, visited)
			}
		}
	}
	ans := make([]string, 0)
	for k := range occur {
		ans = append(ans, k)
	}
	return ans
}

type trie struct {
	root  *trieNode
	stack []*trieNode
	buf   []byte
}

type trieNode struct {
	char     byte
	end      bool
	children []*trieNode
}

func check(board [][]byte, i, j int, t *trie, occur map[string]struct{}, visited [][]bool) {
	cur := t.stack[len(t.stack)-1]
	t.buf = append(t.buf, cur.char)
	visited[i][j] = true
	if cur.end {
		occur[string(t.buf)] = struct{}{}
	}
	for _, v := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		ni, nj := i+v[0], j+v[1]
		if ni >= 0 && ni < len(board) && nj >= 0 && nj < len(board[0]) {
			if has, ind := hasChild(cur, board[ni][nj]); has && !visited[ni][nj] {
				t.stack = append(t.stack, cur.children[ind])
				check(board, ni, nj, t, occur, visited)
			}
		}
	}
	t.stack = t.stack[:len(t.stack)-1]
	t.buf = t.buf[:len(t.buf)-1]
	visited[i][j] = false
}

func hasChild(node *trieNode, ch byte) (has bool, ind int) {
	for i := range node.children {
		if node.children[i].char == ch {
			has = true
			ind = i
			return
		}
	}
	return
}

func (t *trie) add(word string) {
	cur := t.root
outer:
	for i := range word {
		for _, child := range cur.children {
			if child.char == word[i] {
				cur = child
				if i == len(word)-1 {
					cur.end = true
				}
				continue outer
			}
		}
		newnode := &trieNode{char: word[i], children: make([]*trieNode, 0)}
		cur.children = append(cur.children, newnode)
		cur = newnode
		if i == len(word)-1 {
			cur.end = true
		}
	}
}

// for debug
func (t *trie) traverse() {
	stack := []*trieNode{t.root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		fmt.Println(string(cur.char), cur.end)
		stack = stack[:len(stack)-1]
		for i := len(cur.children) - 1; i >= 0; i-- {
			stack = append(stack, cur.children[i])
		}
	}
}

func main() {
	fmt.Println(findWords([][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}, []string{"oath", "pea", "eat", "rain", "ra", "rainy"}))
	fmt.Println(findWords([][]byte{
		{'a'},
	}, []string{"a"}))
	fmt.Println(findWords([][]byte{
		{'a', 'a'},
	}, []string{"a"}))
	fmt.Println(findWords([][]byte{
		{'a', 'a'},
	}, []string{"aaa"}))
}

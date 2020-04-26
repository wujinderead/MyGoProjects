package main

import "fmt"

// https://leetcode.com/problems/stream-of-characters/

// Implement the StreamChecker class as follows:
// StreamChecker(words): Constructor, init the data structure with the given words.
// query(letter): returns true if and only if for some k >= 1, the last k characters queried
// (in order from oldest to newest, including this letter just queried) spell one of
// the words in the given list.
// Example:
//   StreamChecker streamChecker = new StreamChecker(["cd","f","kl"]); // init the dictionary.
//   streamChecker.query('a');          // return false
//   streamChecker.query('b');          // return false
//   streamChecker.query('c');          // return false
//   streamChecker.query('d');          // return true, because 'cd' is in the word list
//   streamChecker.query('e');          // return false
//   streamChecker.query('f');          // return true, because 'f' is in the word list
//   streamChecker.query('g');          // return false
//   streamChecker.query('h');          // return false
//   streamChecker.query('i');          // return false
//   streamChecker.query('j');          // return false
//   streamChecker.query('k');          // return false
//   streamChecker.query('l');          // return true, because 'kl' is in the word list
// Note:
//   1 <= words.length <= 2000
//   1 <= words[i].length <= 2000
//   Words will only consist of lowercase English letters.
//   Queries will only consist of lowercase English letters.
//   The number of queries is at most 40000.

type StreamChecker struct {
	child   [26]*StreamChecker
	cursors []*StreamChecker
	hasword bool
}

//create a trie
func Constructor(words []string) StreamChecker {
	root := StreamChecker{}
	for i := range words {
		cur := &root
		for j := range words[i] {
			if cur.child[int(words[i][j]-'a')] == nil {
				cur.child[int(words[i][j]-'a')] = &StreamChecker{}
			}
			cur = cur.child[int(words[i][j]-'a')]
			if j == len(words[i])-1 {
				cur.hasword = true
			}
		}
	}
	root.cursors = make([]*StreamChecker, 0)
	return root
}

// use a list ot store current cursors
func (this *StreamChecker) Query(letter byte) bool {
	this.cursors = append(this.cursors, this)
	has := false
	newcursors := make([]*StreamChecker, 0)
	for _, cur := range this.cursors {
		child := cur.child[int(letter-'a')]
		if child != nil {
			newcursors = append(newcursors, child)
			has = has || child.hasword
		}
	}
	this.cursors = newcursors
	return has
}

func main() {
	streamChecker := Constructor([]string{"cad", "f", "c", "kl"}) // init the dictionary.
	fmt.Println(streamChecker.Query('a'))                         // return false
	fmt.Println(streamChecker.Query('b'))                         // return false
	fmt.Println(streamChecker.Query('c'))                         // return true, because 'c' is in the word list
	fmt.Println(streamChecker.Query('a'))                         // return false
	fmt.Println(streamChecker.Query('d'))                         // return true, because 'cd' is in
	fmt.Println(streamChecker.Query('e'))                         // return false
	fmt.Println(streamChecker.Query('f'))                         // return true, because 'f' is in
	fmt.Println(streamChecker.Query('g'))                         // return false
	fmt.Println(streamChecker.Query('h'))                         // return false
	fmt.Println(streamChecker.Query('i'))                         // return false
	fmt.Println(streamChecker.Query('j'))                         // return false
	fmt.Println(streamChecker.Query('k'))                         // return false
	fmt.Println(streamChecker.Query('l'))                         // return true, because 'kl' is in
	fmt.Println()
	streamChecker = Constructor([]string{"dfghi", "fgc", "g"})
	fmt.Println(streamChecker.Query('d'))
	fmt.Println(streamChecker.Query('f'))
	fmt.Println(streamChecker.Query('g')) // true
	fmt.Println(streamChecker.Query('c')) // true
	fmt.Println(streamChecker.Query('a'))
	fmt.Println(streamChecker.Query('d'))
	fmt.Println(streamChecker.Query('f'))
	fmt.Println(streamChecker.Query('g')) // true
	fmt.Println(streamChecker.Query('s'))
	fmt.Println(streamChecker.Query('d'))
	fmt.Println(streamChecker.Query('f'))
	fmt.Println(streamChecker.Query('g')) // true
	fmt.Println(streamChecker.Query('h'))
	fmt.Println(streamChecker.Query('i')) // true
	streamChecker = Constructor([]string{"ab", "ba", "aaab", "abab", "baa"})
	bools := make([]bool, 0)
	s := "aaaaabababbbababbbbababaaabaaa"
	for i := range s {
		bools = append(bools, streamChecker.Query(s[i]))
	}
	fmt.Println(bools)
}

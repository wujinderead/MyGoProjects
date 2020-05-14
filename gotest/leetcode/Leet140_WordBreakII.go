package main

import "fmt"

// https://leetcode.com/problems/word-break-ii/

// Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
// add spaces in s to construct a sentence where each word is a valid dictionary word.
// Return all such possible sentences.
// Note:
//   The same word in the dictionary may be reused multiple times in the segmentation.
//   You may assume the dictionary does not contain duplicate words.
// Example 1:
//   Input:
//     s = "catsanddog"
//     wordDict = ["cat", "cats", "and", "sand", "dog"]
//   Output:
//     [
//       "cats and dog",
//       "cat sand dog"
//     ]
// Example 2:
//   Input:
//     s = "pineapplepenapple"
//     wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
//   Output:
//     [
//       "pine apple pen apple",
//       "pineapple pen apple",
//       "pine applepen apple"
//     ]
//   Explanation: Note that you are allowed to reuse a dictionary word.
// Example 3:
//   Input:
//     s = "catsandog"
//     wordDict = ["cats", "dog", "sand", "and", "cat"]
//   Output: []

func wordBreak(s string, wordDict []string) []string {
	// make a trie of dictionary
	root := &trieNode{}
    for i := range wordDict {
    	addTrie(root, wordDict[i], i+1)
	}

	// check if can break first
	if !canBreak(s, wordDict, root) {
		return []string{}
	}
	buf := make([][]string, 0)

	// make a map to contained computed result. e.g., remap[4]=[3]int{0,2,1}
	// (0,2,1) mean buf[0:2][1:], i.e.,
	// means s[4:]="applepenapple" can be form by buf[0][1:] or buf[1][1:]
	// remap[9]=[3]int{0,1,2} means s[9:]="penapple" formed by buf[0][2:]
	// buf = {{"pine", "apple", "pen", "apple"},
	//        {"pine", "applepen", "apple"},
	//        {"pineapple", "pen", "apple"}}
	remap := make(map[int][3]int)
	strs := make([]string, 0, 10)

	// break word backtracking function
	breakWord(s, 0, root, remap, &buf, &strs)
	//fmt.Println(remap)

	// from buf to answer
	ans := make([]string, len(buf))
	for i := range buf {
		tmp := make([]byte, 0, len(s)+len(buf[i])-1)
		tmp = append(tmp, buf[i][0]...)
		for j:=1; j<len(buf[i]); j++ {
			tmp = append(tmp, ' ')
			tmp = append(tmp, buf[i][j]...)
		}
		ans[i] = string(tmp)
	}
	return ans
}

func breakWord(s string, st int, root *trieNode, remap map[int][3]int, buf *[][]string, strs *[]string) {
	if st == len(s) {
		tmp := make([]string, len(*strs))
		copy(tmp, *strs)
		*buf = append(*buf, tmp)
		return
	}
	cur := root
	for i:=st; i<len(s); i++ {
		ord := int(s[i]-'a')
		if cur.childs[ord] == nil {
			break
		}
		cur = cur.childs[ord]
		if cur.ind>0 {   // found s[st: i+1] in dictionary, continue to suffix
			if v, ok := remap[i+1]; ok {  // if suffix computed before, use it
				//fmt.Println("remap:", append(*strs, s[st: i+1]), s[i+1:])
				for j:=v[0]; j<v[1]; j++ {
					tmp := make([]string, 0, len(*strs)+1+len((*buf)[j][v[2]:]))
					tmp = append(tmp, *strs...)
					tmp = append(tmp, s[st: i+1])
					tmp = append(tmp, (*buf)[j][v[2]:]...)
					*buf = append(*buf, tmp)
				}
				continue
			}
			// continue to s[i+1:]
			*strs = append(*strs, s[st: i+1])
			prevlen := len(*buf)
			//if i+1<len(s) { fmt.Println("first:", *strs, s[i+1:])}
			breakWord(s, i+1, root, remap, buf, strs)
			if len(*buf)>prevlen && i+1<len(s) {  // we can break s[i+1:], record its result
				remap[i+1] = [3]int{prevlen, len(*buf), len(*strs)}
			}
			*strs = (*strs)[:len(*strs)-1]
		}
	}
}

func canBreak(s string, wordDict []string, root *trieNode) bool {
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
	for _, v := range []struct{s string; dict []string}{
		{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}},
		{"pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		[]string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}},
	} {
		ans := wordBreak(v.s, v.dict)
		if len(ans)<100 {
			for i := range ans {
				fmt.Printf("%q\n", ans[i])
			}
		}
		fmt.Println(len(ans))
		fmt.Println()
	}
}
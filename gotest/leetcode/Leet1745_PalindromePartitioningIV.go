package main

import "fmt"

// https://leetcode.com/problems/palindrome-partitioning-iv/

// Given a string s, return true if it is possible to split the string s into three
// non-empty palindromic substrings. Otherwise, return false.
// A string is said to be palindrome if it the same string when reversed.
// Example 1:
//   Input: s = "abcbdd"
//   Output: true
//   Explanation: "abcbdd" = "a" + "bcb" + "dd", and all three substrings are palindromes.
// Example 2:
//   Input: s = "bcbddxy"
//   Output: false
//   Explanation: s cannot be split into 3 palindromes.
// Constraints:
//   3 <= s.length <= 2000
//   s consists only of lowercase English letters.

// we need O(n^2) to build a lookup table
// and O(n^2) time to check 3 parts
func checkPartitioning(s string) bool {
	// smap[i] is the length for all palindromic subarrays that start at i
	smap := make([][]int, len(s))
	ssmap := make(map[[2]int]struct{})
	for i := range smap {
		smap[i] = make([]int, 1, 3)
		smap[i][0] = 1 // always has len 1 palindrome
		ssmap[[2]int{i, 1}] = struct{}{}
	}

	// get all palindromic subarray
	for i := 1; i < len(s)-1; i++ {
		l, r := i-1, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			smap[l] = append(smap[l], r-l+1)
			ssmap[[2]int{l, r - l + 1}] = struct{}{}
			l--
			r++
		}
	}
	for i := 0; i < len(s)-1; i++ {
		l, r := i, i+1
		for l >= 0 && r < len(s) && s[l] == s[r] {
			smap[l] = append(smap[l], r-l+1)
			ssmap[[2]int{l, r - l + 1}] = struct{}{}
			l--
			r++
		}
	}

	// check all pairs
	lengs := smap[0]
	for _, leng := range lengs {
		// s[:leng] is palindromic
		if leng == len(s) {
			continue
		}
		leng1s := smap[leng]
		for _, leng1 := range leng1s {
			if leng+leng1 == len(s) {
				continue
			}
			// s[leng: leng+leng1] is palindromic
			// check if s[leng+leng1:] is palindromic
			if _, ok := ssmap[[2]int{leng + leng1, len(s) - leng - leng1}]; ok {
				return true
			}
		}
	}
	return false
}

func main() {
	for _, v := range []struct {
		s   string
		ans bool
	}{
		{"abcbdd", true},
		{"bcbddxy", false},
		{"bbab", true},
	} {
		fmt.Println(checkPartitioning(v.s), v.ans)
	}
}

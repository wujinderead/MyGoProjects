package main

import "fmt"

// https://leetcode.com/problems/minimum-window-substring

// Given a string S and a string T, find the minimum window in S which will contain
// all the characters in T in complexity O(n).
// Example:
//   Input: S = "ADOBECODEBANC", T = "ABC"
//   Output: "BANC"
// Note:
//   If there is no such window in S that covers all characters in T, return the empty string "".
//   If there is such window, you are guaranteed that there will always be only one unique minimum window in S.

func minWindow(s string, t string) string {
	if t=="" {
		return ""
	}
    sset, tset := make([]int, 128), make([]int, 128)
	for i := range t {
		tset[int(t[i])]++
	}
	shortest, sstart := 0x7fffffff, 0
	start, end := 0, 0
	for end<=len(s) {
		if !contain(sset, tset) {
			// fmt.Println("noo", s[start:end], sset)
			if end==len(s) {
				break
			}
			cur := int(s[end])
			if tset[cur]>0 {
				sset[cur]++
			}
			end++
		} else {
			l := end-start
			if l<shortest {
				shortest = l
				sstart = start
			}
			// fmt.Println("con", s[start:end], sset, shortest, start)
			cur := int(s[start])
			if sset[cur]>0 {
				sset[cur]--
			}
			start++
		}
	}
	if shortest == 0x7fffffff {
		return ""
	}
	return s[sstart: sstart+shortest]
}

func contain(sset, tset []int) bool {
	for i := range sset {
		if sset[i]<tset[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("ADEFADEBCADAE", "AABC"))
	fmt.Println(minWindow("ABC", "A"))
	fmt.Println(minWindow("ABC", "D"))
	fmt.Println(minWindow("ABC", "AA"))
	fmt.Println(minWindow("ABC", "ABCD"))
	fmt.Println(minWindow("ABC", ""))
	fmt.Println(minWindow("A", "A"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("!@#$%^&&*", "&&$"))
}
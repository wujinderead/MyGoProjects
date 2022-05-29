package main

import "fmt"

// https://leetcode.com/problems/minimum-window-substring

// Given two strings s and t of lengths m and n respectively, return the minimum window substring
// of s such that every character in t (including duplicates) is included in the window. If there
// is no such substring, return the empty string "".
// The testcases will be generated such that the answer is unique.
// A substring is a contiguous sequence of characters within the string.
// Example 1:
//   Input: s = "ADOBECODEBANC", t = "ABC"
//   Output: "BANC"
//   Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
// Example 2:
//   Input: s = "a", t = "a"
//   Output: "a"
//   Explanation: The entire string s is the minimum window.
// Example 3:
//   Input: s = "a", t = "aa"
//   Output: ""
//   Explanation: Both 'a's from t must be included in the window.
//     Since the largest window of s only has one 'a', return empty string.
// Constraints:
//   m == s.length
//   n == t.length
//   1 <= m, n <= 10⁵
//   s and t consist of uppercase and lowercase English letters

func minWindow(s string, t string) string {
	if t == "" {
		return ""
	}
	sset, tset := make([]int, 128), make([]int, 128)
	for i := range t {
		tset[int(t[i])]++
	}
	shortest, sstart := 0x7fffffff, 0
	start, end := 0, 0
	for end <= len(s) {
		if !contain(sset, tset) {
			// fmt.Println("noo", s[start:end], sset)
			if end == len(s) {
				break
			}
			cur := int(s[end])
			if tset[cur] > 0 {
				sset[cur]++
			}
			end++
		} else {
			l := end - start
			if l < shortest {
				shortest = l
				sstart = start
			}
			// fmt.Println("con", s[start:end], sset, shortest, start)
			cur := int(s[start])
			if sset[cur] > 0 {
				sset[cur]--
			}
			start++
		}
	}
	if shortest == 0x7fffffff {
		return ""
	}
	return s[sstart : sstart+shortest]
}

func contain(sset, tset []int) bool {
	for i := range sset {
		if sset[i] < tset[i] {
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

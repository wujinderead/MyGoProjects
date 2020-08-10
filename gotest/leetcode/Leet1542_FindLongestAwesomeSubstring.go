package main

import (
	"fmt"
)

// https://leetcode.com/problems/find-longest-awesome-substring

// Given a string s. An awesome substring is a non-empty substring of s such that
// we can make any number of swaps in order to make it palindrome.
// Return the length of the maximum length awesome substring of s.
// Example 1:
//   Input: s = "3242415"
//   Output: 5
//   Explanation: "24241" is the longest awesome substring, we can form the palindrome "24142" with some swaps.
// Example 2:
//   Input: s = "12345678"
//   Output: 1
// Example 3:
//   Input: s = "213123"
//   Output: 6
//   Explanation: "213123" is the longest awesome substring, we can form the palindrome "231132" with some swaps.
// Example 4:
//   Input: s = "00"
//   Output: 2
// Constraints:
//   1 <= s.length <= 10^5
//   s consists only of digits.

// O(10N)
// mask means the occurrence of digits of prefix. 
// if some bit is 1, means that the corresponding digit occurs odd times; if 0, even times.
// for a mask, check if has occurred; and then flip each bit and check if has occurred.
// e.g., for 1100, if 1100 has occurred, means the between mask 0000 is all even.
// if 1101 has occurred, means the between mask is 0001, which is palindromic.
func longestAwesome(s string) int {
	mapp := make(map[int]int)
	mapp[0] = -1
	mask := 0
	maxlen := 0
	for i:=0; i<len(s); i++ {
		mask = mask ^ (1<<uint(s[i]-'0'))   // flip this bit
		for j:=0; j<10; j++ {
			mm := mask ^ (1<<uint(j))       // flip each bit and check
			if v, ok := mapp[mm]; ok {
				maxlen = max(maxlen, i-v)
			} 
		}
		if v, ok := mapp[mask]; ok {
			maxlen = max(maxlen, i-v)
		} else {
			mapp[mask] = i	
		}
	}
	return maxlen
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct{s string; ans int} {
		{"3242415", 5},
		{"12345678", 1},
		{"213123", 6},
		{"00", 2},
		{"51224", 3},
		{"350844", 3},
		{"185801630663498", 5},
	} {
		fmt.Println(longestAwesome(v.s), v.ans)
	}
}
package main

import "fmt"

// https://leetcode.com/problems/count-number-of-homogenous-substrings/

// Given a string s, return the number of homogenous substrings of s. Since the
// answer may be too large, return it modulo 10^9 + 7.
// A string is homogenous if all the characters of the string are the same.
// A substring is a contiguous sequence of characters within a string.
// Example 1:
//   Input: s = "abbcccaa"
//   Output: 13
//   Explanation: The homogenous substrings are listed as below:
//     "a"   appears 3 times.
//     "aa"  appears 1 time.
//     "b"   appears 2 times.
//     "bb"  appears 1 time.
//     "c"   appears 3 times.
//     "cc"  appears 2 times.
//     "ccc" appears 1 time.
//     3 + 1 + 2 + 1 + 3 + 2 + 1 = 13.
// Example 2:
//   Input: s = "xy"
//   Output: 2
//   Explanation: The homogenous substrings are "x" and "y".
// Example 3:
//   Input: s = "zzzzz"
//   Output: 15
// Constraints:
//   1 <= s.length <= 105
//   s consists of lowercase letters.

func countHomogenous(s string) int {
	mod := int(1e9) + 7
	var i, e, all, l int
	for i < len(s) {
		e = i
		for e+1 < len(s) && s[e+1] == s[i] {
			e++
		}
		l = e - i + 1
		all += l * (l + 1) / 2
		all = all % mod
		i = e + 1
	}
	return all
}

func main() {
	for _, v := range []struct {
		s   string
		ans int
	}{
		{"abbcccaa", 13},
		{"xy", 2},
		{"zzzzz", 15},
	} {
		fmt.Println(countHomogenous(v.s), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/splitting-a-string-into-descending-consecutive-values/

// You are given a string s that consists of only digits.
// Check if we can split s into two or more non-empty substrings such that the numerical
// values of the substrings are in descending order and the difference between numerical values
// of every two adjacent substrings is equal to 1.
// For example, the string s = "0090089" can be split into ["0090", "089"] with
// numerical values [90,89]. The values are in descending order and adjacent values
// differ by 1, so this way is valid.
// Another example, the string s = "001" can be split into ["0", "01"], ["00", "1"], or
// ["0", "0", "1"]. However all the ways are invalid because they have numerical values
// [0,1], [0,1], and [0,0,1] respectively, all of which are not in descending order.
// Return true if it is possible to split s as described above, or false otherwise.
// A substring is a contiguous sequence of characters in a string.
// Example 1:
//   Input: s = "1234"
//   Output: false
//   Explanation: There is no valid way to split s.
// Example 2:
//   Input: s = "050043"
//   Output: true
//   Explanation: s can be split into ["05", "004", "3"] with numerical values [5,4,3].
//   The values are in descending order with adjacent values differing by 1.
// Example 3:
//   Input: s = "9080701"
//   Output: false
//   Explanation: There is no valid way to split s.
// Example 4:
//   Input: s = "10009998"
//   Output: true
//   Explanation: s can be split into ["100", "099", "98"] with numerical values [100,99,98].
//     The values are in descending order with adjacent values differing by 1.
// Constraints:
//   1 <= s.length <= 20
//   s only consists of digits.

// just backtracking
func splitString(s string) bool {
	pre := 0
	for i := 0; i < len(s)-1; i++ {
		pre = pre*10 + int(s[i]-'0')
		if pre == 0 {
			continue
		}
		if findNext(pre-1, i+1, s) {
			return true
		}
	}
	return false
}

func findNext(target, start int, s string) (ans bool) {
	if start == len(s) {
		return true
	}
	if target == 0 {
		for i := start; i < len(s); i++ {
			if s[i] != '0' {
				return false
			}
		}
		return true
	}
	if target < 0 {
		return false
	}
	pre := 0
	for i := start; i < len(s); i++ {
		pre = pre*10 + int(s[i]-'0')
		if pre == target {
			return findNext(target-1, i+1, s)
		}
		if pre > target {
			return false
		}
	}
	return false
}

func main() {
	for _, v := range []struct {
		s   string
		ans bool
	}{
		{"1234", false},
		{"050043", true},
		{"9080701", false},
		{"000010", true},
		{"000098", true},
		{"987", true},
		{"10009998", true},
		{"4", false},
		{"004", false},
		{"0043", true},
		{"004032", true},
		{"0040320", false},
		{"00403201", true},
		{"004032010", true},
		{"0040320100", true},
		{"00000", false},
		{"00100", true},
	} {
		fmt.Println(splitString(v.s), v.ans)
	}
}

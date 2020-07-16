package main

import (
	"fmt"
)

// https://leetcode.com/problems/number-of-substrings-with-only-1s/

// Given a binary string s (a string consisting only of '0' and '1's).
// Return the number of substrings with all characters 1's.
// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: s = "0110111"
//   Output: 9
//   Explanation: There are 9 substring in total with only 1's characters.
//     "1" -> 5 times.
//     "11" -> 3 times.
//     "111" -> 1 time.
// Example 2:
//   Input: s = "101"
//   Output: 2
//   Explanation: Substring "1" is shown 2 times in s.
// Example 3:
//   Input: s = "111111"
//   Output: 21
//   Explanation: Each substring contains only 1's characters.
// Example 4:
//   Input: s = "000"
//   Output: 0
// Constraints:
//   s[i] == '0' or s[i] == '1'
//   1 <= s.length <= 10^5

func numSub(s string) int {
	total := 0
	c := 0
	for _, v := range s {
		if v=='1' {  // find an '1', add current consecutive 1s count
			c++
			total += c
			total = total%int(1e9+7)
		} else {
			c = 0
		}
	}
	return total
}

func main() {
	fmt.Println(numSub("0110111"))
	fmt.Println(numSub("101"))
	fmt.Println(numSub("111111"))
	fmt.Println(numSub("000"))
	fmt.Println(numSub("01111111111001111111110011"))
	fmt.Println(numSub("1"))
}
package main

import "fmt"

// https://leetcode.com/problems/decode-ways/

// A message containing letters from A-Z is being encoded to numbers using the following mapping:
// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// Given a non-empty string containing only digits, determine the total number of ways to decode it.
// Example 1:
//   Input: "12"
//   Output: 2
//   Explanation: It could be decoded as "AB" (1 2) or "L" (12).
// Example 2:
//   Input: "226"
//   Output: 3
//   Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

// basic idea is like fibonacci array. "1" 1 way, "12" 2 ways; then add "3", if we dont attach "2" and "3",
// we got "1,2 | 3" and "12 | 3"; if we attach "2" and "3", we got "1 | 23". it's thus like fibonacci.
// consider "112", there is 3 ways; if add "0", "0" must attach "2", so it become "11 | 20", only 2 ways.
func numDecodings(s string) int {
	ways := 1
	s0 := 0
	s1 := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '1' || s[i] == '2' {
			if i == len(s)-1 { // s ends with "...12"
				ways *= s0 + s1
			}
			s0, s1 = s1, s0+s1
			continue
		} else if s[i] == '0' { // "...10" or "...20"
			ways *= s0
		} else if i > 0 && s[i-1] == '2' && s[i] > '6' { // "...2(7|8|9)"
			ways *= s1
		} else {
			ways *= s0 + s1 // "...1(3-9)", "...2(3-6)"
		}
		s0 = 0
		s1 = 1
	}
	return ways
}

func main() {
	fmt.Println(numDecodings("261221631712204510"))
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("120"))
	fmt.Println(numDecodings("27"))
	fmt.Println(numDecodings("20"))
	fmt.Println(numDecodings("227"))
	fmt.Println(numDecodings("7"))
}

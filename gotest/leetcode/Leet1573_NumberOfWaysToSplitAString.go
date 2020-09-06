package main

import "fmt"

// https://leetcode.com/problems/number-of-ways-to-split-a-string

// Given a binary string s (a string consisting only of '0's and '1's), we can split s into 3
// non-empty strings s1, s2, s3 (s1+ s2+ s3 = s).
// Return the number of ways s can be split such that the number of characters '1' is the same
// in s1, s2, and s3.
// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: s = "10101"
//   Output: 4
//   Explanation: There are four ways to split s in 3 parts where each part contain the same number of letters '1'.
//     "1|010|1"
//     "1|01|01"
//     "10|10|1"
//     "10|1|01"
// Example 2:
//   Input: s = "1001"
//   Output: 0
// Example 3:
//   Input: s = "0000"
//   Output: 3
//   Explanation: There are three ways to split s in 3 parts.
//     "0|0|00"
//     "0|00|0"
//     "00|0|0"
// Example 4:
//   Input: s = "100100010100110"
//   Output: 12
// Constraints:
//   s[i] == '0' or s[i] == '1'
//   3 <= s.length <= 10^5

func numWays(s string) int {
	c := 0
    for i:=0; i<len(s); i++ {
    	if s[i]=='1' {
    		c++
		}
	}
	if c%3 != 0 {
		return 0
	}
	mod := int(1e9+7)
	if c == 0 {
		return ((len(s)-1)*(len(s)-2)/2) % mod
	}
	c = c/3
	var p1, p2, p3, p4 = -1, -1, -1, -1
	p := 0
	for i:=0; i<len(s); i++ {
		if s[i]=='1' {
			p++
		}
		if p1==-1 && p==c {
			p1 = i
		}
		if p2==-1 && p==c+1 {
			p2 = i
		}
		if p3==-1 && p==2*c {
			p3 = i
		}
		if p4==-1 && p==2*c+1 {
			p4 = i
		}
	}
	return ((p2-p1)*(p4-p3)) % mod
}

func main() {
	for _, v := range []struct{s string; ans int} {
		{"10101", 4},
		{"1001", 0},
		{"0000", 3},
		{"100100010100110", 12},
	} {
		fmt.Println(numWays(v.s), v.ans)
	}
}
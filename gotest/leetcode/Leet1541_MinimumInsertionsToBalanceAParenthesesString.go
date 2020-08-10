package main

import "fmt"

// https://leetcode.com/problems/minimum-insertions-to-balance-a-parentheses-string

// Given a parentheses string s containing only the characters '(' and ')'.
// A parentheses string is balanced if:
// .  Any left parenthesis '(' must have a corresponding two consecutive right parenthesis '))'.
// .  Left parenthesis '(' must go before the corresponding two consecutive right parenthesis '))'.
// For example, "())", "())(())))" and "(())())))" are balanced, ")()", "()))" and "(()))" are not balanced.
// You can insert the characters '(' and ')' at any position of the string to balance it if needed.
// Return the minimum number of insertions needed to make s balanced.
// Example 1:
//   Input: s = "(()))"
//   Output: 1
//   Explanation: The second '(' has two matching '))', but the first '(' has only ')' matching.
//     We need to to add one more ')' at the end of the string to be "(())))" which is balanced.
// Example 2:
//   Input: s = "())"
//   Output: 0
//   Explanation: The string is already balanced.
// Example 3:
//   Input: s = "))())("
//   Output: 3
//   Explanation: Add '(' to match the first '))', Add '))' to match the last '('.
// Example 4:
//   Input: s = "(((((("
//   Output: 12
//   Explanation: Add 12 ')' to balance the string.
// Example 5:
//   Input: s = ")))))))"
//   Output: 5
//   Explanation: Add 4 '(' at the beginning of the string and one ')' at the end.
//     The string becomes "(((())))))))".
// Constraints:
//   1 <= s.length <= 10^5
//   s consists of '(' and ')' only.

func minInsertions(s string) int {
	var l, count, i int
    for i<len(s) {
		if s[i]=='(' {
			l++
			i++
			continue
		} 
		if s[i]==')' {
			if i+1<len(s) && s[i+1]==')' {   // double ')'
				if l>0 {    // "())", just decrease l
					l--
				} else {    // "))", make it "())"
					count++
				}
				i+=2
			} else {   // single ')'
				if l>0 {       // "()", make it "())"  
					l--
					count++
				} else {       // ")", make it "())"
					count+=2
				}
				i++
			}
		}
	}
	return count+2*l
}

func main() {
	for _, v := range []struct{s string; ans int} {
		{"(()))", 1},
		{"())", 0},
		{"))())(", 3},
		{"((((((", 12},
		{")))))))", 5},
		{"(()))(()))()())))", 4},
	} {
		fmt.Println(minInsertions(v.s), v.ans)
	}
}
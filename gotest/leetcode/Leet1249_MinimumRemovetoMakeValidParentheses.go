package main

import "fmt"

// https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/

// Given a string s of '(' , ')' and lowercase English characters. Your task is to remove
// the minimum number of parentheses ( '(' or ')', in any positions ) so that the resulting
// parentheses string is valid and return any valid string.
// Formally, a parentheses string is valid if and only if:
//   It is the empty string, contains only lowercase characters, or
//   It can be written as AB (A concatenated with B), where A and B are valid strings, or
//   It can be written as (A), where A is a valid string.
// Example 1:
//   Input: s = "lee(t(c)o)de)"
//   Output: "lee(t(c)o)de"
//   Explanation: "lee(t(co)de)" , "lee(t(c)ode)" would also be accepted.
// Example 2:
//   Input: s = "a)b(c)d"
//   Output: "ab(c)d"
// Example 3:
//   Input: s = "))(("
//   Output: ""
//   Explanation: An empty string is also valid.
// Example 4:
//   Input: s = "(a(b(c)d)"
//   Output: "a(b(c)d)"
// Constraints:
//   1 <= s.length <= 10^5
//   s[i] is one of '(' , ')' and lowercase English letters.

func minRemoveToMakeValid(s string) string {
	todelete := make([]int, 0)
	stack := make([]int, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		}
		if s[i] == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				todelete = append(todelete, i)
			}
		}
	}
	re := make([]byte, len(s)-len(todelete)-len(stack))
	tode := make([]bool, len(s))
	for _, v := range todelete {
		tode[v] = true
	}
	for _, v := range stack {
		tode[v] = true
	}
	j := 0
	for i := 0; i < len(s); i++ {
		if !tode[i] {
			re[j] = s[i]
			j++
		}
	}
	return string(re)
}

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
	fmt.Println(minRemoveToMakeValid("a)b(c)d"))
	fmt.Println(minRemoveToMakeValid("))(("))
	fmt.Println(minRemoveToMakeValid("(a(b(c)d)"))
}

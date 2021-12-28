package main

import "fmt"

// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/

// A parentheses string is a non-empty string consisting only of '(' and ')'.
// It is valid if any of the following conditions is true:
//   It is ().
//   It can be written as AB (A concatenated with B), where A and B are valid parentheses strings.
//   It can be written as (A), where A is a valid parentheses string.
// You are given a parentheses string s and a string locked, both of length n.
// locked is a binary string consisting only of '0's and '1's. For each index i of locked,
//   If locked[i] is '1', you cannot change s[i].
//   But if locked[i] is '0', you can change s[i] to either '(' or ')'.
// Return true if you can make s a valid parentheses string. Otherwise, return false.
// Example 1:
//   Input: s = "))()))", locked = "010100"
//   Output: true
//   Explanation: locked[1] == '1' and locked[3] == '1', so we cannot change s[1] or s[3].
//   We change s[0] and s[4] to '(' while leaving s[2] and s[5] unchanged to makes valid.
// Example 2:
//   Input: s = "()()", locked = "0000"
//   Output: true
//   Explanation: We do not need to make any changes because s is already valid.
// Example 3:
//   Input: s = ")", locked = "0"
//   Output: false
//   Explanation: locked permits us to change s[0].
//     Changing s[0] to either '(' or ')' will not make s valid.
// Constraints:
//   n == s.length == locked.length
//   1 <= n <= 10⁵
//   s[i] is either '(' or ')'.
//   locked[i] is either '0' or '1'.

func canBeValid(s string, locked string) bool {
	if len(s)%2 == 1 {
		return false
	}
	// from left to right
	d, can := 0, 0
	for i := 0; i < len(s); i++ {
		if locked[i] == '1' {
			if s[i] == '(' {
				d++
			} else {
				d--
			}
		} else {
			can++
		}
		if can+d < 0 {
			return false
		}
	}
	if can < d {
		return false
	}
	// from right to left
	d, can = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if locked[i] == '1' {
			if s[i] == ')' {
				d++
			} else {
				d--
			}
		} else {
			can++
		}
		if can+d < 0 {
			return false
		}
	}
	if can < d {
		return false
	}
	return true
}

func main() {
	for _, v := range []struct {
		s, l string
		ans  bool
	}{
		{"))()))", "010100", true},
		{"()()", "0000", true},
		{")", "0", false},
		{"())()))()(()(((())(()()))))((((()())(())", "1011101100010001001011000000110010100101", true},
	} {
		fmt.Println(canBeValid(v.s, v.l), v.ans)
	}
}

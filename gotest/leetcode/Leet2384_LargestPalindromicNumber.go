package main

import "fmt"

// https://leetcode.com/problems/largest-palindromic-number/

// You are given a string num consisting of digits only.
// Return the largest palindromic integer (in the form of a string) that can be
// formed using digits taken from num. It should not contain leading zeroes.
// Notes:
// You do not need to use all the digits of num, but you must use at least one digit.
// The digits can be reordered.
// Example 1:
//   Input: num = "444947137"
//   Output: "7449447"
//   Explanation:
//     Use the digits "4449477" from "444947137" to form the palindromic integer "7449447".
//     It can be shown that "7449447" is the largest palindromic integer that can be formed.
// Example 2:
//   Input: num = "00009"
//   Output: "9"
//   Explanation:
//     It can be shown that "9" is the largest palindromic integer that can be formed.
//     Note that the integer returned should not contain leading zeroes.
// Constraints:
//  1 <= num.length <= 10⁵
//  num consists of digits.

func largestPalindromic(num string) string {
	count := [10]int{}
	for i := range num {
		count[int(num[i]-'0')]++
	}
	mid := -1 // the max odd count number
	buf := make([]byte, 0)
	for i := 9; i >= 0; i-- {
		if count[i]%2 == 1 && mid == -1 {
			mid = i
		}
		if i == 0 && len(buf) == 0 {
			break // all count[n]<=1 for n>=1
		}
		for x := 0; x < count[i]/2; x++ { // append half to buf
			buf = append(buf, byte(i)+'0')
		}
	}
	if mid == -1 && len(buf) == 0 {
		return "0"
	}
	if mid == -1 {
		for i := len(buf) - 1; i >= 0; i-- { // mirror the string
			buf = append(buf, buf[i])
		}
	} else {
		buf = append(buf, byte(mid)+'0')
		for i := len(buf) - 2; i >= 0; i-- {
			buf = append(buf, buf[i])
		}
	}
	return string(buf)
}

func main() {
	for _, v := range []struct {
		num, ans string
	}{
		{"444947137", "7449447"},
		{"9999977744222200", "997422090224799"},
		{"0009", "9"},
		{"0000000", "0"},
		{"00136", "6"},
	} {
		fmt.Println(largestPalindromic(v.num), v.ans)
	}
}

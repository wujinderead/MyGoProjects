package main

import (
	"fmt"
	"math/big"
)

// https://leetcode.com/problems/add-strings/

// Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.
// Note:
//   The length of both num1 and num2 is < 5100.
//   Both num1 and num2 contains only digits 0-9.
//   Both num1 and num2 does not contain any leading zero.
//   You must not use any built-in BigInteger library or convert the inputs to integer directly.

func addStrings(num1 string, num2 string) string {
	buf := make([]byte, max(len(num1), len(num2))+1)
	extra := 0
	i, j, ind := len(num1)-1, len(num2)-1, len(buf)-1
	for i>=0 || j>=0 {
		d := extra
		if i>=0 {
			d += int(num1[i]-'0')
		}
		if j>=0 {
			d += int(num2[j]-'0')
		}
		buf[ind] = byte(d%10)+'0'
		extra = d/10
		ind, i, j = ind-1, i-1, j-1
	}
	buf[0] = byte(extra)+'0'
	if buf[0] > '0' {
		return string(buf[0:])
	}
    return string(buf[1:])
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	for _, v := range [][2]string{
		{"9", "99"},
		{"1", "999"},
		{"987","4334"},
		{"111","222"},
		{"1", "8"},
		{"1", "9"},
		{"11", "222"},
		{"1234567890865432472377", "23636127638726378263782618361283687"},
	} {
		a, b := new(big.Int), new(big.Int)
		a.SetString(v[0], 10)
		b.SetString(v[1], 10)
		fmt.Println(addStrings(v[0], v[1])==a.Add(a, b).String())
	}
}

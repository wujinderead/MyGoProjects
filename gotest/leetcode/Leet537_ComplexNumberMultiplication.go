package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/complex-number-multiplication/submissions/

// Given two strings representing two complex numbers.
// You need to return a string representing their multiplication.
// Note i^2 = -1 according to the definition.
// Example 1:
//   Input: "1+1i", "1+1i"
//   Output: "0+2i"
//   Explanation: (1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i, and you need convert it to the form of 0+2i.
// Example 2:
//   Input: "1+-1i", "1+-1i"
//   Output: "0+-2i"
//   Explanation: (1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i, and you need convert it to the form of 0+-2i.
// Note:
//   The input strings will not have extra blank.
//   The input strings will be given in the form of a+bi,
//   where the integer a and b will both belong to the range of [-100, 100].
//   And the output should be also in this form.

func complexNumberMultiply(a string, b string) string {
	aa := strings.Split(a, "+")
	a1, _ := strconv.Atoi(aa[0])
	a2, _ := strconv.Atoi(aa[1][:len(aa[1])-1])
	bb := strings.Split(b, "+")
	b1, _ := strconv.Atoi(bb[0])
	b2, _ := strconv.Atoi(bb[1][:len(bb[1])-1])
	real := a1*b1 - a2*b2
	virtual := a1*b2 + a2*b1
	return fmt.Sprintf("%d+%di", real, virtual)
}

func main() {
	fmt.Println(complexNumberMultiply("1+1i", "1+1i"))
	fmt.Println(complexNumberMultiply("1+-1i", "1+-1i"))
}

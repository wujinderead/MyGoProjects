package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/simplified-fractions/

// Given an integer n, return a list of all simplified fractions between 0 and 1
// (exclusive) such that the denominator is less-than-or-equal-to n. The fractions
// can be in any order.
// Example 1:
//   Input: n = 2
//   Output: ["1/2"]
//   Explanation: "1/2" is the only unique fraction with a denominator less-than-or-equal-to 2.
// Example 2:
//   Input: n = 3
//   Output: ["1/2","1/3","2/3"]
// Example 3:
//   Input: n = 4
//   Output: ["1/2","1/3","1/4","2/3","3/4"]
//   Explanation: "2/4" is not a simplified fraction because it can be simplified to "1/2".
// Example 4:
//   Input: n = 1
//   Output: []
// Constraints:
//   1 <= n <= 100

func simplifiedFractions(n int) []string {
	ans := make([]string, 0)
	buf := make([]byte, 0)
	strs := make([]string, n)
	for i := range strs {
		strs[i] = strconv.Itoa(i+1)
	}
	for i:=1; i<n; i++ {
		for j:=i+1; j<=n; j++ {
			if gcd(i, j)==1 {
				buf = append(buf, strs[i-1]...)
				buf = append(buf, '/')
				buf = append(buf, strs[j-1]...)
				ans = append(ans, string(buf))
				buf = buf[:0]
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	if b==0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	fmt.Println(simplifiedFractions(1))
	fmt.Println(simplifiedFractions(2))
	fmt.Println(simplifiedFractions(3))
	fmt.Println(simplifiedFractions(4))
	fmt.Println(simplifiedFractions(5))
	fmt.Println(simplifiedFractions(6))
}
package main

import (
	"fmt"
	"strings"
)

// Given a non-empty string check if it can be constructed by taking a substring of it and
// appending multiple copies of the substring together. You may assume the given string
// consists of lowercase English letters only and its length will not exceed 10000.
// Example 1:
//   Input: "abab"
//   Output: True
//   Explanation: It's the substring "ab" twice.
// Example 2:
//   Input: "aba"
//   Output: False
// Example 3:
//   Input: "abcabcabcabc"
//   Output: True
//   Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)

// the most easiest way: check whether s in (s+s)[1:-1]
func repeatedSubstringPattern(s string) bool {
	s2 := (s + s)[1 : 2*len(s)-1]
	return strings.Index(s2, s) >= 0
}

func repeatedSubstringPattern1(s string) bool {
	n := len(s)
	if n == 1 {
		return false
	}
	factors := make([]int, 0)
	getFactors(n, &factors)
loop:
	for _, p := range factors {
		plen := n / p // 'p' parts, each part len 'plen'
		for i := 0; i < plen; i++ {
			for j := 1; j < p; j++ {
				if s[i+j*plen] != s[i] {
					continue loop
				}
			}
		}
		return true
	}
	return false
}

func getFactors(n int, factors *[]int) {
	if n == 1 {
		return
	}
	isprime := true
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			isprime = false
			*factors = append(*factors, i)
			for n%i == 0 {
				n = n / i
			}
			getFactors(n, factors)
			break
		}
	}
	if isprime {
		*factors = append(*factors, n)
	}
}

func main() {
	factors := make([]int, 0)
	getFactors(2, &factors)
	fmt.Println(factors)
	factors = make([]int, 0)
	getFactors(3, &factors)
	fmt.Println(factors)
	factors = make([]int, 0)
	getFactors(6, &factors)
	fmt.Println(factors)
	factors = make([]int, 0)
	getFactors(99, &factors)
	fmt.Println(factors)
	factors = make([]int, 0)
	getFactors(5*5*37*97, &factors)
	fmt.Println(factors)
	factors = make([]int, 0)
	getFactors(10007, &factors)
	fmt.Println(factors)

	fmt.Println(repeatedSubstringPattern("abab"))
	fmt.Println(repeatedSubstringPattern("aba"))
	fmt.Println(repeatedSubstringPattern("abcabcabcabc"))
	fmt.Println(repeatedSubstringPattern("aaaab"))
	fmt.Println(repeatedSubstringPattern("aaaaa"))
	fmt.Println(repeatedSubstringPattern("abcdeabcde"))
}

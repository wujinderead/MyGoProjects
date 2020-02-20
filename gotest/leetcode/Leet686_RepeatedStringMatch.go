package main

import (
	"fmt"
	"strings"
)

// Given two strings A and B, find the minimum number of times A has to be repeated
// such that B is a substring of it. If no such solution, return -1.
// For example, with A = "abcd" and B = "cdabcdab".
// Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it;
// and B is not a substring of A repeated two times ("abcdabcd").
// Note:
//   The length of A and B will be between 1 and 10000.

// t = ceil(len(b) / len(a))
// B can only be in either A*t or A*(t+1)
func repeatedStringMatch(A string, B string) int {
	if strings.Index(A, B) > -1 {
		return 1
	}
	if strings.Index(A+A, B) > -1 {
		return 2
	}
	ind := strings.Index(B, A)
	if ind < 0 {
		return -1
	}
	i := ind
	p := 0
	for i < len(B) && i+len(A)-1 < len(B) {
		if B[i:i+len(A)] != A {
			return -1
		}
		// fmt.Println("log1:", i, i+len(A))
		p++
		i += len(A)
	}
	// fmt.Println("log2:", B[i:], A[: len(B)-i], B[:ind], A[len(A)-ind:])
	if (i == len(B) || B[i:] == A[:len(B)-i]) && B[:ind] == A[len(A)-ind:] {
		if i < len(B) {
			p++
		}
		if ind > 0 {
			p++
		}
		return p
	}
	return -1

}

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
	fmt.Println(repeatedStringMatch("abcd", "dabcdabcdabcd"))
	fmt.Println(repeatedStringMatch("abcd", "abcdabcdabcdab"))
	fmt.Println(repeatedStringMatch("abcd", "abcdabcdabcd"))
}

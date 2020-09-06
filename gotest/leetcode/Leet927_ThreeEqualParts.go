package main

import "fmt"

// https://leetcode.com/problems/three-equal-parts/

// Given an array A of 0s and 1s, divide the array into 3 non-empty parts such that all of 
// these parts represent the same binary value.
// If it is possible, return any [i, j] with i+1 < j, such that:
//   A[0], A[1], ..., A[i] is the first part;
//   A[i+1], A[i+2], ..., A[j-1] is the second part, and
//   A[j], A[j+1], ..., A[A.length - 1] is the third part.
//   All three parts have equal binary value.
// If it is not possible, return [-1, -1].
// Note that the entire part is used when considering what binary value it represents.  
// For example, [1,1,0] represents 6 in decimal, not 3.  Also, leading zeros are allowed, 
// so [0,1,1] and [1,1] represent the same value.
// Example 1:
//   Input: [1,0,1,0,1]
//   Output: [0,3]
// Example 2:
//   Input: [1,1,0,1,1]
//   Output: [-1,-1]
// Note:
//   3 <= A.length <= 30000
//   A[i] == 0 or A[i] == 1

// count the number of 1's, denote as n, if n%3!=0, can't part. each part must have n/3 1's. 
// and the answer is determined to be the last part.
// e.g. [0,1,1,0,0,1,1,0,0,1,1,0], the result binary value must be "110", we can found it in previous array.
// if [0,1,1,0,0,1,1,0,1,1,0,0], the last value is "1100", 
// the second is "110" which can't be the answer.
func threeEqualParts(A []int) []int {
	// check 1's number
	n := 0
	for _, v := range A {
		n += v
	}
	if n==0 {
		return []int{0, len(A)-1}
	}
	if n%3!=0 {
		return []int{-1, -1}
	}

	n = n/3
	c := 0
	i := len(A)
	for c < n {      // A[i:] is the last value
		i--
		c += A[i]
	}

	j := i
	c = 0
	for c < n {      // A[j:] is the second part
		j--
		c += A[j]
	}
	for ki, kj := i, j; ki<len(A); ki, kj = ki+1, kj+1 {
		if A[ki] != A[kj] {   // check if equal
			return []int{-1, -1}
		}
	}

	m := j
	c = 0
	for c < n {      // A[m:] is the first part
		m--
		c += A[m]
	}
	for ki, kj := i, m; ki<len(A); ki, kj = ki+1, kj+1 {
		if A[ki] != A[kj] {   // check if equal
			return []int{-1, -1}
		}
	}
	return []int{m+len(A)-i-1, j+len(A)-i}
}

func main() {
	for _, v := range []struct{A, ans []int} {
		{[]int{1,0,1,0,1}, []int{0,3}},
		{[]int{1,1,0,1,1}, []int{-1,-1}},
		{[]int{0,0,1,1,0,1,1}, []int{-1,-1}},
		{[]int{0,0,1,1,1,0,0,0,0,0,1,1,1,0,0,1,1,1}, []int{4,13}},
		{[]int{0,0,1,1,1,0,0,0,0,0,1,1,1,0,0,1,1,1,0}, []int{5,14}},
		{[]int{0,0,1,1,1,0,0,0,0,0,1,1,1,0,1,1,1,0,0}, []int{-1,-1}},
		{[]int{0,0,0}, []int{0,2}},
	} {
		fmt.Println(threeEqualParts(v.A), v.ans)
	}
}
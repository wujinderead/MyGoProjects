package main

import "fmt"

// https://leetcode.com/problems/smallest-string-with-a-given-numeric-value/

// The numeric value of a lowercase character is defined as its position (1-indexed)
// in the alphabet, so the numeric value of a is 1, the numeric value of b is 2,
// the numeric value of c is 3, and so on.
// The numeric value of a string consisting of lowercase characters is defined as the sum
// of its characters' numeric values. For example, the numeric value of the string "abe"
// is equal to 1 + 2 + 5 = 8.
// You are given two integers n and k. Return the lexicographically smallest string
// with length equal to n and numeric value equal to k.
// Note that a string x is lexicographically smaller than string y if x comes before y
// in dictionary order, that is, either x is a prefix of y, or if i is the first position
// such that x[i] != y[i], then x[i] comes before y[i] in alphabetic order.
// Example 1:
//   Input: n = 3, k = 27
//   Output: "aay"
//   Explanation: The numeric value of the string is 1 + 1 + 25 = 27, and it is the
//     smallest string with such a value and length equal to 3.
// Example 2:
//   Input: n = 5, k = 73
//   Output: "aaszz"
// Constraints:
//   1 <= n <= 10^5
//   n <= k <= 26 * n

// e.g., for n=5, k=73, because 4*26>=72, so ans[0] can be 'a'
// then n=4, k=72, because 3*26>=71, so ans[1] can be 'a'
// then n=3, k=71, because 2*26<71, so ans[2] can't be 'a', it can be 's'
// then n=2, k=52, so last 2 letters is 'zz'
func getSmallestString(n int, k int) string {
	ans := make([]byte, n)
	var i int
	for i = 0; i < n-1; i++ {
		nn := n - i - 1
		if k-1 > nn*26 {
			break
		}
		ans[i] = 'a'
		k--
	}
	if k%26 != 0 {
		ans[i] = byte((k%26)-1) + 'a'
		i++
	}
	for i < n {
		ans[i] = 'z'
		i++
	}
	return string(ans)
}

func main() {
	for _, v := range []struct {
		n, k int
		ans  string
	}{
		{3, 27, "aay"},
		{5, 73, "aaszz"},
		{1, 1, "a"},
		{1, 2, "b"},
		{1, 26, "z"},
		{2, 27, "aa"},
		{3, 4, "aab"},
		{2, 52, "zz"},
	} {
		fmt.Println(getSmallestString(v.n, v.k), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/count-sorted-vowel-strings/

// Given an integer n, return the number of strings of length n that consist only of vowels
// (a, e, i, o, u) and are lexicographically sorted.
// A string s is lexicographically sorted if for all valid i, s[i] is the same as or comes
// before s[i+1] in the alphabet.
// Example 1:
//   Input: n = 1
//   Output: 5
//   Explanation: The 5 sorted strings that consist of vowels only are ["a","e","i","o","u"].
// Example 2:
//   Input: n = 2
//   Output: 15
//   Explanation: The 15 sorted strings that consist of vowels only are
//     ["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"].
//     Note that "ea" is not a valid string since 'e' comes after 'a' in the alphabet。
// Example 3:
//   Input: n = 33
//   Output: 66045
// Constraints:
//   1 <= n <= 50

func countVowelStrings(n int) int {
	// the number of string end as certain letter
	a, e, i, o, u := 1, 1, 1, 1, 1
	for x := 2; x <= n; x++ {
		a, e, i, o, u = a, a+e, a+e+i, a+e+i+o, a+e+i+o+u
	}
	return a + e + i + o + u
}

func main() {
	for _, v := range []struct {
		n, ans int
	}{
		{1, 5},
		{2, 15},
		{3, 35},
		{4, 70},
		{33, 66045},
	} {
		fmt.Println(countVowelStrings(v.n), v.ans)
	}
}

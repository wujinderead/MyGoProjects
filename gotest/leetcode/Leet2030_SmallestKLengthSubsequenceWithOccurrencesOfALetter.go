package main

import "fmt"

// https://leetcode.com/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter/

// You are given a string s, an integer k, a letter letter, and an integer repetition.
// Return the lexicographically smallest subsequence of s of length k that has the letter letter
// appear at least repetition times. The test cases are generated so that the letter appears in s
// at least repetition times.
// A subsequence is a string that can be derived from another string by deleting some or no characters
// without changing the order of the remaining characters.
// A string a is lexicographically smaller than a string b if in the first position where a and b differ,
// string a has a letter that appears earlier in the alphabet than the corresponding letter in b.
// Example 1:
//   Input: s = "leet", k = 3, letter = "e", repetition = 1
//   Output: "eet"
//   Explanation: There are four subsequences of length 3 that have the letter 'e' appear at least 1 time:
//     - "lee" (from "leet")
//     - "let" (from "leet")
//     - "let" (from "leet")
//     - "eet" (from "leet")
//     The lexicographically smallest subsequence among them is "eet".
// Example 2:
//   Input: s = "leetcode", k = 4, letter = "e", repetition = 2
//   Output: "ecde"
//   Explanation: "ecde" is the lexicographically smallest subsequence of length 4 that
//     has the letter "e" appear at least 2 times.
// Example 3:
//   Input: s = "bb", k = 2, letter = "b", repetition = 2
//   Output: "bb"
//   Explanation: "bb" is the only subsequence of length 2 that has the letter "b" appear at least 2 times.
// Constraints:
//   1 <= repetition <= k <= s.length <= 5 * 104
//   s consists of lowercase English letters.
//   letter is a lowercase English letter, and appears in s at least repetition times.

func smallestSubsequence(s string, k int, L byte, repetition int) string {
	// how many occurrence of L in s[i:]
	ll := make([]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		ll[i] = ll[i+1]
		if s[i] == L {
			ll[i]++
		}
	}
	stack := make([]int, k)
	size := 0  // stack size, stack[size-1] is the top element
	count := 0 // the count of letter in stack
	for i := 0; i < len(s); i++ {
		for size-1 >= 0 { // check if we can pop the stack top
			if s[i] >= s[stack[size-1]] { // can't improve subsequence's lexicographical order
				break
			}
			if size-1+len(s)-i < k { // if pop, no enough letters to make subsequence k-size
				break
			}
			if s[stack[size-1]] == L && count-1+ll[i+1] < repetition { // if pop L, no enough L to make repetition L's
				break
			}
			// we can pop here
			if s[stack[size-1]] == L { // if pop a L, decrement count
				count--
			}
			size--
		}
		// we may have popped some elements, we need to decide whether to push s[i]
		// first check if stack full (check size < k)
		if size < k && s[i] != L && k-size <= repetition-count {
			// if push s[i] letter, we will have no room for enough L, the don't push
			// e.g., s="aaaaee", k=4, L=e, repetition=2; when stack="aa", for s[2]='a',
			// we can't push it because we need push the last two "ee" later
			continue
		}
		if size < k {
			stack[size] = i // push the index
			size++
			if s[i] == L { // if push a letter, increment count
				count++
			}
		}
	}
	// make answer
	ans := make([]byte, k)
	for i := 0; i < k; i++ {
		ans[i] = s[stack[i]]
	}
	return string(ans)
}

func main() {
	for _, v := range []struct {
		s   string
		k   int
		l   byte
		r   int
		ans string
	}{
		{"leet", 3, 'e', 1, "eet"},
		{"leetcode", 4, 'e', 2, "ecde"},
		{"bb", 2, 'b', 2, "bb"},
		{"aaaaee", 2, 'e', 2, "ee"},
		{"aaaaee", 3, 'e', 2, "aee"},
		{"aaaaee", 4, 'e', 2, "aaee"},
		{"aaaaee", 4, 'e', 1, "aaae"},
		{"aaaaee", 2, 'e', 1, "ae"},
		{"mmmxmxymmm", 8, 'm', 4, "mmmmxmmm"},
		{"wuynymkihfdcbabefiiymnoyyytywzy", 16, 'y', 4, "abefiimnoyytywzy"},
	} {
		fmt.Println(smallestSubsequence(v.s, v.k, v.l, v.r), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/

// You are given a string s consisting only of characters 'a' and 'b'.
// You can delete any number of characters in s to make s balanced. s is balanced
// if there is no pair of indices (i,j) such that i < j and s[i] = 'b' and s[j]= 'a'.
// Return the minimum number of deletions needed to make s balanced.
// Example 1:
//   Input: s = "aababbab"
//   Output: 2
//   Explanation: You can either:
//     Delete the characters at 0-indexed positions 2 and 6 ("aababbab" -> "aaabbb"), or
//     Delete the characters at 0-indexed positions 3 and 6 ("aababbab" -> "aabbbb").
// Example 2:
//   Input: s = "bbaaaaabb"
//   Output: 2
//   Explanation: The only solution is to delete the first two characters.
// Constraints:
//   1 <= s.length <= 10^5
//   s[i] is 'a' or 'b'.

// another idea:
// let the min cost for prev string is c, then:
// if current char is b, the string is still valid, the cost is still c;
// if current char is a, we can:
//   keep this char, then we need remove all b's before, cost is bcount
//   remove a, the prev is still valid, so cost is c+1
// so the cost is min(c+1, bcount)
func minimumDeletions(s string) int {
	var alla, allb int
	// get a, b count
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			alla++
		} else {
			allb++
		}
	}
	var a, b int
	// move backward, check how many a's after b, these a's need to be removed
	min := len(s)
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			a++
			// remove b's before a, and all a's after
			if b+(alla-a) < min {
				min = b + (alla - a)
			}
		} else {
			if b+(alla-a) < min {
				min = b + (alla - a)
			}
			b++
		}
	}
	return min
}

func main() {
	for _, v := range []struct {
		s   string
		ans int
	}{
		{"aababbab", 2},
		{"bbaaaaabb", 2},
		{"ababaaaabbbbbaaababbbbbbaaabbaababbabbbbaabbbbaabbabbabaabbbababaa", 25},
		{"ab", 0},
		{"aa", 0},
		{"bb", 0},
		{"b", 0},
		{"ba", 1},
		{"baa", 1},
	} {
		fmt.Println(minimumDeletions(v.s), v.ans)
	}
}

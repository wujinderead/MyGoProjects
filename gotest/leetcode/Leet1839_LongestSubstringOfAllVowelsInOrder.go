package main

import "fmt"

// https://leetcode.com/problems/longest-substring-of-all-vowels-in-order/

// A string is considered beautiful if it satisfies the following conditions:
// Each of the 5 English vowels ('a', 'e', 'i', 'o', 'u') must appear at least once in it.
// The letters must be sorted in alphabetical order (i.e. all 'a's before 'e's, all 'e's before 'i's, etc.).
// For example, strings "aeiou" and "aaaaaaeiiiioou" are considered beautiful, but "uaeio", "aeoiu", and "aaaeeeooo" are not beautiful.
// Given a string word consisting of English vowels, return the length of the longest beautiful
// substring of word. If no such substring exists, return 0.
// A substring is a contiguous sequence of characters in a string.
// Example 1:
//   Input: word = "aeiaaioaaaaeiiiiouuuooaauuaeiu"
//   Output: 13
//   Explanation: The longest beautiful substring in word is "aaaaeiiiiouuu" of length 13.
// Example 2:
//   Input: word = "aeeeiiiioooauuuaeiou"
//   Output: 5
//   Explanation: The longest beautiful substring in word is "aeiou" of length 5.
// Example 3:
//   Input: word = "a"
//   Output: 0
//   Explanation: There is no beautiful substring, so return 0.
// Constraints:
//   1 <= word.length <= 5 * 10^5
//   word consists of characters 'a', 'e', 'i', 'o', and 'u'.

func longestBeautifulSubstring(word string) int {
	max := 0
	start := 0
	for start < len(word) {
		if word[start] != 'a' {
			start++
			continue
		}
		i := start + 1
		for i < len(word) {
			if (word[i-1] == 'a' && (word[i] == 'a' || word[i] == 'e')) ||
				(word[i-1] == 'e' && (word[i] == 'e' || word[i] == 'i')) ||
				(word[i-1] == 'i' && (word[i] == 'i' || word[i] == 'o')) ||
				(word[i-1] == 'o' && word[i] == 'o') {
				i++
			} else if (word[i-1] == 'o' || word[i-1] == 'u') && word[i] == 'u' {
				if i-start+1 > max {
					max = i - start + 1
				}
				i++
			} else {
				break
			}
		}
		start = i
	}
	return max
}

func main() {
	for _, v := range []struct {
		w   string
		ans int
	}{
		{"aeiaaioaaaaeiiiiouuuooaauuaeiu", 13},
		{"aeiou", 5},
		{"aeeou", 5},
		{"a", 0},
		{"aaaaeiou", 8},
		{"aaaaieo", 0},
		{"aeaeiou", 5},
	} {
		fmt.Println(longestBeautifulSubstring(v.w), v.ans)
	}
}

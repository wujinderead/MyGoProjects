package main

import "fmt"

// https://leetcode.com/problems/remove-letter-to-equalize-frequency/

// You are given a 0-indexed string word, consisting of lowercase English letters.
// You need to select one index and remove the letter at that index from word so that the frequency of
// every letter present in word is equal.
// Return true if it is possible to remove one letter so that the frequency of all letters in word are
// equal, and false otherwise.
// Note:
//   The frequency of a letter x is the number of times it occurs in the string.
//   You must remove exactly one letter and cannot chose to do nothing.
// Example 1:
//   Input: word = "abcc"
//   Output: true
//   Explanation: Select index 3 and delete it: word becomes "abc" and each
//     character has a frequency of 1.
// Example 2:
//   Input: word = "aazz"
//   Output: false
//   Explanation: We must delete a character, so either the frequency of "a" is 1
//     and the frequency of "z" is 2, or vice versa. It is impossible to make all present
//     letters have equal frequency.
// Constraints:
//   2 <= word.length <= 100
//   word consists of lowercase English letters only.

// so many edge case, so just following the instruction
func equalFrequency(word string) bool {
	count := [26]int{}
	for _, v := range word {
		count[int(v-'a')]++
	}
	for _, v := range word {
		count[int(v-'a')]-- // delete a word each time
		// check if all count equal
		can := true
		var first int
		for i := range count {
			if count[i] > 0 {
				if first == 0 {
					first = count[i]
				} else if count[i] != first {
					can = false
					break
				}
			}
		}
		if can { // find a valid deletion, return true
			return true
		}
		count[int(v-'a')]++
		// continue to next char
	}
	return false
}

func main() {
	for _, v := range []struct {
		s   string
		ans bool
	}{
		{"abcc", true},
		{"abbcc", true},
		{"aaaabbb", true},
		{"cccaaaabbb", true},
		{"cccaaaabbbddd", true},
		{"abc", true},
		{"aaaaa", true},
		{"aabbb", true},
		{"aabbbccc", false},
		{"aabbbcccd", false},
		{"aabbcc", false},
		{"aaabbb", false},
		{"aaaaabbb", false},
	} {
		fmt.Println(equalFrequency(v.s), v.ans)
	}
}

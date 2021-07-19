package main

import "fmt"

// https://leetcode.com/problems/number-of-wonderful-substrings/

// A wonderful string is a string where at most one letter appears an odd number of times.
// For example, "ccjjc" and "abab" are wonderful, but "ab" is not.
// Given a string word that consists of the first ten lowercase English letters ('a' through 'j'),
// return the number of wonderful non-empty substrings in word. If the same substring appears
// multiple times in word, then count each occurrence separately.
// A substring is a contiguous sequence of characters in a string.
// Example 1:
//   Input: word = "aba"
//   Output: 4
//   Explanation: The four wonderful substrings are underlined below:
//     - "aba" -> "a"
//     - "aba" -> "b"
//     - "aba" -> "a"
//     - "aba" -> "aba"
// Example 2:
//   Input: word = "aabb"
//   Output: 9
//   Explanation: The nine wonderful substrings are underlined below:
//     - "aabb" -> "a"
//     - "aabb" -> "aa"
//     - "aabb" -> "aab"
//     - "aabb" -> "aabb"
//     - "aabb" -> "a"
//     - "aabb" -> "abb"
//     - "aabb" -> "b"
//     - "aabb" -> "bb"
//     - "aabb" -> "b"
// Example 3:
//   Input: word = "he"
//   Output: 2
//   Explanation: The two wonderful substrings are underlined below:
//     - "he" -> "h"
//     - "he" -> "e"
// Constraints:
//   1 <= word.length <= 10^5
//   word consists of lowercase English letters from 'a' to 'j'.

// "a b c d e f g h i j" at most 10 letters
func wonderfulSubstrings(word string) int64 {
	var count int64 = 1
	mapp := make(map[int]int64)
	mask := 1 << uint(word[0]-'a') // mask is the prefix sum of occurrence
	mapp[mask] = 1
	mapp[0] = 1
	for i := 1; i < len(word); i++ {
		mask = mask ^ (1 << uint(word[i]-'a')) // XOR to make odd as 1, even as 0
		tmpmask := 1
		// tmpmask is single 1 mask, let x = mask ^ tmpmask,
		// so that mask ^ x = tmpmask, means we find a segment with single-odd substring
		for i := 0; i <= 9; i++ { // 0101  mask
			tmpmask = 1 << uint(i)      // 0001  tmpmask
			count += mapp[mask^tmpmask] // 0100  mask ^ tmpmask
		}
		count += mapp[mask] // same mask, result in none-odd substring
		mapp[mask] = mapp[mask] + 1
	}
	return count
}

func main() {
	for _, v := range []struct {
		w   string
		ans int
	}{
		{"aba", 4},
		{"aabb", 9},
		{"he", 2},
	} {
		fmt.Println(wonderfulSubstrings(v.w), v.ans)
	}
}

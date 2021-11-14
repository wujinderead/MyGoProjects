package main

import "fmt"

// https://leetcode.com/problems/vowels-of-all-substrings/

// Given a string word, return the sum of the number of vowels ('a', 'e', 'i', 'o', and 'u')
// in every substring of word.
// A substring is a contiguous (non-empty) sequence of characters within a string.
// Note: Due to the large constraints, the answer may not fit in a signed 32-bit integer.
// Please be careful during the calculations.
// Example 1:
//   Input: word = "aba"
//   Output: 6
//   Explanation:
//     All possible substrings are: "a", "ab", "aba", "b", "ba", and "a".
//     - "b" has 0 vowels in it
//     - "a", "ab", "ba", and "a" have 1 vowel each
//     - "aba" has 2 vowels in it
//     Hence, the total sum of vowels = 0 + 1 + 1 + 1 + 1 + 2 = 6.
// Example 2:
//   Input: word = "abc"
//   Output: 3
//   Explanation:
//     All possible substrings are: "a", "ab", "abc", "b", "bc", and "c".
//     - "a", "ab", and "abc" have 1 vowel each
//     - "b", "bc", and "c" have 0 vowels each
//     Hence, the total sum of vowels = 1 + 1 + 1 + 0 + 0 + 0 = 3.
// Example 3:
//   Input: word = "ltcd"
//   Output: 0
//   Explanation: There are no vowels in any substring of "ltcd".
// Example 4:
//   Input: word = "noosabasboosa"
//   Output: 237
//   Explanation: There are a total of 237 vowels in all the substrings.
// Constraints:
//   1 <= word.length <= 10^5
//   word consists of lowercase English letters.

// just check each vowel and check how many substrings can contain this substring
func countVowels(word string) int64 {
	sum := int64(0)
	for i, c := range word {
		if c == 'a' || c == 'i' || c == 'e' || c == 'o' || c == 'u' {
			sum += int64(i+1) * int64(len(word)-i)
		}
	}
	return sum
}

func main() {
	for _, v := range []struct {
		s   string
		ans int
	}{
		{"aba", 6},
		{"abc", 3},
		{"ltcd", 0},
		{"noosabasboosa", 237},
	} {
		fmt.Println(countVowels(v.s), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/check-if-string-is-a-prefix-of-array/

// Given a string s and an array of strings words, determine whether s is a prefix string of words.
// A string s is a prefix string of words if s can be made by concatenating the first k strings in words
// for some positive k no larger than words.length.
// Return true if s is a prefix string of words, or false otherwise.
// Example 1:
//   Input: s = "iloveleetcode", words = ["i","love","leetcode","apples"]
//   Output: true
//   Explanation:
//     s can be made by concatenating "i", "love", and "leetcode" together.
// Example 2:
//   Input: s = "iloveleetcode", words = ["apples","i","love","leetcode"]
//   Output: false
//   Explanation:
//     It is impossible to make s using a prefix of arr.
// Constraints:
//   1 <= words.length <= 100
//   1 <= words[i].length <= 20
//   1 <= s.length <= 1000
//   words[i] and s consist of only lowercase English letters.

// need to concatenate with whole word
func isPrefixString(s string, words []string) bool {
	var j, k int
	for i := range s {
		if k == len(words[j]) {
			j++
			k = 0
		}
		if j == len(words) {
			return false
		}
		if s[i] != words[j][k] {
			return false
		}
		k++
	}
	return k == len(words[j])
}

func main() {
	for _, v := range []struct {
		s   string
		w   []string
		ans bool
	}{
		{"iloveleetcode", []string{"i", "love", "leetcode", "apples"}, true},
		{"iloveleetcode", []string{"i", "love", "leetcode"}, true},
		{"iloveleetcode", []string{"i", "love", "leetcod"}, false},
		{"i", []string{"i"}, true},
		{"is", []string{"i"}, false},
		{"i", []string{"is"}, false},
		{"is", []string{"i", "s"}, true},
		{"is", []string{"i", "ss"}, false},
	} {
		fmt.Println(isPrefixString(v.s, v.w), v.ans)
	}
}

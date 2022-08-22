package main

import "fmt"

// https://leetcode.com/problems/shifting-letters-ii/

// You are given a string s of lowercase English letters and a 2D integer array shifts where
// shifts[i] = [starti, endi, directioni]. For every i, shift the characters in s from the index
// starti to the index endi (inclusive) forward if directioni = 1, or shift the characters
// backward if directioni = 0.
// Shifting a character forward means replacing it with the next letter in the alphabet (wrapping
// around so that 'z' becomes 'a'). Similarly, shifting a character backward means replacing it
// with the previous letter in the alphabet (wrapping around so that 'a' becomes 'z').
// Return the final string after all such shifts to s are applied.
// Example 1:
//   Input: s = "abc", shifts = [[0,1,0],[1,2,1],[0,2,1]]
//   Output: "ace"
//   Explanation: Firstly, shift the characters from index 0 to index 1 backward.
//     Now s = "zac".
//     Secondly, shift the characters from index 1 to index 2 forward. Now s = "zbd".
//     Finally, shift the characters from index 0 to index 2 forward. Now s = "ace".
// Example 2:
//   Input: s = "dztz", shifts = [[0,0,0],[1,1,1]]
//   Output: "catz"
//   Explanation: Firstly, shift the characters from index 0 to index 0 backward.
//     Now s = "cztz".
//     Finally, shift the characters from index 1 to index 1 forward. Now s = "catz".
// Constraints:
//   1 <= s.length, shifts.length <= 5 * 10⁴
//   shifts[i].length == 3
//   0 <= starti <= endi < s.length
//   0 <= directioni <= 1
//   s consists of lowercase English letters.

// use prefix to perform range update
func shiftingLetters(s string, shifts [][]int) string {
	prefix := make([]int, len(s)+1)
	for _, v := range shifts {
		if v[2] == 1 {
			prefix[v[0]] += 1
			prefix[v[1]+1] -= 1
		} else {
			prefix[v[0]] -= 1
			prefix[v[1]+1] += 1
		}
	}
	// accumulate to get update on each position
	for i := 1; i < len(prefix); i++ {
		prefix[i] += prefix[i-1]
	}
	ans := make([]byte, len(s))
	for i := range s {
		ans[i] = byte((((int(s[i]-'a')+prefix[i])%26)+26)%26) + 'a'
	}
	return string(ans)
}

func main() {
	for _, v := range []struct {
		s      string
		shifts [][]int
		ans    string
	}{
		{"abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}, "ace"},
		{"dztz", [][]int{{0, 0, 0}, {1, 1, 1}}, "catz"},
		{"dztz", [][]int{}, "dztz"},
	} {
		fmt.Println(shiftingLetters(v.s, v.shifts), v.ans)
	}
}

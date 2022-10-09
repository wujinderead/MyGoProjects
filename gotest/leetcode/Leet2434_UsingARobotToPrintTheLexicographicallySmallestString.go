package main

import "fmt"

// https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string/

// You are given a string s and a robot that currently holds an empty string t.
// Apply one of the following operations until s and t are both empty:
//   Remove the first character of a string s and give it to the robot. The robot
//     will append this character to the string t.
//   Remove the last character of a string t and give it to the robot. The robot
//     will write this character on paper.
// Return the lexicographically smallest string that can be written on the paper.
// Example 1:
//   Input: s = "zza"
//   Output: "azz"
//   Explanation: Let p denote the written string.
//     Initially p="", s="zza", t="".
//     Perform first operation three times p="", s="", t="zza".
//     Perform second operation three times p="azz", s="", t="".
// Example 2:
//   Input: s = "bac"
//   Output: "abc"
//   Explanation: Let p denote the written string.
//     Perform first operation twice p="", s="c", t="ba".
//     Perform second operation twice p="ab", s="c", t="".
//     Perform first operation p="ab", s="", t="c".
//     Perform second operation p="abc", s="", t="".
// Example 3:
//   Input: s = "bdda"
//   Output: "addb"
//   Explanation: Let p denote the written string.
//     Initially p="", s="bdda", t="".
//     Perform first operation four times p="", s="", t="bdda".
//     Perform second operation four times p="addb", s="", t="".
// Constraints:
//   1 <= s.length <= 10⁵
//   s consists of only English lowercase letters.

func robotWithString(s string) string {
	// get rightmin[i] = min(s[i], s[i+1], ...)
	rightmin := make([]byte, len(s))
	rightmin[len(s)-1] = s[len(s)-1]
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] < rightmin[i+1] {
			rightmin[i] = s[i]
		} else {
			rightmin[i] = rightmin[i+1]
		}
	}
	// main logic
	var stack, ans []byte
	i := 0
	for i < len(s) {
		// push chars in stack that <= current right minimal
		for len(stack) > 0 && stack[len(stack)-1] <= rightmin[i] {
			ans = append(ans, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
		for s[i] != rightmin[i] { // push to stack until find minimal char
			stack = append(stack, s[i])
			i++
		}
		ans = append(ans, s[i]) // push minimal char to answer
		i++
	}
	for len(stack) > 0 {
		ans = append(ans, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return string(ans)
}

func main() {
	for _, v := range []struct {
		s, ans string
	}{
		{"zza", "azz"},
		{"bac", "abc"},
		{"bdda", "addb"},
		{"s", "s"},
		{"fcadbfe", "abdceff"},
	} {
		fmt.Println(robotWithString(v.s), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/longest-substring-without-repeating-characters

// Given a string, find the length of the longest substring without repeating characters. 
// Example 1: 
//   Input: "abcabcbb"
//   Output: 3 
//   Explanation: The answer is "abc", with the length of 3. 
// Example 2: 
//   Input: "bbbbb"
//   Output: 1
//   Explanation: The answer is "b", with the length of 1.
// Example 3:  
//   Input: "pwwkew"
//   Output: 3
//   Explanation: The answer is "wke", with the length of 3. Note that the answer 
//     must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	if s=="" {
		return 0
	}
	l := 0
	max := 1
	last := [128]int{}  
	for i := range last {
		last[i] = -1
	}
	last[int(s[0])] = 0  // last occurrence of character
	for i:=1; i<len(s); i++ {
		ind := int(s[i])
		if last[ind]>-1 && l<=last[ind] {  // move l right to avoid duplicate
			l = last[ind]+1
		}
		last[ind] = i
		if i-l+1 > max {
			max = i-l+1
		}
	} 
	return max
}

func main() {
	for _,v := range []struct{str string; ans int} {
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"a", 1},
		{"aba", 2},
		{"", 0},
		{" ", 1},
	} {
		fmt.Println(lengthOfLongestSubstring(v.str), v.ans)
	}
}
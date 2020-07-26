package main

import "fmt"

// https://leetcode.com/problems/number-of-good-ways-to-split-a-string/

// You are given a string s, a split is called good if you can split s into 2 non-empty strings p and q 
// where its concatenation is equal to s and the number of distinct letters in p and q are the same.
// Return the number of good splits you can make in s.
// Example 1:
//   Input: s = "aacaba"
//   Output: 2
//   Explanation: There are 5 ways to split "aacaba" and 2 of them are good. 
//     ("a", "acaba") Left string and right string contains 1 and 3 different letters respectively.
//     ("aa", "caba") Left string and right string contains 1 and 3 different letters respectively.
//     ("aac", "aba") Left string and right string contains 2 and 2 different letters respectively (good split).
//     ("aaca", "ba") Left string and right string contains 2 and 2 different letters respectively (good split).
//     ("aacab", "a") Left string and right string contains 3 and 1 different letters respectively.
// Example 2:
//   Input: s = "abcd"
//   Output: 1
//   Explanation: Split the string as follows ("ab", "cd").
// Example 3:
//   Input: s = "aaaaa"
//   Output: 4
//   Explanation: All possible splits are good.
// Example 4:
//   Input: s = "acbadbaada"
//   Output: 2
// Constraints:
//   s contains only lowercase English letters.
//   1 <= s.length <= 10^5

func numSplits(s string) int {
	freq := [26]int{}
	right := 0     // distinct characters in right part
	for i:=0; i<len(s); i++ {
		freq[int(s[i]-'a')]++
		if freq[int(s[i]-'a')]==1 {
			right++
		}
	}
	count := 0
	left := 0       // distinct characters in right part
	leftfreq := [26]int{}
	for i:=0; i<len(s); i++ {
		freq[int(s[i]-'a')]--
		if freq[int(s[i]-'a')]==0 {
			right--
		}
		if leftfreq[int(s[i]-'a')] == 0 {
			leftfreq[int(s[i]-'a')] = 1
			left++
		}
		if left==right {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(numSplits("aacaba"), 2)
	fmt.Println(numSplits("abcd"), 1)
	fmt.Println(numSplits("aaaaa"), 4)
	fmt.Println(numSplits("acbadbaada"), 2)
	fmt.Println(numSplits("a"), 0)
}

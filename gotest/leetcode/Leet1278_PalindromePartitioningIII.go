package main

import "fmt"

// https://leetcode.com/problems/palindrome-partitioning-iii/

// You are given a string s containing lowercase letters and an integer k. You need to :
//   First, change some characters of s to other lowercase English letters.
//   Then divide s into k non-empty disjoint substrings such that each substring is palindrome.
// Return the minimal number of characters that you need to change to divide the string.
// Example 1:
//   Input: s = "abc", k = 2
//   Output: 1
//   Explanation: You can split the string into "ab" and "c", and change 1 character
//     in "ab" to make it palindrome.
// Example 2:
//   Input: s = "aabbc", k = 3
//   Output: 0
//   Explanation: You can split the string into "aa", "bb" and "c", all of them are palindrome.
// Example 3:
//   Input: s = "leetcode", k = 8
//   Output: 0
// Constraints:
//   1 <= k <= s.length <= 100.
//   s only contains lowercase English letters.

func palindromePartition(s string, k int) int {
	// get the cost to make substring palindromic
	cost := make(map[[2]int]int)
	for i := 1; i < len(s)-1; i++ {
		// s[i] as center
		cur := 0
		l := 1
		for i-l >= 0 && i+l < len(s) {
			if s[i-l] != s[i+l] {
				cur++
			}
			cost[[2]int{i - l, i + l}] = cur
			l++
		}
	}
	for i := 0; i < len(s)-1; i++ {
		cur := 0
		l := 0
		for i-l >= 0 && i+1+l < len(s) {
			if s[i-l] != s[i+1+l] {
				cur++
			}
			cost[[2]int{i - l, i + 1 + l}] = cur
			l++
		}
	}
	// let dp(s[0:], k) be the minimal change split s to k palindromic parts, then we have candidates:
	// cost(s[0]) + dp(s[1:], k-1) , cost(s[0...1]) + dp(s[2:], k-1) , cost(s[0...2]) + dp(s[3:], k-1)
	// we want the minimal value. base case dp(s[x:], 1) = cost(s[x:])
	old, new := make([]int, len(s)), make([]int, len(s))
	for i := len(s) - 2; i >= 0; i-- {
		old[i] = cost[[2]int{i, len(s) - 1}] // dp(s, 1)
	}
	for i := 2; i <= k; i++ { // split i to parts
		for j := 0; j <= len(s)-i; j++ { // dp(s[j:], i)
			curmin := 0xffffffff
			// dp(s[j:], i) = cost(s[j...e]) + dp(s[e+1:], i-1)
			// we need len(s[e+1:])>=i-1, so len(s)-(e+1) >= i-1, e <= len(s)-i
			for e := j; e <= len(s)-i; e++ {
				curmin = min(curmin, cost[[2]int{j, e}]+old[e+1])
			}
			new[j] = curmin
		}
		old, new = new, old
	}
	return old[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(palindromePartition("abc", 2))
	fmt.Println(palindromePartition("aabbc", 3))
	fmt.Println(palindromePartition("leetcode", 8))
	fmt.Println(palindromePartition("a", 1))
	fmt.Println(palindromePartition("aa", 1))
	fmt.Println(palindromePartition("ab", 1))
	fmt.Println(palindromePartition("ab", 2))
}

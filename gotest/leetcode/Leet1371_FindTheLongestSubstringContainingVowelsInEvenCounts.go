package main

import "fmt"

// https://leetcode.com/problems/find-the-longest-substring-containing-vowels-in-even-counts/

// Given the string s, return the size of the longest substring containing each vowel an 
// even number of times. That is, 'a', 'e', 'i', 'o', and 'u' must appear an even number of times. 
// Example 1: 
//   Input: s = "eleetminicoworoep"
//   Output: 13
//   Explanation: The longest substring is "leetminicowor" which contains two each 
//     of the vowels: e, i and o and zero of the vowels: a and u.
// Example 2: 
//   Input: s = "leetcodeisgreat"
//   Output: 5
//   Explanation: The longest substring is "leetc" which contains two e's.
// Example 3:  
//   Input: s = "bcbcbc"
//   Output: 6
//   Explanation: In this case, the given string "bcbcbc" is the longest because 
//     all vowels: a, e, i, o and u appear zero times.
// Constraints: 
//   1 <= s.length <= 5 x 10^5 
//   s contains only lowercase English letters. 

// use a bitmask to represent the even or odd occurrence of vowel letters in prefix. 
// if we find two same masks, the substring between them are all even occurrence of vowel letters. 
func findTheLongestSubstring(s string) int {
	mapp := [32]int{}        // for 5 vowels, need 2^5=32 bitmasks
	for i := range mapp {
		mapp[i] = len(s)
	}
	mapp[0] = -1
	mask := 0
	ans := 0
	for i:=0; i<len(s); i++ {
		if bit(s[i])>4 {
			ans = max(ans, i-mapp[mask])
			continue
		}
		mask = mask ^ (1<<bit(s[i]))
		ans = max(ans, i-mapp[mask])
		mapp[mask] = min(mapp[mask], i)   // set to first occurrence of this mask
	}
	return ans
}

func bit(b byte) uint {
	switch b {
	case 'a': return 0
	case 'e': return 1
	case 'i': return 2
	case 'o': return 3
	case 'u': return 4
	}
	return 5
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"), 13)
	fmt.Println(findTheLongestSubstring("leetcodeisgreat"), 5)
	fmt.Println(findTheLongestSubstring("bcbcbc"), 6)
	fmt.Println(findTheLongestSubstring("ccacc"), 2)
}
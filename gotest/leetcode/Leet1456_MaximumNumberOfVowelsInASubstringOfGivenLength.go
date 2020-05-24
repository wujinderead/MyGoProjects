package main

import "fmt"

// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/

// Given a string s and an integer k.
// Return the maximum number of vowel letters in any substring of s with length k.
// Vowel letters in English are (a, e, i, o, u).
// Example 1:
//   Input: s = "abciiidef", k = 3
//   Output: 3
//   Explanation: The substring "iii" contains 3 vowel letters.
// Example 2:
//   Input: s = "aeiou", k = 2
//   Output: 2
//   Explanation: Any substring of length 2 contains 2 vowels.
// Example 3:
//   Input: s = "leetcode", k = 3
//   Output: 2
//   Explanation: "lee", "eet" and "ode" contain 2 vowels.
// Example 4:
//   Input: s = "rhythms", k = 4
//   Output: 0
//   Explanation: We can see that s doesn't have any vowel letters.
// Example 5:
//   Input: s = "tryhard", k = 4
//   Output: 1
// Constraints:
//   1 <= s.length <= 10^5
//   s consists of lowercase English letters.
//   1 <= k <= s.length

func maxVowels(s string, k int) int {
	max, cur := 0, 0
	for i:=0; i<k; i++ {
		if s[i]=='a' || s[i]=='e' || s[i]=='i' || s[i]=='o' || s[i]=='u' {
			max++
		}
	}
	cur = max
	for i:=k; i<len(s); i++ {
		if s[i]=='a' || s[i]=='e' || s[i]=='i' || s[i]=='o' || s[i]=='u' {
			cur++
		}
		if s[i-k]=='a' || s[i-k]=='e' || s[i-k]=='i' || s[i-k]=='o' || s[i-k]=='u' {
			cur--
		}
		if cur>max {
			max = cur
		}
	}
	return max
}

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("aeiou", 2))
	fmt.Println(maxVowels("leetcode", 3))
	fmt.Println(maxVowels("rhythms", 4))
	fmt.Println(maxVowels("tryhard", 4))
}
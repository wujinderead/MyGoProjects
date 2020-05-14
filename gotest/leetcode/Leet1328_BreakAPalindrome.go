package main

import "fmt"

// https://leetcode.com/problems/break-a-palindrome/

// Given a palindromic string palindrome, replace exactly one character by any lowercase
// English letter so that the string becomes the lexicographically smallest
// possible string that isn't a palindrome. After doing so, return the final string.
// If there is no way to do so, return the empty string.
// Example 1:
//   Input: palindrome = "abccba"
//   Output: "aaccba"
// Example 2:
//   Input: palindrome = "a"
//   Output: ""
// Constraints:
//   1 <= palindrome.length <= 1000
//   palindrome consists of only lowercase English letters.

func breakPalindrome(palindrome string) string {
	if len(palindrome)==1 {
		return ""
	}
    buf := []byte(palindrome)
    for i:=0; i<len(palindrome)/2; i++ {
    	if buf[i] != 'a' {
    		buf[i] = 'a'
    		return string(buf)
		}
	}
	buf[len(buf)-1] = 'b'
	return string(buf)
}

func main() {
	fmt.Println(breakPalindrome("abccba"))
	fmt.Println(breakPalindrome("a"))
	fmt.Println(breakPalindrome("b"))
	fmt.Println(breakPalindrome("bb"))
	fmt.Println(breakPalindrome("aa"))
	fmt.Println(breakPalindrome("aaa"))
	fmt.Println(breakPalindrome("aba"))
	fmt.Println(breakPalindrome("aabaa"))
	fmt.Println(breakPalindrome("cac"))
	fmt.Println(breakPalindrome("aca"))
}
package main

import "fmt"

// https://leetcode.com/problems/reverse-words-in-a-string/

// Given an input string, reverse the string word by word.
// Example 1:
//   Input: "the sky is blue"
//   Output: "blue is sky the"
// Example 2:
//   Input: "  hello world!  "
//   Output: "world! hello"
//   Explanation: Your reversed string should not contain leading or trailing spaces.
// Example 3:
//   Input: "a good   example"
//   Output: "example good a"
//   Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.
// Note:
//   A word is defined as a sequence of non-space characters.
//   Input string may contain leading or trailing spaces. However, your reversed string should not
//     contain leading or trailing spaces.
//   You need to reduce multiple spaces between two words to a single space
//     in the reversed string.
// Follow up:
// For C programmers, try to solve it in-place in O(1) extra space.

func reverseWords(s string) string {
	buf := make([]byte, 0, len(s)/2)
	i := len(s)-1
	count := 0
	for i>=0 {
		for i>=0 && s[i]==' ' {
			i--
		}
		end := i
		for i>=0 && s[i]!=' ' {
			i--
		}
		if end>i {
			buf = append(buf, s[i+1: end+1]...)
			buf = append(buf, ' ')
			count++
		}
	}
	if count>0 {
		buf = buf[:len(buf)-1]
	}
	return string(buf)
}

func main() {
	fmt.Printf("%q\n", reverseWords("the sky is blue"))
	fmt.Printf("%q\n", reverseWords("  hello world!  "))
	fmt.Printf("%q\n", reverseWords("    hello   world!  "))
	fmt.Printf("%q\n", reverseWords("a good   example"))
	fmt.Printf("%q\n", reverseWords(" "))
	fmt.Printf("%q\n", reverseWords("   "))
	fmt.Printf("%q\n", reverseWords("  gg  "))
	fmt.Printf("%q\n", reverseWords("     gg  jj    "))
	fmt.Printf("%q\n", reverseWords("   gg  jj"))
	fmt.Printf("%q\n", reverseWords("   a"))
	fmt.Printf("%q\n", reverseWords("a   "))
	fmt.Printf("%q\n", reverseWords("   a   "))
	fmt.Printf("%q\n", reverseWords("   a   a"))
	fmt.Printf("%q\n", reverseWords("   a   a  "))
}
package leetcode

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}
	maxlen := 0
	maxind := -1
	for i := 0; i < len(s); i++ {
		low, high := i, i
		for low-1 >= 0 && high+1 < len(s) && s[low-1] == s[high+1] {
			low--
			high++
		}
		if high-low+1 > maxlen {
			maxlen = high - low + 1
			maxind = low
		}
	}
	for i := 0; i < len(s)-1; i++ {
		l := 0
		for i-l >= 0 && i+1+l < len(s) && s[i-l] == s[i+1+l] {
			l++
		}
		if 2*l > maxlen {
			maxlen = 2 * l
			maxind = i - l + 1
		}
	}
	return s[maxind : maxind+maxlen]
}

func main() {
	fmt.Println(longestPalindrome("xxabay"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("abcba"))
}

package generalized_suffix_tree

import (
	"fmt"
	"testing"
)

func TestLongestCommon(t *testing.T) {
	strs := [][]string{{"xabxac", "abcabxabcd", "abxa"},
		{"xabxaabxa", "babxba", "abx"},
		{"GeeksforGeeks", "GeeksQuiz", "Geeks"},
		{"OldSite:GeeksforGeeks.org", "NewSite:GeeksQuiz.com", "Site:Geeks"},
		{"abcde", "fgefg", "e"},
		{"pqrst", "uvwxyz", ""}}
	for i := range strs {
		a, b, exp := strs[i][0], strs[i][1], strs[i][2]
		as, bs, length := longestCommonSubstring(a, b)
		if a[as:as+length] != exp || b[bs:bs+length] != exp {
			t.Errorf("error for '%s' and '%s', expect '%s', got '%s'",
				a, b, exp, a[as:as+length])
		}
		fmt.Println(as, bs, length)
		fmt.Println(a, b, exp)
		fmt.Println()
	}
}

func TestLongestPalindromic(t *testing.T) {
	strs := [][]string{{"cabbaabb", "bbaabb"},
		{"forgeeksskeegfor", "geeksskeeg"},
		{"abcde", "a"},
		{"abcdae", "a"},
		{"abacd", "aba"},
		{"abcdc", "cdc"},
		{"abacdfgdcaba", "aba"},
		{"xyabacdfgdcaba", "aba"},
		{"xababayz", "ababa"},
		{"xabax", "xabax"},
		{"", ""}}
	for i := range strs {
		str, exp := strs[i][0], strs[i][1]
		start, length := longestPalindromicSubstring(str)
		if str[start:start+length] != exp {
			t.Errorf("error for '%s', expect '%s', got '%s'",
				str, exp, str[start:start+length])
		}
		fmt.Println(start, length)
		fmt.Println(str, str[start:start+length])
		fmt.Println()
	}
}

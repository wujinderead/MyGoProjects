package longest

import (
	"fmt"
	"testing"
)

func TestLongestPalindromicSubarray(t *testing.T) {
	cases := [][]interface{}{
		{"cabbaabb", "bbaabb"},
		{"forgeeksskeegfor", "geeksskeeg"},
		{"abcde", "a"},
		{"abcdae", "a"},
		{"abacd", "aba"},
		{"xababayz", "ababa"},
		{"xabaabayz", "abaaba"},
		{"", ""},
		{"a", "a"},
		{"aa", "aa"},
		{"ab", "a"},
	}
	for i := range cases {
		str := cases[i][0].(string)
		expect := cases[i][1].(string)
		re := longestPalindromicSubarray(str)
		fmt.Println(re, expect)
		if len(re) != len(expect) {
			t.Error(str, expect, re)
		}
	}
}

func TestLongestCommonSubarray(t *testing.T) {
	cases := [][]interface{}{
		{"xabxac", "abcabxabcd", "abxa"},
		{"xabxaabxa", "babxba", "abx"},
		{"GeeksforGeeks", "GeeksQuiz", "Geeks"},
		{"OldSite:GeeksforGeeks.org", "NewSite:GeeksQuiz.com", "Site:Geeks"},
		{"abcde", "fghie", "e"},
		{"pqrst", "uvwxyz", ""},
		{"a", "bcde", ""},
		{"a", "bcade", "a"},
		{"adsd", "", ""},
	}
	for i := range cases {
		a := cases[i][0].(string)
		b := cases[i][1].(string)
		exp := cases[i][2].(string)
		re := longestCommonSubarray(a, b)
		fmt.Println(re, exp)
		if len(re) != len(exp) {
			t.Error(a, b, re)
		}
	}
}

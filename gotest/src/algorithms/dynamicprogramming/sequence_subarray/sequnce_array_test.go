package longest

import (
	"fmt"
	"testing"
)

func TestLongestIncreasingSubsequence(t *testing.T) {
	cases := [][]interface{}{
		{[]int{10, 22, 9, 33, 21, 50, 41, 60, 80}, 6},
		{[]int{50, 3, 10, 7, 40, 80}, 4},
		{[]int{3, 2}, 1},
	}
	for i := range cases {
		seq := cases[i][0].([]int)
		expect := cases[i][1].(int)
		re := longestIncreasingSubsequence(seq)
		if re != expect {
			t.Error(seq, expect, re)
		}
	}
}

func TestLongestPalindromicSubsequence(t *testing.T) {
	cases := [][]interface{}{
		{"", 0},
		{"a", 1},
		{"aa", 2},
		{"ab", 1},
		{"aba", 3},
		{"abca", 3},
		{"aaca", 3},
		{"BBABCBCAB", 7},
	}
	for i := range cases {
		str := cases[i][0].(string)
		expect := cases[i][1].(int)
		re := longestPalindromicSubsequence(str)
		if re != expect {
			t.Error(str, expect, re)
		}
	}
	// test lps o(n) space
	for i := range cases {
		str := cases[i][0].(string)
		expect := cases[i][1].(int)
		re := longestPalindromicSubsequenceSpaceOn(str)
		if re != expect {
			t.Error(str, expect, re)
		}
	}
}

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

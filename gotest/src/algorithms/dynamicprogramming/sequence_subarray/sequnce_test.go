package longest

import (
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

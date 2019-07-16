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

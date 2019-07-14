package set_problems

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	cases := [][]interface{}{
		// use [1,2,3] to get 4, 4 ways: {1,1,1,1}, {1,1,2}, {2,2}, {1,3}
		{[]int{1, 2, 3}, 4, 4},
		// use [2,5,3,6] to get 10, 5 ways: {2,2,2,2,2}, {2,2,3,3}, {2,2,6}, {2,3,5}, {5,5}
		{[]int{2, 5, 3, 6}, 10, 5},
	}
	for i := range cases {
		coins := cases[i][0].([]int)
		target := cases[i][1].(int)
		expect := cases[i][2].(int)
		re := coinChange(coins, target)
		if re != expect {
			t.Errorf("expect: %d, got: %d", expect, re)
		}
	}
}

func TestSubsetSum(t *testing.T) {
	cases := [][]interface{}{
		{[]int{3, 34, 4, 12, 5, 2}, 9, true}, // 4, 5
		{[]int{2, 6, 5, 3}, 10, true},        // 2, 5, 3
		{[]int{2, 6, 5, 3}, 15, false},       // cant get 15
	}
	for i := range cases {
		coins := cases[i][0].([]int)
		target := cases[i][1].(int)
		expect := cases[i][2].(bool)
		re := subsetSum(coins, target)
		if re != expect {
			t.Errorf("expect: %v, got: %v", expect, re)
		}
	}
}

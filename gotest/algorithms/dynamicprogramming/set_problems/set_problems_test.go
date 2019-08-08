package set_problems

import (
	"fmt"
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
		{[]int{2, 6, 5, 3}, 15, false},       // can't get 15
	}
	for i := range cases {
		set := cases[i][0].([]int)
		target := cases[i][1].(int)
		expect := cases[i][2].(bool)
		re := subsetSum(set, target)
		if re != expect {
			t.Errorf("expect: %v, got: %v", expect, re)
		}
	}
}

func TestPerfectSum(t *testing.T) {
	cases := [][]interface{}{
		{[]int{1, 2, 3, 4, 5}, 10, [][]int{{1, 2, 3, 4}, {2, 3, 5}, {1, 4, 5}}},
		{[]int{2, 3, 5, 6, 8, 10}, 10, [][]int{{5, 2, 3}, {2, 8}, {10}}},
	}
	for i := range cases {
		set := cases[i][0].([]int)
		target := cases[i][1].(int)
		expect := cases[i][2].([][]int)
		re := perfectSum(set, target)
		fmt.Println(re, expect)
	}
}

func TestMatrixChainMultiply(t *testing.T) {
	cases := [][]interface{}{
		{[]int{40, 20, 30, 10, 30}, 26000},
		{[]int{10, 20, 30, 40, 30}, 30000},
		{[]int{10, 30, 5, 60}, 4500},
		{[]int{10, 30, 5}, 1500},
	}
	for i := range cases {
		mats := cases[i][0].([]int)
		expect := cases[i][1].(int)
		re := matrixChainMultiplication(mats)
		if re != expect {
			t.Errorf("expect: %v, got: %v", expect, re)
		}
	}
}

func Test01Knapsack(t *testing.T) {
	cases := [][]interface{}{
		{[]int{60, 100, 120}, []int{10, 20, 30}, 50, 220},
		{[]int{3, 4, 5, 6}, []int{2, 3, 4, 5}, 8, 10},
	}
	for i := range cases {
		values := cases[i][0].([]int)
		weights := cases[i][1].([]int)
		W := cases[i][2].(int)
		expect := cases[i][3].(int)
		re := knapsack01(weights, values, W)
		if re != expect {
			t.Errorf("expect: %v, got: %v", expect, re)
		}
	}
}

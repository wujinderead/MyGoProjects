package main

import "fmt"

// https://leetcode.com/problems/minimum-money-required-before-transactions/

// You are given a 0-indexed 2D integer array transactions, where transactions[i] = [costi, cashbacki].
// The array describes transactions, where each transaction must be completed exactly once in some order.
// At any given moment, you have a certain amount of money. In order to complete transaction i,
// money >= costi must hold true. After performing a transaction, money becomes money - costi + cashbacki.
// Return the minimum amount of money required before any transaction so that all of the transactions can
// be completed regardless of the order of the transactions.
// Example 1:
//   Input: transactions = [[2,1],[5,0],[4,2]]
//   Output: 10
//   Explanation:
//     Starting with money = 10, the transactions can be performed in any order.
//     It can be shown that starting with money < 10 will fail to complete all transactions in some order.
// Example 2:
//   Input: transactions = [[3,0],[0,3]]
//   Output: 3
//   Explanation:
//     - If transactions are in the order [[3,0],[0,3]], the minimum money required to
//       complete the transactions is 3.
//     - If transactions are in the order [[0,3],[3,0]], the minimum money required to
//       complete the transactions is 0.
//     Thus, starting with money = 3, the transactions can be performed in any order.
// Constraints:
//   1 <= transactions.length <= 10⁵
//   transactions[i].length == 2
//   0 <= costi, cashbacki <= 10⁹

// we actually want the minimal initial money of the most-cost order.
// to cost most, we firstly do the losing TXs, then the earning TXs.
// for losing TXs, make the TX with largest cashback be the last TX;
// for earning TXs, make the TX with largest cost be the first TX, i.e. in this manner:
// [losingTX, ..., losingTX_with_max_cashback, earningTX_with_max_cost, ..., earningTX]
func minimumMoney(transactions [][]int) int64 {
	var lost, x int
	for _, t := range transactions {
		lost += max(t[0]-t[1], 0) // accumulate lost money in losing TXs
		// find the smaller of max_cashback_in_losingTXs or max_cost_in_earning_TXs
		x = max(x, min(t[0], t[1]))
	}
	return int64(lost + x)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		transactions [][]int
		ans          int64
	}{
		{[][]int{{2, 1}, {5, 0}, {4, 2}}, 10},
		{[][]int{{3, 0}, {0, 3}}, 3},
	} {
		fmt.Println(minimumMoney(v.transactions), v.ans)
	}
}

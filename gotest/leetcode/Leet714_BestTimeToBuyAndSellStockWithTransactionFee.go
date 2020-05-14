package main

import "fmt"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

// Your are given an array of integers prices, for which the i-th element is the
// price of a given stock on day i; and a non-negative integer fee representing a
// transaction fee. You may complete as many transactions as you like, but you need
// to pay the transaction fee for each transaction. You may not buy more than 1 share
// of a stock at a time (ie. you must sell the stock share before you buy again.)
// Return the maximum profit you can make.
// Example 1:
//   Input: prices = [1, 3, 2, 8, 4, 9], fee = 2
//   Output: 8
//   Explanation: The maximum profit can be achieved by:
//    Buying at prices[0] = 1 Selling at prices[3] = 8
//    Buying at prices[4] = 4 Selling at prices[5] = 9
//    The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
// Note:
//   0 < prices.length <= 50000.
//   0 < prices[i] < 50000.
//   0 <= fee < 50000.

func maxProfit(prices []int, fee int) int {
	buy, sell := -prices[0]-fee, 0
	fmt.Println(buy, sell)
	for i:=1; i<len(prices); i++ {
		buy = max(buy, sell-prices[i]-fee)
		sell = max(sell, buy+prices[i])
		fmt.Println(buy, sell)
	}
	return sell
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	fmt.Println()
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9, 2}, 2))
	fmt.Println()
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9}, 0))
	fmt.Println()
	fmt.Println(maxProfit([]int{1, 3, 2, 8, 4, 9, 2}, 0))
	fmt.Println()
	fmt.Println(maxProfit([]int{1, 4, 3, 8, 6, 9}, 2))
	fmt.Println()
	fmt.Println(maxProfit([]int{1, 5, 2, 8, 6, 9}, 2))
	fmt.Println()
	fmt.Println(maxProfit([]int{1, 4, 3, 8, 6, 9}, 0))
}
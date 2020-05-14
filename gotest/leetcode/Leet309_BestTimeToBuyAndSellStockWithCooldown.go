package main

import "fmt"

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

// Say you have an array for which the ith element is the price of a given stock on day i.
// Design an algorithm to find the maximum profit. You may complete as many transactions as
// you like (ie, buy one and sell one share of the stock multiple times) with the
// following restrictions:
// You may not engage in multiple transactions at the same time (ie, you must sell the
//   stock before you buy again).
// After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
// Example:
//   Input: [1,2,3,0,2]
//   Output: 3
//   Explanation: transactions = [buy, sell, cooldown, buy, sell]

// dp process:
// buy[i] = max(buy[i-1], sell[i-2]-price)   // with cool down, we can only use the profit of i-2 day
// sell[i] = max(sell[i-1], buy[i-1]+price)
func maxProfit(prices []int) int {
	if len(prices)<2 {
		return 0
	}
	buy, sell := -prices[0], 0
	prev_sell := 0
    for i:=1; i<len(prices); i++ {
    	prev_buy := buy
		buy = max(buy, prev_sell-prices[i])
		prev_sell = sell
		sell = max(sell, prev_buy+prices[i])
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
	fmt.Println(maxProfit([]int{1,2,3,0,2}))
}
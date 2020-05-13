package main

import (
	"fmt"
)

// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/

// Say you have an array for which the ith element is the price of a given stock on day i.
// Design an algorithm to find the maximum profit. You may complete at most two transactions.
// Note: You may not engage in multiple transactions at the same time (i.e., you
// must sell the stock before you buy again).
// Example 1:
//   Input: [3,3,5,0,0,3,1,4]
//   Output: 6
//   Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
//     Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.
// Example 2:
//   Input: [1,2,3,4,5]
//   Output: 4
//   Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
//     Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
//     engaging multiple transactions at the same time. You must sell before buying again.
// Example 3:
//   Input: [7,6,4,3,1]
//   Output: 0
//   Explanation: In this case, no transaction is done, i.e. max profit = 0.

// split the array to 2 parts. find the best result for 2 part and sum it.
// O(N) time, O(N) space. there is O(N) time, O(1) space solution, see:
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/39615/My-explanation-for-O(N)-solution!
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/149383/Easy-DP-solution-using-state-machine-O(n)-time-complexity-O(1)-space-complexity
func maxProfit(prices []int) int {
    if len(prices)==0 {
    	return 0
	}
	// the left part max, right part max
    left, right := make([]int, len(prices)), make([]int, len(prices))
    p := prices[0]
    for i:=1; i<len(prices); i++ {
    	if prices[i]-p > left[i-1] {
			left[i] = prices[i]-p
		} else {
			left[i] = left[i-1]
		}
		if prices[i] < p {
			p = prices[i]
		}
	}
	p = prices[len(prices)-1]
	for i:=len(prices)-2; i>=0; i-- {
		if p-prices[i] > right[i+1] {
			right[i] = p-prices[i]
		} else {
			right[i] = right[i+1]
		}
		if prices[i]>p {
			p = prices[i]
		}
	}
	p = left[len(prices)-1]    // the whole array max (one transaction)
	for i:=0; i+1<len(prices); i++ {
		if left[i]+right[i+1]>p {
			p = left[i]+right[i+1]
		}
	}
	return p
}

func main() {
	fmt.Println(maxProfit([]int{3,3,5,0,0,3,1,4}))
	fmt.Println(maxProfit([]int{1,2,3,4,5}))
	fmt.Println(maxProfit([]int{7,6,4,3,1}))
	fmt.Println(maxProfit([]int{6,1,3,2,4,7}))
	fmt.Println(maxProfit([]int{1,2,4,2,5,7,2,4,9,0}))
	fmt.Println(maxProfit([]int{11,4,1,12,7,13,15,9,2,8,3,5,10,14,6,0}))
}
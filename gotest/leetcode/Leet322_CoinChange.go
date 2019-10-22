package leetcode

import "fmt"

// https://leetcode.com/problems/coin-change/
// You are given coins of different denominations and a total amount of money amount.
// Write a function to compute the fewest number of coins that you need to make up that amount.
// If that amount of money cannot be made up by any combination of the coins, return -1.
// Example 1:
//   Input: coins = [1, 2, 5], amount = 11
//   Output: 3
//   Explanation: 11 = 5 + 5 + 1
// Example 2:
//   Input: coins = [2], amount = 3
//   Output: -1
// Note: You may assume that you have an infinite number of each kind of coin.

func coinChange(coins []int, amount int) int {
	// let a(i, j) be the result for first i coins to get j amount.
	// then a(i, j) = min( a(i-1, j), a(i, j-coins[i])+1)
	// base case a(i, 0)=0, a(i, coins[i])=1
	if amount == 0 {
		return 0
	}
	a0 := make([]int, amount+1)
	a1 := make([]int, amount+1)
	for i := 0; i < len(coins); i++ {
		a1[0] = 0 // a(i, 0) = 0
		for j := 1; j <= amount; j++ {
			if j == coins[i] {
				a1[j] = 1 // a(i, coins[i]) = 1
				continue
			}
			a1[j] = a0[j] // a(i, j) = a(i-1, j)
			if j-coins[i] > 0 && a1[j-coins[i]] > 0 {
				if a0[j] == 0 || a1[j-coins[i]]+1 < a0[j] {
					a1[j] = a1[j-coins[i]] + 1
				}
			}
		}
		//fmt.Println(a1)
		a0, a1 = a1, a0
	}
	if a0[amount] == 0 {
		return -1
	}
	return a0[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2, 1, 5}, 11))
	fmt.Println(coinChange([]int{1, 6, 2}, 4))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{4}, 3))
}

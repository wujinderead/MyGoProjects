package main

import (
    "fmt"
)

// https://leetcode.com/problems/profitable-schemes/

// Alex and Lee play a game with piles of stones. There are an even number of piles arranged in a row, 
// and each pile has a positive integer number of stones piles[i]. The objective of the game is to end 
// with the most stones. The total number of stones is odd, so there are no ties. 
// Alex and Lee take turns, with Alex starting first. Each turn, a player takes the entire pile of stones 
// from either the beginning or the end of the row. This continues until there are no more piles left, 
// at which point the person with the most stones wins. 
// Assuming Alex and Lee play optimally, return True if and only if Alex wins the game. 
// Example 1: 
//   Input: [5,3,4,5]
//   Output: true
//   Explanation: 
//     Alex starts first, and can only take the first 5 or the last 5.
//     Say he takes the first 5, so that the row becomes [3, 4, 5].
//     If Lee takes 3, then the board is [4, 5], and Alex takes 5 to win with 10 points.
//     If Lee takes the last 5, then the board is [3, 4], and Alex takes 4 to win with 9 points.
//     This demonstrated that taking the first 5 was a winning move for Alex, so we return true.
// Note: 
//   2 <= piles.length <= 500 
//   piles.length is even. 
//   1 <= piles[i] <= 500 
//   sum(piles) is odd. 

// O(1) solution:
// alex can always win. Alex can always get the 1st-3rd-5th... elements,
// or he can always get the 2nd-4th-6th... elements. either sum is the max.

func stoneGame(piles []int) bool {
	// let dp(i, j) be the max stones Alex can get from piles[i...j]
	// if Alex takes piles[i], Lee can get dp(i+1, j), Alex get get total sum(piles[i...j]) - dp(i+1, j)
	// if Alex takes piles[j], Lee can get dp(i, j-1), Alex get get total sum(piles[i...j]) - dp(i, j-1)
	// we want the max. time, space O(n^2).
	n := len(piles)
	dp := make(map[[2]int]int)
	acc := make([]int, n+1)
	for i:=0; i<n-1; i++ {
		acc[i+1] = acc[i] + piles[i]
		dp[[2]int{i,i}] = piles[i]
		dp[[2]int{i, i+1}] = max(piles[i], piles[i+1])
	}
	dp[[2]int{n-1, n-1}] = piles[n-1]
	acc[n] = acc[n-1]+piles[n-1]

	// dp
	for diff:=2; diff<n; diff++ {
		for i:=0; i+diff<n; i++ {
			j := i+diff
			sum := acc[j+1]-acc[i]
			dp[[2]int{i, j}] = max(sum-dp[[2]int{i+1,j}], sum-dp[[2]int{i,j-1}])
		}
	}
	return dp[[2]int{0, n-1}] > acc[n]-acc[0]-dp[[2]int{0, n-1}] 
}

func max(a,b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(stoneGame([]int{5,3,4,5}))
	fmt.Println(stoneGame([]int{5,3,100,50}))
	fmt.Println(stoneGame([]int{5,3}))
	fmt.Println(stoneGame([]int{3,5}))
}
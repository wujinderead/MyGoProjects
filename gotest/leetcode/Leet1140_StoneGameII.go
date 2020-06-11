package main

import (
    "fmt"
)

// https://leetcode.com/problems/stone-game-ii/

// Alex and Lee continue their games with piles of stones. There are a number of piles arranged in a row, 
// and each pile has a positive integer number of stones piles[i]. The objective of the game is to end 
// with the most stones. Alex and Lee take turns, with Alex starting first. Initially, M = 1. 
// On each player's turn, that player can take all the stones in the first X remaining piles, 
// where 1 <= X <= 2M. Then, we set M = max(M, X). 
// The game continues until all the stones have been taken. 
// Assuming Alex and Lee play optimally, return the maximum number of stones Alex can get. 
// Example 1: 
//   Input: piles = [2,7,9,4,4]
//   Output: 10
//   Explanation:  If Alex takes one pile at the beginning, Lee takes two piles, then Alex takes 2 piles again. 
//     Alex can get 2 + 4 + 4 = 10 piles in total. If Alex takes two piles at the beginning, then Lee can take 
//     all three piles left. In this case, Alex get 2 + 7 = 9 piles in total. So we return 10 since it's larger. 
// Constraints: 
//   1 <= piles.length <= 100 
//   1 <= piles[i] <= 10 ^ 4 

func stoneGameII(piles []int) int {
    // let dp(i, j) be the most stones alex can get from piles[i:] with M=j, then, 
    // if Alex takes 1<=k<=2j piles, Lee can get dp(i+k, max(k,j)) stones from piles[i+k:],
    // so the total stones Alex can get is 
    // dp(i,j) = sum(piles[i: i+k]) + sum(piles[i+k:]) - dp(i+k, max(k,j)) = sum(piles[i:]) - dp(i+k, max(k,j)
    // we need find the best k that makes dp(i,j) maximal.
    n := len(piles)
    acc := make([]int, n+1)
    for i:=1; i<=n; i++ {
    	acc[i] = acc[i-1]+piles[i-1]
    }
 	dp := [100][65]int{}          // since we have at most 100 piles, M won't exceed 64
 	for i:=n-1; i>=0; i-- {
 		for j:=64; j>=1; j-- {
 			if 2*j>=n-i {  // we can take all, we take all
 				dp[i][j] = acc[n]-acc[i]
 				continue
 			}
 			for k:=1; k<=2*j; k++ {
 				dp[i][j] = max(dp[i][j], acc[n]-acc[i] - dp[i+k][min(64, max(k,j))])
 			}
 		}
 	}

 	return dp[0][1]
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	fmt.Println(stoneGameII([]int{2,7,9,4,4}))
	fmt.Println(stoneGameII([]int{2}))
	fmt.Println(stoneGameII([]int{2,3}))
	fmt.Println(stoneGameII([]int{2,3,5}))
	fmt.Println(stoneGameII([]int{2,3,4,4}))
	fmt.Println(stoneGameII([]int{2,3,4,2}))
	fmt.Println(stoneGameII([]int{287,4481,7269,4526,6341,7931,5551,7077,5888,6119,3083,6740,8401,2617,611,9077,720,5242,
		8815,4151,4227,7177,6110,1712,2311,2721,7427,6291,1622,9955,7071,9436,1424,5733,9271,9387,9078,8532,5193,3703,
		7503,1269,7549,3374,1053,7367,7357,4142,8586,3411,9958,8757,5394,3958,353,5556,6742,1596,1840,8856,4440,7078,
		5728,1147,5204,8119,6930,6016,9302,4599}))
}
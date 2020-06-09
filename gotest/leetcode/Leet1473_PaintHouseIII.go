package main

import (
    "fmt"
)

// https://leetcode.com/problems/paint-house-iii/

// There is a row of m houses in a small city, each house must be painted with one of the n colors 
// (labeled from 1 to n), some houses that has been painted last summer should not be painted again. 
// A neighborhood is a maximal group of continuous houses that are painted with the same color. 
// (For example: houses = [1,2,2,3,3,2,1,1] contains 5 neighborhoods [{1}, {2,2}, {3,3}, {2}, {1,1}]). 
// Given an array houses, an m * n matrix cost and an integer target where: 
//   houses[i]: is the color of the house i, 0 if the house is not painted yet. 
//   cost[i][j]: is the cost of paint the house i with the color j+1. 
// Return the minimum cost of painting all the remaining houses in such a way that there are exactly 
// target neighborhoods, if not possible return -1. 
// Example 1: 
//   Input: houses = [0,0,0,0,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
//   Output: 9
//   Explanation: Paint houses of this way [1,2,2,1,1]
//     This array contains target = 3 neighborhoods, [{1}, {2,2}, {1,1}].
//     Cost of paint all houses (1 + 1 + 1 + 1 + 5) = 9.
// Example 2: 
//   Input: houses = [0,2,1,2,0], cost = [[1,10],[10,1],[10,1],[1,10],[5,1]], m = 5, n = 2, target = 3
//   Output: 11
//   Explanation: Some houses are already painted, Paint the houses of this way [2,2,1,2,2]
//     This array contains target = 3 neighborhoods, [{2,2}, {1}, {2,2}]. 
//     Cost of paint the first and last house (10 + 1) = 11.
// Example 3: 
//   Input: houses = [0,0,0,0,0], cost = [[1,10],[10,1],[1,10],[10,1],[1,10]], m = 5, n = 2, target = 5
//   Output: 5
// Example 4: 
//   Input: houses = [3,1,2,3], cost = [[1,1,1],[1,1,1],[1,1,1],[1,1,1]], m = 4, n = 3, target = 3
//   Output: -1
//   Explanation: Houses are already painted with a total of 4 neighborhoods [{3},{1},{2},{3}] different of target = 3.
// Constraints:  
//   m == houses.length == cost.length 
//   n == cost[i].length 
//   1 <= m <= 100 
//   1 <= n <= 20 
//   1 <= target <= m 
//   0 <= houses[i] <= n 
//   1 <= cost[i][j] <= 10^4 

func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	// let dp[i][j][k] be the minimum cost where we have k neighborhoods in the first i houses 
	// and the i-th house is painted with the color j.
	// dp[i][j][k] = dp[i-1][j][k] + cost[i][j]
	//               min(dp[i-1][x!=j][k-1]) + cost[i][j]
	// if i-th color is determined as x, dp[i][x!=j][k] = 1e9
	// dp[i][j][k] = dp[i-1][j][k]
	//               min(dp[i-1][x!=j][k-1]),      no need to add cost
	dp := [101][21][101]int{}

	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			for k:=1; k<=i && k<=target; k++ {
				dp[i][j][k] = int(1e9)
				if houses[i-1]>0 && j != houses[i-1] {
					continue
				}
				// house[i-1]==0 || house[i-1]>0 && j==house[i-1]
				
				c := cost[i-1][j-1]
				if houses[i-1]>0 {   // already painted, not need to add cost
					c = 0
				}
				if i-1>=k {   // only i-1>=k meaningful
					dp[i][j][k] = min(dp[i][j][k], dp[i-1][j][k] + c)
				} 
				if i==1 || k-1>=1 {  // only k-1>=0 meaningful, except i==1 
					for x:=1; x<=n; x++ {
						if x != j {
							dp[i][j][k] = min(dp[i][j][k], dp[i-1][x][k-1] + c)
						}
					}
				}
			}
		}
	}
	
	allmin := int(1e9)
	for j:=1; j<=n; j++ {
		allmin = min(allmin, dp[m][j][target])
	}
	if allmin==int(1e9) {
		return -1
	}
	return allmin
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
    fmt.Println(minCost([]int{0,0,0,0,0}, [][]int{{1,10},{10,1},{10,1},{1,10},{5,1}}, 5, 2, 3))
    fmt.Println(minCost([]int{0,2,1,2,0}, [][]int{{1,10},{10,1},{10,1},{1,10},{5,1}}, 5, 2, 3))
	fmt.Println(minCost([]int{0,0,0,0,0}, [][]int{{1,10},{10,1},{1,10},{10,1},{1,10}}, 5, 2, 5))
	fmt.Println(minCost([]int{3,1,2,3}, [][]int{{1,1,1},{1,1,1},{1,1,1},{1,1,1}}, 4, 3, 3))
}
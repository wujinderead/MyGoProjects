package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-swaps-to-arrange-a-binary-grid/

// Given an n x n binary grid, in one step you can choose two adjacent rows of the grid and swap them.
// A grid is said to be valid if all the cells above the main diagonal are zeros.
// Return the minimum number of steps needed to make the grid valid, or -1 if the grid cannot be valid.
// The main diagonal of a grid is the diagonal that starts at cell (1, 1) and ends at cell (n, n).
// Example 1:
//      001      110      110      100
//      110  ->  001  ->  100  ->  110
//      100      100      001      001
//   Input: grid = [[0,0,1],[1,1,0],[1,0,0]]
//   Output: 3
// Example 2:
//   Input: grid = [[0,1,1,0],[0,1,1,0],[0,1,1,0],[0,1,1,0]]
//   Output: -1
//   Explanation: All rows are similar, swaps have no effect on the grid.
// Example 3:
//   Input: grid = [[1,0,0],[1,1,0],[1,1,1]]
//   Output: 0
// Constraints:
//   n == grid.length
//   n == grid[i].length
//   1 <= n <= 200
//   grid[i][j] is 0 or 1

func minSwaps(grid [][]int) int {
	n := len(grid)
	zeros := make([]int, n)
	swap := 0
	// how many consecutive 0's to the right of each row
	for i:=0; i<n; i++ {
		for j:=n-1; j>=0; j-- {
			if grid[i][j]!=0 {
				break
			}
			zeros[i] += 1
		} 
	}
	outer: for i:=0; i<n-1; i++ {
		for j:=i; j<n; j++ {
			if zeros[j] >= n-i-1 {      // for i-th row, we need at least n-i-1 0's
				for k:=j; k>i; k--{     // do swaps
					zeros[k] = zeros[k-1]
				}
				swap += j-i
				continue outer
			}
		}
		return -1     // can't find, return -1		
	}
	return swap
}

func main() {
	fmt.Println(minSwaps([][]int{{0,0,1},{1,1,0},{1,0,0}}), 3)
	fmt.Println(minSwaps([][]int{{0,1,1,0},{0,1,1,0},{0,1,1,0},{0,1,1,0}}), -1)
	fmt.Println(minSwaps([][]int{{1,0,0},{1,1,0},{1,1,1}}), 0)
	fmt.Println(minSwaps([][]int{{1}}), 0)
	fmt.Println(minSwaps([][]int{{0}}), 0)
	fmt.Println(minSwaps([][]int{{1,1},{1,0}}), 1)
	fmt.Println(minSwaps([][]int{{1,1},{0,1}}), -1)
	fmt.Println(minSwaps([][]int{{0,0},{1,1}}), 0)	
}
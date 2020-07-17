package main

import (
	"fmt"
)

// https://leetcode.com/problems/largest-1-bordered-square/

// Given a 2D grid of 0s and 1s, return the number of elements in the largest square subgrid 
// that has all 1s on its border, or 0 if such a subgrid doesn't exist in the grid.
// Example 1:
//   Input: grid = [[1,1,1],[1,0,1],[1,1,1]]
//   Output: 9
// Example 2:
//   Input: grid = [[1,1,0,0]]
//   Output: 1
// Constraints:
//   1 <= grid.length <= 100
//   1 <= grid[0].length <= 100
//   grid[i][j] is 0 or 1

func largest1BorderedSquare(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// how many consecutive 1s to the left and up of a cell
	left, up := [100][100]int{}, [100][100]int{} 
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if grid[i][j]==1 {
				if i>0 {
					up[i][j] = up[i-1][j]+1
				} else {
					up[i][j] = grid[i][j]
				}
				if j>0 {
					left[i][j] = left[i][j-1]+1
				} else {
					left[i][j] = grid[i][j]
				}
			}
		}
	}

	maxsquare := 0
	// dp 
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {     // let grid[i][j] be the bottom-right corner of a square
			m := min(left[i][j], up[i][j])
			for k:=m; k>0; k-- {
				// if the top-right corner has more than k 1s to its left, and,
				// if the bottom-left corner has more than k 1s to its up, 
				// we can find a 1-border square with size k  
				if left[i-(k-1)][j]>=k && up[i][j-(k-1)]>=k {
					maxsquare = max(maxsquare, k*k)
					break
				}
			}
		}
	} 
	return maxsquare
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
	fmt.Println(largest1BorderedSquare([][]int{{0,0,0}}))	
	fmt.Println(largest1BorderedSquare([][]int{{1,1,1},{1,0,1},{1,1,1}}))
	fmt.Println(largest1BorderedSquare([][]int{{1,1,0,0}}))
	fmt.Println(largest1BorderedSquare([][]int{
		{1,1,1,1,1},
		{1,1,1,0,1},
		{1,0,1,0,1},
		{1,1,1,0,1},
		{1,1,1,1,1},
	}))
	fmt.Println(largest1BorderedSquare([][]int{
		{1,1,1,1,1},
		{1,1,1,1,1},
		{1,1,1,0,1},
		{0,1,0,0,1},
		{1,1,1,1,1},
	}))
}
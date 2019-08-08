package main

import "fmt"

// https://leetcode.com/problems/unique-paths-ii/
// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time.
// The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// Now consider if some obstacles are added to the grids.
// How many unique paths would there be?
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	arr := make([]int, n)
	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		arr[i] = 1
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			arr[0] = 0
		}
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				arr[j] = arr[j] + arr[j-1]
			} else {
				arr[j] = 0
			}
		}
	}
	return arr[n-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 1, 0, 0},
		{0, 0, 0, 0},
		{1, 0, 0, 1},
		{0, 0, 0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{1},
		{0},
	}))
}

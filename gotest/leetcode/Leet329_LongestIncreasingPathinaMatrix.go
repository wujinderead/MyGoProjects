package main

import "fmt"

// https://leetcode.com/problems/longest-increasing-path-in-a-matrix/

// Given an integer matrix, find the length of the longest increasing path.
// From each cell, you can either move to four directions: left, right, up or down.
// You may NOT move diagonally or move outside of the boundary
// (i.e. wrap-around is not allowed).
// Example 1:
//   Input: nums =
//     [
//       [9,9,4],
//       [6,6,8],
//       [2,1,1]
//     ]
//   Output: 4
//   Explanation: The longest increasing path is [1, 2, 6, 9].
// Example 2:
//   Input: nums =
//     [
//       [3,4,5],
//       [3,2,6],
//       [2,2,1]
//     ]
//   Output: 4
//   Explanation: The longest increasing path is [3, 4, 5, 6]. Moving diagonally is not allowed.

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	longest := 1
	leng := make([][]int, len(matrix))
	for i := range leng {
		leng[i] = make([]int, len(matrix[0]))
	}
	for i := range matrix {
		for j := range matrix[0] {
			longest = max(longest, lip(matrix, leng, i, j))
		}
	}
	return longest
}

func lip(matrix, leng [][]int, i, j int) int {
	if leng[i][j] > 0 {
		return leng[i][j]
	}
	leng[i][j] = 1
	for _, v := range [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		ni, nj := i+v[0], j+v[1]
		if ni >= 0 && nj >= 0 && ni < len(matrix) && nj < len(matrix[0]) && matrix[i][j] < matrix[ni][nj] {
			leng[i][j] = max(leng[i][j], 1+lip(matrix, leng, ni, nj))
		}
	}
	return leng[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}))
	fmt.Println(longestIncreasingPath([][]int{
		{3, 4, 5},
		{3, 2, 6},
		{2, 2, 1},
	}))
}

package main

import (
	"fmt"
	"container/list"
)

// https://leetcode.com/problems/shortest-path-in-binary-matrix/

// In an N by N square grid, each cell is either empty (0) or blocked (1).
// A clear path from top-left to bottom-right has length k if and only if it is 
// composed of cells C_1, C_2, ..., C_k such that:
//   Adjacent cells C_i and C_{i+1} are connected 8-directionally (ie., they are
//     different and share an edge or corner)
//   C_1 is at location (0, 0) (ie. has value grid[0][0])
//   C_k is at location (N-1, N-1) (ie. has value grid[N-1][N-1])
//   If C_i is located at (r, c), then grid[r][c] is empty (ie. grid[r][c] == 0).
// Return the length of the shortest such clear path from top-left to bottom-right.
// If such a path does not exist, return -1.
// Example 1:
//   Input: [[0,1],[1,0]]
//   Output: 2
//      0 1
//      1 0
// Example 2:
//   Input: [[0,0,0],[1,1,0],[1,1,0]]
//   Output: 4
//      0 0 0
//      1 1 0
//      1 1 0
// Note:
//   1 <= grid.length == grid[0].length <= 100
//   grid[r][c] is 0 or 1

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0]==1 {
		return -1
	}
	if n==1 && grid[0][0]==0 {
		return 1
	}
	queue := list.New()
	queue.PushBack([2]int{0,0})
	grid[0][0] = 1
	for queue.Len()>0 {
		t := queue.Remove(queue.Front()).([2]int)
		i, j := t[0], t[1]
		for _, v := range [][2]int{{0,1}, {0,-1}, {1,0}, {-1,0}, {1,1}, {1,-1}, {-1,1}, {-1,-1}} {
			ni, nj := i+v[0], j+v[1]
			if ni>=0 && nj>=0 && ni<n && nj<n && grid[ni][nj]==0 {
				grid[ni][nj] = grid[i][j]+1
				if ni==n-1 && nj==n-1 {
					return grid[ni][nj]
				}
				queue.PushBack([2]int{ni, nj})
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(shortestPathBinaryMatrix([][]int{{0,1},{1,0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{1}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{0,0,0},{1,1,0},{1,1,0}}))
	fmt.Println(shortestPathBinaryMatrix([][]int{{1,0,0},{1,1,0},{1,1,0}}))
}
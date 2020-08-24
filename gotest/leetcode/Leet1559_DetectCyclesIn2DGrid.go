package main

import "fmt"

// https://leetcode.com/problems/detect-cycles-in-2d-grid/

// Given a 2D array of characters grid of size m x n, you need to find if there exists any cycle
// consisting of the same value in grid.
// A cycle is a path of length 4 or more in the grid that starts and ends at the same cell. From a
// given cell, you can move to one of the cells adjacent to it - in one of the four directions
// (up, down, left, or right), if it has the same value of the current cell.
// Also, you cannot move to the cell that you visited in your last move. For example, the cycle
// (1, 1) -> (1, 2) -> (1, 1) is invalid because from (1, 2) we visited (1, 1) which was the last
// visited cell.
// Return true if any cycle of the same value exists in grid, otherwise, return false.
// Example 1:
//   Input: grid = [['a','a','a','a'],['a','b','b','a'],['a','b','b','a'],['a','a','a','a']]
//   Output: true
//   Explanation: There are two valid cycles shown in different colors in the image below:
//         A-A-A-A
//         |     |
//         A B-B A
//         | | | |
//         A B-B A
//         |     |
//         A-A-A-A
// Example 2:
//   Input: grid = [['c','c','c','a'],['c','d','c','c'],['c','c','e','c'],['f','c','c','c']]
//   Output: true
//   Explanation: There is only one valid cycle highlighted in the image below:
//         C-C-C A
//         |   |
//         C D C-C
//         |     |
//         C C E C
//           |   |
//         F C-C-C
// Example 3:
//   Input: grid = [['a','b','b'],['b','z','b'],['b','b','a']]
//   Output: false
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m <= 500
//   1 <= n <= 500
//   grid consists only of lowercase English letters.

func containsCycle(grid [][]byte) bool {
	m, n := len(grid), len(grid[0])
    visited := make([][]bool, m)
    for i := range visited {
    	visited[i] = make([]bool, n)
	}

	// dfs
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			if !visited[i][j] {
				if visit(grid, visited, i, j, -1, -1) {
					return true
				}
			}
		}
	}
	return false
}

func visit(grid [][]byte, visited [][]bool, i, j, pi, pj int) bool {
	visited[i][j] = true
	m, n := len(grid), len(grid[0])
	for _, v := range [][2]int{{-1,0}, {1,0}, {0,1}, {0,-1}} {
		ni, nj := i+v[0], j+v[1]
		if ni>=0 && ni<m && nj>=0 && nj<n && grid[ni][nj]==grid[i][j] {  // only visit same char
			if !visited[ni][nj] {   // if unvisited, visit it
				if visit(grid, visited, ni, nj, i, j) {
					return true
				}
			} else if pi!=ni && pj!=nj {
				// if can visit a visited node, ans that node is not the from node,
				// then there is a circle.
				return true
			}
		}
	}
	return false
}

func main() {
	for _, v := range []struct{grid [][]byte; ans bool} {
		{[][]byte{{'a','a','a','a'},{'a','b','b','a'},{'a','b','b','a'},{'a','a','a','a'}}, true},
		{[][]byte{{'c','c','c','a'},{'c','d','c','c'},{'c','c','e','c'},{'f','c','c','c'}}, true},
		{[][]byte{{'a','b','b'},{'b','z','b'},{'b','b','a'}}, false},
		{[][]byte{{'d','b','b'},{'c','a','a'},{'b','a','c'},{'c','c','c'},{'d','d','a'}}, false},
		{[][]byte{{'c','a','d'},{'a','a','a'},{'a','a','d'},{'a','c','d'},{'a','b','c'}}, true},
	} {
		fmt.Println(containsCycle(v.grid), v.ans)
	}
}
package main

import "fmt"

// https://leetcode.com/problems/count-sub-islands/

// You are given two m x n binary matrices grid1 and grid2 containing only 0's (representing water)
// and 1's (representing land). An island is a group of 1's connected 4-directionally (horizontal
// or vertical). Any cells outside of the grid are considered water cells.
// An island in grid2 is considered a sub-island if there is an island in grid1 that contains all
// the cells that make up this island in grid2.
// Return the number of islands in grid2 that are considered sub-islands.
// Example 1:
//   Input: grid1 = [[1,1,1,0,0],[0,1,1,1,1],[0,0,0,0,0],[1,0,0,0,0],[1,1,0,1,1]],
//     grid2 = [[1,1,1,0,0],[0,0,1,1,1],[0,1,0,0,0],[1,0,1,1,0],[0,1,0,1,0]]
//   Output: 3
//   Explanation: In the picture above, the grid on the left is grid1 and the grid
//     on the right is grid2. The 1s colored red in grid2 are those considered to
//     be part of a sub-island. There are three sub-islands.
// Example 2:
//   Input: grid1 = [[1,0,1,0,1],[1,1,1,1,1],[0,0,0,0,0],[1,1,1,1,1],[1,0,1,0,1]],
//     grid2 = [[0,0,0,0,0],[1,1,1,1,1],[0,1,0,1,0],[0,1,0,1,0],[1,0,0,0,1]]
//   Output: 2
//   Explanation: In the picture above, the grid on the left is grid1 and the grid
//     on the right is grid2. The 1s colored red in grid2 are those considered to be
//     part of a sub-island. There are two sub-islands.
// Constraints:
//   m == grid1.length == grid2.length
//   n == grid1[i].length == grid2[i].length
//   1 <= m, n <= 500
//   grid1[i][j] and grid2[i][j] are either 0 or 1.

// find every islands of grid2, check if all land is in grid1
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	visited := make([][]bool, len(grid1))
	for i := range visited {
		visited[i] = make([]bool, len(grid1[0]))
	}
	count := 0
	for i := range grid2 {
		for j := range grid2[0] {
			if grid2[i][j] == 1 && !visited[i][j] {
				if visit(grid1, grid2, visited, i, j) {
					count++
				}
			}
		}
	}
	return count
}

// visit an island in grid2, return true if this island is sub-island in grid1
func visit(grid1, grid2 [][]int, visited [][]bool, i, j int) bool {
	visited[i][j] = true
	ret := grid1[i][j] == 1
	for _, v := range [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}} {
		ni, nj := i+v[0], j+v[1]
		if ni >= 0 && ni < len(grid1) && nj >= 0 && nj < len(grid1[0]) &&
			grid2[ni][nj] == 1 && !visited[ni][nj] {
			// order matters here, if use `ret = ret & visit(...)`, then if ret=false, visit(...) won't be called
			ret = visit(grid1, grid2, visited, ni, nj) && ret
		}
	}
	return ret
}

func main() {
	for _, v := range []struct {
		g1, g2 [][]int
		ans    int
	}{
		{[][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}},
			[][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}, 3},
		{[][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}},
			[][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}}, 2},
	} {
		fmt.Println(countSubIslands(v.g1, v.g2), v.ans)
	}
}

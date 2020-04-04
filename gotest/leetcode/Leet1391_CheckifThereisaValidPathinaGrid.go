package main

import "fmt"

// https://leetcode.com/problems/check-if-there-is-a-valid-path-in-a-grid

// Given a m x n grid. Each cell of the grid represents a street. The street of
// grid[i][j] can be:
//   1 which means a street connecting the left cell and the right cell.
//   2 which means a street connecting the upper cell and the lower cell.
//   3 which means a street connecting the left cell and the lower cell.
//   4 which means a street connecting the right cell and the lower cell.
//   5 which means a street connecting the left cell and the upper cell.
//   6 which means a street connecting the right cell and the upper cell.
// You will initially start at the street of the upper-left cell (0,0). A valid
// path in the grid is a path which starts from the upper left cell (0,0) and ends
// at the bottom-right cell (m - 1, n - 1). The path should only follow the streets.
// Notice that you are not allowed to change any street.
// Return true if there is a valid path in the grid or false otherwise.
// Example 1:
//   Input: grid = [[2,4,3],[6,5,2]]
//   Output: true
//   Explanation: As shown you can start at cell (0, 0) and visit all the cells of
//     the grid to reach (m - 1, n - 1).
// Example 2:
//   Input: grid = [[1,2,1],[1,2,1]]
//   Output: false
//   Explanation: As shown you the street at cell (0, 0) is not connected with any
//     street of any other cell and you will get stuck at cell (0, 0)
// Example 3:
//   Input: grid = [[1,1,2]]
//   Output: false
//   Explanation: You will get stuck at cell (0, 1) and you cannot reach cell (0, 2).
// Example 4:
//   Input: grid = [[1,1,1,1,1,1,3]]
//   Output: true
// Example 5:
//   Input: grid = [[2],[2],[2],[2],[2],[2],[6]]
//   Output: true
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 300
//   1 <= grid[i][j] <= 6

func hasValidPath(grid [][]int) bool {
	// dfs to traverse the graph
	canVisit := false
	m, n := len(grid), len(grid[0])
	visited := make([]bool, m*n)
	visited[0] = true
	dfs(0, 0, grid, visited, &canVisit)
	return canVisit
}

func dfs(i, j int, grid [][]int, visited []bool, canVisit *bool) {
	if i == len(grid)-1 && j == len(grid[0])-1 {
		*canVisit = true
		return
	}
	// 135 left, 146 right, 256 upper, 234 lower
	if !*canVisit && i+1 < len(grid) && !get2d(visited, i+1, j, len(grid[0])) &&
		(grid[i][j] == 2 || grid[i][j] == 3 || grid[i][j] == 4) &&
		(grid[i+1][j] == 2 || grid[i+1][j] == 5 || grid[i+1][j] == 6) {
		set2d(visited, i+1, j, len(grid[0]))
		dfs(i+1, j, grid, visited, canVisit)
	}
	if !*canVisit && i-1 >= 0 && !get2d(visited, i-1, j, len(grid[0])) &&
		(grid[i][j] == 2 || grid[i][j] == 5 || grid[i][j] == 6) &&
		(grid[i-1][j] == 2 || grid[i-1][j] == 3 || grid[i-1][j] == 4) {
		set2d(visited, i-1, j, len(grid[0]))
		dfs(i-1, j, grid, visited, canVisit)
	}
	if !*canVisit && j+1 < len(grid[0]) && !get2d(visited, i, j+1, len(grid[0])) &&
		(grid[i][j] == 1 || grid[i][j] == 4 || grid[i][j] == 6) &&
		(grid[i][j+1] == 1 || grid[i][j+1] == 3 || grid[i][j+1] == 5) {
		set2d(visited, i, j+1, len(grid[0]))
		dfs(i, j+1, grid, visited, canVisit)
	}
	if !*canVisit && j-1 >= 0 && !get2d(visited, i, j-1, len(grid[0])) &&
		(grid[i][j] == 1 || grid[i][j] == 3 || grid[i][j] == 5) &&
		(grid[i][j-1] == 1 || grid[i][j-1] == 4 || grid[i][j-1] == 6) {
		set2d(visited, i, j-1, len(grid[0]))
		dfs(i, j-1, grid, visited, canVisit)
	}
}

func set2d(arr []bool, i, j, col int) {
	arr[i*col+j] = true
}

func get2d(arr []bool, i, j, col int) bool {
	return arr[i*col+j]
}

func main() {
	fmt.Println(hasValidPath([][]int{{2, 4, 3}, {6, 5, 2}}))
	fmt.Println(hasValidPath([][]int{{1, 2, 1}, {1, 2, 1}}))
	fmt.Println(hasValidPath([][]int{{1, 1, 2}}))
	fmt.Println(hasValidPath([][]int{{1, 1, 1, 1, 1, 1, 3}}))
	fmt.Println(hasValidPath([][]int{{2}, {2}, {2}, {2}, {2}, {2}, {6}}))
}

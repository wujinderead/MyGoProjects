package main

import "fmt"

// https://leetcode.com/problems/path-with-maximum-gold/

// In a gold mine grid of size m * n, each cell in this mine has an integer
// representing the amount of gold in that cell, 0 if it is empty.
// Return the maximum amount of gold you can collect under the conditions:
// Every time you are located in a cell you will collect all the gold in that cell.
// From your position you can walk one step to the left, right, up or down.
// You can't visit the same cell more than once.
// Never visit a cell with 0 gold.
// You can start and stop collecting gold from any position in the grid that has some gold.
// Example 1:
//   Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
//   Output: 24
//   Explanation:
//     [[0,6,0],
//      [5,8,7],
//      [0,9,0]]
//     Path to get the maximum gold, 9 -> 8 -> 7.
// Example 2:
//   Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
//   Output: 28
//   Explanation:
//     [[1,0,7],
//      [2,0,6],
//      [3,4,5],
//      [0,3,0],
//      [9,0,20]]
//     Path to get the maximum gold, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.
// Constraints:
//   1 <= grid.length, grid[i].length <= 15
//   0 <= grid[i][j] <= 100
//   There are at most 25 cells containing gold.

func getMaximumGold(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	max, cur := new(int), new(int)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// only start from edge cell
			if grid[i][j] > 0 {
				visit(grid, visited, i, j, max, cur)
			}
		}
	}
	return *max
}

func visit(grid [][]int, visited [][]bool, i, j int, max, cur *int) {
	visited[i][j] = true
	*cur += grid[i][j]
	if *cur > *max {
		*max = *cur
	}
	for _, neighbor := range neighbors(grid, i, j) {
		if !visited[neighbor[0]][neighbor[1]] {
			visit(grid, visited, neighbor[0], neighbor[1], max, cur)
		}
	}
	*cur -= grid[i][j]
	visited[i][j] = false
}

func neighbors(grid [][]int, i, j int) [][2]int {
	dir := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	neighbor := make([][2]int, 0, 4)
	for d := range dir {
		ni, nj := i+dir[d][0], j+dir[d][1]
		if ni >= 0 && nj >= 0 && ni < len(grid) && nj < len(grid[0]) && grid[ni][nj] > 0 {
			neighbor = append(neighbor, [2]int{ni, nj})
		}
	}
	return neighbor
}

func main() {
	fmt.Println(getMaximumGold([][]int{
		{0, 6, 0},
		{5, 8, 7},
		{0, 9, 0},
	}))
	fmt.Println(getMaximumGold([][]int{
		{1, 0, 7},
		{2, 0, 6},
		{3, 4, 5},
		{0, 3, 0},
		{9, 0, 20},
	}))
	fmt.Println(getMaximumGold([][]int{
		{23, 21, 38, 12, 18, 36, 0, 7, 30, 29, 20, 3, 28},
		{23, 3, 19, 2, 1, 11, 4, 8, 9, 24, 6, 5, 35},
	}))
}

package main

import "fmt"

// https://leetcode.com/problems/unique-paths-iii/

// On a 2-dimensional grid, there are 4 types of squares:
//   1 represents the starting square. There is exactly one starting square.
//   2 represents the ending square. There is exactly one ending square.
//   0 represents empty squares we can walk over.
//   -1 represents obstacles that we cannot walk over.
// Return the number of 4-directional walks from the starting square to the ending square,
// that walk over every non-obstacle square exactly once.
// Example 1:
//   Input: [[1,0,0,0],[0,0,0,0],[0,0,2,-1]]
//   Output: 2
//   Explanation: We have the following two paths:
//     1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2)
//     2. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2)
// Example 2:
//   Input: [[1,0,0,0],[0,0,0,0],[0,0,0,2]]
//   Output: 4
//   Explanation: We have the following four paths:
//     1. (0,0),(0,1),(0,2),(0,3),(1,3),(1,2),(1,1),(1,0),(2,0),(2,1),(2,2),(2,3)
//     2. (0,0),(0,1),(1,1),(1,0),(2,0),(2,1),(2,2),(1,2),(0,2),(0,3),(1,3),(2,3)
//     3. (0,0),(1,0),(2,0),(2,1),(2,2),(1,2),(1,1),(0,1),(0,2),(0,3),(1,3),(2,3)
//     4. (0,0),(1,0),(2,0),(2,1),(1,1),(0,1),(0,2),(0,3),(1,3),(1,2),(2,2),(2,3)
// Example 3:
//   Input: [[0,1],[2,0]]
//   Output: 0
//   Explanation:
//     There is no path that walks over every empty square exactly once.
//     Note that the starting and ending square can be anywhere in the grid.
// Note:
//   1 <= grid.length * grid[0].length <= 20

func uniquePathsIII(grid [][]int) int {
	// use backtracking to get all possible solution
	m, n := len(grid), len(grid[0])
	nob := 0
	si, sj := -1, -1 // must exist
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				nob++
			}
			if grid[i][j] == 1 {
				si, sj = i, j
			}
		}
	}
	pathcount := 0
	grid[si][sj] = -1
	toNextPos(grid, si, sj, 0, nob, &pathcount)
	grid[si][sj] = 1
	return pathcount
}

func toNextPos(grid [][]int, i, j, curnob, nob int, pathcount *int) {
	for _, v := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		ni, nj := i+v[0], j+v[1]
		if ni >= 0 && nj >= 0 && ni < len(grid) && nj < len(grid[0]) && grid[ni][nj] != -1 {
			if grid[ni][nj] == 0 {
				curnob++
				grid[ni][nj] = -1
				toNextPos(grid, ni, nj, curnob, nob, pathcount)
				grid[ni][nj] = 0
				curnob--
			}
			if grid[ni][nj] == 2 {
				if curnob == nob { // find a valid path that walk through all non-obstacles
					*pathcount++
				}
			}
		}
	}
}

func main() {
	g := [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}}
	fmt.Println(uniquePathsIII(g))
	fmt.Println(g)
	g = [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}}
	fmt.Println(uniquePathsIII(g))
	fmt.Println(g)
	g = [][]int{{0, 1}, {2, 0}}
	fmt.Println(uniquePathsIII(g))
	fmt.Println(g)

	fmt.Println(uniquePathsIII([][]int{{1}, {0}, {2}}))
	fmt.Println(uniquePathsIII([][]int{{0}, {1}, {2}}))
}

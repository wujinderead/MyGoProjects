package main

import "fmt"

// https://leetcode.com/problems/stamping-the-grid/

// You are given an m x n binary matrix grid where each cell is either 0 (empty) or 1 (occupied).
// You are then given stamps of size stampHeight x stampWidth. We want to fit the stamps such that
// they follow the given restrictions and requirements:
//   Cover all the empty cells.
//   Do not cover any of the occupied cells.
//   We can put as many stamps as we want.
//   Stamps can overlap with each other.
//   Stamps are not allowed to be rotated.
//   Stamps must stay completely inside the grid.
// Return true if it is possible to fit the stamps while following the given restrictions
// and requirements. Otherwise, return false.
// Example 1:
//   Input: grid = [[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0],[1,0,0,0]], stampHeight = 4, stampWidth = 3
//   Output: true
//   Explanation: We have two overlapping stamps (labeled 1 and 2 in the image)
//     that are able to cover all the empty cells.
//    X  1   1   1
//    X  12  12  12
//    X  12  12  12
//    X  12  12  12
//    X   2   2   2
// Example 2:
//   Input: grid = [[1,0,0,0],[0,1,0,0],[0,0,1,0],[0,0,0,1]], stampHeight = 2, stampWidth = 2
//   Output: false
//     Explanation: There is no way to fit the stamps onto all the empty cells
//     without the stamps going outside the grid.
// Constraints:
//   m == grid.length
//   n == grid[r].length
//   1 <= m, n <= 10⁵
//   1 <= m * n <= 2 * 10⁵
//   grid[r][c] is either 0 or 1.
//   1 <= stampHeight, stampWidth <= 10⁵

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	// pre check
	if stampHeight == 1 && stampWidth == 1 {
		return true
	}
	if stampHeight > len(grid) || stampWidth > len(grid[0]) {
		return false
	}

	// get accumulated occupied cells number in right and bottom of (i, j)
	count := make([][]int, len(grid)+1)
	for i := range count {
		count[i] = make([]int, len(grid[0])+1)
	}
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			count[i][j] = grid[i][j] + count[i+1][j] + count[i][j-1] - count[i+1][j+1]
		}
	}

	// check each cell
	for i := len(grid) - stampHeight; i >= 0; i-- {
		for j := len(grid[0]) - stampWidth; j >= 0; j-- {
			if grid[i][j] == 1 {
				continue
			}
			// how many occupied in this cell
			c := count[i][j] - count[i+stampHeight][j] - count[i][j+stampWidth] + count[i+stampHeight][j+stampWidth]

		}
	}

	// check the answer
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	for _, v := range []struct {
		grid        [][]int
		stampHeight int
		stampWidth  int
		ans         bool
	}{
		{[][]int{{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 4, 3, true},
		{[][]int{{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}}, 2, 2, false},
	} {
		fmt.Println(possibleToStamp(v.grid, v.stampHeight, v.stampWidth), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/grid-game/

// You are given a 0-indexed 2D array grid of size 2 x n, where grid[r][c] represents the number of
// points at position (r, c) on the matrix. Two robots are playing a game on this matrix.
// Both robots initially start at (0, 0) and want to reach (1, n-1). Each robot may only move to the
// right ((r, c) to (r, c + 1)) or down ((r, c) to (r + 1, c)).
// At the start of the game, the first robot moves from (0, 0) to (1, n-1), collecting all the points
// from the cells on its path. For all cells (r, c) traversed on the path, grid[r][c] is set to 0. Then,
// the second robot moves from (0, 0) to (1, n-1), collecting the points on its path. Note that their
// paths may intersect with one another.
// The first robot wants to minimize the number of points collected by the second robot. In contrast,
// the second robot wants to maximize the number of points it collects. If both robots play optimally,
// return the number of points collected by the second robot.
// Example 1:
//   Input: grid = [[2,5,4],[1,5,1]]
//   Output: 4
//   Explanation: The optimal path taken by the first robot is shown in red, and the optimal path taken by the second robot is shown in blue.
//     The cells visited by the first robot are set to 0.
//     The second robot will collect 0 + 0 + 4 + 0 = 4 points.
// Example 2:
//   Input: grid = [[3,3,1],[8,5,2]]
//   Output: 4
//   Explanation: The optimal path taken by the first robot is shown in red,
//     and the optimal path taken by the second robot is shown in blue.
//     The cells visited by the first robot are set to 0.
//     The second robot will collect 0 + 3 + 1 + 0 = 4 points.
// Example 3:
//   Input: grid = [[1,3,1,15],[1,3,3,1]]
//   Output: 7
//   Explanation: The optimal path taken by the first robot is shown in red,
//     and the optimal path taken by the second robot is shown in blue.
//     The cells visited by the first robot are set to 0.
//     The second robot will collect 0 + 1 + 3 + 3 + 0 = 7 points.
// Constraints:
//   grid.length == 2
//   n == grid[r].length
//   1 <= n <= 5 * 104
//   1 <= grid[r][c] <= 105

// original board:
//   xxxxxxx
//   xxxxxxx
// after first:
//   0000aaa
//   bbb0000
// so the second's best choice is to collect all a's or all b's
// so the first's object is to make max(aaa, bbb) minimal
func gridGame(grid [][]int) int64 {
	if len(grid[0]) == 1 {
		return 0
	}
	for i := len(grid[0]) - 2; i >= 0; i-- {
		grid[0][i] = grid[0][i] + grid[0][i+1]
	}
	ans := grid[0][1]
	sum := 0
	for i := 0; i+2 < len(grid[0]); i++ {
		sum += grid[1][i]
		ans = min(ans, max(sum, grid[0][i+2]))
	}
	ans = min(ans, sum+grid[1][len(grid[0])-2])
	return int64(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		g   [][]int
		ans int64
	}{
		{[][]int{{2, 5, 4}, {1, 5, 1}}, 4},
		{[][]int{{3, 3, 1}, {8, 5, 2}}, 4},
		{[][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}, 7},
		{[][]int{{1}, {2}}, 0},
		{[][]int{{1, 2}, {2, 1}}, 2},
		{[][]int{{1, 1}, {2, 2}}, 1},
	} {
		fmt.Println(gridGame(v.g), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/difference-between-ones-and-zeros-in-row-and-column/

// You are given a 0-indexed m x n binary matrix grid.
// A 0-indexed m x n difference matrix diff is created with the following procedure:
//   Let the number of ones in the iᵗʰ row be onesRowi.
//   Let the number of ones in the jᵗʰ column be onesColj.
//   Let the number of zeros in the iᵗʰ row be zerosRowi.
//   Let the number of zeros in the jᵗʰ column be zerosColj.
//   diff[i][j] = onesRowi + onesColj - zerosRowi - zerosColj
// Return the difference matrix diff.
// Example 1:
//   Input: grid = [[0,1,1],[1,0,1],[0,0,1]]
//   Output: [[0,0,4],[0,0,4],[-2,-2,2]]
//   Explanation:
//     - diff[0][0] = onesRow0 + onesCol0 - zerosRow0 - zerosCol0 = 2 + 1 - 1 - 2 = 0
//     - diff[0][1] = onesRow0 + onesCol1 - zerosRow0 - zerosCol1 = 2 + 1 - 1 - 2 = 0
//     - diff[0][2] = onesRow0 + onesCol2 - zerosRow0 - zerosCol2 = 2 + 3 - 1 - 0 = 4
//     - diff[1][0] = onesRow1 + onesCol0 - zerosRow1 - zerosCol0 = 2 + 1 - 1 - 2 = 0
//     - diff[1][1] = onesRow1 + onesCol1 - zerosRow1 - zerosCol1 = 2 + 1 - 1 - 2 = 0
//     - diff[1][2] = onesRow1 + onesCol2 - zerosRow1 - zerosCol2 = 2 + 3 - 1 - 0 = 4
//     - diff[2][0] = onesRow2 + onesCol0 - zerosRow2 - zerosCol0 = 1 + 1 - 2 - 2 = -2
//     - diff[2][1] = onesRow2 + onesCol1 - zerosRow2 - zerosCol1 = 1 + 1 - 2 - 2 = -2
//     - diff[2][2] = onesRow2 + onesCol2 - zerosRow2 - zerosCol2 = 1 + 3 - 2 - 0 = 2
// Example 2:
//   Input: grid = [[1,1,1],[1,1,1]]
//   Output: [[5,5,5],[5,5,5]]
//   Explanation:
//     - diff[0][0] = onesRow0 + onesCol0 - zerosRow0 - zerosCol0 = 3 + 2 - 0 - 0 = 5
//     - diff[0][1] = onesRow0 + onesCol1 - zerosRow0 - zerosCol1 = 3 + 2 - 0 - 0 = 5
//     - diff[0][2] = onesRow0 + onesCol2 - zerosRow0 - zerosCol2 = 3 + 2 - 0 - 0 = 5
//     - diff[1][0] = onesRow1 + onesCol0 - zerosRow1 - zerosCol0 = 3 + 2 - 0 - 0 = 5
//     - diff[1][1] = onesRow1 + onesCol1 - zerosRow1 - zerosCol1 = 3 + 2 - 0 - 0 = 5
//     - diff[1][2] = onesRow1 + onesCol2 - zerosRow1 - zerosCol2 = 3 + 2 - 0 - 0 = 5
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 10⁵
//   1 <= m * n <= 10⁵
//   grid[i][j] is either 0 or 1.

func onesMinusZeros(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	rowsOne := make([]int, m)
	colsOne := make([]int, n)
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 1 {
				rowsOne[i] += 1
				colsOne[j] += 1
			}
		}
	}
	for i := range grid {
		for j := range grid[0] {
			ans[i][j] = 2*(rowsOne[i]+colsOne[j]) - m - n
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		grid, ans [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 1, 1}}, [][]int{{5, 5, 5}, {5, 5, 5}}},
		{[][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}, [][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}}},
	} {
		fmt.Println(onesMinusZeros(v.grid), v.ans)
	}
}

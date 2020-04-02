package main

import (
	"container/list"
	"fmt"
)

// Given a m x n grid. Each cell of the grid has a sign pointing to the next cell
// you should visit if you are currently in this cell. The sign of grid[i][j] can be:
//   1 which means go to the cell to the right. (i.e go from grid[i][j] to grid[i][j+1])
//   2 which means go to the cell to the left. (i.e go from grid[i][j] to grid[i][j-1])
//   3 which means go to the lower cell. (i.e go from grid[i][j] to grid[i+1][j])
//   4 which means go to the upper cell. (i.e go from grid[i][j] to grid[i-1][j])
// Notice that there could be some invalid signs on the cells of the grid which
// points outside the grid.
// You will initially start at the upper left cell (0,0). A valid path in the grid
// is a path which starts from the upper left cell (0,0) and ends at the bottom-right
// cell (m - 1, n - 1) following the signs on the grid. The valid path doesn't
// have to be the shortest.
// You can modify the sign on a cell with cost = 1. You can modify the sign on a
// cell one time only.
// Return the minimum cost to make the grid have at least one valid path.
// Example 1:
//   Input: grid = [[1,1,1,1],[2,2,2,2],[1,1,1,1],[2,2,2,2]]
//   Output: 3
//   Explanation: You will start at point (0, 0).
//        → → → →
//        ← ← ← ←
//        → → → →
//        ← ← ← ←
//     The path to (3, 3) is as follows. (0, 0) --> (0, 1) --> (0, 2) --> (0, 3)
//     change the arrow to down with cost = 1 --> (1, 3) --> (1, 2) --> (1, 1) --> (1, 0)
//     change the arrow to down with cost = 1 --> (2, 0) --> (2, 1) --> (2, 2) --> (2, 3)
//     change the arrow to down with cost = 1 --> (3, 3)
//     The total cost = 3.
// Example 2:
//   Input: grid = [[1,1,3],[3,2,2],[1,1,4]]
//   Output: 0
//        → → ↓
//        ↓ ← ←
//        → → ↑
//   Explanation: You can follow the path from (0, 0) to (2, 2).
// Example 3:
//   Input: grid = [[1,2],[4,3]]
//   Output: 1
// Example 4:
//   Input: grid = [[2,2,2],[2,2,2]]
//   Output: 3
// Example 5:
//   Input: grid = [[4]]
//   Output: 0
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 100

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 0
	}
	c := make([]bool, m*n) // if we visited c[i][j]
	// 1→ 2← 3↓ 4↑
	// from bottom-right position, check the minimal cost of adjacent position
	queueCanReach := list.New()
	queueCanExpend := list.New()
	queueCanReach.PushBack([3]int{m - 1, n - 1, 0}) // (row, col, cost) pair
	set2D(c, m-1, n-1, n)                           // c[m-1][n-1]=true
	flag := true                                    // which queue to pop
	for {                                           // since we can always reach target, this loop always return
		if flag && queueCanReach.Len() > 0 {
			pair := queueCanReach.Remove(queueCanReach.Front()).([3]int)
			i, j, cost := pair[0], pair[1], pair[2]
			queueCanExpend.PushBack(pair)
			if i+1 < m && !get2D(c, i+1, j, n) && grid[i+1][j] == 4 { // ↑
				set2D(c, i+1, j, n)
				queueCanReach.PushBack([3]int{i + 1, j, cost})
			}
			if j+1 < n && !get2D(c, i, j+1, n) && grid[i][j+1] == 2 { // ←
				set2D(c, i, j+1, n)
				queueCanReach.PushBack([3]int{i, j + 1, cost})
			}
			if i-1 >= 0 && !get2D(c, i-1, j, n) && grid[i-1][j] == 3 { // ↓
				if i-1 == 0 && j == 0 {
					return cost
				}
				set2D(c, i-1, j, n)
				queueCanReach.PushBack([3]int{i - 1, j, cost})
			}
			if j-1 >= 0 && !get2D(c, i, j-1, n) && grid[i][j-1] == 1 { // →
				if i == 0 && j-1 == 0 {
					return cost
				}
				set2D(c, i, j-1, n)
				queueCanReach.PushBack([3]int{i, j - 1, cost})
			}
		} else {
			flag = false
		}
		// else pop from queueCanExpend
		if !flag && queueCanExpend.Len() > 0 {
			pair := queueCanExpend.Remove(queueCanExpend.Front()).([3]int)
			i, j, cost := pair[0], pair[1], pair[2]
			if i+1 < m && !get2D(c, i+1, j, n) {
				set2D(c, i+1, j, n)
				queueCanReach.PushBack([3]int{i + 1, j, cost + 1})
			}
			if j+1 < n && !get2D(c, i, j+1, n) {
				set2D(c, i, j+1, n)
				queueCanReach.PushBack([3]int{i, j + 1, cost + 1})
			}
			if i-1 >= 0 && !get2D(c, i-1, j, n) {
				if i-1 == 0 && j == 0 {
					return cost + 1
				}
				set2D(c, i-1, j, n)
				queueCanReach.PushBack([3]int{i - 1, j, cost + 1})
			}
			if j-1 >= 0 && !get2D(c, i, j-1, n) {
				if i == 0 && j-1 == 0 {
					return cost + 1
				}
				set2D(c, i, j-1, n)
				queueCanReach.PushBack([3]int{i, j - 1, cost + 1})
			}
		} else {
			flag = true
		}
	}
}

func set2D(arr []bool, i, j, col int) {
	arr[i*col+j] = true
}

func get2D(arr []bool, i, j, col int) bool {
	return arr[i*col+j]
}

func main() {
	// 1→ 2← 3↓ 4↑
	fmt.Println(minCost([][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 1, 1, 1}, {2, 2, 2, 2}}))
	fmt.Println(minCost([][]int{{1, 1, 1, 1}, {2, 2, 2, 2}, {1, 2, 1, 3}, {2, 1, 1, 4}}))
	fmt.Println(minCost([][]int{{1, 1, 3}, {3, 2, 2}, {1, 1, 4}}))
	fmt.Println(minCost([][]int{{1, 2}, {4, 3}}))
	fmt.Println(minCost([][]int{{2, 2, 2}, {2, 2, 2}}))
	fmt.Println(minCost([][]int{{4}}))
}

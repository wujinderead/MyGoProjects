package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/map-of-highest-peak/

// You are given an integer matrix isWater of size m x n that represents a map of land
// and water cells.
//   If isWater[i][j] == 0, cell (i, j) is a land cell.
//   If isWater[i][j] == 1, cell (i, j) is a water cell.
// You must assign each cell a height in a way that follows these rules:
//   The height of each cell must be non-negative.
//   If the cell is a water cell, its height must be 0.
//   Any two adjacent cells must have an absolute height difference of at most 1.
// A cell is adjacent to another cell if the former is directly north, east, south,
// or west of the latter (i.e., their sides are touching).
// Find an assignment of heights such that the maximum height in the matrix is maximized.
// Return an integer matrix height of size m x n where height[i][j] is cell (i, j)'s height.
// If there are multiple solutions, return any of them.
// Example 1:
//   Input: isWater = [[0,1],
//                     [0,0]]
//   Output: [[1,0],
//            [2,1]]
//   Explanation: The image shows the assigned heights of each cell.
//     The blue cell is the water cell, and the green cells are the land cells.
// Example 2:
//   Input: isWater = [[0,0,1],
//                     [1,0,0],
//                     [0,0,0]]
//   Output: [[1,1,0],
//            [0,1,1],
//            [1,2,2]]
//   Explanation: A height of 2 is the maximum possible height of any assignment.
//     Any height assignment that has a maximum height of 2 while still meeting the
//     rules will also be accepted.
// Constraints:
//   m == isWater.length
//   n == isWater[i].length
//   1 <= m, n <= 1000
//   isWater[i][j] is 0 or 1.
//   There is at least one water cell.

func highestPeak(isWater [][]int) [][]int {
	ans := make([][]int, len(isWater))
	for i := range ans {
		ans[i] = make([]int, len(isWater[0]))
	}

	// push water to queue
	queue := list.New()
	for i := range isWater {
		for j := range isWater[0] {
			if isWater[i][j] == 1 {
				queue.PushBack([2]int{i, j})
			}
		}
	}

	// bfs set land
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).([2]int)
		curi, curj := cur[0], cur[1]
		for _, v := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			newi, newj := curi+v[0], curj+v[1]
			if newi >= 0 && newi < len(isWater) && newj >= 0 && newj < len(isWater[0]) &&
				ans[newi][newj] == 0 && isWater[newi][newj] == 0 { // unset and is land
				ans[newi][newj] = ans[curi][curj] + 1
				queue.PushBack([2]int{newi, newj})
			}
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		wa, ans [][]int
	}{
		{[][]int{{0, 1}, {0, 0}}, [][]int{{1, 0}, {2, 1}}},
		{[][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}, [][]int{{1, 1, 0}, {0, 1, 1}, {1, 2, 2}}},
	} {
		fmt.Println(highestPeak(v.wa), v.ans)
	}
}

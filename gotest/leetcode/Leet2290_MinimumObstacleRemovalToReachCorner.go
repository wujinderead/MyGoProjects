package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/minimum-obstacle-removal-to-reach-corner/

// You are given a 0-indexed 2D integer array grid of size m x n. Each cell has one of two values:
//   0 represents an empty cell,
//   1 represents an obstacle that may be removed.
// You can move up, down, left, or right from and to an empty cell.
// Return the minimum number of obstacles to remove so you can move from the upper left corner (0, 0)
// to the lower right corner (m - 1, n - 1).
// Example 1:
//   Input: grid = [[0,1,1],[1,1,0],[1,1,0]]
//   Output: 2
//   Explanation: We can remove the obstacles at (0, 1) and (0, 2) to create a path from (0, 0) to (2, 2).
//     It can be shown that we need to remove at least 2 obstacles, so we return 2.
//     Note that there may be other ways to remove 2 obstacles to create a path.
// Example 2:
//   Input: grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]]
//   Output: 0
//   Explanation: We can move from (0, 0) to (2, 4) without removing any obstacles, so we return 0.
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 10⁵
//   2 <= m * n <= 10⁵
//   grid[i][j] is either 0 or 1.
//   grid[0][0] == grid[m - 1][n - 1] == 0

// dijkstra, go through empty node need 0 step, and obstacle need 1 step: O(mnlog(mn))
// can use dfs+bfs to achieve O(mn): dfs to traverse empty node and bfs to find shortest path
func minimumObstacles(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	h := pairs{[3]int{0, 0, 0}}
	visited := make(map[[2]int]struct{})
	visited[[2]int{0, 0}] = struct{}{}
	for h.Len() > 0 {
		tmp := heap.Pop(&h).([3]int)
		i, j, step := tmp[0], tmp[1], tmp[2]
		for _, v := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			ni, nj := i+v[0], j+v[1]
			if _, ok := visited[[2]int{ni, nj}]; ni >= 0 && ni < m && nj >= 0 && nj < n && !ok {
				if ni == m-1 && nj == n-1 {
					return step
				}
				heap.Push(&h, [3]int{ni, nj, step + grid[i][j]})
				visited[[2]int{ni, nj}] = struct{}{}
			}
		}
	}
	return 0
}

type pairs [][3]int

func (p pairs) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p pairs) Len() int            { return len(p) }
func (p pairs) Less(i, j int) bool  { return p[i][2] < p[j][2] }
func (p *pairs) Push(x interface{}) { *p = append(*p, x.([3]int)) }
func (p *pairs) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func main() {
	for _, v := range []struct {
		grid [][]int
		ans  int
	}{
		{[][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}, 2},
		{[][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}, 0},
		{[][]int{{0, 1, 1, 1, 0}}, 3},
	} {
		fmt.Println(minimumObstacles(v.grid), v.ans)
	}
}

package main

import (
	"fmt"
	"container/list"
)

// https://leetcode.com/problems/shortest-path-in-a-grid-with-obstacles-elimination/

// Given a m * n grid, where each cell is either 0 (empty) or 1 (obstacle). In one step, 
// you can move up, down, left or right from and to an empty cell. Return the minimum number 
// of steps to walk from the upper left corner (0, 0) to the lower right corner (m-1, n-1) given 
// that you can eliminate at most k obstacles. If it is not possible to find such walk return -1.
// Example 1:
//   Input: 
//     grid = 
//     [[0,0,0],
//      [1,1,0],
//      [0,0,0],
//      [0,1,1],
//      [0,0,0]], 
//     k = 1
//   Output: 6
//   Explanation: 
//     The shortest path without eliminating any obstacle is 10. 
//     The shortest path with one obstacle elimination at position (3,2) is 6. 
//     Such path is (0,0) -> (0,1) -> (0,2) -> (1,2) -> (2,2) -> (3,2) -> (4,2).
// Example 2:
//   Input: 
//     grid = 
//     [[0,1,1],
//      [1,1,1],
//      [1,0,0]], 
//     k = 1
//   Output: -1
//   Explanation: 
//     We need to eliminate at least two obstacles to find such a walk.
// Constraints:
//   grid.length == m
//   grid[0].length == n
//   1 <= m, n <= 40
//   1 <= k <= m*n
//   grid[i][j] == 0 or 1
//   grid[0][0] == grid[m-1][n-1] == 0

func shortestPath(grid [][]int, K int) int {
	m, n := len(grid), len(grid[0])
	if m==1 && n==1 {
		return 0
	}
	visited := make(map[int]int)
	visited[pack(0,0,0)] = 0          // (row, col, k) state we have seen
	queue := list.New()
	queue.PushBack(pack(0,0,0))
	for queue.Len()>0 {
		key := queue.Remove(queue.Front()).(int)
		i, j, k := unpack(key)
		step := visited[key]
		for _, v := range [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}} {
			ni, nj := i+v[0], j+v[1]
			if !(ni>=0 && ni<m && nj>=0 && nj<n) {  // invalid ni, nj
				continue
			}
			if grid[ni][nj]==0 {     // visit grid
				if ni==m-1 && nj==n-1 {
					return step+1
				}
				key = pack(ni, nj, k)
				if _, ok := visited[key]; !ok {
					visited[key] = step+1
					queue.PushBack(key)
				}
			}
			if grid[ni][nj]==1 && k+1<=K {     // visit obstacle
				key = pack(ni, nj, k+1)        // need increment k when visit obstacle
				if _, ok := visited[key]; !ok {
					visited[key] = step+1
					queue.PushBack(key)
				}
			}
		}
	}
	return -1
}

func pack(i, j, k int) int {
	return i+(j<<8)+(k<<16)
}

func unpack(a int) (int, int, int) {
	return a&0xff, (a&0xff00)>>8, a>>16 
}

func main() {
	fmt.Println(shortestPath([][]int{
		{0,0,0},
		{1,1,0},
		{0,0,0},
		{0,1,1},
		{0,0,0},
	}, 1))
	fmt.Println(shortestPath([][]int{
		{0,0,0},
		{1,1,0},
		{0,0,0},
		{1,1,1},
		{0,0,0},
	}, 1))
	fmt.Println(shortestPath([][]int{
		{0,1,1},
		{1,1,1},
		{1,0,0},
	}, 1))

	fmt.Println(shortestPath([][]int{
		{0,0,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,0},
		{0,1,0,0,0,0,0,0,0,0},
		{0,1,0,1,1,1,1,1,1,1},
		{0,1,0,0,0,0,0,0,0,0},
		{0,1,1,1,1,1,1,1,1,0},
		{0,1,0,0,0,0,0,0,0,0},
		{0,1,0,1,1,1,1,1,1,1},
		{0,1,0,1,1,1,1,0,0,0},
		{0,1,0,0,0,0,0,0,1,0},
		{0,1,1,1,1,1,1,0,1,0},
		{0,0,0,0,0,0,0,0,1,0},
	}, 1))
}
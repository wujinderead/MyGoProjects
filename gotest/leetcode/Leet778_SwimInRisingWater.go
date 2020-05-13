package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/swim-in-rising-water/

// On an N x N grid, each square grid[i][j] represents the elevation at that point (i,j).
// Now rain starts to fall. At time t, the depth of the water everywhere is t. You can
// swim from a square to another 4-directionally adjacent square if and only
// if the elevation of both squares individually are at most t. You can swim infinite
// distance in zero time. Of course, you must stay within the boundaries of the
// grid during your swim.
// You start at the top left square (0, 0). What is the least time until you can
// reach the bottom right square (N-1, N-1)?
// Example 1:
//   Input: [[0,2],[1,3]]
//   Output: 3
//   Explanation:
//     At time 0, you are in grid location (0, 0).
//     You cannot go anywhere else because 4-directionally adjacent neighbors have a
//     higher elevation than t = 0.
//     You cannot reach point (1, 1) until time 3.
//     When the depth of water is 3, we can swim anywhere inside the grid.
// Example 2:
//   Input: [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]
//   Output: 16
//   Explanation:
//        0  1  2  3  4
//       24 23 22 21  5
//       12 13 14 15 16
//       11 17 18 19 20
//       10  9  8  7  6
//     The final route is marked in bold.
//     We need to wait until time 16 so that (0, 0) and (4, 4) are connected.
// Note:
//   2 <= N <= 50.
//   grid[i][j] is a permutation of [0, ..., N*N - 1].

// another method: binary search from grid[0][0] to n*n-1, for each mid,
// use dfs to check if we can reach grid[n-1][n-1], that's also n*n*logn time.
func swimInWater(grid [][]int) int {
    // from (0,0) point, find the minimal value which we can visit.
    // we have n*n points, the heap max is n*n, the time is n*n*logn
    n := len(grid)
	h := heaper(make([][3]int, 0, n*n/2))
	min := 0
	heap.Push(&h, [3]int{grid[0][0], 0, 0})
	visited := make([]bool, n*n)
	visited[0] = true
	for h.Len()>0 {
		pair := heap.Pop(&h).([3]int)
		v, i, j := pair[0], pair[1], pair[2]
		if v>min {
			min = v
		}
		if i==n-1 && j==n-1 {
			return min
		}
		for _, v := range [][2]int{{0,1}, {0,-1}, {1,0}, {-1,0}} {
			ni, nj := i+v[0], j+v[1]
			if ni>=0 && ni<n && nj>=0 && nj<n && !visited[ni*n+nj] {
				visited[ni*n+nj] = true
				heap.Push(&h, [3]int{grid[ni][nj], ni, nj})
			}
		}
	}
	return n*n-1
}

type heaper [][3]int

func (h heaper) Len() int {
	return len(h)
}

func (h heaper) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heaper) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h *heaper) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *heaper) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func main() {
	fmt.Println(swimInWater([][]int{{0,1}, {2,3}}))
	fmt.Println(swimInWater([][]int{{0,2}, {3,1}}))
	fmt.Println(swimInWater([][]int{{2,1}, {3,0}}))
	fmt.Println(swimInWater([][]int{{0,1,2,3,4},{24,23,22,21,5},{12,13,14,15,16},{11,17,18,19,20},{10,9,8,7,6}}))
}
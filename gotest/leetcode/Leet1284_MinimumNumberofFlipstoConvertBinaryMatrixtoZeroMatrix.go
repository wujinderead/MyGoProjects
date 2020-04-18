package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/minimum-number-of-flips-to-convert-binary-matrix-to-zero-matrix/

// Given a m x n binary matrix mat. In one step, you can choose one cell and flip
// it and all the four neighbours of it if they exist (Flip is changing 1 to 0 and
// 0 to 1). A pair of cells are called neighboors if they share one edge.
// Return the minimum number of steps required to convert mat to a zero matrix
// or -1 if you cannot.
// Binary matrix is a matrix with all cells equal to 0 or 1 only.
// Zero matrix is a matrix with all cells equal to 0.
// Example 1:
//   Input: mat = [[0,0],[0,1]]
//   Output: 3
//   Explanation: One possible solution is to flip (1, 0) then (0, 1) and finally (1, 1) as shown.
//     0 0  ->  1 0  ->  0 1  ->  0 0
//     0 1      1 0      1 1      0 0
// Example 2:
//   Input: mat = [[0]]
//   Output: 0
//   Explanation: Given matrix is a zero matrix. We don't need to change it.
// Example 3:
//   Input: mat = [[1,1,1],[1,0,1],[0,0,0]]
//   Output: 6
// Example 4:
//   Input: mat = [[1,0,0],[1,0,0]]
//   Output: -1
//   Explanation: Given matrix can't be a zero matrix
// Constraints:
//   m == mat.length
//   n == mat[0].length
//   1 <= m <= 3
//   1 <= n <= 3
//   mat[i][j] is 0 or 1.

func minFlips(mat [][]int) int {
	// since m and n are small, we can use an integer to represent the states
	// and use bfs to find the minimal flip
	m, n := len(mat), len(mat[0])
	cur := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ind := uint(i*n + j)
			cur = cur | (mat[i][j] << ind)
		}
	}
	if cur == 0 {
		return 0
	}
	visited := make([]bool, 1<<uint(m*n))
	queue := list.New()
	queue.PushBack([2]int{cur, 0})
	visited[cur] = true
	for queue.Len() > 0 {
		tmp := queue.Remove(queue.Front()).([2]int)
		cur, step := tmp[0], tmp[1]
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				t := cur
				// flip mat[i][j]
				t = t ^ (1 << uint(i*n+j))
				for _, v := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
					ni, nj := i+v[0], j+v[1]
					if ni >= 0 && nj >= 0 && ni < m && nj < n {
						t = t ^ (1 << uint(ni*n+nj))
					}
				}
				if t == 0 {
					return step + 1
				}
				if !visited[t] {
					visited[t] = true
					queue.PushBack([2]int{t, step + 1})
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(minFlips([][]int{{0, 0}, {0, 1}}))
	fmt.Println(minFlips([][]int{{0}}))
	fmt.Println(minFlips([][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 0}}))
	fmt.Println(minFlips([][]int{{1, 0, 0}, {1, 0, 0}}))
}

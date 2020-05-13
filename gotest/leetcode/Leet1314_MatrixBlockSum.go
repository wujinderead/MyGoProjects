package main

import (
	"fmt"
)

// https://leetcode.com/problems/matrix-block-sum/

// Given a m * n matrix mat and an integer K, return a matrix answer where each
// answer[i][j] is the sum of all elements mat[r][c] for i - K <= r <= i + K,
// j - K <= c <= j + K, and (r, c) is a valid position in the matrix.
// Example 1:
//   Input: mat = [[1,2,3],[4,5,6],[7,8,9]], K = 1
//   Output: [[12,21,16],[27,45,33],[24,39,28]]
// Example 2:
//   Input: mat = [[1,2,3],[4,5,6],[7,8,9]], K = 2
//   Output: [[45,45,45],[45,45,45],[45,45,45]]
// Constraints:
//   m == mat.length
//   n == mat[i].length
//   1 <= m, n, K <= 100
//   1 <= mat[i][j] <= 100

func matrixBlockSum(mat [][]int, K int) [][]int {
	m, n := len(mat), len(mat[0])
	inter, ans := make([][]int, m), make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		inter[i] = make([]int, n)
	}

	// populate inter matrix by accumulate by col
	for j:=0; j<n; j++ {   // for col j
		for i:=0; i<=K && i<m; i++ {
			inter[0][j] += mat[i][j]
		}
		for i:=1; i<m; i++ {
			inter[i][j] = inter[i-1][j]
			if i+K<m {
				inter[i][j] += mat[i+K][j]
			}
			if i-K-1>=0 {
				inter[i][j] -= mat[i-K-1][j]
			}
		}
	}

	// populate ans matrix by accumulate by row
	for i:=0; i<m; i++ {   // for row i
		for j:=0; j<=K && j<n; j++ {
			ans[i][0] += inter[i][j]
		}
		for j:=1; j<n; j++ {
			ans[i][j] = ans[i][j-1]
			if j+K<n {
				ans[i][j] += inter[i][j+K]
			}
			if j-K-1>=0 {
				ans[i][j] -= inter[i][j-K-1]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(matrixBlockSum([][]int{{1,2,3},{4,5,6},{7,8,9}}, 1))
	fmt.Println(matrixBlockSum([][]int{{1,2,3},{4,5,6},{7,8,9}}, 2))
	fmt.Println(matrixBlockSum([][]int{{1,2,3},{4,5,6},{7,8,9}}, 3))
}
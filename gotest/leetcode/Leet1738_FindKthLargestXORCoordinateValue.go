package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/find-kth-largest-xor-coordinate-value/

// You are given a 2D matrix of size m x n, consisting of non-negative integers.
// You are also given an integer k. The value of coordinate (a, b) of the matrix
// is the XOR of all matrix[i][j] where 0 <= i <= a < m and 0 <= j <= b < n (0-indexed).
// Find the kth largest value (1-indexed) of all the coordinates of matrix.
// Example 1:
//   Input: matrix = [[5,2],[1,6]], k = 1
//   Output: 7
//   Explanation: The value of coordinate (0,1) is 5 XOR 2 = 7, which is the largest value.
// Example 2:
//   Input: matrix = [[5,2],[1,6]], k = 2
//   Output: 5
//   Explanation: The value of coordinate (0,0) is 5 = 5, which is the 2nd largest value.
// Example 3:
//   Input: matrix = [[5,2],[1,6]], k = 3
//   Output: 4
//   Explanation: The value of coordinate (1,0) is 5 XOR 1 = 4, which is the 3rd largest value.
// Example 4:
//   Input: matrix = [[5,2],[1,6]], k = 4
//   Output: 0
//   Explanation: The value of coordinate (1,1) is 5 XOR 2 XOR 1 XOR 6 = 0, which is the 4th largest value.
// Constraints:
//   m == matrix.length
//   n == matrix[i].length
//   1 <= m, n <= 1000
//   0 <= matrix[i][j] <= 10^6
//   1 <= k <= m * n

// for matrix, accumulated XOR is
//   A B         A       A xor B
//   C D      A xor C    (A xor B) xor (A xor C) xor D Xor A = A xor B xor C xor D
func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	valMat := make([][]int, m)
	vals := make([]int, 1, m*n)
	for i := range valMat {
		valMat[i] = make([]int, n)
	}
	valMat[0][0] = matrix[0][0]
	vals[0] = valMat[0][0]
	for i := 1; i < m; i++ {
		valMat[i][0] = matrix[i][0] ^ valMat[i-1][0]
		vals = append(vals, valMat[i][0])
	}
	for j := 1; j < n; j++ {
		valMat[0][j] = matrix[0][j] ^ valMat[0][j-1]
		vals = append(vals, valMat[0][j])
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			valMat[i][j] = matrix[i][j] ^ valMat[i][j-1] ^ valMat[i-1][j] ^ valMat[i-1][j-1]
			vals = append(vals, valMat[i][j])
		}
	}
	sort.Ints(vals)
	return vals[len(vals)-k]
}

func main() {
	for _, v := range []struct {
		mat    [][]int
		k, ans int
	}{
		{[][]int{{5, 2}, {1, 6}}, 1, 7},
		{[][]int{{5, 2}, {1, 6}}, 2, 5},
		{[][]int{{5, 2}, {1, 6}}, 3, 4},
		{[][]int{{5, 2}, {1, 6}}, 4, 0},
	} {
		fmt.Println(kthLargestValue(v.mat, v.k), v.ans)
	}
}

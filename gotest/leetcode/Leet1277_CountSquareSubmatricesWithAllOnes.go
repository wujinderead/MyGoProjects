package main

import (
	"fmt"
)

// https://leetcode.com/problems/count-square-submatrices-with-all-ones

// Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.
// Example 1:
//   Input: matrix =
//     [
//       [0,1,1,1],
//       [1,1,1,1],
//       [0,1,1,1]
//     ]
//   Output: 15
//   Explanation: 
//     There are 10 squares of side 1.
//     There are 4 squares of side 2.
//     There is  1 square of side 3.
//     Total number of squares = 10 + 4 + 1 = 15.
// Example 2:
//   Input: matrix = 
//     [
//       [1,0,1],
//       [1,1,0],
//       [1,1,0]
//     ]
//   Output: 7
//   Explanation: 
//     There are 6 squares of side 1.  
//     There is 1 square of side 2. 
//     Total number of squares = 6 + 1 = 7.
// Constraints:
//     1 <= arr.length <= 300
//     1 <= arr[0].length <= 300
//     0 <= arr[i][j] <= 1

func countSquares(matrix [][]int) int {
	count := 0
	m, n := len(matrix), len(matrix[0])
	
	// initialize max size
	size := make([][]int, m)
	for i := range size {
		size[i] = make([]int, n)
	}
	for i := range matrix {
		size[i][0] = matrix[i][0]
		count += matrix[i][0]
	}
	for j:=1; j<n; j++ {
		size[0][j] = matrix[0][j]
		count += matrix[0][j]
	}

	// dp
	for i:=1; i<m; i++ {
		for j:=1; j<n; j++ {
			if matrix[i][j]==1 {
				size[i][j] = min(size[i-1][j], min(size[i][j-1], size[i-1][j-1]))+1
				count += size[i][j]
			}
		}
	}
	return count
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	fmt.Println(countSquares([][]int{
		{0,1,1,1},
		{1,1,1,1},
		{0,1,1,1},
	}))
	fmt.Println(countSquares([][]int{
		{1,0,1},
		{1,1,0},
		{1,1,0},
	}))
}
package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-falling-path-sum/

// Given a square array of integers A, we want the minimum sum of a falling path through A.
// A falling path starts at any element in the first row, and chooses one element from each row.
// The next row's choice must be in a column that is different from the previous row's column
// by at most one.
// Example 1:
//   Input: [[1,2,3],[4,5,6],[7,8,9]]
//   Output: 12
//   Explanation:
//     The possible falling paths are:
//     [1,4,7], [1,4,8], [1,5,7], [1,5,8], [1,5,9]
//     [2,4,7], [2,4,8], [2,5,7], [2,5,8], [2,5,9], [2,6,8], [2,6,9]
//     [3,5,7], [3,5,8], [3,5,9], [3,6,8], [3,6,9]
//     The falling path with the smallest sum is [1,4,7], so the answer is 12.
// Note:
//   1 <= A.length == A[0].length <= 100
//   -100 <= A[i][j] <= 100

func minFallingPathSum(A [][]int) int {
	prev := make([]int, len(A[0]))
	cur := make([]int, len(A[0]))
	copy(prev, A[0])
	for i := 1; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			cur[j] = prev[j] + A[i][j]
			if j-1 >= 0 {
				cur[j] = min(cur[j], prev[j-1]+A[i][j])
			}
			if j+1 < len(A[0]) {
				cur[j] = min(cur[j], prev[j+1]+A[i][j])
			}
		}
		prev, cur = cur, prev
	}
	amin := prev[0]
	for i := 1; i < len(A[0]); i++ {
		amin = min(amin, prev[i])
	}
	return amin
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	ints := [][]int{
		{12, 4, 2, 13, 10},
		{0, 19, 11, 7, 5},
		{15, 18, 9, 14, 6},
		{8, 1, 16, 17, 3},
	}
	fmt.Println(minFallingPathSum(ints))
	ints = [][]int{
		{-19, 57},
		{-40, -5},
	}
	fmt.Println(minFallingPathSum(ints))
}

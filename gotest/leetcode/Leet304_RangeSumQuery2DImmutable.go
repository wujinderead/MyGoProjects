package main

import "fmt"

// https://leetcode.com/problems/range-sum-query-2d-immutable/

// Given a 2D matrix matrix, find the sum of the elements inside the rectangle defined by its 
// upper left corner (row1, col1) and lower right corner (row2, col2).
// The above rectangle (with the red border) is defined by (row1, col1) = (2, 1) and (row2, col2) = (4, 3), 
// which contains sum = 8.
// Example:
//   Given matrix = [
//     [3, 0, 1, 4, 2],
//     [5, 6, 3, 2, 1],
//     [1, 2, 0, 1, 5],
//     [4, 1, 0, 1, 7],
//     [1, 0, 3, 0, 5]
//    ]
//     sumRegion(2, 1, 4, 3) -> 8
//     sumRegion(1, 1, 2, 2) -> 11
//     sumRegion(1, 2, 2, 4) -> 12
// Note:
//   You may assume that the matrix does not change.
//   There are many calls to sumRegion function.
//   You may assume that row1 ≤ row2 and col1 ≤ col2.

type NumMatrix struct {
    acc [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix)==0 || len(matrix[0])==0 {
		return NumMatrix{make([][]int, 0)}
	}
	m, n := len(matrix), len(matrix[0])
	acc := make([][]int, m+1)
	for i := range acc {
		acc[i] = make([]int, n+1)
	}
	// acc[i][j] is the sum of matrix[i:][j:]
	for i:=m-1; i>=0; i-- {
		for j:=n-1; j>=0; j-- {
			acc[i][j] = matrix[i][j] + acc[i+1][j] + acc[i][j+1] - acc[i+1][j+1]
		}
	}
	return NumMatrix{acc}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if len(this.acc)==0 {
		return 0
	}
	return this.acc[row1][col1] - this.acc[row2+1][col1] - this.acc[row1][col2+1] + this.acc[row2+1][col2+1]
}

func main() {
	obj := Constructor([][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	})
	fmt.Println(obj.SumRegion(2, 1, 4, 3), 8)
	fmt.Println(obj.SumRegion(1, 1, 2, 2), 11)
	fmt.Println(obj.SumRegion(1, 2, 2, 4), 12)
	fmt.Println(obj.SumRegion(1, 1, 1, 1), 6)
	fmt.Println(obj.SumRegion(1, 1, 2, 1), 8)
	fmt.Println(obj.SumRegion(4, 4, 4, 4), 5)
}
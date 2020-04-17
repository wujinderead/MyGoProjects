package main

import "fmt"

// https://leetcode.com/problems/n-queens-ii/

// The n-queens puzzle is the problem of placing n queens on an n×n chessboard such
// that no two queens attack each other.
// Given an integer n, return the number of distinct solutions to the n-queens puzzle.
// Example:
//   Input: 4
//   Output: 2
//   Explanation: There are two distinct solutions to the 4-queens puzzle as shown below.
//     [
//      [".Q..",  // Solution 1
//       "...Q",
//       "Q...",
//       "..Q."],
//
//      ["..Q.",  // Solution 2
//       "Q...",
//       "...Q",
//       ".Q.."]
//     ]

func totalNQueens(n int) int {
	buf := make([][]byte, n)
	for i := range buf {
		buf[i] = make([]byte, n)
		for j := range buf[i] {
			buf[i][j] = '.'
		}
	}
	total := 0
	col := make([]bool, n)
	dia := make([]bool, 2*n-1)
	rev := make([]bool, 2*n-1)
	next(buf, 0, col, dia, rev, &total)
	return total
}

func next(buf [][]byte, row int, col, dia, rev []bool, total *int) {
	if row == len(buf) {
		*total++
	}
	for j := 0; j < len(buf); j++ {
		if !col[j] && !rev[row+j] && !dia[row-j+len(buf)-1] {
			buf[row][j] = 'Q'
			col[j] = true
			rev[row+j] = true
			dia[row-j+len(buf)-1] = true
			next(buf, row+1, col, dia, rev, total)
			buf[row][j] = '.'
			col[j] = false
			rev[row+j] = false
			dia[row-j+len(buf)-1] = false
		}
	}
}

func main() {
	fmt.Println(totalNQueens(3))
	fmt.Println(totalNQueens(4))
	fmt.Println(totalNQueens(5))
	fmt.Println(totalNQueens(6))
	fmt.Println(totalNQueens(7))
	fmt.Println(totalNQueens(8))
	fmt.Println(totalNQueens(9))
	fmt.Println(totalNQueens(10))
}

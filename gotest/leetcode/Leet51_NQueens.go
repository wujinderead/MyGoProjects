package main

import "fmt"

// https://leetcode.com/problems/n-queens/

// The n-queens puzzle is the problem of placing n queens on an n×n chessboard such
// that no two queens attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle.
// Each solution contains a distinct board configuration of the n-queens' placement,
// where 'Q' and '.' both indicate a queen and an empty space respectively.
// Example:
//   Input: 4
//   Output: [
//      [".Q..",  // Solution 1
//       "...Q",
//       "Q...",
//       "..Q."],
//      ["..Q.",  // Solution 2
//       "Q...",
//       "...Q",
//       ".Q.."]
//     ]
//   Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.

func solveNQueens(n int) [][]string {
	buf := make([][]byte, n)
	for i := range buf {
		buf[i] = make([]byte, n)
		for j := range buf[i] {
			buf[i][j] = '.'
		}
	}
	ret := make([][]string, 0)
	col := make([]bool, n)
	dia := make([]bool, 2*n-1)
	rev := make([]bool, 2*n-1)
	next(buf, 0, col, dia, rev, &ret)
	return ret
}

func next(buf [][]byte, row int, col, dia, rev []bool, ret *[][]string) {
	if row == len(buf) {
		aret := make([]string, len(buf))
		for i := range buf {
			aret[i] = string(buf[i])
		}
		*ret = append(*ret, aret)
	}
	for j := 0; j < len(buf); j++ {
		if !col[j] && !rev[row+j] && !dia[row-j+len(buf)-1] {
			buf[row][j] = 'Q'
			col[j] = true
			rev[row+j] = true
			dia[row-j+len(buf)-1] = true
			next(buf, row+1, col, dia, rev, ret)
			buf[row][j] = '.'
			col[j] = false
			rev[row+j] = false
			dia[row-j+len(buf)-1] = false
		}
	}
}

func main() {
	fmt.Println(solveNQueens(3))
	fmt.Println(solveNQueens(4))
	fmt.Println(solveNQueens(5))
	fmt.Println(solveNQueens(6))
}

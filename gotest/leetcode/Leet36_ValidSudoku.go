package main

import "fmt"

// https://leetcode.com/problems/valid-sudoku/

// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated
// according to the following rules:
//   Each row must contain the digits 1-9 without repetition.
//   Each column must contain the digits 1-9 without repetition.
//   Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 with out repetition.
// Note:
//   A Sudoku board (partially filled) could be valid but is not necessarily solva ble.
//   Only the filled cells need to be validated according to the mentioned rules.
// Example 1:
//   Input: board =
//     [["5","3",".",".","7",".",".",".","."],
//     ["6",".",".","1","9","5",".",".","."],
//     [".","9","8",".",".",".",".","6","."],
//     ["8",".",".",".","6",".",".",".","3"],
//     ["4",".",".","8",".","3",".",".","1"],
//     ["7",".",".",".","2",".",".",".","6"],
//     [".","6",".",".",".",".","2","8","."],
//     [".",".",".","4","1","9",".",".","5"],
//     [".",".",".",".","8",".",".","7","9"]]
//   Output: true
// Example 2:
//   Input: board =
//     [["8","3",".",".","7",".",".",".","."],
//     ["6",".",".","1","9","5",".",".","."],
//     [".","9","8",".",".",".",".","6","."],
//     ["8",".",".",".","6",".",".",".","3"],
//     ["4",".",".","8",".","3",".",".","1"],
//     ["7",".",".",".","2",".",".",".","6"],
//     [".","6",".",".",".",".","2","8","."],
//     [".",".",".","4","1","9",".",".","5"],
//     [".",".",".",".","8",".",".","7","9"]]
//   Output: false
//   Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8.
//     Since there are two 8's in the top left 3x3 sub-box, it is invalid.
// Constraints:
//   board.length == 9
//   board[i].length == 9
//   board[i][j] is a digit or '.'

func isValidSudoku(board [][]byte) bool {
	// for each row
	for i := 0; i < 9; i++ {
		mask := 0
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			if mask&(1<<uint(board[i][j]-'0')) > 0 {
				return false
			}
			mask = mask | (1 << uint(board[i][j]-'0'))
		}
	}
	// for each column
	for j := 0; j < 9; j++ {
		mask := 0
		for i := 0; i < 9; i++ {
			if board[i][j] == '.' {
				continue
			}
			if mask&(1<<uint(board[i][j]-'0')) > 0 {
				return false
			}
			mask = mask | (1 << uint(board[i][j]-'0'))
		}
	}
	// for each 3x3 cell
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			mask := 0
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					ii, jj := i+k, j+l
					if board[ii][jj] == '.' {
						continue
					}
					if mask&(1<<uint(board[ii][jj]-'0')) > 0 {
						return false
					}
					mask = mask | (1 << uint(board[ii][jj]-'0'))
				}
			}
		}
	}
	return true
}

func main() {
	for _, v := range []struct {
		b   [][]byte
		ans bool
	}{
		{[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}, true},
		{[][]byte{
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}, false},
		{[][]byte{
			{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
			{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
			{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
			{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
			{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
			{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
			{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
		}, false},
	} {
		fmt.Println(isValidSudoku(v.b), v.ans)
	}
}

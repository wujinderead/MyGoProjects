package main

import (
	"fmt"
)

// https://leetcode.com/problems/surrounded-regions/

// Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
// Example:
//   X X X X
//   X O O X
//   X X O X
//   X O X X
// After running your function, the board should be:
//   X X X X
//   X X X X
//   X X X X
//   X O X X
// Explanation:
// Surrounded regions shouldn’t be on the border,
// which means that any 'O' on the border of the board are not flipped to 'X'.
// Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'.
// Two cells are connected if they are adjacent cells connected horizontally or vertically.

func solve(board [][]byte) {
	// row or col <=2, must be on the border, no need for
	if len(board) <= 2 || len(board[0]) <= 2 {
		return
	}
	marked := make(map[int]struct{}) // mark 'O' connected to the border
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		// first col
		if _, ok := marked[(i<<16)+0]; !ok && board[i][0] == 'O' {
			dfs(i, 0, marked, board)
		}
		// last col
		if _, ok := marked[(i<<16)+n-1]; !ok && board[i][n-1] == 'O' {
			dfs(i, n-1, marked, board)
		}
	}
	for j := 1; j < n-1; j++ {
		// first row
		if _, ok := marked[(0<<16)+j]; !ok && board[0][j] == 'O' {
			dfs(0, j, marked, board)
		}
		// last row
		if _, ok := marked[((m-1)<<16)+j]; !ok && board[m-1][j] == 'O' {
			dfs(m-1, j, marked, board)
		}
	}
	// flip unmarked non-bordered 'O'
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if _, ok := marked[(i<<16)+j]; !ok && board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(i, j int, marked map[int]struct{}, board [][]byte) {
	marked[(i<<16)+j] = struct{}{}
	if _, ok := marked[((i+1)<<16)+j]; !ok && i+1 < len(board) && board[i+1][j] == 'O' {
		dfs(i+1, j, marked, board)
	}
	if _, ok := marked[((i-1)<<16)+j]; !ok && i-1 >= 0 && board[i-1][j] == 'O' {
		dfs(i-1, j, marked, board)
	}
	if _, ok := marked[(i<<16)+j+1]; !ok && j+1 < len(board[0]) && board[i][j+1] == 'O' {
		dfs(i, j+1, marked, board)
	}
	if _, ok := marked[(i<<16)+j-1]; !ok && j-1 >= 0 && board[i][j-1] == 'O' {
		dfs(i, j-1, marked, board)
	}
}

func main() {
	board := [][]byte{
		{'X', 'O', 'X', 'X'},
		{'X', 'O', 'O', 'O'},
		{'X', 'O', 'X', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	for i := 0; i < len(board); i++ {
		fmt.Println(string(board[i]))
	}
	board = [][]byte{
		{'O', 'X', 'O'},
		{'X', 'O', 'X'},
		{'O', 'X', 'O'},
	}
	solve(board)
	for i := 0; i < len(board); i++ {
		fmt.Println(string(board[i]))
	}
}

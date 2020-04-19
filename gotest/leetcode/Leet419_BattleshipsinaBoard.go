package main

import "fmt"

// https://leetcode.com/problems/battleships-in-a-board/

// Given an 2D board, count how many battleships are in it. The battleships are
// represented with 'X's, empty slots are represented with '.'s. You may assume the
// following rules:
// 1. You receive a valid board, made of only battleships or empty slots.
// 2. Battleships can only be placed horizontally or vertically. In other words, they
//    can only be made of the shape 1xN (1 row, N columns) or Nx1 (N rows, 1 column),
//    where N can be of any size.
// 3. At least one horizontal or vertical cell separates between two battleships -
//    there are no adjacent battleships.
// Example:
//   X..X
//   ...X
//   ...X
//   In the above board there are 2 battleships.
// Invalid Example:
//   ...X
//   XXXX
//   ...X
//   This is an invalid board that you will not receive - as battleships will always
//   have a cell separating between them.
// Follow up: Could you do it in one-pass, using only O(1) extra memory and with
// out modifying the value of the board?

func countBattleships(board [][]byte) int {
	// O(1) extra space
	count1, single1 := 0, 0
	for i := 0; i < len(board); i++ {
		good, bad := 0, 0
		for j := 0; j <= len(board[0]); j++ {
			if j == len(board[0]) || board[i][j] == '.' {
				if good > 0 && bad == 0 {
					count1++
					if good == 1 {
						single1++
					}
				}
				good, bad = 0, 0
			}
			if j < len(board[0]) && board[i][j] == 'X' {
				if (i-1 >= 0 && board[i-1][j] == 'X') || (i+1 < len(board) && board[i+1][j] == 'X') {
					bad++
				} else {
					good++
				}
			}
		}
	}
	count2, single2 := 0, 0
	for j := 0; j < len(board[0]); j++ {
		good, bad := 0, 0
		for i := 0; i <= len(board); i++ {
			if i == len(board) || board[i][j] == '.' {
				if good > 0 && bad == 0 {
					count2++
					if good == 1 {
						single2++
					}
				}
				good, bad = 0, 0
			}
			if i < len(board) && board[i][j] == 'X' {
				if (j-1 >= 0 && board[i][j-1] == 'X') || (j+1 < len(board[0]) && board[i][j+1] == 'X') {
					bad++
				} else {
					good++
				}
			}
		}
	}
	return count1 + count2 - single1
}

func main() {
	fmt.Println(countBattleships([][]byte{
		{'X', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
		{'.', '.', '.', 'X'},
	}))
	fmt.Println(countBattleships([][]byte{
		{'.', '.', '.', 'X'},
		{'X', 'X', 'X', 'X'},
		{'.', '.', '.', 'X'},
	}))
	fmt.Println(countBattleships([][]byte{
		{'X', 'X', '.', 'X', 'X', '.', 'X', 'X'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', 'X', '.', 'X', '.', '.', '.', 'X'},
		{'X', '.', '.', 'X', '.', '.', 'X', '.'},
		{'X', 'X', '.', 'X', '.', '.', '.', '.'},
		{'X', '.', '.', 'X', '.', '.', 'X', 'X'},
		{'X', '.', '.', '.', 'X', 'X', '.', 'X'},
	}))
}

package backtracking

import (
	"fmt"
)

var deltas = [][]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

// the order of solution would greatly affect the result
// use the following delta sequence for size 8, it take hours to get the result
// var deltas = [][]int{{2, 1}, {2, -1}, {-2, 1}, {-2, -1}, {1, 2}, {1, -2}, {-1, 2}, {-1, -2}}

func knightTour(size int) {
	board := make([][]int, size, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size, size)
	}

	board[0][0] = 1
	if knightTourNext(board, 0, 0, 1) {
		fmt.Println(board)
	} else {
		fmt.Println("no solution.")
	}
}

func knightTourNext(board [][]int, curx, cury, count int) bool {
	if count == len(board)*len(board) {
		return true
	}
	for i := range deltas {
		dx := deltas[i][0]
		dy := deltas[i][1]
		if canGoto(board, curx+dx, cury+dy) {
			//fmt.Println(count, curx+dx, cury+dy)
			board[curx+dx][cury+dy] = count + 1
			if knightTourNext(board, curx+dx, cury+dy, count+1) {
				return true
			}
			board[curx+dx][cury+dy] = 0
		}
	}
	return false
}

func canGoto(board [][]int, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board) && board[x][y] == 0
}

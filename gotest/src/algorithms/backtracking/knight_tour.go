package backtracking

import (
	"fmt"
)

var deltas = [][]int{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}}

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

func getDegree(board [][]int, x, y int) int {
	unvisitedAdjacency := 0
	for i := range deltas {
		dx := deltas[i][0]
		dy := deltas[i][1]
		if canGoto(board, x+dx, y+dy) {
			unvisitedAdjacency++
		}
	}
	return unvisitedAdjacency
}

// Warnsdorff’s algorithm for Knight’s tour problem:
// start from any initial position of the knight on the board,
// always move to an adjacent, unvisited square with minimal degree
// (minimum number of unvisited adjacent). it will guarantee to get a solution.
// On an 8 × 8 board, there are exactly 26,534,728,821,064 directed closed tours.
func knightTourWarnsdorff(size int) {
	board := make([][]int, size, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size, size)
	}

	board[0][0] = 1
	prevx, prevy := 0, 0
	for i := 2; i < size*size; i++ {
		minDegree := 8
		minx := 0
		miny := 0
		for i := range deltas {
			dx := deltas[i][0]
			dy := deltas[i][1]
			if canGoto(board, prevx+dx, prevy+dy) {
				curDegree := getDegree(board, prevx+dx, prevy+dy)
				if curDegree <= minDegree {
					minDegree = curDegree
					minx = prevx + dx
					miny = prevy + dy
				}
			}
		}
		board[minx][miny] = i
		prevx = minx
		prevy = miny
	}
	fmt.Println(board)
}

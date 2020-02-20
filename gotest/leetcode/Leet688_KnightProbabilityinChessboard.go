package main

import "fmt"

// https://leetcode.com/problems/knight-probability-in-chessboard/

// On an NxN chessboard, a knight starts at the r-th row and c-th column
// and attempts to make exactly K moves. The rows and columns are 0 indexed,
// so the top-left square is (0, 0), and the bottom-right square is (N-1, N-1).
// A chess knight has 8 possible moves it can make, as illustrated below.
// Each move is two squares in a cardinal direction, then one square in an orthogonal direction.
// Each time the knight is to move, it chooses one of eight possible moves uniformly
// at random (even if the piece would go off the chessboard) and moves there.
// The knight continues moving until it has made exactly K moves or has moved off the chessboard.
// Return the probability that the knight remains on the board after it has stopped moving.
// Example:
//   Input: 3, 2, 0, 0
//   Output: 0.0625
//   Explanation:
//     There are two moves (to (1,2), (2,1)) that will keep the knight on the board.
//     From each of those positions, there are also two moves that will keep the knight on the board.
//     The total probability the knight stays on the board is 0.0625.
// Note:
//    N will be between 1 and 25.
//    K will be between 0 and 100.
//    The knight always initially starts on the board.

func knightProbability(N int, K int, r int, c int) float64 {
	if K == 0 {
		return 1
	}
	moves := [][2]int{{1, 2}, {2, 1}, {-1, 2}, {-2, 1}, {1, -2}, {2, -1}, {-1, -2}, {-2, -1}}
	old := make([]float64, N*N)
	neu := make([]float64, N*N)
	add2d(old, r, c, N, 1)
	prob := 1.0
	for k := 1; k <= K; k++ {
		in := 0.0
		out := 0.0
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				num := getclear2d(old, i, j, N)
				if num == 0 {
					continue
				}
				for m := range moves {
					if i+moves[m][0] >= N || i+moves[m][0] < 0 || j+moves[m][1] >= N || j+moves[m][1] < 0 {
						out += num
						continue
					}
					in += num
					add2d(neu, i+moves[m][0], j+moves[m][1], N, num)
				}
			}
		}
		if in == 0 {
			return 0
		}
		fmt.Println(neu)
		old, neu = neu, old
		prob *= in / (in + out)
	}
	return prob
}

func getclear2d(arr []float64, i, j, col int) float64 {
	a := arr[i*col+j]
	arr[i*col+j] = 0
	return a
}

func add2d(arr []float64, i, j, col int, value float64) {
	arr[i*col+j] += value
}

func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
	fmt.Println(knightProbability(3, 2, 1, 1))
	fmt.Println(knightProbability(8, 30, 6, 4))
}

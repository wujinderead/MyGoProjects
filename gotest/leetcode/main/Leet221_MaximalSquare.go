package main

import (
	"fmt"
)

func maximalSquare(matrix [][]byte) int {
	maximal := make([][]byte, len(matrix))
	for i := range maximal {
		maximal[i] = make([]byte, len(matrix[0]))
	}
	var max byte = 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if i-1 >= 0 && j-1 >= 0 && matrix[i][j] == '1' {
				maximal[i][j] = min(maximal[i-1][j-1], min(maximal[i-1][j], maximal[i][j-1])) + 1
			} else {
				maximal[i][j] = matrix[i][j] - '0'
			}
			if maximal[i][j] > max {
				max = maximal[i][j]
			}
		}
	}
	fmt.Println(maximal)
	return int(max) * int(max)
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

func main() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	fmt.Println(maximalSquare(matrix))
	matrix = [][]byte{
		{'1', '1', '1', '0', '0'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '1'},
	}
	fmt.Println(maximalSquare(matrix))
	matrix = [][]byte{
		{'1', '0', '1', '0', '0', '1', '1', '1', '0'},
		{'1', '1', '1', '0', '0', '0', '0', '0', '1'},
		{'0', '0', '1', '1', '0', '0', '0', '1', '1'},
		{'0', '1', '1', '0', '0', '1', '0', '0', '1'},
		{'1', '1', '0', '1', '1', '0', '0', '1', '0'},
		{'0', '1', '1', '1', '1', '1', '1', '0', '1'},
		{'1', '0', '1', '1', '1', '0', '0', '1', '0'},
		{'1', '1', '1', '0', '1', '0', '0', '0', '1'},
		{'0', '1', '1', '1', '1', '0', '0', '1', '0'},
		{'1', '0', '0', '1', '1', '1', '0', '0', '0'},
	}
	fmt.Println(maximalSquare(matrix))
}

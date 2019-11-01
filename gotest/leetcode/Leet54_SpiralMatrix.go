package main

import "fmt"

// https://leetcode.com/problems/spiral-matrix/

// Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.
// Example:
//   Input:
//   [
//    [1, 2, 3, 4],
//    [5, 6, 7, 8],
//    [9,10,11,12]
//   ]
//   Output: [1,2,3,4,8,12,11,10,9,5,6,7]
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	leftmost, upmost := 0, 0
	rightmost, downmost := len(matrix[0])-1, len(matrix)-1
	ans := make([]int, len(matrix)*len(matrix[0]))
	index := 0
	for {
		// to right
		if index < len(ans) {
			for i := leftmost; i <= rightmost; i++ {
				ans[index] = matrix[upmost][i]
				index++
			}
			//fmt.Println(ans, leftmost, rightmost, upmost, downmost)
			upmost++
		}
		// to down
		if index < len(ans) {
			for i := upmost; i <= downmost; i++ {
				ans[index] = matrix[i][rightmost]
				index++
			}
			//fmt.Println(ans, leftmost, rightmost, upmost, downmost)
			rightmost--
		}
		// to left
		if index < len(ans) {
			for i := rightmost; i >= leftmost; i-- {
				ans[index] = matrix[downmost][i]
				index++
			}
			//fmt.Println(ans, leftmost, rightmost, upmost, downmost)
			downmost--
		}
		// to up
		if index < len(ans) {
			for i := downmost; i >= upmost; i-- {
				ans[index] = matrix[i][leftmost]
				index++
			}
			//fmt.Println(ans, leftmost, rightmost, upmost, downmost)
			leftmost++
		}
		// terminate
		if index == len(ans) {
			break
		}
	}
	return ans
}

func main() {
	fmt.Println(spiralOrder([][]int{{1}}))
	fmt.Println(spiralOrder([][]int{{1, 2}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3}}))
	fmt.Println(spiralOrder([][]int{{1}, {2}}))
	fmt.Println(spiralOrder([][]int{{1}, {2}, {3}}))
	fmt.Println(spiralOrder([][]int{{1, 2}, {3, 4}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}}))
	fmt.Println(spiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}))
	fmt.Println(spiralOrder([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}))
}

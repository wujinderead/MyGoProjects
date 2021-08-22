package main

import "fmt"

// https://leetcode.com/problems/maximum-matrix-sum/

// You are given an n x n integer matrix. You can do the following operation any number of times:
//   Choose any two adjacent elements of matrix and multiply each of them by -1.
//   Two elements are considered adjacent if and only if they share a border.
// Your goal is to maximize the summation of the matrix's elements. Return the maximum sum
// of the matrix's elements using the operation mentioned above.
// Example 1:
//   Input: matrix = [[1,-1],[-1,1]]
//   Output: 4
//   Explanation: We can follow the following steps to reach sum equals 4:
//     - Multiply the 2 elements in the first row by -1.
//     - Multiply the 2 elements in the first column by -1.
// Example 2:
//   Input: matrix = [[1,2,3],[-1,-2,-3],[1,2,3]]
//   Output: 16
//   Explanation: We can follow the following step to reach sum equals 16:
//     - Multiply the 2 last elements in the second row by -1.
// Constraints:
//   n == matrix.length == matrix[i].length
//   2 <= n <= 250
//   -10^5 <= matrix[i][j] <= 10^5

// we can always pair two negative numbers together and revert it
func maxMatrixSum(matrix [][]int) int64 {
	sum := 0
	countNeg := 0
	minabs := int(10e6)
	for i := range matrix {
		for j := range matrix[0] {
			if matrix[i][j] < 0 {
				countNeg++
				sum += -matrix[i][j]
				if -matrix[i][j] < minabs {
					minabs = -matrix[i][j]
				}
			} else if matrix[i][j] == 0 { // an zero can flip a negative
				countNeg--
			} else {
				sum += matrix[i][j]
				if matrix[i][j] < minabs {
					minabs = matrix[i][j]
				}
			}
		}
	}
	if countNeg > 0 && countNeg%2 == 1 { // if odd negative number, minus the negative number with minimal abs
		sum -= 2 * minabs
	}
	return int64(sum)
}

func main() {
	for _, v := range []struct {
		m   [][]int
		ans int
	}{
		{[][]int{{1, -1}, {-1, 1}}, 4},
		{[][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}, 16},
		{[][]int{{-1, 0, -1}, {-2, 1, 3}, {3, 2, 2}}, 15},
		{[][]int{{2, 9, 3}, {5, 4, -4}, {1, 7, 1}}, 34},
	} {
		fmt.Println(maxMatrixSum(v.m), v.ans)
	}
}

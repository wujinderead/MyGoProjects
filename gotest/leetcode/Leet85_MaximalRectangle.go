package main

import "fmt"

// https://leetcode.com/problems/maximal-rectangle/

// Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle
// containing only 1's and return its area.
// Example:
//   Input:
//     [
//       ["1","0","1","0","0"],
//       ["1","0","1","1","1"],
//       ["1","1","1","1","1"],
//       ["1","0","0","1","0"]
//     ]
//   Output: 6

// let the matrix be m lines n cols, then for the first k (1<=k<=m) lines,
// we can deem it as a histogram with n cols, and we can calculate the max rectangle in O(n) time.
// thus we can get the max rectangle in the matrix in O(mn) time.
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	hist := make([]int, n)
	stack := make([]int, n+1)
	allmax1 := 0
	for i := 0; i < m; i++ {
		// update histogram
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				hist[j] += 1
			} else {
				hist[j] = 0
			}
		}
		// reset stack for each iteration
		slen := 0
		// calculate max rectangle in histogram
		for j := 0; j <= n; j++ {
			h := 0 // add 0 as the last element to trigger final compute
			if j < n {
				h = hist[j]
			}
			for slen > 0 && h < hist[stack[slen-1]] {
				p := stack[slen-1] // pop the index
				slen--
				wid := j
				if slen > 0 {
					wid = j - 1 - stack[slen-1] // wid = (slen==0 ? j : j-1-stack.peek)
				}
				cur := hist[p] * wid
				if cur > allmax1 {
					allmax1 = cur
				}
			}
			stack[slen] = j
			slen++
		}
	}
	return allmax1
}

func main() {
	fmt.Println(maximalRectangle([][]byte{
		{1, 0, 1, 0, 0},
		{1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1},
		{1, 0, 0, 1, 0},
	}))
	fmt.Println(maximalRectangle([][]byte{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 1, 1, 1},
		{0, 1, 1, 1, 1},
		{1, 1, 1, 1, 1},
	}))
}

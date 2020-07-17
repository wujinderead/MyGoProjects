package main

import (
	"fmt"
)

// https://leetcode.com/problems/count-submatrices-with-all-ones/

// Given a rows * columns matrix mat of ones and zeros, return how many submatrices have all ones.
// Example 1:
//   Input: mat = [[1,0,1],
//                 [1,1,0],
//                 [1,1,0]]
//   Output: 13
//   Explanation:
//     There are 6 rectangles of side 1x1.
//     There are 2 rectangles of side 1x2.
//     There are 3 rectangles of side 2x1.
//     There is 1 rectangle of side 2x2. 
//     There is 1 rectangle of side 3x1.
//     Total number of rectangles = 6 + 2 + 3 + 1 + 1 = 13.
// Example 2:
//   Input: mat = [[0,1,1,0],
//                 [0,1,1,1],
//                 [1,1,1,0]]
//   Output: 24
//   Explanation:
//     There are 8 rectangles of side 1x1.
//     There are 5 rectangles of side 1x2.
//     There are 2 rectangles of side 1x3. 
//     There are 4 rectangles of side 2x1.
//     There are 2 rectangles of side 2x2. 
//     There are 2 rectangles of side 3x1. 
//     There is 1 rectangle of side 3x2. 
//     Total number of rectangles = 8 + 5 + 2 + 4 + 2 + 2 + 1 = 24.
// Example 3:
//   Input: mat = [[1,1,1,1,1,1]]
//   Output: 21
// Example 4:
//   Input: mat = [[1,0,1],[0,1,0],[1,0,1]]
//   Output: 5
// Constraints:
//   1 <= rows <= 150
//   1 <= columns <= 150
//   0 <= mat[i][j] <= 1

// naive dp, O(m*n*m)
func numSubmat1(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	left := make([][]int, m)   // how many consecutive 1s left to a node
	for i:=0; i<m; i++ {
		left[i] = make([]int, n)
		for j:=0; j<n; j++ {
			if j>0 && mat[i][j]==1 {
				left[i][j] = left[i][j-1]+1
			} else {
				left[i][j] = mat[i][j]
			}
		}
	}

	// dp 
	count := 0
	for i:=0; i<m; i++ {
		for j:=0; j<n; j++ {
			k := i
			min := left[k][j]
			for k>=0 && left[k][j]>0 {
				if left[k][j]<min {
					min = left[k][j]
				}
				count += min
				k--
			}
		}
	}
	return count
}

// histogram-like method, O(m*n)
func numSubmat(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	hist := make([]int, n)
	count := 0
	stack := make([]int, n)
	sum := make([]int, n)
	for i:=0; i<m; i++ {
		// make histogram for each line
		for j:=0; j<n; j++ {
			if mat[i][j]==1 {
				hist[j] += 1
			} else {
				hist[j] = 0
			}
		}
		// process histogram
		slen := 0
		for j:=0; j<n; j++ {
			for slen>0 && hist[j]<=hist[stack[slen-1]] {
				slen--       // pop larger height
			}
			if slen>0 {
				peek := stack[slen-1]
				sum[j] = sum[peek]
				sum[j] += hist[j]*(j-peek)
			} else {
				sum[j] = hist[j]*(j+1)
			}
			stack[slen] = j   // push j
			slen++
		}
		for j:=0; j<n; j++ {
			count += sum[j]
			sum[j] = 0
		}
	}
	return count
}

func main() {
	for _, v := range []struct{arr [][]int;res int}{
		{[][]int{{1,1,1,1,1,1}}, 21},
		{[][]int{{1,1,1},{1,0,1},{1,1,1}}, 20},
		{[][]int{{1,0,1},{0,1,0},{1,0,1}}, 5},
		{[][]int{{1,0,1,0}}, 2},
		{[][]int{{1,0,1},{1,1,0},{1,1,0}}, 13},
		{[][]int{{0,1,1,0},{0,1,1,1},{1,1,1,0}}, 24},
		{[][]int{{0,1,1,1,1},{1,1,1,1,0},{0,1,0,0,1},{1,1,1,1,1},{0,1,1,1,1},{1,1,1,1,1}}, 112},
	} {
		fmt.Println(numSubmat(v.arr), v.res)
		fmt.Println(numSubmat1(v.arr), v.res)		
	}
}
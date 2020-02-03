package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-falling-path-sum-ii/

// Given a square grid of integers arr, a falling path with non-zero shifts is a
// choice of exactly one element from each row of arr, such that no two elements
// chosen in adjacent rows are in the same column.
// Return the minimum sum of a falling path with non-zero shifts.
// Example 1:
//   Input: arr = [[1,2,3],[4,5,6],[7,8,9]]
//   Output: 13
//   Explanation:
//     The possible falling paths are:
//     [1,5,9], [1,5,7], [1,6,7], [1,6,8],
//     [2,4,8], [2,4,9], [2,6,7], [2,6,8],
//     [3,4,8], [3,4,9], [3,5,7], [3,5,9]
//     The falling path with the smallest sum is [1,5,7], so the answer is 13.
// Constraints:
//   1 <= arr.length == arr[i].length <= 200
//   -99 <= arr[i][j] <= 99

func minFallingPathSum(arr [][]int) int {
	if len(arr) == 1 {
		return arr[0][0]
	}
	prev := make([]int, len(arr[0]))
	copy(prev, arr[0])
	for i := 1; i < len(arr); i++ {
		// find the minimal and second minimal value
		min1 := 0
		min2 := 1
		if prev[min2] < prev[min1] {
			min1, min2 = min2, min1
		}
		for j := 2; j < len(arr[0]); j++ {
			if prev[j] < prev[min1] {
				min2 = min1
				min1 = j
			} else if prev[j] < prev[min2] {
				min2 = j
			}
		}
		pmin1 := prev[min1]
		pmin2 := prev[min2]
		for j := 0; j < len(arr[0]); j++ {
			if j == min1 { // if the same col, add second minimal value
				prev[j] = pmin2 + arr[i][j]
			} else { // otherwise, add minimal value
				prev[j] = pmin1 + arr[i][j]
			}
		}
	}
	amin := prev[0]
	for i := 1; i < len(arr[0]); i++ {
		if prev[i] < amin {
			amin = prev[i]
		}
	}
	return amin
}

func main() {
	ints := [][]int{
		{21, 4, 2, 13, 10},
		{25, 19, 11, 7, 5},
		{23, 18, 9, 14, 6},
		{8, 1, 20, 17, 3},
		{16, 22, 24, 15, 12},
	}
	fmt.Println(minFallingPathSum(ints))
	ints = [][]int{{2}}
	fmt.Println(minFallingPathSum(ints))
	ints = [][]int{
		{2, 4},
		{6, 7},
	}
	fmt.Println(minFallingPathSum(ints))
	ints = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(minFallingPathSum(ints))
}

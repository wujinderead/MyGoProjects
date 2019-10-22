package leetcode

import "fmt"

// https://leetcode.com/problems/triangle/

// Given a triangle, find the minimum path sum from top to bottom.
// Each step you may move to adjacent numbers on the row below.
// For example, given the following triangle
// [
//      [2],
//     [3,4],
//    [6,5,7],
//   [4,1,8,3]
// ]
// The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
// Note:
// Bonus point if you are able to do this using only O(n) extra space,
// where n is the total number of rows in the triangle.

func minimumTotal(triangle [][]int) int {
	sums1 := make([]int, len(triangle))
	sums2 := make([]int, len(triangle))
	sums1[0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		sums2[0] = sums1[0] + triangle[i][0]
		sums2[i] = sums1[i-1] + triangle[i][i]
		for j := 1; j < i; j++ {
			sums2[j] = min(sums1[j], sums1[j-1]) + triangle[i][j]
		}
		// fmt.Println(sums1, sums2)
		sums1, sums2 = sums2, sums1
	}
	min := 0x7ffffffff
	for i := range sums1 {
		if sums1[i] < min {
			min = sums1[i]
		}
	}
	return min
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
	fmt.Println(minimumTotal([][]int{
		{2},
	}))
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
	}))
}

package leetcode

import "fmt"

// https://leetcode.com/problems/unique-paths/
// A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
// The robot can only move either down or right at any point in time.
// The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
// How many possible unique paths are there?
func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	if m < n {
		m, n = n, m // make m smaller
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			arr[j] = arr[j-1] + arr[j]
		}
	}
	return arr[n-1]
}

func main() {
	fmt.Println(uniquePaths(5, 3))
	fmt.Println(uniquePaths(3, 5))
	fmt.Println(uniquePaths(1, 5))
	fmt.Println(uniquePaths(2, 2))
}

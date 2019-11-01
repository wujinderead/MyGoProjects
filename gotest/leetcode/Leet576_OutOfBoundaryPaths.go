package main

import "fmt"

// https://leetcode.com/problems/out-of-boundary-paths/

// There is an m by n grid with a ball. Given the start coordinate (i,j) of the ball,
// you can move the ball to adjacent cell or cross the grid boundary in four directions
// (up, down, left, right). However, you can at most move N times. Find out the number
// of paths to move the ball out of grid boundary. The answer may be very large,
// return it after mod 10^9 + 7.
// Example 1:
//   Input: m = 2, n = 2, N = 2, i = 0, j = 0
//   Output: 6
//   Explanation:
// Example 2:
//   Input: m = 1, n = 3, N = 3, i = 0, j = 1
//   Output: 12
//   Explanation:
// Note:
//   Once you move the ball out of boundary, you cannot move it back.
//   The length and height of the grid is in range [1,50].
//   N is in range [0,50].

func findPaths(m int, n int, N int, i0 int, j0 int) int {
	// O(mnN)
	modulus := 1000000007
	re := 0
	old := make([]int, m*n)
	new := make([]int, m*n)
	add2d(old, i0, j0, n, 1)
	for k := 0; k < N; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				oldv := getset2d(old, i, j, n) // get then set 0
				oldv %= modulus
				if oldv == 0 {
					continue
				}
				// down
				if i+1 >= m {
					re += oldv
				} else {
					add2d(new, i+1, j, n, oldv)
				}
				// up
				if i-1 < 0 {
					re += oldv
				} else {
					add2d(new, i-1, j, n, oldv)
				}
				// right
				if j+1 >= n {
					re += oldv
				} else {
					add2d(new, i, j+1, n, oldv)
				}
				// left
				if j-1 < 0 {
					re += oldv
				} else {
					add2d(new, i, j-1, n, oldv)
				}
			}
		}
		old, new = new, old
		re %= modulus
	}
	return re
}

func getset2d(arr []int, i, j, col int) int {
	a := arr[i*col+j]
	arr[i*col+j] = 0
	return a
}

func add2d(arr []int, i, j, col, value int) {
	arr[i*col+j] += value
}

func main() {
	fmt.Println(findPaths(1, 1, 1, 0, 0))      // 4
	fmt.Println(findPaths(2, 2, 2, 0, 0))      // 6
	fmt.Println(findPaths(1, 3, 3, 0, 1))      // 12
	fmt.Println(findPaths(3, 3, 5, 1, 1))      // 180
	fmt.Println(findPaths(3, 3, 6, 1, 0))      // 365
	fmt.Println(findPaths(30, 40, 50, 13, 29)) // 438333906
	fmt.Println(findPaths(49, 49, 20, 24, 24)) // 0
}

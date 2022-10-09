package main

import "fmt"

// https://leetcode.com/problems/paths-in-matrix-whose-sum-is-divisible-by-k/

// You are given a 0-indexed m x n integer matrix grid and an integer k. You are currently at
// position (0, 0) and you want to reach position (m - 1, n - 1) moving only down or right.
// Return the number of paths where the sum of the elements on the path is divisible by k.
// Since the answer may be very large, return it modulo 10⁹ + 7.
// Example 1:
//   Input: grid = [[5,2,4],[3,0,5],[0,7,2]], k = 3
//   Output: 2
//   Explanation: There are two paths where the sum of the elements on the path is divisible by k.
//     The first path highlighted in red has a sum of 5 + 2 + 4 + 5 + 2 = 18 which is divisible by 3.
//     The second path highlighted in blue has a sum of 5 + 3 + 0 + 5 + 2 = 15 which is divisible by 3.
// Example 2:
//   Input: grid = [[0,0]], k = 5
//   Output: 1
//   Explanation: The path highlighted in red has a sum of 0 + 0 = 0 which is divisible by 5.
// Example 3:
//   Input: grid = [[7,3,4,9],[2,3,6,2],[2,3,7,0]], k = 1
//   Output: 10
//   Explanation: Every integer is divisible by 1 so the sum of the elements on
//     every possible path is divisible by k.
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m, n <= 5 * 10⁴
//   1 <= m * n <= 5 * 10⁴
//   0 <= grid[i][j] <= 100
//   1 <= k <= 50

func numberOfPaths(grid [][]int, k int) int {
	const P = int(1e9 + 7)
	m, n := len(grid), len(grid[0])
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}

	// initialize
	dp[0][0][grid[0][0]%k] = 1
	s := grid[0][0]
	for i := 1; i < m; i++ {
		s = (s + grid[i][0]) % k
		dp[i][0][s] = 1
	}
	s = grid[0][0]
	for j := 1; j < n; j++ {
		s = (s + grid[0][j]) % k
		dp[0][j][s] = 1
	}

	// dp
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			for x := 0; x < k; x++ {
				xx := (x + grid[i][j]) % k
				dp[i][j][xx] = (dp[i][j-1][x] + dp[i-1][j][x]) % P
			}
		}
	}
	return dp[m-1][n-1][0]
}

func main() {
	for _, v := range []struct {
		grid   [][]int
		k, ans int
	}{
		{[][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}, 3, 2},
		{[][]int{{0, 0}}, 5, 1},
		{[][]int{{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0}}, 1, 10},
	} {
		fmt.Println(numberOfPaths(v.grid, v.k), v.ans)
	}
}

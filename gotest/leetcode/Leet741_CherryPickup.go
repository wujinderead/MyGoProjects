package main

import "fmt"

// https://leetcode.com/problems/cherry-pickup/

// In a N x N grid representing a field of cherries, each cell is one of three possible integers.
// 0 means the cell is empty, so you can pass through; 
// 1 means the cell contains a cherry, that you can pick up and pass through; 
// -1 means the cell contains a thorn that blocks your way.
// Your task is to collect maximum number of cherries possible by following the rules below:
// 1. Starting at the position (0, 0) and reaching (N-1, N-1) by moving right or down
//    through valid path cells (cells with value 0 or 1);
// 2. After reaching (N-1, N-1), returning to (0, 0) by moving left or up through
//    valid path cells;
// 3. When passing through a path cell containing a cherry, you pick it up and the
//    cell becomes an empty cell (0);
// 4. If there is no valid path between (0, 0) and (N-1, N-1), then no cherries can
//    be collected.
// Example 1:
//   Input: grid =
//     [[0, 1, -1],
//      [1, 0, -1],
//      [1, 1,  1]]
//   Output: 5
//   Explanation:
//     The player started at (0, 0) and went down, down, right right to reach (2, 2).
//     4 cherries were picked up during this single trip, and the matrix becomes
//     [[0,1,-1],[0,0,-1],[0,0,0]]. Then, the player went left, up, up, left to return
//     home, picking up one more cherry.
//     The total number of cherries picked up is 5, and this is the maximum possible.
// Note:
//   grid is an N by N 2D array, with 1 <= N <= 50.
//   Each grid[i][j] is an integer in the set {-1, 0, 1}.
//   It is guaranteed that grid[0][0] and grid[N-1][N-1] are not -1.

// https://leetcode.com/problems/cherry-pickup/solution/

// we can see the problem as we go two separated trips: (0,0)->(i,j) and (0,0)->(p,q).
// denote the max cherry of the two trips as T(i,j,p,q). those two trips won't overlap,
// or overlap if only when i==p and j==q (when this happens, we only add one grid[i][j])
// after n steps, we got i+j=p+q=n, we can reduce it to 3 variables, as T(n, i, p).
// T(n, i, p) = grid[i][n-i] + (i==p ? 0 : grid[p][n-p]) +
//     max(T(n-1, i-1, p-1), T(n-1, i-1, p), T(n-1, i, p-1), T(n-1, i, p))
// then we can iterate n from 1 to 2N-1, to get the final answer T(2N-1, N-1, N-1)
func cherryPickup(grid [][]int) int {
	N := len(grid)
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, N)
	}

	dp[0][0] = grid[0][0]
	for n := 1; n < 2*N-1; n++ {
		for i:=N-1; i>=0; i-- {
			for p:=N-1; p>=0; p-- {
				j := n - i
				q := n - p

				if j < 0 || j >= N || q < 0 || q >= N || grid[i][j] < 0 || grid[p][q] < 0 {
					dp[i][p] = -1
					continue
				}

				if i > 0 {
					dp[i][p] = max(dp[i][p], dp[i-1][p])
				}
				if p > 0 {
					dp[i][p] = max(dp[i][p], dp[i][p-1])
				}
				if i > 0 && p > 0 {
					dp[i][p] = max(dp[i][p], dp[i-1][p-1])
				}

				if dp[i][p] >= 0 {
					dp[i][p] += grid[i][j]
					if i != p {
						dp[i][p] +=	grid[p][q]
					}
				}
			}
		}
	}
	return max(dp[N - 1][N - 1], 0)
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(cherryPickup([][]int{
		{0,1,-1},
		{1,0,-1},
		{1,1,1},
	}))

	fmt.Println(cherryPickup([][]int{
		{1,  0, -1, 1},
		{-1, 1, -1, 1},
		{1,  1,  1, 1},
		{1,  1,  1, 1},
	}))

	fmt.Println(cherryPickup([][]int{
		{1,  1,  1, 1},
		{1,  1,  1, 1},
		{1,  1,  1, 1},
		{1,  1,  1, 1},
	}))

	fmt.Println(cherryPickup([][]int{
		{1,1,-1},
		{1,-1,1},
		{-1,1,1},
	}))
	
	fmt.Println(cherryPickup([][]int{
		{1,1,1,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,1,0,0,1},
		{1,0,0,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,1,1,1,1},
	}))

	fmt.Println(cherryPickup([][]int{
		{1,1,1,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,1,0,0,1},
		{0,0,0,1,0,0,0},
		{1,0,0,1,0,0,0},
		{0,0,0,1,0,0,0},
		{0,0,0,1,1,1,1},
	}))
}
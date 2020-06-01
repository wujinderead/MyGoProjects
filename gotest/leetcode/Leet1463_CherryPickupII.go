package main

import "fmt"

// https://leetcode.com/problems/cherry-pickup-ii/

// Given a rows x cols matrix grid representing a field of cherries. Each cell in grid represents the number 
// of cherries that you can collect. You have two robots that can collect cherries for you, Robot #1 is 
// located at the top-left corner (0,0) , and Robot #2 is located at the top-right corner (0,cols-1) of the grid. 
// Return the maximum number of cherries collection using both robots by following the rules below: 
//   From a cell (i,j), robots can move to cell (i+1, j-1) , (i+1, j) or (i+1, j+1). 
//   When any robot is passing through a cell, It picks it up all cherries, and the cell becomes an empty cell (0). 
//   When both robots stay on the same cell, only one of them takes the cherries. 
//   Both robots cannot move outside of the grid at any moment. 
//   Both robots should reach the bottom row in the grid. 
// Example 1: 
//          (3)  2  [1]
//          (2) [5]  1
//           1  (5) [5]
//          (2)  1  [1]
//   Input: grid = [[3,1,1],[2,5,1],[1,5,5],[2,1,1]]
//   Output: 24
//   Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
//     Cherries taken by Robot #1, (3 + 2 + 5 + 2) = 12.
//     Cherries taken by Robot #2, (1 + 5 + 5 + 1) = 12.
//     Total of cherries: 12 + 12 = 24.
// Example 2: 
//   Input: grid = [[1,0,0,0,0,0,1],[2,0,0,0,0,3,0],[2,0,9,0,0,0,0],[0,3,0,5,4,0,0],[1,0,2,3,0,0,6]]
//   Output: 28
//   Explanation: Path of robot #1 and #2 are described in color green and blue respectively.
//     Cherries taken by Robot #1, (1 + 9 + 5 + 2) = 17.
//     Cherries taken by Robot #2, (1 + 3 + 4 + 3) = 11.
//     Total of cherries: 17 + 11 = 28.
// Example 3:  
//   Input: grid = [[1,0,0,3],[0,0,0,3],[0,0,3,3],[9,0,3,3]]
//   Output: 22
// Example 4:  
//   Input: grid = [[1,1],[1,1]]
//   Output: 4
// Constraints:  
//   rows == grid.length 
//   cols == grid[i].length 
//   2 <= rows, cols <= 70 
//   0 <= grid[i][j] <= 100 

func cherryPickup(grid [][]int) int {
	// the state of a line is only related to the previous line. each line has n*n states. dp those states.
	m, n := len(grid), len(grid[0])
	old, new := make([]int, n*n), make([]int, n*n)
	for i := range old {
		old[i] = -0xffffff               // set negative to indicate we haven't visit it 
	}
	old[n-1] = grid[0][0]+grid[0][n-1]   // line 0, col 0 + col n-1

	// dp
	for k:=1; k<m; k++ {
		for s:=0; s<n*n; s++ {
			new[s] = -0xffffff
			c1, c2 := s/n, s%n   // robot1 at grid[k][c1], robot2 at grid[k][c2]
			toadd := grid[k][c1]
			if c1 != c2 {
				toadd += grid[k][c2]
			}
			// it relate to at most 9 states in line k-1
			for _, d1 := range []int{1,0,-1} {
				for _, d2 := range []int{1,0,-1} {
					nc1, nc2 := c1+d1, c2+d2
					if nc1>=0 && nc1<n && nc2>=0 && nc2<n && old[nc1*n+nc2]>=0 {
						new[s] = max(new[s], old[nc1*n+nc2]+toadd)
					}
				}
			}
		}
		old, new = new, old
	}
	allmax := 0
	for _, v := range old {
		if v>allmax {
			allmax = v
		}
	}
    return allmax
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(cherryPickup([][]int{{3,1,1},{2,5,1},{1,5,5},{2,1,1}}), 24)
	fmt.Println(cherryPickup([][]int{{1,0,0,0,0,0,1},{2,0,0,0,0,3,0},{2,0,9,0,0,0,0},{0,3,0,5,4,0,0},{1,0,2,3,0,0,6}}), 28)
	fmt.Println(cherryPickup([][]int{{1,0,0,3},{0,0,0,3},{0,0,3,3},{9,0,3,3}}), 22)
	fmt.Println(cherryPickup([][]int{{1,1},{1,1}}), 4)
	fmt.Println(cherryPickup([][]int{
		{0,8,7,10,9,10,0,9,6},
		{8,7,10,8,7,4,9,6,10},
		{8,1,1,5,1,5,5,1,2},
		{9,4,10,8,8,1,9,5,0},
		{4,3,6,10,9,2,4,8,10},
		{7,3,2,8,3,3,5,9,8},
		{1,2,6,5,6,2,0,10,0},
	}), 96)
	fmt.Println(cherryPickup([][]int{
		{0,0,10,2,8,4,0},
		{7,9,3,5,4,8,3},
		{6,9,8,3,5,6,0},
		{0,4,1,1,9,3,7},
		{5,6,9,8,8,10,10},
		{9,2,9,7,4,8,3},
		{1,6,1,2,0,9,9},
	}), 96)
}
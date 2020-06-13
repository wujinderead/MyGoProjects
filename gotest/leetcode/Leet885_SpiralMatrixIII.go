package main

import (
    "fmt"
)

// https://leetcode.com/problems/spiral-matrix-iii/

// On a 2 dimensional grid with R rows and C columns, we start at (r0, c0) facing east. 
// Here, the north-west corner of the grid is at the first row and column, and the south-east corner of 
// the grid is at the last row and column. Now, we walk in a clockwise spiral shape to visit every position 
// in this grid. Whenever we would move outside the boundary of the grid, we continue our walk outside the grid 
// (but may return to the grid boundary later.) Eventually, we reach all R * C spaces of the grid. 
// Return a list of coordinates representing the positions of the grid in the order they were visited. 
// Example 1: 
//   Input: R = 1, C = 4, r0 = 0, c0 = 0
//   Output: [[0,0],[0,1],[0,2],[0,3]]
// Example 2: 
//   Input: R = 5, C = 6, r0 = 1, c0 = 4
//   Output: [[1,4],[1,5],[2,5],[2,4],[2,3],[1,3],[0,3],[0,4],[0,5],[3,5],[3,4],[3,3],[3,2],[2,2],[1,2],
//     [0,2],[4,5],[4,4],[4,3],[4,2],[4,1],[3,1],[2,1],[1,1],[0,1],[4,0],[3,0],[2,0],[1,0],[0,0]]
// Note:  
//   1 <= R <= 100 
//   1 <= C <= 100 
//   0 <= r0 < R 
//   0 <= c0 < C 

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
    // to right, to down, to left, to up; two moves 1 step, two moves 2 steps, two moves 3 steps. 
    // until all nodes are visited
    count := 1
    step := 2
    ans := make([][]int, 0, R*C)
    ans = append(ans, []int{r0, c0})
    if R*C==1 {
    	return ans
    }
    i, j := r0, c0
    for {
    	for _, v := range [][2]int{{0,1}, {1,0}, {0,-1}, {-1,0}} {
    		ni, nj := i+v[0]*(step/2), j+v[1]*(step/2)    // every 2 steps we change direction
    		for s:=1; s<=step/2; s++ {
    			ii, jj := i+v[0]*s, j+v[1]*s
    			if ii>=0 && ii<R && jj>=0 && jj<C {
    				ans = append(ans, []int{ii, jj})
    				count++
    				if count==R*C {
    					return ans
    				}
    			}
    		}
    		step++
    		i, j = ni, nj
    	}
	}
    return ans
}

func main() {
	fmt.Println(spiralMatrixIII(1,4,0,0))
	fmt.Println(spiralMatrixIII(5,6,1,4))
	fmt.Println(spiralMatrixIII(1,1,0,0))
	fmt.Println(spiralMatrixIII(1,2,0,0))
	fmt.Println(spiralMatrixIII(1,2,0,1))
}
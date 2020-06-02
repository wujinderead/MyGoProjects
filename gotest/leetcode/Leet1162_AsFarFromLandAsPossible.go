package main

import "fmt"

// https://leetcode.com/problems/as-far-from-land-as-possible/

// Given an N x N grid containing only values 0 and 1, where 0 represents water and 1 represents land, 
// find a water cell such that its distance to the nearest land cell is maximized and return the distance. 
// The distance used in this problem is the Manhattan distance: the distance between two cells 
// (x0, y0) and (x1, y1) is |x0 - x1| + |y0 - y1|. 
// If no land or water exists in the grid, return -1. 
// Example 1:  
//   Input: [[1,0,1],[0,0,0],[1,0,1]]
//   Output: 2
//   Explanation: 
//     The cell (1, 1) is as far as possible from all the land with distance 2.
// Example 2:  
//   Input: [[1,0,0],[0,0,0],[0,0,0]]
//   Output: 4
//   Explanation: 
//     The cell (2, 2) is as far as possible from all the land with distance 4.
// Note: 
//   1 <= grid.length == grid[0].length <= 100 
//   grid[i][j] is 0 or 1 

func maxDistance(grid [][]int) int {
	n := len(grid)
    cur, next := make([][2]int, 0, 10), make([][2]int, 0, 10)
    for i := range grid {
    	for j := range grid[0] {
    		if grid[i][j] == 1 {
    			cur = append(cur, [2]int{i, j})   // add all 1 to cur queue
    		}
    	}
    }
    if len(cur)==0 || len(cur)==n*n {
    	return -1
    }
    max := 0
    for len(cur)>0 {
    	max++
    	for _, v := range cur {   // add cur's neighbor to next queue
    		i, j := v[0], v[1]
    		for _, vv := range [][2]int{{1,0}, {-1,0}, {0,1}, {0,-1}} {
    			ni, nj := i+vv[0], j+vv[1]
    			if ni>=0 && ni<n && nj>=0 && nj<n && grid[ni][nj]==0 {
    				grid[ni][nj] = max
    				next = append(next, [2]int{ni, nj}) 
    			}
    		}
    	}
    	cur, next = next, cur[:0]
    }
    return max-1
}

func main() {
	fmt.Println(maxDistance([][]int{{1,0,1},{0,0,0},{1,0,1}}))
	fmt.Println(maxDistance([][]int{{1,0,0},{0,0,0},{0,0,0}}))
	fmt.Println(maxDistance([][]int{{1,1,1},{1,1,1},{1,1,1}}))
	fmt.Println(maxDistance([][]int{{1,1,1},{1,1,1},{1,1,0}}))	
}
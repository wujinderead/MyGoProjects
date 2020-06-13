package main

import (
	"container/list"
    "fmt"
)

// https://leetcode.com/problems/making-a-large-island/

// In a 2D grid of 0s and 1s, we change at most one 0 to a 1. 
// After, what is the size of the largest island? (An island is a 4-directionally connected group of 1s). 
// Example 1: 
//   Input: [[1, 0], [0, 1]]
//   Output: 3
//   Explanation: Change one 0 to 1 and connect two 1s, then we get an island with area = 3.
// Example 2: 
//   Input: [[1, 1], [1, 0]]
//   Output: 4
//   Explanation: Change the 0 to 1 and make the island bigger, only one island with area = 4. 
// Example 3: 
//   Input: [[1, 1], [1, 1]]
//   Output: 4
//   Explanation: Can't change any 0 to 1, only one island with area = 4. 
// Notes: 
//   1 <= grid.length = grid[0].length <= 50. 
//   0 <= grid[i][j] <= 1. 

func largestIsland(grid [][]int) int {
    label := [50][50]int{}
    m, n := len(grid), len(grid[0])
    L := 1
    issize := []int{-1}
    queue := list.New()
    maxisland := 0

    // label each single island
    for i:=0; i<m; i++ {
    	for j:=0; j<n; j++ {
    		if grid[i][j]==1 && label[i][j]==0 {
    			queue.PushBack([2]int{i, j})
    			label[i][j] = L
    			size := 0
    			for queue.Len()>0 {
    				tmp := queue.Remove(queue.Front()).([2]int)
    				size++
    				for _, v := range[][2]int{{1,0}, {-1,0}, {0,1}, {0,-1}} {
    					ni, nj := tmp[0]+v[0], tmp[1]+v[1]
    					if ni>=0 && ni<m && nj>=0 && nj<n && grid[ni][nj]==1 && label[ni][nj]==0 {
    						label[ni][nj] = L
    						queue.PushBack([2]int{ni, nj})
    					}
    				}
    			}
    			maxisland = max(maxisland, size)
    			issize = append(issize, size)
    			L++
    		}
    	}
    }

    // try make a larger island
    cand := make([]int, 0, 4)
    for i:=0; i<m; i++ {
    	for j:=0; j<n; j++ {
    		if grid[i][j]==0 {
    			outer: for _, v := range[][2]int{{1,0}, {-1,0}, {0,1}, {0,-1}} {
    				ni, nj := i+v[0], j+v[1]
    				if ni>=0 && ni<m && nj>=0 && nj<n && label[ni][nj]>0 {
    					for k := range cand {
    						if cand[k] == label[ni][nj] {   // for a 0's neighbor, find all distinct island
    							continue outer
    						}
    					}
    					cand = append(cand, label[ni][nj])
    				}
    			}
    			sum := 1
    			for k := range cand {
    				sum += issize[cand[k]]  // sum all distinct islands' size
    			}
    			maxisland = max(maxisland, sum)
    			cand = cand[:0]
    		}
    	}
    }
    return maxisland
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(largestIsland([][]int{
		{1,0},
		{0,1},
	}), 3)
	fmt.Println(largestIsland([][]int{
		{1,0},
		{1,1},
	}), 4)
	fmt.Println(largestIsland([][]int{
		{1,1},
		{1,1},
	}), 4)
	fmt.Println(largestIsland([][]int{
		{0,1,0},
		{1,0,1},
		{0,1,1},
	}), 6)
	fmt.Println(largestIsland([][]int{
		{0,0,0},
		{0,0,0},
		{0,0,0},
	}), 1)
}
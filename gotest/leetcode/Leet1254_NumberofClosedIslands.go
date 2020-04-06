package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/number-of-closed-islands

// Given a 2D grid consists of 0s (land) and 1s (water). An island is a maximal
// 4-directionally connected group of 0s and a closed island is an island totally
// (all left, top, right, bottom) surrounded by 1s.
// Return the number of closed islands.
// Example 1:
//   Input: grid = [[1,1,1,1,1,1,1,0],[1,0,0,0,0,1,1,0],[1,0,1,0,1,1,1,0],[1,0,0,0,
//     0,1,0,1],[1,1,1,1,1,1,1,0]]
//   Output: 2
//   Explanation:
//     Islands in gray are closed because they are completely surrounded by water
//     (group of 1s).
// Example 2:
//   Input: grid = [[0,0,1,0,0],[0,1,0,1,0],[0,1,1,1,0]]
//   Output: 1
// Example 3:
//   Input: grid = [[1,1,1,1,1,1,1],
//               [1,0,0,0,0,0,1],
//               [1,0,1,1,1,0,1],
//               [1,0,1,0,1,0,1],
//               [1,0,1,1,1,0,1],
//               [1,0,0,0,0,0,1],
//               [1,1,1,1,1,1,1]]
//   Output: 2
// Constraints:
//   1 <= grid.length, grid[0].length <= 100
//   0 <= grid[i][j] <=1

func closedIsland(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	count := 0
	queue := list.New()
	dir := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}}
	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if grid[i][j]==0 && !visited[i][j] {
				isIsle := true
				visited[i][j] = true
				queue.PushBack([2]int{i, j})
				for queue.Len()>0 {
					pair := queue.Remove(queue.Front()).([2]int)
					r, c := pair[0], pair[1]
					if r==0 || r==len(grid)-1 || c==0 || c==len(grid[0])-1 {
						isIsle = false  // if reach the border, it's not closed island
					}
					for _, v := range dir {
						newi, newj := r+v[0], c+v[1]
						if newi>=0 && newi<len(grid) && newj>=0 && newj<len(grid[0]) &&
							!visited[newi][newj] && grid[newi][newj]==0 {
							queue.PushBack([2]int{newi, newj})
							visited[newi][newj] = true
						}
					}
				}
				if isIsle {
					count++
				}
			}
		}
	}
	return count
}

func main() {
    fmt.Println(closedIsland([][]int{
		{1,1,1,1,1,1,1,0},
		{1,0,0,0,0,1,1,0},
		{1,0,1,0,1,1,1,0},
		{1,0,0,0,0,1,0,1},
		{1,1,1,1,1,1,1,0},
	}))
	fmt.Println(closedIsland([][]int{
		{0,0,1,0,0},
		{0,1,0,1,0},
		{0,1,1,1,0},
	}))
	fmt.Println(closedIsland([][]int{
		{1,1,1,1,1,1,1},
		{1,0,0,0,0,0,1},
		{1,0,1,1,1,0,1},
		{1,0,1,0,1,0,1},
		{1,0,1,1,1,0,1},
		{1,0,0,0,0,0,1},
		{1,1,1,1,1,1,1},
	}))
}
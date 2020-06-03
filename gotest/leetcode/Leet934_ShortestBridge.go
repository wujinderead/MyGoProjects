package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/shortest-bridge/

// In a given 2D binary array A, there are two islands. (An island is a 4-directionally 
// connected group of 1s not connected to any other 1s.) 
// Now, we may change 0s to 1s so as to connect the two islands together to form 1 island. 
// Return the smallest number of 0s that must be flipped. (It is guaranteed that the answer is at least 1.) 
// Example 1: 
//   Input: A = [[0,1],[1,0]]
//   Output: 1
// Example 2: 
//   Input: A = [[0,1,0],[0,0,0],[0,0,1]]
//   Output: 2
// Example 3: 
//   Input: A = [[1,1,1,1,1],[1,0,0,0,1],[1,0,1,0,1],[1,0,0,0,1],[1,1,1,1,1]]
//   Output: 1
// Constraints: 
//   2 <= A.length == A[0].length <= 100 
//   A[i][j] == 0 or A[i][j] == 1 

func shortestBridge(A [][]int) int {
	n := len(A)
	queue := list.New() 
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var fi, fj int
	outer: for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			if A[i][j] == 1 {
				fi, fj = i, j
				break outer
			}
		}
	}
	visit(A, fi, fj, visited, queue)   // dfs to get all island

	// use bfs to get the minimal distance
	var i, j, dist, v int
	for queue.Len()>0 {
		v = queue.Remove(queue.Front()).(int)
		dist = v>>20
		i = (v-(dist<<20)) >> 10
		j = v-(dist<<20)-(i<<10)
		for _, vv := range [][2]int{{0,1}, {0,-1}, {-1,0}, {1,0}} {
			ni, nj := i+vv[0], j+vv[1]
			if ni>=0 && ni<len(A) && nj>=0 && nj<len(A) && !visited[ni][nj] {
				visited[ni][nj] = true
				if A[ni][nj] == 1 {
					return dist
				}
				queue.PushBack((ni<<10) + nj + (dist+1)<<20)
			}
		}
	}
	return -1 // unreachable
}

func visit(A [][]int, i, j int, visited [][]bool, queue *list.List) {
	visited[i][j] = true
	queue.PushBack((i<<10)+j)
	for _, v := range [][2]int{{0,1}, {0,-1}, {-1,0}, {1,0}} {
		ni, nj := i+v[0], j+v[1]
		if ni>=0 && ni<len(A) && nj>=0 && nj<len(A) && A[ni][nj]==1 && !visited[ni][nj] {
			visit(A, ni, nj, visited, queue)
		}
	}
}

func main() {
	fmt.Println(shortestBridge([][]int{{0,1}, {1,0}}))
	fmt.Println(shortestBridge([][]int{{0,1,0},{0,0,0},{0,0,1}}))
	fmt.Println(shortestBridge([][]int{{1,1,1,1,1},{1,0,0,0,1},{1,0,1,0,1},{1,0,0,0,1},{1,1,1,1,1}}))
	fmt.Println(shortestBridge([][]int{
		{1,1,0,0,0},
		{1,0,0,0,0},
		{1,0,0,0,1},
		{1,0,0,0,0},
		{1,0,0,0,0},
	}))
}
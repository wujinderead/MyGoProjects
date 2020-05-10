package main

import "fmt"

// https://leetcode.com/problems/minimum-time-to-collect-all-apples-in-a-tree/

// Given an undirected tree consisting of n vertices numbered from 0 to n-1, which
// has some apples in their vertices. You spend 1 second to walk over one edge of
// the tree. Return the minimum time in seconds you have to spend in order to collect
// all apples in the tree starting at vertex 0 and coming back to this vertex.
// The edges of the undirected tree are given in the array edges, where edges[i]=[fromi, toi]
// means that exists an edge connecting the vertices fromi and toi.
// Additionally, there is a boolean array hasApple, where hasApple[i] = true means
// that vertex i has an apple, otherwise, it does not have any apple.
// Example 1:
//   Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]],
//     hasApple = [false,false,true,false,true,true,false]
//           0
//         /   \
//        1    2A
//       / \   / \
//     4A  5A 3  6
//   Output: 8
//   Explanation: The figure above represents the given tree where red vertices have
//     an apple. One optimal path to collect all apples is shown by the green arrows.
// Example 2:
//   Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]],
//     hasApple = [false,false,true,false,false,true,false]
//           0
//         /   \
//        1    2A
//       / \   / \
//      4  5A 3  6
//   Output: 6
//   Explanation: The figure above represents the given tree where red vertices have
//     an apple. One optimal path to collect all apples is shown by the green arrows.
// Example 3:
//   Input: n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]],
//     hasApple = [false,false,false,false,false,false,false]
//   Output: 0
// Constraints:
//   1 <= n <= 10^5
//   edges.length == n-1
//   edges[i].length == 2
//   0 <= fromi, toi <= n-1
//   fromi < toi
//   hasApple.length == n

func minTime(n int, edges [][]int, hasApple []bool) int {
    graph := make([][]int, n)
    for i := range edges {
    	graph[edges[i][0]] = append(graph[edges[i][0]], edges[i][1])
    	graph[edges[i][1]] = append(graph[edges[i][1]], edges[i][0])
	}
	time := 0
	visited := make([]bool, n)
	visited[0] = true
	visit(0, graph, hasApple, visited, &time)
	return time
}

func visit(i int, graph [][]int, hasApple, visited []bool, time *int) bool {
	has := hasApple[i]
	for _, j := range graph[i] {
		if !visited[j] {
			visited[j] = true
			*time++
			childhas := visit(j, graph, hasApple, visited, time)
			if childhas {
				*time++
			} else {
				*time--
			}
			has = has ||  childhas
		}
	}
	return has
}

func main() {
	fmt.Println(minTime(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}},
		[]bool{false,false,true,false,true,true,false}), 8)
	fmt.Println(minTime(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}},
		[]bool{false,false,true,false,false,true,false}), 6)
	fmt.Println(minTime(7, [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}},
		[]bool{false,false,false,false,false,false,false}), 0)
}
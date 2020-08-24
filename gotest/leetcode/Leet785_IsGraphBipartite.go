package main

import "fmt"

// https://leetcode.com/problems/is-graph-bipartite/

// Given an undirected graph, return true if and only if it is bipartite.
// Recall that a graph is bipartite if we can split it's set of nodes into two independent subsets 
// A and B such that every edge in the graph has one node in A and another node in B.
// The graph is given in the following form: graph[i] is a list of indexes j for which the edge 
// between nodes i and j exists.  Each node is an integer between 0 and graph.length - 1.  
// There are no self edges or parallel edges: graph[i] does not contain i, and it doesn't contain 
// any element twice.
// Example 1:
//   Input: [[1,3], [0,2], [1,3], [0,2]]
//   Output: true
//   Explanation: 
//     The graph looks like this:
//       0----1
//       |    |
//       |    |
//       3----2
//     We can divide the vertices into two groups: {0, 2} and {1, 3}.
// Example 2:
//   Input: [[1,2,3], [0,2], [0,1,3], [0,2]]
//   Output: false
//   Explanation: 
//     The graph looks like this:
//       0----1
//       | \  |
//       |  \ |
//       3----2
//     We cannot find a way to divide the set of nodes into two independent subsets.
// Note:
//   graph will have length in range [1, 100].
//   graph[i] will contain integers in range [0, graph.length - 1].
//   graph[i] will not contain i or duplicate values.
//   The graph is undirected: if any element j is in graph[i], then i will be in graph[j].

func isBipartite(graph [][]int) bool {
	part := make([]int, len(graph))
	for i := range graph {
		if part[i] == 0 {
			part[i] = 1   // color unvisited node 1
			if !visit(graph, part, i) {
				return false
			}
		}
	}
	return true
}

func visit(graph [][]int, part []int, i int) bool {
	for _, v := range graph[i] {
		if part[v] == 0 {        // if node's neighbor is unvisited
			part[v] = -part[i]   // color it with reverse color
			if !visit(graph, part, v) {
				return false
			}
		} else if part[v] == part[i] {  // if two nodes of an edge are in the same set, return false 
			return false
		}
	}
	return true
}

func main() {
	for _, v := range []struct{graph [][]int; ans bool} {
		{[][]int{{1,3}, {0,2}, {1,3}, {0,2}}, true},
		{[][]int{{1,2,3}, {0,2}, {0,1,3}, {0,2}}, false},
	} {
		fmt.Println(isBipartite(v.graph), v.ans)
	}
}

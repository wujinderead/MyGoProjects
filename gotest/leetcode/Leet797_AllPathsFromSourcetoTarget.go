package main

import "fmt"

// https://leetcode.com/problems/all-paths-from-source-to-target/

// Given a directed, acyclic graph of N nodes. Find all possible paths from
// node 0 to node N-1, and return them in any order.
// The graph is given as follows: the nodes are 0, 1, ..., graph.length - 1.
// graph[i] is a list of all nodes j for which the edge (i, j) exists.
//   Example:
//   Input: [[1,2], [3], [3], []]
//   Output: [[0,1,3],[0,2,3]]
//   Explanation:
//     The graph looks like this:
//       0--->1
//       |    |
//       v    v
//       2--->3
//     There are two paths: 0 -> 1 -> 3 and 0 -> 2 -> 3.
// Note:
//   The number of nodes in the graph will be in the range [2, 15].
//   You can print different paths in any order, but you should keep the order of nodes inside one path.

func allPathsSourceTarget(graph [][]int) [][]int {
	paths := make([][]int, 0)
	visited := make([]bool, len(graph))
	path := make([]int, len(graph))
	visit(0, 0, graph, path, visited, &paths)
	return paths
}

func visit(curind, pathind int, graph [][]int, path []int, visited []bool, paths *[][]int) {
	path[pathind] = curind
	if curind == len(graph)-1 {
		tmp := make([]int, pathind+1)
		copy(tmp, path[:pathind+1])
		*paths = append(*paths, tmp)
		return
	}
	for _, nextind := range graph[curind] {
		if !visited[nextind] {
			visited[nextind] = true
			visit(nextind, pathind+1, graph, path, visited, paths)
			visited[nextind] = false
		}
	}
}

func main() {
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
	fmt.Println(allPathsSourceTarget([][]int{{1, 2}, {4}, {1, 3}, {4, 6}, {5}, {6}, {7}, {}}))
}

package main

import "fmt"

// https://leetcode.com/problems/reachable-nodes-with-restrictions/

// There is an undirected tree with n nodes labeled from 0 to n - 1 and n - 1 edges.
// You are given a 2D integer array edges of length n - 1 where edges[i] = [ai, bi] indicates that
// there is an edge between nodes ai and bi in the tree. You are also given an integer array
// restricted which represents restricted nodes.
// Return the maximum number of nodes you can reach from node 0 without visiting a restricted node.
// Note that node 0 will not be a restricted node.
// Example 1:
//   Input: n = 7, edges = [[0,1],[1,2],[3,1],[4,0],[0,5],[5,6]], restricted = [4,5]
//   Output: 4
//   Explanation: The diagram above shows the tree.
//     We have that [0,1,2,3] are the only nodes that can be reached from node 0
//     without visiting a restricted node.
// Example 2:
//   Input: n = 7, edges = [[0,1],[0,2],[0,5],[0,4],[3,2],[6,5]], restricted = [4,2,1]
//   Output: 3
//   Explanation: The diagram above shows the tree.
//     We have that [0,5,6] are the only nodes that can be reached from node 0
//     without visiting a restricted node.
//  Constraints:
//    2 <= n <= 10⁵
//    edges.length == n - 1
//    edges[i].length == 2
//    0 <= ai, bi < n
//    ai != bi
//    edges represents a valid tree.
//    1 <= restricted.length < n
//    1 <= restricted[i] < n
//    All the values of restricted are unique.

// just DFS
func reachableNodes(n int, edges [][]int, restricted []int) int {
	visited := make([]bool, n)
	graph := make([][]int, n)
	count := 0
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	for _, r := range restricted {
		visited[r] = true // mark restricted as visited
	}
	// dfs
	visit(0, graph, visited, &count)
	return count
}

func visit(ind int, graph [][]int, visited []bool, count *int) {
	visited[ind] = true
	*count++
	for _, v := range graph[ind] {
		if !visited[v] {
			visit(v, graph, visited, count)
		}
	}
}

func main() {
	for _, v := range []struct {
		n          int
		edges      [][]int
		restricted []int
		ans        int
	}{
		{7, [][]int{{0, 1}, {1, 2}, {3, 1}, {4, 0}, {0, 5}, {5, 6}}, []int{4, 5}, 4},
		{7, [][]int{{0, 1}, {0, 2}, {0, 5}, {0, 4}, {3, 2}, {6, 5}}, []int{4, 2, 1}, 3},
	} {
		fmt.Println(reachableNodes(v.n, v.edges, v.restricted), v.ans)
	}
}

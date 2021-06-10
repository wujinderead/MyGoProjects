package main

import "fmt"

// https://leetcode.com/problems/largest-color-value-in-a-directed-graph/

// There is a directed graph of n colored nodes and m edges. The nodes are numbered from 0 to n - 1.
// You are given a string colors where colors[i] is a lowercase English letter representing
// the color of the ith node in this graph (0-indexed). You are also given a 2D array edges
// where edges[j] = [aj, bj] indicates that there is a directed edge from node aj to node bj.
// A valid path in the graph is a sequence of nodes x1 -> x2 -> x3 -> ... -> xk such that
// there is a directed edge from xi to xi+1 for every 1 <= i < k. The color value of the path
// is the number of nodes that are colored the most frequently occurring color along that path.
// Return the largest color value of any valid path in the given graph, or -1 if the graph
// contains a cycle.
// Example 1:
//   Input: colors = "abaca", edges = [[0,1],[0,2],[2,3],[3,4]]
//   Output: 3
//           0a -> 2a -> 3c -> 4a
//           |
//           1b
//   Explanation: The path 0 -> 2 -> 3 -> 4 contains 3 nodes that are colored "a" (red in the above image).
// Example 2:
//   Input: colors = "a", edges = [[0,0]]
//   Output: -1
//   Explanation: There is a cycle from 0 to 0.
// Constraints:
//   n == colors.length
//   m == edges.length
//   1 <= n <= 10^5
//   0 <= m <= 10^5
//   colors consists of lowercase English letters.
//   0 <= aj, bj < n

// use topological sort
// let dp[v][c] be the max-frequency of color c to the paths end at vertex v
func largestPathValue(colors string, edges [][]int) int {
	graph := make([][]int, len(colors))
	dp := make([][26]int, len(colors))
	degree := make([]int, len(colors))

	// make graph
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		degree[edge[1]]++
	}
	max := 0

	// topological sort
	stack := make([]int, 0)
	for i := range degree {
		if degree[i] == 0 {
			stack = append(stack, i)
		}
	}
	visited := 0
	for len(stack) > 0 {
		// pop a vertex from queue
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited++
		dp[cur][colors[cur]-'a']++
		if dp[cur][colors[cur]-'a'] > max {
			max = dp[cur][colors[cur]-'a']
		}
		// for each edge going out of cur
		for _, to := range graph[cur] {
			// topological sort
			degree[to]--
			if degree[to] == 0 {
				stack = append(stack, to)
			}
			// dp
			for c := 0; c < 26; c++ {
				if dp[cur][c] > dp[to][c] {
					dp[to][c] = dp[cur][c]
				}
			}
		}
	}
	if visited < len(colors) { // has circle, can't traverse all vertices
		return -1
	}
	return max
}

func main() {
	for _, v := range []struct {
		c   string
		e   [][]int
		ans int
	}{
		{"abaca", [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}, 3},
		{"a", [][]int{{0, 0}}, -1},
	} {
		fmt.Println(largestPathValue(v.c, v.e), v.ans)
	}
}

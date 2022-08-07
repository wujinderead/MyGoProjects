package main

import "fmt"

// https://leetcode.com/problems/longest-cycle-in-a-graph/

// You are given a directed graph of n nodes numbered from 0 to n - 1, where each node has
// at most one outgoing edge.
// The graph is represented with a given 0-indexed array edges of size n, indicating that there is
// a directed edge from node i to node edges[i]. If there is no outgoing edge from node i, then
// edges[i] == -1.
// Return the length of the longest cycle in the graph. If no cycle exists, return -1.
// A cycle is a path that starts and ends at the same node.
// Example 1:
//   Input: edges = [3,3,4,2,3]
//   Output: 3
//   Explanation: The longest cycle in the graph is the cycle: 2 -> 4 -> 3 -> 2.
//     The length of this cycle is 3, so 3 is returned.
// Example 2:
//   Input: edges = [2,-1,3,1]
//   Output: -1
//   Explanation: There are no cycles in this graph.
// Constraints:
//   n == edges.length
//   2 <= n <= 10⁵
//   -1 <= edges[i] < n
//   edges[i] != i

// just dfs, there can't be two cycle in a single connected components
func longestCycle(edges []int) int {
	n := len(edges)
	dist := make([]int, n)
	group := make([]int, n)
	for i := range group {
		group[i] = -1
	}

	ans := -1
	for i := range edges {
		if group[i] != -1 {
			continue
		}
		// for unvisited node
		curg := i
		group[i] = curg
		dist[i] = 1
		curi := i
		for edges[curi] != -1 {
			next := edges[curi]
			if group[next] == -1 { // next unvisited
				dist[next] = dist[curi] + 1
				curi = next
				group[curi] = curg
			} else if group[next] == curg { // found a cycle
				cycle := dist[curi] - dist[next] + 1
				if cycle > ans {
					ans = cycle
				}
				break
			} else { // group[next] < curg, reach a visited component
				break
			}
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		edges []int
		ans   int
	}{
		{[]int{3, 3, 4, 2, 3}, 3},
		{[]int{2, -1, 3, 1}, -1},
	} {
		fmt.Println(longestCycle(v.edges), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/find-closest-node-to-given-two-nodes/

// You are given a directed graph of n nodes numbered from 0 to n - 1, where each node has at most
// one outgoing edge.
// The graph is represented with a given 0-indexed array edges of size n, indicating that there is
// a directed edge from node i to node edges[i]. If there is no outgoing edge from i, then edges[i]
// == -1.
// You are also given two integers node1 and node2.
// Return the index of the node that can be reached from both node1 and node2, such that the maximum
// between the distance from node1 to that node, and from node2 to that node is minimized. If there
// are multiple answers, return the node with the smallest index, and if no possible answer exists,
// return -1.
// Note that edges may contain cycles.
// Example 1:
//   Input: edges = [2,2,3,-1], node1 = 0, node2 = 1
//   Output: 2
//   Explanation: The distance from node 0 to node 2 is 1, and the distance from node 1 to node 2 is 1.
//     The maximum of those two distances is 1. It can be proven that we cannot get a
//     node with a smaller maximum distance than 1, so we return node 2.
// Example 2:
//   Input: edges = [1,2,-1], node1 = 0, node2 = 2
//   Output: 2
//   Explanation: The distance from node 0 to node 2 is 2, and the distance from node 2 to itself is 0.
//     The maximum of those two distances is 2. It can be proven that we cannot get a
//     node with a smaller maximum distance than 2, so we return node 2.
// Constraints:
//   n == edges.length
//   2 <= n <= 10⁵
//   -1 <= edges[i] < n
//   edges[i] != i
//   0 <= node1, node2 < n

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	dist1 := make([]int, n)
	dist2 := make([]int, n)
	visited1 := make([]bool, n)
	visited2 := make([]bool, n)

	// dfs node1
	i := node1
	visited1[i] = true
	dist1[i] = 1
	for edges[i] != -1 && !visited1[edges[i]] {
		dist1[edges[i]] = dist1[i] + 1
		i = edges[i]
		visited1[i] = true
	}

	// dfs node2
	i = node2
	visited2[i] = true
	dist2[i] = 1
	for edges[i] != -1 && !visited2[edges[i]] {
		dist2[edges[i]] = dist2[i] + 1
		i = edges[i]
		visited2[i] = true
	}
	fmt.Println(dist1)
	fmt.Println(dist2)

	// find answer
	ans := -1
	min := 100000
	for i := range edges {
		if dist1[i] > 0 && dist2[i] > 0 { // both reachable
			max := dist1[i]
			if dist2[i] > max { // get max of two distances
				max = dist2[i]
			}
			if max < min { // update minimal
				min = max
				ans = i
			}
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		edges       []int
		n1, n2, ans int
	}{
		{[]int{2, 2, 3, -1}, 0, 1, 2},
		{[]int{1, 2, -1}, 0, 2, 2},
		{[]int{5, 3, 1, 0, 2, 4, 5}, 3, 2, 3},
	} {
		fmt.Println(closestMeetingNode(v.edges, v.n1, v.n2), v.ans)
	}

}

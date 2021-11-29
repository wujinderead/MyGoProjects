package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/reachable-nodes-in-subdivided-graph/

// You are given an undirected graph (the "original graph") with n nodes labeled from 0 to n - 1.
// You decide to subdivide each edge in the graph into a chain of nodes, with the number of new nodes
// varying between each edge.
// The graph is given as a 2D array of edges where edges[i] = [ui, vi, cnti] indicates that there is an
// edge between nodes ui and vi in the original graph, and cnti is the total number of new nodes that
// you will subdivide the edge into. Note that cnti == 0 means you will not subdivide the edge.
// To subdivide the edge [ui, vi], replace it with (cnti + 1) new edges and cnti new nodes. The new nodes
// are x1, x2, ..., xcnti, and the new edges are [ui, x1], [x1, x2], [x2, x3], ..., [xcnti-1, xcnti], [xcnti, vi].
// In this new graph, you want to know how many nodes are reachable from the node 0, where a node is
// reachable if the distance is maxMoves or less.
// Given the original graph and maxMoves, return the number of nodes that are reachable from node 0 in the new graph.
// Example 1:
//   Input: edges = [[0,1,10],[0,2,1],[1,2,2]], maxMoves = 6, n = 3
//   Output: 13
//   Explanation: The edge subdivisions are shown in the image above.
//     The nodes that are reachable are highlighted in yellow.
// Example 2:
//   Input: edges = [[0,1,4],[1,2,6],[0,2,8],[1,3,1]], maxMoves = 10, n = 4
//   Output: 23
// Example 3:
//   Input: edges = [[1,2,4],[1,4,5],[1,3,1],[2,3,4],[3,4,5]], maxMoves = 17, n = 5
//   Output: 1
//   Explanation: Node 0 is disconnected from the rest of the graph, so only node 0 is reachable.
// Constraints:
//   0 <= edges.length <= min(n * (n - 1) / 2, 10⁴)
//   edges[i].length == 3
//   0 <= ui < vi < n
//   There are no multiple edges in the graph.
//   0 <= cnti <= 10⁴
//   0 <= maxMoves <= 10⁹
//   1 <= n <= 3000

// single source shortest path, do it in dijkstra manner
func reachableNodes(es [][]int, maxMoves int, n int) int {
	ans := 1
	// make a graph
	graph := make([][][2]int, n)
	for _, e := range es {
		graph[e[0]] = append(graph[e[0]], [2]int{e[1], e[2] + 1})
		graph[e[1]] = append(graph[e[1]], [2]int{e[0], e[2] + 1})
	}

	shortest := make([]int, n)
	shortest[0] = 1 // set 0's shortest path to 1 for easy compute; so need also increment maxMoves
	maxMoves++
	// push all node0's edges to heap
	h := &edges{}
	for _, e := range graph[0] {
		heap.Push(h, [4]int{0, e[0], e[1] + 1, e[1]}) // from, to, shortest path length, edge length
	}

	// main process
	for h.Len() > 0 {
		cur := heap.Pop(h).([4]int)
		n1, n2, sp, eg := cur[0], cur[1], cur[2], cur[3]

		// both nodes reachable
		if shortest[n1] > 0 && shortest[n2] > 0 {
			canReach := maxMoves - shortest[n1] + maxMoves - shortest[n2]
			if canReach > eg-1 {
				canReach = eg - 1
			}
			ans += canReach
			continue
		}
		// if only one reachable, let n1 be source, n2 be candidate
		if shortest[n2] > 0 {
			n1, n2 = n2, n1
		}
		if sp <= maxMoves { // n2 can be reachable
			shortest[n2] = sp
			ans += sp - shortest[n1]
			for _, e := range graph[n2] { // for n2's unvisited neighbor
				if shortest[e[0]] == 0 {
					heap.Push(h, [4]int{n2, e[0], sp + e[1], e[1]})
				}
			}
		} else { // n2 unreachable
			ans += maxMoves - shortest[n1]
		}
	}
	return ans
}

type edges [][4]int // (vertex1, vertex2, shortest path, edge length) tuple

func (t edges) Len() int {
	return len(t)
}

func (t edges) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t edges) Less(i, j int) bool {
	return t[i][2] < t[j][2]
}

func (t *edges) Push(x interface{}) {
	*t = append(*t, x.([4]int))
}

func (t *edges) Pop() interface{} {
	x := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return x
}

func main() {
	for _, v := range []struct {
		e         [][]int
		m, n, ans int
	}{
		{[][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}, 6, 3, 13},
		{[][]int{{0, 1, 4}, {1, 2, 6}, {0, 2, 8}, {1, 3, 1}}, 10, 4, 23},
		{[][]int{{1, 2, 4}, {1, 4, 5}, {1, 3, 1}, {2, 3, 4}, {3, 4, 5}}, 17, 5, 1},
		{[][]int{{1, 2, 5}, {0, 3, 3}, {1, 3, 2}, {2, 3, 4}, {0, 4, 1}}, 7, 5, 13},
		{[][]int{{2, 4, 2}, {3, 4, 5}, {2, 3, 1}, {0, 2, 1}, {0, 3, 5}}, 14, 5, 18},
	} {
		fmt.Println(reachableNodes(v.e, v.m, v.n), v.ans)
	}
}

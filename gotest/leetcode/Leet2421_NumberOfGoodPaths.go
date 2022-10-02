package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/number-of-good-paths/

// There is a tree (i.e. a connected, undirected graph with no cycles) consisting of n nodes numbered
// from 0 to n - 1 and exactly n - 1 edges.
// You are given a 0-indexed integer array vals of length n where vals[i] denotes the value of the
// iᵗʰ node. You are also given a 2D integer array edges where edges[i] = [ai, bi] denotes that there
// exists an undirected edge connecting nodes ai and bi.
// A good path is a simple path that satisfies the following conditions:
//   The starting node and the ending node have the same value.
//   All nodes between the starting node and the ending node have values less than or equal to the
//     starting node (i.e. the starting node's value should be the maximum value along the path).
// Return the number of distinct good paths.
// Note that a path and its reverse are counted as the same path. For example, 0 - > 1 is considered
// to be the same as 1 -> 0. A single node is also considered as a valid path.
// Example 1:
//   Input: vals = [1,3,2,1,3], edges = [[0,1],[0,2],[2,3],[2,4]]
//   Output: 6
//   Explanation: There are 5 good paths consisting of a single node.
//     There is 1 additional good path: 1 -> 0 -> 2 -> 4.
//     (The reverse path 4 -> 2 -> 0 -> 1 is treated as the same as 1 -> 0 -> 2 -> 4.)
//     Note that 0 -> 2 -> 3 is not a good path because vals[2] > vals[0].
// Example 2:
//   Input: vals = [1,1,2,2,3], edges = [[0,1],[1,2],[2,3],[2,4]]
//   Output: 7
//   Explanation: There are 5 good paths consisting of a single node.
//     There are 2 additional good paths: 0 -> 1 and 2 -> 3.
// Example 3:
//   Input: vals = [1], edges = []
//   Output: 1
//   Explanation: The tree consists of only one node, so there is one good path.
// Constraints:
//   n == vals.length
//   1 <= n <= 3 * 10⁴
//   0 <= vals[i] <= 10⁵
//   edges.length == n - 1
//   edges[i].length == 2
//   0 <= ai, bi < n
//   ai != bi
//   edges represents a valid tree.

// union-find: add vertices to graph from smaller value
func numberOfGoodPaths(vals []int, edges [][]int) int {
	visited := make([]bool, len(vals))
	// sort edges by max vals of two vertices
	sort.Slice(edges, func(i, j int) bool {
		return max(vals[edges[i][0]], vals[edges[i][1]]) < max(vals[edges[j][0]], vals[edges[j][1]])
	})

	// union-find
	ans := len(vals)
	root := make([]int, len(vals))
	for i := range root {
		root[i] = -1
	}
	start := 0 // group edges with same max val
	for start < len(edges) {
		curmax := max(vals[edges[start][1]], vals[edges[start][0]])
		end := start
		for end+1 < len(edges) && max(vals[edges[end+1][1]], vals[edges[end+1][0]]) == curmax {
			end++
		}
		// add the edges in a group to graph
		for i := start; i <= end; i++ {
			r0 := getRoot(root, edges[i][0])
			r1 := getRoot(root, edges[i][1])
			root[r1] = r0
		}
		// count pair
		count := make(map[int]int, 4)
		for i := start; i <= end; i++ {
			if !visited[edges[i][0]] {
				if vals[edges[i][0]] == curmax { // for vertex with value = curmax, group them
					r := getRoot(root, edges[i][0])
					count[r] = count[r] + 1
				}
				visited[edges[i][0]] = true
			}
			if !visited[edges[i][1]] {
				if vals[edges[i][1]] == curmax {
					r := getRoot(root, edges[i][1])
					count[r] = count[r] + 1
				}
				visited[edges[i][1]] = true
			}
		}
		for _, v := range count { // the vertices in same group can pair
			ans += v * (v - 1) / 2
		}
		start = end + 1
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getRoot(root []int, ind int) int {
	if root[ind] != -1 {
		root[ind] = getRoot(root, root[ind])
		return root[ind]
	}
	return ind
}

func main() {
	for _, v := range []struct {
		vals  []int
		edges [][]int
		ans   int
	}{
		{[]int{1, 3, 2, 1, 3}, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}}, 6},
		{[]int{1, 1, 2, 2, 3}, [][]int{{0, 1}, {1, 2}, {2, 3}, {2, 4}}, 7},
		{[]int{1}, [][]int{}, 1},
	} {
		fmt.Println(numberOfGoodPaths(v.vals, v.edges), v.ans)
	}
}

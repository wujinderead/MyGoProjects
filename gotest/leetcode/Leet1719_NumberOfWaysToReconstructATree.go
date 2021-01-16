package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/number-of-ways-to-reconstruct-a-tree/

// You are given an array pairs, where pairs[i] = [xi, yi], and:
//   There are no duplicates.
//   xi < yi
// Let ways be the number of rooted trees that satisfy the following conditions:
//   The tree consists of nodes whose values appeared in pairs.
//   A pair [xi, yi] exists in pairs if and only if xi is an ancestor of yi or yi is an ancestor of xi.
// Note: the tree does not have to be a binary tree.
// Two ways are considered to be different if there is at least one node that has different parents in both ways.
// Return:
//   0 if ways == 0
//   1 if ways == 1
//   2 if ways > 1
// A rooted tree is a tree that has a single root node, and all edges are oriented to be outgoing from the root.
// An ancestor of a node is any node on the path from the root to that node (excluding the node itself).
// The root has no ancestors.
// Example 1:
//   Input: pairs = [[1,2],[2,3]]
//   Output: 1
//   Explanation: There is exactly one valid rooted tree, which is shown in the above figure.
// Example 2:
//   Input: pairs = [[1,2],[2,3],[1,3]]
//   Output: 2
//   Explanation: There are multiple valid rooted trees. Three of them are shown in the above figures.
// Example 3:
//   Input: pairs = [[1,2],[2,3],[2,4],[1,5]]
//   Output: 0
//   Explanation: There are no valid rooted trees.
// Constraints:
//   1 <= pairs.length <= 10^5
//   1 <= xi < yi <= 500
//   The elements in pairs are unique.

// the root node must have the most number of adjacent node, we can sort by degree and
// construct the tree.
func checkWays(pairs [][]int) int {
	visited := make(map[[2]int]bool)
	parent := make(map[int]int)
	graph := make(map[int][]int)
	for _, v := range pairs {
		graph[v[0]] = append(graph[v[0]], v[1])
		graph[v[1]] = append(graph[v[1]], v[0])
		parent[v[0]] = 0
		parent[v[1]] = 0
	}
	degree := make([][2]int, 0, len(graph))
	for k, v := range graph {
		degree = append(degree, [2]int{k, len(v)}) // (node, degree) pair
	}
	sort.Slice(degree, func(i, j int) bool {
		return degree[i][1] > degree[j][1]
	})
	for _, v := range degree {
		cur := v[0]
		curparent := parent[cur]
		children := graph[cur]
		for _, ch := range children {
			key := [2]int{ch, curparent}
			if ch > curparent {
				key = [2]int{curparent, ch}
			}
			if curparent != 0 && curparent != ch && !visited[key] {
				return 0
			}
			key = [2]int{cur, ch}
			if cur > ch {
				key = [2]int{ch, cur}
			}
			if !visited[key] {
				visited[key] = true
				parent[ch] = cur
			}
		}
	}
	// count parent
	count := make(map[int]int)
	for _, v := range parent {
		count[v] = count[v] + 1
	}
	if count[0] > 1 { // multiple nodes have parent = 0, means multiple roots
		return 0
	}
	for k, v := range count {
		if k != 0 && v == 1 { // if a non-root node has only one child, it's switchable with parent
			return 2
		}
	}
	return 1
}

func main() {
	for _, v := range []struct {
		p   [][]int
		ans int
	}{
		{[][]int{{1, 2}, {2, 3}}, 1},
		{[][]int{{1, 2}, {2, 3}, {1, 3}}, 2},
		{[][]int{{1, 2}, {2, 3}, {2, 4}, {1, 5}}, 0},
		{[][]int{{1, 2}, {3, 4}}, 0},
		{[][]int{{1, 2}, {2, 3}, {3, 4}}, 0},
		{[][]int{{1, 2}, {2, 3}, {2, 4}, {3, 4}}, 2},
	} {
		fmt.Println(checkWays(v.p), v.ans)
	}
}

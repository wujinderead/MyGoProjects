package main

import "fmt"

// https://leetcode.com/problems/valid-arrangement-of-pairs/

// You are given a 0-indexed 2D integer array pairs where pairs[i] = [starti, endi].
// An arrangement of pairs is valid if for every index i where 1 <= i < pairs.length,
// we have endi-1 == starti.
// Return any valid arrangement of pairs.
// Note: The inputs will be generated such that there exists a valid arrangement of pairs.
// Example 1:
//   Input: pairs = [[5,1],[4,5],[11,9],[9,4]]
//   Output: [[11,9],[9,4],[4,5],[5,1]]
//   Explanation:
//     This is a valid arrangement since endi-1 always equals starti.
//     end0 = 9 == 9 = start1
//     end1 = 4 == 4 = start2
//     end2 = 5 == 5 = start3
// Example 2:
//   Input: pairs = [[1,3],[3,2],[2,1]]
//   Output: [[1,3],[3,2],[2,1]]
//   Explanation:
//     This is a valid arrangement since endi-1 always equals starti.
//     end0 = 3 == 3 = start1
//     end1 = 2 == 2 = start2
//     The arrangements [[2,1],[1,3],[3,2]] and [[3,2],[2,1],[1,3]] are also valid.
// Example 3:
//   Input: pairs = [[1,2],[1,3],[2,1]]
//   Output: [[1,2],[2,1],[1,3]]
//   Explanation:
//     This is a valid arrangement since endi-1 always equals starti.
//     end0 = 2 == 2 = start1
//     end1 = 1 == 1 = start2
// Constraints:
//   1 <= pairs.length <= 10⁵
//   pairs[i].length == 2
//   0 <= starti, endi <= 10⁹
//   starti != endi
//   No two pairs are exactly the same.
//   There exists a valid arrangement of pairs.

// see the pairs as a graph, we want to find an Euler Path, which traverse every edges exact once.
// if for every vertices in a graph, inDegree == outDegree, we can find an Euler Circuit;
// or every other vertices has inDegree == outDegree, except for two vertices,
// one with inDegree==outDegree+1, one with inDegree==outDegree-1.
// the vertex with inDegree==outDegree-1 is the start vertex of an Euler Path.
func validArrangement(pairs [][]int) [][]int {
	indegree, outdegree := make(map[int]int), make(map[int]int)
	graph := make(map[int][]int)
	nodes := make(map[int]struct{})
	for _, v := range pairs {
		graph[v[0]] = append(graph[v[0]], v[1])
		outdegree[v[0]] = outdegree[v[0]] + 1
		indegree[v[1]] = indegree[v[1]] + 1
		nodes[v[0]] = struct{}{}
		nodes[v[1]] = struct{}{}
	}
	start := -1
	for k := range nodes {
		if outdegree[k] == indegree[k]+1 {
			start = k
			break
		}
	}
	path := make([][]int, 0, len(pairs))
	if start == -1 {
		visit(pairs[0][0], graph, &path)
	} else {
		visit(start, graph, &path)
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func visit(node int, graph map[int][]int, path *[][]int) {
	for {
		nexts := graph[node]
		if len(nexts) == 0 {
			break
		}
		next := nexts[len(nexts)-1]
		graph[node] = nexts[:len(nexts)-1]
		visit(next, graph, path)
		*path = append(*path, []int{node, next}) // post-order visit, because we want goes back
	}
}

func main() {
	for _, v := range []struct {
		p, ans [][]int
	}{
		{[][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}, [][]int{{11, 9}, {9, 4}, {4, 5}, {5, 1}}},
		{[][]int{{1, 3}, {3, 2}, {2, 1}}, [][]int{{1, 3}, {3, 2}, {2, 1}}},
		{[][]int{{1, 2}, {1, 3}, {2, 1}}, [][]int{{1, 2}, {2, 1}, {1, 3}}},
	} {
		fmt.Println(validArrangement(v.p), v.ans)
	}
}

package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/checking-existence-of-edge-length-limited-paths/

// An undirected graph of n nodes is defined by edgeList, where edgeList[i] = [ui, vi, disi]
// denotes an edge between nodes ui and vi with distance disi. Note that there may be
// multiple edges between two nodes.
// Given an array queries, where queries[j] = [pj, qj, limitj], your task is to determine
// for each queries[j] whether there is a path between pj and qj such that each edge on
// the path has a distance strictly less than limitj.
// Return a boolean array answer, where answer.length == queries.length and the jth value
// of answer is true if there is a path for queries[j] is true, and false otherwise.
// Example 1:
//   Input: n = 3, edgeList = [[0,1,2],[1,2,4],[2,0,8],[1,0,16]], queries = [[0,1,2],[0,2,5]]
//   Output: [false,true]
//   Explanation: The above figure shows the given graph. Note that there are two overlapping
//     edges between 0 and 1 with distances 2 and 16.
//     For the first query, between 0 and 1 there is no path where each distance is less than 2,
//     thus we return false for this query. For the second query, there is a path (0 -> 1 -> 2)
//     of two edges with distances less than 5, thus we return true for this query.
// Example 2:
//   Input: n = 5, edgeList = [[0,1,10],[1,2,5],[2,3,9],[3,4,13]], queries = [[0,4,14],[1,4,13]]
//   Output: [true,false]
//   Explanation: The above figure shows the given graph.
// Constraints:
//   2 <= n <= 10^5
//   1 <= edgeList.length, queries.length <= 10^5
//   edgeList[i].length == 3
//   queries[j].length == 3
//   0 <= ui, vi, pj, qj <= n - 1
//   ui != vi
//   pj != qj
//   1 <= disi, limitj <= 10^9
//   There may be multiple edges between two nodes.

// sort edges and queries, use union-find to add edges to sets.
// for a query, add edges less than limit to sets,
// and check if two query vertices are in same set, as we only care about connectivity.
func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	ans := make([]bool, len(queries))
	arr := make([]int, n)
	for i := range arr {
		arr[i] = -1
	}
	mapp := make(map[[2]int]struct{})
	newqueries := make([][]int, len(queries))
	for i := range newqueries {
		newqueries[i] = []int{queries[i][0], queries[i][1], queries[i][2], i}
	}

	// sort edges and queries
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})
	sort.Slice(newqueries, func(i, j int) bool {
		return newqueries[i][2] < newqueries[j][2]
	})

	// main logic
	ei := 0
	for _, q := range newqueries {
		// union: add edge[ei] that less than query to union set
		for ei < len(edgeList) && edgeList[ei][2] < q[2] {
			a, b := edgeList[ei][0], edgeList[ei][1]
			ei++
			if _, ok := mapp[[2]int{a, b}]; ok { // ignore duplicated edge
				continue
			}
			mapp[[2]int{a, b}] = struct{}{}

			// union
			ra, rb := root(arr, a), root(arr, b)
			if ra == rb { // already in same set
				continue
			}
			arr[rb] = ra
		}
		// find: test if two query points are in same set
		r0, r1 := root(arr, q[0]), root(arr, q[1])
		if r0 == r1 {
			ans[q[3]] = true // set original index
		}
	}
	return ans
}

func root(arr []int, i int) int {
	if arr[i] == -1 {
		return i
	}
	x := root(arr, arr[i])
	arr[i] = x
	return x
}

func main() {
	for _, v := range []struct {
		n      int
		es, qs [][]int
		ans    []bool
	}{
		{3, [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}},
			[][]int{{0, 1, 2}, {0, 2, 5}}, []bool{false, true}},
		{5, [][]int{{0, 1, 10}, {1, 2, 5}, {2, 3, 9}, {3, 4, 13}},
			[][]int{{0, 4, 14}, {1, 4, 13}}, []bool{true, false}},
	} {
		fmt.Println(distanceLimitedPathsExist(v.n, v.es, v.qs), v.ans)
	}
}

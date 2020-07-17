package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/path-with-maximum-probability/

// You are given an undirected weighted graph of n nodes (0-indexed), represented by an edge list 
// where edges[i] = [a, b] is an undirected edge connecting the nodes a and b with a probability 
// of success of traversing that edge succProb[i]. Given two nodes start and end, find the path with 
// the maximum probability of success to go from start to end and return its success probability.
// If there is no path from start to end, return 0. Your answer will be accepted if it differs from 
// the correct answer by at most 1e-5.
// Example 1:
//           0
//   0.5  /    \  0.2
//       1 ---- 2
//         0.5
//   Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.2], start = 0, end = 2
//   Output: 0.25000
//   Explanation: There are two paths from start to end, one having a probability of success = 0.2 
//     and the other has 0.5 * 0.5 = 0.25.
// Example 2:
//   Input: n = 3, edges = [[0,1],[1,2],[0,2]], succProb = [0.5,0.5,0.3], start = 0, end = 2
//   Output: 0.30000
// Example 3:
//   Input: n = 3, edges = [[0,1]], succProb = [0.5], start = 0, end = 2
//   Output: 0.00000
//   Explanation: There is no path between 0 and 2.
// Constraints:
//   2 <= n <= 10^4
//   0 <= start, end < n
//   start != end
//   0 <= a, b < n
//   a != b
//   0 <= succProb.length == edges.length <= 2*10^4
//   0 <= succProb[i] <= 1
//   There is at most one edge between every two nodes.

// do it in a dijkstra manner
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	// form the graph
	graph := make([][]pair, n)
	for i := range edges {
		from, to, prob := edges[i][0], edges[i][1], succProb[i]
		graph[from] = append(graph[from], pair{to, prob})
		graph[to] = append(graph[to], pair{from, prob})
	}

	// dijkstra-like method: put all possible probability to a heap,
	// get the most probability and update others
	h := pairs(make([]pair, 1, 20))
	inset := make([]bool, n)
	h[0] = pair{start, 1.0}   // the prob for to start point is 1.0
	for len(h)>0 {
		p := heap.Pop(&h).(pair)
		if inset[p.ind] {
			continue      // skip visited point
		}
		inset[p.ind] = true   // mark current point visited
		if p.ind==end {       // if current point is target, return
			return p.prob
		}
		for _, v := range graph[p.ind] {   // for current point's unvisited neighbor
			if !inset[v.ind] {
				heap.Push(&h, pair{v.ind, p.prob*v.prob})   // push new prob for neighbor
			}
		}
	}
	return 0.0
}

type pair struct {
	ind int
	prob float64
}

type pairs []pair

func (p pairs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p pairs) Len() int { return len(p) }
func (p pairs) Less(i, j int) bool { return p[i].prob > p[j].prob }
func (p *pairs) Push(x interface{}) { *p = append(*p, x.(pair)) }
func (p *pairs) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func main() {
	fmt.Println(maxProbability(3, [][]int{{0,1},{1,2},{0,2}}, []float64{0.5,0.5,0.2}, 0, 2))
	fmt.Println(maxProbability(3, [][]int{{0,1},{1,2},{0,2}}, []float64{0.5,0.5,0.3}, 0, 2))
	fmt.Println(maxProbability(3, [][]int{{0,1}}, []float64{0.5}, 0, 2))
	fmt.Println(maxProbability(5, [][]int{{0,1},{1,2},{2,3},{0,4},{1,4},{2,4},{3,4},{2,0}}, 
		[]float64{0.9, 0.9, 0.9, 0.1, 0.2, 0.3, 0.4, 0.8}, 0, 4))
}
package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/second-minimum-time-to-reach-destination/

// A city is represented as a bi-directional connected graph with n vertices where each vertex is
// labeled from 1 to n (inclusive). The edges in the graph are represented as a 2D integer array
// edges, where each edges[i] = [ui, vi] denotes a bi-directional edge between vertex ui and vertex vi.
// Every vertex pair is connected by at most one edge, and no vertex has an edge to itself. The time
// taken to traverse any edge is time minutes.
// Each vertex has a traffic signal which changes its color from green to red and vice versa every change
// minutes. All signals change at the same time. You can enter a vertex at any time, but can leave a vertex
// only when the signal is green. You cannot wait at a vertex if the signal is green.
// The second minimum value is defined as the smallest value strictly larger than the minimum value.
// For example the second minimum value of [2, 3, 4] is 3, and the second minimum value of [2, 2, 4] is 4.
// Given n, edges, time, and change, return the second minimum time it will take to go from vertex 1 to vertex n.
// Notes:
//   You can go through any vertex any number of times, including 1 and n.
//   You can assume that when the journey starts, all signals have just turned green.
// Example 1:
//   Input: n = 5, edges = [[1,2],[1,3],[1,4],[3,4],[4,5]], time = 3, change = 5
//   Output: 13
//   Explanation:
//     The figure on the left shows the given graph.
//     The blue path in the figure on the right is the minimum time path.
//     The time taken is:
//     - Start at 1, time elapsed=0
//     - 1 -> 4: 3 minutes, time elapsed=3
//     - 4 -> 5: 3 minutes, time elapsed=6
//     Hence the minimum time needed is 6 minutes.
//     The red path shows the path to get the second minimum time.
//     - Start at 1, time elapsed=0
//     - 1 -> 3: 3 minutes, time elapsed=3
//     - 3 -> 4: 3 minutes, time elapsed=6
//     - Wait at 4 for 4 minutes, time elapsed=10
//     - 4 -> 5: 3 minutes, time elapsed=13
//     Hence the second minimum time is 13 minutes.
// Example 2:
//   Input: n = 2, edges = [[1,2]], time = 3, change = 2
//   Output: 11
//   Explanation:
//     The minimum time path is 1 -> 2 with time = 3 minutes.
//     The second minimum time path is 1 -> 2 -> 1 -> 2 with time = 11 minutes.
// Constraints:
//   2 <= n <= 10^4
//   n - 1 <= edges.length <= min(2 * 10^4, n * (n - 1) / 2)
//   edges[i].length == 2
//   1 <= ui, vi <= n
//   ui != vi
//   There are no duplicate edges.
//   Each vertex can be reached directly or indirectly from every other vertex.
//   1 <= time, change <= 10^3

// Another method:
// https://leetcode.com/problems/second-minimum-time-to-reach-destination/discuss/1525154/No-Dijkstra%3A-(n-%2B-1)-or-(n-%2B-2)
// since `time` for each edge is the same, actually `change` and `time` does not affect the algorithm
// use BFS to find the shortest path from 1 to n, say x steps;
// then the second shortest path must be the follwong 2 situations:
//  x+1 steps, if their is a detour
//  x+2 steps, go back-and-forth once

// alike dijkstra, in this problem, we need the shortest and second shortest path
func secondMinimum(n int, edges [][]int, time int, change int) int {
	graph := make([][]int, n+1)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	h := vs([][2]int{{1, 2 * change}}) // vertex 0, time 2*change to let firstMin[1]>0
	firstMin := make([]int, 1+n)       // minimal time to node[i]
	secondMin := make([]int, 1+n)      // second minimal time to node[i]
	for len(h) > 0 {                   // use a heap to get the shortest path
		cur := heap.Pop(&h).([2]int)
		cind, ctime := cur[0], cur[1] // pop node with minimal time
		nexttime := ctime + time
		if (ctime/change)%2 == 1 { // add the waiting time due to red signal
			nexttime += change - ctime%change
		}
		for _, next := range graph[cind] { // to adjacent nodes
			if firstMin[next] == 0 { // if first reach, set minimal time
				firstMin[next] = nexttime
			} else if secondMin[next] == 0 && nexttime > firstMin[next] {
				secondMin[next] = nexttime // if second reach, set second minimal time
			} else {
				continue // do nothing
			}
			heap.Push(&h, [2]int{next, nexttime}) // push to heap
		}
	}
	return secondMin[n] - 2*change
}

type vs [][2]int

func (v vs) Len() int {
	return len(v)
}

func (v vs) Less(i, j int) bool {
	return v[i][1] < v[j][1]
}

func (v vs) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

func (v *vs) Push(x interface{}) {
	*v = append(*v, x.([2]int))
}

func (v *vs) Pop() interface{} {
	x := (*v)[len(*v)-1]
	*v = (*v)[:len(*v)-1]
	return x
}

func main() {
	for _, v := range []struct {
		n         int
		e         [][]int
		t, c, ans int
	}{
		{5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5, 13},
		{2, [][]int{{1, 2}}, 3, 2, 11},
	} {
		fmt.Println(secondMinimum(v.n, v.e, v.t, v.c), v.ans)
	}
}

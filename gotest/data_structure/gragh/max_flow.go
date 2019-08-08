package gragh

import (
	"container/list"
	"fmt"
	"math"
)

func maxFlowFordFulkerson(g *matrix, s, t int) int {
	// initialize residual graph as the same as g
	residual := new(matrix)
	residual.n = g.n
	residual.matrix = make([]int, g.n*g.n)
	copy(residual.matrix, g.matrix)

	parent := make([]int, g.n) // store the path from s to t

	maxFlow := 0

	// Augment the flow while there is path from source to sink
	for maxFlowBfs(residual, s, t, parent) {
		// Find minimum residual capacity of the edges along the
		// path filled by BFS. Or we can say find the maximum flow
		// through the path found.
		fmt.Println("path:")
		for v := t; v != -1; v = parent[v] {
			fmt.Print(v, " ")
		}
		fmt.Println()
		pathFlow := math.MaxInt64
		for v := t; v != s; v = parent[v] {
			u := parent[v]
			pathFlow = min(pathFlow, residual.get(u, v))
		}
		fmt.Println(pathFlow)

		// update residual capacities of the edges and reverse edges
		// along the path
		for v := t; v != s; v = parent[v] {
			u := parent[v]
			residual.set(u, v, residual.get(u, v)-pathFlow)
			residual.set(v, u, residual.get(v, u)+pathFlow)
		}
		residual.print()

		// Add path flow to overall flow
		maxFlow += pathFlow
	}

	// Return the overall flow
	return maxFlow
}

func maxFlowBfs(residual *matrix, s, t int, parent []int) bool {
	queue := list.New()
	queue.PushBack(s)
	parent[s] = -1
	visited := make([]bool, residual.n)
	visited[s] = true
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(int)
		for i := 0; i < residual.n; i++ {
			if residual.get(cur, i) > 0 && !visited[i] {
				visited[i] = true
				parent[i] = cur
				queue.PushBack(i)
			}
		}
	}
	return visited[t]
}

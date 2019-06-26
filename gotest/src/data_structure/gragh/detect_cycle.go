package gragh

import "fmt"

func detectCycleUndirectedDfs(g *graph) bool {
	// dfs to store nods in stack, assume no parallel edge between two verticesß
	stack := make([]int, g.n)
	parent := make([]int, g.n)
	stackind := -1
	visited := make([]bool, g.n)
	stackind++
	stack[stackind] = 0
	parent[stackind] = 0 // initial point's parent is self
	for stackind >= 0 {
		cur := stack[stackind]
		curparent := parent[stackind] // pop stack
		stackind--
		visited[cur] = true
		// traverse cur's edges
		for nb := g.adjacency[cur]; nb != nil; nb = nb.next {
			// if cur has a visited adjacent vertex and this vertex is not cur's parent
			// then there must be a cycle
			fmt.Println(cur, "->", nb.id, curparent, visited[nb.id])
			if visited[nb.id] && nb.id != curparent {
				return true
			}
			if !visited[nb.id] {
				stackind++
				stack[stackind] = nb.id // add cur's child to stack
				parent[stackind] = cur
			}
		}
	}
	return false
}

func detectCycleDirectedDfs(g *graph) {

}

// union is for undirected graph
func detectCycleUnionFind(g *graph) {

}

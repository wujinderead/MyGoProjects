package gragh

import "fmt"

// assume no parallel edge between two vertices
func detectCycleUndirectedDfs(g *graph) bool {
	// use stack to perform dfs
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

func detectCycleDirectedDfs(g *graph) bool {
	// use stack to perform dfs
	instack := make([]bool, g.n)
	visited := make([]bool, g.n)
	for i := 0; i < g.n; i++ {
		if !visited[i] && detectCycleDirectedUtil(g, i, visited, instack) {
			return true
		}
	}
	return false
}

func detectCycleDirectedUtil(g *graph, cur int, visited, instack []bool) bool {
	visited[cur] = true
	instack[cur] = true
	// defer func(){instack[cur]=false}()   // can use defer here, buf defer will affect performance
	for nb := g.adjacency[cur]; nb != nil; nb = nb.next {
		// if cur has a adjacent vertex that is already in the stack
		// then there is a cycle
		fmt.Println(cur, "->", nb.id)
		if instack[nb.id] {
			fmt.Println("back edge:", cur, nb.id)
			// set instack[cur] to false when leave current func call
			instack[cur] = false
			return true
		}
		if !visited[nb.id] && detectCycleDirectedUtil(g, nb.id, visited, instack) {
			instack[cur] = false
			return true
		}
	}
	instack[cur] = false // set in stack to false
	return false
}

// union-find is for undirected graph
// add edges to disjoint set one by one, if two vertices of an edge is in the same set
// then there is a cycle
func detectCycleUnionFind(g *graph) bool {
	disjointset := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		disjointset[i] = -1
	}
	for i := 0; i < g.n; i++ {
		for nb := g.adjacency[i]; nb != nil; nb = nb.next {
			if i < nb.id { // only process edge for once
				x := find(disjointset, i)
				y := find(disjointset, nb.id)
				fmt.Println(i, nb.id, x, y)
				if x == y {
					return true
				}
				union(disjointset, x, y)
			}
		}
	}
	return false
}

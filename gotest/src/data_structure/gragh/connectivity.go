package gragh

// for directed graph, check whether the whole graph is strongly connected
func checkStronglyConnected(g *graph) bool {
	// use two runs of dfs, the second one use the transposed graph of original graph
	visited := make([]bool, g.n)
	count := new(int) // how many vertices are visited
	dfsUtil(0, g, visited, count)
	if *count < g.n {
		return false // can't go to every other vertices from 0, not strongly connected
	}

	// transpose graph
	*count = 0
	trans := new(graph) // to store the transposed
	trans.n = g.n
	trans.adjacency = make([]*neighbor, g.n)
	for i := 0; i < g.n; i++ {
		visited[i] = false
		for nb := g.adjacency[i]; nb != nil; nb = nb.next {
			trans.adjacency[nb.id] = &neighbor{i, 0, trans.adjacency[nb.id]}
		}
	}
	trans.print()

	// dfs transpose graph
	dfsUtil(0, trans, visited, count)
	if *count < g.n {
		// can't go to every other vertices from 0 in transposed graph
		// not strongly connected
		return false
	}
	return true
}

func dfsUtil(cur int, g *graph, visited []bool, count *int) {
	visited[cur] = true
	*count++
	for nb := g.adjacency[cur]; nb != nil; nb = nb.next {
		if !visited[nb.id] {
			dfsUtil(nb.id, g, visited, count)
		}
	}
}

// find all strongly connected components of a graph
func findStronglyConnectedComponents(g *graph) [][]int {
	// dfs the graph and push visted vertices in stack
	index := new(int)
	*index = -1
	stack := make([]int, g.n)
	visited := make([]bool, g.n)
	for i := 0; i < g.n; i++ {
		if !visited[i] {
			sccDfsUtil1(0, g, visited, stack, index)
		}
	}
	// dfs the transposed graph according to the popped order of stack
	trans := new(graph)
	trans.n = g.n
	trans.adjacency = make([]*neighbor, g.n)
	for i := 0; i < g.n; i++ {
		visited[i] = false
		for nb := g.adjacency[i]; nb != nil; nb = nb.next {
			// copy each edge to transposed graph
			trans.adjacency[nb.id] = &neighbor{i, 0, trans.adjacency[nb.id]}
		}
	}
	sccs := make([][]int, 0)
	for *index >= 0 {
		cur := stack[*index]
		*index--
		if !visited[cur] {
			scc := make([]int, 0)
			sccDfsUtil2(cur, trans, visited, &scc)
			sccs = append(sccs, scc)
		}
	}
	return sccs
}

func sccDfsUtil1(cur int, g *graph, visited []bool, stack []int, index *int) {
	visited[cur] = true
	for nb := g.adjacency[cur]; nb != nil; nb = nb.next {
		if !visited[nb.id] {
			sccDfsUtil1(nb.id, g, visited, stack, index)
		}
	}
	// like topological sort, when dfs walk back, we push current vertex to stack.
	// this will place the deepest vertex of dfs search path at the bottom of stack.
	// when we pop the stack, we get the uppermost vertex first
	*index++
	stack[*index] = cur
}

func sccDfsUtil2(cur int, g *graph, visited []bool, scc *[]int) {
	visited[cur] = true
	*scc = append(*scc, cur)
	for nb := g.adjacency[cur]; nb != nil; nb = nb.next {
		if !visited[nb.id] {
			sccDfsUtil2(nb.id, g, visited, scc)
		}
	}
}

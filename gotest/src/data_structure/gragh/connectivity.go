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

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
	// dfs the graph and push visited vertices in stack
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

// for undirected graph, find all its articulation vertices
func findArticulationPoints(g *graph) []int {
	ap := make([]bool, g.n)
	disc := make([]int, g.n) // the time when vertex is visited first time
	low := make([]int, g.n)  // the topmost reachable ancestor of vertex
	for i := 0; i < g.n; i++ {
		if disc[i] == 0 { // disc[i]==0 means hasn't been visited
			time := 1                              // set discovery time=1 for root
			apUtil(i, -1, g, ap, low, disc, &time) // root's parent is -1
		}
	}
	aps := make([]int, 0)
	for i := 0; i < g.n; i++ {
		if ap[i] {
			aps = append(aps, i)
		}
	}
	return aps
}

func apUtil(u, parent int, g *graph, ap []bool, low, disc []int, time *int) {
	childCount := 0
	// current vertex is u
	disc[u] = *time
	low[u] = *time // set disc[u] = low[u] when visited first time
	*time++
	for nb := g.adjacency[u]; nb != nil; nb = nb.next {
		v := nb.id // current child vertex is v
		childCount++
		if disc[v] == 0 { // not visited
			apUtil(v, u, g, ap, low, disc, time)
			// because of recursion, the deepest vertex get processed first
			low[u] = min(low[u], low[v]) // if u's child can go back to u's ancestor, u also can

			// u is an articulation point in following cases:
			// 1. u is root of DFS tree and has two or more children.
			if parent == -1 && childCount > 1 {
				ap[u] = true
			}
			// 2. u is not root, and u's child v can not go back to u's ancestor.
			// low[v]<disc[u] means that v can go back to u's ancestor, low[v]>=disc[u] means can't.
			if parent != -1 && low[v] >= disc[u] {
				ap[u] = true
			}
		} else if v != parent {
			low[u] = min(low[u], disc[v]) // if v can go back to u's ancestor, update low
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// find all strongly connected components of a graph use Tarjan’s algorithm
func findAllSccTarjan(g *graph) [][]int {
	sccs := make([][]int, 0)
	disc := make([]int, g.n)  // the time when vertex is visited first time
	low := make([]int, g.n)   // the topmost reachable ancestor of vertex
	stack := make([]int, g.n) // stack to store searched vertices
	index := -1
	instack := make([]bool, g.n) // whether vertex is in stack
	for i := 0; i < g.n; i++ {
		if disc[i] == 0 { // disc[i]==0 means not visited
			time := 1 // set discovery time=1 for root
			sccTarjanUtil(i, g, instack, low, disc, stack, &time, &index, &sccs)
		}
	}
	return sccs
}

func sccTarjanUtil(u int, g *graph, instack []bool, low, disc, stack []int, time, index *int, sccs *[][]int) {
	disc[u] = *time
	low[u] = *time // initialize disc and low equally
	*time++
	*index++
	stack[*index] = u // push to stack
	instack[u] = true
	for nb := g.adjacency[u]; nb != nil; nb = nb.next {
		v := nb.id
		if disc[v] == 0 { // v not visited
			sccTarjanUtil(v, g, instack, low, disc, stack, time, index, sccs)
			low[u] = min(low[u], low[v])
		} else if instack[v] { // v is visited and in stack, means u to v is a back edge, not cross edge
			low[u] = min(low[u], disc[v])
		}
	}

	// if after walk down and walk back, low is still equal to disc,
	// then current vertex is a head of a strongly connected components
	var w int
	if low[u] == disc[u] {
		scc := make([]int, 0)
		for stack[*index] != u { // pop the stack until popped vertex is current vertex
			w = stack[*index]
			*index--
			scc = append(scc, w)
			instack[w] = false
		}
		w = stack[*index]
		*index--
		scc = append(scc, w)
		instack[w] = false
		*sccs = append(*sccs, scc)
	}
}

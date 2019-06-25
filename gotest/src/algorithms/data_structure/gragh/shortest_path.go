package gragh

func dijkstraMatrix(g *matrix, src int) (dist []int, sptSet []int) {
	sptSet = make([]int, g.n)
	dist = make([]int, g.n)
	pred := make([]int, g.n)
	for i := range dist {
		dist[i] = INF
		sptSet[i] = -1
		pred[i] = -1
	}
	dist[src] = 0
	pred[src] = src
	for i := 0; i < g.n; i++ {
		mindist := INF
		minvert := src
		// find shortest distance vertex that not in spt set
		for j := 0; j < g.n; j++ {
			if sptSet[j] == -1 && dist[j] < mindist {
				mindist = dist[j]
				minvert = j
			}
		}
		if minvert == INF {
			break // the remaining vertex are unreachable from src, thus we can break here
		}
		// update shortest distance
		for j := 0; j < g.n; j++ {
			curdist := g.get(minvert, j) // this also works from directed graph
			if curdist > 0 && dist[minvert]+curdist < dist[j] {
				dist[j] = dist[minvert] + curdist
				pred[j] = minvert
			}
		}
		sptSet[minvert] = pred[minvert]
	}
	return
}

func dijkstraAdjacent(g *graph, src int) (mindist []int, sptSet []int) {
	sptSet = make([]int, g.n)
	indices := make([]int, g.n)
	mindist = make([]int, g.n)
	pred := make([]int, g.n)
	position := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		sptSet[i] = -1
		indices[i] = i
		position[i] = i
		mindist[i] = INF
	}
	h := &heap{indices, pred, mindist, position}
	h.mindist[src] = 0   // minimum distance for source vertex is 0
	h.minfrom[src] = src // source vertex's pred is itself
	h.siftUp(src)        // src's value is least, sift up to stack top
	for num := 0; num < g.n; num++ {
		i := h.pop() // pop vertex with minimum distance with mst set
		if mindist[i] == INF {
			break
		}
		for nb := g.adjacency[i]; nb != nil; nb = nb.next {
			j := nb.id // update shortest distance between src and j via i
			if h.mindist[i]+nb.weight < h.mindist[j] {
				h.mindist[j] = h.mindist[i] + nb.weight
				h.minfrom[j] = i
				if !h.siftDown(h.position[j]) { // need to sift after modification
					h.siftUp(h.position[j])
				}
			}
		}
		// set mst set
		sptSet[i] = h.minfrom[i]
	}
	return
}

func floydWarshall(g *matrix) [][]int {
	// this solution has not provide path information
	// however it can be achieved by using another 2D array to store the predecessor.
	dist := make([][]int, g.n)
	for i := range dist {
		dist[i] = make([]int, g.n)
		for j := range dist[i] {
			// initialize the distance matrix
			if i == j {
				dist[i][j] = 0 // dist[i][i]=0
			} else if g.get(i, j) == 0 {
				dist[i][j] = INF // dist[i][j]=INF if i, j is not directed linked
			} else {
				dist[i][j] = g.get(i, j) // real distance
			}
		}
	}
	// floyd-warshall algorithm
	for k := 0; k < g.n; k++ { // k represent the intermediate vertex, outermost loop
		for i := 0; i < g.n; i++ {
			for j := 0; j < g.n; j++ {
				// dist[i][k] and dist[k][j] should not be INF to avoid overflow
				if dist[i][k] != INF && dist[k][j] != INF && dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j] // this also works from directed graph
				}
			}
		}
	}
	return dist
}

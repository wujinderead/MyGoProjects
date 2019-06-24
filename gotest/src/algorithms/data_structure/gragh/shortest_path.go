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

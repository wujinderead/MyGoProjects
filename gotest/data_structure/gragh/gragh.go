package gragh

import (
	"fmt"
	"math"
)

const (
	INF  = math.MaxInt8
	NINF = math.MinInt8
)

type graph struct {
	n         int
	adjacency []*neighbor
}

type neighbor struct {
	id     int
	weight int
	next   *neighbor
}

type matrix struct {
	n      int
	matrix []int
}

func newDirectedGraph(size int, edges [][]int) *graph {
	g := new(graph)
	g.n = size
	g.adjacency = make([]*neighbor, size)
	isWeighted := len(edges[0]) == 3
	for i := range edges {
		from := edges[i][0]
		to := edges[i][1]
		weight := 0
		if isWeighted {
			weight = edges[i][2]
		}
		// always add at head
		g.adjacency[from] = &neighbor{id: to, next: g.adjacency[from], weight: weight}
	}
	return g
}

func newUndirectedGraph(size int, edges [][]int) *graph {
	g := new(graph)
	g.n = size
	g.adjacency = make([]*neighbor, size)
	isWeighted := len(edges[0]) == 3
	for i := range edges {
		from := edges[i][0]
		to := edges[i][1]
		weight := 0
		if isWeighted {
			weight = edges[i][2]
		}
		// always add at head
		g.adjacency[from] = &neighbor{id: to, next: g.adjacency[from], weight: weight}
		g.adjacency[to] = &neighbor{id: from, next: g.adjacency[to], weight: weight}
	}
	return g
}

func (g *graph) print() {
	for i := range g.adjacency {
		fmt.Print(i, " -> ")
		for nb := g.adjacency[i]; nb != nil; nb = nb.next {
			fmt.Printf("%d(%d), ", nb.id, nb.weight)
		}
		fmt.Println()
	}
}

// unweighted graph, 0 means no edge, 1 means exist an edge
func newUndirectedMatrix(size int, edges [][]int) *matrix {
	g := new(matrix)
	g.n = size
	g.matrix = make([]int, size*size)
	isWeighted := len(edges[0]) == 3
	for i := range edges {
		from := edges[i][0]
		to := edges[i][1]
		weight := 1
		if isWeighted {
			weight = edges[i][2]
		}
		g.matrix[from*size+to] = weight
		g.matrix[to*size+from] = weight
	}
	return g
}

// weighted graph, 0 means no edge, value means edge weight
func newDirectedMatrix(size int, edges [][]int) *matrix {
	g := new(matrix)
	g.n = size
	g.matrix = make([]int, size*size)
	isWeighted := len(edges[0]) == 3
	for i := range edges {
		from := edges[i][0]
		to := edges[i][1]
		weight := 1
		if isWeighted {
			weight = edges[i][2]
		}
		g.matrix[from*size+to] = weight
	}
	return g
}

func (g *matrix) get(i, j int) int {
	return g.matrix[i*g.n+j]
}

func (g *matrix) set(i, j, k int) {
	g.matrix[i*g.n+j] = k
}

func (g *matrix) print() {
	for i := 0; i < g.n; i++ {
		for j := 0; j < g.n; j++ {
			if g.matrix[i*g.n+j] > 0 {
				fmt.Println(i, "->", j, ",", g.matrix[i*g.n+j])
			}
		}
	}
}

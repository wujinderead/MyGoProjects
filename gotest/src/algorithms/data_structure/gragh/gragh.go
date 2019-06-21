package gragh

import (
	"fmt"
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

func newGraph(size int, edges [][]int) *graph {
	g := new(graph)
	g.n = size
	g.adjacency = make([]*neighbor, size)
	for i := range edges {
		from := edges[i][0]
		to := edges[i][1]
		// always add at head
		g.adjacency[from] = &neighbor{id: to, next: g.adjacency[from]}
	}
	return g
}

func newWeightedGraph(size int, edges [][]int) *graph {
	g := new(graph)
	g.n = size
	g.adjacency = make([]*neighbor, size)
	for i := range edges {
		from := edges[i][0]
		to := edges[i][1]
		weight := edges[i][2]
		// always add at head
		g.adjacency[from] = &neighbor{id: to, next: g.adjacency[from], weight: weight}
	}
	return g
}

func (g *graph) print() {
	for i := range g.adjacency {
		fmt.Print(i, " -> ")
		for nb := g.adjacency[i]; nb != nil; nb = nb.next {
			fmt.Print(nb.id, ", ")
		}
		fmt.Println()
	}
}

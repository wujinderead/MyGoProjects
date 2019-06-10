package gragh

import (
	"container/list"
	"fmt"
)

type graph struct {
	n         int
	adjacency []*neighbor
}

type neighbor struct {
	id   int
	next *neighbor
}

func newGraph(size int, edges [][]int) *graph {
	g := new(graph)
	g.n = size
	g.adjacency = make([]*neighbor, size, size)
	for i := range edges {
		from := edges[i][0]
		to := edges[i][1]
		// always add at head
		g.adjacency[from] = &neighbor{to, g.adjacency[from]}
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

func topologicalOne(g *graph) []int {
	top := make([]int, g.n, g.n)
	inDegree := make([]int, g.n, g.n)
	for i := range g.adjacency {
		for nb := g.adjacency[i]; nb != nil; nb = nb.next {
			inDegree[nb.id]++
		}
	}
	queue := list.New()
	for i, in := range inDegree {
		if in == 0 {
			queue.PushBack(i)
		}
	}
	count := 0
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(int)
		top[count] = cur
		count++
		for nb := g.adjacency[cur]; nb != nil; nb = nb.next {
			inDegree[nb.id]--
			if inDegree[nb.id] == 0 {
				queue.PushBack(nb.id)
			}
		}
	}
	if count < g.n {
		return nil
	}
	return top
}

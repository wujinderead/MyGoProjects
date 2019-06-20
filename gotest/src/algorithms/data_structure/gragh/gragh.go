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

func topologicalQueue(g *graph) []int {
	top := make([]int, g.n)
	inDegree := make([]int, g.n)
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

func topologicalDfs(g *graph) []int {
	top := make([]int, g.n)
	visited := make([]bool, g.n)
	tindex := g.n - 1
	for i := range g.adjacency {
		if !visited[i] {
			topologicalDfsHelper(g, i, &tindex, top, visited)
		}
	}
	return top
}

func topologicalDfsHelper(g *graph, i int, tindex *int, top []int, visited []bool) {
	visited[i] = true
	for nb := g.adjacency[i]; nb != nil; nb = nb.next {
		if !visited[nb.id] {
			fmt.Println(i, nb.id)
			topologicalDfsHelper(g, nb.id, tindex, top, visited)
		}
	}
	top[*tindex] = i // push the deepest vertex to stack bottom
	*tindex--
}

// todo topological all
// utilize backtracking
func topologicalAll(g *graph) [][]int {
	tops := make([][]int, 0)
	//top := make([]int, g.n)
	return tops
}

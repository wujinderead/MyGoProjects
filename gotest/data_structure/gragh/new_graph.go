package gragh

import (
	conheap "container/heap"
)

func dijkstra(g *graph, src int) ([]int, []int) {
	shortest := make([]int, g.n)
	pred := make([]int, g.n)
	for i := range shortest {
		shortest[i] = 0xffffffff
		pred[i] = -1
	}
	inset := make([]bool, g.n)
	heap := edges([][2]int{{0, src}})
	shortest[src] = 0
	for heap.Len() > 0 {
		tmp := conheap.Pop(&heap).([2]int)
		_, ind := tmp[0], tmp[1]
		if inset[ind] {
			continue
		}
		inset[ind] = true
		for node := g.adjacency[ind]; node != nil; node = node.next {
			if shortest[ind]+node.weight < shortest[node.id] {
				shortest[node.id] = shortest[ind] + node.weight
				conheap.Push(&heap, [2]int{shortest[node.id], node.id})
				pred[node.id] = ind
			}
		}
	}
	return shortest, pred
}

type edges [][2]int // distance and index pair

func (h edges) Len() int {
	return len(h)
}

func (h edges) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h edges) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h *edges) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *edges) Pop() (x interface{}) {
	x = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func newFloydUndirected(n int, graph [][]int) [][]int {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
	}
	for i := range graph {
		dist[graph[i][0]][graph[i][1]] = graph[i][2]
		dist[graph[i][1]][graph[i][0]] = graph[i][2]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				d1, d2 := dist[i][k], dist[k][j]
				if d1 == 0 || d2 == 0 { // no route from i to k, or k to j
					continue
				}
				if dist[i][j] == 0 || d1+d2 < dist[i][j] {
					dist[i][j] = d1 + d2
					dist[j][i] = d1 + d2
				}
			}
		}
	}
	return dist
}

func newFloydDirected(n int, graph [][]int) [][]int {
	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
	}
	for i := range graph {
		dist[graph[i][0]][graph[i][1]] = graph[i][2]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				d1, d2 := dist[i][k], dist[k][j]
				if d1 == 0 || d2 == 0 { // no route from i to k, or k to j
					continue
				}
				if dist[i][j] == 0 || d1+d2 < dist[i][j] {
					dist[i][j] = d1 + d2
				}
			}
		}
	}
	return dist
}

func newTopologicalSortAll(g *graph) [][]int {
	indgree := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		for node := g.adjacency[i]; node != nil; node = node.next {
			indgree[node.id]++
		}
	}
	set := make(map[int]struct{})
	for i := range indgree {
		if indgree[i] == 0 {
			set[i] = struct{}{}
		}
	}
	tops := make([][]int, 0)
	top := make([]int, g.n)
	topAllHelper(g, top, indgree, 0, set, &tops)
	return tops
}

func topAllHelper(g *graph, top, indgree []int, ind int, set map[int]struct{}, tops *[][]int) {
	if ind == len(top) {
		tmp := make([]int, len(top))
		copy(tmp, top)
		*tops = append(*tops, tmp)
		return
	}
	curset := make([]int, 0, len(set))
	for i := range set {
		curset = append(curset, i)
	}
	for _, i := range curset {
		top[ind] = i
		delete(set, i)
		for node := g.adjacency[i]; node != nil; node = node.next {
			indgree[node.id]--
			if indgree[node.id] == 0 {
				set[node.id] = struct{}{}
			}
		}
		topAllHelper(g, top, indgree, ind+1, set, tops)

		// revert
		for node := g.adjacency[i]; node != nil; node = node.next {
			if indgree[node.id] == 0 {
				delete(set, node.id)
			}
			indgree[node.id]++
		}
		set[i] = struct{}{}
	}
}

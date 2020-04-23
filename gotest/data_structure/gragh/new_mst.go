package gragh

import (
	conheap "container/heap"
)

func newKruskal(g *graph) [][]int {
	heap := edgeHeap([][3]int{})
	disset := make([]int, g.n)
	mst := make([][]int, 0)
	for i := range g.adjacency {
		disset[i] = i
		for node := g.adjacency[i]; node != nil; node = node.next {
			if i < node.id {
				conheap.Push(&heap, [3]int{node.weight, i, node.id})
			}
		}
	}
	for heap.Len() > 0 {
		tmp := conheap.Pop(&heap).([3]int)
		_, i, j := tmp[0], tmp[1], tmp[2]
		ri, rj := root(disset, i), root(disset, j)
		if ri == rj {
			continue
		}
		disset[ri] = rj
		mst = append(mst, []int{i, j})
	}
	return mst
}

func root(disset []int, i int) int {
	for disset[i] != i {
		i = disset[i]
	}
	return i
}

type edgeHeap [][3]int // edge length, from index, to index

func (h edgeHeap) Len() int {
	return len(h)
}

func (h edgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h edgeHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
}

func (h *edgeHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *edgeHeap) Pop() (x interface{}) {
	x = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func newPrim(g *graph) [][]int {
	mst := make([][]int, 0)
	heap := edgeHeap([][3]int{}) // start from vertex 0
	inset := make([]bool, g.n)
	for node := g.adjacency[0]; node != nil; node = node.next {
		conheap.Push(&heap, [3]int{node.weight, 0, node.id})
	}
	for heap.Len() > 0 {
		tmp := conheap.Pop(&heap).([3]int)
		_, from, ind := tmp[0], tmp[1], tmp[2]
		if inset[ind] {
			continue
		}
		inset[ind] = true
		mst = append(mst, []int{from, ind})
		if len(mst) == g.n-1 {
			break
		}
		for node := g.adjacency[ind]; node != nil; node = node.next {
			conheap.Push(&heap, [3]int{node.weight, ind, node.id})
		}
	}
	return mst
}

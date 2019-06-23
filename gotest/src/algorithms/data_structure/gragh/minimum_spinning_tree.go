package gragh

import (
	"fmt"
	"math"
	"sort"
)

const INF = math.MaxInt8

// O(V²)
func mstPrimMatrix(g *matrix) []int {
	mstSet := make([]int, g.n)
	for i := range mstSet {
		mstSet[i] = -1 // mstSet[i]=-1 means vertex i not in mst set
	}
	minDist := make([]int, g.n)
	for i := range minDist {
		minDist[i] = INF
	}
	minDist[0] = 0 // use 0 as initial vertex, actually we can use arbitrary vertex to initiate
	for i := 0; i < g.n; i++ {
		min := INF
		minj := 0
		for j := 0; j < g.n; j++ {
			// not in mst set and has the minimum dist to the mst set
			if mstSet[j] == -1 && minDist[j] < min {
				min = minDist[j]
				minj = j
			}
		}
		// this value equals to initial vertex, initial vertex will has a self loop
		cor := 0
		// update minDist that go out minj
		for j := 0; j < g.n; j++ {
			dist := g.get(j, minj)
			if dist > 0 && dist < minDist[j] {
				minDist[j] = dist
			}
			// find the vertex in mst set that correspond to minj
			if dist == min && mstSet[j] != -1 {
				cor = j // edge [cor, minj] is added
			}
		}
		mstSet[minj] = cor
	}
	return mstSet
}

// O(ElogE)
// prim for adjacent list need a heap
func mstPrimAdjacent(g *graph) []int {
	mstSet := make([]int, g.n)
	indices := make([]int, g.n)
	mindist := make([]int, g.n)
	minfrom := make([]int, g.n)
	position := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		mstSet[i] = -1
		indices[i] = i
		position[i] = i
		mindist[i] = INF
	}
	h := &heap{indices, minfrom, mindist, position}
	h.mindist[0] = 0 // vertex 0 as initial vertex
	h.minfrom[0] = 0 // initial vertex need a self loop
	for num := 0; num < g.n; num++ {
		i := h.pop() // pop vertex with minimum distance with mst set
		fmt.Println(i, h)
		for nb := g.adjacency[i]; nb != nil; nb = nb.next {
			j := nb.id // update minimum distance between i and j
			if mstSet[j] == -1 && nb.weight < h.mindist[j] {
				h.mindist[j] = nb.weight
				h.minfrom[j] = i
				if !h.siftDown(h.position[j]) { // need to sift after modification
					h.siftUp(h.position[j])
				}
			}
		}
		// set mst set
		fmt.Println(i, h)
		mstSet[i] = h.minfrom[i]
	}
	return mstSet
}

// prim for adjacent list need a heap to get min length edge
// heap contains the vertex index, the heap is shifted base on
// min dist of vertex
type heap struct {
	indices  []int // indices[i] = j, vertex j is in heap position i
	minfrom  []int // minfrom[i] = j, distance between vertex i, j is minimum
	mindist  []int // mindist[i] = j, minimum distance for vertex i is j
	position []int // position[i] = j, vertex i is at indices position j
}

func (h *heap) siftDown(i int) bool {
	i0 := i
	for i <= len(h.indices)/2-1 {
		r := 2*i + 1
		if r+1 < len(h.indices) {
			if h.mindist[h.indices[r+1]] < h.mindist[h.indices[r]] {
				r = r + 1
			}
		}
		if h.mindist[h.indices[r]] < h.mindist[h.indices[i]] {
			h.position[h.indices[r]], h.position[h.indices[i]] = i, r // switch position
			h.indices[r], h.indices[i] = h.indices[i], h.indices[r]
			i = r
		} else {
			break
		}
	}
	return i > i0
}

func (h *heap) siftUp(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if h.mindist[h.indices[i]] < h.mindist[h.indices[p]] {
			h.position[h.indices[i]], h.position[h.indices[p]] = p, i // switch position
			h.indices[i], h.indices[p] = h.indices[p], h.indices[i]
			i = p
		} else {
			break
		}
	}
}

func (h *heap) pop() int {
	top := h.indices[0]
	h.indices[0] = h.indices[len(h.indices)-1]
	h.position[h.indices[0]] = 0
	h.indices = h.indices[:len(h.indices)-1]
	h.siftDown(0)
	return top
}

// O(ElogE)
func mstKruskalAdjacent(g *graph) [][]int {
	mstSet := make([][]int, 0)
	disjoint := make([]int, g.n)
	edges := make([][]int, 0)
	for i := 0; i < g.n; i++ {
		disjoint[i] = -1
		for nb := g.adjacency[i]; nb != nil; nb = nb.next {
			if i < nb.id {
				edges = append(edges, []int{i, nb.id, nb.weight})
			}
		}
	}
	fmt.Println(edges)
	sort.Sort(edgeSorter(edges))
	count := 0
	edgeind := 0
	fmt.Println(edges)
	for count < g.n-1 {
		curedge := edges[edgeind]
		v1 := curedge[0]
		v2 := curedge[1]
		s1 := find(disjoint, v1)
		s2 := find(disjoint, v2)
		fmt.Println(v1, v2, s1, s2, edgeind, count)
		if s1 == s2 { // same subset, include this edge will make a circle, discard it
			edgeind++
			continue
		} else { // different subset, union these two subsets
			union(disjoint, s1, s2)
			mstSet = append(mstSet, []int{v1, v2})
			edgeind++
			count++
		}
	}
	return mstSet
}

func find(disjoint []int, i int) int {
	if disjoint[i] == -1 {
		return i
	}
	return find(disjoint, disjoint[i])
}

func union(disjoint []int, v1, v2 int) {
	disjoint[v1] = v2
}

type edgeSorter [][]int

func (e edgeSorter) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e edgeSorter) Len() int {
	return len(e)
}

func (e edgeSorter) Less(i, j int) bool {
	return e[i][2] < e[j][2]
}

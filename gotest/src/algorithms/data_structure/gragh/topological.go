package gragh

import (
	"container/list"
	"fmt"
)

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

// utilize backtracking
func topologicalAll(g *graph) [][]int {
	tops := make([][]int, 0)
	top := make([]int, g.n)
	inDegree := make([]int, g.n)
	for i := range g.adjacency {
		for nb := g.adjacency[i]; nb != nil; nb = nb.next {
			inDegree[nb.id]++
		}
	}
	set := make(map[int]struct{}, g.n)
	for i := range g.adjacency {
		if inDegree[i] == 0 {
			set[i] = struct{}{}
		}
	}
	topi := 0
	topologicalStep(g, &tops, set, top, inDegree, &topi)
	return tops
}

func topologicalStep(g *graph, tops *[][]int, set map[int]struct{}, top, inDegree []int, topi *int) {
	if *topi == len(top) {
		newtop := make([]int, len(top))
		copy(newtop, top)
		*tops = append(*tops, newtop)
		return
	}
	curset := make([]int, 0, len(set))
	for i := range set {
		curset = append(curset, i)
	}
	for _, k := range curset {
		top[*topi] = k
		delete(set, k)
		*topi++
		for nb := g.adjacency[k]; nb != nil; nb = nb.next {
			inDegree[nb.id]--
			if inDegree[nb.id] == 0 {
				set[nb.id] = struct{}{}
			}
		}
		topologicalStep(g, tops, set, top, inDegree, topi)
		// revert for backtracking
		set[k] = struct{}{}
		*topi--
		for nb := g.adjacency[k]; nb != nil; nb = nb.next {
			if inDegree[nb.id] == 0 {
				delete(set, nb.id)
			}
			inDegree[nb.id]++
		}
	}
}

// longest path for DAG. longest path in a weighted graph is a NP-hard problem,
// but if the graph is a DAG, we can solve it in O(N) time use topological sort.
func longestPathDag(g *graph, src, dst int) int {
	// first topological sort
	top := topologicalQueue(g)

	// est[i] to store the longest path from src to i
	est := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		est[i] = NINF
	}
	est[src] = 0
	for i := 0; i < len(top); i++ {
		from := top[i]
		if from == dst {
			break // no need to process succeeded vertices in topological sequence as they cannot go back to 'dst'
		}
		if est[from] == NINF {
			continue // skip those vertices that can not reach from src
		}
		for nb := g.adjacency[from]; nb != nil; nb = nb.next {
			to := nb.id
			wei := nb.weight
			if est[to] < est[from]+wei {
				est[to] = est[from] + wei // update the longest between 'from' and 'to'
			}
		}
	}
	return est[dst] // return NINF means src cannot reach dst
}

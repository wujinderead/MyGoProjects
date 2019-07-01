package gragh

import (
	"fmt"
	"testing"
)

func TestDetectCycleUndirected(t *testing.T) {
	g := newUndirectedGraph(8, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4},
		{4, 5}, {5, 0}, {0, 6}, {1, 7}})
	fmt.Println(detectCycleUndirectedDfs(g))
	g = newUndirectedGraph(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4},
		{0, 3}})
	fmt.Println(detectCycleUndirectedDfs(g))
}

func TestDetectCycleDirected(t *testing.T) {
	g := newDirectedGraph(8, [][]int{{1, 0}, {0, 2}, {0, 6}, {2, 3},
		{3, 0}, {1, 4}, {4, 5}, {5, 3}, {5, 5}})
	fmt.Println(detectCycleDirectedDfs(g))
}

func TestDetectCycleUnionFind(t *testing.T) {
	g := newUndirectedGraph(8, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4},
		{4, 5}, {5, 0}, {0, 6}, {1, 7}})
	g.print()
	fmt.Println(detectCycleUnionFind(g))
}

package gragh

import (
	"fmt"
	"testing"
)

func TestDetectCycleUndirected(t *testing.T) {
	g := newUndirectedGraph(8, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4},
		{4, 5}, {5, 0}, {0, 6}, {1, 7}})
	fmt.Println(detectCycleUndirectedDfs(g))
}

func TestDetectCycleDirected(t *testing.T) {
	g := newUndirectedGraph(8, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4},
		{4, 5}, {5, 0}, {0, 6}, {1, 7}})
	fmt.Println(detectCycleUndirectedDfs(g))
}

func TestDetectCycleUnionFind(t *testing.T) {
	g := newUndirectedGraph(8, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4},
		{4, 5}, {5, 0}, {0, 6}, {1, 7}})
	fmt.Println(detectCycleUndirectedDfs(g))
}

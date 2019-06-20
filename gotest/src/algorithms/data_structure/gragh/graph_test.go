package gragh

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := newGraph(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {2, 4}, {4, 5}, {2, 5}})
	g.print()
}

func TestGraphTopologicalQueue(t *testing.T) {
	g := newGraph(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {2, 4}, {4, 5}, {2, 5}})
	fmt.Println(topologicalQueue(g))
}

func TestGraphTopologicalDfs(t *testing.T) {
	g := newGraph(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {2, 4}, {4, 5}, {2, 5}})
	fmt.Println(topologicalDfs(g))
}

func TestGraphTopologicalAll(t *testing.T) {
	g := newGraph(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {4, 5}, {2, 5}})
	allTops := topologicalAll(g)
	for i := range allTops {
		fmt.Println(allTops[i])
	}
}

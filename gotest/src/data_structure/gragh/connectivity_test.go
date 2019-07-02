package gragh

import (
	"fmt"
	"testing"
)

func TestStronglyConnected(t *testing.T) {
	g := newDirectedGraph(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {4, 2}, {2, 4}})
	g.print()
	fmt.Println(checkStronglyConnected(g))
	g = newDirectedGraph(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 0}, {4, 2}})
	g.print()
	fmt.Println(checkStronglyConnected(g))
}

func TestStronglyConnectedComponents(t *testing.T) {
	g := newDirectedGraph(4, [][]int{{0, 1}, {1, 2}, {2, 3}})
	fmt.Println(findStronglyConnectedComponents(g))
	g = newDirectedGraph(6, [][]int{{0, 2}, {1, 0}, {2, 1}, {0, 3}, {3, 4}, {1, 5}})
	fmt.Println(findStronglyConnectedComponents(g))
}

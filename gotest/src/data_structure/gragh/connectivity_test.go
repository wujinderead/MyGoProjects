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

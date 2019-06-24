package gragh

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := newDirectedGraph(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {2, 4}, {4, 5}, {2, 5}})
	g.print()
	fmt.Println()
	g = newDirectedGraph(6, [][]int{{0, 2, 1}, {1, 3, 2}, {3, 2, 3},
		{3, 4, 4}, {2, 4, 5}, {4, 5, 6}, {2, 5, 7}})
	g.print()
	fmt.Println()
	g = newUndirectedGraph(6, [][]int{{0, 2, 1}, {1, 3, 2}, {3, 2, 3},
		{3, 4, 4}, {2, 4, 5}, {4, 5, 6}, {2, 5, 7}})
	g.print()
}

func TestNewGraphMatrix(t *testing.T) {
	g := newDirectedMatrix(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {2, 4}, {4, 5}, {2, 5}})
	g.print()
	fmt.Println()
	g = newUndirectedMatrix(6, [][]int{{0, 2, 1}, {1, 3, 2}, {3, 2, 3},
		{3, 4, 4}, {2, 4, 5}, {4, 5, 6}, {2, 5, 7}})
	g.print()
	fmt.Println()
	g = newDirectedMatrix(6, [][]int{{0, 2, 1}, {1, 3, 2}, {3, 2, 3},
		{3, 4, 4}, {2, 4, 5}, {4, 5, 6}, {2, 5, 7}})
	g.print()
}

package gragh

import (
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := newGraph(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {2, 4}, {4, 5}, {2, 5}})
	g.print()
}

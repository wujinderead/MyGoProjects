package gragh

import (
	"fmt"
	"testing"
)

func TestMaxFlowFordFulkerson(t *testing.T) {
	g := newDirectedMatrix(6, [][]int{{0, 1, 16},
		{0, 2, 13},
		{1, 2, 10},
		{2, 1, 4},
		{1, 3, 12},
		{2, 4, 14},
		{3, 2, 9},
		{4, 3, 7},
		{3, 5, 20},
		{4, 5, 4},
	})
	fmt.Println(maxFlowFordFulkerson(g, 0, 5))
}

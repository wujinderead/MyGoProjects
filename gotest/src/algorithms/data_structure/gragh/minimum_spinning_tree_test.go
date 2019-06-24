package gragh

import (
	"fmt"
	"testing"
)

func TestMstPrimMatrix(t *testing.T) {
	g := newUndirectedMatrix(9, [][]int{
		{0, 1, 4},
		{0, 7, 8},
		{1, 7, 11},
		{1, 2, 8},
		{2, 8, 2},
		{7, 8, 7},
		{6, 7, 1},
		{6, 8, 6},
		{2, 3, 7},
		{2, 5, 4},
		{5, 6, 2},
		{3, 5, 14},
		{3, 4, 9},
		{4, 5, 10},
	})
	g.print()
	mst := mstPrimMatrix(g)
	for i := range mst {
		if mst[i] == i { // ignore self loop
			continue
		}
		fmt.Println(i, "->", mst[i])
	}
}

func TestMstPrimAdjacent(t *testing.T) {
	g := newUndirectedGraph(9, [][]int{
		{0, 1, 4},
		{0, 7, 8},
		{1, 7, 11},
		{1, 2, 8},
		{2, 8, 2},
		{7, 8, 7},
		{6, 7, 1},
		{6, 8, 6},
		{2, 3, 7},
		{2, 5, 4},
		{5, 6, 2},
		{3, 5, 14},
		{3, 4, 9},
		{4, 5, 10},
	})
	g.print()
	mst := mstPrimAdjacent(g)
	for i := range mst {
		fmt.Println(i, "->", mst[i])
	}
}

func TestMstKruskalAdjacent(t *testing.T) {
	g := newUndirectedGraph(9, [][]int{
		{0, 1, 4},
		{0, 7, 8},
		{1, 7, 11},
		{1, 2, 8},
		{2, 8, 2},
		{7, 8, 7},
		{6, 7, 1},
		{6, 8, 6},
		{2, 3, 7},
		{2, 5, 4},
		{5, 6, 2},
		{3, 5, 14},
		{3, 4, 9},
		{4, 5, 10},
	})
	g.print()
	mst := mstKruskalAdjacent(g)
	for i := range mst {
		fmt.Println(mst[i][0], "->", mst[i][1])
	}
}

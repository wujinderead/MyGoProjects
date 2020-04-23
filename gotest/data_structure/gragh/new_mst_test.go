package gragh

import (
	"fmt"
	"testing"
)

func TestMstKruskalNew(t *testing.T) {
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
	mst := mstKruskalAdjacent(g)
	for i := range mst {
		fmt.Println(i, "->", mst[i])
	}
	fmt.Println()
	mst1 := newKruskal(g)
	for i := range mst {
		fmt.Println(mst1[i])
	}
}

func TestMstPrimNew(t *testing.T) {
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
	mst := mstPrimAdjacent(g)
	for i := range mst {
		fmt.Println(i, "->", mst[i])
	}
	fmt.Println()
	mst1 := newPrim(g)
	for i := range mst1 {
		fmt.Println(mst1[i])
	}
}

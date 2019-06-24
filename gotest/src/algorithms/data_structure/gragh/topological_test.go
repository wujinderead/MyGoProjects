package gragh

import (
	"fmt"
	"testing"
)

func TestGraphTopologicalQueue(t *testing.T) {
	g := newDirectedGraph(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {2, 4}, {4, 5}, {2, 5}})
	fmt.Println(topologicalQueue(g))
}

func TestGraphTopologicalDfs(t *testing.T) {
	g := newDirectedGraph(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {2, 4}, {4, 5}, {2, 5}})
	fmt.Println(topologicalDfs(g))
}

func TestGraphTopologicalAll(t *testing.T) {
	for _, g := range []*graph{
		newDirectedGraph(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {4, 5}, {2, 5}}),
		newDirectedGraph(4, [][]int{{0, 3}, {1, 2}, {2, 3}}),
		newDirectedGraph(6, [][]int{{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1}}),
	} {
		allTops := topologicalAll(g)
		for i := range allTops {
			fmt.Println(allTops[i])
		}
		fmt.Println()
	}
}

func TestLongestPathDag(t *testing.T) {
	g := newDirectedGraph(6, [][]int{
		{0, 1, 5},
		{0, 2, 3},
		{1, 3, 6},
		{1, 2, 2},
		{2, 4, 4},
		{2, 5, 2},
		{2, 3, 7},
		{3, 5, 1},
		{3, 4, -1},
		{4, 5, -2},
	})
	fmt.Println(longestPathDag(g, 0, 0))
	fmt.Println(longestPathDag(g, 0, 1))
	fmt.Println(longestPathDag(g, 0, 2))
	fmt.Println(longestPathDag(g, 0, 3))
	fmt.Println(longestPathDag(g, 0, 4))
	fmt.Println(longestPathDag(g, 0, 5))
	fmt.Println(longestPathDag(g, 1, 0))
	fmt.Println(longestPathDag(g, 1, 1))
	fmt.Println(longestPathDag(g, 1, 2))
	fmt.Println(longestPathDag(g, 1, 3))
	fmt.Println(longestPathDag(g, 1, 4))
	fmt.Println(longestPathDag(g, 1, 5))
}

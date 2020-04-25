package gragh

import (
	"fmt"
	"testing"
)

func TestDijkstraAdjacentUndirectedNew(t *testing.T) {
	fmt.Println("for contrast:")

	fmt.Println("new test:")
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
	src := 0
	dist, pred := dijkstra(g, src)
	dist1, pred1 := dijkstraAdjacent(g, src)
	fmt.Println("new:", dist, pred)
	fmt.Println("con:", dist1, pred1)

	src = 2
	dist, pred = dijkstra(g, src)
	dist1, pred1 = dijkstraAdjacent(g, src)
	fmt.Println("new:", dist, pred)
	fmt.Println("con:", dist1, pred1)

	src = 8
	dist, pred = dijkstra(g, src)
	dist1, pred1 = dijkstraAdjacent(g, src)
	fmt.Println("new:", dist, pred)
	fmt.Println("con:", dist1, pred1)
}

func TestDijkstraAdjacentDirectedNew(t *testing.T) {
	fmt.Println("for contrast:")

	fmt.Println("new test:")
	g := newDirectedGraph(9, [][]int{
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
	src := 0
	dist, pred := dijkstra(g, src)
	dist1, pred1 := dijkstraAdjacent(g, src)
	fmt.Println("new:", dist, pred)
	fmt.Println("con:", dist1, pred1)

	src = 2
	dist, pred = dijkstra(g, src)
	dist1, pred1 = dijkstraAdjacent(g, src)
	fmt.Println("new:", dist, pred)
	fmt.Println("con:", dist1, pred1)

	src = 8
	dist, pred = dijkstra(g, src)
	dist1, pred1 = dijkstraAdjacent(g, src)
	fmt.Println("new:", dist, pred)
	fmt.Println("con:", dist1, pred1)
}

func TestFloydUndirectedNew(t *testing.T) {
	n := 9
	edges := [][]int{
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
	}
	g := newUndirectedMatrix(n, edges)
	dist := floydWarshall(g)
	fmt.Println("\nshortest path:")
	for i := range dist {
		fmt.Println(i, ":", dist[i])
	}

	// new test
	dist1 := newFloydUndirected(n, edges)
	for i := range dist1 {
		fmt.Println(i, ":", dist1[i])
	}

	for i := 0; i < len(dist); i++ {
		for j := 0; j < len(dist); j++ {
			if dist[i][j] != dist1[i][j] {
				t.Fail()
			}
		}
	}
}

func TestFloydDirectedNew(t *testing.T) {
	n := 9
	edges := [][]int{
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
	}
	g := newDirectedMatrix(n, edges)
	dist := floydWarshall(g)
	fmt.Println("\nshortest path:")
	for i := range dist {
		fmt.Println(i, ":", dist[i])
	}

	// new test
	dist1 := newFloydDirected(n, edges)
	for i := range dist1 {
		fmt.Println(i, ":", dist1[i])
	}

	for i := 0; i < len(dist); i++ {
		for j := 0; j < len(dist); j++ {
			if dist[i][j] != 127 && dist[i][j] != dist1[i][j] {
				t.Fail()
			}
		}
	}
}

func TestTopologicalAllNew(t *testing.T) {
	for _, g := range []*graph{
		newDirectedGraph(6, [][]int{{0, 2}, {1, 3}, {3, 2}, {3, 4}, {4, 5}, {2, 5}}),
		newDirectedGraph(4, [][]int{{0, 3}, {1, 2}, {2, 3}}),
		newDirectedGraph(6, [][]int{{5, 2}, {5, 0}, {4, 0}, {4, 1}, {2, 3}, {3, 1}}),
	} {
		allTops := topologicalAll(g)
		fmt.Println(allTops)
		newAll := newTopologicalSortAll(g)
		fmt.Println(newAll)
		fmt.Println()
	}
}

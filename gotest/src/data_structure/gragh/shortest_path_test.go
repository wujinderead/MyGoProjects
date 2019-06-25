package gragh

import (
	"fmt"
	"testing"
)

func TestDijkstraMatrixUndirected(t *testing.T) {
	// undirected
	fmt.Println("undirected graph:")
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
	src := 0
	dist, pred := dijkstraMatrix(g, src)
	fmt.Println("\nshortest path:")
	for i := range dist {
		fmt.Println(src, "->", i, ",", dist[i])
	}
	fmt.Println("\npath:")
	for i := range pred {
		fmt.Println(pred[i], "->", i)
	}
}

func TestDijkstraMatrixDirected(t *testing.T) {
	// directed
	fmt.Println("directed graph:")
	g := newDirectedMatrix(9, [][]int{
		{0, 1, 4},
		{0, 7, 8},
		{1, 7, 11},
		{1, 2, 8},
		{2, 8, 2},
		{8, 7, 7},
		{7, 6, 1},
		{8, 6, 6},
		{3, 2, 7},
		{5, 2, 4},
		{6, 5, 2},
		{3, 5, 14},
		{4, 3, 9},
		{5, 4, 10},
	})
	g.print()
	src := 0
	dist, pred := dijkstraMatrix(g, src)
	fmt.Println("\nshortest path 0:")
	for i := range dist {
		fmt.Println(src, "->", i, ",", dist[i])
	}
	fmt.Println("\npath 0:")
	for i := range pred {
		fmt.Println(pred[i], "->", i)
	}
	src = 4
	dist, pred = dijkstraMatrix(g, src)
	fmt.Println("\nshortest path 4:")
	for i := range dist {
		fmt.Println(src, "->", i, ",", dist[i])
	}
	fmt.Println("\npath 4:")
	for i := range pred {
		fmt.Println(pred[i], "->", i)
	}
}

func TestDijkstraAdjacentUndirected(t *testing.T) {
	// undirected
	fmt.Println("undirected graph:")
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
	dist, pred := dijkstraAdjacent(g, src)
	fmt.Println("\nshortest path:")
	for i := range dist {
		fmt.Println(src, "->", i, ",", dist[i])
	}
	fmt.Println("\npath:")
	for i := range pred {
		fmt.Println(pred[i], "->", i)
	}
}

func TestDijkstraAdjacentDirected(t *testing.T) {
	// directed
	fmt.Println("directed graph:")
	g := newDirectedGraph(9, [][]int{
		{0, 1, 4},
		{0, 7, 8},
		{1, 7, 11},
		{1, 2, 8},
		{2, 8, 2},
		{8, 7, 7},
		{7, 6, 1},
		{8, 6, 6},
		{3, 2, 7},
		{5, 2, 4},
		{6, 5, 2},
		{3, 5, 14},
		{4, 3, 9},
		{5, 4, 10},
	})
	g.print()
	src := 0
	dist, pred := dijkstraAdjacent(g, src)
	fmt.Println("\nshortest path 0:")
	for i := range dist {
		fmt.Println(src, "->", i, ",", dist[i])
	}
	fmt.Println("\npath 0:")
	for i := range pred {
		fmt.Println(pred[i], "->", i)
	}
	src = 4
	dist, pred = dijkstraAdjacent(g, src)
	fmt.Println("\nshortest path 4:")
	for i := range dist {
		fmt.Println(src, "->", i, ",", dist[i])
	}
	fmt.Println("\npath 4:")
	for i := range pred {
		fmt.Println(pred[i], "->", i)
	}
}

func TestFloydWarshallUndirected(t *testing.T) {
	// undirected
	fmt.Println("undirected graph:")
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
	dist := floydWarshall(g)
	fmt.Println("\nshortest path:")
	for i := range dist {
		fmt.Println(i, ":", dist[i])
	}
}

func TestFloydWarshallDirected(t *testing.T) {
	// directed
	fmt.Println("directed graph:")
	g := newDirectedMatrix(9, [][]int{
		{0, 1, 4},
		{0, 7, 8},
		{1, 7, 11},
		{1, 2, 8},
		{2, 8, 2},
		{8, 7, 7},
		{7, 6, 1},
		{8, 6, 6},
		{3, 2, 7},
		{5, 2, 4},
		{6, 5, 2},
		{3, 5, 14},
		{4, 3, 9},
		{5, 4, 10},
	})
	g.print()
	dist := floydWarshall(g)
	fmt.Println("\nshortest path:")
	for i := range dist {
		fmt.Println(i, ":", dist[i])
	}
}

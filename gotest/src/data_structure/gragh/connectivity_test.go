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

func TestFindArticulationPoints(t *testing.T) {
	g := newUndirectedGraph(7, [][]int{{0, 3}, {0, 1}, {1, 2}, {2, 3}, {3, 4}, {2, 5}, {5, 6}})
	fmt.Println(findArticulationPoints(g))
	g = newUndirectedGraph(7, [][]int{{0, 3}, {0, 1}, {1, 2}, {2, 3}, {3, 4}, {2, 5}, {5, 6}, {6, 1}})
	fmt.Println(findArticulationPoints(g))
	g = newUndirectedGraph(4, [][]int{{0, 1}, {1, 2}, {2, 3}})
	fmt.Println(findArticulationPoints(g))
}

func TestSccTarjan(t *testing.T) {
	g := newDirectedGraph(4, [][]int{{0, 1}, {1, 2}, {2, 3}})
	fmt.Println(findAllSccTarjan(g))
	g = newDirectedGraph(6, [][]int{{0, 2}, {1, 0}, {2, 1}, {0, 3}, {3, 4}, {1, 5}})
	fmt.Println(findAllSccTarjan(g))
	g = newDirectedGraph(7, [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}, {1, 4}, {1, 6}, {3, 5}, {4, 5}})
	fmt.Println(findAllSccTarjan(g))
	g = newDirectedGraph(11, [][]int{{0, 1}, {0, 3}, {1, 2}, {1, 4}, {2, 0}, {2, 6}, {3, 2}, {4, 5},
		{4, 6}, {5, 6}, {5, 7}, {5, 8}, {5, 9}, {6, 4}, {7, 9}, {8, 9}, {9, 8}})
	fmt.Println(findAllSccTarjan(g))
}

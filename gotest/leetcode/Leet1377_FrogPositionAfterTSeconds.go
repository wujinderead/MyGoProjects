package main

import (
	"container/list"
	"fmt"
)

// Given an undirected tree consisting of n vertices numbered from 1 to n. A frog
// starts jumping from the vertex 1. In one second, the frog jumps from its current
// vertex to another unvisited vertex if they are directly connected. The frog can
// not jump back to a visited vertex. In case the frog can jump to several vertices
// it jumps randomly to one of them with the same probability, otherwise, when the
// frog can not jump to any unvisited vertex it jumps forever on the same vertex.
// The edges of the undirected tree are given in the array edges, where
// edges[i] = [fromi, toi] means that exists an edge connecting directly the
// vertices fromi and toi.
// Return the probability that after t seconds the frog is on the vertex target.
// Example 1:
//   Input: n = 7, edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]], t = 2, target = 4
//   Output: 0.16666666666666666
//   Explanation: The figure above shows the given graph. The frog starts at vertex
//     1, jumping with 1/3 probability to the vertex 2 after second 1 and then jumping
//     with 1/2 probability to vertex 4 after second 2. Thus the probability for the
//     frog is on the vertex 4 after 2 seconds is 1/3 * 1/2 = 1/6 = 0.16666666666666666.
//               1
//           /   |   \
//          2    3   7
//     	   / \   |
//        4  6   5
// Example 2:
//   Input: n = 7, edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]], t = 1, target = 7
//   Output: 0.3333333333333333
//   Explanation: The figure above shows the given graph. The frog starts at vertex
//     1, jumping with 1/3 = 0.3333333333333333 probability to the vertex 7 after
//     second 1.
// Example 3:
//  Input: n = 7, edges = [[1,2],[1,3],[1,7],[2,4],[2,6],[3,5]], t = 20, target = 6
//  Output: 0.16666666666666666
// Constraints:
//   1 <= n <= 100
//   edges.length == n-1
//   edges[i].length == 2
//   1 <= edges[i][0], edges[i][1] <= n
//   1 <= t <= 50
//   1 <= target <= n
//   Answers within 10^-5 of the actual value will be accepted as correct.

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	mapp := make(map[int][]int)
	for i := range edges {
		from, to := edges[i][0], edges[i][1]
		succ, ok := mapp[from]
		if ok {
			mapp[from] = append(succ, to)
		} else {
			mapp[from] = []int{to}
		}
		succ, ok = mapp[to]
		if ok {
			mapp[to] = append(succ, from)
		} else {
			mapp[to] = []int{from}
		}
	}
	mapp[1] = append(mapp[1], 0) // add a dummy 0 as 1's predecessor
	prob := make([]float64, n+1)
	time := make([]int, n+1)
	visited := make([]bool, n+1)
	queue := list.New()
	queue.PushBack(1)
	prob[1] = 1
	time[1] = 0
	visited[0] = true
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(int)
		succ, _ := mapp[cur]
		visited[cur] = true
		if cur == target {
			if len(succ) > 1 && t > time[cur] {
				return 0
			} else {
				return prob[cur]
			}
		}
		if time[cur] == t && time[target] == 0 {
			return 0
		}
		for _, v := range succ {
			if !visited[v] {
				prob[v] = prob[cur] / float64(len(succ)-1) // minus 1 to exclude the predecessor
				time[v] = time[cur] + 1
				queue.PushBack(v)
			}
		}
	}
	return prob[target]
}

func main() {
	fmt.Println("a", frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 2, 4))
	fmt.Println("b", frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 1, 7))
	fmt.Println("c", frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 20, 6))
	fmt.Println("d", frogPosition(7, [][]int{{1, 2}, {1, 3}, {1, 7}, {2, 4}, {2, 6}, {3, 5}}, 1, 4))
	fmt.Println("e", frogPosition(3, [][]int{{2, 1}, {3, 2}}, 1, 2))
	fmt.Println("e", frogPosition(3, [][]int{{2, 1}, {3, 2}}, 2, 2))
	fmt.Println("f", frogPosition(1, [][]int{}, 0, 1))
	fmt.Println("g", frogPosition(1, [][]int{}, 2, 1))
	tree := [][]int{{1, 2}, {3, 1}, {4, 1},
		{2, 5}, {6, 2}, {2, 7},
		{3, 8}, {3, 9}, {3, 10},
		{11, 4}, {12, 4}, {4, 13},
	}
	fmt.Println("h", frogPosition(13, tree, 0, 1))
	fmt.Println("h", frogPosition(13, tree, 1, 1))
	fmt.Println("i", frogPosition(13, tree, 1, 2))
	fmt.Println("j", frogPosition(13, tree, 1, 3))
	fmt.Println("k", frogPosition(13, tree, 1, 5))
	fmt.Println("l", frogPosition(13, tree, 2, 7))
	fmt.Println("m", frogPosition(13, tree, 2, 13))
	fmt.Println("m", frogPosition(13, tree, 2, 4))
	fmt.Println("n", frogPosition(13, tree, 3, 1))
	fmt.Println("n", frogPosition(13, tree, 3, 2))
	fmt.Println("o", frogPosition(13, tree, 3, 13))
	fmt.Println("o", frogPosition(13, tree, 4, 12))
	fmt.Println("p", frogPosition(8, [][]int{{2, 1}, {3, 2}, {4, 1}, {5, 1}, {6, 4}, {7, 1}, {8, 7}}, 7, 7))
}

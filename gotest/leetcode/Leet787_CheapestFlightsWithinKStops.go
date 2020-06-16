package main

import (
	"container/heap"
    "fmt"
)

// https://leetcode.com/problems/cheapest-flights-within-k-stops/

// There are n cities connected by m flights. Each flight starts from city u and arrives at v with a price w.
// Now given all the cities and flights, together with starting city src and the destination dst, your task 
// is to find the cheapest price from src to dst with up to k stops. If there is no such route, output -1.
// Example 1:
//   Input: 
//     n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//     src = 0, dst = 2, k = 1
//   Output: 200
//   Explanation: 
//     The graph looks like this:
//     The cheapest price from city 0 to city 2 with at most 1 stop costs 200, as marked red in the picture.
// Example 2:
//   Input: 
//     n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
//     src = 0, dst = 2, k = 0
//   Output: 500
//   Explanation: 
//     The graph looks like this:
//     The cheapest price from city 0 to city 2 with at most 0 stop costs 500, as marked blue in the picture.
// Constraints:
//   The number of nodes n will be in range [1, 100], with nodes labeled from 0 to n - 1.
//   The size of flights will be in range [0, n * (n - 1) / 2].
//   The format of each flight will be (src, dst, price).
//   The price of each flight will be in the range [1, 10000].
//   k is in the range of [0, n - 1].
//   There will not be any duplicated flights or self cycles.

// use a heap to track the minimal dist to src (like dijkstra), with respect to step
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	// make a graph
    graph := make([][][2]int, n)  // graph[src] = []{{dst1, price1}, ...}
    for _, v := range flights {
    	graph[v[0]] = append(graph[v[0]], [2]int{v[1], v[2]})
    }

    // add src edges to heap
    h := tuples(make([][3]int, 0, len(graph[src])))
    for _, v := range graph[src] {   
    	h = append(h, [3]int{v[0], v[1], 0})   // (dst, shortest dist to src, step) tuple
    }

    // min heap to get min dist
    for h.Len()>0 {
    	tmp := heap.Pop(&h).([3]int)
    	d, p, step := tmp[0], tmp[1], tmp[2]
    	if d==dst {
    		return p
    	}
    	for _, v := range graph[d] {
    		if step+1<=K {
    			heap.Push(&h, [3]int{v[0], v[1]+p, step+1})
    		}
    	}
    } 
    return -1
}

type tuples [][3]int   (dst, shortest dist to src, step) tuple

func (t tuples) Len() int {
	return len(t)
}

func (t tuples) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t tuples) Less(i, j int) bool {
	return t[i][1] < t[j][1]   // sort by price
}

func (t *tuples) Push(x interface{}) {
	*t = append(*t, x.([3]int))
}

func (t *tuples) Pop() interface{} {
	x := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return x
}

func main() {
	fmt.Println(findCheapestPrice(3, [][]int{{0,1,100},{1,2,100},{0,2,500}}, 0, 2, 1))
	fmt.Println(findCheapestPrice(3, [][]int{{0,1,100},{1,2,100},{0,2,500}}, 0, 2, 0))
}
package main

import "fmt"

// https://leetcode.com/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/

// There are n cities numbered from 0 to n-1 and n-1 roads such that there is only 
// one way to travel between two different cities (this network form a tree). Last year, 
// The ministry of transport decided to orient the roads in one direction because they are too narrow. 
// Roads are represented by connections where connections[i] = [a, b] represents a road from city a to b. 
// This year, there will be a big event in the capital (city 0), and many people want to travel to this city. 
// Your task consists of reorienting some roads such that each city can visit the city 0. 
// Return the minimum number of edges changed. 
// It's guaranteed that each city can reach the city 0 after reorder. 
// Example 1: 
//   Input: n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
//   Output: 3
//   Explanation: Change the direction of edges show in red such that each node can
//     reach the node 0 (capital). 
// Example 2: 
//       0 <-- 1 --> 2 <-- 3 --> 4
//                x           x						
//   Input: n = 5, connections = [[1,0],[1,2],[3,2],[3,4]]
//   Output: 2
//   Explanation: Change the direction of edges show in red such that each node can
//     reach the node 0 (capital). 
// Example 3:  
//   Input: n = 3, connections = [[1,0],[2,0]]
//   Output: 0
// Constraints: 
//   2 <= n <= 5 * 10^4 
//   connections.length == n-1 
//   connections[i].length == 2 
//   0 <= connections[i][0], connections[i][1] <= n-1 
//   connections[i][0] != connections[i][1] 

func minReorder(n int, connections [][]int) int {
    graph := make([][][2]int, n)
    count := 0
    for _, v := range connections {
    	graph[v[0]] = append(graph[v[0]], [2]int{v[1], 0})   // 0 means origin edge
    	graph[v[1]] = append(graph[v[1]], [2]int{v[0], 1})   // 1 means reverse edge
    }
    visited := make([]bool, n)
    visit(0, graph, visited, &count)
    return count
}

func visit(i int, graph [][][2]int, visited []bool, count *int) {
	visited[i] = true
	for _, v := range graph[i] {
		if !visited[v[0]] {
			if v[1] == 0 {       // reverse the original edge
				*count += 1
			}
			visit(v[0], graph, visited, count)
		}
	}
}

func main() {
	fmt.Println(minReorder(6, [][]int{{0,1}, {1,3}, {2,3}, {4,0}, {4,5}}), 3)
	fmt.Println(minReorder(5, [][]int{{1,0}, {1,2}, {3,2}, {3,4}}), 2)
	fmt.Println(minReorder(3, [][]int{{1,0}, {2,0}}), 0)
}
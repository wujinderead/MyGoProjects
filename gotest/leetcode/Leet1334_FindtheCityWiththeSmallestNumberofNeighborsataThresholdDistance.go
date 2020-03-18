package main

import "fmt"

// https://leetcode.com/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/

//There are n cities numbered from 0 to n-1. Given the array edges where
// edges[i] = [fromi, toi, weighti] represents a bidirectional and weighted edge
// between cities fromi and toi, and given the integer distanceThreshold.
// Return the city with the smallest number of cities that are reachable through
// some path and whose distance is at most distanceThreshold,
// If there are multiple such cities, return the city with the greatest number.
// Notice that the distance of a path connecting cities i and j is equal to the
// sum of the edges' weights along that path.
// Example 1:
//   Input: n = 4, edges = [[0,1,3],[1,2,1],[1,3,4],[2,3,1]], distanceThreshold = 4
//   Output: 3
//   Explanation: The figure above describes the graph.
//           3
//     ⓪ ------- ①              0=>1  3
//             /  |              0=>2  4
//         4 /    | 1            0=>3  5
//         /      |              1=>2  1
//       /        |              1=>3  2
//     ③ ------- ②              2=>3  1
//           1
//
//     The neighboring cities at a distanceThreshold = 4 for each city are:
//       City 0 -> [City 1, City 2]
//       City 1 -> [City 0, City 2, City 3]
//       City 2 -> [City 0, City 1, City 3]
//       City 3 -> [City 1, City 2]
//     Cities 0 and 3 have 2 neighboring cities at a distanceThreshold = 4,
//     but we have to return city 3 since it has the greatest number.
// Example 2:
//   Input: n = 5, edges = [[0,1,2],[0,4,8],[1,2,3],[1,4,2],[2,3,1],[3,4,1]], distanceThreshold = 2
//   Output: 0
//   Explanation: The figure above describes the graph.
//     The neighboring cities at a distanceThreshold = 2 for each city are:
//       City 0 -> [City 1]
//       City 1 -> [City 0, City 4]
//       City 2 -> [City 3, City 4]
//       City 3 -> [City 2, City 4]
//       City 4 -> [City 1, City 2, City 3]
//     The city 0 has 1 neighboring city at a distanceThreshold = 2.
// Constraints:
//   2 <= n <= 100
//   1 <= edges.length <= n * (n - 1) / 2
//   edges[i].length == 3
//   0 <= fromi < toi < n
//   1 <= weighti, distanceThreshold <= 10^4
//   All pairs (fromi, toi) are distinct.

// use floyd-warshall algorithm to get the nearest distance of city pairs
func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	dis := make([]int, n*n)
	for i := range edges {
		set2d(dis, edges[i][0], edges[i][1], edges[i][2], n)
		set2d(dis, edges[i][1], edges[i][0], edges[i][2], n)
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				d1 := get2d(dis, i, k, n)
				d2 := get2d(dis, k, j, n)
				dd := get2d(dis, i, j, n)
				if d1 > 0 && d2 > 0 && (d1+d2 < dd || dd == 0) {
					set2d(dis, i, j, d1+d2, n)
					set2d(dis, j, i, d1+d2, n)
				}
			}
		}
	}
	minnum := n
	minind := n
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			d := get2d(dis, i, j, n)
			if d > 0 && d <= distanceThreshold {
				count++
			}
		}
		if count <= minnum {
			minnum = count
			minind = i
		}
	}
	return minind
}

//go:inline
func get2d(arr []int, i, j, col int) int {
	return arr[i*col+j]
}

//go:inline
func set2d(arr []int, i, j, v, col int) {
	arr[i*col+j] = v
}

func main() {
	fmt.Println(findTheCity(4, [][]int{{0, 1, 3}, {1, 2, 1}, {1, 3, 4}, {2, 3, 1}}, 4))
	fmt.Println(findTheCity(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2))
}

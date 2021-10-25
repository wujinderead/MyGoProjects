package main

import "fmt"

// https://leetcode.com/problems/minimum-score-triangulation-of-polygon/

// You have a convex n-sided polygon where each vertex has an integer value. You are given an integer array
// values where values[i] is the value of the ith vertex (i.e., clockwise order).
// You will triangulate the polygon into n - 2 triangles. For each triangle, the value of that triangle is
// the product of the values of its vertices, and the total score of the triangulation is the sum of these
// values over all n - 2 triangles in the triangulation.
// Return the smallest possible total score that you can achieve with some triangulation of the polygon.
// Example 1:
//   Input: values = [1,2,3]
//   Output: 6
//   Explanation: The polygon is already triangulated, and the score of the only triangle is 6.
// Example 2:
//   Input: values = [3,7,4,5]
//   Output: 144
//   Explanation: There are two triangulations, with possible scores: 3*7*5 + 4*5*7 = 245, or 3*4*5 + 3*4*7 = 144.
//     The minimum score is 144.
// Example 3:
//   Input: values = [1,3,1,4,1,5]
//   Output: 13
//   Explanation: The minimum score triangulation has score 1*1*3 + 1*1*4 + 1*1*5 + 1*1*1 = 13.
// Constraints:
//   n == values.length
//   3 <= n <= 50
//   1 <= values[i] <= 100

// for polygon with vertices V[i...j], we use edge V[i] and V[j] as an edge of triangle, then we can use
// V[i+1], V[i+2], ..., V[j-1] as the third vertex of the triangle. let's say V[x] is the third vertex,
// then we split the polygon into three parts: polygon V[i...x], triangle (V[i], V[x], V[j]), and polygon V[x...j]
// the score is the total sum of the three parts
func minScoreTriangulation(values []int) int {
	dp := make(map[[2]int]int)
	for i := 0; i+2 < len(values); i++ {
		dp[[2]int{i, i + 2}] = values[i] * values[i+1] * values[i+2]
	}
	var min, j, cand int
	for diff := 3; diff < len(values); diff++ { // update diagonally
		for i := 0; i+diff < len(values); i++ {
			j = i + diff
			min = int(1e8)
			// v[i] and v[j] as an edge; v[k] is the third vertex
			for k := i + 1; k < j; k++ {
				cand = dp[[2]int{i, k}] + values[i]*values[j]*values[k] + dp[[2]int{k, j}]
				if cand < min {
					min = cand
				}
			}
			dp[[2]int{i, j}] = min
		}
	}
	return dp[[2]int{0, len(values) - 1}]
}

func main() {
	for _, v := range []struct {
		v   []int
		ans int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{3, 7, 4, 5}, 144},
		{[]int{1, 3, 1, 4, 1, 5}, 13},
	} {
		fmt.Println(minScoreTriangulation(v.v), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/build-a-matrix-with-conditions/
// You are given a positive integer k. You are also given:
//   a 2D integer array rowConditions of size n where rowConditions[i] = [abovei, belowi], and
//   a 2D integer array colConditions of size m where colConditions[i] = [lefti, righti].
// The two arrays contain integers from 1 to k.
// You have to build a k x k matrix that contains each of the numbers from 1 to k
// exactly once. The remaining cells should have the value 0.
// The matrix should also satisfy the following conditions:
//   The number abovei should appear in a row that is strictly above the row at
//     which the number belowi appears for all i from 0 to n - 1.
//   The number lefti should appear in a column that is strictly left of the column
//     at which the number righti appears for all i from 0 to m - 1.
// Return any matrix that satisfies the conditions. If no answer exists, return an empty matrix.
// Example 1:
//   Input: k = 3, rowConditions = [[1,2],[3,2]], colConditions = [[2,1],[3,2]]
//   Output: [[3,0,0],[0,0,1],[0,2,0]]
//   Explanation:
//     The diagram above shows a valid example of a matrix that satisfies all the conditions.
//     The row conditions are the following:
//     - Number 1 is in row 1, and number 2 is in row 2, so 1 is above 2 in the matrix.
//     - Number 3 is in row 0, and number 2 is in row 2, so 3 is above 2 in the matrix.
//     The column conditions are the following:
//     - Number 2 is in column 1, and number 1 is in column 2, so 2 is left of 1 in
//     the matrix.
//     - Number 3 is in column 0, and number 2 is in column 1, so 3 is left of 2 in
//     the matrix.
//     Note that there may be multiple correct answers.
// Example 2:
//   Input: k = 3, rowConditions = [[1,2],[2,3],[3,1],[2,3]], colConditions = [[2,1]]
//   Output: []
//   Explanation: From the first two conditions, 3 has to be below 1 but the third
//     conditions needs 3 to be above 1 to be satisfied.
//     No matrix can satisfy all the conditions, so we return the empty matrix.
// Constraints:
//   2 <= k <= 400
//   1 <= rowConditions.length, colConditions.length <= 10⁴
//   rowConditions[i].length == colConditions[i].length == 2
//   1 <= abovei, belowi, lefti, righti <= k
//   abovei != belowi
//   lefti != righti

// topological sort
func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	// get topological sort
	row, ok := findTopSort(k, rowConditions)
	if !ok {
		return [][]int{}
	}
	col, ok := findTopSort(k, colConditions)
	if !ok {
		return [][]int{}
	}

	// make matrix
	matrix := make([][]int, k)
	for i := range matrix {
		matrix[i] = make([]int, k)
	}
	for i := 1; i <= k; i++ {
		matrix[row[i]][col[i]] = i
	}
	return matrix
}

func findTopSort(k int, rowConditions [][]int) ([]int, bool) {
	top := make([]int, k+1)    // top[i] the topological order of vertex i
	degree := make([]int, k+1) // in-degree of node
	graph := make([][]int, k+1)
	for _, v := range rowConditions {
		// no need de-dup edges, topological sort can cover duplicated edge
		graph[v[0]] = append(graph[v[0]], v[1])
		degree[v[1]]++
	}
	queue := make([]int, 0)
	for i := 1; i <= k; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}
	ind := 0 // the order for each node been visited
	for len(queue) > 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		top[cur] = ind
		ind++
		for _, next := range graph[cur] {
			degree[next]--
			if degree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}
	if ind < k { // can't visit all node, has cycle
		return nil, false
	}
	return top, true
}

func main() {
	for _, v := range []struct {
		k                                 int
		rowConditions, colConditions, ans [][]int
	}{
		{3, [][]int{{1, 2}, {3, 2}}, [][]int{{2, 1}, {3, 2}}, [][]int{{3, 0, 0}, {0, 0, 1}, {0, 2, 0}}},
		{3, [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}, [][]int{{2, 1}}, [][]int{}},
	} {
		fmt.Println(buildMatrix(v.k, v.rowConditions, v.colConditions), v.ans)
	}
}

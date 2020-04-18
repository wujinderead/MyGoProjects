package main

import "fmt"

// https://leetcode.com/problems/grid-illumination/

// On a N x N grid of cells, each cell (x, y) with 0 <= x < N and 0 <= y < N has a lamp.
// Initially, some number of lamps are on. lamps[i] tells us the location of the
// i-th lamp that is on. Each lamp that is on illuminates every square on its x-axis,
// y-axis, and both diagonals (similar to a Queen in chess).
// For the i-th query queries[i] = (x, y), the answer to the query is 1 if the cell (x, y)
// is illuminated, else 0.
// After each query (x, y) [in the order given by queries], we turn off any lamps
// that are at cell (x, y) or are adjacent 8-directionally (ie., share a corner or
// edge with cell (x, y).)
// Return an array of answers. Each value answer[i] should be equal to the answer of
// the i-th query queries[i].
// Example 1:
//   Input: N = 5, lamps = [[0,0],[4,4]], queries = [[1,1],[1,0]]
//   Output: [1,0]
//   Explanation:
//     Before performing the first query we have both lamps [0,0] and [4,4] on.
//     The grid representing which cells are lit looks like this, where [0,0] is the
//     top left corner, and [4,4] is the bottom right corner:
//         1 1 1 1 1
//         1 1 0 0 1
//         1 0 1 0 1
//         1 0 0 1 1
//         1 1 1 1 1
//     Then the query at [1, 1] returns 1 because the cell is lit.  After this query,
//     the lamp at [0, 0] turns off, and the grid now looks like this:
//         1 0 0 0 1
//         0 1 0 0 1
//         0 0 1 0 1
//         0 0 0 1 1
//         1 1 1 1 1
//     Before performing the second query we have only the lamp [4,4] on.  Now the query
//     at [1,0] returns 0, because the cell is no longer lit.
// Note:
//   1 <= N <= 10^9
//   0 <= lamps.length <= 20000
//   0 <= queries.length <= 20000
//   lamps[i].length == queries[i].length == 2

func gridIllumination(N int, lamps [][]int, queries [][]int) []int {
	row, col := make(map[int]int), make(map[int]int)
	dia1, dia2 := make(map[int]int), make(map[int]int)
	lits := make(map[[2]int]struct{})
	for i := range lamps {
		r, c := lamps[i][0], lamps[i][1]
		if _, ok := lits[[2]int{r, c}]; !ok {
			row[r] = row[r] + 1
			col[c] = col[c] + 1
			dia1[r+c] = dia1[r+c] + 1
			dia2[r-c] = dia2[r-c] + 1
			lits[[2]int{r, c}] = struct{}{}
		}
	}
	rs := make([]int, len(queries))
	for i := range queries {
		r, c := queries[i][0], queries[i][1]
		if row[r] > 0 || col[c] > 0 || dia1[r+c] > 0 || dia2[r-c] > 0 {
			rs[i] = 1
		}
		for _, dr := range [3]int{-1, 0, 1} {
			for _, dc := range [3]int{-1, 0, 1} {
				nr, nc := r+dr, c+dc
				if _, ok := lits[[2]int{nr, nc}]; ok {
					// turn off (nr, nc)
					delete(lits, [2]int{nr, nc})
					row[nr] = row[nr] - 1
					col[nc] = col[nc] - 1
					dia1[nr+nc] = dia1[nr+nc] - 1
					dia2[nr-nc] = dia2[nr-nc] - 1
				}
			}
		}
	}
	return rs
}

func main() {
	fmt.Println(gridIllumination(5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 0}}))
	fmt.Println(gridIllumination(5, [][]int{{0, 0}, {4, 4}, {0, 4}, {4, 0}}, [][]int{{2, 2}, {3, 1}, {1, 3}, {3, 1}}))
	fmt.Println(gridIllumination(5, [][]int{{0, 0}, {4, 4}}, [][]int{{1, 1}, {1, 1}}))
}

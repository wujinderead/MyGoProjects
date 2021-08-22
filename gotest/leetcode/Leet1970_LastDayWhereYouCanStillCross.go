package main

import "fmt"

// https://leetcode.com/problems/last-day-where-you-can-still-cross/

// There is a 1-based binary matrix where 0 represents land and 1 represents water.
// You are given integers row and col representing the number of rows and columns
// in the matrix, respectively.
// Initially on day 0, the entire matrix is land. However, each day a new cell becomes flooded
// with water. You are given a 1-based 2D array cells, where cells[i] = [ri, ci] represents
// that on the ith day, the cell on the rith row and cith column (1-based coordinates) will be
// covered with water (i.e., changed to 1).
// You want to find the last day that it is possible to walk from the top to the bottom by only
// walking on land cells. You can start from any cell in the top row and end at any cell in the
// bottom row. You can only travel in the four cardinal directions (left, right, up, and down).
// Return the last day where it is possible to walk from the top to the bottom by only walking
// on land cells.
// Example 1:
//   Input: row = 2, col = 2, cells = [[1,1],[2,1],[1,2],[2,2]]
//   Output: 2
//   Explanation: The above image depicts how the matrix changes each day starting from day 0.
//     The last day where it is possible to cross from top to bottom is on day 2.
// Example 2:
//   Input: row = 2, col = 2, cells = [[1,1],[1,2],[2,1],[2,2]]
//   Output: 1
//   Explanation: The above image depicts how the matrix changes each day starting from day 0.
//     The last day where it is possible to cross from top to bottom is on day 1.
// Example 3:
//   Input: row = 3, col = 3, cells = [[1,2],[2,1],[3,3],[2,2],[1,1],[1,3],[2,3],[3,2],[3,1]]
//   Output: 3
//      d0      d1      d2      d3      d4
//     000     010     010     010     010
//     000     000     100     100     110
//     000     000     000     001     001
//   Explanation: The above image depicts how the matrix changes each day starting from day 0.
//     The last day where it is possible to cross from top to bottom is on day 3.
// Constraints:
//   2 <= row, col <= 2 * 10^4
//   4 <= row * col <= 2 * 10^4
//   cells.length == row * col
//   1 <= ri <= row
//   1 <= ci <= col
//   All the values of cells are unique.

// union-find to group flooded cells and keep track their width, if any flooded cells has
// a width equal to matrix width, we can't travel
func latestDayToCross(row int, col int, cells [][]int) int {
	flooded := make(map[[2]int]struct{}) // set for already flooded cells
	root := make([]int, (row+1)*(col+1))
	left := make([]int, (row+1)*(col+1))
	right := make([]int, (row+1)*(col+1))
	for i := range cells {
		firstRoot := -1
		x, y := cells[i][0], cells[i][1] // current cell
		flooded[[2]int{x, y}] = struct{}{}
		for _, dx := range []int{-1, 0, 1} {
			for _, dy := range []int{-1, 0, 1} { // for current cell's 8 neighbors
				nx, ny := x+dx, y+dy
				if (dx == 0 && dy == 0) || nx < 1 || ny < 1 || nx > row || ny > col {
					continue
				}
				if _, ok := flooded[[2]int{nx, ny}]; !ok {
					continue
				}
				// find first flooded neighbor, find its root firstRoot,
				// set current cell and all neighbors' root as firstRoot
				if firstRoot == -1 {
					firstRoot = getRoot(root, nx*col+ny)
					root[x*col+y] = firstRoot // set root for current cell
					if y < left[firstRoot] {
						left[firstRoot] = y
					}
					if y > right[firstRoot] {
						right[firstRoot] = y
					}
				} else {
					neighborRoot := getRoot(root, nx*col+ny)
					root[neighborRoot] = firstRoot //set all neighbor's root as firstRoot
					// update leftmost and rightmost col for firstRoot
					if left[neighborRoot] < left[firstRoot] {
						left[firstRoot] = left[neighborRoot]
					}
					if right[neighborRoot] > right[firstRoot] {
						right[firstRoot] = right[neighborRoot]
					}
				}
				// if after add current cell, it can't cross, return
				if left[firstRoot] == 1 && right[firstRoot] == col {
					return i
				}
			}
		}
		if firstRoot == -1 { // no neighbor, self root
			root[x*col+y] = x*col + y // x==root[x], means x is root
			left[x*col+y] = y
			right[x*col+y] = y
		}
	}
	return 1
}

func getRoot(root []int, index int) int {
	if root[index] != index {
		x := getRoot(root, root[index])
		root[index] = x
		return x
	}
	return index
}

func main() {
	for _, v := range []struct {
		r, c int
		cc   [][]int
		ans  int
	}{
		{2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}, 2},
		{2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}, 1},
		{3, 3, [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}, 3},
		{2, 6, [][]int{{1, 4}, {1, 3}, {2, 1}, {2, 5}, {2, 2}, {1, 5}, {2, 4}, {1, 2}, {1, 6}, {2, 3}, {2, 6}, {1, 1}}, 3},
	} {
		fmt.Println(latestDayToCross(v.r, v.c, v.cc), v.ans)
	}
}

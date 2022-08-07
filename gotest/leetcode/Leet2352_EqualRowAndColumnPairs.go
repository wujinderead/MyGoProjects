package main

import "fmt"

// https://leetcode.com/problems/equal-row-and-column-pairs/

// Given a 0-indexed n x n integer matrix grid, return the number of pairs (Ri, Cj) such that row Ri
// and column Cj are equal.
// A row and column pair is considered equal if they contain the same elements in the same order
// (i.e. an equal array).
// Example 1:
//   Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
//   Output: 1
//   Explanation: There is 1 equal row and column pair:
//     - (Row 2, Column 1): [2,7,7]
// Example 2:
//   Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
//   Output: 3
//   Explanation: There are 3 equal row and column pairs:
//     - (Row 0, Column 0): [3,1,2,2]
//     - (Row 2, Column 2): [2,4,2,2]
//     - (Row 3, Column 2): [2,4,2,2]
// Constraints:
//   n == grid.length == grid[i].length
//   1 <= n <= 200
//   1 <= grid[i][j] <= 10⁵

func equalPairs(grid [][]int) int {
	const p = int(1e9) + 7
	hash := make(map[int]int)
	for i := range grid {
		h := 0
		for j := range grid {
			h = (h*10 + grid[i][j]) % p
		}
		hash[h] = hash[h] + 1
	}
	ans := 0
	for j := range grid {
		h := 0
		for i := range grid {
			h = (h*10 + grid[i][j]) % p
		}
		ans += hash[h]
	}
	return ans
}

func main() {
	for _, v := range []struct {
		grid [][]int
		ans  int
	}{
		{[][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, 1},
		{[][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, 3},
	} {
		fmt.Println(equalPairs(v.grid), v.ans)
	}

}

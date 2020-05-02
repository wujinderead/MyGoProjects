package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/diagonal-traverse-ii

// Given a list of lists of integers, nums, return all elements of nums in diagonal
// order as shown in the below images.
// Example 1:
//   Input: nums = [[1,2,3],[4,5,6],[7,8,9]]
//   Output: [1,4,2,7,5,3,8,6,9]
//       1 2 3
//       4 5 6
//       7 8 9
// Example 2:
//   Input: nums = [[1,2,3,4,5],[6,7],[8],[9,10,11],[12,13,14,15,16]]
//   Output: [1,6,2,8,7,3,9,4,12,10,5,13,11,14,15,16]
//       1 2 3 4 5
//       6 7
//       8
//       9 A B
//       C D E F 16
// Example 3:
//   Input: nums = [[1,2,3],[4],[5,6,7],[8],[9,10,11]]
//   Output: [1,4,2,5,3,8,6,9,7,10,11]
// Example 4:
//   Input: nums = [[1,2,3,4,5,6]]
//   Output: [1,2,3,4,5,6]
// Constraints:
//   1 <= nums.length <= 10^5
//   1 <= nums[i].length <= 10^5
//   1 <= nums[i][j] <= 10^9
//   There at most 10^5 elements in nums.

// IMPROVEMENT: no need to sort, just iterate the nums.
func findDiagonalOrder(nums [][]int) []int {
	// sort the numbers by the sum of row and col, then by col
	indexes := make([][2]int, 0)
	for i := range nums {
		for j := range nums[i] {
			indexes = append(indexes, [2]int{i, j})
		}
	}
	sort.Sort(coor(indexes))
	ans := make([]int, len(indexes))
	for i := range indexes {
		ans[i] = nums[indexes[i][0]][indexes[i][1]]
	}
	return ans
}

type coor [][2]int

func (c coor) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c coor) Len() int {
	return len(c)
}

func (c coor) Less(i, j int) bool {
	if c[i][0]+c[i][1]==c[j][0]+c[j][1] {
		return c[i][0]>c[j][0]
	}
	return c[i][0]+c[i][1]<c[j][0]+c[j][1]
}

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1,2,3}, {4,5,6}, {7,8,9}}))
	fmt.Println(findDiagonalOrder([][]int{{1,2,3,4,5}, {6,7}, {8}, {9,10,11}, {12,13,14,15,16}}))
	fmt.Println(findDiagonalOrder([][]int{{1,2,3}, {4}, {5,6,7}, {8}, {9,10,11}}))
	fmt.Println(findDiagonalOrder([][]int{{1,2,3,4,5,6}}))
}
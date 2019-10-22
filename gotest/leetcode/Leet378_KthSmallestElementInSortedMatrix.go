package leetcode

import "fmt"

// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

// Given a n x n matrix where each of the rows and columns are sorted in ascending order,
// find the kth smallest element in the matrix.
// Note that it is the kth smallest element in the sorted order, not the kth distinct element.
// Example:
// matrix = [
//    [ 1,  2,  3],
//    [ 4,  6,  8],
//    [ 5,  7,  9]
// ],
// k = 8,
// return 8.
func kthSmallest(matrix [][]int, k int) int {
	return 0
}

// todo
func main() {
	fmt.Println([][]int{{1}}, 1)
	for i := 1; i <= 4; i++ {
		fmt.Println([][]int{{1, 3}, {2, 4}}, i)
	}
	for i := 1; i <= 4; i++ {
		fmt.Println([][]int{{1, 2}, {3, 4}}, i)
	}
	for i := 1; i <= 9; i++ {
		fmt.Println([][]int{{1, 2, 5}, {3, 6, 7}, {4, 8, 9}}, i)
	}
	for i := 1; i <= 16; i++ {
		fmt.Println([][]int{{1, 2, 5}, {3, 6, 7}, {4, 8, 9}}, i)
	}
}

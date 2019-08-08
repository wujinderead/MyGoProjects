package main

import (
	"fmt"
	"math/rand"
)

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

func main() {
	n := 5
	a := rand.Perm(n*n)
	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			fmt.Print()
		}
	}
}
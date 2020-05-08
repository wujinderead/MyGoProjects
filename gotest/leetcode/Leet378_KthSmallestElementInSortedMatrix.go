package main

import (
	"fmt"
	"container/heap"
)

// https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

// Given a n x n matrix where each of the rows and columns are sorted in ascending order,
// find the kth smallest element in the matrix. Note that it is the kth smallest element
// in the sorted order, not the kth distinct element.
// Example:
//   matrix = [
//     [ 1,  2,  3],
//     [ 4,  6,  8],
//     [ 5,  7,  9]
//   ],
//   k = 8,
//   return 8.

// other method: binary search, lo=matrix[0][0], hi=matrix[n-1][n-1], mid=(lo+hi)/2,
// count the numbers that less than mid, if the number is k, we found the k-th is mid.
func kthSmallest(matrix [][]int, k int) int {
	p := pair(make([][3]int, 0))
	for i:=0; i<len(matrix[0]) && i<k; i++ {   // push first row
		heap.Push(&p, [3]int{matrix[0][i], 0, i})
	}
	count := 0
	for {
		pp := heap.Pop(&p).([3]int)
		count++
		if count==k {
			return pp[0]
		}
		i, j := pp[1], pp[2]
		if i+1<len(matrix) {
			heap.Push(&p, [3]int{matrix[i+1][j], i+1, j})
		}
	}
	return 0
}

type pair [][3]int

func (p pair) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p pair) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pair) Len() int {
	return len(p)
}

func (p *pair) Push(x interface{}) {
	*p = append(*p, x.([3]int))
}

func (p *pair) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

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

package main

import (
	"fmt"
	"container/heap"
)

// https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows

// You are given an m * n matrix, mat, and an integer k, which has its rows sorted
// in non-decreasing order. You are allowed to choose exactly 1 element from each row
// to form an array. Return the Kth smallest array sum among all possible arrays.
// Example 1:
//   Input: mat = [[1,3,11],[2,4,6]], k = 5
//   Output: 7
//   Explanation: Choosing one element from each row, the first k smallest sum are:
//     [1,2], [1,4], [3,2], [3,4], [1,6]. Where the 5th sum is 7.
// Example 2:
//   Input: mat = [[1,3,11],[2,4,6]], k = 9
//   Output: 17
// Example 3:
//   Input: mat = [[1,10,10],[1,4,5],[2,3,6]], k = 7
//   Output: 9
//   Explanation: Choosing one element from each row, the first k smallest sum are:
//     [1,1,2], [1,1,3], [1,4,2], [1,4,3], [1,1,6], [1,5,2], [1,5,3]. Where the 7th sum is 9.
// Example 4:
//   Input: mat = [[1,1,10],[2,2,9]], k = 7
//   Output: 12
// Constraints:
//   m == mat.length
//   n == mat.length[i]
//   1 <= m, n <= 40
//   1 <= k <= min(200, n ^ m)
//   1 <= mat[i][j] <= 5000
//   mat[i] is a non decreasing array.

// find the k smallest pair of first two rows, then pair with third row, then pair with fourth row ...
func kthSmallest(mat [][]int, k int) int {
	arr := make([]int, len(mat[0]))
	copy(arr, mat[0])
	for i:=1; i<len(mat); i++ {
		arr = kSmallestPairs(arr, mat[i], k)
	}
	return arr[k-1]
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) []int {
	if k > len(nums1)*len(nums2) {
		k = len(nums1)*len(nums2)
	}
	ans := make([]int, 0, k)
	// push nums1[0] and nums2[i] to heap
	h := heaper(make([]element, min(len(nums2), k)))
	for i:=0; i<len(nums2) && i<k; i++ {
		h[i] = element{nums1[0]+nums2[i], 0, i}
	}
	// pop from heap and add new candidate
	for len(ans)<k {
		e := heap.Pop(&h).(element)
		val, i0, i1 := e.val, e.i0, e.i1
		// initially, we push [0,0], [0,1], [0,2] ... to heap; when pop it, we only increment i0.
		//                      |      |      |
		//                    [1,0]  [1,1]  [1,2] ...
		ans = append(ans, val)
		if i0+1<len(nums1) {
			heap.Push(&h, element{nums1[i0+1]+nums2[i1], i0+1, i1})
		}
	}
	return ans
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

type element struct {
	val int
	i0, i1 int
}

type heaper []element

func (h heaper) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heaper) Len() int {
	return len(h)
}

func (h heaper) Less(i, j int) bool {
	return h[i].val < h[j].val
}

func (h *heaper) Push(x interface{}) {
	*h = append(*h, x.(element))
}

func (h *heaper) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func main() {
	fmt.Println(kthSmallest([][]int{{1,3,11},{2,4,6}}, 5))
	fmt.Println(kthSmallest([][]int{{1,3,11},{2,4,6}}, 9))
	fmt.Println(kthSmallest([][]int{{1,10,10},{1,4,5},{2,3,6}}, 7))
	fmt.Println(kthSmallest([][]int{{1,10,10},{1,4,5},{2,3,6}}, 27))
	fmt.Println(kthSmallest([][]int{{1,1,10},{2,2,9}}, 7))
}
package main

import (
	"fmt"
	"container/heap"
)

// https://leetcode.com/problems/find-k-pairs-with-smallest-sums/

// You are given two integer arrays nums1 and nums2 sorted in ascending order and
// an integer k. Define a pair (u,v) which consists of one element from the first array
// and one element from the second array.
// Find the k pairs (u1,v1),(u2,v2) ...(uk,vk) with the smallest sums.
// Example 1:
//   Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
//   Output: [[1,2],[1,4],[1,6]]
//   Explanation: The first 3 pairs are returned from the sequence:
//     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
// Example 2:
//   Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
//   Output: [1,1],[1,1]
//   Explanation: The first 2 pairs are returned from the sequence:
//     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
// Example 3:
//   Input: nums1 = [1,2], nums2 = [3], k = 3
//   Output: [1,3],[2,3]
//   Explanation: All possible pairs are returned from the sequence: [1,3],[2,3]

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
    if k > len(nums1)*len(nums2) {
    	k = len(nums1)*len(nums2)
	}
	ans := make([][]int, 0, k)
	// push nums1[0] and nums2[i] to heap
	h := heaper(make([]element, min(len(nums2), k)))
	for i:=0; i<len(nums2) && i<k; i++ {
		h[i] = element{nums1[0]+nums2[i], 0, i}
	}
	// pop from heap and add new candidate
	for len(ans)<k {
		e := heap.Pop(&h).(element)
		i0, i1 := e.i0, e.i1
		// initially, we push [0,0], [0,1], [0,2] ... to heap; when pop it, we only increment i0.
		//                      |      |      |
		//                    [1,0]  [1,1]  [1,2] ...
		ans = append(ans, []int{nums1[i0], nums2[i1]})
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
	fmt.Println(kSmallestPairs([]int{1,7,11}, []int{2,4,6}, 1))
	fmt.Println(kSmallestPairs([]int{1,7,11}, []int{2,4,6}, 2))
	fmt.Println(kSmallestPairs([]int{1,7,11}, []int{2,4,6}, 3))
	fmt.Println(kSmallestPairs([]int{1,7,11}, []int{2,4,6}, 6))
	fmt.Println(kSmallestPairs([]int{1,7,11}, []int{2,4,6}, 9))
	fmt.Println(kSmallestPairs([]int{1,1,2}, []int{1,2,3}, 2))
	fmt.Println(kSmallestPairs([]int{1,2}, []int{3}, 3))
}
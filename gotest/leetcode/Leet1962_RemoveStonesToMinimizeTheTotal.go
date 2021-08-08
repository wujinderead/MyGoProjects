package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/remove-stones-to-minimize-the-total/

// You are given a 0-indexed integer array piles, where piles[i] represents the number of stones
// in the ith pile, and an integer k. You should apply the following operation exactly k times:
//   Choose any piles[i] and remove floor(piles[i] / 2) stones from it.
//   Notice that you can apply the operation on the same pile more than once.
//   Return the minimum possible total number of stones remaining after applying the k operations.
// floor(x) is the greatest integer that is smaller than or equal to x (i.e., rounds x down).
// Example 1:
//   Input: piles = [5,4,9], k = 2
//   Output: 12
//   Explanation: Steps of a possible scenario are:
//     - Apply the operation on pile 2. The resulting piles are [5,4,5].
//     - Apply the operation on pile 0. The resulting piles are [3,4,5].
//     The total number of stones in [3,4,5] is 12.
// Example 2:
//   Input: piles = [4,3,6,7], k = 3
//   Output: 12
//   Explanation: Steps of a possible scenario are:
//     - Apply the operation on pile 3. The resulting piles are [4,3,3,7].
//     - Apply the operation on pile 4. The resulting piles are [4,3,3,4].
//     - Apply the operation on pile 0. The resulting piles are [2,3,3,4].
//     The total number of stones in [2,3,3,4] is 12.
// Constraints:
// 1 <= piles.length <= 10^5
// 1 <= piles[i] <= 10^4
// 1 <= k <= 10^5

// just use heap
func minStoneSum(p []int, k int) int {
	pp := piles(p)
	heap.Init(&pp)
	for i := 0; i < k; i++ {
		max := heap.Pop(&pp).(int)
		max -= max / 2
		heap.Push(&pp, max)
		if max == 1 {
			break
		}
	}
	sum := 0
	for i := range pp {
		sum += pp[i]
	}
	return sum
}

type piles []int

func (p piles) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p piles) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p piles) Len() int {
	return len(p)
}

func (p *piles) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *piles) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func main() {
	for _, v := range []struct {
		p      []int
		k, ans int
	}{
		{[]int{5, 4, 9}, 2, 12},
		{[]int{4, 3, 6, 7}, 3, 12},
		{[]int{1, 1, 1, 1, 1, 1}, 10, 6},
		{[]int{1, 1, 1, 1, 1, 2}, 10, 6},
	} {
		fmt.Println(minStoneSum(v.p, v.k), v.ans)
	}
}

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// There are n engineers numbered from 1 to n and two arrays: speed and efficiency,
// where speed[i] and efficiency[i] represent the speed and efficiency for the i-th
// engineer respectively. Return the maximum performance of a team composed of
// at most k engineers, since the answer can be a huge number,
// return this modulo 10^9 + 7.
// The performance of a team is the sum of their engineers' speeds multiplied by
// the minimum efficiency among their engineers.
// Example 1:
//   Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 2
//   Output: 60
//   Explanation:
//     We have the maximum performance of the team by selecting engineer 2 (with
//     speed=10 and efficiency=4) and engineer 5 (with speed=5 and efficiency=7).
//     That is, performance = (10 + 5) * min(4, 7) = 60.
// Example 2:
//   Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 3
//   Output: 68
//   Explanation:
//     This is the same example as the first but k = 3. We can select engineer 1,
//     engineer 2 and engineer 5 to get the maximum performance of the team. That is,
//     performance = (2 + 10 + 5) * min(5, 4, 7) = 68.
// Example 3:
//   Input: n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 4
//   Output: 72
// Constraints:
//   1 <= n <= 10^5
//   speed.length == n
//   efficiency.length == n
//   1 <= speed[i] <= 10^5
//   1 <= efficiency[i] <= 10^8
//   1 <= k <= n

// take the example in problem
// speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 4
// sort the engineers by efficiency, we got:
// S:  1 5 2 10 5 8
// E:  9 7 5 4  3 2
// for the first k elements, we got sum(S) increasing, min(E) decreasing,
// that's all candidates for the answer. we establish a heap for S.
// for next S, if we can get sum(S) increase, then it's a candidate.
func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	eles := pairs(make([][2]int, n))
	for i := range speed {
		eles[i][0] = speed[i]
		eles[i][1] = efficiency[i]
	}
	sort.Sort(eles)
	rmax := 0
	sum := 0
	minheap := minheap(make([]int, 0, k))
	for i := 0; i < k; i++ {
		sum += eles[i][0]
		if sum*eles[i][1] > rmax {
			rmax = sum * eles[i][1]
		}
		heap.Push(&minheap, eles[i][0])
	}
	for i := k; i < n; i++ {
		peek := minheap[0]
		// if current s > heap.top, and we can get a larger sum(s)
		if eles[i][0] > peek {
			sum += eles[i][0] - peek

			// substitute heap top with new s
			minheap[0] = eles[i][0]
			heap.Fix(&minheap, 0)

			// check if rmax changes
			if sum*eles[i][1] > rmax {
				rmax = sum * eles[i][1]
			}
		}
	}
	return rmax % (1000000007)
}

type pairs [][2]int
type minheap []int

func (m minheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minheap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m minheap) Len() int {
	return len(m)
}

func (m *minheap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minheap) Pop() interface{} {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pairs) Less(i, j int) bool {
	return p[i][1] > p[j][1]
}

func (p pairs) Len() int {
	return len(p)
}

func (p *pairs) Push(x interface{}) {
	*p = append(*p, x.([2]int))
}

func (p *pairs) Pop() (x interface{}) {
	ele := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return ele
}

func main() {
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2))
	fmt.Println()
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 3))
	fmt.Println()
	fmt.Println(maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 4))
}

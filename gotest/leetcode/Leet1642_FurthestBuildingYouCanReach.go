package main

import (
	"container/heap"
	"fmt"
)

// https://leetcode.com/problems/furthest-building-you-can-reach/

// You are given an integer array heights representing the heights of buildings, some bricks, and some ladders.
// You start your journey from building 0 and move to the next building by possibly using bricks or ladders.
// While moving from building i to building i+1 (0-indexed), If the current building's height is greater than
// or equal to the next building's height, you do not need a ladder or bricks.
// If the current building's height is less than the next building's height, you can either use one ladder
// or (h[i+1] - h[i]) bricks.
// Return the furthest building index (0-indexed) you can reach if you use the given ladders and bricks optimally.
// Example 1:
//   Input: heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
//   Output: 4
//   Explanation: Starting at building 0, you can follow these steps:
//     - Go to building 1 without using ladders nor bricks since 4 >= 2.
//     - Go to building 2 using 5 bricks. You must use either bricks or ladders because 2 < 7.
//     - Go to building 3 without using ladders nor bricks since 7 >= 6.
//     - Go to building 4 using your only ladder. You must use either bricks or ladders because 6 < 9.
//     It is impossible to go beyond building 4 because you do not have any more bricks or ladders.
// Example 2:
//   Input: heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
//   Output: 7
// Example 3:
//   Input: heights = [14,3,19,3], bricks = 17, ladders = 0
//   Output: 3
// Constraints:
//   1 <= heights.length <= 10^5
//   1 <= heights[i] <= 10^6
//   0 <= bricks <= 10^9
//   0 <= ladders <= heights.length

/*
https://leetcode.com/problems/furthest-building-you-can-reach/discuss/918515/JavaC%2B%2BPython-Priority-Queue
A more concise implement with same idea:
	public int furthestBuilding(int[] A, int bricks, int ladders) {
        PriorityQueue<Integer> pq = new PriorityQueue<>();
        for (int i = 0; i < A.length - 1; i++) {
            int d = A[i + 1] - A[i];
            if (d > 0)
                pq.add(d);
            if (pq.size() > ladders)
                bricks -= pq.poll();
            if (bricks < 0)
                return i;
        }
        return A.length - 1;
    }
*/
func furthestBuilding(heights []int, bricks int, ladders int) int {
	// make a min-heap of ladders size, keep the heap contains the large numbers to use ladders.
	// check if the remain numbers can be covered by bricks
	canReach := 0
	heaper := minheap(make([]int, 0, ladders))
	sum := 0
	for i := 1; i < len(heights); i++ {
		if heights[i] <= heights[i-1] {
			canReach = i
			continue
		}
		// handle ascending
		diff := heights[i] - heights[i-1]
		if ladders > 0 {
			if len(heaper) < ladders {
				heap.Push(&heaper, diff)
				canReach = i
				continue
			} else if diff > heaper[0] { // diff > min ele of heap, replace it
				pop := heap.Pop(&heaper).(int)
				heap.Push(&heaper, diff)
				diff = pop
			}
		}
		// check if brick enough
		sum += diff
		if sum > bricks {
			break
		}
		canReach = i
	}
	return canReach
}

type minheap []int

func (m minheap) Len() int {
	return len(m)
}

func (m minheap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m minheap) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m *minheap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minheap) Pop() interface{} {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func main() {
	for _, v := range []struct {
		h         []int
		b, l, ans int
	}{
		{[]int{4, 2, 7, 6, 9, 14, 12}, 5, 1, 4},
		{[]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2, 7},
		{[]int{14, 3, 19, 3}, 17, 0, 3},
	} {
		fmt.Println(furthestBuilding(v.h, v.b, v.l), v.ans)
	}
}

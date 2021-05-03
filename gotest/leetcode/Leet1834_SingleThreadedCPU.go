package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.com/problems/single-threaded-cpu/

// You are given n tasks labeled from 0 to n - 1 represented by a 2D integer array tasks,
// where tasks[i] = [enqueueTimei, processingTimei] means that the ith task will be available
// to process at enqueueTimei and will take processingTimei to finish processing.
// You have a single-threaded CPU that can process at most one task at a time and will
// act in the following way:
// - If the CPU is idle and there are no available tasks to process, the CPU remains idle.
// - If the CPU is idle and there are available tasks, the CPU will choose the one
//   with the shortest processing time. If multiple tasks have the same shortest
//   processing time, it will choose the task with the smallest index.
// - Once a task is started, the CPU will process the entire task without stopping.
// - The CPU can finish a task then start a new one instantly.
// Return the order in which the CPU will process the tasks.
// Example 1:
//   Input: tasks = [[1,2],[2,4],[3,2],[4,1]]
//   Output: [0,2,3,1]
//   Explanation: The events go as follows:
//     - At time = 1, task 0 is available to process. Available tasks = {0}.
//     - Also at time = 1, the idle CPU starts processing task 0. Available tasks = {}.
//     - At time = 2, task 1 is available to process. Available tasks = {1}.
//     - At time = 3, task 2 is available to process. Available tasks = {1, 2}.
//     - Also at time = 3, the CPU finishes task 0 and starts processing task 2 as it is the shortest.
//       Available tasks = {1}.
//     - At time = 4, task 3 is available to process. Available tasks = {1, 3}.
//     - At time = 5, the CPU finishes task 2 and starts processing task 3 as it is the shortest.
//       Available tasks = {1}.
//     - At time = 6, the CPU finishes task 3 and starts processing task 1. Available tasks = {}.
//     - At time = 10, the CPU finishes task 1 and becomes idle.
// Example 2:
//   Input: tasks = [[7,10],[7,12],[7,5],[7,4],[7,2]]
//   Output: [4,3,2,0,1]
//   Explanation: The events go as follows:
//     - At time = 7, all the tasks become available. Available tasks = {0,1,2,3,4}.
//     - Also at time = 7, the idle CPU starts processing task 4. Available tasks = {0,1,2,3}.
//     - At time = 9, the CPU finishes task 4 and starts processing task 3. Available tasks = {0,1,2}.
//     - At time = 13, the CPU finishes task 3 and starts processing task 2. Available tasks = {0,1}.
//     - At time = 18, the CPU finishes task 2 and starts processing task 0. Available tasks = {1}.
//     - At time = 28, the CPU finishes task 0 and starts processing task 1. Available tasks = {}.
//     - At time = 40, the CPU finishes task 1 and becomes idle.
// Constraints:
//   tasks.length == n
//   1 <= n <= 10^5
//   1 <= enqueueTimei, processingTimei <= 10^9

func getOrder(tasks [][]int) []int {
	ts := make(objs, len(tasks))
	for i := range ts {
		ts[i] = [3]int{tasks[i][0], tasks[i][1], i}
	}
	// first sort task by start
	sort.SliceStable(ts, func(i, j int) bool {
		return ts[i][0] < ts[j][0]
	})
	order := make([]int, 0)
	end := ts[0][0] // end is the earliest time
	h := &objs{}
outer:
	for i := 0; i < len(ts); i++ {
		if ts[i][0] <= end { // push runnable task in queue to heap
			heap.Push(h, ts[i])
			continue
		}
		for h.Len() > 0 { // no runnable task in queue, pop a task and run
			x := heap.Pop(h).([3]int)
			order = append(order, x[2])
			end = end + x[1]     // extend current time
			if ts[i][0] <= end { // add to heap if current task become runnable
				heap.Push(h, ts[i])
				continue outer
			}
		}
		end = ts[i][0] // no task in heap, current task is the first task we can run, add it to heap
		heap.Push(h, ts[i])
	}
	for h.Len() > 0 {
		x := heap.Pop(h).([3]int)
		order = append(order, x[2])
	}
	return order
}

type objs [][3]int

func (t objs) Len() int {
	return len(t)
}

func (t objs) Less(i, j int) bool {
	if t[i][1] == t[j][1] {
		return t[i][2] < t[j][2]
	}
	return t[i][1] < t[j][1]
}

func (t objs) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *objs) Push(x interface{}) {
	*t = append(*t, x.([3]int))
}

func (t *objs) Pop() interface{} {
	x := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return x
}

func main() {
	for _, v := range []struct {
		tasks [][]int
		ans   []int
	}{
		{[][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}, []int{0, 2, 3, 1}},
		{[][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}, []int{4, 3, 2, 0, 1}},
		{[][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24},
			{2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}},
			[]int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7}},
	} {
		fmt.Println(getOrder(v.tasks), v.ans)
	}
}

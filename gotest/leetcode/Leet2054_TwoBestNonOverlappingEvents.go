package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/two-best-non-overlapping-events/

// You are given a 0-indexed 2D integer array of events where events[i] = [startTimei, endTimei, valuei].
// The ith event starts at startTimei and ends at endTimei, and if you attend this event, you will
// receive a value of valuei. You can choose at most two non-overlapping events to attend such that
// the sum of their values is maximized.
// Return this maximum sum.
// Note that the start time and end time is inclusive: that is, you cannot attend two events where one
// of them starts and the other ends at the same time. More specifically, if you attend an event with
// end time t, the next event must start at or after t + 1.
// Example 1:
//   Input: events = [[1,3,2],[4,5,2],[2,4,3]]
//   Output: 4
//   Explanation: Choose the green events, 0 and 1 for a sum of 2 + 2 = 4.
// Example 2:
//   Input: events = [[1,3,2],[4,5,2],[1,5,5]]
//   Output: 5
//   Explanation: Choose event 2 for a sum of 5.
// Example 3:
//   Input: events = [[1,5,3],[1,5,1],[6,6,5]]
//   Output: 8
//   Explanation: Choose events 0 and 2 for a sum of 3 + 5 = 8.
// Constraints:
//   2 <= events.length <= 10^5
//   events[i].length == 3
//   1 <= startTimei <= endTimei <= 10^9
//   1 <= valuei <= 10^6

func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool { // sort by start index
		return events[i][0] < events[j][0]
	})
	max := make([]int, len(events)+1) // max[i] is the max value in events[i:]
	for i := len(events) - 1; i >= 0; i-- {
		max[i] = max[i+1]
		if events[i][2] > max[i] {
			max[i] = events[i][2]
		}
	}
	ans := 0
	for i := 0; i < len(events); i++ {
		target := events[i][1] + 1 // the other event's startTime must > current event's endTime
		left, right := i, len(events)
		// search first x that event[x][0] > = target
		for left < right {
			mid := (left + right) / 2
			if events[mid][0] >= target {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if events[i][2]+max[left] > ans {
			ans = events[i][2] + max[left]
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		e   [][]int
		ans int
	}{
		{[][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}}, 4},
		{[][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}}, 5},
		{[][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}}, 8},
		{[][]int{{1, 5, 3}, {2, 4, 7}}, 7},
		{[][]int{{1, 5, 6}, {2, 4, 3}}, 6},
		{[][]int{{1, 5, 3}, {5, 7, 2}}, 3},
		{[][]int{{1, 5, 3}, {6, 7, 2}}, 5},
	} {
		fmt.Println(maxTwoEvents(v.e), v.ans)
	}
}

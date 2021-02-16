package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/maximum-number-of-events-that-can-be-attended-ii/

// You are given an array of events where events[i] = [startDayi, endDayi, valuei].
// The ith event starts at startDayi and ends at endDayi, and if you attend this event,
// you will receive a value of valuei. You are also given an integer k which represents
// the maximum number of events you can attend.
// You can only attend one event at a time. If you choose to attend an event, you must
// attend the entire event. Note that the end day is inclusive: that is, you cannot attend
// two events where one of them starts and the other ends on the same day.
// Return the maximum sum of values that you can receive by attending events.
// Example 1:
//   Input: events = [[1,2,4],[3,4,3],[2,3,1]], k = 2
//   Output: 7
//      Time       1    2    3    4
//      Event0  |----4----|
//      Event1            |----3----|
//      Event2       |----1----|
//   Explanation: Choose the green events, 0 and 1 (0-indexed) for a total value of 4 + 3 = 7.
// Example 2:
//   Input: events = [[1,2,4],[3,4,3],[2,3,10]], k = 2
//   Output: 10
//      Time       1    2    3    4
//      Event0  |----4----|
//      Event1            |----3----|
//      Event2       |----10---|
//   Explanation: Choose event 2 for a total value of 10.
//     Notice that you cannot attend any other event as they overlap,
//     and that you do not have to attend k events.
// Example 3:
//   Input: events = [[1,1,1],[2,2,2],[3,3,3],[4,4,4]], k = 3
//   Output: 9
//   Explanation: Although the events do not overlap, you can only attend 3 events.
//     Pick the highest valued three.
// Constraints:
//    1 <= k <= events.length
//    1 <= k * events.length <= 10^6
//    1 <= startDayi <= endDayi <= 10^9
//    1 <= valuei <= 10^6

// let V(i, k) be the value we can get from events[i:] with max k events.
// for ith event, if we attend it, we can get vi+V(i+x, k-1)
// if we skip it, we can get V(i+1, k), we want the larger.
func maxValue(es [][]int, K int) int {
	evts := events(es)
	sort.Sort(evts)
	old, new := make([]int, len(evts)), make([]int, len(evts))

	// calculate k=1
	old[len(evts)-1] = evts[len(evts)-1][2]
	for i := len(evts) - 2; i >= 0; i-- {
		old[i] = max(old[i+1], evts[i][2])
	}

	// dp
	for k := 2; k <= K; k++ {
		new[len(evts)-1] = old[len(evts)-1]
		for i := len(evts) - 2; i >= 0; i-- {
			new[i] = new[i+1] // V(i+1, k)
			// if use event[i], we can only use the event whose start > event[i].end
			e := evts[i][1]
			if evts[len(evts)-1][0] <= e {
				new[i] = max(new[i], evts[i][2])
				continue
			}
			// binary search to find first index that evts[index].start > end
			l, r := i+1, len(evts)-1
			for l < r {
				mid := (l + r) / 2
				if evts[mid][0] <= e {
					l = mid + 1
				} else {
					r = mid
				}
			}
			new[i] = max(new[i], old[l]+evts[i][2]) // V(l, k-1) + value[i]
		}
		old, new = new, old
	}
	return old[0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type events [][]int

func (e events) Len() int {
	return len(e)
}

func (e events) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e events) Less(i, j int) bool {
	if e[i][0] == e[j][0] {
		return e[i][1] < e[j][1]
	}
	return e[i][0] < e[j][0]
}

func main() {
	for _, v := range []struct {
		e      [][]int
		k, ans int
	}{
		//{[][]int{{1,2,4},{3,4,3},{2,3,1}}, 2, 7},
		//{[][]int{{1,2,4},{3,4,3},{2,3,10}}, 2, 10},
		//{[][]int{{1,1,1},{2,2,2},{3,3,3},{4,4,4}}, 3, 9},
		{[][]int{{11, 17, 56}, {24, 40, 53}, {5, 62, 67}, {66, 69, 84}, {56, 89, 15}}, 2, 151},
	} {
		fmt.Println(maxValue(v.e, v.k), v.ans)
	}
}

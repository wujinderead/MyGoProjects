package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.com/problems/meeting-rooms-iii/

// You are given an integer n. There are n rooms numbered from 0 to n - 1.
// You are given a 2D integer array meetings where meetings[i] = [starti, endi] means that a meeting
// will be held during the half-closed time interval [starti, endi). All the values of starti are unique.
// Meetings are allocated to rooms in the following manner:
//   Each meeting will take place in the unused room with the lowest number.
//   If there are no available rooms, the meeting will be delayed until a room becomes free. The delayed
//     meeting should have the same duration as the original meeting.
//   When a room becomes unused, meetings that have an earlier original start tim额should be given the room.
// Return the number of the room that held the most meetings. If there are multiple rooms, return the room
// with the lowest number.
// A half-closed interval [a, b) is the interval between a and b including a and not including b.
// Example 1:
//   Input: n = 2, meetings = [[0,10],[1,5],[2,7],[3,4]]
//   Output: 0
//   Explanation:
//     - At time 0, both rooms are not being used. The first meeting starts in room 0.
//     - At time 1, only room 1 is not being used. The second meeting starts in room 1.
//     - At time 2, both rooms are being used. The third meeting is delayed.
//     - At time 3, both rooms are being used. The fourth meeting is delayed.
//     - At time 5, the meeting in room 1 finishes. The third meeting starts in room 1
//       for the time period [5,10).
//     - At time 10, the meetings in both rooms finish. The fourth meeting starts in
//       room 0 for the time period [10,11).
//     Both rooms 0 and 1 held 2 meetings, so we return 0.
// Example 2:
//   Input: n = 3, meetings = [[1,20],[2,10],[3,5],[4,9],[6,8]]
//   Output: 1
//   Explanation:
//     - At time 1, all three rooms are not being used. The first meeting starts in room 0.
//     - At time 2, rooms 1 and 2 are not being used. The second meeting starts in room 1.
//     - At time 3, only room 2 is not being used. The third meeting starts in room 2.
//     - At time 4, all three rooms are being used. The fourth meeting is delayed.
//     - At time 5, the meeting in room 2 finishes. The fourth meeting starts in room 2
//       for the time period [5,10).
//     - At time 6, all three rooms are being used. The fifth meeting is delayed.
//     - At time 10, the meetings in rooms 1 and 2 finish. The fifth meeting starts in
//       room 1 for the time period [10,12).
//     Room 0 held 1 meeting while rooms 1 and 2 each held 2 meetings, so we return 1.
// Constraints:
//   1 <= n <= 100
//   1 <= meetings.length <= 10⁵
//   meetings[i].length == 2
//   0 <= starti < endi <= 5 * 10⁵
//   All the values of starti are unique.

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool { // sort meetings by start
		return meetings[i][0] < meetings[j][0]
	})

	max := 0                          // max index
	held := make([]int, n)            // number of meetings held in room[i]
	occupied := make(pairs, 0, n)     // heap of (ordinal, available time) pairs. order by available time, ordinal
	available := make(ordinals, 0, n) // heap of available ordinals
	for i := 0; i < n; i++ {
		heap.Push(&available, i) // initially, all rooms are available
	}

	for _, m := range meetings {
		start, end := m[0], m[1]
		// for current meeting, assume current time is start, pop finished meetings
		for occupied.Len() > 0 && occupied[0][1] <= start {
			pop := heap.Pop(&occupied).([2]int)
			heap.Push(&available, pop[0]) // push empty meeting room to available
		}
		var ind int
		if available.Len() > 0 { // if current time has available room
			pop := heap.Pop(&available).(int)
			heap.Push(&occupied, [2]int{pop, end})
			ind = pop
			held[ind]++
		} else { // wait until a meeting finished, current time become pop[1]
			pop := heap.Pop(&occupied).([2]int)
			heap.Push(&occupied, [2]int{pop[0], pop[1] + end - start}) // the finish time become pop[1]+end-start
			ind = pop[0]
			held[ind]++
		}
		if held[ind] > held[max] || (held[ind] == held[max] && ind < max) { // update max index
			max = ind
		}
	}

	return max
}

type pairs [][2]int // (ordinal, available time) pair

func (p pairs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p pairs) Len() int      { return len(p) }
func (p pairs) Less(i, j int) bool {
	if p[i][1] == p[j][1] {
		return p[i][0] < p[j][0]
	}
	return p[i][1] < p[j][1]
}
func (p *pairs) Push(x interface{}) { *p = append(*p, x.([2]int)) }
func (p *pairs) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

type ordinals []int // ordinals

func (p ordinals) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p ordinals) Len() int            { return len(p) }
func (p ordinals) Less(i, j int) bool  { return p[i] < p[j] }
func (p *ordinals) Push(x interface{}) { *p = append(*p, x.(int)) }
func (p *ordinals) Pop() interface{} {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func main() {
	for _, v := range []struct {
		n        int
		meetings [][]int
		ans      int
	}{
		{2, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}, 0},
		{3, [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}, 1},
		{4, [][]int{{18, 19}, {3, 12}, {17, 19}, {2, 13}, {7, 10}}, 0},
	} {
		fmt.Println(mostBooked(v.n, v.meetings), v.ans)
	}
}

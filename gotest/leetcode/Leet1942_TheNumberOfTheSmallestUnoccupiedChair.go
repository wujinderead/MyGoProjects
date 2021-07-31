package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode.com/problems/the-number-of-the-smallest-unoccupied-chair/

// There is a party where n friends numbered from 0 to n - 1 are attending. There is an infinite
// number of chairs in this party that are numbered from 0 to infinity. When a friend arrives at
// the party, they sit on the unoccupied chair with the smallest number.
// For example, if chairs 0, 1, and 5 are occupied when a friend comes, they will sit on chair number 2.
// When a friend leaves the party, their chair becomes unoccupied at the moment they leave. If
// another friend arrives at that same moment, they can sit in that chair.
// You are given a 0-indexed 2D integer array times where times[i] = [arrivali, leavingi], indicating
// the arrival and leaving times of the ith friend respectively, and an integer targetFriend.
// All arrival times are distinct.
// Return the chair number that the friend numbered targetFriend will sit on.
// Example 1:
//   Input: times = [[1,4],[2,3],[4,6]], targetFriend = 1
//   Output: 1
//   Explanation:
//     - Friend 0 arrives at time 1 and sits on chair 0.
//     - Friend 1 arrives at time 2 and sits on chair 1.
//     - Friend 1 leaves at time 3 and chair 1 becomes empty.
//     - Friend 0 leaves at time 4 and chair 0 becomes empty.
//     - Friend 2 arrives at time 4 and sits on chair 0.
//     Since friend 1 sat on chair 1, we return 1.
// Example 2:
//   Input: times = [[3,10],[1,5],[2,6]], targetFriend = 0
//   Output: 2
//   Explanation:
//     - Friend 1 arrives at time 1 and sits on chair 0.
//     - Friend 2 arrives at time 2 and sits on chair 1.
//     - Friend 0 arrives at time 3 and sits on chair 2.
//     - Friend 1 leaves at time 5 and chair 0 becomes empty.
//     - Friend 2 leaves at time 6 and chair 1 becomes empty.
//     - Friend 0 leaves at time 10 and chair 2 becomes empty.
//     Since friend 0 sat on chair 2, we return 2.
// Constraints:
//   n == times.length
//   2 <= n <= 10^4
//   times[i].length == 2
//   1 <= arrivali < leavingi <= 10^5
//   0 <= targetFriend <= n - 1
//   Each arrivali time is distinct.

func smallestChair(times [][]int, targetFriend int) int {
	events := make([][3]int, 0, len(times)*2)
	for i := range times {
		// occur time, person index, type 1=arrive, 0=leave
		if times[i][0] <= times[targetFriend][0] {
			events = append(events, [3]int{times[i][0], i, 1})
		}
		if times[i][1] <= times[targetFriend][0] {
			events = append(events, [3]int{times[i][1], i, 0})
		}
	}
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][2] < events[j][2]
	})
	maxSeat := 0
	freeHeap := &minHeap{}       // min heap of free seats
	seatMap := make(map[int]int) // map[person_index]seat
	for _, e := range events {
		// leave event
		person, event := e[1], e[2]
		if event == 0 {
			seat := seatMap[person]
			delete(seatMap, person)
			heap.Push(freeHeap, seat) // got a free seat, push to heap
			continue
		}
		// arrive event
		var seat int
		if freeHeap.Len() > 0 {
			seat = heap.Pop(freeHeap).(int) // find a seat in heap
			seatMap[person] = seat
		} else {
			seatMap[person] = maxSeat // sit in the max seat
			seat = maxSeat
			maxSeat++
		}
		if person == targetFriend {
			return seat
		}
	}
	return -1
}

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() (x interface{}) {
	x = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func main() {
	for _, v := range []struct {
		ti      [][]int
		ta, ans int
	}{
		{[][]int{{1, 4}, {2, 3}, {4, 6}}, 1, 1},
		{[][]int{{3, 10}, {1, 5}, {2, 6}}, 0, 2},
		{[][]int{{1, 8}, {2, 6}, {3, 7}, {4, 5}, {6, 9}, {7, 10}}, 4, 1},
		{[][]int{{1, 8}, {2, 6}, {3, 7}, {4, 5}, {5, 9}, {7, 10}}, 4, 3},
	} {
		fmt.Println(smallestChair(v.ti, v.ta), v.ans)
	}
}

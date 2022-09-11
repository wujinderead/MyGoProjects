package main

import "fmt"

// https://leetcode.com/problems/divide-intervals-into-minimum-number-of-groups/

// You are given a 2D integer array intervals where intervals[i] = [lefti, righti] represents the
// inclusive interval [lefti, righti].
// You have to divide the intervals into one or more groups such that each interval is in exactly
// one group, and no two intervals that are in the same group intersect each other.
// Return the minimum number of groups you need to make.
// Two intervals intersect if there is at least one common number between them.
// For example, the intervals [1, 5] and [5, 8] intersect.
// Example 1:
//   Input: intervals = [[5,10],[6,8],[1,5],[2,3],[1,10]]
//   Output: 3
//   Explanation: We can divide the intervals into the following groups:
//     - Group 1: [1, 5], [6, 8].
//     - Group 2: [2, 3], [5, 10].
//     - Group 3: [1, 10].
//     It can be proven that it is not possible to divide the intervals into fewer than 3 groups.
// Example 2:
//   Input: intervals = [[1,3],[5,6],[8,10],[11,13]]
//   Output: 1
//   Explanation: None of the intervals overlap, so we can put all of them in one group.
// Constraints:
//   1 <= intervals.length <= 10⁵
//   intervals[i].length == 2
//   1 <= lefti <= righti <= 10⁶

// find the point that shared by most intervals
// line-sweep, can ignore input range use a map and sort.
func minGroups(intervals [][]int) int {
	max := 0
	prefix := make([]int, 1000000+1)
	for _, v := range intervals {
		prefix[v[0]]++
		prefix[v[1]+1]--
	}
	for i := 1; i < len(prefix); i++ {
		prefix[i] += prefix[i-1]
		if prefix[i] > max {
			max = prefix[i]
		}
	}
	return max
}

func main() {
	for _, v := range []struct {
		intervals [][]int
		ans       int
	}{
		{[][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}, 3},
		{[][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}, 1},
		{[][]int{{441459, 446342}, {801308, 840640}, {871890, 963447}, {228525, 336985}, {807945, 946787}, {479815, 507766}, {693292, 944029}, {751962, 821744}}, 4},
	} {
		fmt.Println(minGroups(v.intervals), v.ans)
	}
}

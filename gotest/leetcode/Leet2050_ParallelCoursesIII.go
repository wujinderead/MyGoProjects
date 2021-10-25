package main

import "fmt"

// https://leetcode.com/problems/parallel-courses-iii/

// You are given an integer n, which indicates that there are n courses labeled from 1 to n.
// You are also given a 2D integer array relations where relations[j] = [prevCoursej, nextCoursej]
// denotes that course prevCoursej has to be completed before course nextCoursej (prerequisite relationship).
// Furthermore, you are give n a 0-indexed integer array time where time[i] denotes how many months it takes
// to complete the (i+1)th course.
// You must find the minimum number of months needed to complete all the courses following these rules:
//   You may start taking a course at any time if the prerequisites are met.
//   Any number of courses can be taken at the same time.
// Return the minimum number of months needed to complete all the courses.
// Note: The test cases are generated such that it is possible to complete every course
// (i.e., the graph is a directed acyclic graph).
// Example 1:
//   Input: n = 3, relations = [[1,3],[2,3]], time = [3,2,5]
//   Output: 8
//   Explanation: The figure above represents the given graph and the time required to complete each course.
//     We start course 1 and course 2 simultaneously at month 0.
//     Course 1 takes 3 months and course 2 takes 2 months to complete respectively.
//     Thus, the earliest time we can start course 3 is at month 3, and the total time required is 3+5=8 months.
// Example 2:
//   Input: n = 5, relations = [[1,5],[2,5],[3,5],[3,4],[4,5]], time = [1,2,3,4,5]
//   Output: 12
//   Explanation: The figure above represents the given graph and the time required to complete each course.
//     You can start courses 1, 2, and 3 at month 0.
//     You can complete them after 1, 2, and 3 months respectively.
//     Course 4 can be taken only after course 3 is completed, i.e., after 3 months.
//     It is completed after 3 + 4 = 7 months.
//     Course 5 can be taken only after courses 1, 2, 3, and 4 have been completed, i.e., after max(1,2,3,7) = 7 months.
//     Thus, the minimum time needed to complete all the courses is 7 + 5 = 12 months.
// Constraints:
//   1 <= n <= 5 * 10^4
//   0 <= relations.length <= min(n * (n - 1) / 2, 5 * 10^4)
//   relations[j].length == 2
//   1 <= prevCoursej, nextCoursej <= n
//   prevCoursej != nextCoursej
//   All the pairs [prevCoursej, nextCoursej] are unique.
//   time.length == n
//   1 <= time[i] <= 10^4
//   The given graph is a directed acyclic graph.

// use topological-sort, when processing, update successor node's earliest start time
func minimumTime(n int, relations [][]int, time []int) int {
	graph := make([][]int, n)
	indgree := make([]int, n)
	earliest := make([]int, n)
	maxTime := 0
	for _, v := range relations {
		graph[v[0]-1] = append(graph[v[0]-1], v[1]-1)
		indgree[v[1]-1]++
	}

	// find node with 0 in-degree
	queue := make([]int, 0, n/2)
	for i := range indgree {
		if indgree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// topological sort
	for len(queue) > 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		curTime := earliest[cur] + time[cur] // current course can finish earliest at curTime
		for _, v := range graph[cur] {
			indgree[v]-- // update successor's in-degree
			if indgree[v] == 0 {
				queue = append(queue, v)
			}
			if curTime > earliest[v] { // update successor's earliest start time
				earliest[v] = curTime
			}
		}
		if curTime > maxTime { // update maxTime (be careful: last processed node may not have the max time)
			maxTime = curTime
		}
	}
	return maxTime
}

func main() {
	for _, v := range []struct {
		n   int
		r   [][]int
		t   []int
		ans int
	}{
		{3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 5}, 8},
		{5, [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}, []int{1, 2, 3, 4, 5}, 12},
		{2, [][]int{{}}, []int{3, 5}, 5},
	} {
		fmt.Println(minimumTime(v.n, v.r, v.t), v.ans)
	}
}

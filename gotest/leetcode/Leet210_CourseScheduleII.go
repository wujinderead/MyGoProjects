package main

import (
	"fmt"
)

// https://leetcode.com/problems/course-schedule-ii/

// There are a total of n courses you have to take, labeled from 0 to n-1.
// Some courses may have prerequisites, for example to take course 0 you have to
// first take course 1, which is expressed as a pair: [0,1].
// Given the total number of courses and a list of prerequisite pairs, return the
// ordering of courses you should take to finish all courses.
// There may be multiple correct orders, you just need to return one of them. If
// it is impossible to finish all courses, return an empty array.
// Example 1:
//   Input: 2, [[1,0]]
//   Output: [0,1]
//   Explanation: There are a total of 2 courses to take. To take course 1 you should have finished
//     course 0. So the correct course order is [0,1].
// Example 2:
//   Input: 4, [[1,0],[2,0],[3,1],[3,2]]
//   Output: [0,1,2,3] or [0,2,1,3]
//   Explanation: There are a total of 4 courses to take. To take course 3 you should have finished both
//     courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0.
//     So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3] .
// Note:
// The input prerequisites is a graph represented by a list of edges, not adjacency matrices.
// Read more about how a graph is represented.
// You may assume that there are no duplicate edges in the input prerequisites. 

func findOrder(numCourses int, prerequisites [][]int) []int {
    // get a topological sort order. if we can't visit all vertices there must be a circle.
    indegree := make([]int, numCourses)
    graph := make([][]int, numCourses)
	for i := range prerequisites {
		// add an edge prerequisites[i][1] -> prerequisites[i][0]
		graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
		indegree[prerequisites[i][0]]++
	}
	stack := make([]int, 0, numCourses)
	for i:=0; i<numCourses; i++ {
		if indegree[i]==0 {
			stack = append(stack, i)
		}
	}
	ans := make([]int, 0, numCourses)
	for len(stack)>0 {
		i := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, i)
		for _, j := range graph[i] {
			indegree[j]--
			if indegree[j]==0 {
				stack = append(stack, j)
			}
		}
	}
	if len(ans)<numCourses {
		return ans[:0]
	}
	return ans
}

func main() {
	fmt.Println(findOrder(2, [][]int{{1,0}}))
	fmt.Println(findOrder(4, [][]int{{1,0},{2,0},{3,1},{3,2}}))
	fmt.Println(findOrder(3, [][]int{{1,0},{0,1},{1,2}}))
}
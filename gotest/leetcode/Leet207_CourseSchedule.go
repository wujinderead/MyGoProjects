package main

import "fmt"

// https://leetcode.com/problems/course-schedule/

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses-1.
// Some courses may have prerequisites, for example to take course 0 you have to first
// take course 1, which is expressed as a pair: [0,1]
// Given the total number of courses and a list of prerequisite pairs, is it possible
// for you to finish all courses?
// Example 1:
//   Input: numCourses = 2, prerequisites = [[1,0]]
//   Output: true
//   Explanation: There are a total of 2 courses to take.
//     To take course 1 you should have finished course 0. So it is possible.
// Example 2:
//   Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
//   Output: false
//   Explanation: There are a total of 2 courses to take.
//     To take course 1 you should have finished course 0,
//     and to take course 0 you should also have finished course 1. So it is impossible.
// Constraints:
//   The input prerequisites is a graph represented by a list of edges, not adjacency matrices.
//   Read more about how a graph is represented.
//   You may assume that there are no duplicate edges in the input prerequisites.
//   1 <= numCourses <= 10^5

func canFinish(numCourses int, prerequisites [][]int) bool {
    // start from a zero-indegree vertex, and traverse the graph, check if we can visit all vertices
    // anyway, if there is a circle in the graph, we can't visit all vertices.
    indegree := make([]int, numCourses)
    visited := make([]bool, numCourses)
    graph := make([][]int, numCourses)
    for i := range prerequisites {
    	// add an edge prerequisites[i][1] -> prerequisites[i][0]
    	graph[prerequisites[i][1]] = append(graph[prerequisites[i][1]], prerequisites[i][0])
    	indegree[prerequisites[i][0]]++
	}
	for i:=0; i<numCourses; i++ {
		if !visited[i] && indegree[i]==0 {  // from an unvisited and non-prerequisite vertex
			visited[i] = true
			visit(i, graph, visited, indegree)
		}
	}
    for i := range visited {
    	if !visited[i] {
    		return false
		}
	}
	return true
}

func visit(i int, graph [][]int, visited []bool, indegree []int) {
	for _, j := range graph[i] {
		indegree[j]--       // decrease its all child
	}
	for _, j := range graph[i] {
		if !visited[j] && indegree[j]==0 {
			visited[j] = true
			visit(j, graph, visited, indegree)
		}
	}
}

func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{0, 1}, {1, 0}}))
	fmt.Println(canFinish(3, [][]int{{1, 2}, {0, 1}, {1, 0}}))
	fmt.Println(canFinish(4, [][]int{{1, 2}, {0, 1}, {1, 0}, {0, 3}}))
	fmt.Println(canFinish(5, [][]int{{1, 2}, {0, 1}, {1, 0}, {0, 3}, {3,4}, {4,3}}))
}
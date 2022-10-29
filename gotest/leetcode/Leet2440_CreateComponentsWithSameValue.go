package main

import "fmt"

// https://leetcode.com/problems/create-components-with-same-value/

// There is an undirected tree with n nodes labeled from 0 to n - 1.
// You are given a 0-indexed integer array nums of length n where nums[i] represents the value
// of the iᵗʰ node. You are also given a 2D integer array edges of length n - 1 where edges[i]
// = [ai, bi] indicates that there is an edge between nodes ai and bi in the tree.
// You are allowed to delete some edges, splitting the tree into multiple connected components.
// Let the value of a component be the sum of all nums[i] for which node i is in the component.
// Return the maximum number of edges you can delete, such that every connected component in the
// tree has the same value.
// Example 1:
//   Input: nums = [6,2,2,2,6], edges = [[0,1],[1,2],[1,3],[3,4]]
//   Output: 2
//   Explanation: The above figure shows how we can delete the edges [0,1] and [3,4].
//     The created components are nodes [0], [1,2,3] and [4]. The sum of the values
//     in each component equals 6. It can be proven that no better deletion exists, so the answer is 2.
// Example 2:
//   Input: nums = [2], edges = []
//   Output: 0
//   Explanation: There are no edges to be deleted.
// Constraints:
//   1 <= n <= 2 * 10⁴
//   nums.length == n
//   1 <= nums[i] <= 50
//   edges.length == n - 1
//   edges[i].length == 2
//   0 <= edges[i][0], edges[i][1] <= n - 1
//   edges represents a valid tree.

// https://leetcode.com/problems/create-components-with-same-value/discuss/2706628/C%2B%2B-easy-DFS
func componentValue(nums []int, edges [][]int) int {
	// get sum
	max := 0
	sum := 0
	for _, v := range nums {
		sum += v
		if v > max {
			max = v
		}
	}

	// precheck: all value is 1 or same
	if max == 1 || sum == max*len(nums) { // can delete all edges
		return len(nums) - 1
	}

	// make the graph
	graph := make([][]int, len(nums))
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}

	// check can split
	ans := 0
	for i := 2; i*i <= sum; i++ {
		if sum/i >= max && sum%i == 0 { // split to i parts, sum/i is target sum
			if canSplit(graph, nums, 0, -1, sum/i) == 0 {
				ans = i - 1 // remove i-1 edges can split the graph to i parts
			}
		}
	}
	return ans
}

func canSplit(graph [][]int, nums []int, cur, parent, target int) int {
	sum := nums[cur] // sum of the subtree
	for _, child := range graph[cur] {
		if child != parent {
			sum += canSplit(graph, nums, child, cur, target)
			if sum > target {
				return sum
			}
		}
	}
	// sum == target means current subtree can be deleted from tree.
	if sum == target {
		return 0
	}
	return sum
}

func main() {
	for _, v := range []struct {
		nums  []int
		edges [][]int
		ans   int
	}{
		{[]int{6, 2, 2, 2, 6}, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 2},
		{[]int{2}, [][]int{}, 0},
	} {
		fmt.Println(componentValue(v.nums, v.edges), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/smallest-missing-genetic-value-in-each-subtree/

// There is a family tree rooted at 0 consisting of n nodes numbered 0 to n - 1. You are given
// a 0-indexed integer array parents, where parents[i] is the parent for node i. Since node 0 is
// the root, parents[0] == -1.
// There are 10^5 genetic values, each represented by an integer in the inclusive range [1, 10^5].
// You are given a 0-indexed integer array nums, where nums[i] is a distinct genetic value for node i.
// Return an array ans of length n where ans[i] is the smallest genetic value that is missing from
// the subtree rooted at node i.
// The subtree rooted at a node x contains node x and all of its descendant nodes.
// Example 1:
//   Input: parents = [-1,0,0,2], nums = [1,2,3,4]
//   Output: [5,1,1,1]
//   Explanation: The answer for each subtree is calculated as follows:
//     - 0: The subtree contains nodes [0,1,2,3] with values [1,2,3,4]. 5 is the smallest missing value.
//     - 1: The subtree contains only node 1 with value 2. 1 is the smallest missing value.
//     - 2: The subtree contains nodes [2,3] with values [3,4]. 1 is the smallest missing value.
//     - 3: The subtree contains only node 3 with value 4. 1 is the smallest missing value.
// Example 2:
//   Input: parents = [-1,0,1,0,3,3], nums = [5,4,6,2,1,3]
//   Output: [7,1,1,4,2,1]
//   Explanation: The answer for each subtree is calculated as follows:
//     - 0: The subtree contains nodes [0,1,2,3,4,5] with values [5,4,6,2,1,3]. 7 is the smallest missing value.
//     - 1: The subtree contains nodes [1,2] with values [4,6]. 1 is the smallest missing value.
//     - 2: The subtree contains only node 2 with value 6. 1 is the smallest missing value.
//     - 3: The subtree contains nodes [3,4,5] with values [2,1,3]. 4 is the smallest missing value.
//     - 4: The subtree contains only node 4 with value 1. 2 is the smallest missing value.
//     - 5: The subtree contains only node 5 with value 3. 1 is the smallest missing value.
// Example 3:
//   Input: parents = [-1,2,3,0,2,4,1], nums = [2,3,4,5,6,7,8]
//   Output: [1,1,1,1,1,1,1]
//   Explanation: The value 1 is missing from all the subtrees.
// Constraints:
//   n == parents.length == nums.length
//   2 <= n <= 10^5
//   0 <= parents[i] <= n - 1 for i != 0
//   parents[0] == -1
//   parents represents a valid tree.
//   1 <= nums[i] <= 10^5
//   Each nums[i] is distinct.

// just consider the nodes on the path from the value-1 node to root,
// for nodes on this path, dfs to get all sub-nodes and put them in a set
func smallestMissingValueSubtree(parents []int, nums []int) []int {
	// make a graph
	graph := make(map[int][]int)
	for i, v := range parents {
		graph[v] = append(graph[v], i)
	}

	// make answer
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = 1
	}

	// find index of value 1
	cur := -1
	for i := range nums {
		if nums[i] == 1 {
			cur = i
			break
		}
	}
	if cur == -1 { // no node contain 1, return
		return ans
	}

	// for the nodes on the path from the value-1 node to root, find all its values
	pre := -1
	firstMiss := 1
	set := make(map[int]struct{})
	for cur != -1 {
		set[nums[cur]] = struct{}{}        // add current node's value to set
		for _, child := range graph[cur] { // dfs current node's unvisited child
			if child == pre {
				continue
			}
			dfs(graph, set, nums, child) // add all sub-tree node's value to set
		}
		for { // find first miss
			if _, ok := set[firstMiss]; !ok {
				break
			}
			firstMiss++
		}
		ans[cur] = firstMiss
		pre = cur // move current node to its parent, current node become pre
		cur = parents[cur]
	}
	return ans
}

func dfs(graph map[int][]int, set map[int]struct{}, nums []int, ind int) {
	set[nums[ind]] = struct{}{}
	for _, v := range graph[ind] {
		dfs(graph, set, nums, v)
	}
}

func main() {
	for _, v := range []struct {
		p, n, ans []int
	}{
		{[]int{-1, 0, 0, 2}, []int{1, 2, 3, 4}, []int{5, 1, 1, 1}},
		{[]int{-1, 0, 1, 0, 3, 3}, []int{5, 4, 6, 2, 1, 3}, []int{7, 1, 1, 4, 2, 1}},
		{[]int{-1, 2, 3, 0, 2, 4, 1}, []int{2, 3, 4, 5, 6, 7, 8}, []int{1, 1, 1, 1, 1, 1, 1}},
	} {
		fmt.Println(smallestMissingValueSubtree(v.p, v.n), v.ans)
	}
}

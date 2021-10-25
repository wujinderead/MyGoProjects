package main

import "fmt"

// https://leetcode.com/problems/count-nodes-with-the-highest-score/

// There is a binary tree rooted at 0 consisting of n nodes. The nodes are labeled from 0 to n - 1.
// You are given a 0-indexed integer array parents representing the tree, where parents[i] is the
// parent of node i. Since node 0 is the root, parents[0] == -1.
// Each node has a score. To find the score of a node, consider if the node and the edges connected
// to it were removed. The tree would become one or more non-empty subtrees. The size of a subtree
// is the number of the nodes in it. The score of the node is the product of the sizes of all those subtrees.
// Return the number of nodes that have the highest score.
// Example 1:
//   Input: parents = [-1,2,0,2,0]
//   Output: 3
//   Explanation:
//     - The score of node 0 is: 3 * 1 = 3
//     - The score of node 1 is: 4 = 4
//     - The score of node 2 is: 1 * 1 * 2 = 2
//     - The score of node 3 is: 4 = 4
//     - The score of node 4 is: 4 = 4
//     The highest score is 4, and three nodes (node 1, node 3, and node 4) have the highest score.
// Example 2:
//   Input: parents = [-1,2,0]
//   Output: 2
//   Explanation:
//     - The score of node 0 is: 2 = 2
//     - The score of node 1 is: 2 = 2
//     - The score of node 2 is: 1 * 1 = 1
//     The highest score is 2, and two nodes (node 0 and node 1) have the highest score.
// Constraints:
//   n == parents.length
//   2 <= n <= 10^5
//   parents[0] == -1
//   0 <= parents[i] <= n - 1 for i != 0
//   parents represents a valid binary tree.

// visit the tree from leaves using topological-sort like manner
func countHighestScoreNodes(parents []int) int {
	ndc := make([]int, len(parents)) // number of direct children for each node: 0, 1, or 2
	nsc := make([]int, len(parents)) // subtree size rooted at node[i]
	nlc := make([]int, len(parents)) // subtree size rooted at node[i].one_child
	nrc := make([]int, len(parents)) // subtree size rooted at node[i].other_child

	// count direct children
	for _, p := range parents {
		if p == -1 {
			continue
		}
		ndc[p]++
	}

	// find all leaves
	queue := make([]int, 0, len(parents)/2+2)
	for i := range ndc {
		if ndc[i] == 0 {
			queue = append(queue, i) // push all leaves
		}
	}

	// topological sort
	max := 0
	scoreMap := make(map[int]int)
	for len(queue) > 0 {
		i := queue[len(queue)-1]
		queue = queue[:len(queue)-1] // got a node
		cur := 1
		remain := len(parents) - 1 - nlc[i] - nrc[i] // the tree nodes exclude current node and its children
		if remain > 0 {
			cur *= remain
		}
		if nlc[i] > 0 { // multiply child subtree
			cur *= nlc[i]
		}
		if nrc[i] > 0 {
			cur *= nrc[i]
		}
		if cur > max { // update max value
			max = cur
		}
		scoreMap[cur] = scoreMap[cur] + 1
		nsc[i] = 1 + nlc[i] + nrc[i] // current node subtree size
		if i == 0 {                  // find root, can break
			break
		}
		parent := parents[i]
		ndc[parent]--         // decrease parent's direct child
		if ndc[parent] == 0 { // if all children of parent traversed, push to queue
			queue = append(queue, parent)
		}
		if nlc[parent] == 0 {
			nlc[parent] = nsc[i]
		} else {
			nrc[parent] = nsc[i]
		}
	}
	return scoreMap[max]
}

func main() {
	for _, v := range []struct {
		p   []int
		ans int
	}{
		{[]int{-1, 2, 0, 2, 0}, 3},
		{[]int{-1, 2, 0}, 2},
	} {
		fmt.Println(countHighestScoreNodes(v.p), v.ans)
	}
}

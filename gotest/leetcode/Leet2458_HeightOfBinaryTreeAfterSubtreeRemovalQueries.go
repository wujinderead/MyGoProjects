package main

import "fmt"

// https://leetcode.com/problems/height-of-binary-tree-after-subtree-removal-queries/

// You are given the root of a binary tree with n nodes. Each node is assigned a unique value from 1 to n.
// You are also given an array queries of size m.
// You have to perform m independent queries on the tree where in the iᵗʰ query you do the following:
// Remove the subtree rooted at the node with the value queries[i] from the tree.
// It is guaranteed that queries[i] will not be equal to the value of the root.
// Return an array answer of size m where answer[i] is the height of the tree after performing the
// iᵗʰ query.
// Note:
//   The queries are independent, so the tree returns to its initial state after each query.
//   The height of a tree is the number of edges in the longest simple path from the root to
//     some node in the tree.
// Example 1:
//   Input: root = [1,3,4,2,null,6,5,null,null,null,null,null,7], queries = [4]
//   Output: [2]
//   Explanation: The diagram above shows the tree after removing the subtree rooted
//     at node with value 4.
//     The height of the tree is 2 (The path 1 -> 3 -> 2).
// Example 2:
//   Input: root = [5,8,9,2,1,3,7,4,6], queries = [3,2,4,8]
//   Output: [3,2,3,2]
//   Explanation: We have the following queries:
//     - Removing the subtree rooted at node with value 3. The height of the tree
//     becomes 3 (The path 5 -> 8 -> 2 -> 4).
//     - Removing the subtree rooted at node with value 2. The height of the tree
//     becomes 2 (The path 5 -> 8 -> 1).
//     - Removing the subtree rooted at node with value 4. The height of the tree
//     becomes 3 (The path 5 -> 8 -> 2 -> 6).
//     - Removing the subtree rooted at node with value 8. The height of the tree
//     becomes 2 (The path 5 -> 9 -> 3).
// Constraints:
//   The number of nodes in the tree is n.
//   2 <= n <= 10⁵
//   1 <= Node.val <= n
//   All the values in the tree are unique.
//   m == queries.length
//   1 <= m <= min(n, 10⁴)
//   1 <= queries[i] <= n
//   queries[i] != root.val

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	ans := make([]int, len(queries))
	height := make(map[int]int)
	getHeight(root, height)

	// dfs
	maxHei := make(map[int]int)
	dfs(root, 0, 0, maxHei, height)

	// get answer
	for i := range queries {
		ans[i] = maxHei[queries[i]]
	}
	return ans
}

// depth: current node depth
// curMax: max height after remove current node
func dfs(r *TreeNode, depth, curMax int, maxHei, height map[int]int) {
	if r == nil {
		return
	}
	maxHei[r.Val] = curMax
	dfs(r.Left, depth+1, max(curMax, depth+getHeight(r.Right, height)), maxHei, height)
	dfs(r.Right, depth+1, max(curMax, depth+getHeight(r.Left, height)), maxHei, height)
}

func getHeight(r *TreeNode, height map[int]int) int {
	if r == nil {
		return 0
	}
	if v, ok := height[r.Val]; ok {
		return v
	}
	height[r.Val] = max(getHeight(r.Left, height), getHeight(r.Right, height)) + 1
	return height[r.Val]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Right = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(treeQueries(root, []int{1, 2, 3, 4, 6, 7, 8, 9}))
	fmt.Println([]int{3, 2, 3, 3, 3, 3, 2, 3})
}

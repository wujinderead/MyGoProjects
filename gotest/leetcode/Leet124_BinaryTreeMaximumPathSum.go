package main

import "fmt"

// https://leetcode.com/problems/binary-tree-maximum-path-sum

// Given a non-empty binary tree, find the maximum path sum. 
// For this problem, a path is defined as any sequence of nodes from some starting node to any 
// node in the tree along the parent-child connections. The path must contain at least one node 
// and does not need to go through the root. 
// Example 1: 
//   Input: [1,2,3]
//       1
//      / \
//     2   3
//   Output: 6
// Example 2: 
//   Input: [-10,9,20,null,null,15,7]
//      -10
//      / \
//     9  20
//       /  \
//      15   7
//   Output: 42

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root==nil {
		return 0
	}
	m := -0x7fffffff
	helper(root, &m)
	return m
}

func helper(t *TreeNode, m *int) int {
	if t==nil {
		return 0
	}
	l, r, v := helper(t.Left, m), helper(t.Right, m), t.Val
	mm := max(max(v, v+l), max(v+r, v+l+r))   // the best we can contribute to max-path-sum
	*m = max(*m, mm)
	return max(max(0, v), max(v+l, v+r))  // the best we can provide to parent
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	r := &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	fmt.Println(maxPathSum(r))

	r = &TreeNode{Val: -10}
	r.Left = &TreeNode{Val: 9}
	r.Right = &TreeNode{Val: 20}
	r.Right.Left = &TreeNode{Val: 15}
	r.Right.Right = &TreeNode{Val: 7}
	fmt.Println(maxPathSum(r))

	r = &TreeNode{Val: -3}
	r.Left = &TreeNode{Val: -2}
	r.Right = &TreeNode{Val: -4}
	fmt.Println(maxPathSum(r))
}
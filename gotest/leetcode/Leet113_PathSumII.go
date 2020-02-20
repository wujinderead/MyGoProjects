package main

import (
	"fmt"
)

// https://leetcode.com/problems/path-sum-ii/

// Given a binary tree and a sum, find all root-to-leaf paths where
// each path's sum equals the given sum. Note: A leaf is a node with no children.
// Example:
//   Given the below binary tree and sum = 22,
//           5
//          / \
//         4   8
//        /   / \
//       11  13  4
//      /  \    / \
//     7    2  5   1
//   Return:
//     [
//      [5,4,11,2],
//      [5,8,4,5]
//     ]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	// post-order traverse, when visit leaf, check the sum
	cursum := 0
	cur := root
	stack := make([]*TreeNode, 0, 10)
	paths := make([][]int, 0)
	var lastvisit *TreeNode = nil
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cursum += cur.Val
			cur = cur.Left
		} else {
			peek := stack[len(stack)-1]
			if peek.Right != nil && lastvisit != peek.Right {
				cur = peek.Right
			} else {
				// visit peek
				if peek.Left == nil && peek.Right == nil && cursum == sum {
					path := make([]int, len(stack))
					for i := range stack {
						path[i] = stack[i].Val
					}
					paths = append(paths, path)
				}
				lastvisit = peek
				cursum -= peek.Val
				stack = stack[:len(stack)-1]
			}
		}
	}
	return paths
}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 1}
	fmt.Println(pathSum(root, 22))
}

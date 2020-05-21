package main

import (
	"fmt"
)

// https://leetcode.com/problems/binary-tree-postorder-traversal/

// Given a binary tree, return the postorder traversal of its nodes' values.
// Example:
//   Input: [1,null,2,3]
//     1
//      \
//       2
//      /
//     3
//   Output: [3,2,1]
// Follow up: Recursive solution is trivial, could you do it iteratively? 

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
    stack := make([]*TreeNode, 0, 10)
    buf := make([]int, 0, 10)
    cur := root
	lastVisited := (*TreeNode)(nil)
    for cur != nil || len(stack)>0 {
    	if cur != nil {
    		stack = append(stack, cur)
    		cur = cur.Left
		} else {
			peek := stack[len(stack)-1]
			if peek.Right != nil && peek.Right != lastVisited {
				cur = peek.Right
			} else {
				buf = append(buf, peek.Val)
				stack = stack[:len(stack)-1]
				lastVisited = peek
			}
		}
	}
	return buf
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(postorderTraversal(root))
	fmt.Println(postorderTraversal(nil))
	root = &TreeNode{Val: 1}
	fmt.Println(postorderTraversal(root))
}
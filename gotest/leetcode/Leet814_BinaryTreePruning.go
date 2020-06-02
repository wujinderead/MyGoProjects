package main

import "fmt"

// https://leetcode.com/problems/binary-tree-pruning/

// We are given the head node root of a binary tree, where additionally every node's value is either a 0 or a 1. 
// Return the same tree where every subtree (of the given tree) not containing a 1 has been removed. 
// (Recall that the subtree of a node X is X, plus every node that is a descendant of X.) 
// Example 1:
//   Input: [1,null,0,0,1]
//   Output: [1,null,0,null,1]
//        1               1
//         \               \
//          0      ->       0
//         / \               \
//        0   1               1
//  Explanation: 
//     Only the red nodes satisfy the property "every subtree not containing a 1".
//     The diagram on the right represents the answer.
// Example 2:
//   Input: [1,0,1,0,0,0,1]
//   Output: [1,null,1,null,1]
// Example 3:
//   Input: [1,1,0,1,1,0,1,0]
//   Output: [1,1,0,1,1,null,1]
// Note: 
//   The binary tree will have at most 100 nodes. 
//   The value of each node will only be 0 or 1. 

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
 }

func pruneTree(root *TreeNode) *TreeNode {
    dummy := &TreeNode{}
    dummy.Left = root
    prune(dummy)
    return dummy.Left
}

func prune(r *TreeNode) bool {  // if all 0 return true, else false
	if r==nil {
		return true
	}
	t1, t2 := prune(r.Left), prune(r.Right)
	if t1 {
		r.Left = nil
	}
	if t2 {
		r.Right = nil
	}
	return r.Val==0 && t1 && t2
}

func main() {
	r := &TreeNode{Val: 1}
	r.Right = &TreeNode{Val: 0}
	r.Right.Left = &TreeNode{Val: 0}
	r.Right.Right = &TreeNode{Val: 1}
	fmt.Println(pruneTree(r))

	r = &TreeNode{Val: 1}
	fmt.Println(pruneTree(r))

	r = &TreeNode{Val: 0}
	fmt.Println(pruneTree(r))

	r = &TreeNode{Val: 0}
	r.Right = &TreeNode{Val: 1}
	fmt.Println(pruneTree(r))
}
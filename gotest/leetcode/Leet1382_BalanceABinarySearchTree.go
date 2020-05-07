package main

import "fmt"

// https://leetcode.com/problems/balance-a-binary-search-tree/

// Given a binary search tree, return a balanced binary search tree with the same node values.
// A binary search tree is balanced if and only if the depth of the two subtrees of every node
// never differ by more than 1. If there is more than one answer, return any of them.
// Example 1:
//   Input: root = [1,null,2,null,3,null,4,null,null]
//   Output: [2,1,3,null,null,null,4]
//         1                      2
//          \                    / \
//           2                  1   3
//            \                      \
//             3                      4
//              \
//               4
//   Explanation: This is not the only correct answer, [3,1,4,null,2,null,null] is also correct.
// Constraints:
//   The number of nodes in the tree is between 1 and 10^4.
//   The tree nodes will have distinct values between 1 and 10^5.

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
    // traverse the tree to get all elements, construct a BST based on all elements
    cur := root
    stack, vals := make([]*TreeNode, 0), make([]int, 0)
    for cur != nil || len(stack)>0 {
    	if cur != nil {
    		stack = append(stack, cur)
    		cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			vals = append(vals, cur.Val)
			cur = cur.Right
		}
	}
	return makeBST(vals)
}

func makeBST(vals []int) *TreeNode {
	if len(vals)==1 {
		return &TreeNode{Val: vals[0]}
	}
	cur := &TreeNode{Val: vals[len(vals)/2]}
	cur.Left = makeBST(vals[:len(vals)/2])
	if len(vals)/2+1 < len(vals) {
		cur.Right = makeBST(vals[len(vals)/2+1:])
	}
	return cur
}

func main() {
	r := &TreeNode{Val: 3}
	r.Right = &TreeNode{Val: 4}
	r.Right.Right = &TreeNode{Val: 5}
	r.Left = &TreeNode{Val: 2}
	r.Left.Left = &TreeNode{Val: 1}
	t := balanceBST(r)
	fmt.Println(t)
	fmt.Println(t.Left)
	fmt.Println(t.Right)
	fmt.Println(t.Left.Left)
	fmt.Println(t.Left.Right)
	fmt.Println(t.Right.Left)
	fmt.Println(t.Right.Right)
}
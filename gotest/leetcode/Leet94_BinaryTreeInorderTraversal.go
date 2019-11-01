package main

import "fmt"

// https://leetcode.com/problems/binary-tree-inorder-traversal/

// Given a binary tree, return the inorder traversal of its nodes' values.
// Example:
//   Input: [1,null,2,3]
//     1
//      \
//       2
//      /
//     3
//
//   Output: [1,3,2]
// Follow up: Recursive solution is trivial, could you do it iteratively?

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorder traversal, left-root-right. stack needed for iterative solution
func inorderTraversal(root *TreeNode) []int {
	re := make([]int, 0)
	c := root
	s := newStack()
	for c != nil || !s.empty() {
		for c != nil {
			s.push(c)
			c = c.Left
		}
		c = s.pop()
		re = append(re, c.Val)
		c = c.Right
	}
	return re
}

type stack struct {
	arr []*TreeNode
}

func newStack() *stack {
	return &stack{make([]*TreeNode, 0)}
}

func (s *stack) push(node *TreeNode) {
	s.arr = append(s.arr, node)
}

func (s *stack) pop() *TreeNode {
	a := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return a
}

func (s *stack) empty() bool {
	return len(s.arr) == 0
}

func main() {
	r := &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Left.Left = &TreeNode{Val: 4}
	r.Left.Right = &TreeNode{Val: 5}
	r.Right.Left = &TreeNode{Val: 6}
	r.Right.Right = &TreeNode{Val: 7}
	fmt.Println(inorderTraversal(r))
}

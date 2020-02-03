package main

import "fmt"

// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

// Given a binary tree, flatten it to a linked list in-place.
// For example, given the following tree:
//
//       1
//      / \
//     2   5
//    / \   \
//   3   4   6
// The flattened tree should look like:
//   1
//    \
//     2
//      \
//       3
//        \
//         4
//          \
//           5
//            \
//             6
//
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// pre-order traverse
	stack := make([]*TreeNode, 1, 10)
	last := &TreeNode{} // a dummy head
	dummy := last
	stack[0] = root
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
		last.Right = cur
		cur.Left = nil
		last = cur
	}
	dummy.Right = nil // gc dummy
	dummy = nil
}

func main() {
	//       1
	//      / \
	//     2   5
	//    / \   \
	//   3   4   6
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 6}
	flatten(root)
	r := root
	for r != nil {
		fmt.Println(r.Val, r.Left)
		r = r.Right
	}
	fmt.Println()

	//           5
	//          / \
	//         4   8
	//        /   / \
	//       11  13  4
	//      /  \    / \
	//     7    2  5   1
	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 11}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 1}
	flatten(root)
	r = root
	for r != nil {
		fmt.Println(r.Val, r.Left)
		r = r.Right
	}
	fmt.Println()

	//      1
	//       \
	//        2
	//       /
	//      3
	//     /
	//    4
	//     \
	//      5
	root = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Left.Left = &TreeNode{Val: 4}
	root.Right.Left.Right = &TreeNode{Val: 5}
	flatten(root)
	r = root
	for r != nil {
		fmt.Println(r.Val, r.Left)
		r = r.Right
	}
}

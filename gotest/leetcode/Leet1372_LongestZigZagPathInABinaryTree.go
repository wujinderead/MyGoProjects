package main

import "fmt"

// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/

// Given a binary tree root, a ZigZag path for a binary tree is defined as follow:
//   Choose any node in the binary tree and a direction (right or left).
//   If the current direction is right then move to the right child of the current
//     node otherwise move to the left child.
//   Change the direction from right to left or right to left.
//   Repeat the second and third step until you can't move in the tree.
// Zigzag length is defined as the number of nodes visited - 1.
// (A single node has a length of 0).
// Return the longest ZigZag path contained in that tree.
// Example 1: 
//   Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1,null,1]
//       1
//        \
//         1
//        / \ R
//       1   1
//        L / \
//         1   1
//          \ R
//           1
//            \
//             1
//   Output: 3
//   Explanation: Longest ZigZag path in blue nodes (right -> left -> right).
// Example 2: 
//   Input: root = [1,1,1,null,1,null,null,1,1,null,1]
//             1
//          L / \
//           1   1
//          R \
//             1
//          L / \
//           1   1
//          R \
//             1
//   Output: 4
//   Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).
// Example 3:
//   Input: root = [1]
//   Output: 0
// Constraints:
//   Each tree has at most 50000 nodes..
//   Each node's value is between [1, 100].

type TreeNode struct {
 	Val int
	Left *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	max := 0
    helper(root, &max)
    return max
}

func helper(t *TreeNode, max *int) (int, int) {
	if t==nil {
		return 0, 0
	}
	l, r := 0, 0
	if t.Left != nil {
		_, rr := helper(t.Left, max)
		l = rr+1
	}
	if t.Right != nil {
		ll, _ := helper(t.Right, max)
		r = ll+1
	}
	if r>*max {
		*max=r
	}
	if l>*max {
		*max=l
	}
	return l, r
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 1}
	root.Right.Right.Left = &TreeNode{Val: 1}
	root.Right.Right.Right = &TreeNode{Val: 1}
	root.Right.Right.Left.Right = &TreeNode{Val: 1}
	root.Right.Right.Left.Right.Right = &TreeNode{Val: 1}
	fmt.Println(longestZigZag(root))

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 1}
	root.Left.Right.Left = &TreeNode{Val: 1}
	root.Left.Right.Right = &TreeNode{Val: 1}
	root.Left.Right.Left.Right = &TreeNode{Val: 1}
	fmt.Println(longestZigZag(root))

	root = &TreeNode{Val: 1}
	fmt.Println(longestZigZag(root))
}
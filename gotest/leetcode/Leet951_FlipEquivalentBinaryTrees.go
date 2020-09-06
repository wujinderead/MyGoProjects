package main

import "fmt"

// https://leetcode.com/problems/flip-equivalent-binary-trees

// For a binary tree T, we can define a flip operation as follows: choose any node,
// and swap the left and right child subtrees.
// A binary tree X is flip equivalent to a binary tree Y if and only if we can make
// X equal to Y after some number of flip operations.
// Given the roots of two binary trees root1 and root2, return true if the two trees
// are flip equivalent or false otherwise.
// Example 1: 
//             1                 1     
//           /   \             /   \       
//          2     3           3     2
//        /  \   /             \   / \
//       4   5  6              6  4   5
//          / \                      / \
//         7  8                     8  7
//   Input: root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
//   Output: true
//   Explanation: We flipped at nodes with values 1, 3, and 5.
// Example 2:
//   Input: root1 = [], root2 = []
//   Output: true
// Example 3:
//   Input: root1 = [], root2 = [1]
//   Output: false
// Example 4:
//   Input: root1 = [0,null,1], root2 = []
//   Output: false
// Example 5:
//   Input: root1 = [0,null,1], root2 = [0,1]
//   Output: true
// Constraints:
//   The number of nodes in each tree is in the range [0, 100].
//   Each value in each tree will be a unique integer in the range [0, 99].

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    if root1==nil && root2==nil {   // both nil, return true
		return true
	} else if (root1==nil || root2==nil) || root1.Val!=root2.Val {
		return false     // one nil, or both non-nil with different values
	}
	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}

func main() {
	r1 := &TreeNode{Val: 1}
	r1.Left = &TreeNode{Val: 2}
	r1.Left.Left = &TreeNode{Val: 4}
	r1.Left.Right = &TreeNode{Val: 5}
	r1.Left.Right.Left = &TreeNode{Val: 7}
	r1.Left.Right.Right = &TreeNode{Val: 8}
	r1.Right = &TreeNode{Val: 3}
	r1.Right.Left = &TreeNode{Val: 6}

	r2 := &TreeNode{Val: 1}
	r2.Left = &TreeNode{Val: 3}
	r2.Left.Right = &TreeNode{Val: 6}
	r2.Right = &TreeNode{Val: 2}
	r2.Right.Left = &TreeNode{Val: 4}
	r2.Right.Right = &TreeNode{Val: 5}
	r2.Right.Right.Left = &TreeNode{Val: 8}
	r2.Right.Right.Right = &TreeNode{Val: 7}
	fmt.Println(flipEquiv(r1, r2), true)

	fmt.Println(flipEquiv(nil, nil), true)

	r2 = &TreeNode{Val: 1}
	fmt.Println(flipEquiv(nil, r2), false)

	r1 = &TreeNode{Val: 1}
	r1.Right = &TreeNode{Val: 2}
	fmt.Println(flipEquiv(r1, nil), false)

	r2.Left = &TreeNode{Val: 2}
	fmt.Println(flipEquiv(r1, r2), true)
}
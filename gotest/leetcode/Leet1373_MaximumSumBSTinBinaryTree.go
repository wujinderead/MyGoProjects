package main

import "fmt"

// Given a binary tree root, the task is to return the maximum sum of all keys of
// any sub-tree which is also a Binary Search Tree (BST).
// Assume a BST is defined as follows:
// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.
// Example 1:
//   Input: root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
//   Output: 20
//   Explanation: Maximum sum in a valid Binary search tree is obtained in root
//     node with key equal to 3.
//               1
//            /    \
//          4        3   <-----  the result bst rooted at 3
//        /   \    /  \
//       2    4   2    5
//                    / \
//                   4   6
// Example 2:
//   Input: root = [4,3,null,1,2]
//   Output: 2
//   Explanation: Maximum sum in a valid Binary search tree is obtained in a single
//     root node with key equal to 2.
//                     4
//                   /
//                 3
//                / \
//               1   2    <----- the largest sum bst is 2
// Example 3:
//   Input: root = [-4,-2,-5]
//   Output: 0
//   Explanation: All values are negatives. Return an empty BST.
// Example 4:
//   Input: root = [2,1,3]
//   Output: 6
// Example 5:
//   Input: root = [5,4,8,3,null,6,3]
//   Output: 7
// Constraints:
//   Each tree has at most 40000 nodes..
//   Each node's value is between [-4 * 10^4 , 4 * 10^4].

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxSumBST(root *TreeNode) int {
	max := 0
	testBst(root, &max)
	return max
}

func testBst(t *TreeNode, max *int) (sum int, isBst bool, maxv, minv int) {
	leftBst, leftSum, rightBst, rightSum := true, 0, true, 0
	lmax, lmin, rmax, rmin := t.Val, t.Val, t.Val, t.Val
	if t.Left != nil {
		leftSum, leftBst, lmax, lmin = testBst(t.Left, max)
		leftBst = leftBst && lmax < t.Val // bst need t.Val larger than the max value in t.Left
	}
	if t.Right != nil {
		rightSum, rightBst, rmax, rmin = testBst(t.Right, max)
		rightBst = rightBst && t.Val < rmin // bst need t.Val less than the max value in t.Right
	}
	if leftBst && rightBst {
		sum += leftSum + rightSum + t.Val
		if sum > *max {
			*max = sum
		}
	}
	return sum, leftBst && rightBst, rmax, lmin
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 5}
	root.Right.Right.Left = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 6}
	fmt.Println(maxSumBST(root))

	root = &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	fmt.Println(maxSumBST(root))

	root = &TreeNode{Val: -4}
	root.Left = &TreeNode{Val: -5}
	root.Right = &TreeNode{Val: -2}
	fmt.Println(maxSumBST(root))

	root = &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(maxSumBST(root))

	root = &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 9}
	fmt.Println(maxSumBST(root))
}

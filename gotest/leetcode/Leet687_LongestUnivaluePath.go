package main

import "fmt"

// https://leetcode.com/problems/longest-univalue-path/

// Given a binary tree, find the length of the longest path where each node
// in the path has the same value. This path may or may not pass through the root.
// The length of path between two nodes is represented by the number of edges between them.
// Example 1:
//   Input:
//         5
//        / \
//       4   5
//      / \   \
//     1   1   5
//   Output: 2
// Example 2:
//   Input:
//         1
//        / \
//       4   5
//      / \   \
//     4   4   5
//   Output: 2
// Note:
//   The given binary tree has not more than 10000 nodes.
//   The height of the tree is not more than 1000.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	_, ml, _ := helper(root)
	if ml > 0 {
		return ml - 1
	}
	return 0
}

func helper(t *TreeNode) (selflen, maxlen, maxele int) {
	if t == nil {
		return 0, 0, -1
	}
	lelelen, lmaxlen, lmaxele := helper(t.Left)
	relelen, rmaxlen, rmaxele := helper(t.Right)
	selflen1 := 1
	selflen2 := 1
	selfmax := 1
	if t.Left != nil && t.Left.Val == t.Val {
		selflen1 += lelelen
		selfmax += lelelen
	}
	if t.Right != nil && t.Right.Val == t.Val {
		selflen2 += relelen
		selfmax += relelen
	}
	selflen = max(selflen1, selflen2)
	maxlen = selfmax
	maxele = t.Val
	if lmaxlen > maxlen {
		maxlen = lmaxlen
		maxele = lmaxele
	}
	if rmaxlen > maxlen {
		maxlen = rmaxlen
		maxele = rmaxele
	}
	fmt.Println(t.Val, selflen, maxlen, maxele)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//     5
	//    / \
	//   4   5
	//  / \   \
	// 1   1   5
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 5}
	fmt.Println(longestUnivaluePath(root))
	fmt.Println()

	//     1
	//    / \
	//   4   5
	//  / \   \
	// 4   4   5
	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	fmt.Println(longestUnivaluePath(root))
	fmt.Println()

	//           3
	//          /
	//         3
	//        / \
	//       3   3
	//      / \   \
	//     3   4   3
	//    / \   \
	//   2  3    4
	//  / \  \   /
	// 2  2  3  4
	//    \     \
	//    2      4
	//            \
	//            4
	root = &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	root.Left.Right.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Right.Right = &TreeNode{Val: 4}
	root.Left.Left.Right.Right.Left = &TreeNode{Val: 4}
	root.Left.Left.Right.Right.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Right.Right.Left.Right.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Left.Right = &TreeNode{Val: 3}
	root.Left.Left.Left.Right.Right = &TreeNode{Val: 3}
	root.Left.Left.Left.Left = &TreeNode{Val: 2}
	root.Left.Left.Left.Left.Left = &TreeNode{Val: 2}
	root.Left.Left.Left.Left.Right = &TreeNode{Val: 2}
	root.Left.Left.Left.Left.Right.Right = &TreeNode{Val: 2}
	fmt.Println(longestUnivaluePath(root))

}

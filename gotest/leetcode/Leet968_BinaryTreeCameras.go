package main

import (
    "fmt"
)

// https://leetcode.com/problems/binary-tree-cameras/

// Given a binary tree, we install cameras on the nodes of the tree. 
// Each camera at a node can monitor its parent, itself, and its immediate children. 
// Calculate the minimum number of cameras needed to monitor all nodes of the tree. 
// Example 1: 
//             0
//            / 
//           cam
//           / \
//          0   0
//   Input: [0,0,null,0,0]
//   Output: 1
//   Explanation: One camera is enough to monitor all nodes if placed as shown.
// Example 2: 
//   Input: [0,0,null,0,null,0,null,null,0]
//   Output: 2
//   Explanation: At least two cameras are needed to monitor all nodes of the tree.
//     The above image shows one of the valid configurations of camera placement.
// Note: 
//   The number of nodes in the given tree will be in the range [1, 1000]. 
//   Every node has value 0. 

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// greedy solution:
// https://leetcode.com/problems/binary-tree-cameras/discuss/211180/JavaC%2B%2BPython-Greedy-DFS
//   Set cameras on all leaves' parents, thenremove all covered nodes.
//   Repeat step 1 until all nodes are covered.

func minCameraCover(root *TreeNode) int {
	_, a, b := helper(root)
	return min(a, b)
}

//   state1:      state2:        state3:
//     X            X              Y
//    / \          / \            / \
//   X   X        X   Y          ?   ?
// we assume the children have been covered by lower nodes. 
// state1 means we don't cover current node by self or child, but by parent.
// then the state for current node and child nodes are the following situations:
//   state1:             state2:                      state3:
//     X              X          X                        Y
//    / \          /    \     /    \               /            \
//   s2 s2      (s2,s3) s3   s3 (s2,s3)      (s1,s2,s3)    (s1,s2,s3)
func helper(t *TreeNode) (int, int, int) {
	if t==nil {
		return 0, 0, 9999
	}
	l1, l2, l3 := helper(t.Left)
	r1, r2, r3 := helper(t.Right)
	lmin, rmin := min(l2, l3), min(r2, r3)
	s1 := l2+r2
	s2 := min(lmin+r3, l3+rmin)
	s3 := 1+min(l1, lmin) + min(r1, rmin)
	return s1, s2, s3
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	r := &TreeNode{}
	r.Left = &TreeNode{}
	r.Left.Left = &TreeNode{}
	r.Left.Right = &TreeNode{}
	fmt.Println(minCameraCover(r))

	r = &TreeNode{}
	r.Left = &TreeNode{}
	r.Left.Left = &TreeNode{}
	r.Left.Left.Left = &TreeNode{}
	r.Left.Left.Left.Right = &TreeNode{}
	fmt.Println(minCameraCover(r))

	r = &TreeNode{}
	fmt.Println(minCameraCover(r))

	r = &TreeNode{}
	r.Left = &TreeNode{}
	r.Right = &TreeNode{}
	r.Left.Left = &TreeNode{}
	fmt.Println(minCameraCover(r))

	r = &TreeNode{}
	r.Left = &TreeNode{}
	r.Right = &TreeNode{}
	r.Left.Left = &TreeNode{}
	r.Left.Left.Left = &TreeNode{}
	r.Left.Left.Right = &TreeNode{}
	r.Right.Left = &TreeNode{}
	r.Right.Right = &TreeNode{}
	fmt.Println(minCameraCover(r))

	r = &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Left.Left = &TreeNode{Val: 3}
	r.Left.Left.Left = &TreeNode{Val: 4}
	r.Left.Left.Left.Right = &TreeNode{Val: 5}
	r.Left.Left.Left.Right.Left = &TreeNode{Val: 6}
	fmt.Println(minCameraCover(r))
}
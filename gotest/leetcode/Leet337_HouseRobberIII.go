package main

import (
	"fmt"
)

// https://leetcode.com/problems/house-robber-iii/

// The thief has found himself a new place for his thievery again.
// There is only one entrance to this area, called the "root".
// Besides the root, each house has one and only one parent house.
// After a tour, the smart thief realized that "all houses in this place forms a binary tree".
// It will automatically contact the police if two directly-linked houses were broken into on the same night.
// Determine the maximum amount of money the thief can rob tonight without alerting the police.
// Example 1:
//   Input: [3,2,3,null,3,null,1]
//        3
//       / \
//      2   3
//       \   \
//        3   1
//   Output: 7
//   Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
// Example 2:
//   Input: [3,4,5,1,3,null,1]
//        3
//       / \
//      4   5
//     / \   \
//    1   3   1
//   Output: 9
//   Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r, c := robHelper(root)
	fmt.Println(r, c)
	return max(r, c)
}

// for tree (1), a1 and a2 are equivalent that can be replaced by a1+a2
// so it can be transferred (2).
//   (1)      x            (2)     x                   (3)
//          /   \                /  \
//         c     y              c    y                       x
//       /   \    \            / \    \                    /  \
//      b     d    z   =>     b  d    z   =>           a+c+e   y
//     / \   / \             /    \                      /     \
//    a1 a2 e1 e2           a     e         max(a,b)+max(d,e)   z
// it actually is house robber problem for array [a, b, c, d, e],
// then recurse to upper node. it becomes (3)
func robHelper(t *TreeNode) (self int, child int) {
	if t == nil {
		return 0, 0
	}
	l, lc := robHelper(t.Left)
	r, rc := robHelper(t.Right)
	child = max(l, lc) + max(r, rc)
	self = t.Val + lc + rc
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	t1 := &TreeNode{3, nil, nil}
	t1.Left = &TreeNode{2, nil, nil}
	t1.Right = &TreeNode{3, nil, nil}
	t1.Left.Right = &TreeNode{3, nil, nil}
	t1.Right.Right = &TreeNode{1, nil, nil}
	fmt.Println(rob(t1))

	t2 := &TreeNode{2, nil, nil}
	t2.Left = &TreeNode{4, nil, nil}
	t2.Right = &TreeNode{5, nil, nil}
	t2.Left.Left = &TreeNode{2, nil, nil}
	t2.Left.Right = &TreeNode{3, nil, nil}
	t2.Right.Right = &TreeNode{4, nil, nil}
	fmt.Println(rob(t2))

	t3 := &TreeNode{4, nil, nil}
	t3.Left = &TreeNode{1, nil, nil}
	t3.Left.Left = &TreeNode{2, nil, nil}
	t3.Left.Left.Left = &TreeNode{3, nil, nil}
	fmt.Println(rob(t3))
}

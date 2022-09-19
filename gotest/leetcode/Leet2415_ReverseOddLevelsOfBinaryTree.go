package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/

// Given the root of a perfect binary tree, reverse the node values at each odd level of the tree.
// For example, suppose the node values at level 3 are [2,1,3,4,7,11,29,18], then it should become
// [18,29,11,7,4,3,1,2].
// Return the root of the reversed tree.
// A binary tree is perfect if all parent nodes have two children and all leaves are on the same level.
// The level of a node is the number of edges along the path between it and the root node.
// Example 1:
//   Input: root = [2,3,5,8,13,21,34]
//   Output: [2,5,3,8,13,21,34]
//   Explanation:
//     The tree has only one odd level.
//     The nodes at level 1 are 3, 5 respectively, which are reversed and become 5, 3.
// Example 2:
//   Input: root = [7,13,11]
//   Output: [7,11,13]
//   Explanation:
//     The nodes at level 1 are 13, 11, which are reversed and become 11, 13.
// Example 3:
//   Input: root = [0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]
//   Output: [0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]
//   Explanation:
//     The odd levels have non-zero values.
//     The nodes at level 1 were 1, 2, and are 2, 1 after the reversal.
//     The nodes at level 3 were 1, 1, 1, 1, 2, 2, 2, 2, and are 2, 2, 2, 2, 1, 1, 1, 1 after the reversal.
// Constraints:
//   The number of nodes in the tree is in the range [1, 2¹⁴].
//   0 <= Node.val <= 10⁵
//   root is a perfect binary tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	tier := 0
	queue := list.New()
	queue.PushBack(root)
	buf := make([]int, 0, 16)
	for queue.Len() > 0 {
		l := queue.Len()
		if tier%2 == 0 { // even tier
			for i := 0; i < l; i++ {
				cur := queue.Remove(queue.Front()).(*TreeNode)
				if cur.Left != nil {
					queue.PushBack(cur.Left)
				}
				if cur.Right != nil {
					queue.PushBack(cur.Right)
				}
			}
			tier++
			continue
		}
		// odd tier
		buf = buf[:0]
		cur := queue.Front()
		for i := 0; i < l; i++ {
			buf = append(buf, cur.Value.(*TreeNode).Val)
			cur = cur.Next()
		}
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*TreeNode)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
			cur.Val = buf[l-i-1]
		}
		tier++
	}
	return root
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 8}
	root.Left.Right = &TreeNode{Val: 13}
	root.Right.Left = &TreeNode{Val: 21}
	root.Right.Right = &TreeNode{Val: 34}
	root.Left.Left.Left = &TreeNode{Val: 1}
	root.Left.Left.Right = &TreeNode{Val: 2}
	root.Left.Right.Left = &TreeNode{Val: 3}
	root.Left.Right.Right = &TreeNode{Val: 4}
	root.Right.Left.Left = &TreeNode{Val: 5}
	root.Right.Left.Right = &TreeNode{Val: 6}
	root.Right.Right.Left = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 8}

	reverseOddLevels(root)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		l := queue.Len()
		for i := 0; i < l; i++ {
			cur := queue.Remove(queue.Front()).(*TreeNode)
			fmt.Println(cur.Val)
			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}
			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}
		}
	}
}

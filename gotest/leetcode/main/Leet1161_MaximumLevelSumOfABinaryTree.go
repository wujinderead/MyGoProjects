package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/

// Given the root of a binary tree, the level of its root is 1,
// the level of its children is 2, and so on.
// Return the smallest level X such that
// the sum of all the values of nodes at level X is maximal.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// use bfs to find the maximal level sum
func maxLevelSum(root *TreeNode) int {
	type node struct {
		n *TreeNode
		l int
	}
	minLevel := 1
	maxSum := root.Val
	queue := list.New()
	queue.PushBack(node{root, 1})
	levelSum := 0
	for queue.Len() > 0 {
		cur := queue.Remove(queue.Front()).(node)
		l := cur.l
		levelSum += cur.n.Val
		if cur.n.Left != nil {
			queue.PushBack(node{cur.n.Left, l + 1})
		}
		if cur.n.Right != nil {
			queue.PushBack(node{cur.n.Right, l + 1})
		}
		if (queue.Len() > 0 && queue.Front().Value.(node).l != l) || queue.Len() == 0 { // last node of current level
			if levelSum > maxSum {
				maxSum = levelSum
				minLevel = l
			}
			levelSum = 0
		}
	}
	return minLevel
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: -8}
	root.Right.Right = &TreeNode{Val: 8}
	fmt.Println(maxLevelSum(root))
}

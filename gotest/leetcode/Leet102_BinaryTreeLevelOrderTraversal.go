package main

import "fmt"

// https://leetcode.com/problems/binary-tree-level-order-traversal/

// Given a binary tree, return the level order traversal of its nodes' values.
// (ie, from left to right, level by level).
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its level order traversal as:
// [
//   [3],
//   [9,20],
//   [15,7]
// ]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	re := make([][]int, 0)
	if root == nil {
		return re
	}
	queue := &linkQueue{}
	queue.pushback(root, 0)
	curlevel := 0
	curlevelelem := make([]int, 0)
	for !queue.empty() {
		node, level := queue.popfront()
		if level == curlevel {
			curlevelelem = append(curlevelelem, node.Val)
		} else {
			curlevel++
			re = append(re, curlevelelem)
			curlevelelem = []int{node.Val}
		}
		if node.Left != nil {
			queue.pushback(node.Left, level+1)
		}
		if node.Right != nil {
			queue.pushback(node.Right, level+1)
		}
	}
	re = append(re, curlevelelem)
	return re
}

type linkQueue struct {
	head, tail *item
}

type item struct {
	node  *TreeNode
	level int
	next  *item
}

func (q *linkQueue) pushback(node *TreeNode, level int) {
	it := &item{node: node, level: level, next: nil}
	if q.tail == nil { // empty
		q.tail = it
		q.head = it
		return
	}
	q.tail.next = it
	q.tail = it
}

func (q *linkQueue) popfront() (*TreeNode, int) {
	if q.head == nil {
		return nil, -1
	}
	re := q.head
	if q.head == q.tail { // last one item
		q.head = nil
		q.tail = nil
		return re.node, re.level
	}
	q.head = q.head.next
	return re.node, re.level
}

func (q *linkQueue) empty() bool {
	return q.head == nil
}

func main() {
	r := &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Left.Left = &TreeNode{Val: 4}
	r.Left.Right = &TreeNode{Val: 5}
	r.Right.Left = &TreeNode{Val: 6}
	r.Right.Right = &TreeNode{Val: 7}
	fmt.Println(levelOrder(r))

	r = &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Left.Right = &TreeNode{Val: 5}
	r.Right.Left = &TreeNode{Val: 6}
	r.Right.Left.Right = &TreeNode{Val: 7}
	fmt.Println(levelOrder(r))
}

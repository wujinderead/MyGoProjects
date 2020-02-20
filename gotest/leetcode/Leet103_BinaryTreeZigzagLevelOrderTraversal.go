package main

import "fmt"

// https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/

// Given a binary tree, return the zigzag level order traversal of its nodes' values.
// (ie, from left to right, then right to left for the next level and alternate between).
// For example:
// Given binary tree [3,9,20,null,null,15,7],
//
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its zigzag level order traversal as:
//  [
//    [3],
//    [20,9],
//    [15,7]
//  ]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	for i := range re {
		if i%2 == 1 {
			reverse(re[i])
		}
	}
	return re
}

func reverse(arr []int) {
	i, j := 0, len(arr)-1
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
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
	fmt.Println(zigzagLevelOrder(r))

	r = &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Left.Right = &TreeNode{Val: 5}
	r.Right.Left = &TreeNode{Val: 6}
	r.Right.Left.Right = &TreeNode{Val: 7}
	fmt.Println(zigzagLevelOrder(r))
}

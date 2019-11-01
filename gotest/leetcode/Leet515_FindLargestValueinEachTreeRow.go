package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/find-largest-value-in-each-tree-row/

// You need to find the largest value in each row of a binary tree.
// Example:
//   Input:
//
//          1
//         / \
//        3   2
//       / \   \
//      5   3   9
//
//   Output: [1, 3, 9]

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	re := make([]int, 0)
	if root == nil {
		return re
	}
	queue := &linkQueue{}
	queue.pushback(root, 0)
	curlevel := 0
	curmax := math.MinInt64
	for !queue.empty() {
		node, level := queue.popfront()
		if level == curlevel {
			if node.Val > curmax {
				curmax = node.Val
			}
		} else {
			curlevel++
			re = append(re, curmax)
			curmax = node.Val
		}
		if node.Left != nil {
			queue.pushback(node.Left, level+1)
		}
		if node.Right != nil {
			queue.pushback(node.Right, level+1)
		}
	}
	re = append(re, curmax)
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
	r := &TreeNode{Val: -12}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Left.Left = &TreeNode{Val: 4}
	r.Left.Right = &TreeNode{Val: 5}
	r.Right.Left = &TreeNode{Val: 6}
	r.Right.Right = &TreeNode{Val: 7}
	fmt.Println(largestValues(r))

	r = &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Left.Right = &TreeNode{Val: 5}
	r.Right.Left = &TreeNode{Val: 6}
	r.Right.Left.Right = &TreeNode{Val: 7}
	fmt.Println(largestValues(r))
}

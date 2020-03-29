package main

import "fmt"

// https://leetcode.com/problems/populating-next-right-pointers-in-each-node

// You are given a perfect binary tree where all leaves are on the same level, and
// every parent has two children. The binary tree has the following definition:
// struct Node {
//   int val;
//   Node *left;
//   Node *right;
//   Node *next;
// }
// Populate each next pointer to point to its next right node. If there is no next
// right node, the next pointer should be set to NULL.
// Initially, all next pointers are set to NULL.
// Follow up:
//   You may only use constant extra space.
//   Recursive approach is fine, you may assume implicit stack space does not count
//   as extra space for this problem.
// Example 1:
//   Input: root = [1,2,3,4,5,6,7]
//   Output: [1,#,2,3,#,4,5,6,7,#]
//   Explanation: Given the above perfect binary tree (Figure A), your function should
//   populate each next pointer to point to its next right node, just like in Figure B.
//   The serialized output is in level order as connected by the next pointers,
//   with '#' signifying the end of each level.
//          1  ->
//       /    \
//     2   ->   3  ->
//    /  \     / \
//   4 -> 5-> 6 ->7 ->
//
// Constraints:
//   The number of nodes in the given tree is less than 4096.
//   -1000 <= node.val <= 1000

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	root.Next = nil
	connectHelper(root)
	return root
}

func connectHelper(cur *Node) {
	l, r := cur.Left, cur.Right
	for l != nil {
		l.Next = r
		l = l.Right
		r = r.Left
	}
	if cur.Left != nil {
		connectHelper(cur.Left)
		connectHelper(cur.Right)
	}
}

func main() {
	root := &Node{Val: 1}
	root.Left = &Node{Val: 2}
	root.Right = &Node{Val: 3}
	root.Left.Left = &Node{Val: 4}
	root.Left.Left.Left = &Node{Val: 8}
	root.Left.Left.Right = &Node{Val: 9}
	root.Left.Right = &Node{Val: 5}
	root.Left.Right.Left = &Node{Val: 10}
	root.Left.Right.Right = &Node{Val: 11}
	root.Right.Left = &Node{Val: 6}
	root.Right.Left.Left = &Node{Val: 12}
	root.Right.Left.Right = &Node{Val: 13}
	root.Right.Right = &Node{Val: 7}
	root.Right.Right.Left = &Node{Val: 14}
	root.Right.Right.Right = &Node{Val: 15}
	connect(root)
	r := root
	for r != nil {
		t := r
		for t != nil {
			fmt.Print(t.Val, " ")
			t = t.Next
		}
		fmt.Println()
		r = r.Left
	}
}

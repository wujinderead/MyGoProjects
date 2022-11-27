package main

import "fmt"

// https://leetcode.com/problems/remove-nodes-from-linked-list/

// You are given the head of a linked list.
// Remove every node which has a node with a strictly greater value anywhere to the right side of it.
// Return the head of the modified linked list.
// Example 1:
//   Input: head = [5,2,13,3,8]
//   Output: [13,8]
//   Explanation: The nodes that should be removed are 5, 2 and 3.
//     - Node 13 is to the right of node 5.
//     - Node 13 is to the right of node 2.
//     - Node 8 is to the right of node 3.
// Example 2:
//   Input: head = [1,1,1,1]
//   Output: [1,1,1,1]
//   Explanation: Every node has value 1, so no nodes are removed.
// Constraints:
//   The number of the nodes in the given list is in the range [1, 10⁵].
//   1 <= Node.val <= 10⁵

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	var stack []int
	for head != nil {
		for len(stack) > 0 && stack[len(stack)-1] < head.Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, head.Val)
		head = head.Next
	}
	newHead := &ListNode{Val: stack[0]}
	head = newHead
	for i := 1; i < len(stack); i++ {
		head.Next = &ListNode{Val: stack[i]}
		head = head.Next
	}
	return newHead
}

func main() {
	f := func(h *ListNode) {
		for h != nil {
			fmt.Print(h.Val, "->")
			h = h.Next
		}
		fmt.Println()
	}
	{
		h := &ListNode{Val: 1}
		h.Next = &ListNode{Val: 1}
		h.Next.Next = &ListNode{Val: 1}
		h.Next.Next.Next = &ListNode{Val: 1}
		f(h)
		hh := removeNodes(h)
		f(hh)
	}
	{
		h := &ListNode{Val: 5}
		h.Next = &ListNode{Val: 2}
		h.Next.Next = &ListNode{Val: 13}
		h.Next.Next.Next = &ListNode{Val: 3}
		h.Next.Next.Next.Next = &ListNode{Val: 8}
		f(h)
		hh := removeNodes(h)
		f(hh)
	}
}

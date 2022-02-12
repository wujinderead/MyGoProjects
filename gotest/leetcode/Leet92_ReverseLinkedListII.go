package main

import "fmt"

// https://leetcode.com/problems/reverse-linked-list-ii/

// Given the head of a singly linked list and two integers left and right where left <= right,
// reverse the nodes of the list from position left to position right, and return the reversed list.
// Example 1:
//   Input: head = [1,2,3,4,5], left = 2, right = 4
//   Output: [1,4,3,2,5]
// Example 2:
//   Input: head = [5], left = 1, right = 1
//   Output: [5]
// Constraints:
//   The number of nodes in the list is n.
//   1 <= n <= 500
//   -500 <= Node.val <= 500
//   1 <= left <= right <= n

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	prev, cur := (*ListNode)(nil), head
	for left-1 > 0 {
		prev = cur
		cur = cur.Next
		left--
		right--
	}
	l, lPrev := cur, prev // find left and its prev
	next := cur.Next
	for right-1 > 0 { // move to right, change direction along
		nn := next.Next
		next.Next = cur
		cur = next
		next = nn
		right--
	}
	r, rNext := cur, next // find right and its next
	// .. lPrev l ... r rNext ..
	// .. lPrev r ... l rNext ..
	l.Next = rNext // it's definite
	if lPrev != nil {
		lPrev.Next = r // it's conditional, depends on whether l is the first
		return head
	}
	return r
}

func main() {
	printList := func(l *ListNode) {
		for l != nil {
			fmt.Print(l.Val, " -> ")
			l = l.Next
		}
		fmt.Println()
	}
	{
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		head.Next.Next.Next = &ListNode{Val: 4}
		head.Next.Next.Next.Next = &ListNode{Val: 5}
		printList(head)
		head = reverseBetween(head, 2, 4)
		printList(head)
	}
	{
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		printList(head)
		head = reverseBetween(head, 1, 2)
		printList(head)
	}
	{
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		printList(head)
		head = reverseBetween(head, 1, 2)
		printList(head)
	}
	{
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		printList(head)
		head = reverseBetween(head, 2, 3)
		printList(head)
	}
	{
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 3}
		printList(head)
		head = reverseBetween(head, 1, 3)
		printList(head)
	}
}

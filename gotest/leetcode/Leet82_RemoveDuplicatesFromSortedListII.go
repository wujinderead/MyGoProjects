package main

import "fmt"

// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/

// Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
// leaving only distinct numbers from the original list. Return the linked list sorted as well.
// Example 1:
//   Input: head = [1,2,3,3,4,4,5]
//   Output: [1,2,5]
// Example 2:
//   Input: head = [1,1,1,2,3]
//   Output: [2,3]
// Constraints:
//   The number of nodes in the list is in the range [0, 300].
//   -100 <= Node.val <= 100
//   The list is guaranteed to be sorted in ascending order.

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := &ListNode{Next: head}
	prev, cur, next := dummy, head, head.Next
	dup := false
	for next != nil {
		if next.Val == cur.Val {
			prev.Next = next
			dup = true
		} else {
			if dup {
				prev.Next = next
			} else {
				prev = cur
			}
			dup = false
		}
		cur = next
		next = next.Next
	}
	if dup {
		prev.Next = nil
	}
	return dummy.Next
}

func deleteDuplicates_doubleLoop(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev, cur := dummy, head
	for cur != nil {
		next := cur.Next
		if next != nil && next.Val == cur.Val {
			for next != nil && cur.Val == next.Val {
				cur = next
				next = next.Next
			}
			prev.Next = next
		} else {
			prev = cur
		}
		cur = next
	}
	return dummy.Next
}

func main() {
	printList := func(l *ListNode) {
		for l != nil {
			fmt.Print(l.Val, " -> ")
			l = l.Next
		}
		fmt.Println()
	}
	makeList := func(arr []int) *ListNode {
		head := &ListNode{Val: arr[0]}
		cur := head
		for i := 1; i < len(arr); i++ {
			cur.Next = &ListNode{Val: arr[i]}
			cur = cur.Next
		}
		return head
	}
	for _, v := range [][]int{
		{1, 1, 1, 2, 5},
		{1, 2, 3, 3, 4, 4, 5},
		{1, 2, 2, 3, 4, 4, 5},
		{1, 2, 2, 2},
		{1, 2, 2},
		{1, 1, 2, 2},
		{1, 1, 2, 2, 3},
		{1, 2, 2, 3, 3, 4},
		{1, 1, 1},
		{1, 1},
		{1},
	} {
		{
			head := makeList(v)
			printList(head)
			head = deleteDuplicates(head)
			printList(head)
		}
		{
			head := makeList(v)
			head = deleteDuplicates_doubleLoop(head)
			printList(head)
		}
	}
}

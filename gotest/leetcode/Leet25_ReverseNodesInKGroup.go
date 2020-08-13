package main

import "fmt"

// https://leetcode.com/problems/reverse-nodes-in-k-group/submissions/

// Given a linked list, reverse the nodes of a linked list k at a time and return its modified list. 
// k is a positive integer and is less than or equal to the length of the linked list. If the number 
// of nodes is not a multiple of k then left-out nodes in the end should remain as it is. 
// Example: 
//   Given this linked list: 1->2->3->4->5 
//   For k = 2, you should return: 2->1->4->3->5 
//   For k = 3, you should return: 3->2->1->4->5 
// Note: 
//   Only constant extra memory is allowed. 
//   You may not alter the values in the list's nodes, only nodes itself may be changed. 

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
    if k<=1 {
		return head
	}
	ret, s, e, pe := head, head, head, (*ListNode)(nil)
	for {
		for i:=0; i<k-1 && e!= nil; i++ {
			e = e.Next
		}
		if e==nil {
			break
		}
		// previous: -> pe -> s -> x -> e -> en, reverse s to e
		// become:   -> pe -> e -> x -> s -> en 
		if pe != nil {
			pe.Next = e
		}
		pe = s
		en := e.Next
		x := s.Next
		s.Next = en
		for x != en {
			xn := x.Next
			x.Next = s
			s = x
			x = xn
		}
		if ret == head {
			ret = e
		}
		s, e = en, en
	}
	return ret
}

func main() {
	l := &ListNode{Val: 1}
	l.Next = &ListNode{Val: 2}
	l.Next.Next = &ListNode{Val: 3}
	l.Next.Next.Next = &ListNode{Val: 4}
	l.Next.Next.Next.Next = &ListNode{Val: 5}
	for i:=1; i<=6; i++ {
		l = reverseKGroup(l, i)
		print(l)
	}
	l.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	for i:=1; i<=7; i++ {
		l = reverseKGroup(l, i)
		print(l)
	}
}

func print(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val, " -> ")
		l = l.Next
	}
	fmt.Println()
}
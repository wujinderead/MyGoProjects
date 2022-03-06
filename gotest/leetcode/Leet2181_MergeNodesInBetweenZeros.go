package main

import "fmt"

// https://leetcode.com/problems/merge-nodes-in-between-zeros/

// You are given the head of a linked list, which contains a series of integers separated by 0's.
// The beginning and end of the linked list will have Node.val == 0.
// For every two consecutive 0's, merge all the nodes lying in between them into a single node
// whose value is the sum of all the merged nodes. The modified list should not contain any 0's.
// Return the head of the modified linked list.
// Example 1:
//   Input: head = [0,3,1,0,4,5,2,0]
//   Output: [4,11]
//   Explanation:
//     The above figure represents the given linked list. The modified list contains
//     - The sum of the nodes marked in green: 3 + 1 = 4.
//     - The sum of the nodes marked in red: 4 + 5 + 2 = 11.
// Example 2:
//   Input: head = [0,1,0,3,0,2,2,0]
//   Output: [1,3,4]
//   Explanation:
//     The above figure represents the given linked list. The modified list contains
//     - The sum of the nodes marked in green: 1 = 1.
//     - The sum of the nodes marked in red: 3 = 3.
//     - The sum of the nodes marked in yellow: 2 + 2 = 4.
// Constraints:
//   The number of nodes in the list is in the range [3, 2 * 10⁵].
//   0 <= Node.val <= 1000
//   There are no two consecutive nodes with Node.val == 0.
//   The beginning and end of the linked list have Node.val == 0.

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	prev := head
	cur := head.Next
	sum := 0
	for cur != nil {
		if cur.Val != 0 {
			sum += cur.Val
			cur = cur.Next
		} else {
			cur.Val = sum
			sum = 0
			prev.Next = cur
			prev = cur
			cur = cur.Next
		}
	}
	return head.Next
}

func sliceToLinkedList(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

func print(h *ListNode) {
	for h != nil {
		fmt.Print(h.Val, "->")
		h = h.Next
	}
	fmt.Println()
}

func main() {
	for _, v := range []*ListNode{
		sliceToLinkedList([]int{0, 1, 0}),
		sliceToLinkedList([]int{0, 3, 1, 0, 4, 5, 2, 0}),
		sliceToLinkedList([]int{0, 2, 2, 0}),
	} {
		print(mergeNodes(v))
	}
}

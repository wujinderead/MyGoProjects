package main

import "fmt"

// https://leetcode.com/problems/merge-in-between-linked-lists/

// You are given two linked lists: list1 and list2 of sizes n and m respectively.
// Remove list1's nodes from the ath node to the bth node, and put list2 in their place.
// The blue edges and nodes in the following figure incidate the result:
// Build the result list and return its head.
// Example 1:
//   Input: list1 = [0,1,2,3,4,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
//   Output: [0,1,2,1000000,1000001,1000002,5]
//   Explanation: We remove the nodes 3 and 4 and put the entire list2 in their place.
//     The blue edges and nodes in the above figure indicate the result.
// Example 2:
//   Input: list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
//   Output: [0,1,1000000,1000001,1000002,1000003,1000004,6]
//   Explanation: The blue edges and nodes in the above figure indicate the result.
// Constraints:
//   3 <= list1.length <= 104
//   1 <= a <= b < list1.length - 1
//   1 <= list2.length <= 104

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	aprev := list1
	for i := 0; i < a-1; i++ {
		aprev = aprev.Next
	}
	bnext := aprev
	for i := a - 1; i <= b; i++ {
		bnext = bnext.Next
	}
	list2tail := list2
	for list2tail.Next != nil {
		list2tail = list2tail.Next
	}
	aprev.Next = list2
	list2tail.Next = bnext
	return list1
}

func main() {
	for _, v := range []struct {
		l1      []int
		a, b    int
		l2, ans []int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 3, 4, []int{1000000, 1000001, 1000002},
			[]int{0, 1, 2, 1000000, 1000001, 1000002, 5}},
		{[]int{0, 1, 2, 3, 4, 5, 6}, 2, 5, []int{1000000, 1000001, 1000002, 1000003, 1000004},
			[]int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}},
	} {
		list1, list2 := makeList(v.l1), makeList(v.l2)
		ans := mergeInBetween(list1, v.a, v.b, list2)
		printList(ans)
		printList(makeList(v.ans))
	}
}

func makeList(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &ListNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println()
}

package main

import "fmt"

// https://leetcode.com/problems/flatten-a-multilevel-doubly-linked-list

// You are given a doubly linked list which in addition to the next and previous pointers,
// it could have a child pointer, which may or may not point to a separate doubly linked list.
// These child lists may have one or more children of their own, and so on, to produce a multilevel
// data structure, as shown in the example below.
// Flatten the list so that all the nodes appear in a single-level, doubly linked list. You are given
// the head of the first level of the list.
// Example 1:
//   Input: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
//   Output: [1,2,3,7,8,11,12,9,10,4,5,6]
//   Explanation:
//     The multilevel linked list in the input is as follows:
//        1---2---3---4---5---6--NULL
//                |
//                7---8---9---10--NULL
//                    |
//                    11--12--NULL
//     After flattening the multilevel linked list it becomes:
//        1--2--3--7--8--11--12--9--10--4--5--6
// Example 2:
//   Input: head = [1,2,null,3]
//   Output: [1,3,2]
//   Explanation:
//     The input multilevel linked list is as follows:
//       1---2---NULL
//       |
//       3---NULL
// Example 3:
//   Input: head = []
//   Output: []
// Constraints:
//   Number of Nodes will not exceed 1000.
//   1 <= Node.val <= 10^5

type Node struct {
    Val int
    Prev *Node
    Next *Node
    Child *Node
}

func flatten(root *Node) *Node {
	helper(root)
    return root
}

func helper(r *Node) *Node {
	n := r
	ret := r
	for n != nil {    // traverse list
		if n.Child != nil {       // if has child
			nc := n.Child
			nn := n.Next
			last := helper(nc)    // proceed to child, last is the last node for flattened child 
			n.Next = nc
			nc.Prev = n
			last.Next = nn
			if nn != nil {        // former state: n <-> nn, current state n <--> nc ... last <--> nn
				nn.Prev = last
			}
			n.Child = nil
			ret = last    // set ret to last 
			n = nn
		} else {
			ret = n       // set ret to last non-nil node, which is 'n' for now
			n = n.Next
		}
	}
	return ret
}

func main() {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2, Prev: head}
	head.Next.Next = &Node{Val: 3, Prev: head.Next}
	head.Next.Next.Next = &Node{Val: 4, Prev: head.Next.Next}
	head.Next.Next.Next.Next = &Node{Val: 5, Prev: head.Next.Next.Next}
	head.Next.Next.Next.Next.Next = &Node{Val: 6, Prev: head.Next.Next.Next.Next}
	head.Next.Next.Child = &Node{Val: 7}
	head.Next.Next.Child.Next = &Node{Val: 8, Prev: head.Next.Next.Child}
	head.Next.Next.Child.Next.Next = &Node{Val: 9, Prev: head.Next.Next.Child.Next}
	head.Next.Next.Child.Next.Next.Next = &Node{Val: 10, Prev: head.Next.Next.Child.Next.Next}
	head.Next.Next.Child.Next.Child = &Node{Val: 11}
	head.Next.Next.Child.Next.Child.Next = &Node{Val: 12, Prev: head.Next.Next.Child.Next.Child}
	nn := flatten(head)
	for nn != nil {
		fmt.Println(nn.Val, nn.Prev)
		nn = nn.Next
	}
	fmt.Println()

	head = &Node{Val: 1}
	head.Next = &Node{Val: 2, Prev: head}
	head.Child = &Node{Val: 3}
	nn = flatten(head)
	for nn != nil {
		fmt.Println(nn.Val, nn.Prev)
		nn = nn.Next
	}
	fmt.Println()

	head = &Node{Val: 1}
	head.Next = &Node{Val: 2, Prev: head}
	head.Child = &Node{Val: 3}
	head.Next.Child = &Node{Val: 4}
	nn = flatten(head)
	for nn != nil {
		fmt.Println(nn.Val, nn.Prev)
		nn = nn.Next
	}
}

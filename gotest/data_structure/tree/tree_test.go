package tree

import (
	"fmt"
	"testing"
)

type Node struct {
	val      int
	next     *Node
	children *Node
}

//                1
//       /      /      \
//    2        3          4
//   /\     / / \ \     / \ \
//  5  6   7 8  9 10   11 12 13
//    /      \    / \    / \
//   14       15 19 18  16 17
//                  /\
//                 20 21
func makeTree() *Node {
	root := &Node{val: 1}
	root.children = &Node{val: 2}
	root.children.next = &Node{val: 3}
	root.children.next.next = &Node{val: 4}
	root.children.children = &Node{val: 5}
	root.children.children.next = &Node{val: 6}
	root.children.next.children = &Node{val: 7}
	root.children.next.children.next = &Node{val: 8}
	root.children.next.children.next.next = &Node{val: 9}
	root.children.next.children.next.next.next = &Node{val: 10}
	root.children.next.next.children = &Node{val: 11}
	root.children.next.next.children.next = &Node{val: 12}
	root.children.next.next.children.next.next = &Node{val: 13}
	root.children.children.next.children = &Node{val: 14}
	root.children.next.children.next.children = &Node{val: 15}
	root.children.next.next.children.next.children = &Node{val: 16}
	root.children.next.next.children.next.children.next = &Node{val: 17}
	root.children.next.children.next.next.next.children = &Node{val: 19}
	root.children.next.children.next.next.next.children.next = &Node{val: 18}
	root.children.next.children.next.next.next.children.next.children = &Node{val: 20}
	root.children.next.children.next.next.next.children.next.children.next = &Node{val: 21}
	return root
}

func TestTreePreOrderTraverse(t *testing.T) {
	root := makeTree()
	stack := []*Node{root}
	tmpstack := make([]*Node, 0)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Println(cur.val)
		// push children to stack in reverse order
		for child := cur.children; child != nil; child = child.next {
			tmpstack = append(tmpstack, child)
		}
		for len(tmpstack) > 0 {
			stack = append(stack, tmpstack[len(tmpstack)-1])
			tmpstack = tmpstack[:len(tmpstack)-1]
		}
	}
}

func TestTreePostOrderTraverse(t *testing.T) {
	root := makeTree()
	stack := make([]*Node, 0)
	cur := root
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.children
		} else {
			cur = stack[len(stack)-1]
			// we can see current stack is the path from root to current node
			for i := range stack {
				fmt.Print(stack[i].val, " ")
			}
			fmt.Println(",", cur.val)
			stack = stack[:len(stack)-1]
			cur = cur.next
		}
	}
}

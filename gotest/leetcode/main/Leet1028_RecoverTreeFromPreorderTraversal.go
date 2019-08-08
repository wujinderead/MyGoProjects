package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverFromPreorder(s string) *TreeNode {
	i := 0
	root := (*TreeNode)(nil)
	stack := list.New()
	for i < len(s) {
		level := 0
		num := 0
		for s[i] == '-' { // get level
			level++
			i++
		}
		for i < len(s) && s[i] != '-' { // get number
			num *= 10
			num += int(s[i]) - '0'
			i++
		}
		node := &TreeNode{num, nil, nil} // make new node
		for level < stack.Len() {
			stack.Remove(stack.Back())
		}
		if root == nil {
			root = node
		} else if stack.Back().Value.(*TreeNode).Left == nil {
			stack.Back().Value.(*TreeNode).Left = node
		} else {
			stack.Back().Value.(*TreeNode).Right = node
		}
		stack.PushBack(node)
	}
	return root
}

func printTree(root *TreeNode) {
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		cur := queue.Front()
		fmt.Print(cur.Value)
		fmt.Printf(" %p\n", cur.Value.(*TreeNode))
		queue.Remove(cur)
		if cur.Value.(*TreeNode) != nil {
			queue.PushBack(cur.Value.(*TreeNode).Left)
			queue.PushBack(cur.Value.(*TreeNode).Right)
		}
	}
}

func main() {
	printTree(recoverFromPreorder("123"))
	printTree(recoverFromPreorder("1-4--5---6--7-8--17--12---13"))
}

package main

import "fmt"

// https://leetcode.com/problems/n-ary-tree-preorder-traversal/

// Given an n-ary tree, return the preorder traversal of its nodes' values.
// Nary-Tree input serialization is represented in their level order traversal, 
// each group of children is separated by the null value (See examples).
// Follow up:
//   Recursive solution is trivial, could you do it iteratively?
// Example 1: 
//            1
//         /  |  \
//        3   2   4
//       / \
//      5  6
//   Input: root = [1,null,3,2,4,null,5,6]
//   Output: [1,3,5,6,2,4]
// Example 2:
//   Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
//   Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]
// Constraints:
//   The height of the n-ary tree is less than or equal to 1000
//   The total number of nodes is between [0, 10^4]

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
    buf := make([]int, 0, 10)
    if root==nil {
    	return buf
	}
    stack := make([]*Node, 0, 10)
    stack = append(stack, root)
    for len(stack)>0 {
    	cur := stack[len(stack)-1]
    	stack = stack[:len(stack)-1]
    	buf = append(buf, cur.Val)
    	for i:=len(cur.Children)-1; i>=0; i-- {
    		stack = append(stack, cur.Children[i])  // push children to stack in reverse order
		}
	}
	return buf
}

func main() {
	root := &Node{Val: 1}
	root.Children = []*Node{{Val: 3}, {Val: 2}, {Val: 4}}
	root.Children[0].Children = []*Node{{Val: 5}, {Val: 6}}
	fmt.Println(preorder(root))
	root = &Node{Val: 3}
	fmt.Println(preorder(root))
	fmt.Println(preorder(nil))
}
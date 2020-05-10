package main

import "fmt"

// https://leetcode.com/problems/deepest-leaves-sum/

// Given a binary tree, return the sum of values of its deepest leaves.
// Example 1: 
//           1
//          / \
//         2  3
//        / \  \
//       4  5  6
//      /       \
//     7        8
//   Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
//   Output: 15
// Constraints:
//   The number of nodes in the tree is between 1 and 10^4.
//   The value of nodes is between 1 and 100.

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// post-order traverse the tree, when we visit a node, the stack is the path back to root
func deepestLeavesSum(root *TreeNode) int {
    cur := root
    stack := make([]*TreeNode, 0, 50)
    max := 0
    lastvisit := (*TreeNode)(nil)

    // get deepest level
    for cur != nil || len(stack)>0 {
    	if cur != nil {
    		stack = append(stack, cur)
    		cur = cur.Left
		} else {
			peek := stack[len(stack)-1]
			if peek.Right != nil && lastvisit != peek.Right {
				cur = peek.Right
			} else {
				if len(stack)>max {    // visit peek
					max = len(stack)
				}
				lastvisit = peek
				stack = stack[:len(stack)-1]
			}
		}
	}

	// get deepest level sum
	cur = root
	all := 0
	lastvisit = nil
	for cur != nil || len(stack)>0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			peek := stack[len(stack)-1]
			if peek.Right != nil && lastvisit != peek.Right {
				cur = peek.Right
			} else {
				if len(stack)==max {    // visit peek
					all += peek.Val
				}
				lastvisit = peek
				stack = stack[:len(stack)-1]
			}
		}
	}
	return all
}

func main() {
	r := &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Left.Left = &TreeNode{Val: 4}
	r.Left.Right = &TreeNode{Val: 5}
	r.Right.Right = &TreeNode{Val: 6}
	r.Left.Left.Left = &TreeNode{Val: 7}
	r.Right.Right.Right = &TreeNode{Val: 8}
	fmt.Println(deepestLeavesSum(r))
	r = &TreeNode{Val: 1}
	fmt.Println(deepestLeavesSum(r))
	r = &TreeNode{Val: 1}
	r.Right = &TreeNode{Val: 2}
	fmt.Println(deepestLeavesSum(r))
}
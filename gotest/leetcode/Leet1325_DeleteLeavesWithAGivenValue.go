package main

import (
	"fmt"
)

// https://leetcode.com/problems/delete-leaves-with-a-given-value

// Given a binary tree root and an integer target, delete all the leaf nodes with
// value target. Note that once you delete a leaf node with value target, if it's parent node
// becomes a leaf node and has the value target, it should also be deleted (you need to continue
// doing that until you can't).
// Example 1:
//   Input: root = [1,2,3,2,null,2,4], target = 2
//          1                  1
//         / \                  \
//        2   3      ->          3
//       /   / \                  \
//      2   2   4                  4
//   Output: [1,null,3,null,4]
//   Explanation: Leaf nodes in green with value (target = 2) are removed (Picture in left).
//     After removing, new nodes become leaf nodes with value (target = 2) (Picture in center).
// Example 2: 
//         1                 1
//        / \               /
//       3  3        ->    3
//      / \                 \
//     3  2                  2
//   Input: root = [1,3,3,3,2], target = 3
//   Output: [1,3,null,null,2]
// Example 3: 
//           1
//          /
//         2
//        /       ->    1
//       2
//      /
//     2
//   Input: root = [1,2,null,2,null,2], target = 2
//   Output: [1]
//   Explanation: Leaf nodes in green with value (target = 2) are removed at each step.
// Example 4:
//   Input: root = [1,1,1], target = 1
//   Output: []
// Example 5:
//   Input: root = [1,2,3], target = 1
//   Output: [1,2,3]
// Constraints:
//   1 <= target <= 1000
//   Each tree has at most 3000 nodes.
//   Each node's value is between [1, 1000].

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    // post-order traverse the tree, visit children before parent, delete leaf by parent
    cur := root
    stack := make([]*TreeNode, 0, 10)
    lastVisited := (*TreeNode)(nil)
    for cur != nil || len(stack)>0 {
    	if cur != nil {
    		stack = append(stack, cur)
    		cur = cur.Left
		} else {
			peek := stack[len(stack)-1]
			if peek.Right != nil && peek.Right != lastVisited {
				cur = peek.Right
			} else {
				// visit peek
				if peek.Left != nil && peek.Left.Val==target && peek.Left.Left==nil && peek.Left.Right==nil {
					peek.Left = nil
				}
				if peek.Right != nil && peek.Right.Val==target && peek.Right.Left==nil && peek.Right.Right==nil {
					peek.Right = nil
				}
				lastVisited = peek
				stack = stack[:len(stack)-1]
			}
		}
	}
	if root.Val==target && root.Left==nil && root.Right==nil {
		return nil
	}
	return root
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 4}
	r := removeLeafNodes(root, 2)
	fmt.Println(r)
	fmt.Println(r.Right)
	fmt.Println(r.Right.Right)

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	r = removeLeafNodes(root, 3)
	fmt.Println(r)
	fmt.Println(r.Left)
	fmt.Println(r.Left.Right)

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Left.Left = &TreeNode{Val: 2}
	fmt.Println(removeLeafNodes(root, 2))

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 1}
	fmt.Println(removeLeafNodes(root, 1))

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	r = removeLeafNodes(root, 1)
	fmt.Println(r)
	fmt.Println(r.Left)
	fmt.Println(r.Right)
}
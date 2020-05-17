package main

import (
	"fmt"
)

// https://leetcode.com/problems/count-good-nodes-in-binary-tree/

// Given a binary tree root, a node X in the tree is named good if in the path from
// root to X there are no nodes with a value greater than X.
// Return the number of good nodes in the binary tree.
// Example 1:
//   Input: root = [3,1,4,3,null,1,5]
//   Output: 4
//   Explanation: Nodes in blue are good.
//           (3)
//          /  \
//         1   (4)
//        /    / \
//      (3)   1  (5)
//     Root Node (3) is always a good node.
//     Node 4 -> (3,4) is the maximum value in the path starting from the root.
//     Node 5 -> (3,4,5) is the maximum value in the path
//     Node 3 -> (3,1,3) is the maximum value in the path.
// Example 2:
//   Input: root = [3,3,null,4,2]
//   Output: 3
//   Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
//          (3)
//          /
//        (3)
//        / \
//      (4)  2
// Example 3:
//   Input: root = [1]
//   Output: 1
//   Explanation: Root is considered as good.
// Constraints:
//   The number of nodes in the binary tree is in the range [1, 10^5].
//   Each node's value is between [-10^4, 10^4].

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
    // use post-order traverse to get a path from root to leaf
    stack := make([]*TreeNode, 0, 10)
    maxstack := make([]int, 1, 10)
    cur := root
    maxstack[0] = -0x7ffffff
    lastVisited := (*TreeNode)(nil)
    count := 0
    for cur != nil || len(stack)>0 {
    	if cur != nil {
			stack = append(stack, cur)
			maxpeek := maxstack[len(maxstack)-1]
			if cur.Val>=maxpeek {
				maxstack = append(maxstack, cur.Val)
			} else {
				maxstack = append(maxstack, maxpeek)
			}
			cur = cur.Left
		} else {
			peek := stack[len(stack)-1]
			maxpeek := maxstack[len(maxstack)-1]
			if peek.Right != nil && lastVisited != peek.Right {
				cur = peek.Right
			} else {
				if peek.Val >= maxpeek {
					count++
				}
				lastVisited = peek
				stack = stack[:len(stack)-1]
				maxstack = maxstack[:len(maxstack)-1]
			}
		}
	}
    return count
}

func main() {
	root := &TreeNode{Val: 1}
	fmt.Println(goodNodes(root))
	root = &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 5}
	fmt.Println(goodNodes(root))
	root = &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 1}
	fmt.Println(goodNodes(root))
	fmt.Println(goodNodes(nil))
}
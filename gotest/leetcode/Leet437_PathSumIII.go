package main

import "fmt"

// https://leetcode.com/problems/path-sum-iii/

// You are given a binary tree in which each node contains an integer value.
// Find the number of paths that sum to a given value.
// The path does not need to start or end at the root or a leaf, but it must go
// downwards (traveling only from parent nodes to child nodes).
// The tree has no more than 1,000 nodes and the values
// are in the range -1,000,000 to 1,000,000.
// Example:
//   root = [10,5,-3,3,2,null,11,3,-2,null,1], sum = 8
//           10
//          /  \
//         5   -3
//        / \    \
//       3   2   11
//      / \   \
//     3  -2   1
//   Return 3. The paths that sum to 8 are:
//     1.  5 -> 3
//     2.  5 -> 2 -> 1
//     3. -3 -> 11

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	// post-order traverse the tree, when visit each node, check sum through the path to root
	count := 0
	stack := make([]*TreeNode, 0, 20)
	cur := root
	var lastvisit *TreeNode = nil
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			peek := stack[len(stack)-1]
			if peek.Right != nil && peek.Right != lastvisit {
				cur = peek.Right
			} else {
				// visit node
				cursum := 0
				for i := len(stack) - 1; i >= 0; i-- {
					cursum += stack[i].Val
					if sum == cursum {
						count++
					}
				}
				lastvisit = peek
				// pop
				stack = stack[:len(stack)-1]
			}
		}
	}
	return count
}

func main() {
	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: -3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 3}
	root.Left.Left.Right = &TreeNode{Val: -2}
	root.Left.Right.Right = &TreeNode{Val: 1}
	fmt.Println(pathSum(root, 8))
}

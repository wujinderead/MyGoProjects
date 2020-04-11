package main

import (
	"fmt"
)

// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/

// Given a binary tree root. Split the binary tree into two subtrees by removing
// 1 edge such that the product of the sums of the subtrees are maximized.
// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: root = [1,2,3,4,5,6]
//   Output: 110
//           1                 1
//         /  \              x  \
//        2   3             2   3
//       / \   \           / \   \
//      4   5   6         4   5   6
//   Explanation: Remove the red edge and get 2 binary trees with sum 11 and 10.
//     Their product is 110 (11*10)
// Example 2:
//   Input: root = [1,null,2,3,4,null,null,5,6]
//   Output: 90
//   Explanation:  Remove the red edge and get 2 binary trees with sum 15 and 6.
//     Their product is 90 (15*6)
//                       1
//                           2
//                         3   4
//                            5  6
// Example 3:
//   Input: root = [2,3,9,10,7,8,6,5,4,11,1]
//   Output: 1025
//                   2
//           3              9
//       10    7        8       6
//     5   4   11  1
// Example 4:
//   Input: root = [1,1]
//   Output: 1
// Constraints:
//   Each tree has at most 50000 nodes and at least 2 nodes.
//   Each node's value is between [1, 10000].

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// first turn to get the total sum, second turn to find the node
// that the sum of all child nodes is the most adjacent to the half
func maxProduct(root *TreeNode) int {
	stack := make([]*TreeNode, 0, 30)
	allsum := getAllSum(root, stack)
	target := new(int)
	getSum(root, allsum, target)
	return *target%1000000007
}

func getSum(t *TreeNode, allsum int, target *int) int {
	sum := t.Val
	if t.Left != nil {
		sum += getSum(t.Left, allsum, target)
	}
	if t.Right != nil {
		sum += getSum(t.Right, allsum, target)
	}
	if sum*(allsum-sum)>*target {
		*target = sum*(allsum-sum)
	}
	return sum
}

func getAllSum(root *TreeNode, stack []*TreeNode) int {
	stack = append(stack, root)
	allsum := 0
	for len(stack)>0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		allsum += cur.Val
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return allsum
}

func main() {
    root := &TreeNode{Val: 1}
    root.Left = &TreeNode{Val: 2}
    root.Right = &TreeNode{Val: 3}
    root.Left.Left = &TreeNode{Val: 4}
    root.Left.Right = &TreeNode{Val: 5}
    root.Right.Right = &TreeNode{Val: 6}
    fmt.Println(maxProduct(root))

    root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 1}
	fmt.Println(maxProduct(root))

	root = &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 9}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: 10}
	root.Right.Left = &TreeNode{Val: 8}
	root.Right.Right = &TreeNode{Val: 6}
	root.Left.Left.Left = &TreeNode{Val: 5}
	root.Left.Left.Right = &TreeNode{Val: 4}
	root.Left.Right.Left = &TreeNode{Val: 11}
	root.Left.Right.Right = &TreeNode{Val: 1}
	fmt.Println(maxProduct(root))

	root = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Left = &TreeNode{Val: 5}
	root.Right.Right.Right = &TreeNode{Val: 6}
	fmt.Println(maxProduct(root))
}
package main

import (
    "fmt"
)

// https://leetcode.com/problems/distribute-coins-in-binary-tree/

// Given the root of a binary tree with N nodes, each node in the tree has node.val coins, 
// and there are N coins total. In one move, we may choose two adjacent nodes and move one 
// coin from one node to another. (The move may be from parent to child, or from child to parent.) 
// Return the number of moves required to make every node have exactly one coin.
// Example 1: 
//       3
//      / \
//     0   0
//   Input: [3,0,0]
//   Output: 2
//   Explanation: From the root of the tree, we move one coin to its left child, and one coin to its right child.
// Example 2: 
//       0
//      / \
//     3   0
//   Input: [0,3,0]
//   Output: 3
//   Explanation: From the left child of the root, we move two coins to the root [taking two moves].  
//     Then, we move one coin from the root of the tree to the right child.
// Example 3: 
//       1
//      / \
//     0   2
//  Input: [1,0,2]
//  Output: 2
// Example 4: 
//       1
//      / \
//     0   0
//      \
//       3
//   Input: [1,0,0,null,3]
//   Output: 4
// Note: 
//   1<= N <= 100 
//   0 <= node.val <= N 

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
    move := 0
    dfs(root, &move)
    return move
}

func dfs(r *TreeNode, move *int) int {
	if r==nil {
		return 0
	}
	a := r.Val+dfs(r.Left, move)+dfs(r.Right, move)  // t coins we get from child
	*move += abs(a-1)   // if a-1>0, we give coins to parent; if a-1<0, we get coins from parent
	return a-1
}

func abs(a int) int {
	if a<0 {
		return -a
	}
	return a
}

func main() {
	r := &TreeNode{Val: 3}
	r.Left = &TreeNode{Val: 0}
	r.Right = &TreeNode{Val: 0}
	fmt.Println(distributeCoins(r))

	r = &TreeNode{Val: 0}
	r.Left = &TreeNode{Val: 3}
	r.Right = &TreeNode{Val: 0}
	fmt.Println(distributeCoins(r))

	r = &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 0}
	r.Right = &TreeNode{Val: 2}
	fmt.Println(distributeCoins(r))

	r = &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 0}
	r.Right = &TreeNode{Val: 0}
	r.Right.Right = &TreeNode{Val: 3}
	fmt.Println(distributeCoins(r))
}
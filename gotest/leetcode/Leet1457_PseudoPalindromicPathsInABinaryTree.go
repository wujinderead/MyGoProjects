package main

import "fmt"

// https://leetcode.com/problems/pseudo-palindromic-paths-in-a-binary-tree/

// Given a binary tree where node values are digits from 1 to 9. A path in the binary
// tree is said to be pseudo-palindromic if at least one permutation of the node values
// in the path is a palindrome.
// Return the number of pseudo-palindromic paths going from the root node to leaf nodes.
// Example 1: 
//               2
//              / \
//             3  1
//            / \  \
//           3  1  1
//   Input: root = [2,3,1,3,1,null,1]
//   Output: 2
//   Explanation: The figure above represents the given binary tree. There are three
//     paths going from the root node to leaf nodes: the red path [2,3,3], the green
//     path [2,1,1], and the path [2,3,1]. Among these paths only red path and green path
//     are pseudo-palindromic paths since the red path [2,3,3] can be rearranged in
//     [3,2,3] (palindrome) and the green path [2,1,1] can be rearranged in [1,2,1] (palindrome).
// Example 2: 
//                2
//               / \
//              1  1
//             / \
//            1  3
//                \
//                 1
//   Input: root = [2,1,1,1,3,null,null,null,null,null,1]
//   Output: 1
//   Explanation: The figure above represents the given binary tree. There are three paths
//     going from the root node to leaf nodes: the green path [2,1,1], the path
//     [2,1,3,1], and the path [2,1]. Among these paths only the green path is pseudo-
//     palindromic since [2,1,1] can be rearranged in [1,2,1] (palindrome).
// Example 3:
//   Input: root = [9]
//   Output: 1
// Constraints:
//   The given binary tree will have between 1 and 10^5 nodes.
//   Node values are digits from 1 to 9.

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pseudoPalindromicPaths (root *TreeNode) int {
	count := 0
	mask := 0
    visit(root, mask, &count)
    return count
}

func visit(cur *TreeNode, mask int, count *int) {
	mask = mask ^ (1<<uint(cur.Val-1))    // xor to indicate even/odd of this bit
	fmt.Println(cur.Val, mask)
	if cur.Left==nil && cur.Right==nil {
		if check(mask) {
			*count++
			return
		}
	}
	if cur.Left != nil {
		visit(cur.Left, mask, count)
	}
	if cur.Right != nil {
		visit(cur.Right, mask, count)
	}
}

func check(mask int) bool {   // check mask is 0 or only one bit is set
	if mask==0 {
		return true
	}
	for i:=0; i<9; i++ {
		if mask==1<<uint(i) {
			return true
		}
	}
	return false
}

func main() {
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 1}
	fmt.Println(pseudoPalindromicPaths(root))
}
package main

import "fmt"

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree

// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
// According to the definition of LCA on Wikipedia: “The lowest common ancestor
// is defined between two nodes p and q as the lowest node in T that has both p and
// q as descendants (where we allow a node to be a descendant of itself).”
// Given the following binary tree: root = [3,5,1,6,2,0,8,null,null,7,4]
// Example 1:
//          3
//        /  \
//       5   1
//      / \ / \
//     6  2 0 8
//       / \
//      7  4
//   Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//   Output: 3
//   Explanation: The LCA of nodes 5 and 1 is 3.
// Example 2:
//   Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
//   Output: 5
//   Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant
//     of itself according to the LCA definition.
// Note:
//   All of the nodes' values will be unique.
//   p and q are different and both values will exist in the binary tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var s1, s2 []*TreeNode
	stack := make([]*TreeNode, 0, 20)
	cur := root
	var lastVisit *TreeNode
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}
		// cur nil
		peek := stack[len(stack)-1]
		if peek.Right != nil && lastVisit != peek.Right {
			cur = peek.Right
		} else { // visit peek, post-order traverse make a path from the root to current node
			if peek == p {
				s1 = make([]*TreeNode, len(stack))
				copy(s1, stack)
			}
			if peek == q {
				s2 = make([]*TreeNode, len(stack))
				copy(s2, stack)
			}
			stack = stack[:len(stack)-1]
			lastVisit = peek
		}
	}
	i := 0
	for i+1 < len(s1) && i+1 < len(s2) && s1[i+1] == s2[i+1] { // find the split point of two paths
		i++
	}
	return s1[i]
}

func lowestCommonAncestorRecursive(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestorRecursive(root.Left, p, q)
	right := lowestCommonAncestorRecursive(root.Right, p, q)
	if left != nil && right != nil {
		return root
	} else if left != nil {
		return left
	}
	return right
}

func main() {
	r := &TreeNode{Val: 3}
	r.Left = &TreeNode{Val: 5}
	r.Right = &TreeNode{Val: 1}
	r.Left.Left = &TreeNode{Val: 6}
	r.Left.Right = &TreeNode{Val: 2}
	r.Right.Left = &TreeNode{Val: 0}
	r.Right.Right = &TreeNode{Val: 8}
	r.Left.Right.Left = &TreeNode{Val: 7}
	r.Left.Right.Right = &TreeNode{Val: 4}
	for _, v := range []struct {
		p, q *TreeNode
	}{
		//{r.Left, r.Right.Left},
		//{r, r.Right.Left},
		//{r.Right.Right, r.Right.Left},
		{r.Left.Left, r.Left.Right.Right},
	} {
		fmt.Println(lowestCommonAncestor(r, v.p, v.q))
		fmt.Println(lowestCommonAncestorRecursive(r, v.p, v.q))
		fmt.Println(lowestCommonAncestorPreOrderStack(r, v.p, v.q))
	}
}

package main

import (
	"fmt"
)

// https://leetcode.com/problems/all-nodes-distance-k-in-binary-tree/

// We are given a binary tree (with root node root), a target node, and an integer value K.
// Return a list of the values of all nodes that have a distance K from the target node.  
// The answer can be returned in any order.
// Example 1:
//   Input: root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, K = 2
//   Output: [7,4,1]
//			 3
//		   /   \
//		  5     1
//		/  \   / \
//	   6   2  0  8
//		  / \
//		 7  4
//   Explanation: 
//     The nodes that are a distance 2 from the target node (with value 5)
//     have values 7, 4, and 1.
//     Note that the inputs "root" and "target" are actually TreeNodes.
//     The descriptions of the inputs above are just serializations of these objects.
// Note:
//   The given tree is non-empty.
//   Each node in the tree has unique values 0 <= node.val <= 500.
//   The target node is a node in the tree.
//   0 <= K <= 1000.

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	ans := make([]int, 0, 10)
	find(root, target.Val, K, &ans)
	return ans
}

// find target in children, if not find return 0, if found return distance to target.
func find(t *TreeNode, target int, K int, ans *[]int) int {
	if t.Val==target {    // find target
		findKChild(t, K, ans)
		return 1  // parent is 1 step from target
	}
	l, r := 0, 0
	if t.Left != nil {    // find target in left
		l = find(t.Left, target, K, ans)
		if l>0 && K-l>0 {
			findKChild(t.Right, K-l-1, ans)  // target is in left, find K-l-1 in right
		}
		if K>0 && K==l {
			*ans = append(*ans, t.Val)
		}
	}
	if t.Right != nil && l==0 {   // no target in left, find in right
		r = find(t.Right, target, K, ans)
		if r>0 && K-r>0 {
			findKChild(t.Left, K-r-1, ans)
		}
		if K>0 && K==r {
			*ans = append(*ans, t.Val)
		}
	}

	if l==0 && r==0 {
		return 0
	}
	if l>r {
		return l+1
	}
	return r+1
}

func findKChild(r *TreeNode, distance int, ans *[]int) {
	if r==nil {
		return
	}
	if distance==0 {
		*ans = append(*ans, r.Val)
		return
	}
	findKChild(r.Left, distance-1, ans)
	findKChild(r.Right, distance-1, ans)
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

	fmt.Println(distanceK(r, r.Left, 2))
	fmt.Println(distanceK(r, r.Left.Left, 3))
	fmt.Println(distanceK(r, r.Right.Left, 3))
	fmt.Println(distanceK(r, r.Right.Left, 4))
	fmt.Println(distanceK(r, r.Right.Left, 2))
	fmt.Println(distanceK(r, r.Right.Left, 6))
	fmt.Println(distanceK(r, r.Right.Left, 1))
	fmt.Println(distanceK(r, r.Right.Left, 0))
}
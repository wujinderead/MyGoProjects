package main

import (
    "fmt"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/

// Given a rooted binary tree, return the lowest common ancestor of its deepest leaves. 
// Recall that: 
//   The node of a binary tree is a leaf if and only if it has no children 
//   The depth of the root of the tree is 0, and if the depth of a node is d, the 
// depth of each of its children is d+1. 
//   The lowest common ancestor of a set S of nodes is the node A with the largest
// depth such that every node in S is in the subtree with root A. 
// Example 1: 
//   Input: root = [1,2,3]
//   Output: [1,2,3]
//   Explanation: 
//     The deepest leaves are the nodes with values 2 and 3.
//     The lowest common ancestor of these leaves is the node with value 1.
//     The answer returned is a TreeNode object (not an array) with serialization "[1,2,3]".
// Example 2: 
//   Input: root = [1,2,3,4]
//   Output: [4]
// Example 3: 
//   Input: root = [1,2,3,4,5]
//   Output: [2,4,5]
// Constraints: 
//   The given tree will have between 1 and 1000 nodes. 
//   Each node of the tree will have a distinct value between 1 and 1000. 

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lcaDeepestLeavesOnePass(root *TreeNode) *TreeNode {
	lca, _ := onePassHelper(root)
    return lca
}

func onePassHelper(root *TreeNode) (*TreeNode, int) {
	if root==nil {
		return nil, 0
	}
    llca, lhei := onePassHelper(root.Left)
    rlca, rhei := onePassHelper(root.Right)
    if lhei>rhei {
    	return llca, lhei+1
    }
    if rhei>lhei {
    	return rlca, rhei+1
    }
    // if two child heights equal, cur node is the lca for subtree cur.
    return root, lhei+1
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	maxlen := 0
    first(0, root, &maxlen)   // get max height
    var lca *TreeNode
    second(0, root, &lca, maxlen)
    return lca
}

func first(plen int, cur *TreeNode, maxlen *int) {
	if cur==nil {
		return
	}
	if plen+1>*maxlen {
		*maxlen = plen+1
	}
	first(plen+1, cur.Left, maxlen)
	first(plen+1, cur.Right, maxlen)
}

func second(plen int, cur *TreeNode, lca **TreeNode, maxlen int) bool {
	if cur == nil {
		return false
	}
	if plen+1==maxlen {
		*lca = cur
		return true
	}
	b1 := second(plen+1, cur.Left, lca, maxlen) 
	b2 := second(plen+1, cur.Right, lca, maxlen) 
	if b1 && b2 {
		*lca = cur
		return true
	}
	return b1 || b2
}

func main() {
	for _, t := range []func(*TreeNode) *TreeNode {lcaDeepestLeaves, lcaDeepestLeavesOnePass} {
		r := &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 2}
		r.Right = &TreeNode{Val: 3}
		fmt.Println(t(r))
	
		r = &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 2}
		r.Right = &TreeNode{Val: 3}
		r.Left.Left = &TreeNode{Val: 4}
		fmt.Println(t(r))
	
		r = &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 2}
		r.Right = &TreeNode{Val: 3}
		r.Left.Left = &TreeNode{Val: 4}
		r.Left.Right = &TreeNode{Val: 5}
		fmt.Println(t(r))
	
		r = &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 2}
		r.Right = &TreeNode{Val: 3}
		r.Left.Left = &TreeNode{Val: 4}
		r.Left.Right = &TreeNode{Val: 5}
		r.Right.Right = &TreeNode{Val: 6}
		fmt.Println(t(r))
	
		r = &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 2}
		r.Right = &TreeNode{Val: 3}
		r.Left.Left = &TreeNode{Val: 4}
		r.Left.Right = &TreeNode{Val: 5}
		r.Right.Right = &TreeNode{Val: 6}
		rr := &TreeNode{Val: 7}
		rr.Left = r
		rr.Right = &TreeNode{Val: 8}
		rr.Right.Right = &TreeNode{Val: 9}
		fmt.Println(t(rr))
	
		r = &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 2}
		r.Right = &TreeNode{Val: 3}
		r.Left.Left = &TreeNode{Val: 4}
		r.Left.Right = &TreeNode{Val: 5}
		r.Right.Right = &TreeNode{Val: 6}
		rr = &TreeNode{Val: 7}
		rr.Left = r
		rr.Right = &TreeNode{Val: 8}
		rr.Right.Right = &TreeNode{Val: 9}
		rr.Right.Right.Left = &TreeNode{Val: 10}
		fmt.Println(t(rr))
	}
}
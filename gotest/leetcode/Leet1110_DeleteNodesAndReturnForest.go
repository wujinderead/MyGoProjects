package main

import (
    "fmt"
)

// https://leetcode.com/problems/delete-nodes-and-return-forest/

// Given the root of a binary tree, each node in the tree has a distinct value. After deleting 
// all nodes with a value in to_delete, we are left with a forest (a disjoint union of trees). 
// Return the roots of the trees in the remaining forest. You may return the result in any order. 
// Example 1:
//   Input: root = [1,2,3,4,5,6,7], to_delete = [3,5]
//   Output: [[1,2,null,4],[6],[7]]
// Constraints: 
//   The number of nodes in the given tree is at most 1000. 
//   Each node has a distinct value between 1 and 1000. 
//   to_delete.length <= 1000 
//   to_delete contains distinct values between 1 and 1000. 

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	delmap := make(map[int]struct{})
	for _, v := range to_delete {
		delmap[v] = struct{}{}
	}
	delmap[0] = struct{}{}
    dummy := &TreeNode{Left: root}
    forest := make([]*TreeNode, 0)
    visit(dummy, dummy, delmap, &forest)   // we need also to delete dummy
    return forest
}

func visit(cur, parent *TreeNode, delmap map[int]struct{}, forest *[]*TreeNode) {
	if cur.Left != nil {
		visit(cur.Left, cur, delmap, forest)
	}
	if cur.Right != nil {
		visit(cur.Right, cur, delmap, forest)
	}
	if _, ok := delmap[cur.Val]; ok {  // cur is to delete
		if cur.Left != nil {
			*forest = append(*forest, cur.Left)
		}
		if cur.Right != nil {
			*forest = append(*forest, cur.Right)
		}
		if cur==parent.Left {
			parent.Left = nil 
		} else {
			parent.Right = nil
		}
	}
}

func main() {
	r := &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	r.Left.Left = &TreeNode{Val: 4}
	r.Left.Right = &TreeNode{Val: 5}
	r.Right.Left = &TreeNode{Val: 6}
	r.Right.Right = &TreeNode{Val: 7}
	for _, v := range delNodes(r, []int{3,5}) {
		fmt.Println(v)
	}
	fmt.Println()

	r = &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	for _, v := range delNodes(r, []int{2}) {
		fmt.Println(v)
	}
	fmt.Println()
	
	r = &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	for _, v := range delNodes(r, []int{1}) {
		fmt.Println(v)
	}
	fmt.Println()
	
	r = &TreeNode{Val: 1}
	r.Left = &TreeNode{Val: 2}
	r.Right = &TreeNode{Val: 3}
	for _, v := range delNodes(r, []int{1,2}) {
		fmt.Println(v)
	}
}
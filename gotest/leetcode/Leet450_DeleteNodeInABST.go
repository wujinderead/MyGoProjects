package main

import "fmt"

// https://leetcode.com/problems/delete-node-in-a-bst/

// Given a root node reference of a BST and a key, delete the node with the given
// key in the BST. Return the root node reference (possibly updated) of the BST. 
// Basically, the deletion can be divided into two stages:
// Search for a node to remove. If the node is found, delete the node. 
// Note: Time complexity should be O(height of tree). 
// Example:
//   root = [5,3,6,2,4,null,7]
//   key = 3     
//         5
//        / \
//       3   6
//      / \   \
//     2   4   7
// Given key to delete is 3. So we find the node with value 3 and delete it.
// One valid answer is [5,4,6,2,null,null,7], shown in the following BST.    
//         5
//        / \
//       4   6
//      /     \
//     2       7    
// Another valid answer is [5,2,6,null,4,null,7].    
//         5
//        / \
//       2   6
//        \   \
//         4   7

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
    stack := make([]*TreeNode, 0)  // space O(h), can be O(1) by use pointers
    cur := root
    for cur != nil {
    	if key < cur.Val {
    		stack = append(stack, cur)
    		cur = cur.Left
    		continue
    	} 
    	if key > cur.Val {
    		stack = append(stack, cur)
    		cur = cur.Right
    		continue
    	}

    	// cur.Val==key, found
    	// if both childs exist, find successor, and the successor must be its child
    	if cur.Left != nil && cur.Right != nil {  
    		t := cur.Right
    		stack = append(stack, cur)
    		for t.Left != nil {
    			stack = append(stack, t)
    			t = t.Left
    		}
    		cur.Val = t.Val
    		cur = t
    	}

    	// delete cur
    	p := (*TreeNode)(nil)               // cur's parent
    	if len(stack)>0 {
    		p = stack[len(stack)-1]
    	}
    	child := cur.Left   // cur's non-nil child; if both child nil, child is nil
    	if child == nil {
    		child = cur.Right
    	}
    	if p==nil {   // we delete root
    		return child
    	}
    	if p.Left == cur {
    		p.Left = child
    	} else {
    		p.Right = child
    	}
    	break
    }
    // default
    return root
}

func main() {
	r := &TreeNode{Val: 5}
	r.Left = &TreeNode{Val: 3}
	r.Right = &TreeNode{Val: 6}
	r.Left.Left = &TreeNode{Val: 2}
	r.Left.Right = &TreeNode{Val: 4}
	r.Right.Right = &TreeNode{Val: 7}
	r = deleteNode(r, 3)
	fmt.Println(r.Left, r.Left.Right)

	r = &TreeNode{Val: 5}
	r = deleteNode(r, 3)
	fmt.Println(r)

	r = &TreeNode{Val: 5}
	r = deleteNode(r, 5)
	fmt.Println(r)

	r = &TreeNode{Val: 5}
	r.Left = &TreeNode{Val: 3}
	r = deleteNode(r, 5)
	fmt.Println(r)

	r = &TreeNode{Val: 5}
	r.Left = &TreeNode{Val: 6}
	r = deleteNode(r, 5)
	fmt.Println(r)

	r = &TreeNode{Val: 4}
	r.Left = &TreeNode{Val: 1}
	r.Right = &TreeNode{Val: 7}
	r.Right.Left = &TreeNode{Val: 5}
	r.Right.Right = &TreeNode{Val: 8}
	r.Right.Left.Right = &TreeNode{Val: 6}
	r = deleteNode(r, 4)
	fmt.Println(r, r.Right, r.Right.Left)

	r = &TreeNode{Val: 1}
	r.Right = &TreeNode{Val: 2}
	r = deleteNode(r, 2)
	fmt.Println(r)
}
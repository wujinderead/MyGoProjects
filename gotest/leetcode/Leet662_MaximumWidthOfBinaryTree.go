package main

import (
	"container/list"
    "fmt"
)

// https://leetcode.com/problems/maximum-width-of-binary-tree/

// Given a binary tree, write a function to get the maximum width of the given tree. 
// The width of a tree is the maximum width among all levels. The binary tree has the same 
// structure as a full binary tree, but some nodes are null. The width of one level is defined 
// as the length between the end-nodes (the leftmost and right most non-null nodes in the level, 
// where the null nodes between the end-nodes are also counted into the length calculation.
// Example 1:
//   Input: 
//              1
//            /   \
//           3     2
//          / \     \  
//         5   3     9  
//   Output: 4
//   Explanation: The maximum width existing in the third level with the length 4 (5,3,null,9).
// Example 2:
//   Input:   
//             1
//            /  
//           3    
//          / \       
//         5   3     
//   Output: 2
//   Explanation: The maximum width existing in the third level with the length 2 (5,3).
// Example 3:
//   Input: 
//             1
//            / \
//           3   2 
//          /        
//         5      
//   Output: 2
//   Explanation: The maximum width existing in the second level with the length 2 (3,2).
// Example 4:
//   Input: 
//             1
//            / \
//           3   2
//          /     \  
//         5       9 
//        /         \
//       6           7
//   Output: 8
//   Explanation:The maximum width existing in the fourth level with the length 8 (6,null,null,null,null,null,null,7).
//   Note: Answer will in the range of 32-bit signed integer.

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 1
    queue := list.New()
    queue.PushBack(struct{*TreeNode;int}{root,1})
    for queue.Len()>0 {
    	L := queue.Len()
    	left, right := 0, 0
    	for k:=0; k<L; k++ {
    		pair := queue.Remove(queue.Front()).(struct{*TreeNode;int})
    		node, number := pair.TreeNode, pair.int
    		if k==0 {
    			left = number
    		} 
    		if k==L-1 {
    			right = number
    		}
    		if node.Left != nil {
    			queue.PushBack(struct{*TreeNode;int}{node.Left, number*2})
    		}
    		if node.Right != nil {
    			queue.PushBack(struct{*TreeNode;int}{node.Right, number*2+1})
    		}
    	}
    	if right-left+1>max {
    		max = right-left+1
    	}
    }
    return max
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 3}
	fmt.Println(widthOfBinaryTree(root))

	root.Right.Right = nil
	root.Right.Left = &TreeNode{Val: 1}
	fmt.Println(widthOfBinaryTree(root))
}
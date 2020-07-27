package main

import (
    "fmt"
)

// https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/

// Given the root of a binary tree and an integer distance. A pair of two different leaf nodes of 
// a binary tree is said to be good if the length of the shortest path between them is less than 
// or equal to distance. Return the number of good leaf node pairs in the tree.
// Example 1:
//           1
//          / \
//         2  3
//         \
//         4
//   Input: root = [1,2,3,null,4], distance = 3
//   Output: 1
//   Explanation: The leaf nodes of the tree are 3 and 4 and the length of the shortest path between them is 3.
//     This is the only good pair.
// Example 2:
//            1
//          /  \
//         2    3
//        / \  / \
//       4  5 6  7
//   Input: root = [1,2,3,4,5,6,7], distance = 3
//   Output: 2
//   Explanation: The good pairs are [4,5] and [6,7] with shortest path = 2. 
//     The pair [4,6] is not good because the length of ther shortest path between them is 4.
// Example 3:
//   Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
//   Output: 1
//   Explanation: The only good pair is [2,5].
// Example 4:
//   Input: root = [100], distance = 1
//   Output: 0
// Example 5:
//   Input: root = [1,1,1], distance = 2
//   Output: 1
// Constraints:
//   The number of nodes in the tree is in the range [1, 2^10].
//   Each node's value is between [1, 100].
//   1 <= distance <= 10

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	if distance==1 {  // distance between 2 leaves can't <= 1
		return 0
	}
	count := new(int)
	helper(root, distance, count)
	return *count
}

func helper(cur *TreeNode, distance int, count *int) [10]int {
	if cur.Left == nil && cur.Right == nil {   // if leaf, height is 1
		return [10]int{1: 1}
	} 
	left, right := [10]int{}, [10]int{}
	if cur.Left != nil {
		left = helper(cur.Left, distance, count)  // get leaves heights for child
	}
	if cur.Right != nil {
		right = helper(cur.Right, distance, count)
	}
	for i:=1; i<distance; i++ {
		for j:=1; i+j<=distance; j++ {    // check how many pairs with distance <= target
			*count += left[i]*right[j]
		}
	}
	for i:=9; i>0; i-- {
		left[i] = left[i-1]+right[i-1]    // increment each leaf's height 
	}
	return left
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	fmt.Println(countPairs(root, 3), 1)

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(countPairs(root, 3), 2)

	root = &TreeNode{Val: 1}
	fmt.Println(countPairs(root, 1), 0)	
	fmt.Println(countPairs(root, 2), 0)	

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(countPairs(root, 2), 1)

	root = &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 3}
	root.Right.Right.Right = &TreeNode{Val: 2}
	fmt.Println(countPairs(root, 3), 1)
}
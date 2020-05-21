package main

import (
	"fmt"
	"container/list"
)

// https://leetcode.com/problems/n-ary-tree-level-order-traversal/

// Given an n-ary tree, return the level order traversal of its nodes' values.
// Nary-Tree input serialization is represented in their level order traversal, 
// each group of children is separated by the null value (See examples).
// Example 1: 
//            1
//         /  |  \
//        3   2   4
//       / \
//      5  6
//   Input: root = [1,null,3,2,4,null,5,6]
//   Output: [[1],[3,2,4],[5,6]]
// Example 2:
//   Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
//   Output: [[1],[2,3,4,5],[6,7,8,9,10],[11,12,13],[14]]
// Constraints:
//   The height of the n-ary tree is less than or equal to 1000
//   The total number of nodes is between [0, 10^4]

type Node struct {
	Val int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
    levels := make([][]int, 0)
    if root==nil {
    	return levels
	}
	queue := list.New()
	queue.PushBack(struct{*Node;int}{root, 1})
	for queue.Len()>0 {
		t := queue.Remove(queue.Front()).(struct{*Node;int})
		node, level := t.Node, t.int
		if level>len(levels) {
			levels = append(levels, []int{node.Val})
		} else {
			levels[level-1] = append(levels[level-1], node.Val)
		}
		for _, v := range node.Children {
			queue.PushBack(struct {*Node;int}{v, level+1})
		}
	}
	return levels
}

func main() {
	root := &Node{Val: 3}
	root.Children = []*Node{{Val: 3}, {Val: 2}, {Val: 4}}
	root.Children[0].Children = []*Node{{Val: 5}, {Val: 6}}
	fmt.Println(levelOrder(root))
	root = &Node{Val: 3}
	fmt.Println(levelOrder(root))
	fmt.Println(levelOrder(nil))
}
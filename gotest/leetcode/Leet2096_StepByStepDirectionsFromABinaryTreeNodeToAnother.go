package main

import "fmt"

// https://leetcode.com/problems/step-by-step-directions-from-a-binary-tree-node-to-another/

// You are given the root of a binary tree with n nodes. Each node is uniquely assigned a value
// from 1 to n. You are also given an integer startValue representing the value of the start node s,
// and a different integer destValue representing the value of the destination node t.
// Find the shortest path starting from node s and ending at node t. Generate step-by-step directions
// of such path as a string consisting of only the uppercase letters 'L', 'R', and 'U'. Each letter
// indicates a specific direction:
//   'L' means to go from a node to its left child node.
//   'R' means to go from a node to its right child node.
//   'U' means to go from a node to its parent node.
// Return the step-by-step directions of the shortest path from node s to node t.
// Example 1:
//   Input: root = [5,1,2,3,null,6,4], startValue = 3, destValue = 6
//   Output: "UURL"
//   Explanation: The shortest path is: 3 → 1 → 5 → 2 → 6.
// Example 2:
//   Input: root = [2,1], startValue = 2, destValue = 1
//   Output: "L"
//   Explanation: The shortest path is: 2 → 1.
// Constraints:
//   The number of nodes in the tree is n.
//   2 <= n <= 10⁵
//   1 <= Node.val <= n
//   All the values in the tree are unique.
//   1 <= startValue, destValue <= n
//   startValue != destValue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// find path from root to start and dest, e.g. "LLRR", "LLLR"
// then exclude the same prefix, it means from their nearest common ancestor,
// the paths are "RR" and "LR", so reverse first path then plus the second path.
// final answer is "UULR"
func getDirections(root *TreeNode, startValue int, destValue int) string {
	sbuf, dbuf := []byte{}, []byte{}
	visit(root, startValue, &sbuf)
	visit(root, destValue, &dbuf)
	//fmt.Println(string(sbuf), string(dbuf))
	i := -1
	for i = -1; i+1 < len(sbuf) && i+1 < len(dbuf); i++ {
		if sbuf[i+1] != dbuf[i+1] {
			break
		}
	}
	ans := []byte{}
	for k := len(sbuf) - 1; k > i; k-- {
		ans = append(ans, 'U')
	}
	for k := i + 1; k < len(dbuf); k++ {
		ans = append(ans, dbuf[k])
	}
	return string(ans)
}

func visit(root *TreeNode, tar int, buf *[]byte) bool {
	if root.Val == tar {
		return true
	}
	var v1, v2 bool
	if root.Left != nil {
		*buf = append(*buf, 'L')
		v1 = visit(root.Left, tar, buf)
		if !v1 {
			*buf = (*buf)[:len(*buf)-1]
		}
	}
	if !v1 && root.Right != nil {
		*buf = append(*buf, 'R')
		v2 = visit(root.Right, tar, buf)
		if !v2 {
			*buf = (*buf)[:len(*buf)-1]
		}
	}
	return v1 || v2
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
	for i := 1; i <= 5; i++ {
		for j := i + 1; j <= 6; j++ {
			fmt.Println(i, j, getDirections(root, i, j))
			fmt.Println(j, i, getDirections(root, j, i))
		}
	}
}

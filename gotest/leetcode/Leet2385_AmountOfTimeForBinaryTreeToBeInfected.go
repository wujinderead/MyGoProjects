package main

import "fmt"

// https://leetcode.com/problems/amount-of-time-for-binary-tree-to-be-infected/

// You are given the root of a binary tree with unique values, and an integer
// start. At minute 0, an infection starts from the node with value start.
// Each minute, a node becomes infected if:
//   The node is currently uninfected.
//   The node is adjacent to an infected node.
// Return the number of minutes needed for the entire tree to be infected.
// Example 1:
//   Input: root = [1,5,3,null,4,10,6,9,2], start = 3
//   Output: 4
//   Explanation: The following nodes are infected during:
//   - Minute 0: Node 3
//   - Minute 1: Nodes 1, 10 and 6
//   - Minute 2: Node 5
//   - Minute 3: Node 4
//   - Minute 4: Nodes 9 and 2
//   It takes 4 minutes for the whole tree to be infected so we return 4.
// Example 2:
//   Input: root = [1], start = 1
//   Output: 0
//   Explanation: At minute 0, the only node in the tree is infected so we return 0.
// Constraints:
//   The number of nodes in the tree is in the range [1, 10⁵].
//   1 <= Node.val <= 10⁵
//   Each node has a unique value.
//   A node with a value of start exists in the tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	max := 0
	maxDist(root, start, &max)
	return max
}

// first return variable: the height of root
// second return variable: the distance between root and start
func maxDist(root *TreeNode, start int, max *int) (int, int) {
	if root == nil {
		return 0, -1
	}
	lHei, lSta := maxDist(root.Left, start, max)
	rHei, rSta := maxDist(root.Right, start, max)
	hei := intMax(lHei, rHei) + 1
	if root.Val == start {
		*max = intMax(*max, hei-1) // find the start, max distance is height
		return hei, 0
	} else if lSta >= 0 { // left child contains start: distance to start + right height
		*max = intMax(*max, lSta+1+rHei)
		return hei, lSta + 1
	} else if rSta >= 0 {
		*max = intMax(*max, rSta+1+lHei)
		return hei, rSta + 1
	} else { // neither child nor self contains start
		return hei, -1
	}
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	{
		r := &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 2}
		r.Left.Left = &TreeNode{Val: 4}
		r.Left.Left.Left = &TreeNode{Val: 6}
		r.Left.Left.Right = &TreeNode{Val: 7}
		r.Left.Right = &TreeNode{Val: 5}
		r.Left.Right.Left = &TreeNode{Val: 8}
		r.Left.Right.Right = &TreeNode{Val: 9}
		r.Right = &TreeNode{Val: 3}
		r.Right.Left = &TreeNode{Val: 10}
		r.Right.Right = &TreeNode{Val: 11}
		for i := 1; i < 12; i++ {
			fmt.Println(i, amountOfTime(r, i))
		}
	}
	{
		r := &TreeNode{Val: 19}
		r.Left = &TreeNode{Val: 4}
		r.Left.Left = &TreeNode{Val: 9}
		r.Left.Right = &TreeNode{Val: 11}
		r.Left.Right.Right = &TreeNode{Val: 17}
		r.Right = &TreeNode{Val: 3}
		r.Right.Left = &TreeNode{Val: 18}
		r.Right.Left.Right = &TreeNode{Val: 8}
		r.Right.Left.Right.Left = &TreeNode{Val: 20}
		r.Right.Right = &TreeNode{Val: 1}
		fmt.Println(amountOfTime(r, 1))
	}
}

package leetcode

import "fmt"

// https://leetcode.com/problems/unique-binary-search-trees-ii/

// Given an integer n, generate all structurally unique
// BST's (binary search trees) that store values 1 ... n.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nilTree = []*TreeNode{nil}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	arr := make([]int, n)
	for i := range arr {
		arr[i] = i + 1
	}
	return generateTreeArray(arr)
}

func generateTreeArray(arr []int) []*TreeNode {
	if len(arr) == 0 {
		return nilTree
	}
	trees := make([]*TreeNode, 0)
	for i := 0; i < len(arr); i++ {
		// arr[i] as root
		lefts := generateTreeArray(arr[0:i])
		rights := generateTreeArray(arr[i+1:])
		for _, left := range lefts {
			for _, right := range rights {
				trees = append(trees, &TreeNode{arr[i], left, right})
			}
		}
	}
	return trees
}

func main() {
	trees := generateTrees(4)
	for _, tree := range trees {
		fmt.Println(tree.Val)
	}
}

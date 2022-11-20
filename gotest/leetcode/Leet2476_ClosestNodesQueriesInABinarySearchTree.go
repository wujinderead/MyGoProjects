package main

import "fmt"

// https://leetcode.com/problems/closest-nodes-queries-in-a-binary-search-tree/

// You are given the root of a binary search tree and an array queries of size n
// consisting of positive integers.
// Find a 2D array answer of size n where answer[i] = [mini, maxi]:
// mini is the largest value in the tree that is smaller than or equal to queries[i]. If a such
// value does not exist, add -1 instead.
// maxi is the smallest value in the tree that is greater than or equal to queries[i]. If a such
// value does not exist, add -1 instead.
// Return the array answer.
// Example 1:
//   Input: root = [6,2,13,1,4,9,15,null,null,null,null,null,null,14], queries = [2,5,16]
//   Output: [[2,2],[4,6],[15,-1]]
//   Explanation: We answer the queries in the following way:
//     - The largest number that is smaller or equal than 2 in the tree is 2, and the
//       smallest number that is greater or equal than 2 is still 2. So the answer for
//       the first query is [2,2].
//     - The largest number that is smaller or equal than 5 in the tree is 4, and the
//       smallest number that is greater or equal than 5 is 6. So the answer for the
//       second query is [4,6].
//     - The largest number that is smaller or equal than 16 in the tree is 15, and
//       the smallest number that is greater or equal than 16 does not exist. So the answer
//       for the third query is [15,-1].
// Example 2:
//   Input: root = [4,null,9], queries = [3]
//   Output: [[-1,4]]
//   Explanation: The largest number that is smaller or equal to 3 in the tree does
//     not exist, and the smallest number that is greater or equal to 3 is 4. So the
//     answer for the query is [-1,4].
// Constraints:
//   The number of nodes in the tree is in the range [2, 10⁵].
//   1 <= Node.val <= 10⁶
//   n == queries.length
//   1 <= n <= 10⁵
//   1 <= queries[i] <= 10⁶

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	// inorder traversal to get all sorted nodes
	stack := make([]*TreeNode, 0, 10)
	arr := make([]int, 0, 10)
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			arr = append(arr, cur.Val)
			cur = cur.Right
		}
	}
	// binary search for each query
	ans := make([][]int, len(queries))
	for i := range queries {
		q := queries[i]
		left, right := 0, len(arr)-1
		for left <= right {
			mid := (left + right) / 2
			if arr[mid] < q {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		// left is the first index that arr[left] >= q
		if left == len(arr) { // no value in arr >= q
			ans[i] = []int{arr[len(arr)-1], -1}
		} else if arr[left] == q { // find q in arr
			ans[i] = []int{q, q}
		} else if left == 0 { // arr[0] > q
			ans[i] = []int{-1, arr[left]}
		} else { // arr[left-1] < q < arr[left]
			ans[i] = []int{arr[left-1], arr[left]}
		}
	}
	return ans
}

func main() {
	{
		r := &TreeNode{Val: 6}
		r.Left = &TreeNode{Val: 2}
		r.Right = &TreeNode{Val: 13}
		r.Left.Left = &TreeNode{Val: 1}
		r.Left.Right = &TreeNode{Val: 4}
		r.Right.Left = &TreeNode{Val: 9}
		r.Right.Right = &TreeNode{Val: 15}
		r.Right.Right.Left = &TreeNode{Val: 14}
		fmt.Println(closestNodes(r, []int{2, 5, 16}), [][]int{{2, 2}, {4, 6}, {15, -1}})
	}
	{
		r := &TreeNode{Val: 4}
		r.Right = &TreeNode{Val: 9}
		fmt.Println(closestNodes(r, []int{3, 4, 5, 9, 10}), [][]int{{-1, 4}, {4, 4}, {4, 9}, {9, 9}, {9, -1}})
	}
	{
		r := &TreeNode{Val: 4}
		fmt.Println(closestNodes(r, []int{3, 4, 5}), [][]int{{-1, 4}, {4, 4}, {4, -1}})
	}
}

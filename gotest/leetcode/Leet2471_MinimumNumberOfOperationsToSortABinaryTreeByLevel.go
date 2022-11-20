package main

import (
	"container/list"
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-number-of-operations-to-sort-a-binary-tree-by-level/

// You are given the root of a binary tree with unique values.
// In one operation, you can choose any two nodes at the same level and swap their values.
// Return the minimum number of operations needed to make the values at each level sorted in
// a strictly increasing order.
// The level of a node is the number of edges along the path between it and the root node.
// Example 1:
//   Input: root = [1,4,3,7,6,8,5,null,null,null,null,9,null,10]
//   Output: 3
//   Explanation:
//     - Swap 4 and 3. The 2ⁿᵈ level becomes [3,4].
//     - Swap 7 and 5. The 3ʳᵈ level becomes [5,6,8,7].
//     - Swap 8 and 7. The 3ʳᵈ level becomes [5,6,7,8].
//     We used 3 operations so return 3.
//     It can be proven that 3 is the minimum number of operations needed.
// Example 2:
//   Input: root = [1,3,2,7,6,5,4]
//   Output: 3
//   Explanation:
//     - Swap 3 and 2. The 2ⁿᵈ level becomes [2,3].
//     - Swap 7 and 4. The 3ʳᵈ level becomes [4,6,5,7].
//     - Swap 6 and 5. The 3ʳᵈ level becomes [4,5,6,7].
//     We used 3 operations so return 3.
//     It can be proven that 3 is the minimum number of operations needed.
// Example 3:
//   Input: root = [1,2,3,4,5,6]
//   Output: 0
//   Explanation: Each level is already sorted in increasing order so return 0.
// Constraints:
//   The number of nodes in the tree is in the range [1, 10⁵].
//   1 <= Node.val <= 10⁵
//   All the values of the tree are unique.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) int {
	l := list.New()
	l.PushBack(root)
	count := 0
	for l.Len() > 0 {
		ll := l.Len()
		origin := make([]int, 0, ll)
		sorted := make([]int, ll)
		rev := make(map[int]int)
		for i := 0; i < ll; i++ {
			cur := l.Remove(l.Front()).(*TreeNode)
			if cur.Left != nil {
				l.PushBack(cur.Left)
			}
			if cur.Right != nil {
				l.PushBack(cur.Right)
			}
			origin = append(origin, cur.Val)
			rev[cur.Val] = i
		}
		copy(sorted, origin)
		sort.Ints(sorted)
		// for example
		// sorted is:  5 6 7 8
		// for origin: 7 6 8 5
		// sorted[0]=5 in wrong place, swap: 5 6 8 7
		// sorted[1]=6 in right place
		// sorted[2]=7 in wrong place, swap: 5 6 7 8
		for i := 0; i < len(origin); i++ {
			if sorted[i] != origin[i] {
				ai, bi := rev[sorted[i]], rev[origin[i]]
				origin[ai], origin[bi] = origin[bi], origin[ai]
				rev[origin[ai]] = ai
				rev[origin[bi]] = bi
				count++
			}
		}
	}
	return count
}

func main() {
	{
		r := &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 4}
		r.Right = &TreeNode{Val: 3}
		r.Left.Left = &TreeNode{Val: 7}
		r.Left.Right = &TreeNode{Val: 6}
		r.Right.Left = &TreeNode{Val: 8}
		r.Right.Right = &TreeNode{Val: 5}
		r.Right.Left.Left = &TreeNode{Val: 9}
		r.Right.Right.Right = &TreeNode{Val: 10}
		fmt.Println(minimumOperations(r), 3)
	}
	{
		r := &TreeNode{Val: 1}
		r.Left = &TreeNode{Val: 3}
		r.Right = &TreeNode{Val: 2}
		r.Left.Left = &TreeNode{Val: 7}
		r.Left.Right = &TreeNode{Val: 6}
		r.Right.Left = &TreeNode{Val: 5}
		r.Right.Right = &TreeNode{Val: 4}
		fmt.Println(minimumOperations(r), 3)
	}
}

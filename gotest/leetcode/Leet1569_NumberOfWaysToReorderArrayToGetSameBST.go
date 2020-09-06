package main

import "fmt"

// https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/

// Given an array nums that represents a permutation of integers from 1 to n. We are going to
// construct a binary search tree (BST) by inserting the elements of nums in order into an
// initially empty BST. Find the number of different ways to reorder nums so that the constructed BST
// is identical to that formed from the original array nums.
// For example, given nums = [2,1,3], we will have 2 as the root, 1 as a left child, and 3 as a right child.
// The array [2,3,1] also yields the same BST but [3,2,1] yields a different BST.
// Return the number of ways to reorder nums such that the BST formed is identical to the
// original BST formed from nums.
// Since the answer may be very large, return it modulo 10^9 + 7.
// Example 1:
//           2
//          / \
//         1  3
//   Input: nums = [2,1,3]
//   Output: 1
//   Explanation: We can reorder nums to be [2,3,1] which will yield the same BST.
//     There are no other ways to reorder nums which will yield the same BST.
// Example 2:
//           3
//          / \
//         1  4
//          \  \
//          2  5
//   Input: nums = [3,4,5,1,2]
//   Output: 5
//   Explanation: The following 5 arrays will yield the same BST:
//     [3,1,2,4,5]
//     [3,1,4,2,5]
//     [3,1,4,5,2]
//     [3,4,1,2,5]
//     [3,4,1,5,2]
// Example 3:
//   Input: nums = [1,2,3]
//   Output: 0
//   Explanation: There are no other orderings of nums that will yield the same BST.
// Example 4:
//            3
//          /   \
//         1     5
//          \   / \
//          2  4  6
//   Input: nums = [3,1,2,5,4,6]
//   Output: 19
// Example 5:
//   Input: nums = [9,4,2,1,3,6,5,7,8,14,11,10,12,13,16,15,17,18]
//   Output: 216212978
//   Explanation: The number of ways to reorder nums to get the same BST is 3216212999.
//     Taking this number modulo 10^9 + 7 gives 216212978.
// Constraints:
//   1 <= nums.length <= 1000
//   1 <= nums[i] <= nums.length
//   All integers in nums are distinct.

type Node struct {
	val int 
	left *Node
	right *Node
}

// for nums=[3,1,2,5,4,6], 3 is the first node, so we know that [1,2] is in left sub-tree,
// [5,4,6] is in right sub-tree. [1,2] and [5,4,6] can be combined arbitrarily, 
// if we can keep the relative order of each part, e.g. [1,5,2,4,6].
// so we have C(5,2)=10 ways to do that. and [1,2] has 1 way, [5,4,6] has 2 ways. 
// so the total ways is 10*2=20.
// to compute large combination, we use YangHui triangle.
func numOfWays(nums []int) int {
	n := len(nums)
    if n < 3 {
		return 0
	}
	mod := int(1e9+7)

	// make a yanghui triangle
	tri := make([][]int, n)
	for i := range tri {
		tri[i] = make([]int, i+1)
		tri[i][0] = 1
		for j:=1; j<i; j++ {
			tri[i][j] = (tri[i-1][j-1] + tri[i-1][j]) % mod
		}
		tri[i][i] = 1
		//fmt.Println(tri[i])
	}

	root := &Node{val: 0}       // make a BST
	ans := 1
	for _, v := range nums {    // add each number to BST
		min, max := 1, n
		p, t := root, root
		for {       // add this number to BST
			p = t
			if v < t.val {
				t = p.left
				max = p.val-1
				// current interval [min, max]
				// insert new node will part it to [min, v-1], [v+1, max] two parts 
				// the interval is the union of two parts with respect to the relative order in each parts
				// so we have Combination(max-min, v-min) ways.
				if t == nil {                
					p.left = &Node{val: v}  
					break
				}
			} else {
				min = p.val+1
				t = p.right
				if t == nil {
					p.right = &Node{val: v}
					break					
				}
			}
		} 
		ans = (ans * tri[max-min][v-min]) % mod
	}
	return ans-1
}

func main() {
	for _, v := range []struct{nums []int; ans int} {
		{[]int{2,1,3}, 1},
		{[]int{3,4,5,1,2}, 5},
		{[]int{1,2,3}, 0},
		{[]int{3,1,2,5,4,6}, 19},
		{[]int{9,4,2,1,3,6,5,7,8,14,11,10,12,13,16,15,17,18}, 216212978},
		{[]int{7,42,40,12,21,22,13,23,28,38,46,32,30,5,45,9,36,33,1,15,8,3,43,41,20,19,10,29,26,34,6,31,17,4,39,37,14,2,35,11,24,16,44,18,27,25}, 252654363},
		{[]int{13,1,10,4,16,21,29,23,22,31,15,19,2,17,5,12,7,32,24,8,9,30,14,20,18,27,26,11,6,3,28,25}, 432023023},
	} {
		fmt.Println(numOfWays(v.nums), v.ans)
	}
}
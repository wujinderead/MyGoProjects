package main

import (
	"sort"
	"fmt"
)

// https://leetcode.com/problems/maximum-length-of-pair-chain/

// You are given n pairs of numbers. In every pair, the first number is always
// smaller than the second number.
// Now, we define a pair (c, d) can follow another pair (a, b) if and only if
// b < c. Chain of pairs can be formed in this fashion.
// Given a set of pairs, find the length longest chain which can be formed.
// You needn't use up all the given pairs. You can select pairs in any order.
// Example 1:
//   Input: [[1,2], [2,3], [3,4]]
//   Output: 2
//   Explanation: The longest chain is [1,2] -> [3,4]
// Note:
//   The number of given pairs will be in the range [1, 1000].

func findLongestChain(p [][]int) int {
	// sort the pairs by second value, and form a longest chain as possible
	sort.Sort(pairs(p))
	min := -0x7fffffff
	count := 0
	for i:=0; i<len(p); i++ {
		if p[i][0] > min {
			count++
			min = p[i][1]
		}
	}
	return count
}

type pairs [][]int

func (p pairs) Less(i, j int) bool {
	return p[i][1] < p[j][1]
}

func (p pairs) Len() int {
	return len(p)
}

func (p pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	fmt.Println(findLongestChain([][]int{{1,2}, {2,3}, {3,4}}))
}
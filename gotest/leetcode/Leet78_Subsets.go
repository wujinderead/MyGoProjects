package main

import "fmt"

// https://leetcode.com/problems/subsets/

// Given a set of distinct integers, nums, return all possible subsets (the power set).
// Note: The solution set must not contain duplicate subsets.
// Example:
//   Input: nums = [1,2,3]
//   Output:
//     [
//       [3],
//       [1],
//       [2],
//       [1,2,3],
//       [1,3],
//       [2,3],
//       [1,2],
//       []
//     ]

// if there are n elements, the bit set for the subsets happens to be from
// 0 (0...0) to pow(2,n)-1 (1...1)
func subsets(nums []int) [][]int {
	numsets := 1 << uint(len(nums))
	subsets := make([][]int, numsets)
	subsets[0] = []int{}
	for i := 1; i < numsets; i++ {
		// count bits set for at most 14-bit values in v:
		c := (i * 0x200040008001 & 0x111111111111111) % 0xf
		fmt.Println(i, c)
		subsets[i] = make([]int, c)
		mask := 1
		ind := 0
		for j := 0; j < len(nums); j++ {
			if i&mask == mask {
				subsets[i][ind] = nums[j]
				ind++
			}
			mask <<= 1
		}
	}
	return subsets
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{2, 3, 4, 5}))
}

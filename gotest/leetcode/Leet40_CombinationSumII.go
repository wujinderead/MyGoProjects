package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/combination-sum-ii/

// Given a collection of candidate numbers (candidates) and a target number (target),
// find all unique combinations in candidates where the candidate numbers sums
// to target. Each number in candidates may only be used once in the combination.
// Note:
//   All numbers (including target) will be positive integers.
//   The solution set must not contain duplicate combinations.
// Example 1:
//   Input: candidates = [10,1,2,7,6,1,5], target = 8,
//   A solution set is:
//    [
//     [1, 7],
//     [1, 2, 5],
//     [2, 6],
//     [1, 1, 6]
//    ]
// Example 2:
//   Input: candidates = [2,5,2,1,2], target = 5,
//   A solution set is:
//    [
//     [1,2,2],
//     [5]
//    ]

func combinationSum2(candidates []int, target int) [][]int {
	sort.Sort(sort.IntSlice(candidates))
	ans := make([][]int, 0)
	buf := make([]int, 0, 10)
	helper(candidates, buf, 0, target, &ans)
	return ans
}

func helper(nums, buf []int, ni int, target int, ans *[][]int) {
	if target==0 {
		tmp := make([]int, len(buf))
		copy(tmp, buf)
		*ans = append(*ans, tmp)
		return
	}
	if target<0 {
		return
	}
	for i:=ni; i<len(nums); i++ {       // when i=ni+1 means we don't use nums[ni] bu use nums[ni+1]
		if i>ni && nums[i] == nums[i-1] {   // i>ni means we didn't include nums[ni...i-1]
			continue
		}
		buf = append(buf, nums[i])
		helper(nums, buf, i+1, target-nums[i], ans)
		buf = buf[:len(buf)-1]
	}
}

func main() {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5}, 8))
	fmt.Println(combinationSum2([]int{2,5,2,1,2}, 5))
	fmt.Println(combinationSum2([]int{3,1,3,5,1,1}, 8))
}
package main

import "fmt"

// https://leetcode.com/problems/combination-sum/

// Given a set of candidate numbers (candidates) (without duplicates) and a target
// number (target), find all unique combinations in candidates where the candidate
// numbers sums to target.
// The same repeated number may be chosen from candidates unlimited number of times.
// Note:
//   All numbers (including target) will be positive integers.
//   The solution set must not contain duplicate combinations.
// Example 1:
//   Input: candidates = [2,3,6,7], target = 7,
//   A solution set is:
//    [
//     [7],
//     [2,2,3]
//    ]
// Example 2:
//   Input: candidates = [2,3,5], target = 8,
//   A solution set is:
//    [
//     [2,2,2,2],
//     [2,3,3],
//     [3,5]
//    ]

func combinationSum(candidates []int, target int) [][]int {
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
	if ni==len(nums) {
		return
	}
	if nums[ni]<=target {
		buf = append(buf, nums[ni])
		helper(nums, buf, ni, target-nums[ni], ans)   // use nums[i]
		buf = buf[:len(buf)-1]
	}
	helper(nums, buf, ni+1, target, ans)    // do not use nums[i]
}

func main() {
	fmt.Println(combinationSum([]int{2,3,6,7}, 7))
	fmt.Println(combinationSum([]int{2,3,5}, 8))
}
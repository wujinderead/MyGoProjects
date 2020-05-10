package main

import "fmt"

// https://leetcode.com/problems/increasing-subsequences/

// Given an integer array, your task is to find all the different possible increasing
// subsequences of the given array, and the length of an increasing subsequence
// should be at least 2.
// Example:
//   Input: [4, 6, 7, 7]
//   Output: [[4, 6], [4, 7], [4, 6, 7], [4, 6, 7, 7], [6, 7], [6, 7, 7], [7,7], [4,7,7]]
// Note:
//   The length of the given array will not exceed 15.
//   The range of integer in the given array is [-100,100].
//   The given array may contain duplicates, and two equal integers should also be
//   considered as a special case of increasing sequence.

func findSubsequences(nums []int) [][]int {
    seqs := make([][]int, 0, len(nums))
    buf := make([]int, 0, len(nums))
    outer: for i:=0; i<len(nums)-1; i++ {
		for j:=i-1; j>=0; j-- {
			if nums[j]==nums[i] {
				continue outer // exclude duplicate
			}
		}
		findSub(i, buf, nums, &seqs)
	}
    return seqs
}

func findSub(s int, buf, nums[] int, seqs *[][]int) {
	buf = append(buf, nums[s])
	if len(buf)>1 {
		tmp := make([]int, len(buf))
		copy(tmp, buf)
		*seqs = append(*seqs, tmp)
	}
	outer: for i:=s+1; i<len(nums); i++ {
		if nums[s]<=nums[i] {
			for j:=i-1; j>s; j-- {
				if nums[j] == nums[i] {   // exclude duplicate
					continue outer
				}
			}
			findSub(i, buf, nums, seqs)
		}
	}
	buf = buf[:len(buf)-1]
}

func main() {
	fmt.Println(findSubsequences([]int{1,2}))
	fmt.Println(findSubsequences([]int{2,1}))
	fmt.Println(findSubsequences([]int{1,2,2}))
	fmt.Println(findSubsequences([]int{1,2,2,2}))
	fmt.Println(findSubsequences([]int{4,6,7,7}))
	a := findSubsequences([]int{1,2,3,4,5,6,7,8,9,10,1,1,1,1,1})
	fmt.Println(len(a), a)
}
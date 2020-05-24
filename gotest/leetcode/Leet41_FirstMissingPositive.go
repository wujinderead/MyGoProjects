package main

import (
	"fmt"
	"math/rand"
)

// https://leetcode.com/problems/first-missing-positive/

// Given an unsorted integer array, find the smallest missing positive integer.
// Example 1:
//  Input: [1,2,0]
//  Output: 3
// Example 2:
//   Input: [3,4,-1,1]
//   Output: 2
// Example 3:
//   Input: [7,8,9,11,12]
//   Output: 1
// Note:
//   Your algorithm should run in O(n) time and uses constant extra space.

func firstMissingPositive(nums []int) int {
	// put the number in the right place
    for i:=range nums {
    	t := nums[i]
    	for t>=1 && t<=len(nums) && nums[t-1] != nums[i]{
			nums[i], nums[t-1] = nums[t-1], nums[i]
			t = nums[i]
		}
	}
	// find the first number that is not in the place
	for i:=range nums {
		if nums[i] != i+1 {
			return i+1
		}
	}
	return len(nums)+1
}

func main() {
	fmt.Println(firstMissingPositive([]int{1,2,0}))
	fmt.Println(firstMissingPositive([]int{1,1}))
	fmt.Println(firstMissingPositive([]int{1,2,3,4}))
	fmt.Println(firstMissingPositive([]int{3,4,-1,1}))
	fmt.Println(firstMissingPositive([]int{7,8,9,11,12}))
	for i:=0; i<5; i++ {
		a := rand.Perm(20)
		fmt.Println(a[:17], a[17:])
		fmt.Println(firstMissingPositive(a[:17]))
	}
}
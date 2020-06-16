package main

import (
    "fmt"
)

// https://leetcode.com/problems/subarray-sum-equals-k/

// Given an array of integers and an integer k, you need to find the total number of 
// continuous subarrays whose sum equals to k.
// Example 1:
//   Input:nums = [1,1,1], k = 2
//   Output: 2
// Constraints:
//   The length of the array is in range [1, 20,000].
//   The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].

func subarraySum(nums []int, k int) int {
    // compute prefix sums and save their count in a map
    mapp := make(map[int]int)
    cursum := 0
   	count := 0
   	mapp[0] = 1
   	for i:=0; i<len(nums); i++ {
   		cursum += nums[i]
   		prev := mapp[cursum-k]   // how many prefix has sum = cursum-k
   		// if there are some x that make sum(nums[0...x])==cursum-k, and sum(nums[0...i])=cursum
   		// then there is sum(nums[x+1...i])=k; the count of such x is already in map.
   		count += prev   
   		mapp[cursum] = mapp[cursum]+1   // add count
   	}
   	return count
}

func main() {
	fmt.Println(subarraySum([]int{1,1,1}, 2))
	fmt.Println(subarraySum([]int{0,0,0}, 0))
	fmt.Println(subarraySum([]int{6,3,-3,5,-5,6}, 6))
	fmt.Println(subarraySum([]int{3,-3,5,-5,0}, 0))
}
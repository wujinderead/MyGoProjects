package main

import (
    "fmt"
)

// https://leetcode.com/problems/continuous-subarray-sum/

// Given a list of non-negative numbers and a target integer k, write a function to check 
// if the array has a continuous subarray of size at least 2 that sums up to a multiple of k, 
// that is, sums up to n*k where n is also an integer.
// Example 1:
//   Input: [23, 2, 4, 6, 7],  k=6
//   Output: True
//   Explanation: Because [2, 4] is a continuous subarray of size 2 and sums up to 6.
// Example 2:  
//   Input: [23, 2, 6, 4, 7],  k=6
//   Output: True
//   Explanation: Because [23, 2, 6, 4, 7] is an continuous subarray of size 5 and sums up to 42.
// Constraints:
//   The length of the array won't exceed 10,000.
//   You may assume the sum of all the numbers is in the range of a signed 32-bit integer.

func checkSubarraySum(nums []int, k int) bool {
    if k<0 {
    	k = -k
    }
    // special case for k==0
    if k==0 {
    	for i:=1; i<len(nums); i++ {
    		if nums[i]==0 && nums[i-1]==0 {
    			return true
    		}
    	}
    	return false
    }

    // get the prefix sum and mod k, if we got two same remains, we know the difference is what we want.
	res := make(map[int]int)
    prefix := make([]int, len(nums))
	prefix[0] = nums[0]
    res[nums[0]%k] = 0 
    for i:=1; i<len(nums); i++ {
    	prefix[i] = prefix[i-1]+nums[i]
    	if _, ok := res[prefix[i]%k]; !ok {
    		res[prefix[i]%k] = i
    	}
    	if i-res[prefix[i]%k]>1 {  // subarray length >= 2 
    		return true
    	}
    	if prefix[i]%k==0 {
    		return true
    	}
    }
    return false
}

func main() {
	fmt.Println(checkSubarraySum([]int{23, 2, 4, 6, 7}, 6))
	fmt.Println(checkSubarraySum([]int{3,6,7}, 6))
	fmt.Println(checkSubarraySum([]int{3,6,9}, -6))
	fmt.Println(checkSubarraySum([]int{0}, 0))
	fmt.Println(checkSubarraySum([]int{0,0}, 0))
	fmt.Println(checkSubarraySum([]int{1000000000}, 1000000000))
	fmt.Println(checkSubarraySum([]int{1,1000000000-1}, 1000000000))
}
package main

import (
    "fmt"
)

// https://leetcode.com/problems/subarray-product-less-than-k/

// Your are given an array of positive integers nums. Count and print the number of (contiguous) subarrays 
// where the product of all the elements in the subarray is less than k.
// Example 1:
//   Input: nums = [10, 5, 2, 6], k = 100
//   Output: 8
//   Explanation: The 8 subarrays that have product less than 100 are: [10], [5], [2], [6], [10, 5], [5, 2], [2, 6], [5, 2, 6].
//     Note that [10, 5, 2] is not included as the product of 100 is not strictly less than k.
// Note:
//   0 < nums.length <= 50000.
//   0 < nums[i] < 1000.
//   0 <= k < 10^6.

func numSubarrayProductLessThanK(nums []int, k int) int {
	if k<=1 {
		return 0
	}
    s := 0
    product := 1
    count := 0
    for i:=0; i<len(nums); i++ {
    	product *= nums[i]
    	for product>=k {           // let product(nums[s...i])<k
    		product /= nums[s]
    		s++
    	}
    	// if product(nums[s...i])<k, for s<=x<=i, product(nums[x...i])<k, add i-s+1 to count
        // if s=i+1, that's ok since i-s+1=0
    	count += i-s+1
    }
    return count
}

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 6))	
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 0))		
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 1))		
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 2))		
	fmt.Println(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 3))		
}
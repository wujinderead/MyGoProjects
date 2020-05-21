package main

import "fmt"

// https://leetcode.com/problems/maximum-product-subarray/

// Given an integer array nums, find the contiguous subarray within an array
// (containing at least one number) which has the largest product.
// Example 1:
//   Input: [2,3,-2,4]
//   Output: 6
//   Explanation: [2,3] has the largest product 6.
// Example 2:
//   Input: [-2,0,-1]
//   Output: 0
//   Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

// IMPROVEMENT: a much easier process:
// https://leetcode.com/problems/maximum-product-subarray/discuss/48230/Possibly-simplest-solution-with-O(n)-time-complexity
func maxProduct(nums []int) int {
	if len(nums)==0 {
		return 0
	}
    max := nums[0]
    allmul := 1
    firstneg := 1
    alln := 0
    for i := range nums {
    	if nums[i]==0 {
    		if max<0 {
    			max = 0
			}
			if allmul/firstneg>max && alln>1 {
				max = allmul/firstneg
			}
			allmul = 1
			alln = 0
			firstneg = 1
			continue
		}
		allmul *= nums[i]
		alln++
		if firstneg>0 {
			firstneg *= nums[i]
		}
		if allmul>max {
			max = allmul
		}
	}
	if allmul/firstneg>max && alln>1{
		max = allmul/firstneg
	}
    return max
}

func main() {
	fmt.Println(maxProduct([]int{2,3,-2,4}))
	fmt.Println(maxProduct([]int{2,3,-2,4,-3,-1,2,0,7,-2,100}))
	fmt.Println(maxProduct([]int{-2,0,-1}))
	fmt.Println(maxProduct([]int{-2,0}))
	fmt.Println(maxProduct([]int{2,0}))
	fmt.Println(maxProduct([]int{0,-1}))
	fmt.Println(maxProduct([]int{-1}))
	fmt.Println(maxProduct([]int{2,-5,-2,-4,3}))
}
package main

import (
    "fmt"
)

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand. 
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]). Find the minimum element. 
// You may assume no duplicate exists in the array. 
// Example 1: 
//   Input: [3,4,5,1,2] 
//   Output: 1
// Example 2: 
//   Input: [4,5,6,7,0,1,2]
//   Output: 0

func findMin(nums []int) int {
    l, r := 0, len(nums)-1
    for l<r {
    	if nums[l]<nums[r] {      // the min is in the wrong part
    		return nums[l]        // if not wrong, we found it
    	}
    	mid := (l+r)/2
    	if nums[l]<=nums[mid] {
    		l = mid+1
    	} else {
    		r = mid
    	}
    }
    return nums[l]
}

func main() {
	fmt.Println(findMin([]int{1,2,3,4,5}))
	fmt.Println(findMin([]int{2,3,4,5,1}))
	fmt.Println(findMin([]int{3,4,5,1,2}))
	fmt.Println(findMin([]int{4,5,1,2,3}))
	fmt.Println(findMin([]int{5,1,2,3,4}))
	fmt.Println(findMin([]int{0,1,2}))
	fmt.Println(findMin([]int{1,2,0}))
	fmt.Println(findMin([]int{2,0,1}))
	fmt.Println(findMin([]int{1,0}))
	fmt.Println(findMin([]int{0,1}))
	fmt.Println(findMin([]int{0}))
}
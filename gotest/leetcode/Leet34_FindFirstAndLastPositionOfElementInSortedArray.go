package main

import (
    "fmt"
)

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

// Given an array of integers nums sorted in ascending order, find the starting and ending 
// position of a given target value. Your algorithm's runtime complexity must be in the order of O(log n).
// If the target is not found in the array, return [-1, -1].
// Example 1:
//   Input: nums = [5,7,7,8,8,10], target = 8
//   Output: [3,4]
// Example 2:
//   Input: nums = [5,7,7,8,8,10], target = 6
//   Output: [-1,-1]

func searchRange(nums []int, target int) []int {
	if len(nums)==0 {  // special case: empty array
		return []int{-1,-1}
	}
    l, r := 0, len(nums)-1

    // this binary search find the first l that nums[l]>=target 
    for l<r {
    	mid := (l+r)/2
    	if nums[mid]>=target {
    		r = mid
    	} else {
    		l = mid+1
    	}
    }
    if nums[l] != target {   // target not exist
    	return []int{-1, -1}
    }
    left := l

    l, r = 0, len(nums)-1
    // this binary search find the first l that nums[l]>target 
    for l<r {
    	mid := (l+r)/2
    	if nums[mid]<=target {
    		l = mid+1
    	} else {
    		r = mid
    	}
    }
    if nums[l] != target {
    	l = l-1
    }
    return []int{left, l}
}

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
	fmt.Println(searchRange([]int{5,7,7,8,8,10,10,10}, 10))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 11))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 3))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 7))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 9))
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 5))
	fmt.Println(searchRange([]int{2}, 1))
	fmt.Println(searchRange([]int{2}, 2))
	fmt.Println(searchRange([]int{2}, 3))
	fmt.Println(searchRange([]int{1,2}, 1))
	fmt.Println(searchRange([]int{1,2}, 2))
	fmt.Println(searchRange([]int{2,2}, 2))
}
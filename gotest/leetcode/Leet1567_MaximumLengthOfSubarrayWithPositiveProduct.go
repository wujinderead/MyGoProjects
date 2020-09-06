package main

import "fmt"

// https://leetcode.com/problems/maximum-length-of-subarray-with-positive-product/

// Given an array of integers nums, find the maximum length of a subarray where the product of
// all its elements is positive. A subarray of an array is a consecutive sequence of zero or
// more values taken out of that array.
// Return the maximum length of a subarray with positive product.
// Example 1:
//   Input: nums = [1,-2,-3,4]
//   Output: 4
//   Explanation: The array nums already has a positive product of 24.
// Example 2:
//   Input: nums = [0,1,-2,-3,-4]
//   Output: 3
//   Explanation: The longest subarray with positive product is [1,-2,-3] which has a product of 6.
//     Notice that we cannot include 0 in the subarray since that'll make the product 0 which
//     is not positive.
// Example 3:
//   Input: nums = [-1,-2,-3,0,1]
//   Output: 2
//   Explanation: The longest subarray with positive product is [-1,-2] or [-2,-3].
// Example 4:
//   Input: nums = [-1,2]
//   Output: 1
// Example 5:
//   Input: nums = [1,2,3,5,-6,4,0,10]
//   Output: 4
// Constraints:
//   1 <= nums.length <= 10^5
//   -10^9 <= nums[i] <= 10^9

// in an interval between two 0's, e.g., [0,p,N,p,N,p,N,p,p,0], 
// if we have even number of negative values in it, then the answer the interval length.
// if odd number of negative values, we compare the first and the last negtive values,
// i.e., the candidates are: the left of last N, or the right of first N.
//    [0,p,N,p,N,p,N,p,p,0]
//       ---------
//           -----------
func getMaxLen(nums []int) int {
	maxlen := 0
	s := -1
	fn, ln, negs := -1, -1, 0
	for i:=0; i<=len(nums); i++ {
		// add zero at end to trigger final compute
		v := 0
		if i<len(nums) {
			v = nums[i]
		}

		// find an interval
		if v==0 {   
			if i-s > 1 {   // non-empty interval
				if negs%2==0 {   // even negs, whole interval is ok
					maxlen = max(maxlen, i-s-1)
				} else {
					maxlen = max(maxlen, max(ln-s-1, i-fn-1))
				}
			}
			fn, ln, s = i, i, i
			negs = 0
		}

		// find a value
		if v < 0 {
			negs++
			if fn==s {
				fn = i
			}
			ln = i
		}
		// if v > 0, do nothing
	}
    return maxlen
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct{nums []int; ans int} {
		{[]int{1,-2,-3,4}, 4},
		{[]int{0,1,-2,-3,-4}, 3},
		{[]int{-1,-2,-3,0,1}, 2},
		{[]int{-1,2}, 1},
		{[]int{1,2,3,5,-6,4,0,10}, 4},
		{[]int{-2}, 0},
		{[]int{0,-2}, 0},
		{[]int{-2,0}, 0},
		{[]int{0,0,0,-2}, 0},
		{[]int{1,0}, 1},
		{[]int{1}, 1},
	} {
		fmt.Println(getMaxLen(v.nums), v.ans)
	}
}
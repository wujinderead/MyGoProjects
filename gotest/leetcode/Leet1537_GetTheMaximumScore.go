package main

import (
	"fmt"
)

// https://leetcode.com/problems/get-the-maximum-score/

// You are given two sorted arrays of distinct integers nums1 and nums2.
// A valid path is defined as follows:
//   Choose array nums1 or nums2 to traverse (from index-0).
//   Traverse the current array from left to right.
//   If you are reading any value that is present in nums1 and nums2 you are allowed to change 
//     your path to the other array. (Only one repeated value is considered in the valid path).
// Score is defined as the sum of uniques values in a valid path.
// Return the maximum score you can obtain of all possible valid paths.
// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: nums1 = [2,4,5,8,10], nums2 = [4,6,8,9]
//   Output: 30
//   Explanation: Valid paths:
//     [2,4,5,8,10], [2,4,5,8,9], [2,4,6,8,9], [2,4,6,8,10],  (starting from nums1)
//     [4,6,8,9], [4,5,8,10], [4,5,8,9], [4,6,8,10]    (starting from nums2)
//     The maximum is obtained with the path in green [2,4,6,8,10].
// Example 2:
//   Input: nums1 = [1,3,5,7,9], nums2 = [3,5,100]
//   Output: 109
//   Explanation: Maximum sum is obtained with the path [1,3,5,100].
// Example 3:
//   Input: nums1 = [1,2,3,4,5], nums2 = [6,7,8,9,10]
//   Output: 40
//   Explanation: There are no common elements between nums1 and nums2.
//     Maximum sum is obtained with the path [6,7,8,9,10].
// Example 4:
//   Input: nums1 = [1,4,5,8,9,11,19], nums2 = [2,3,4,11,12]
//   Output: 61
// Constraints:
//   1 <= nums1.length <= 10^5
//   1 <= nums2.length <= 10^5
//   1 <= nums1[i], nums2[i] <= 10^7
//   nums1 and nums2 are strictly increasing.

// e.g., for two arrays [1,4,5,8,9,11,19] and [2,3,4,11,12], align equal numbers.
//   1, 4, 5,8,9, 11, 19
// 2,3, 4         11, 12
// for unaligned numbers, compare and add larger to total sum.
func maxSum(nums1 []int, nums2 []int) int {
	prime := int(1e9+7)
	var i, j, sumi, sumj, allsum int
	for i<=len(nums1) && j<=len(nums2) {   
		ai, aj := prime, prime    // add a larger number at the end of arrays to trigger final compute 
		if i < len(nums1) {
			ai = nums1[i]
		}
		if j < len(nums2) {
			aj = nums2[j]
		}

		// use two pointers to align equal numberss
		if ai < aj {
			sumi += ai
			i++
		} else if ai > aj {
			sumj += aj
			j++
		} else {   // equal
			if sumi > sumj {
				allsum += sumi+ai
			} else {
				allsum += sumj+aj
			}
			sumi, sumj = 0, 0
			i++
			j++
		}
	}
	return allsum%prime
}

func main() {
	for _, v := range []struct{a1, a2 []int; ans int} {
		{[]int{2,4,5,8,10}, []int{4,6,8,9}, 30},
		{[]int{1,3,5,7,9}, []int{3,5,100}, 109},
		{[]int{1,2,3,4,5}, []int{6,7,8,9,10}, 40},
		{[]int{1,4,5,8,9,11,19}, []int{2,3,4,11,12}, 61},
	} {
		fmt.Println(maxSum(v.a1, v.a2), v.ans)
	}
}
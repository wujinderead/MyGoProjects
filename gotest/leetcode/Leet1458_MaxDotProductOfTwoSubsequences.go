package main

import "fmt"

// https://leetcode.com/problems/max-dot-product-of-two-subsequences/

// Given two arrays nums1 and nums2.
// Return the maximum dot product between non-empty subsequences of nums1 and nums2
// with the same length.
// A subsequence of a array is a new array which is formed from the original array
// by deleting some (can be none) of the characters without disturbing the relative
// positions of the remaining characters. (ie, [2,3,5] is a subsequence of [1,2,3,4,5]
// while [1,5,3] is not).
// Example 1:
//   Input: nums1 = [2,1,-2,5], nums2 = [3,0,-6]
//   Output: 18
//   Explanation: Take subsequence [2,-2] from nums1 and subsequence [3,-6] from nums2.
//     Their dot product is (2*3 + (-2)*(-6)) = 18.
// Example 2:
//   Input: nums1 = [3,-2], nums2 = [2,-6,7]
//   Output: 21
//   Explanation: Take subsequence [3] from nums1 and subsequence [7] from nums2.
//     Their dot product is (3*7) = 21.
// Example 3:
//   Input: nums1 = [-1,-1], nums2 = [1,1]
//   Output: -1
//   Explanation: Take subsequence [-1] from nums1 and subsequence [1] from nums2.
//     Their dot product is -1.
// Constraints:
//    1 <= nums1.length, nums2.length <= 500
//    -1000 <= nums1[i], nums2[i] <= 1000

func maxDotProduct(nums1 []int, nums2 []int) int {
    // let dp(i, j) be the maximal dot product of nums1[0...i] and nums2[0...j], then 
    // dp(i+1, j+1) has candidates: 
    //    dp(i, j+1), 
    //    dp(i+1, j), 
    //    nums1[i+1]*nums[j+1]_(IF>0) + dp(i,j)_(IF>0)
    old, new := make([]int, len(nums2)), make([]int, len(nums2))
    old[0] = nums1[0]*nums2[0]
    for i:=1; i<len(nums2); i++ {
    	old[i] = old[i-1]
    	if nums1[0]*nums2[i] > old[i-1] {
    		old[i] = nums1[0]*nums2[i]
    	}
    }
    //fmt.Println(old)
    for i:=1; i<len(nums1); i++ {
    	new[0] = old[0]
    	if nums1[i]*nums2[0] > old[0] {
    		new[0] = nums1[i]*nums2[0]
    	}
    	for j:=1; j<len(nums2); j++ {
    		new[j] = max(old[j], new[j-1])     // dp(i, j-1), dp(i-1, j) 
    		tmp := old[j-1]
    		if nums1[i]*nums2[j]>=0 {
    			tmp = nums1[i]*nums2[j]
    			if old[j-1]>=0 {
    				tmp += old[j-1]
    			}
    		}
    		new[j] = max(new[j], tmp)
    	}
    	//fmt.Println(new)
    	old, new = new, old
    }
    return old[len(nums2)-1]
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxDotProduct([]int{2,1,-2,5}, []int{3,0,-6}))
	fmt.Println(maxDotProduct([]int{3,-2}, []int{2,-6,7}))
	fmt.Println(maxDotProduct([]int{-1,-1}, []int{1,1}))
	fmt.Println(maxDotProduct([]int{-1}, []int{-2,3}))
	fmt.Println(maxDotProduct([]int{1}, []int{-3}))
	fmt.Println(maxDotProduct([]int{-3, 2}, []int{-1}))
	fmt.Println(maxDotProduct([]int{5,-4,-3}, []int{-4,-3,0,-4,2}))
}
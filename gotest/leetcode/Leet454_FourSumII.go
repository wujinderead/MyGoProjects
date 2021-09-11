package main

import "fmt"

// https://leetcode.com/problems/4sum-ii/

// Given four integer arrays nums1, nums2, nums3, and nums4 all of length n, return the
// number of tuples (i, j, k, l) such that:
//   0 <= i, j, k, l < n
//   nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0
// Example 1:
//   Input: nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
//   Output: 2
//   Explanation:
//     The two tuples are:
//     1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
//     2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
// Example 2:
//   Input: nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
//   Output: 1
// Constraints:
//   n == nums1.length
//   n == nums2.length
//   n == nums3.length
//   n == nums4.length
//   1 <= n <= 200
//   -2^28 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2^28

// O(n²) times to get (nums1*nums2) pairs' sums, and find -sum in (nums3*nums4) pairs
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	mapp := make(map[int]int)
	for _, v := range nums1 {
		for _, vv := range nums2 {
			mapp[v+vv] = mapp[v+vv] + 1
		}
	}
	for _, v := range nums3 {
		for _, vv := range nums4 {
			count += mapp[-v-vv]
		}
	}
	return count
}

func main() {
	for _, v := range []struct {
		n1, n2, n3, n4 []int
		ans            int
	}{
		{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
		{[]int{0}, []int{0}, []int{0}, []int{0}, 1},
	} {
		fmt.Println(fourSumCount(v.n1, v.n2, v.n3, v.n4), v.ans)
	}
}

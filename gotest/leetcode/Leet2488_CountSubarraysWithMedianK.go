package main

import "fmt"

// https://leetcode.com/problems/count-subarrays-with-median-k/

// You are given an array nums of size n consisting of distinct integers from 1 to n and a positive
// integer k.
// Return the number of non-empty subarrays in nums that have a median equal to k.
// Note:
//   The median of an array is the middle element after sorting the array in ascending order.
//   If the array is of even length, the median is the left middle element.
// For example, the median of [2,3,1,4] is 2, and the median of [8,4,3,5,1] is 4.
// A subarray is a contiguous part of an array.
// Example 1:
//   Input: nums = [3,2,1,4,5], k = 4
//   Output: 3
//   Explanation: The subarrays that have a median equal to 4 are: [4], [4,5] and [1,4,5].
// Example 2:
//   Input: nums = [2,3,1], k = 3
//   Output: 1
//   Explanation: [3] is the only subarray that has a median equal to 3.
// Constraints:
//   n == nums.length
//   1 <= n <= 10⁵
//   1 <= nums[i], k <= n
//   The integers in nums are distinct.

// for nums, make all number greater than k as 1, all numbers less than k as -1, k as 0
// then the expected subarrays should have subarray_sum of 0 or 1.
func countSubarrays(nums []int, k int) int {
	kind := 0 // index of k
	// normalize nums
	for i := range nums {
		if nums[i] > k {
			nums[i] = 1
		} else if nums[i] < k {
			nums[i] = -1
		} else {
			nums[i] = 0
			kind = i
		}
	}
	// prefixSum-count map
	mapp := make(map[int]int)
	s := 0
	mapp[0] = 1                 // for empty prefix
	for i := 0; i < kind; i++ { // get prefix before kind
		s += nums[i]
		mapp[s] += 1
	}
	count := 0
	for i := kind; i < len(nums); i++ { // start from kind, this ensures that all subarrays contains k
		s += nums[i]
		count += mapp[s] + mapp[s-1]
	}
	return count
}

func main() {
	for _, v := range []struct {
		nums   []int
		k, ans int
	}{
		{[]int{3, 2, 1, 4, 5}, 4, 3},
		{[]int{2, 3, 1}, 3, 1},
		{[]int{5, 2, 4, 3, 6, 1, 7}, 3, 12},
	} {
		fmt.Println(countSubarrays(v.nums, v.k), v.ans)
	}
}

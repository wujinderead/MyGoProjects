package main

import "fmt"

// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/

// You are given an integer array nums and an integer k. Find the maximum subarray sum of all the
// subarrays of nums that meet the following conditions:
//   The length of the subarray is k, and
//   All the elements of the subarray are distinct.
// Return the maximum subarray sum of all the subarrays that meet the conditions.
// If no subarray meets the conditions, return 0.
// A subarray is a contiguous non-empty sequence of elements within an array.
// Example 1:
//   Input: nums = [1,5,4,2,9,9,9], k = 3
//   Output: 15
//   Explanation: The subarrays of nums with length 3 are:
//     - [1,5,4] which meets the requirements and has a sum of 10.
//     - [5,4,2] which meets the requirements and has a sum of 11.
//     - [4,2,9] which meets the requirements and has a sum of 15.
//     - [2,9,9] which does not meet the requirements because the element 9 is repeated.
//     - [9,9,9] which does not meet the requirements because the element 9 is repeated.
//     We return 15 because it is the maximum subarray sum of all the subarrays that meet the conditions
// Example 2:
//   Input: nums = [4,4,4], k = 3
//   Output: 0
//   Explanation: The subarrays of nums with length 3 are:
//     - [4,4,4] which does not meet the requirements because the element 4 is repeated.
//     We return 0 because no subarrays meet the conditions.
// Constraints:
//  1 <= k <= nums.length <= 10⁵
//  1 <= nums[i] <= 10⁵

func maximumSubarraySum(nums []int, k int) int64 {
	max := 0
	sum := 0
	count := make(map[int]int)
	dup := 0 // duplicated value count
	for i := 0; i < k; i++ {
		sum += nums[i]
		count[nums[i]] = count[nums[i]] + 1
		if count[nums[i]] == 2 {
			dup++
		}
	}
	if dup == 0 && sum > max {
		max = sum
	}
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		count[nums[i]] = count[nums[i]] + 1
		if count[nums[i]] == 2 {
			dup++
		}
		count[nums[i-k]] = count[nums[i-k]] - 1
		if count[nums[i-k]] == 1 {
			dup--
		}
		if dup == 0 && sum > max {
			max = sum
		}
	}
	return int64(max)
}

func main() {
	for _, v := range []struct {
		nums []int
		k    int
		ans  int64
	}{
		{[]int{1, 5, 4, 2, 9, 9, 9}, 3, 15},
		{[]int{4, 4, 4}, 3, 0},
		{[]int{11, 11, 7, 2, 9, 4, 17, 1}, 1, 17},
	} {
		fmt.Println(maximumSubarraySum(v.nums, v.k), v.ans)
	}
}

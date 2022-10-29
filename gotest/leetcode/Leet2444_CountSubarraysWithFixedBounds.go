package main

import "fmt"

// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/

// You are given an integer array nums and two integers minK and maxK.
// A fixed-bound subarray of nums is a subarray that satisfies the following conditions:
//   The minimum value in the subarray is equal to minK.
//   The maximum value in the subarray is equal to maxK.
// Return the number of fixed-bound subarrays.
// A subarray is a contiguous part of an array.
// Example 1:
//   Input: nums = [1,3,5,2,7,5], minK = 1, maxK = 5
//   Output: 2
//   Explanation: The fixed-bound subarrays are [1,3,5] and [1,3,5,2].
// Example 2:
//   Input: nums = [1,1,1,1], minK = 1, maxK = 1
//   Output: 10
//   Explanation: Every subarray of nums is a fixed-bound subarray. There are 10 possible subarrays.
// Constraints:
//   2 <= nums.length <= 10⁵
//   1 <= nums[i], minK, maxK <= 10⁶

// only consider the subarrays that has with minK <= values <= maxK
// then for each subarray, start from right, find the closest minK and maxK to a number's right.
// for example: minK=1, maxK=5
//  2,4,1,5,3,5,1,4
//          s   e e
//  s     e e e e e
func countSubarrays(nums []int, minK int, maxK int) int64 {
	count := 0
	s := 0
	for s < len(nums) {
		if nums[s] < minK || nums[s] > maxK {
			s++
			continue
		}
		e := s
		for e+1 < len(nums) && nums[e+1] >= minK && nums[e+1] <= maxK {
			e++
		}
		// found a segment nums[s...e] with minK <= values <= maxK
		minI, maxI := e+1, e+1 // the closest minK and maxK index to nums[i]'s right
		for i := e; i >= s; i-- {
			if nums[i] == minK {
				minI = i
			}
			if nums[i] == maxK {
				maxI = i
			}
			// for subarray start at nums[i], max(minI, maxI) <= x <= e can be the end
			// we need both occurrence of maxK and minK; however when minI or maxI == e+1,
			// e - max(minI, maxI) + 1 is still 0
			count += e - max(minI, maxI) + 1
		}
		s = e + 1 // to find a new segment
	}
	return int64(count)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		nums []int
		minK int
		maxK int
		ans  int64
	}{
		{[]int{1, 3, 5, 2, 7, 5}, 1, 5, 2},
		{[]int{1, 1, 1, 1}, 1, 1, 10},
	} {
		fmt.Println(countSubarrays(v.nums, v.minK, v.maxK), v.ans)
	}
}

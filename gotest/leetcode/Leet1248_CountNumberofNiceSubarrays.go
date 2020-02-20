package main

import (
	"fmt"
)

// https://leetcode.com/problems/count-number-of-nice-subarrays/

// Given an array of integers nums and an integer k. A subarray is called nice if
// there are k odd numbers on it.
// Return the number of nice sub-arrays.
// Example 1:
//   Input: nums = [1,1,2,1,1], k = 3
//   Output: 2
//   Explanation: The only sub-arrays with 3 odd numbers are [1,1,2,1] and [1,2,1,1].
// Example 2:
//   Input: nums = [2,4,6], k = 1
//   Output: 0
//   Explanation: There is no odd numbers in the array.
// Example 3:
//   Input: nums = [2,2,2,1,2,2,1,2,2,2], k = 2
//   Output: 16
// Constraints:
//   1 <= nums.length <= 50000
//   1 <= nums[i] <= 10^5
//   1 <= k <= nums.length

func numberOfSubarrays(nums []int, k int) int {
	all := 0
	odds := []int{-1}
	for i := range nums {
		if nums[i]%2 == 1 {
			odds = append(odds, i)
		}
	}
	odds = append(odds, len(nums))
	for i := 1; i+k < len(odds); i++ {
		all += (odds[i] - odds[i-1]) * (odds[i+k] - odds[i+k-1])
	}
	return all
}

func main() {
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1, 2, 2, 1}, 2))
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
	fmt.Println(numberOfSubarrays([]int{2, 4, 6}, 1))
	fmt.Println(numberOfSubarrays([]int{1, 1}, 1))
	fmt.Println(numberOfSubarrays([]int{1}, 1))
	fmt.Println(numberOfSubarrays([]int{1, 2}, 1))
	fmt.Println(numberOfSubarrays([]int{1, 1}, 2))
	fmt.Println(numberOfSubarrays([]int{1, 1}, 1))
	fmt.Println(numberOfSubarrays([]int{2}, 2))
}

// 2,2,2,1,2,2,1,2,2,2,2,1,1,2,1,2,2,1
// ---------------------
// 2,2,2,1,2,2,1,2,2,2,2,1,1,2,1,2,2,1
//         ---------------
// 2,2,2,1,2,2,1,2,2,2,2,1,1,2,1,2,2,1
//               -----------
// 2,2,2,1,2,2,1,2,2,2,2,1,1,2,1,2,2,1
//                         ---------
// 2,2,2,1,2,2,1,2,2,2,2,1,1,2,1,2,2,1
//                           ---------

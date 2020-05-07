package main

import "fmt"

// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/

// Given an array nums of 0s and 1s and an integer k, return True if all 1's are
// at least k places away from each other, otherwise return False.
// Example 1:
//   Input: nums = [1,0,0,0,1,0,0,1], k = 2
//   Output: true
//   Explanation: Each of the 1s are at least 2 places away from each other.
// Example 2:
//   Input: nums = [1,0,0,1,0,1], k = 2
//   Output: false
//   Explanation: The second 1 and third 1 are only one apart from each other.
// Example 3:
//   Input: nums = [1,1,1,1,1], k = 0
//   Output: true
// Example 4:
//   Input: nums = [0,1,0,1], k = 1
//   Output: true
// Constraints:
//   1 <= nums.length <= 10^5
//   0 <= k <= nums.length
//   nums[i] is 0 or 1

func kLengthApart(nums []int, k int) bool {
	prev := 0
    for i := range nums {
		if nums[i] == 1 {
			prev = i
			break
		}
	}
	for j := prev+1; j<len(nums); j++ {
		if nums[j] == 1 {
			if j-prev-1 < k {
				return false
			}
			prev = j
		}
	}
	return true
}

func main() {
	fmt.Println(kLengthApart([]int{1,0,0,0,1,0,0,1}, 2))
	fmt.Println(kLengthApart([]int{1,0,0,1,0,1}, 2))
	fmt.Println(kLengthApart([]int{1,1,1,1,1}, 0))
	fmt.Println(kLengthApart([]int{0,1,0,1}, 1))
}
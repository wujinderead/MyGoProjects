package main

import "fmt"

// https://leetcode.com/problems/check-if-there-is-a-valid-partition-for-the-array/

// You are given a 0-indexed integer array nums. You have to partition the array
// into one or more contiguous subarrays.
// We call a partition of the array valid if each of the obtained subarrays
// satisfies one of the following conditions:
//   The subarray consists of exactly 2 equal elements. For example, the subarray [2,2] is good.
//   The subarray consists of exactly 3 equal elements. For example, the subarray [4,4,4] is good.
//   The subarray consists of exactly 3 consecutive increasing elements, that is,
//     the difference between adjacent elements is 1. For example, the subarray [3,4,5]
//     is good, but the subarray [1,3,5] is not.
// Return true if the array has at least one valid partition. Otherwise, return false.
// Example 1:
//   Input: nums = [4,4,4,5,6]
//   Output: true
//   Explanation: The array can be partitioned into the subarrays [4,4] and [4,5,6].
//     This partition is valid, so we return true.
// Example 2:
//   Input: nums = [1,1,1,2]
//   Output: false
//   Explanation: There is no valid partition for this array.
// Constraints:
//   2 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁶

func validPartition(nums []int) bool {
	dp := [4]bool{false, false, false, true} // initialize dp[-1] as true
	for i := range nums {
		dp[i%4] = false
		// if nums[i]==num[i-1], dp[i] ||= dp[i-2]
		if i-1 >= 0 && nums[i] == nums[i-1] {
			dp[i%4] = dp[i%4] || dp[(i+2)%4] // dp[(i-2)%4] == dp[(i+2)%4]
		}
		// if nums[i]==num[i-1]=nums[i-2], dp[i] ||= dp[i-3]
		if i-2 >= 0 && nums[i] == nums[i-1] && nums[i-1] == nums[i-2] {
			dp[i%4] = dp[i%4] || dp[(i+1)%4] // dp[(i-3)%4] == dp[(i+1)%4]
		}
		if i-2 >= 0 && nums[i] == nums[i-1]+1 && nums[i-1] == nums[i-2]+1 {
			dp[i%4] = dp[i%4] || dp[(i+1)%4]
		}
	}
	return dp[(len(nums)-1)%4]
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  bool
	}{
		{[]int{4, 4, 4, 5, 6}, true},
		{[]int{4, 4, 4, 5, 5}, true},
		{[]int{1, 1, 1, 2}, false},
		{[]int{1}, false},
		{[]int{1, 1}, true},
		{[]int{1, 1, 2}, false},
		{[]int{1, 1, 2, 3}, false},
		{[]int{1, 1, 1, 2, 3}, true},
		{[]int{1, 1, 1, 2, 3}, true},
		{[]int{1, 1, 1, 2, 3, 4}, true},
		{[]int{1, 1, 1, 2, 3, 3}, false},
		{[]int{1, 1, 1, 2, 3, 3, 3}, true},
		{[]int{1, 1, 1, 2, 3, 3, 4}, false},
		{[]int{1, 1, 1, 2, 3, 3, 4, 5}, true},
		{[]int{1, 1, 1, 2, 3, 3, 3, 4, 5}, false},
		{[]int{1, 1, 1, 2, 3, 3, 3, 3, 4, 5}, true},
		{[]int{1, 1, 1, 2, 3, 3, 3, 3, 3, 4, 5}, true},
	} {
		fmt.Println(validPartition(v.nums), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/find-all-good-indices/

// You are given a 0-indexed integer array nums of size n and a positive integer k.
// We call an index i in the range k <= i < n - k good if the following conditions are satisfied:
//   The k elements that are just before the index i are in non-increasing order.
//   The k elements that are just after the index i are in non-decreasing order.
// Return an array of all good indices sorted in increasing order.
// Example 1:
//   Input: nums = [2,1,1,1,3,4,1], k = 2
//   Output: [2,3]
//   Explanation: There are two good indices in the array:
//     - Index 2. The subarray [2,1] is in non-increasing order, and the subarray [1,3]
//       is in non-decreasing order.
//     - Index 3. The subarray [1,1] is in non-increasing order, and the subarray [3,4]
//       is in non-decreasing order.
//     Note that the index 4 is not good because [4,1] is not non-decreasing.
// Example 2:
//   Input: nums = [2,1,1,2], k = 2
//   Output: []
//   Explanation: There are no good indices in this array.
// Constraints:
//   n == nums.length
//   3 <= n <= 10⁵
//   1 <= nums[i] <= 10⁶
//   1 <= k <= n / 2

func goodIndices(nums []int, k int) []int {
	inc, dec, ans := make([]int, len(nums)), make([]int, len(nums)), []int{}
	dec[0] = 1
	inc[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			dec[i] = dec[i-1] + 1
		} else {
			dec[i] = 1
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			inc[i] = inc[i+1] + 1
		} else {
			inc[i] = 1
		}
	}
	for i := 1; i < len(nums)-1; i++ {
		if dec[i-1] >= k && inc[i+1] >= k {
			ans = append(ans, i)
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		nums []int
		k    int
		ans  []int
	}{
		{[]int{2, 1, 1, 1, 3, 4, 1}, 2, []int{2, 3}},
		{[]int{2, 1, 1, 2}, 2, []int{}},
		{[]int{253747, 459932, 263592, 354832, 60715, 408350, 959296}, 2, []int{3}},
	} {
		fmt.Println(goodIndices(v.nums, v.k), v.ans)
	}
}

package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-number-of-operations-to-make-array-continuous/

// You are given an integer array nums. In one operation, you can replace any element in nums
// with any integer.
// nums is considered continuous if both of the following conditions are fulfilled:
//   All elements in nums are unique.
//   The difference between the maximum element and the minimum element in nums equals nums.length - 1.
// For example, nums = [4, 2, 5, 3] is continuous, but nums = [1, 2, 3, 5, 6] is not continuous.
// Return the minimum number of operations to make nums continuous.
// Example 1:
//   Input: nums = [4,2,5,3]
//   Output: 0
//   Explanation: nums is already continuous.
// Example 2:
//   Input: nums = [1,2,3,5,6]
//   Output: 1
//   Explanation: One possible solution is to change the last element to 4.
//     The resulting array is [1,2,3,5,4], which is continuous.
// Example 3:
//   Input: nums = [1,10,100,1000]
//   Output: 3
//   Explanation: One possible solution is to:
//     - Change the second element to 2.
//     - Change the third element to 3.
//     - Change the fourth element to 4.
//     The resulting array is [1,2,3,4], which is continuous.
// Constraints:
//   1 <= nums.length <= 10^5
//   1 <= nums[i] <= 10^9

// sort unique elements, and for each element x,
// find how many numbers are in the range [x, x+len(nums)-1]
// that means use binary search to find the index of value x+len(nums)-1
// e.g. nums = [1,2,10,100,1000], find how many elements in these ranges:
// 1-5, 2-6, 10-14, 100-104, 1000-1004, respectively
func minOperations(nums []int) int {
	set := make(map[int]struct{})
	for _, v := range nums {
		set[v] = struct{}{}
	}
	keys := make([]int, 0)
	for k := range set {
		keys = append(keys, k)
	}
	keys = append(keys, 1e10) // append a max number
	// get unique elements and then sort
	sort.Sort(sort.IntSlice(keys))

	min := len(nums) - 1
	for i := 0; i < len(keys)-1; i++ {
		target := keys[i] + len(nums) - 1
		// binary search to find the index to insert target,
		// i.e., find first index x that keys[x] > target
		// NOTE: can also use sliding window here, as the target index always go right
		l, r := i, len(keys)-1
		for l < r {
			mid := (l + r) / 2
			if keys[mid] <= target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		// if we find n unique numbers in range [nums[i], nums[i]+len(nums)-1]
		// we need len(nums)-n operations to make array continuous
		if len(nums)-(l-i) < min {
			min = len(nums) - (l - i)
		}
	}
	return min
}

func main() {
	for _, v := range []struct {
		n   []int
		ans int
	}{
		{[]int{4, 2, 5, 3}, 0},
		{[]int{1, 2, 3, 5, 6}, 1},
		{[]int{1, 10, 100, 1000}, 3},
		{[]int{1, 1, 2, 2, 3}, 2},
	} {
		fmt.Println(minOperations(v.n), v.ans)
	}
}

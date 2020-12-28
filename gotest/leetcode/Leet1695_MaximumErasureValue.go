package main

import "fmt"

// https://leetcode.com/problems/maximum-erasure-value/

// You are given an array of positive integers nums and want to erase a subarray containing
// unique elements. The score you get by erasing the subarray is equal to the sum of its
// elements. Return the maximum score you can get by erasing exactly one subarray.
// An array b is called to be a subarray of a if it forms a contiguous subsequence of a,
// that is, if it is equal to a[l],a[l+1],...,a[r] for some (l,r).
// Example 1:
//   Input: nums = [4,2,4,5,6]
//   Output: 17
//   Explanation: The optimal subarray here is [2,4,5,6].
// Example 2:
//   Input: nums = [5,2,1,2,5,2,1,2,5]
//   Output: 8
//   Explanation: The optimal subarray here is [5,2,1] or [1,2,5].
// Constraints:
//   1 <= nums.length <= 10^5
//   1 <= nums[i] <= 10^4

// just use two pointer
func maximumUniqueSubarray(nums []int) int {
	var l, r int
	ans := 0
	sum := 0
	count := make(map[int]struct{})
	for r < len(nums) {
		if _, ok := count[nums[r]]; !ok {
			sum += nums[r]
			count[nums[r]] = struct{}{}
			if sum > ans {
				ans = sum
			}
			r++
			continue
		}
		for nums[l] != nums[r] {
			delete(count, nums[l])
			sum -= nums[l]
			l++
		}
		delete(count, nums[l])
		sum -= nums[l]
		l++
	}
	return ans
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{4, 2, 4, 5, 6}, 17},
		{[]int{5, 2, 1, 2, 5, 2, 1, 2, 5}, 8},
		{[]int{2}, 2},
		{[]int{2, 2}, 2},
		{[]int{1, 2}, 3},
		{[]int{1, 2, 2}, 3},
		{[]int{5, 1, 1}, 6},
	} {
		fmt.Println(maximumUniqueSubarray(v.nums), v.ans)
	}
}

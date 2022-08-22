package main

import "fmt"

// https://leetcode.com/problems/maximum-segment-sum-after-removals/

// You are given two 0-indexed integer arrays nums and removeQueries, both of length n.
// For the iᵗʰ query, the element in nums at the index removeQueries[i] is removed,
// splitting nums into different segments.
// A segment is a contiguous sequence of positive integers in nums. A segment sum is the sum
// of every element in a segment.
// Return an integer array answer, of length n, where answer[i] is the maximum segment sum after
// applying the iᵗʰ removal.
// Note: The same index will not be removed more than once.
// Example 1:
//   Input: nums = [1,2,5,6,1], removeQueries = [0,3,2,4,1]
//   Output: [14,7,2,2,0]
//   Explanation: Using 0 to indicate a removed element, the answer is as follows:
//     Query 1: Remove the 0th element, nums becomes [0,2,5,6,1] and the maximum
//     segment sum is 14 for segment [2,5,6,1].
//     Query 2: Remove the 3rd element, nums becomes [0,2,5,0,1] and the maximum
//     segment sum is 7 for segment [2,5].
//     Query 3: Remove the 2nd element, nums becomes [0,2,0,0,1] and the maximum
//     segment sum is 2 for segment [2].
//     Query 4: Remove the 4th element, nums becomes [0,2,0,0,0] and the maximum
//     segment sum is 2 for segment [2].
//     Query 5: Remove the 1st element, nums becomes [0,0,0,0,0] and the maximum
//     segment sum is 0, since there are no segments.
//     Finally, we return [14,7,2,2,0].
// Example 2:
//   Input: nums = [3,2,11,1], removeQueries = [3,2,1,0]
//   Output: [16,5,3,0]
//   Explanation: Using 0 to indicate a removed element, the answer is as follows:
//     Query 1: Remove the 3rd element, nums becomes [3,2,11,0] and the maximum
//     segment sum is 16 for segment [3,2,11].
//     Query 2: Remove the 2nd element, nums becomes [3,2,0,0] and the maximum segment
//     sum is 5 for segment [3,2].
//     Query 3: Remove the 1st element, nums becomes [3,0,0,0] and the maximum segment
//     sum is 3 for segment [3].
//     Query 4: Remove the 0th element, nums becomes [0,0,0,0] and the maximum segment
//     sum is 0, since there are no segments.
//     Finally, we return [16,5,3,0].
// Constraints:
//   n == nums.length == removeQueries.length
//   1 <= n <= 10⁵
//   1 <= nums[i] <= 10⁹
//   0 <= removeQueries[i] < n
//   All the values of removeQueries are unique.

// reverse the procedure by adding numbers
func maximumSegmentSum(nums []int, removeQueries []int) []int64 {
	ans := make([]int64, len(nums))
	prefix := make([]int64, len(nums)+1)
	border := make([]int, len(nums)) // border[i]: the length of segment with i as the border
	for i := range nums {
		border[i] = -1
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	// start by adding last deleted number
	var max int64
	for i := len(nums) - 1; i > 0; i-- {
		ind := removeQueries[i] // the index of added number
		left, right := ind, ind // left border, right border, length of current segment
		if ind+1 < len(nums) && border[ind+1] >= 0 {
			right += border[ind+1]
		}
		if ind-1 >= 0 && border[ind-1] >= 0 {
			left -= border[ind-1]
		}
		border[left] = right - left + 1
		border[right] = right - left + 1
		if prefix[right+1]-prefix[left] > max {
			max = prefix[right+1] - prefix[left]
		}
		ans[i-1] = max
	}
	return ans
}

func main() {
	for _, v := range []struct {
		nums          []int
		removeQueries []int
		ans           []int64
	}{
		{[]int{1, 2, 5, 6, 1}, []int{0, 3, 2, 4, 1}, []int64{14, 7, 2, 2, 0}},
		{[]int{3, 2, 11, 1}, []int{3, 2, 1, 0}, []int64{16, 5, 3, 0}},
	} {
		fmt.Println(maximumSegmentSum(v.nums, v.removeQueries), v.ans)
	}
}

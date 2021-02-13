package main

import "fmt"

// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/

// You are given an integer array nums. The absolute sum of a subarray
// [numsl, numsl+1, ..., numsr-1, numsr] is abs(numsl + numsl+1 + ... + numsr-1 + numsr).
// Return the maximum absolute sum of any (possibly empty) subarray of nums.
// Note that abs(x) is defined as follows:
//   If x is a negative integer, then abs(x) = -x.
//   If x is a non-negative integer, then abs(x) = x.
// Example 1:
//   Input: nums = [1,-3,2,3,-4]
//   Output: 5
//   Explanation: The subarray [2,3] has absolute sum = abs(2+3) = abs(5) = 5.
// Example 2:
//   Input: nums = [2,-5,1,-4,3,-2]
//   Output: 8
//   Explanation: The subarray [-5,1,-4] has absolute sum = abs(-5+1-4) = abs(-8) = 8.
// Constraints:
//   1 <= nums.length <= 10^5
//   -10^4 <= nums[i] <= 10^4

// another method:
// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/discuss/1052527/JavaC%2B%2BPython-O(1)-Space
// as subarray_sum = one_prefix_sum - another_prefix_sum, thus,
// max(subarray_sum) = max(prefix_sum) - min(prefix_sum)
func maxAbsoluteSum(nums []int) int {
	// pos is the max positive sum of sum(nums[0...i]), sum(nums[1...i]), ..., sum(nums[i...i])
	// neg is the min negative sum of sum(nums[0...i]), sum(nums[1...i]), ..., sum(nums[i...i])
	var pos, neg, ans int
	for _, v := range nums {
		if v > 0 {
			ans = max(ans, pos+v)
		} else {
			ans = max(ans, -neg-v)
		}
		pos = max(pos+v, 0)
		neg = min(neg+v, 0)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{1, -3, 2, 3, -4}, 5},
		{[]int{2, -5, 1, -4, 3, -2}, 8},
	} {
		fmt.Println(maxAbsoluteSum(v.nums), v.ans)
	}
}

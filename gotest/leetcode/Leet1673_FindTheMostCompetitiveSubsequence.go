package main

import "fmt"

// https://leetcode.com/problems/find-the-most-competitive-subsequence/

// Given an integer array nums and a positive integer k, return the most competitive subsequence
// of nums of size k.
// An array's subsequence is a resulting sequence obtained by erasing some (possibly zero)
// elements from the array.
// We define that a subsequence a is more competitive than a subsequence b (of the same length)
// if in the first position where a and b differ, subsequence a has a number less than the
// corresponding number in b. For example, [1,3,4] is more competitive than [1,3,5] because
// the first position they differ is at the final number, and 4 is less than 5.
// Example 1:
//   Input: nums = [3,5,2,6], k = 2
//   Output: [2,6]
//   Explanation: Among the set of every possible subsequence:
//     {[3,5], [3,2], [3,6], [5,2], [5,6], [2,6]}, [2,6] is the most competitive.
// Example 2:
//   Input: nums = [2,4,3,3,5,4,9,6], k = 4
//   Output: [2,3,3,4]
// Constraints:
//   1 <= nums.length <= 10^5
//   0 <= nums[i] <= 10^9
//   1 <= k <= nums.length

// stack and greedy: make a array of size k, push to tail;
// if cur_value < stack tail, replace for best effort
func mostCompetitive(nums []int, k int) []int {
	ans := make([]int, 0, k)
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		for len(ans) > 0 && v < ans[len(ans)-1] && len(nums)-i >= k-(len(ans)-1) {
			ans = ans[:len(ans)-1]
		}
		if len(ans) < k {
			ans = append(ans, v)
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
		{[]int{3, 5, 2, 6}, 2, []int{2, 6}},
		{[]int{2, 4, 3, 3, 5, 4, 9, 6}, 4, []int{2, 3, 3, 4}},
		{[]int{3, 5, 6, 2}, 2, []int{3, 2}},
		{[]int{3, 5, 6, 2, 7}, 3, []int{3, 2, 7}},
	} {
		fmt.Println(mostCompetitive(v.nums, v.k), v.ans)
	}
}

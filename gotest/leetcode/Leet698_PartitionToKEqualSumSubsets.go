package main

import "fmt"

// https://leetcode.com/problems/partition-to-k-equal-sum-subsets/

// Given an integer array nums and an integer k, return true if it is possible to
// divide this array into k non-empty subsets whose sums are all equal.
// Example 1:
//   Input: nums = [4,3,2,3,5,2,1], k = 4
//   Output: true
//   Explanation: It is possible to divide it into 4 subsets (5), (1,4), (2,3), (2,3) with equal sums.
// Example 2:
//   Input: nums = [1,2,3,4], k = 3
//   Output: false
// Constraints:
//   1 <= k <= nums.length <= 16
//   1 <= nums[i] <= 10⁴

func canPartitionKSubsets(nums []int, k int) bool {
	// pre determine
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}
	avg := sum / k
	for _, v := range nums {
		if v > avg { // some single value >= avg, return false
			return false
		}
	}
	// dp[mask]: using mask nums, the residue we part these numbers
	dp := make([]int, 1<<len(nums))
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 0; i < 1<<len(nums); i++ { // for each mask
		if dp[i] < 0 { // if not valid
			continue
		}
		for j := 0; j < len(nums); j++ {
			if i|(1<<j) == i {
				continue
			}
			// set current mask's jth bit to 1 (i.e. add j-th number to current mask)
			newmask := i | (1 << j)
			if dp[i]+nums[j] <= avg { // for example, avg = 5 and dp[i]=2, the we can add 1,2,3 to current set, but not 4
				dp[newmask] = (dp[i] + nums[j]) % avg
			}
		}
	}
	return dp[(1<<len(nums))-1] == 0
}

func main() {
	for _, v := range []struct {
		nums []int
		k    int
		ans  bool
	}{
		{[]int{5, 4, 3, 2, 3, 5, 2, 1, 5}, 6, true},
		{[]int{4, 3, 2, 3, 5, 2, 1}, 4, true},
		{[]int{1, 2, 3, 4}, 3, false},
		{[]int{2, 2, 2, 4}, 2, false},
	} {
		fmt.Println(canPartitionKSubsets(v.nums, v.k), v.ans)
	}
}

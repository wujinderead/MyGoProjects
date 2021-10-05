package main

import "fmt"

// https://leetcode.com/problems/maximum-number-of-ways-to-partition-an-array/

// You are given a 0-indexed integer array nums of length n. The number of ways to partition nums is
// the number of pivot indices that satisfy both conditions:
//   1 <= pivot < n
//   nums[0] + nums[1] + ... + nums[pivot - 1] == nums[pivot] + nums[pivot + 1] + ... + nums[n - 1]
// You are also given an integer k. You can choose to change the value of one element of nums to k,
// or to leave the array unchanged.
// Return the maximum possible number of ways to partition nums to satisfy both conditions after
// changing at most one element.
// Example 1:
//   Input: nums = [2,-1,2], k = 3
//   Output: 1
//   Explanation: One optimal approach is to change nums[0] to k. The array becomes [3,-1,2].
//     There is one way to partition the array:
//     - For pivot = 2, we have the partition [3,-1 | 2]: 3 + -1 == 2.
// Example 2:
//   Input: nums = [0,0,0], k = 1
//   Output: 2
//   Explanation: The optimal approach is to leave the array unchanged.
//     There are two ways to partition the array:
//     - For pivot = 1, we have the partition [0 | 0,0]: 0 == 0 + 0.
//     - For pivot = 2, we have the partition [0,0 | 0]: 0 + 0 == 0.
// Example 3:
//   Input: nums = [22,4,-25,-20,-15,15,-16,7,19,-10,0,-13,-14], k = -33
//   Output: 4
//   Explanation: One optimal approach is to change nums[2] to k. The array becomes
//     [22,4,-33,-20,-15,15,-16,7,19,-10,0,-13,-14].
//     There are four ways to partition the array.
// Constraints:
//   n == nums.length
//   2 <= n <= 10^5
//   -10^5 <= k, nums[i] <= 10^5

func waysToPartition(nums []int, k int) int {
	max := 0                         // the answer
	prefix := make([]int, len(nums)) // prefix sum array
	prefix[0] = nums[0]
	mapp := make(map[int]int) // occurrence of prefix sum
	mapp[prefix[0]] = 1
	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i]
		mapp[prefix[i]] = mapp[prefix[i]] + 1
	}
	sum := prefix[len(nums)-1]

	// for unchanged, find how many prefix_sum == sum/2
	if sum%2 == 0 {
		if sum == 0 && mapp[sum] > 1 {
			max = mapp[sum] - 1
		} else if sum != 0 {
			max = mapp[sum/2]
		}
	}

	// replace for each number with k
	leftMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// change nums[i] with k, we got newsum; we need find prefix_sum == newsum/2
		// for nums[0...i-1], prefix sums are unchanged, search target is newsum/2
		// for nums[i...], each prefix_sum = each original prefix_sum+k-nums[i],
		// so we need change the search target to newsum/2-k+nums[i]
		newsum := sum - nums[i] + k
		if newsum%2 != 0 {
			leftMap[prefix[i]] = leftMap[prefix[i]] + 1
			mapp[prefix[i]] = mapp[prefix[i]] - 1
			continue
		}
		n := leftMap[newsum/2] + mapp[newsum/2-k+nums[i]]
		if newsum == 0 {
			n--
		}
		if n > max {
			max = n
		}
		leftMap[prefix[i]] = leftMap[prefix[i]] + 1
		mapp[prefix[i]] = mapp[prefix[i]] - 1
	}
	return max
}

func main() {
	for _, v := range []struct {
		n      []int
		k, ans int
	}{
		{[]int{2, -1, 2}, 3, 1},
		{[]int{0, 0, 0}, 1, 2},
		{[]int{22, 4, -25, -20, -15, 15, -16, 7, 19, -10, 0, -13, -14}, -33, 4},
		{[]int{-2, 2, 3, -3}, 10, 1},
		{[]int{-1, -2, -3, 6}, 20, 0},
	} {
		fmt.Println(waysToPartition(v.n, v.k), v.ans)
	}
}

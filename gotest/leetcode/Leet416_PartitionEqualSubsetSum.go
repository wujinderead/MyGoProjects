package main

import "fmt"

// Given a non-empty array containing only positive integers, find if the array can be
// partitioned into two subsets such that the sum of elements in both subsets is equal.
// Note: Each of the array element will not exceed 100. The array size will not exceed 200.
// Example 1:
//   Input: [1, 5, 11, 5]
//   Output: true
//   Explanation: The array can be partitioned as [1, 5, 5] and [11].
// Example 2:
//   Input: [1, 2, 3, 5]
//   Output: false
//   Explanation: The array cannot be partitioned into equal sum subsets.

func canPartition(nums []int) bool {
	s := 0
	for i := range nums {
		s += nums[i]
	}
	if s%2 == 1 || len(nums) == 1 {
		return false
	}
	// find if some numbers can sum to s/2
	s = s / 2
	// let c(i, s) be whether we can sum to s with nums[:i], then
	// c(i, s) = c(i-1, s) || c(i-1, s-nums[i])
	// base case c(x, 0)=true, c(0, nums[0])=true
	old := make([]bool, s+1)
	new := make([]bool, s+1)
	old[0] = true
	if nums[0] <= s {
		old[nums[0]] = true
	}
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= s; j++ {
			new[j] = old[j]
			if j-nums[i] >= 0 {
				new[j] = new[j] || old[j-nums[i]]
			}
		}
		old, new = new, old
	}
	return old[s]
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
	fmt.Println(canPartition([]int{1, 6, 3, 2}))
	fmt.Println(canPartition([]int{2, 6}))
	fmt.Println(canPartition([]int{3}))
}

package leetcode

import "fmt"

// https://leetcode.com/problems/target-sum/

// You are given a list of non-negative integers, a1, a2, ..., an, and a target, S.
// Now you have 2 symbols + and -. For each integer,
// you should choose one from + and - as its new symbol.
// Find out how many ways to assign symbols to make sum of integers equal to target S.
// Example 1:
//   Input: nums is [1, 1, 1, 1, 1], S is 3.
//   Output: 5
//   Explanation: There are 5 ways to assign symbols to make the sum of nums be target 3.
//     -1+1+1+1+1 = 3
//     +1-1+1+1+1 = 3
//     +1+1-1+1+1 = 3
//     +1+1+1-1+1 = 3
//     +1+1+1+1-1 = 3
// Note:
//   The length of the given array is positive and will not exceed 20.
//   The sum of elements in the given array will not exceed 1000.
//   Your output answer is guaranteed to be fitted in a 32-bit integer.

func findTargetSumWays(nums []int, S int) int {
	// let c[i, j] be the number of ways to get sum j for nums[0...i]
	// then c[i, j] = c[i-1, j+nums[i]] + c[i-1, j-nums[i]]
	// base case c[0, ±nums[0]] = 1, c[x, 0] = 0
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if S > sum || S < -sum {
		return 0
	}
	c0 := make([]int, 2*sum+1)
	c1 := make([]int, 2*sum+1)
	if nums[0] != 0 {
		c0[nums[0]+sum] = 1
		c0[-nums[0]+sum] = 1
	} else {
		c0[sum] = 2
	}
	fmt.Println(c0)
	for i := 1; i < len(nums); i++ {
		for j := -sum; j <= sum; j++ {
			c1[j+sum] = 0 // clear c1
			if j+nums[i] <= sum {
				c1[j+sum] += c0[j+nums[i]+sum]
			}
			if j-nums[i] >= -sum {
				c1[j+sum] += c0[j-nums[i]+sum]
			}
		}
		fmt.Println(c1)
		c0, c1 = c1, c0
	}
	return c0[S+sum]
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 4))
	fmt.Println(findTargetSumWays([]int{1, 0, 1, 1, 1, 1}, 3))
}

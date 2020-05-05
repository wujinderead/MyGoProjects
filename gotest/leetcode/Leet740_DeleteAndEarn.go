package main

import "fmt"

// https://leetcode.com/problems/delete-and-earn/

// Given an array nums of integers, you can perform operations on the array.
// In each operation, you pick any nums[i] and delete it to earn nums[i] points.
// After, you must delete every element equal to nums[i] - 1 or nums[i] + 1.
// You start with 0 points. Return the maximum number of points you can earn by
// applying such operations.
// Example 1:
//   Input: nums = [3, 4, 2]
//   Output: 6
//   Explanation:
//     Delete 4 to earn 4 points, consequently 3 is also deleted.
//     Then, delete 2 to earn 2 points. 6 total points are earned.
// Example 2:
//   Input: nums = [2, 2, 3, 3, 3, 4]
//   Output: 9
//   Explanation:
//     Delete 3 to earn 3 points, deleting both 2's and the 4.
//     Then, delete 3 again to earn 3 points, and 3 again to earn 3 points.
//     9 total points are earned.
// Note:
//   The length of nums is at most 20000.
//   Each element nums[i] is an integer in the range [1, 10000].

func deleteAndEarn(nums []int) int {
	// count the occurrence of each integer
    count := make([]int, 10001)
    for _, v := range nums {
    	count[v]++
	}
	dp := make([]int, 10001)
	dp[1] = count[1]
	for i:=2; i<len(count); i++ {
		dp[i] = dp[i-1]
		if dp[i-2]+i*count[i] > dp[i-1] {
			dp[i] = dp[i-2]+i*count[i]
		}
	}
	return dp[len(count)-1]
}

func main() {
	fmt.Println(deleteAndEarn([]int{3,4,2}))
	fmt.Println(deleteAndEarn([]int{2,2,2,3,3,4,4,4,5,5,7,7,7,8,8,8,8,9,11}))
	fmt.Println(deleteAndEarn([]int{10,8,4,2,1,3,4,8,2,9,10,4,8,5,9,1,5,1,6,8,1,1,6,7,8,9,1,7,6,8,4,5,4,1,5,9,8,6,10,6,4,3,8,4,10,8,8,10,6,4,4,4,9,6,9,10,7,1,5,3,4,4,8,1,1,2,1,4,1,1,4,9,4,7,1,5,1,10,3,5,10,3,10,2,1,10,4,1,1,4,1,2,10,9,7,10,1,2,7,5}))
}
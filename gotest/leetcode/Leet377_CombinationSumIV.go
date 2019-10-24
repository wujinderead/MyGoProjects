package leetcode

import "fmt"

// https://leetcode.com/problems/combination-sum-iv/

// Given an integer array with all positive numbers and no duplicates,
// find the number of possible combinations that add up to a positive integer target.
// Example:
//   nums = [1, 2, 3]
//   target = 4
//   The possible combination ways are:
//     (1, 1, 1, 1)
//     (1, 1, 2)
//     (1, 2, 1)
//     (1, 3)
//     (2, 1, 1)
//     (2, 2)
//     (3, 1)
//   Note that different sequences are counted as different combinations.
//   Therefore the output is 7.
// Follow up: 
//   What if negative numbers are allowed in the given array?
//   How does it change the problem?
//   What limitation we need to add to the question to allow negative numbers?

func combinationSum4(nums []int, target int) int {
	if len(nums)==0 || target==0 {
		return 0
	}
	// let c(j) be the number of ways to sum j with integer set {A0, A1, ..., An-1}
	// then c(j)=c(j-A0) + c(j-A1) + ... c(j-An-1), base case c(0)=1
	c := make([]int, target+1)
	c[0] = 1
	for j:=1; j<=target; j++ {
		for i := range nums {
			if j-nums[i]>=0 {
				c[j] += c[j-nums[i]]
			}
		}
	}
    return c[target]
}

func main() {
    fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
    fmt.Println(combinationSum4([]int{1}, 4))
}
package main

import (
    "fmt"
)

// https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/

// In a given array nums of positive integers, find three non-overlapping subarrays with maximum sum.
// Each subarray will be of size k, and we want to maximize the sum of all 3*k entries.
// Return the result as a list of indices representing the starting position of each interval (0-indexed). 
// If there are multiple answers, return the lexicographically smallest one.
// Example:
//   Input: [1,2,1,2,6,7,5,1], 2
//   Output: [0, 3, 5]
//   Explanation: Subarrays [1, 2], [2, 6], [7, 5] correspond to the starting indices [0, 3, 5].
//     We could have also taken [2, 1], but an answer of [1, 3, 5] would be lexicographically larger.
// Note:
//   nums.length will be between 1 and 20000.
//   nums[i] will be between 1 and 65535.
//   k will be between 1 and floor(nums.length / 3).

// another greedy method, just use three pointers and O(1) space:
// https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/discuss/108238/Python-o(n)-time-o(1)-space.-Greedy-solution.

func maxSumOfThreeSubarrays(nums []int, k int) []int {
    // to find 3 non-overlapping subarrays in nums[:i], 
    // if we use nums[i-k: i], we need to find 2 subarrays in nums[:i-k];
    // if we don't use, we need to find 3 subrarrys in nums[:i-1]
    sum := make([]int, len(nums)+1)
    for i:=1; i<=len(nums); i++ {
    	sum[i] = sum[i-1]+nums[i-1]
    }

    one := make([][2]int, len(nums))
    one[k-1] = [2]int{sum[k], 0}    // (sum, start) pair
    for i:=k; i<len(nums); i++ {    // from [0...i], select one k-length subarray with max sum
    	one[i] = one[i-1]
    	if sum[i+1]-sum[i-k+1] > one[i][0] {
    		one[i][0] = sum[i+1]-sum[i-k+1]
    		one[i][1] = i-k+1
    	}
    } 
    // fmt.Println(one)

    two := make([][3]int, len(nums)) 
    two[2*k-1] = [3]int{sum[2*k], 0, k}   // (sum, start1, starrt2) pair           
    for i:=2*k; i<len(nums); i++ {        // from [0...i], select two k-length subarrays with max sum
    	two[i] = two[i-1]
    	if one[i-k][0]+sum[i+1]-sum[i-k+1] > two[i][0] {
    		two[i][0] = one[i-k][0]+sum[i+1]-sum[i-k+1]
    		two[i][1] = one[i-k][1]
    		two[i][2] = i-k+1
    	}
    }
    // fmt.Println(two)

    three := make([][4]int, len(nums))
    three[3*k-1] = [4]int{sum[3*k], 0, k, 2*k}  // (sum, start1, starrt2, start3) pair    
    for i:=3*k; i<len(nums); i++ {              // from [0...i], select three k-length subarrays with max sum
    	// select 3 subarrays from [0...i-k], or,
    	// select 2 from [0...i-k], plus [i-k+1...i]
    	three[i] = three[i-1]
    	if two[i-k][0]+sum[i+1]-sum[i-k+1] > three[i][0] {
    		three[i][0] = two[i-k][0]+sum[i+1]-sum[i-k+1]
    		three[i][1] = two[i-k][1]
    		three[i][2] = two[i-k][2]
    		three[i][3] = i-k+1
    	}
    }
    // fmt.Println(three)

    return three[len(nums)-1][1:]
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSumOfThreeSubarrays([]int{1,2,1,2,6,7,5,1}, 2))
	fmt.Println(maxSumOfThreeSubarrays([]int{1,2,1,2,1,7,2}, 2))
}
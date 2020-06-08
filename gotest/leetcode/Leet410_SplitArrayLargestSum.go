package main

import (
    "fmt"
)

// https://leetcode.com/problems/split-array-largest-sum/

// Given an array which consists of non-negative integers and an integer m, you can split the 
// array into m non-empty continuous subarrays. Write an algorithm to minimize the largest sum 
// among these m subarrays.
// Note: 
//   If n is the length of array, assume the following constraints are satisfied:
//   1 ≤ n ≤ 1000 
//   1 ≤ m ≤ min(50, n) 
// Examples: 
//   Input:
//     nums = [7,2,5,10,8]
//     m = 2
//   Output:
//     18
//   Explanation:
//     There are four ways to split nums into two subarrays.
//     The best way is to split it into [7,2,5] and [10,8],
//     where the largest sum among the two subarrays is only 18.

// another method, binary search:
// https://leetcode.com/problems/split-array-largest-sum/discuss/89817/Clear-Explanation%3A-8ms-Binary-Search-Java

func splitArray(nums []int, m int) int {
    // let dp(i, j) be the minimal largest sum to part nums[0...i] to j parts, then,
    // we can part nums[0...i-1] to j-1 parts, nums[i] be the j-th part; or,
    // part nums[0...i-1] to j-2 parts, nums[i-1...i] be the j-th part...; thus,
    // dp(i, j) = min( max(dp(i-1, j-1), nums[i]), max(dp(i-2, j-1), nums[i-1]+nums[i]), ...) 
    // time O(m*n^2)
    dp := make([][]int, len(nums))
    for i:=range dp {
    	dp[i] = make([]int, min(i+1,m)+1)
    }
    dp[0][1] = nums[0]   // to split nums[0...0] to 1 part 
    for i:=1; i<len(nums); i++ {   // to split nums[0...i] to j part
    	dp[i][1] = dp[i-1][1]+nums[i]
    	for j:=2; j<len(dp[i]); j++ {
    		s := nums[i]
    		dp[i][j] = max(dp[i-1][j-1], s)
    		for k:=i-2; k+1>=j-1 && s+nums[k+1]<dp[i][j]; k-- {
    			s += nums[k+1]
    			dp[i][j] = min(dp[i][j], max(dp[k][j-1], s))
    		}
    	}
    }
    return dp[len(nums)-1][m]
}

// have l = max number of array; r = sum of all numbers in the array. the target is between l and r.
// mid = (l+r)/2. for mid, we calculate the most parts it can be partitioned. 
func splitArrayBinarySearch(nums []int, m int) int {
    r := 0
    l := 0
    for _, v := range nums {
        r += v
        l = max(l, v)
    }
    for l<=r {
        mid := (l+r)/2
        s := 0
        part := 1
        for _, v := range nums {
            if s+v<=mid {
                s += v
            } else {
                s = v
                part++
            }
        }
        if part>m {  // part>m means mid is to low, increase l
            l = mid+1
        } else {
            r = mid-1
        }
    }
    return l
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
    for _, s := range []func([]int, int) int {splitArray, splitArrayBinarySearch} {
	   fmt.Println(s([]int{7,2,5,10,8}, 1))
	   fmt.Println(s([]int{7,2,5,10,8}, 2))
	   fmt.Println(s([]int{7,2,5,10,8}, 3))
	   fmt.Println(s([]int{7,2,5,10,8}, 4))
	   fmt.Println(s([]int{7,2,5,10,8}, 5))
	   fmt.Println(s([]int{3,5}, 1))
	   fmt.Println(s([]int{3,5}, 2))
    }
}
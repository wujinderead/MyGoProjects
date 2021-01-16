package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-operations-to-make-a-subsequence/

// You are given an array target that consists of 𝐝𝐢𝐬𝐭𝐢𝐧𝐜𝐭 integers and another integer
// array arr that 𝐜𝐚𝐧 have duplicates.
// In one operation, you can insert any integer at any position in arr. For example, if
// arr = [1,4,1,2], you can add 3 in the middle and make it [1,4,3,1,2]. Note that you can
// insert the integer at the very beginning or end of the array.
// Return the minimum number of operations needed to make target a subsequence of arr.
// A subsequence of an array is a new array generated from the original array by deleting
// some elements (possibly none) without changing the remaining elements' relative order.
// For example, [2,7,4] is a subsequence of [4,2,3,7,2,1,4] (the underlined elements),
// while [2,4,2] is not.
// Example 1:
//   Input: target = [5,1,3], arr = [9,4,2,3,4]
//   Output: 2
//   Explanation: You can add 5 and 1 in such a way that makes arr = [5,9,4,1,2,3,4],
//     then target will be a subsequence of arr.
// Example 2:
//   Input: target = [6,4,8,1,3,2], arr = [4,7,6,2,3,8,6,1]
//   Output: 3
// Constraints:
//   1 <= target.length, arr.length <= 10^5
//   1 <= target[i], arr[i] <= 10^9
//   target contains no duplicates.

// the straightforward thought is that find the longest common subsequence of target and arr,
// however, the complexity will be O(mn). then we found that the target has DISTINCT numbers,
// so the numbers in target have orders. Therefore, we use this order to find the
// longest increasing subsequence in arr.
func minOperations(target []int, arr []int) int {
	mapp := make(map[int]int)
	for i, v := range target {
		mapp[v] = i
	}
	// map elements in arr to the order of target
	newarr := make([]int, 0)
	for _, v := range arr {
		if order, ok := mapp[v]; ok {
			newarr = append(newarr, order)
		}
	}
	// find LIS of newarr
	// dp[i] is the minimal last value that an (i+1)-length increasing subsequence can have
	// e.g., dp[3] = 10, means there exists an increasing subsequence that has length 4,
	// and the minimal last value is 10.
	dp := make([]int, 0, len(newarr))
	for i := range newarr {
		ind := sort.SearchInts(dp, newarr[i])
		if ind == len(dp) {
			dp = append(dp, newarr[i])
		} else {
			dp[ind] = newarr[i]
		}
	}
	return len(target) - len(dp)
}

func main() {
	for _, v := range []struct {
		tar, arr []int
		ans      int
	}{
		{[]int{5, 1, 3}, []int{9, 4, 2, 3, 4}, 2},
		{[]int{6, 4, 8, 1, 3, 2}, []int{4, 7, 6, 2, 3, 8, 6, 1}, 3},
	} {
		fmt.Println(minOperations(v.tar, v.arr), v.ans)
	}
}

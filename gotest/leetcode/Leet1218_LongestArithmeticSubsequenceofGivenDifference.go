package main

import "fmt"

// https://leetcode.com/problems/longest-arithmetic-subsequence-of-given-difference

// Given an integer array arr and an integer difference, return the length of the
// longest subsequence in arr which is an arithmetic sequence such that the
// difference between adjacent elements in the subsequence equals difference.
// Example 1:
//   Input: arr = [1,2,3,4], difference = 1
//   Output: 4
//   Explanation: The longest arithmetic subsequence is [1,2,3,4].
// Example 2:
//   Input: arr = [1,3,5,7], difference = 1
//   Output: 1
//   Explanation: The longest arithmetic subsequence is any single element.
// Example 3:
//   Input: arr = [1,5,7,8,5,3,4,2,1], difference = -2
//   Output: 4
//   Explanation: The longest arithmetic subsequence is [7,5,3,1].
// Constraints:
//   1 <= arr.length <= 10^5
//   -10^4 <= arr[i], difference <= 10^4

func longestSubsequence(arr []int, difference int) int {
	mapp := make(map[int]int)
	allmax := 1
	for i := range arr {
		v, ok := mapp[arr[i]-difference]
		if !ok {
			mapp[arr[i]] = 1
		} else {
			mapp[arr[i]] = v + 1
			allmax = max(allmax, v+1)
		}
	}
	return allmax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestSubsequence([]int{1, 2, 3, 4}, 1))
	fmt.Println(longestSubsequence([]int{1, 3, 5, 7}, 1))
	fmt.Println(longestSubsequence([]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2))
}

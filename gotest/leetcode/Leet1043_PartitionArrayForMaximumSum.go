package main

import "fmt"

// Given an integer array A, you partition the array into (contiguous) subarrays
// of length at most K. After partitioning, each subarray has their values changed
// to become the maximum value of that subarray.
// Return the largest sum of the given array after partitioning.
// Example 1:
//   Input: A = [1,15,7,9,2,5,10], K = 3
//   Output: 84
//   Explanation: A becomes [15,15,15,9,10,10,10]
// Note:
//   1 <= K <= A.length <= 500
//   0 <= A[i] <= 10^6

// O(nk)
func maxSumAfterPartitioning(A []int, K int) int {
	// let dp(i) be the answer to split A[0...i].  
	dp := make([]int, len(A))
	m := 0
	for i:=0; i<K; i++ {
		m = max(m, A[i])
		dp[i] = m*(i+1)   // dp(i) = max(A[0...i])*(i+1)
	}
	for i:=K; i<len(A); i++ {
		m = 0
		for j:=i-1; i-j<=K; j-- {
			m = max(m, A[j+1])
			dp[i] = max(dp[i], dp[j]+m*(i-j))
		}
	}
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct{arr []int; K, ans int} {
		{[]int{1,15,7,9,2,5,10}, 3, 84},
	} {
		fmt.Println(maxSumAfterPartitioning(v.arr, v.K), v.ans)
	}
}
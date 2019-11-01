package main

import (
	"fmt"
	"math/rand"
)

// https://leetcode.com/problems/largest-sum-of-averages

// We partition a row of numbers A into at most K adjacent (non-empty) groups,
// then our score is the sum of the average of each group.
// What is the largest score we can achieve?
// Note that our partition must use every number in A,
// and that scores are not necessarily integers.
// Example:
//   Input:
//   A = [9,1,2,3,9]
//   K = 3
//   Output: 20
//   Explanation: The best choice is to partition A into [9], [1, 2, 3], [9].
//     The answer is 9 + (1 + 2 + 3) / 3 + 9 = 20.
//     We could have also partitioned A into [9, 1], [2], [3, 9], for example.
//     That partition would lead to a score of 5 + 2 + 6 = 13, which is worse.
// Note:
//   1 <= A.length <= 100.
//   1 <= A[i] <= 10000.
//   1 <= K <= A.length.
//   Answers within 10^-6 of the correct answer will be accepted as correct.

func largestSumOfAverages(A []int, K int) float64 {
	// let dp(i, k) be the largest sum of avgs of array A[i:],
	// then for each j>i, we split at j and the new sum is avg(A[i:j])+dp(j, k-1),
	// it's a candidate and we want to find the max.

	// prefix array to compute avg[i:j] = (prefix[j]-prefix[i])/(j-i)
	N := len(A)
	prefix := make([]int, N+1)
	for i := 1; i <= N; i++ {
		prefix[i] = prefix[i-1] + A[i-1]
	}
	// if K>=len(A), split to all single number, answer is sum(A)
	if K >= N {
		return float64(prefix[N])
	}

	dp := make([]float64, N*(K+1))
	for i := 0; i < N; i++ {
		// dp(i, 1) = avg(A[i:])
		set2d(dp, i, 1, K+1, float64(prefix[N]-prefix[i])/float64(N-i))
	}
	// computation order:
	// we already know dp(0...N-1, 1),
	// then compute dp(N-1, 2), dp(N-2, 2), ..., dp(0, 2)
	// then compute dp(N-1, 3), dp(N-2, 3), ..., dp(0, 3)
	// ..., finally,
	// then compute dp(N-1, K), dp(N-2, K), ..., dp(0, K), return dp(0, K)
	for k := 2; k <= K; k++ {
		for i := N - 1; i >= 0; i-- {
			if N-i < k {
				// no need to split A[i:N] to ks parts if N-i<k, just return the sum A[i:N]
				set2d(dp, i, k, K+1, float64(prefix[N]-prefix[i]))
				continue
			}
			maxavg := 0.0
			for j := i + 1; j < N; j++ {
				// a candidate, avg(A[i:j])+dp(j, k-1)
				curavg := float64(prefix[j]-prefix[i])/float64(j-i) + get2d(dp, j, k-1, K+1)
				if curavg > maxavg {
					maxavg = curavg
				}
			}
			set2d(dp, i, k, K+1, maxavg)
		}
	}
	return get2d(dp, 0, K, K+1)
}

func set2d(arr []float64, i, j, col int, v float64) {
	arr[i*col+j] = v
}

func get2d(arr []float64, i, j, col int) float64 {
	return arr[i*col+j]
}

func main() {
	fmt.Println(largestSumOfAverages([]int{9, 1, 2, 3, 9}, 3))       // 20
	fmt.Println(largestSumOfAverages([]int{1, 2, 3, 4, 5, 6, 7}, 4)) // 20.5
	fmt.Println(largestSumOfAverages([]int{4, 1, 7, 5, 6, 2, 3}, 4)) // 18.1667
	a := rand.Perm(20)
	fmt.Println(a)
	fmt.Println(largestSumOfAverages(a, 6))  // 75.75
	fmt.Println(largestSumOfAverages(a, 8))  // 100.1
	fmt.Println(largestSumOfAverages(a, 12)) // 141.46667
}

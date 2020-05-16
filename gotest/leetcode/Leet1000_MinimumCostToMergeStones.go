package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-cost-to-merge-stones/

// There are N piles of stones arranged in a row. The i-th pile has stones[i] stones. 
// A move consists of merging exactly K consecutive piles into one pile, and the
// cost of this move is equal to the total number of stones in these K piles. 
// Find the minimum cost to merge all piles of stones into one pile. If it is 
// impossible, return -1. 
// Example 1: 
//   Input: stones = [3,2,4,1], K = 2
//   Output: 20
//   Explanation: 
//     We start with [3, 2, 4, 1].
//     We merge [3, 2] for a cost of 5, and we are left with [5, 4, 1].
//     We merge [4, 1] for a cost of 5, and we are left with [5, 5].
//     We merge [5, 5] for a cost of 10, and we are left with [10].
//     The total cost was 20, and this is the minimum possible.
// Example 2: 
//   Input: stones = [3,2,4,1], K = 3
//   Output: -1
//   Explanation: After any merge operation, there are 2 piles left, 
//     and we can't merge anymore.  So the task is impossible. 
// Example 3: 
//   Input: stones = [3,5,1,2,6], K = 3
//   Output: 25
//   Explanation: 
//     We start with [3, 5, 1, 2, 6].
//     We merge [5, 1, 2] for a cost of 8, and we are left with [3, 8, 6].
//     We merge [3, 8, 6] for a cost of 17, and we are left with [17].
//     The total cost was 25, and this is the minimum possible.
// Note: 
//   1 <= stones.length <= 30 
//   2 <= K <= 30 
//   1 <= stones[i] <= 100 

// time O(k*n^2), space O(n^2)
func mergeStones(stones []int, K int) int {
	n := len(stones)
	if (n-1)%(K-1) != 0 {
		return -1
	}

	// prefix for accumulated sum of stones
	prefix := make([]int, n+1)
	for i:=1; i<=n; i++ {
		prefix[i] = prefix[i-1]+stones[i-1]
	}

	// initialize
	cost := make([][]int, n)
	for i := range cost {
		cost[i] = make([]int, n)
	}

	// dynamic programming, incrementing the diff between s[i...j], i.e., update diagonally
	// cost[i][j] means the cost to merge s[i...j] to (j-i)%(K-1)+1 piles
	// e.g., s=[1,2,3,4], K=3; cost([1,2,3,4]) means to merge to 2 parts.
	// cand1 = cost([1]) + cost([2,3,4])
	// cand2 = cost([1,2,3]) + cost([4])
	for diff:=K-1; diff<n; diff++ {   // the difference between i, j
		for i:=0; i<n && i+diff<n; i++ {
			j := i+diff
			cost[i][j] = 0x7fffffff
			for k:=i; k<j; k+=K-1 {
				cost[i][j] = min(cost[i][j], cost[i][k]+cost[k+1][j])
			}
			if (j-i)%(K-1)==0 {  // if we merge it to one pile, plus the sum
				cost[i][j] += prefix[j+1] - prefix[i]
			}
		}
	}
	for i:=range cost {
		fmt.Println(cost[i])
	}
	return cost[0][n-1]
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	fmt.Println(mergeStones([]int{1,2,2,4,5,6,7,8,9}, 3))
	fmt.Println(mergeStones([]int{3,2,4,1}, 2))
	fmt.Println(mergeStones([]int{3,2,4,1}, 3))
	fmt.Println(mergeStones([]int{3,5,1,2,6}, 3))
	fmt.Println(mergeStones([]int{3,4}, 2))
	fmt.Println(mergeStones([]int{3}, 2))
	fmt.Println(mergeStones([]int{1,2,3,2,3,4,7,8,2}, 3))
	fmt.Println(mergeStones([]int{6,4,4,6}, 2))
	fmt.Println(mergeStones([]int{1,2,3,4,5}, 3))
}
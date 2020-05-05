package main

import (
	"fmt"
)

// https://leetcode.com/problems/number-of-ways-to-wear-different-hats-to-each-other/

// There are n people and 40 types of hats labeled from 1 to 40.
// Given a list of list of integers hats, where hats[i] is a list of all hats
// preferred by the i-th person.
// Return the number of ways that the n people wear different hats to each other.
// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: hats = [[3,4],[4,5],[5]]
//   Output: 1
//   Explanation: There is only one way to choose hats given the conditions.
//     First person choose hat 3, Second person choose hat 4 and last one hat 5.
// Example 2:
//   Input: hats = [[3,5,1],[3,5]]
//   Output: 4
//   Explanation: There are 4 ways to choose hats (3,5), (5,3), (1,3) and (1,5)
// Example 3:
//   Input: hats = [[1,2,3,4],[1,2,3,4],[1,2,3,4],[1,2,3,4]]
//   Output: 24
//   Explanation: Each person can choose hats labeled from 1 to 4.
//     Number of Permutations of (1,2,3,4) = 24.
// Example 4:
//   Input: hats = [[1,2,3],[2,3,5,6],[1,3,7,9],[1,8,9],[2,5,7]]
//   Output: 111
// Constraints:
//   n == hats.length
//   1 <= n <= 10
//   1 <= hats[i].length <= 40
//   1 <= hats[i][j] <= 40
//   hats[i] contains a list of unique integers.

func numberWays(hats [][]int) int {
	dp := make([][]int, 40)
	h2p := make([][]int, 40)
	for i := range dp {
		dp[i] = make([]int, 1<<uint(len(hats)))
		h2p[i] = make([]int, 0)    // hat to people mapping
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
    for i := range hats {
    	for _, v := range hats[i] {
			h2p[v-1] = append(h2p[v-1], i)  // each hat can serve which people
		}
	}
	return dfs(dp, h2p, 1<<uint(len(hats))-1, 0)
}

// number of ways to wear people (in bit mask, e.g., people 0,1,3 = 0b1011) with hat[i:40]
func dfs(dp [][]int, h2p [][]int, people, i int) int {
	if people==0 {     // no people need to wear, found a valid solution
		return 1
	}
	if i==40 {         // no more hat, can't wear
		return 0
	}
	if dp[i][people] > -1 {
		return dp[i][people]
	}
	ans := dfs(dp, h2p, people, i+1)     // don't use i-th hat
	for _, j := range h2p[i] {           // let person j wear i-th hat
		if people & (1<<uint(j)) > 0 {   // person j in people
			ans += dfs(dp, h2p, people ^ (1<<uint(j)), i+1)   // no need to consider person j, set bit to 0
			ans %= 1000000007
		}
	}
	dp[i][people] = ans
	return ans
}

func main() {
	fmt.Println(numberWays([][]int{{3,4}, {4,5}, {5}}))
	fmt.Println(numberWays([][]int{{3,5,1}, {3,5}}))
	fmt.Println(numberWays([][]int{{1,2,3,4}, {1,2,3,4}, {1,2,3,4}, {1,2,3,4}}))
	fmt.Println(numberWays([][]int{{1,2,3}, {2,3,5,6}, {1,3,7,9}, {1,8,9}, {2,5,7}}))
	ss := make([][]int, 10)
	for i := range ss {
		ss[i] = make([]int, 40)
		for j := range ss[i] {
			ss[i][j] = j+1
		}
	}
	fmt.Println(numberWays(ss))
}
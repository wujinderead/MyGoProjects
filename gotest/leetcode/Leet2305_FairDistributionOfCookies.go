package main

import "fmt"

// https://leetcode.com/problems/fair-distribution-of-cookies/

// You are given an integer array cookies, where cookies[i] denotes the number of cookies in the
// iᵗʰ bag. You are also given an integer k that denotes the number of children to distribute all
// the bags of cookies to. All the cookies in the same bag must go to the same child and cannot be
// split up.
// The unfairness of a distribution is defined as the maximum total cookies obtained by a single
// child in the distribution.
// Return the minimum unfairness of all distributions.
// Example 1:
//   Input: cookies = [8,15,10,20,8], k = 2
//   Output: 31
//   Explanation: One optimal distribution is [8,15,8] and [10,20]
//     - The 1ˢᵗ child receives [8,15,8] which has a total of 8 + 15 + 8 = 31cookies.
//     - The 2ⁿᵈ child receives [10,20] which has a total of 10 + 20 = 30 cookies.
//     The unfairness of the distribution is max(31,30) = 31.
//     It can be shown that there is no distribution with an unfairness less than 31.
// Example 2:
//   Input: cookies = [6,1,3,2,2,4,1,2], k = 3
//   Output: 7
//   Explanation: One optimal distribution is [6,1], [3,2,2], and [4,1,2]
//     - The 1ˢᵗ child receives [6,1] which has a total of 6 + 1 = 7 cookies.
//     - The 2ⁿᵈ child receives [3,2,2] which has a total of 3 + 2 + 2 = 7 cookies.
//     - The 3ʳᵈ child receives [4,1,2] which has a total of 4 + 1 + 2 = 7 cookies.
//     The unfairness of the distribution is max(7,7,7) = 7.
//     It can be shown that there is no distribution with an unfairness less than 7.
// Constraints:
//   2 <= cookies.length <= 8
//   1 <= cookies[i] <= 10⁵
//   2 <= k <= cookies.length

func distributeCookies(cookies []int, k int) int {
	allMask := 1 << len(cookies)
	dp := make([]int, allMask)
	next := make([]int, allMask)
	M := int(1e6)
	for i := range dp { // initial dp[0][mask]=MAX_VALUE
		dp[i] = M
	}
	dp[0] = 0

	// sum[mask] = all cookies with respective mask
	sum := make([]int, allMask)
	for i := 1; i < allMask; i++ {
		for j := 0; j < len(cookies); j++ {
			if (1<<j)&i > 0 {
				sum[i] += cookies[j]
			}
		}
	}

	// for each submask of mask, dp[i][mask] = min(dp[i][mask], max(sum[submask], dp[i-1][mask-submask]))
	// i.e., give submask cookies to i-th person, give mask-submask cookies to remain i-1 persons.
	for x := 1; x <= k; x++ { // from 1 person tp k person
		for mask := 1; mask < allMask; mask++ {
			next[mask] = M                                                     // initial as MAX_VALUE
			for submask := mask; submask > 0; submask = (submask - 1) & mask { // for each submask of i
				next[mask] = min(next[mask], max(sum[submask], dp[mask-submask]))
			}
		}
		dp, next = next, dp
	}
	return dp[allMask-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		c      []int
		k, ans int
	}{
		{[]int{8, 15, 10, 20, 8}, 2, 31},
		{[]int{8, 15, 10, 20, 8}, 1, 61},
		{[]int{8, 15, 10, 20, 8}, 5, 20},
		{[]int{6, 1, 3, 2, 2, 4, 1, 2}, 3, 7},
		{[]int{4, 2, 1, 2, 2, 6, 3, 1}, 3, 7},
	} {
		fmt.Println(distributeCookies(v.c, v.k), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/

// There are n uniquely-sized sticks whose lengths are integers from 1 to n. You
// want to arrange the sticks such that exactly k sticks are visible from the left.
// A stick is visible from the left if there are no longer sticks to the left of it.
// For example, if the sticks are arranged [1,3,2,5,4], then the sticks with lengths
// 1, 3, and 5 are visible from the left.
// Given n and k, return the number of such arrangements. Since the answer may be large,
// return it modulo 10^9 + 7.
// Example 1:
//   Input: n = 3, k = 2
//   Output: 3
//   Explanation: [1,3,2], [2,3,1], and [2,1,3] are the only arrangements such that
//     exactly 2 sticks are visible. The visible sticks are underlined.
// Example 2:
//   Input: n = 5, k = 5
//   Output: 1
//   Explanation: [1,2,3,4,5] is the only arrangement such that all 5 sticks are visible.
//     The visible sticks are underlined.
// Example 3:
//   Input: n = 20, k = 11
//   Output: 647427950
//   Explanation: There are 647427950 (mod 10^9 + 7) ways to rearrange the sticks such
//     that exactly 11 sticks are visible.
// Constraints:
//   1 <= n <= 1000
//   1 <= k <= n

// let f(k,n) be the number of ways to show k visible sticks with n different numbers.
// initially,the numbers are A1 <= A2, ..., <= An, n sticks are visible.
// we can move x numbers to the right of An, these x numbers are hide by An,
// we need the n-x-1 numbers left of An to show k-1 visible sticks.
// f(1,n)=(n-1)!, f(n,n)=1
// time O(nkk)
func rearrangeSticks1(N int, K int) int {
	dp := make([][]int, K+1)
	mod := int(1e9 + 7)
	for k := 1; k <= K; k++ {
		dp[k] = make([]int, N+1)
	}
	dp[1][1] = 1
	for n := 2; n <= N; n++ { // f(1,n) = (n-1)!
		dp[1][n] = ((n - 1) * dp[1][n-1]) % mod
	}
	for k := 2; k <= K; k++ {
		dp[k][k] = 1
		// f(k,n) = f(k-1,i) * (n-1)! / i!    for k-1 <= i <= n-1
		for n := k + 1; n <= N; n++ {
			dp[k][n] += dp[k-1][n-1] // i==n-1
			x := n - 1
			for i := n - 2; i >= k-1; i-- {
				dp[k][n] += dp[k-1][i] * x
				dp[k][n] = dp[k][n] % mod
				x = (x * i) % mod
			}
		}
	}
	return dp[K][N]
}

// time O(nk) method:
// https://leetcode.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/discuss/1211169/JavaC%2B%2BPython-Concise-DP-Solution
// dp(n,k) is use n sticks to show k visible sticks
// case 1, longest stick at last position, we need n-1 sticks to see k-1 visible (excluding the last).
//   dp[n][k] += dp[n-1][k-1]
// case 2: longest not at last position, then n-1 choices for the last elements,
// we need then n-1 sticks to see k visible sticks (because last must be invisible)
//   dp[n][k] += dp[n-1][k] * (n-1)
func rearrangeSticks(n int, k int) int {
	dp := [1001][1001]int{}
	return f(&dp, n, k)
}

func f(dp *[1001][1001]int, n, k int) int {
	if n == k {
		return 1
	}
	if k == 0 {
		return 0
	}
	if dp[n][k] == 0 {
		dp[n][k] = (f(dp, n-1, k-1) + (n-1)*f(dp, n-1, k)) % 1000000007
	}
	return dp[n][k]
}

func main() {
	for _, v := range []struct {
		n, k, ans int
	}{
		{3, 2, 3},
		{5, 5, 1},
		{5, 1, 24},
		{6, 3, 225},
		{6, 4, 85},
		{20, 11, 647427950},
		{105, 20, 680986848},
	} {
		fmt.Println(rearrangeSticks(v.n, v.k), v.ans)
	}
}

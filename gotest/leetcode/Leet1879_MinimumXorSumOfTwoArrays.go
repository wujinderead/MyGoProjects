package main

import "fmt"

// https://leetcode.com/problems/minimum-xor-sum-of-two-arrays/

// You are given two integer arrays nums1 and nums2 of length n.
// The XOR sum of the two integer arrays is (nums1[0] XOR nums2[0]) + (nums1[1] XOR nums2[1]) + ... +
// (nums1[n - 1] XOR nums2[n - 1]) (0-indexed).
// For example, the XOR sum of [1,2,3] and [3,2,1] is equal to (1 XOR 3) + (2 XOR 2) + (3 XOR 1) =
// 2 + 0 + 2 = 4.
// Rearrange the elements of nums2 such that the resulting XOR sum is minimized.
// Return the XOR sum after the rearrangement.
// Example 1:
//   Input: nums1 = [1,2], nums2 = [2,3]
//   Output: 2
//   Explanation: Rearrange nums2 so that it becomes [3,2].
//     The XOR sum is (1 XOR 3) + (2 XOR 2) = 2 + 0 = 2.
// Example 2:
//   Input: nums1 = [1,0,3], nums2 = [5,3,4]
//   Output: 8
//   Explanation: Rearrange nums2 so that it becomes [5,4,3].
//     The XOR sum is (1 XOR 5) + (0 XOR 4) + (3 XOR 3) = 4 + 4 + 0 = 8.
// Constraints:
//   n == nums1.length
//   n == nums2.length
//   1 <= n <= 14
//   0 <= nums1[i], nums2[i] <= 10⁷

func minimumXORSum(nums1 []int, nums2 []int) int {
	// dp[i][mask] means use masked nums2 to XOR nums1[:i]
	// for each j-th bit in mask, dp[i-1][mask-(1<<j)] + nums1[i]^nums2[j]
	n := len(nums1)
	dp := make([][]int, n)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 1<<n)
		for j := range dp[i] {
			dp[i][j] = 1 << 31
		}
	}
	//recursive + memorial
	return dynamicProgramming(nums1, nums2, dp, n-1, (1<<n)-1)
}

func dynamicProgramming(nums1, nums2 []int, dp [][]int, i, mask int) int {
	if i == -1 {
		return 0
	}
	if dp[i][mask] != 1<<31 {
		return dp[i][mask]
	}
	for j := 0; j < len(nums1); j++ {
		if (1<<j)&mask > 0 { // j-th bit of mask is 1
			// use nums2[j] to XOR nums1[i], remain bits to XOR nums1[...i-1]
			cand := (nums1[i] ^ nums2[j]) + dynamicProgramming(nums1, nums2, dp, i-1, mask-(1<<j))
			if cand < dp[i][mask] {
				dp[i][mask] = cand
			}
		}
	}
	return dp[i][mask]
}

func main() {
	for _, v := range []struct {
		n1, n2 []int
		ans    int
	}{
		{[]int{1, 2}, []int{2, 3}, 2},
		{[]int{1, 0, 3}, []int{5, 3, 4}, 8},
		{[]int{3}, []int{1}, 2},
	} {
		fmt.Println(minimumXORSum(v.n1, v.n2), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/

// Given an array A of non-negative integers, return the maximum sum of elements in two
// non-overlapping (contiguous) subarrays, which have lengths L and M. (For clarification,
// the L-length subarray could occur before or after the M-length suarray.)
// Formally, return the largest V for which
//   V = (A[i] + A[i+1] + ... + A[i+L-1]) + (A[j] + A[j+1] + ... + A[j+M-1]) and either:
//   0 <= i < i + L - 1 < j < j + M - 1 < A.length, or
//   0 <= j < j + M - 1 < i < i + L - 1 < A.length.
// Example 1:
//   Input: A = [0,6,5,2,2,5,1,9,4], L = 1, M = 2
//   Output: 20
//   Explanation: One choice of subarrays is [9] with length 1, and [6,5] with length 2.
// Example 2:
//   Input: A = [3,8,1,3,2,1,8,9,0], L = 3, M = 2
//   Output: 29
//   Explanation: One choice of subarrays is [3,8,1] with length 3, and [8,9] with length 2.
// Example 3:
//   Input: A = [2,1,5,6,0,9,5,0,3,8], L = 4, M = 3
//   Output: 31
//   Explanation: One choice of subarrays is [5,6,0,9] with length 4, and [3,8] with length 3.
// Note:
//   L >= 1
//   M >= 1
//   L + M <= A.length <= 1000
//   0 <= A[i] <= 1000

// https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/discuss/278251/JavaC%2B%2BPython-O(N)Time-O(1)-Space
func maxSumTwoNoOverlap(A []int, L int, M int) int {
	B := make([]int, len(A))    // B is the accumulate sum of A
	copy(B, A)
	for i:=1; i<len(B); i++ {
		B[i] += B[i-1]
	}
	res := B[L + M - 1]
	Lmax := B[L - 1]   // Lmax: max sum of contiguous L elements before the last M elements
	Mmax := B[M - 1]   // Mmax: max sum of contiguous M elements before the last L elements
	for i:=L+M; i<len(B); i++ {
		// initially, last M elements before i is A[L: L+M] = B[i]-B[i-M],
		// the max L elements before M is A[0: L] = B[L-1].
		// we increment i, last M elements before i is A[L+1: L+M+1] = B[i]-B[i-M],
		// the max L elements before M has a new candidate: A[1: L+1] = B[i-M]-B[i-L-M]
		Lmax = max(Lmax, B[i-M] - B[i-L-M])
		Mmax = max(Mmax, B[i-L] - B[i-L-M])
		res = max(res, max(Lmax + B[i]-B[i-M], Mmax+B[i]-B[i-L]))
	}
	return res
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSumTwoNoOverlap([]int{0,6,5,2,2,5,1,9,4}, 1, 2))
	fmt.Println(maxSumTwoNoOverlap([]int{3,8,1,3,2,1,8,9,0}, 3, 2))
	fmt.Println(maxSumTwoNoOverlap([]int{2,1,5,6,0,9,5,0,3,8}, 4, 3))
}
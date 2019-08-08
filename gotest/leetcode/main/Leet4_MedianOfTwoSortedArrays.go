package main

import "fmt"

// https://leetcode.com/problems/median-of-two-sorted-arrays/

// There are two sorted arrays A and B of size m and n respectively.
// Find the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).
func findMedianSortedArrays(A []int, B []int) float64 {
	// to find median of two sorted arrays, we need to part the arrays as
	// A[0...i-1] | A[i...m-1]
	// B[0...j-1] | B[j...n-1]
	// the two parts should have the same length, i.e., i+j==(m+n)/2,
	// and A[i-1]≤B[j] and B[j-1]≤A[i]
	// then we can use binary search
	if len(A) > len(B) {
		A, B = B, A // need to search the shorter array, so that for 0≤i≤m,
	}
	m, n := len(A), len(B)
	min := 0
	max := m
	for min <= max {
		i := (max + min) / 2
		j := (m+n+1)/2 - i
		if i > 0 && A[i-1] > B[j] {
			// i too large, decrease i, j will increase, A[i-1] will decrease, B[j] will increase
			max = i - 1
		} else if i < m && B[j-1] > A[i] {
			// j too large, decrease j, i will increase, A[i] will increase, B[j-1] will decrease
			min = i + 1
		} else {
			// got perfect i
			maxLeft := 0
			if i == 0 {
				maxLeft = B[j-1]
			} else if j == 0 {
				maxLeft = A[i-1]
			} else {
				maxLeft = intmax(A[i-1], B[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			if i == m {
				minRight = B[j]
			} else if j == n {
				minRight = A[i]
			} else {
				minRight = intmin(B[j], A[i])
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0
}

func intmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intmin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{7}, []int{1, 2, 3, 4, 5, 6}))
	fmt.Println(findMedianSortedArrays([]int{7}, []int{1, 2, 3, 4, 5}))
	fmt.Println(findMedianSortedArrays([]int{2}, []int{1, 3, 4, 5}))
}

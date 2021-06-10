package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-absolute-sum-difference/

// You are given two positive integer arrays nums1 and nums2, both of length n.
// The absolute sum difference of arrays nums1 and nums2 is defined as the sum of
// |nums1[i] - nums2[i]| for each 0 <= i < n (0-indexed).
// You can replace at most one element of nums1 with any other element in nums1 to
// minimize the absolute sum difference.
// Return the minimum absolute sum difference after replacing at most one element
// in the array nums1. Since the answer may be large, return it modulo 10^9 + 7.
// |x| is defined as:
//   x if x >= 0, or
//   -x if x < 0.
// Example 1:
//   Input: nums1 = [1,7,5], nums2 = [2,3,5]
//   Output: 3
//   Explanation: There are two possible optimal solutions:
//     - Replace the second element with the first: [1,7,5] => [1,1,5], or
//     - Replace the second element with the third: [1,7,5] => [1,5,5].
//     Both will yield an absolute sum difference of |1-2| + (|1-3| or |5-3|) + |5-5| = 3.
// Example 2:
//   Input: nums1 = [2,4,6,8,10], nums2 = [2,4,6,8,10]
//   Output: 0
//   Explanation: nums1 is equal to nums2 so no replacement is needed. This will result in an
//     absolute sum difference of 0.
// Example 3:
//   Input: nums1 = [1,10,4,4,2,7], nums2 = [9,3,5,1,7,4]
//   Output: 20
//   Explanation: Replace the first element with the second: [1,10,4,4,2,7] => [10,10,4,4,2,7].
//     This yields an absolute sum difference of
//     |10-9| + |10-3| + |4-5| + |4-1| + |2-7| + |7-4| = 20
// Constraints:
//   n == nums1.length
//   n == nums2.length
//   1 <= n <= 10^5
//   1 <= nums1[i], nums2[i] <= 10^5

// for each nums2[i], binary search in nums1 to find the nearest number
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	const mod = int(1e9 + 7)
	sorted := make([]int, len(nums1))
	copy(sorted, nums1)
	sort.Sort(sort.IntSlice(sorted))
	sum := 0
	max := 0 // the max decrease we can get
	diff := make([]int, len(sorted))
	for i := range nums1 {
		sum = (sum + abs(nums1[i]-nums2[i])) % mod
		// search in sorted nums1, find the closest value to nums2[i]
		target := nums2[i]
		l, r := 0, len(nums2)-1
		// termination logic:
		// if sorted[-1]<target, l=r=len(sorted-1); else
		// l=r is the first index that makes sorted[index] >= target
		for l < r {
			mid := (l + r) / 2
			if sorted[mid] < target {
				l = mid + 1
			} else {
				r = mid
			}
		}
		// l or l-1 is the potential closest to target
		closest := nums1[i]
		if abs(sorted[l]-target) < abs(closest-target) {
			closest = sorted[l]
		}
		if l-1 > 0 && abs(sorted[l-1]-target) < abs(closest-target) {
			closest = sorted[l-1]
		}
		diff[i] = abs(closest - target)
		decrease := abs(nums1[i]-target) - abs(closest-target)
		if decrease > max {
			max = decrease
		}
	}
	return (sum - max + mod) % mod
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	a1 := []int{53, 48, 14, 71, 31, 55, 6, 80, 28, 19, 15, 40, 7, 21, 69, 15, 5, 42, 86, 15, 11, 54, 44, 62, 9, 100, 2, 26, 81, 87, 87, 18, 45, 29, 46, 100, 20, 87, 49, 86, 14, 74, 74, 52, 52, 60, 8, 25, 21, 96, 7, 90, 91, 42, 32, 34, 55, 20, 66, 36, 64, 67, 44, 51, 4, 46, 25, 57, 84, 23, 10, 84, 99, 33, 51, 28, 59, 88, 50, 41, 59, 69, 59, 65, 78, 50, 78, 50, 39, 91, 44, 78, 90, 83, 55, 5, 74, 96, 77, 46}
	a2 := []int{39, 49, 64, 34, 80, 26, 44, 3, 92, 46, 27, 88, 73, 55, 66, 10, 4, 72, 19, 37, 40, 49, 40, 58, 82, 32, 36, 91, 62, 21, 68, 65, 66, 55, 44, 24, 78, 56, 12, 79, 38, 53, 36, 90, 40, 73, 92, 14, 73, 89, 28, 53, 52, 46, 84, 47, 51, 31, 53, 22, 24, 14, 83, 75, 97, 87, 66, 42, 45, 98, 29, 82, 41, 36, 57, 95, 100, 2, 71, 34, 43, 50, 66, 52, 6, 43, 94, 71, 93, 61, 28, 84, 7, 79, 23, 48, 39, 27, 48, 79}
	for _, v := range []struct {
		n1, n2 []int
		ans    int
	}{
		{[]int{1, 7, 5}, []int{2, 3, 5}, 3},
		{[]int{2}, []int{2}, 0},
		{[]int{2}, []int{6}, 4},
		{[]int{1, 6}, []int{4, 5}, 3},
		{[]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10}, 0},
		{a1, a2, 3156},
	} {
		fmt.Println(minAbsoluteSumDiff(v.n1, v.n2), v.ans)
	}
}

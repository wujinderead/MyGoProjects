package main

import "fmt"

// https://leetcode.com/problems/number-of-pairs-satisfying-inequality/

// You are given two 0-indexed integer arrays nums1 and nums2, each of size n, and an integer diff.
// Find the number of pairs (i, j) such that:
//   0 <= i < j <= n - 1 and
//   nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff.
// Return the number of pairs that satisfy the conditions.
// Example 1:
//   Input: nums1 = [3,2,5], nums2 = [2,2,1], diff = 1
//   Output: 3
//   Explanation:
//     There are 3 pairs that satisfy the conditions:
//     1. i = 0, j = 1: 3 - 2 <= 2 - 2 + 1. Since i < j and 1 <= 1, this pair satisfies the conditions.
//     2. i = 0, j = 2: 3 - 5 <= 2 - 1 + 1. Since i < j and -2 <= 2, this pair satisfies the conditions.
//     3. i = 1, j = 2: 2 - 5 <= 2 - 1 + 1. Since i < j and -3 <= 2, this pair satisfies the conditions.
//     Therefore, we return 3.
// Example 2:
//   Input: nums1 = [3,-1], nums2 = [-2,2], diff = -1
//   Output: 0
//   Explanation:
//     Since there does not exist any pair that satisfies the conditions, we return 0.
// Constraints:
//   n == nums1.length == nums2.length
//   2 <= n <= 10⁵
//   -10⁴ <= nums1[i], nums2[i] <= 10⁴
//   -10⁴ <= diff <= 10⁴

// let d[x] = nums1[x]-nums2[x]
// nums1[i] - nums1[j] <= nums2[i] - nums2[j] + diff  --->
// (nums1[i]-nums2[i]) - (nums1[j]-nums2[j]) <= diff  --->
// d[i] <= d[j]+diff
func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
	d := make([]int, len(nums1))
	buf := make([]int, len(d))
	count := new(int)
	for i := range d {
		d[i] = nums1[i] - nums2[i]
	}
	mergeSort(d, buf, count, 0, len(d)-1, diff)
	return int64(*count)
}

func mergeSort(d, buf []int, count *int, start, end, diff int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	mergeSort(d, buf, count, start, mid, diff)
	mergeSort(d, buf, count, mid+1, end, diff)
	buf = buf[:0]
	i, j := start, mid+1
	// the logic:s left part and right part are sorted, so if d[i]<=d[j]+diff, then
	// for all j<=x<=end, d[i] <= d[x]+diff are satisfied. so it contribute end-j+1 to count.
	for i <= mid && j <= end {
		if d[i] <= d[j]+diff {
			*count += end - j + 1
			i++
		} else {
			j++
		}
	}
	for i <= mid {
		*count += end - j + 1
		i++
	}
	// sort the array by merge two sorted list
	i, j = start, mid+1
	for i <= mid && j <= end {
		if d[j] < d[i] {
			buf = append(buf, d[j])
			j++
		} else {
			buf = append(buf, d[i])
			i++
		}
	}
	for i <= mid {
		buf = append(buf, d[i])
		i++
	}
	for j <= end {
		buf = append(buf, d[j])
		j++
	}
	copy(d[start:end+1], buf)
}

func main() {
	for _, v := range []struct {
		nums1 []int
		nums2 []int
		diff  int
		ans   int64
	}{
		{[]int{3, 2, 5}, []int{2, 2, 1}, 1, 3},
		{[]int{3, -1}, []int{-2, 2}, -1, 0},
		{[]int{-58, 100, -93, -70, 42, -32, 10, 60, 12, 26, 14, -2, -5, 6, 87, 15, -54, -75, -95, -53, 73, -7, -58, 26, -12, 63, -57, -81, -34, -85, -57, 70, 40, -81, 34, -79, -42, -96, -29, 11, -92, 81, -97, -18, -41, 40, 8, 44, 53, 28, -50, 37, 0, 85, -67, -37, 82, 67, -65, 62, 15, -11, 28, -44, 74, 31, 71, 94, 100, -10, -8, -14, -41, -23, 62, 53, -37, -7, 15, -4, -93}, []int{-29, -4, 17, 81, -58, 13, 44, -62, -2, -76, -16, 82, 41, 46, 23, -7, -62, 79, -45, -24, 62, -7, 52, 49, -9, 52, 42, 53, -30, -25, -16, 0, -72, 34, 99, -49, -56, -2, 19, -94, -31, -81, -47, -24, 81, 35, 60, -6, 87, -74, 61, 38, -75, 84, -15, -96, 8, -6, 19, -44, -4, 11, -49, -5, 42, 64, -72, -23, 9, -45, 21, -24, -35, -81, -28, -63, -79, 44, 78, 32, 88}, -44, 1314},
	} {
		fmt.Println(numberOfPairs(v.nums1, v.nums2, v.diff), v.ans)
	}
}

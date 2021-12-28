package main

import "fmt"

// https://leetcode.com/problems/minimum-swaps-to-make-sequences-increasing/

// You are given two integer arrays of the same length nums1 and nums2. In one operation,
// you are allowed to swap nums1[i] with nums2[i].
// For example, if nums1 = [1,2,3,8], and nums2 = [5,6,7,4], you can swap the element at i = 3
// to obtain nums1 = [1,2,3,4] and nums2 = [5,6,7,8].
// Return the minimum number of needed operations to make nums1 and nums2 strictly increasing.
// The test cases are generated so that the given input always makes it possible.
// An array arr is strictly increasing if and only if arr[0] < arr[1] < arr[2] < ... < arr[arr.length - 1].
// Example 1:
//   Input: nums1 = [1,3,5,4], nums2 = [1,2,3,7]
//   Output: 1
//   Explanation:
//     Swap nums1[3] and nums2[3]. Then the sequences are:
//     nums1 = [1, 3, 5, 7] and nums2 = [1, 2, 3, 4]
//     which are both strictly increasing.
// Example 2:
//   Input: nums1 = [0,3,5,8,9], nums2 = [2,1,4,6,9]
//   Output: 1
// Constraints:
//   2 <= nums1.length <= 10⁵
//   nums2.length == nums1.length
//   0 <= nums1[i], nums2[i] <= 2 * 10⁵

// for nums1[0...i-1] and nums2[0...i-1], we can compute the min operation to make them increasing,
// with only 2 states: either we swapped at index i-1, or NOT.
// then for index i, we can decide if we swap or NOT, that will have 4 states:
// i-1 swapped, i swapped     swap[i]=swap[i-1]+1
// i-1 swapped, i not         no_swap[i]=swap[i-1]
// i-1 not, i swapped         swap[i]=no_swap[i-1]+1
// i-1 not, i not             no_swap[i]=no_swap[i]
// and we can reduce it to 2 states: either we swapped at index i+1, or NOT.
// thus we can solve this problem in a dp manner.
func minSwap(nums1 []int, nums2 []int) int {
	// two arrays below are min operations to make nums1[0...i] and nums2[0...i] increasing,
	// sw[i]: when nums1[i] and nums2[i] swapped
	// no[i]: when nums1[i] and nums2[i] not swapped
	n := len(nums1)
	max := int(1e6)
	sw, no := make([]int, n), make([]int, n)
	sw[0] = 1
	no[0] = 0
	var sw1, sw2, no1, no2 int
	for i := 1; i < n; i++ {
		p1, p2, c1, c2 := nums1[i-1], nums2[i-1], nums1[i], nums2[i]
		// no swap i-1, no swap i
		// swap i-1, swap i
		if p1 < c1 && p2 < c2 {
			no1 = no[i-1]
			sw1 = sw[i-1] + 1
		} else {
			no1 = max
			sw1 = max
		}
		// swap i-1, no swap i
		// no swap i-1, swap i
		if p2 < c1 && p1 < c2 {
			no2 = sw[i-1]
			sw2 = no[i-1] + 1
		} else {
			no2 = max
			sw2 = max
		}
		no[i] = min(no1, no2)
		sw[i] = min(sw1, sw2)
	}
	return min(no[n-1], sw[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		n1, n2 []int
		ans    int
	}{
		{[]int{1, 3, 5, 4}, []int{1, 2, 3, 7}, 1},
		{[]int{0, 3, 5, 8, 9}, []int{2, 1, 4, 6, 9}, 1},
	} {
		fmt.Println(minSwap(v.n1, v.n2), v.ans)
	}
}

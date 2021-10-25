package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/

// You are given an integer array nums of 2 * n integers. You need to partition nums into two arrays
// of length n to minimize the absolute difference of the sums of the arrays. To partition nums, put
// each element of nums into one of the two arrays.
// Return the minimum possible absolute difference.
// Example 1:
//   Input: nums = [3,9,7,3]
//   Output: 2
//   Explanation: One optimal partition is: [3,9] and [7,3]. The absolute difference between the sums
//     of the arrays is abs((3 + 9) - (7 + 3)) = 2.
// Example 2:
//   Input: nums = [-36,36]
//   Output: 72
//   Explanation: One optimal partition is: [-36] and [36].
//     The absolute difference between the sums of the arrays is abs((-36) - (36)) = 72.
// Example 3:
//   Input: nums = [2,-1,0,4,-2,-9]
//   Output: 0
//   Explanation: One optimal partition is: [2,4,-9] and [-1,0,-2].
//     The absolute difference between the sums of the arrays is abs((2 + 4 + -9) - (-1 + 0 + -2)) = 0.
// Constraints:
//   1 <= n <= 15
//   nums.length == 2 * n
//   -10^7 <= nums[i] <= 10^7

// meet-in-the-middle:
// for len(nums)=2n, split nums to 2 parts: len(left)=n and len(right)=n.
// for each part, there are 2^n subsets, so there are 2^n possible sums.
// get sum for each subset, and aggregate the subsets based on subset size.
// e.g. left[x]={l1,l2,...}, means we select x numbers from left parts,
// and all possible sums of these x numbers are {l1,l2,...}.
// then for each X in {l1,l2,...}, we need to find a value Y in right[n-x]={r1,r2,...},
// to make X+Y closest to sum(nums)/2 (by sorting right[n-x] and use binary search).
func minimumDifference(nums []int) int {
	n := len(nums) / 2
	sum := 0
	for i := range nums {
		nums[i] *= 2 // enlarge to avoid 0.5
		sum += nums[i]
	}
	target := sum / 2

	// get all sums for each part
	leftMaps := make([]map[int]struct{}, n+1)
	rightMaps := make([]map[int]struct{}, n+1)
	for i := 0; i < n+1; i++ {
		leftMaps[i] = make(map[int]struct{})
		rightMaps[i] = make(map[int]struct{})
	}
	for mask := 0; mask < (1 << n); mask++ {
		lsum := 0
		rsum := 0
		cnt := 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) > 0 {
				lsum += nums[i]
				rsum += nums[i+n]
				cnt++
			}
		}
		leftMaps[cnt][lsum] = struct{}{}
		rightMaps[cnt][rsum] = struct{}{}
	}

	// meet in the middle
	minabs := int(1e9)
	for i := 0; i < n+1; i++ {
		rnums := make([]int, 0, len(rightMaps[n-i]))
		for k := range rightMaps[n-i] {
			rnums = append(rnums, k)
		}
		sort.Sort(sort.IntSlice(rnums))
		for k := range leftMaps[i] { // for each k in left[i], find closest to target-k in right[n-i]
			ind := sort.SearchInts(rnums, target-k)
			if ind == 0 {
				if abs(k+rnums[0]-target) < minabs {
					minabs = abs(k + rnums[0] - target)
				}
				continue
			}
			if ind == len(rnums) {
				if abs(k+rnums[len(rnums)-1]-target) < minabs {
					minabs = abs(k + rnums[len(rnums)-1] - target)
				}
				continue
			}
			// for ind=sort.SearchInts(a, target), a[ind] is the first that >= target, a[ind-1]<target
			// a[ind] and a[ind-1] are candidates that closest to target
			if abs(k+rnums[ind]-target) < minabs {
				minabs = abs(k + rnums[ind] - target)
			}
			if abs(k+rnums[ind-1]-target) < minabs {
				minabs = abs(k + rnums[ind-1] - target)
			}
		}
	}
	// minabs is the minimal |sum(halfArray)-target|, so the difference between 2 half array sums is 2*minabs
	// as we have enlarged the original array. so the answer is minabs
	return minabs
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	for _, v := range []struct {
		n   []int
		ans int
	}{
		{[]int{3, 7, 3, 9}, 2},
		{[]int{-36, 36}, 72},
		{[]int{2, -1, 0, 4, -2, -9}, 0},
		{[]int{1, 2, 3, 4, 5, 6}, 1},
	} {
		fmt.Println(minimumDifference(v.n), v.ans)
	}
}

package main

import (
	"sort"
	"fmt"
)

// https://leetcode.com/problems/find-k-th-smallest-pair-distance/

// Given an integer array, return the k-th smallest distance among all the pairs.
// The distance of a pair (A, B) is defined as the absolute difference between A and B.
// Example 1:
//   Input:
//     nums = [1,3,1]
//     k = 1
//   Output: 0
//   Explanation:
//     Here are all the pairs:
//     (1,3) -> 2
//     (1,1) -> 0
//     (3,1) -> 2
//     Then the 1st smallest distance pair is (1,1), and its distance is 0.
// Note:
//   2 <= len(nums) <= 10000.
//   0 <= nums[i] < 1000000.
//   1 <= k <= len(nums) * (len(nums) - 1) / 2.

func smallestDistancePair(nums []int, k int) int {
    // sort numbers
    sort.Sort(sort.IntSlice(nums))

    // max diff is nums[n-1]-nums[0], min diff is the minimal of nums[i]-nums[i-1]
    lo, hi := 0x7fffffff, nums[len(nums)-1]-nums[0]
    for i:=1; i<len(nums); i++ {
    	if nums[i]-nums[i-1]<lo {
    		lo = nums[i]-nums[i-1]
		}
	}

	// binary search, time log(max-min)
	for lo<hi {
		mid := (lo+hi)/2
		// count distances <= mid, time nlogn
		// IMPROVEMENT: we do count in O(n) using two pointers
		count := 0
		for i:=1; i<len(nums); i++ {
			count += countLessAndEqual(nums[i:], mid+nums[i-1])
		}
		//fmt.Println("lo:", lo, "hi:", hi, "m:", mid, "co:", count, "k:", k)
		if count<k {
			lo = mid+1
		} else {
			hi = mid
		}
	}
    return lo
}

func countLessAndEqual(nums []int, n int) int {
	if n>=nums[len(nums)-1] {
		return len(nums)
	}
	// find first nums[i] that > n, return i
	lo, hi := 0, len(nums)-1
	for lo<hi {
		mid := (lo+hi)/2
		if nums[mid]<=n {
			lo = mid+1
		} else {
			hi = mid
		}
	}
	return lo
}

func main() {
	for _, v := range [][]int{
		{1,3,1},
		{1,3,4,7},
		{1,2},
		{8,12,21,26,29},
	} {
		diff := make([]int, 0, len(v)*(len(v)-1)/2)
		for i:=1; i<len(v); i++ {
			for j:=0; j<i; j++ {
				abs := v[i]-v[j]
				if abs<0 {
					abs = -abs
				}
				diff = append(diff, abs)
			}
		}
		sort.Sort(sort.IntSlice(diff))
		for i:=diff[0]-1; i<=diff[len(diff)-1]+1; i++ {
			fmt.Println(diff, "<=", i, "=", countLessAndEqual(diff, i))
		}
		for i:=1; i<=len(v)*(len(v)-1)/2; i++ {
			fmt.Println(smallestDistancePair(v, i), diff[i-1])
		}
	}
}
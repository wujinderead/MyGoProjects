package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/

// Given the array nums consisting of n positive integers. You computed the sum of all non-empty 
// continous subarrays from the array and then sort them in non-decreasing order, creating a 
// new array of n * (n + 1) / 2 numbers. Return the sum of the numbers from index left to index right 
// (indexed from 1), inclusive, in the new array. Since the answer can be a huge number 
// return it modulo 10^9 + 7.
// Example 1:
//   Input: nums = [1,2,3,4], n = 4, left = 1, right = 5
//   Output: 13 
//   Explanation: All subarray sums are 1, 3, 6, 10, 2, 5, 9, 3, 7, 4. After sorting them in 
//     non-decreasing order we have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. 
//     The sum of the numbers from index le = 1 to ri = 5 is 1 + 2 + 3 + 3 + 4 = 13. 
// Example 2:
//   Input: nums = [1,2,3,4], n = 4, left = 3, right = 4
//   Output: 6
//   Explanation: The given array is the same as example 1. 
//     We have the new array [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]. 
//     The sum of the numbers from index le = 3 to ri = 4 is 3 + 3 = 6.
// Example 3:
//   Input: nums = [1,2,3,4], n = 4, left = 1, right = 10
//   Output: 50
// Constraints:
//   1 <= nums.length <= 10^3
//   nums.length == n
//   1 <= nums[i] <= 100
//   1 <= left <= right <= n * (n + 1) / 2

// time O(Nlog(sum(nums))), space O(n)
func rangeSum(nums []int, n int, left int, right int) int {
	prefix := make([]int, n+1)  // prefix sum of nums
	prepre := make([]int, n+1)  // prefix sum of prefix
	for i:=range nums {
		prefix[i+1] = prefix[i] + nums[i]
		prepre[i+1] = prefix[i+1] + prepre[i]
	}
	return (sumKSmallest(prefix, prepre, right) - sumKSmallest(prefix, prepre, left-1)) % int(1e9+7)
}

// the sum of first k numbers in the sorted list of range sum.
func sumKSmallest(prefix, prepre []int, k int) int {
	val := getKth(prefix, k)
	i, res := 0, 0
	for j:=1; j<len(prefix); j++ {
		for prefix[j]-prefix[i] > val {
			i++
		}
		x := 0       // x := i-1>=0 ? prepre[i-1] : 0
		if i-1>=0 {
			x = prepre[i-1]
		}
		res += (j-i)*prefix[j] - (prepre[j-1]-x)   // it's easy to get by drawing graphs
	}
	return res - (countLess(prefix, val)-k)*val    // minus duplicated
}

// how many range sum <= target, by scan prefix.
func countLess(prefix []int, target int) int {
	i, count := 0, 0
	for j:=1; j<len(prefix); j++ {
		for prefix[j]-prefix[i] > target {
			i++
		}
		count += j-i
	}
	return count
}

// get kth range sum, by scan prefix, binary search.
func getKth(prefix []int, k int) int {
	l, r := 1, prefix[len(prefix)-1]
	for l < r {
		mid := l+(r-l)/2
		count := countLess(prefix, mid)
		if count < k {
			l = mid+1
		} else {
			r = mid
		}
	}
	return l
}

func verify(nums []int, n int, left int, right int) int {
	prefix := make([]int, n+1)
	for i:=range nums {
		prefix[i+1] = prefix[i]+nums[i]
	}
	sums := make([]int, 0, n*(n+1)/2)
	for i:=1; i<=n; i++ {
		for j:=0; j<i; j++ {
			sums = append(sums, prefix[i]-prefix[j])
		}
	}
	sort.Sort(sort.IntSlice(sums))
	ans := 0
	for i:=left-1; i<right; i++ {
		ans += sums[i]
	}
	return ans
}

func main() {
	for _, v := range [][]int{
		{3},
		{7,2},
		{1,2,3,4},
		{5,7,4},
		{3,1,4,2},
	} {
		n := len(v)
		for i:=1; i<=n*(n+1)/2; i++ {
			for j:=i; j<=n*(n+1)/2; j++ {
				fmt.Println(rangeSum(v, n, i, j), verify(v, n, i, j))
			}
		}
	}
	long := make([]int, 1000)
	for i := range long {
		long[i] = 100
	}
	fmt.Println(rangeSum(long, 1000, 1, 500500))
}
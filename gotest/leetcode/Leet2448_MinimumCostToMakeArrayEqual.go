package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-cost-to-make-array-equal/

// You are given two 0-indexed arrays nums and cost consisting each of n positive integers.
// You can do the following operation any number of times:
// Increase or decrease any element of the array nums by 1.
// The cost of doing one operation on the iᵗʰ element is cost[i].
// Return the minimum total cost such that all the elements of the array nums become equal.
// Example 1:
//   Input: nums = [1,3,5,2], cost = [2,3,1,14]
//   Output: 8
//   Explanation: We can make all the elements equal to 2 in the following way:
//     - Increase the 0th element one time. The cost is 2.
//     - Decrease the 1st element one time. The cost is 3.
//     - Decrease the 2nd element three times. The cost is 1 + 1 + 1 = 3.
//     The total cost is 2 + 3 + 3 = 8.
//     It can be shown that we cannot make the array equal with a smaller cost.
// Example 2:
//   Input: nums = [2,2,2,2,2], cost = [4,2,8,1,3]
//   Output: 0
//   Explanation: All the elements are already equal, so no operations are needed.
// Constraints:
//   n == nums.length == cost.length
//   1 <= n <= 1e5
//   1 <= nums[i], cost[i] <= 1e6

// O(nlogn): sort the array first
func minCost(nums []int, cost []int) int64 {
	pair := make([][2]int, len(nums))
	for i := range pair {
		pair[i] = [2]int{nums[i], cost[i]}
	}
	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] < pair[j][0]
	})
	sum := 0
	suffix := 0
	for i := range pair { // initial cost: reduce all numbers to 0
		sum += pair[i][0] * pair[i][1]
		suffix += pair[i][1]
	}
	min := sum
	prefix := 0
	prenum := 0
	for i := range pair { // make all numbers to nums[i]
		sum = sum - suffix*(pair[i][0]-prenum) + prefix*(pair[i][0]-prenum)
		suffix -= pair[i][1]
		prefix += pair[i][1]
		prenum = pair[i][0]
		if sum < min {
			min = sum
		}
	}
	return int64(min)
}

// binary search: let x be the equal number, f(x) be the total cost,
// then in the interval [min(nums), max(nums)], f(x) is a convex curve.
// binary search by f(mid) and f(mid+1) to decide the direction of shrink.
func minCostBinarySearch(nums []int, cost []int) int64 {
	left, right := 1, int(1e6)
	res := 0
	for left < right {
		mid := (left + right) / 2
		rmid := getCost(nums, cost, mid)
		rmid1 := getCost(nums, cost, mid+1)
		res = min(rmid, rmid1)
		if rmid > rmid1 {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return int64(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getCost(nums, cost []int, value int) int {
	sum := 0
	for i := range cost {
		sum += abs(nums[i]-value) * cost[i]
	}
	return sum
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	for _, v := range []struct {
		nums []int
		cost []int
		ans  int64
	}{
		{[]int{1, 3, 5, 2}, []int{2, 3, 1, 14}, 8},
		{[]int{2, 2, 2, 2, 2}, []int{4, 2, 8, 1, 3}, 0},
		{[]int{1, 4, 8, 15}, []int{1, 1, 1, 1}, 18},
	} {
		fmt.Println(minCost(v.nums, v.cost), v.ans)
		fmt.Println(minCostBinarySearch(v.nums, v.cost), v.ans)
	}
}

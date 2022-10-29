package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-number-of-operations-to-make-arrays-similar/

// You are given two positive integer arrays nums and target, of the same length.
// In one operation, you can choose any two distinct indices i and j where 0 <= i,j < nums.length and:
//   set nums[i] = nums[i] + 2 and
//   set nums[j] = nums[j] - 2.
// Two arrays are considered to be similar if the frequency of each element is the same.
// Return the minimum number of operations required to make nums similar to target. The test cases are
// generated such that nums can always be similar to target.
// Example 1:
//   Input: nums = [8,12,6], target = [2,14,10]
//   Output: 2
//   Explanation: It is possible to make nums similar to target in two operations:
//     - Choose i = 0 and j = 2, nums = [10,12,4].
//     - Choose i = 1 and j = 2, nums = [10,14,2].
//     It can be shown that 2 is the minimum number of operations needed.
// Example 2:
//   Input: nums = [1,2,5], target = [4,1,3]
//   Output: 1
//   Explanation: We can make nums similar to target in one operation:
//     - Choose i = 1 and j = 2, nums = [1,4,3].
// Example 3:
//   Input: nums = [1,1,1,1,1], target = [1,1,1,1,1]
//   Output: 0
//   Explanation: The array nums is already similiar to target.
// Constraints:
//   n == nums.length == target.length
//   1 <= n <= 10⁵
//   1 <= nums[i], target[i] <= 10⁶
//   It is possible to make nums similar to target.

func makeSimilar(nums []int, target []int) int64 {
	var oddN, evenN, oddT, evenT []int
	for i := range nums {
		if nums[i]%2 == 0 {
			evenN = append(evenN, nums[i])
		} else {
			oddN = append(oddN, nums[i])
		}
		if target[i]%2 == 0 {
			evenT = append(evenT, target[i])
		} else {
			oddT = append(oddT, target[i])
		}
	}
	// split to odd and event, then sort
	sort.Ints(oddN)
	sort.Ints(evenN)
	sort.Ints(oddT)
	sort.Ints(evenT)
	count := 0
	for i := range oddN {
		count += abs(oddT[i]-oddN[i]) / 2
	}
	for i := range evenT {
		count += abs(evenT[i]-evenN[i]) / 2
	}
	return int64(count) / 2 // need divide by 2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	for _, v := range []struct {
		nums   []int
		target []int
		ans    int64
	}{
		{[]int{1, 2, 5}, []int{4, 1, 3}, 1},
		{[]int{8, 12, 6}, []int{2, 14, 10}, 2},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}, 0},
		{
			[]int{758, 334, 402, 1792, 1112, 1436, 1534, 1702, 1538, 1427, 720, 1424, 114, 1604, 564, 120, 578},
			[]int{1670, 216, 1392, 1828, 1104, 464, 678, 1134, 644, 1178, 1150, 1608, 1799, 1156, 244, 2, 892},
			645,
		},
	} {
		fmt.Println(makeSimilar(v.nums, v.target), v.ans)
	}
}

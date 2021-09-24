package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/find-the-kth-largest-integer-in-the-array/

// You are given an array of strings nums and an integer k. Each string in nums represents
// an integer without leading zeros.
// Return the string that represents the kth largest integer in nums.
// Note: Duplicate numbers should be counted distinctly. For example, if nums is ["1","2","2"],
// "2" is the first largest integer, "2" is the second-largest integer, and "1" is the third-largest integer.
// Example 1:
//   Input: nums = ["3","6","7","10"], k = 4
//   Output: "3"
//   Explanation:
//     The numbers in nums sorted in non-decreasing order are ["3","6","7","10"].
//     The 4th largest integer in nums is "3".
// Example 2:
//   Input: nums = ["2","21","12","1"], k = 3
//   Output: "2"
//   Explanation:
//     The numbers in nums sorted in non-decreasing order are ["1","2","12","21"].
//     The 3rd largest integer in nums is "2".
// Example 3:
//   Input: nums = ["0","0"], k = 2
//   Output: "0"
//   Explanation:
//     The numbers in nums sorted in non-decreasing order are ["0","0"].
//     The 2nd largest integer in nums is "0".
// Constraints:
//   1 <= k <= nums.length <= 10^4
//   1 <= nums[i].length <= 100
//   nums[i] consists of only digits.
//   nums[i] will not have any leading zeros.

// just sort by string length; if same length, compare string
func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) != len(nums[j]) {
			return len(nums[i]) < len(nums[j])
		}
		for x := range nums[i] {
			if nums[i][x] != nums[j][x] {
				return nums[i][x] < nums[j][x]
			}
		}
		return false
	})
	return nums[len(nums)-k]
}

func main() {
	for _, v := range []struct {
		ns  []string
		k   int
		ans string
	}{
		{[]string{"3", "6", "7", "10"}, 1, "10"},
		{[]string{"3", "6", "7", "10"}, 2, "7"},
		{[]string{"3", "6", "7", "10"}, 3, "6"},
		{[]string{"3", "6", "7", "10"}, 4, "3"},
		{[]string{"2", "21", "12", "1"}, 3, "2"},
		{[]string{"0", "0"}, 2, "0"},
		{[]string{"423", "521", "2", "42"}, 2, "423"},
		{[]string{"423", "521", "2", "521", "423", "42"}, 2, "521"},
	} {
		fmt.Println(kthLargestNumber(v.ns, v.k), v.ans)
	}
}

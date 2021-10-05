package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/number-of-pairs-of-strings-with-concatenation-equal-to-target/

// Given an array of digit strings nums and a digit string target, return the number of pairs of
// indices (i, j) (where i != j) such that the concatenation of nums[i] + nums[j] equals target.
// Example 1:
//   Input: nums = ["777","7","77","77"], target = "7777"
//   Output: 4
//   Explanation: Valid pairs are:
//     - (0, 1): "777" + "7"
//     - (1, 0): "7" + "777"
//     - (2, 3): "77" + "77"
//     - (3, 2): "77" + "77"
// Example 2:
//   Input: nums = ["123","4","12","34"], target = "1234"
//   Output: 2
//   Explanation: Valid pairs are:
//     - (0, 1): "123" + "4"
//     - (2, 3): "12" + "34"
// Example 3:
//   Input: nums = ["1","1","1"], target = "11"
//   Output: 6
//   Explanation: Valid pairs are:
//     - (0, 1): "1" + "1"
//     - (1, 0): "1" + "1"
//     - (0, 2): "1" + "1"
//     - (2, 0): "1" + "1"
//     - (1, 2): "1" + "1"
//     - (2, 1): "1" + "1"
// Constraints:
//   2 <= nums.length <= 100
//   1 <= nums[i].length <= 100
//   2 <= target.length <= 100
//   nums[i] and target consist of digits.
//   nums[i] and target do not have leading zeros.

func numOfPairs(nums []string, target string) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if len(nums[i])+len(nums[j]) == len(target) {
				if strings.HasPrefix(target, nums[i]) && strings.HasSuffix(target, nums[j]) {
					count++
				}
				if strings.HasPrefix(target, nums[j]) && strings.HasSuffix(target, nums[i]) {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	for _, v := range []struct {
		n   []string
		t   string
		ans int
	}{
		{[]string{"777", "7", "77", "77"}, "7777", 4},
		{[]string{"123", "4", "12", "34"}, "1234", 2},
		{[]string{"1", "1", "1"}, "11", 6},
	} {
		fmt.Println(numOfPairs(v.n, v.t), v.ans)
	}
}

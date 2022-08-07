package main

import "fmt"

// https://leetcode.com/problems/count-number-of-bad-pairs/

// You are given a 0-indexed integer array nums. A pair of indices (i, j) is a bad pair if
// i < j and j - i != nums[j] - nums[i].
// Return the total number of bad pairs in nums.
// Example 1:
//   Input: nums = [4,1,3,3]
//   Output: 5
//   Explanation: The pair (0, 1) is a bad pair since 1 - 0 != 1 - 4.
//     The pair (0, 2) is a bad pair since 2 - 0 != 3 - 4, 2 != -1.
//     The pair (0, 3) is a bad pair since 3 - 0 != 3 - 4, 3 != -1.
//     The pair (1, 2) is a bad pair since 2 - 1 != 3 - 1, 1 != 2.
//     The pair (2, 3) is a bad pair since 3 - 2 != 3 - 3, 1 != 0.
//     There are a total of 5 bad pairs, so we return 5.
// Example 2:
//   Input: nums = [1,2,3,4,5]
//   Output: 0
//   Explanation: There are no bad pairs.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i] <= 10⁹

// `nums[i]-nums[j] = i-j` equals to `nums[i]-i = nums[j]-j`
func countBadPairs(nums []int) int64 {
	all := len(nums) * (len(nums) - 1) / 2
	good := 0
	mapp := make(map[int]int) // map to store the count of nums[i]-i
	for i, v := range nums {
		good += mapp[v-i]
		mapp[v-i] += 1
	}
	return int64(all - good)
}

func main() {
	for _, v := range []struct {
		nums []int
		ans  int
	}{
		{[]int{4, 1, 3, 3}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	} {
		fmt.Println(countBadPairs(v.nums), v.ans)
	}
}

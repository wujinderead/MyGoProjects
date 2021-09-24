package main

import "fmt"

// https://leetcode.com/problems/sum-of-beauty-in-the-array/

// You are given a 0-indexed integer array nums. For each index i (1 <= i <= nums.length - 2)
// the beauty of nums[i] equals:
//   2, if nums[j] < nums[i] < nums[k], for all 0 <= j < i and for all i < k <= nums.length - 1.
//   1, if nums[i - 1] < nums[i] < nums[i + 1], and the previous condition is not satisfied.
//   0, if none of the previous conditions holds.
// Return the sum of beauty of all nums[i] where 1 <= i <= nums.length - 2.
// Example 1:
//   Input: nums = [1,2,3]
//   Output: 2
//   Explanation: For each index i in the range 1 <= i <= 1:
//     - The beauty of nums[1] equals 2.
// Example 2:
//   Input: nums = [2,4,6,4]
//   Output: 1
//   Explanation: For each index i in the range 1 <= i <= 2:
//     - The beauty of nums[1] equals 1.
//     - The beauty of nums[2] equals 0.
// Example 3:
//   Input: nums = [3,2,1]
//   Output: 0
//   Explanation: For each index i in the range 1 <= i <= 1:
//     - The beauty of nums[1] equals 0.
// Constraints:
//   3 <= nums.length <= 10^5
//   1 <= nums[i] <= 10^5

// just compute prefix max and suffix min
func sumOfBeauties(nums []int) int {
	left := make([]int, len(nums))  // left[i] = max(nums[0...i])
	right := make([]int, len(nums)) // right[i] = min(nums[i...])
	left[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1]
		if nums[i] > left[i] {
			left[i] = nums[i]
		}
	}
	right[len(nums)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1]
		if nums[i] < right[i] {
			right[i] = nums[i]
		}
	}
	sum := 0
	for i := 1; i < len(nums)-1; i++ {
		if left[i-1] < nums[i] && nums[i] < right[i+1] {
			sum += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			sum += 1
		}
	}
	return sum
}

func main() {
	for _, v := range []struct {
		n   []int
		ans int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{2, 4, 6, 4}, 1},
		{[]int{3, 2, 1}, 0},
	} {
		fmt.Println(sumOfBeauties(v.n), v.ans)
	}
}

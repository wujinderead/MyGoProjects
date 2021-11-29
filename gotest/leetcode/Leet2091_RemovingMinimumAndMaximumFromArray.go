package main

import "fmt"

// https://leetcode.com/problems/removing-minimum-and-maximum-from-array/

// You are given a 0-indexed array of distinct integers nums.
// There is an element in nums that has the lowest value and an element that has the highest value.
// We call them the minimum and maximum respectively. Your goal is to remove both these elements from the array.
// A deletion is defined as either removing an element from the front of the array or removing an element from
// the back of the array.
// Return the minimum number of deletions it would take to remove both the minimum and maximum element from the array.
// Example 1:
//   Input: nums = [2,10,7,5,4,1,8,6]
//   Output: 5
//   Explanation:
//   The minimum element in the array is nums[5], which is 1.
//   The maximum element in the array is nums[1], which is 10.
//   We can remove both the minimum and maximum by removing 2 elements from the front and 3 elements from the back.
//   This results in 2 + 3 = 5 deletions, which is the minimum number possible.
// Example 2:
//   Input: nums = [0,-4,19,1,8,-2,-3,5]
//   Output: 3
//   Explanation:
//     The minimum element in the array is nums[1], which is -4.
//     The maximum element in the array is nums[2], which is 19.
//     We can remove both the minimum and maximum by removing 3 elements from the front.
//     This results in only 3 deletions, which is the minimum number possible.
// Example 3:
//   Input: nums = [101]
//   Output: 1
//   Explanation:
//     There is only one element in the array, which makes it both the minimum and maximum element.
//     We can remove it with 1 deletion.
// Constraints:
//   1 <= nums.length <= 10⁵
//   -10⁵ <= nums[i] <= 10⁵
//   The integers in nums are distinct.

func minimumDeletions(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	mini, maxi := 0, 0
	for i := range nums {
		if nums[i] < nums[mini] {
			mini = i
		}
		if nums[i] > nums[maxi] {
			maxi = i
		}
	}
	if mini > maxi { // assume mini < maxi
		maxi, mini = mini, maxi
	}
	// x x mini y y maxi z z, we can delete nums[0...maxi] or nums[mini...] or union(nums[0...mini], nums[maxi...])
	ans := maxi + 1
	if len(nums)-mini < ans {
		ans = len(nums) - mini
	}
	if len(nums)-maxi+mini+1 < ans {
		ans = len(nums) - maxi + mini + 1
	}
	return ans
}

func main() {
	for _, v := range []struct {
		n   []int
		ans int
	}{
		{[]int{2, 10, 7, 5, 4, 1, 8, 6}, 5},
		{[]int{0, -4, 19, 1, 8, -2, -3, 5}, 3},
		{[]int{101}, 1},
	} {
		fmt.Println(minimumDeletions(v.n), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/number-of-subarrays-with-gcd-equal-to-k/

// Given an integer array nums and an integer k, return the number of subarrays of nums where the
// greatest common divisor of the subarray's elements is k.
// A subarray is a contiguous non-empty sequence of elements within an array.
// The greatest common divisor of an array is the largest integer that evenly
// divides all the array elements.
// Example 1:
//   Input: nums = [9,3,1,2,6,3], k = 3
//   Output: 4
//   Explanation: The subarrays of nums where 3 is the greatest common divisor of
//     all the subarray's elements are:
//     - [9,3,1,2,6,3]
//     - [9,3,1,2,6,3]
//     - [9,3,1,2,6,3]
//     - [9,3,1,2,6,3]
// Example 2:
// Input: nums = [4], k = 7
// Output: 0
// Explanation: There are no subarrays of nums where 7 is the greatest common
//   divisor of all the subarray's elements.
// Constraints:
//   1 <= nums.length <= 1000
//   1 <= nums[i], k <= 10⁹

// O(N^2): for substring start at nums[i], check the gcd is always k
func subarrayGCD(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%k != 0 {
			continue
		}
		if nums[i] == k {
			count++
		}
		gc := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j]%k != 0 {
				break
			}
			gc = gcd(gc, nums[j])
			if gc == k {
				count++
			}
		}
	}
	return count
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func main() {
	for _, v := range []struct {
		nums []int
		k    int
		ans  int
	}{
		{[]int{9, 3, 1, 2, 6, 3}, 3, 4},
		{[]int{4}, 7, 0},
		{[]int{12, 6, 9, 2, 3, 6, 12, 6}, 3, 6},
	} {
		fmt.Println(subarrayGCD(v.nums, v.k), v.ans)
	}
}

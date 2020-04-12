package main

import "fmt"

// https://leetcode.com/problems/reverse-subarray-to-maximize-array-value/

// You are given an integer array nums. The value of this array is defined as the
// sum of |nums[i]-nums[i+1]| for all 0 <= i < nums.length-1.
// You are allowed to select any subarray of the given array and reverse it. You
// can perform this operation only once.
// Find maximum possible value of the final array.
// Example 1:
//   Input: nums = [2,3,1,5,4]
//   Output: 10
//   Explanation: By reversing the subarray [3,1,5] the array becomes [2,5,1,3,4]
//     whose value is 10.
// Example 2:
//   Input: nums = [2,4,9,24,2,1,10]
//   Output: 68
// Constraints:
//   1 <= nums.length <= 3*10^4
//   -10^5 <= nums[i] <= 10^5

// IMPROVEMENT, O(n) solution, utilize some math about absolute difference:
// https://leetcode.com/problems/reverse-subarray-to-maximize-array-value/discuss/489882/O(n)-Solution-with-explanation

func maxValueAfterReverse(nums []int) int {
	// for array [A0, ..., Ai-1, Ai, ..., Aj, Aj+1, ..., An], if we reverse A[i...j],
	// the difference between A[0...i-1], A[i...j], A[j+1...n] is unchanged.
	// the only change is that |A[i]-A[i-1]| + |A[j+1]-A[j]| become |A[j]-A[i-1]| + |A[j+1]-A[i]|
	// we want this change is the most large.
	// time complexity O(n²)
	sum := 0
	for i := 1; i < len(nums); i++ {
		sum += abs(nums[i] - nums[i-1])
	}
	maxdiff := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ { // for each pair
			ai, aj := nums[i], nums[j]
			t1, t2, t3, t4 := 0, 0, 0, 0
			if i-1 >= 0 {
				t1 = aj - nums[i-1]
				t3 = ai - nums[i-1]
			}
			if j+1 < len(nums) {
				t2 = nums[j+1] - ai
				t4 = nums[j+1] - aj
			}
			diff := (abs(t1) + abs(t2)) - (abs(t3) + abs(t4))
			if diff > maxdiff {
				maxdiff = diff
			}
		}
	}
	return sum + maxdiff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(maxValueAfterReverse([]int{2, 3, 1, 5, 4}))
	fmt.Println(maxValueAfterReverse([]int{2, 4, 9, 24, 2, 1, 10}))
}

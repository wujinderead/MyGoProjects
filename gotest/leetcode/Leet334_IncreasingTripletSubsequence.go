package leetcode

import "fmt"

// Given an unsorted array return whether an increasing
// subsequence of length 3 exists or not in the array.
// Formally the function should:
// Return true if there exists i, j, k
// such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.
// Note: Your algorithm should run in O(n) time complexity and O(1) space complexity.
// Example 1:
//   Input: [1,2,3,4,5]
//   Output: true
// Example 2:
//   Input: [5,4,3,2,1]
//   Output: false

func increasingTriplet(nums []int) bool {
	// find longest increasing sequence that >= 3
	if len(nums) <= 2 {
		return false
	}
	var a, b int
	a = nums[0]
	lis := 1
	for i := 1; i < len(nums); i++ {
		if lis == 1 {
			if nums[i] > a {
				b = nums[i]
				lis++
			} else {
				a = nums[i]
			}
		}
		if lis == 2 {
			if nums[i] > b {
				return true
			} else if nums[i] > a {
				b = nums[i]
			} else {
				a = nums[i]
			}
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{1, 2, 3, 4, 5}))
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
	fmt.Println(increasingTriplet([]int{1, 2}))
	fmt.Println(increasingTriplet([]int{1, 2, 3}))
	fmt.Println(increasingTriplet([]int{2, 1, 3}))
}

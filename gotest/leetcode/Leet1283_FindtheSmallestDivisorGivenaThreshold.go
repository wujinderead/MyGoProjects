package main

import "fmt"

// Given an array of integers nums and an integer threshold, we will choose a positive integer
// divisor and divide all the array by it and sum the result of the division. Find the smallest
// divisor such that the result mentioned above is less than or equal to threshold.
// Each result of division is rounded to the nearest integer greater than or equal to
// that element. (For example: 7/3 = 3 and 10/2 = 5).
// It is guaranteed that there will be an answer.
// Example 1:
//   Input: nums = [1,2,5,9], threshold = 6
//   Output: 5
//   Explanation:
//     We can get a sum to 17 (1+2+5+9) if the divisor is 1.
//     If the divisor is 4 we can get a sum to 7 (1+1+2+3) and
//     if the divisor is 5 the sum will be 5 (1+1+1+2).
// Example 2:
//   Input: nums = [2,3,5,7,11], threshold = 11
//   Output: 3
// Example 3:
//   Input: nums = [19], threshold = 5
//   Output: 4
// Constraints:
//   1 <= nums.length <= 5 * 10^4
//   1 <= nums[i] <= 10^6
//   nums.length <= threshold <= 10^6

func smallestDivisor(nums []int, threshold int) int {
	// the possible answer is from 1 to the max value of nums
	// so use binary search, time complexity is O(len(nums)*log(max(nums)))
	high := 0
	for i := range nums {
		if nums[i] > high {
			high = nums[i]
		}
	}
	low := 1
tag:
	for low < high { // until low==high
		mid := (low + high) / 2
		sum := 0
		for i := range nums {
			sum += (nums[i] + mid - 1) / mid
			if sum > threshold { // mid is too low
				low = mid + 1
				continue tag
			}
		}
		high = mid
	}
	return high
}

func main() {
	fmt.Println(smallestDivisor([]int{1, 2, 5, 9}, 6))
	fmt.Println(smallestDivisor([]int{2, 3, 5, 7, 11}, 11))
	fmt.Println(smallestDivisor([]int{19}, 5))
}

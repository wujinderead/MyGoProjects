package main

import "fmt"

// https://leetcode.com/problems/sequential-digits/

// An integer has sequential digits if and only if each digit
// in the number is one more than the previous digit.
// Return a sorted list of all the integers in the range
// [low, high] inclusive that have sequential digits.
// Example 1:
//   Input: low = 100, high = 300
//   Output: [123,234]
// Example 2:
//   Input: low = 1000, high = 13000
//   Output: [1234,2345,3456,4567,5678,6789,12345]
// Constraints:
//   10 <= low <= high <= 10^9

func sequentialDigits(low int, high int) []int {
	nums := []int{
		12, 23, 34, 45, 56, 67, 78, 89,
		123, 234, 345, 456, 567, 678, 789,
		1234, 2345, 3456, 4567, 5678, 6789,
		12345, 23456, 34567, 45678, 56789,
		123456, 234567, 345678, 456789,
		1234567, 2345678, 3456789,
		12345678, 23456789,
		123456789,
	}
	start, end := 0, 0
	if low > nums[len(nums)-1] || high < nums[0] {
		return []int{}
	}
	for i := range nums {
		if low <= nums[i] {
			start = i
			break
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if high >= nums[i] {
			end = i
			break
		}
	}
	return nums[start : end+1]
}

func main() {
	fmt.Println(sequentialDigits(10, 11))
	fmt.Println(sequentialDigits(10, 12))
	fmt.Println(sequentialDigits(11, 12))
	fmt.Println(sequentialDigits(12, 12))
	fmt.Println(sequentialDigits(12, 13))
	fmt.Println(sequentialDigits(12, 23))
	fmt.Println(sequentialDigits(12, 22))
	fmt.Println(sequentialDigits(11, 122))
	fmt.Println(sequentialDigits(12, 123))
	fmt.Println(sequentialDigits(13, 123))
	fmt.Println(sequentialDigits(123456789, 123456790))
	fmt.Println(sequentialDigits(123456788, 123456789))
	fmt.Println(sequentialDigits(123456789, 123456789))
	fmt.Println(sequentialDigits(123456790, 123456790))
	fmt.Println(sequentialDigits(10, 123456790))
}

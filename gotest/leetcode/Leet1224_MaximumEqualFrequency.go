package leetcode

import "fmt"

// https://leetcode.com/problems/maximum-equal-frequency/

// Given an array nums of positive integers, return the longest possible
// length of an array prefix of nums, such that it is possible to remove exactly
// one element from this prefix so that every number that has appeared in it
// will have the same number of occurrences.
// If after removing one element there are no remaining elements, it's still
// considered that every appeared number has the same number of occurrences (0).
// Example 1:
//   Input: nums = [2,2,1,1,5,3,3,5]
//   Output: 7
//   Explanation: For the subarray [2,2,1,1,5,3,3] of length 7,
//     if we remove nums[4]=5, we will get [2,2,1,1,3,3],
//     so that each number will appear exactly twice.
// Example 2:
//   Input: nums = [1,1,1,2,2,2,3,3,3,4,4,4,5]
//   Output: 13
// Example 3:
//   Input: nums = [1,1,1,2,2,2]
//   Output: 5
// Example 4:
//   Input: nums = [10,2,8,9,3,8,1,5,2,3,7,6]
//   Output: 8
// Constraints:
//   2 <= nums.length <= 10^5
//   1 <= nums[i] <= 10^5

func maxEqualFreq(nums []int) int {
	if len(nums) <= 3 {
		return len(nums)
	}
	occur := make(map[int]int)
	oofo := make(map[int]map[int]struct{})
	for i := range nums {
		occur[nums[i]] += 1
	}
	for k, v := range occur {
		if _, ok := oofo[v]; ok {
			oofo[v][k] = struct{}{}
		} else {
			oofo[v] = map[int]struct{}{k: {}}
		}
	}
	for i := len(nums) - 1; i > 2; i-- {
		if len(oofo) == 1 { // one type of occurrence
			var o int
			var set map[int]struct{}
			for k, v := range oofo {
				o = k
				set = v
			}
			if o == 1 || len(set) == 1 {
				return i + 1
			}
		}
		if len(oofo) == 2 { // two type of occurrence
			j := 0
			var o1, o2 int
			var set1, set2 map[int]struct{}
			for k, v := range oofo {
				if j == 0 {
					o1 = k
					set1 = v
					j++
				} else {
					o2 = k
					set2 = v
				}
			}
			if (o1 == 1 && len(set1) == 1) || (o2 == 1 && len(set2) == 1) {
				return i + 1
			}
			if (o1 == o2+1 && len(set1) == 1) || (o2 == o1+1 && len(set2) == 1) {
				return i + 1
			}
		}
		// can not meet the condition, delete last integer
		n := nums[i]
		delete(oofo[occur[n]], n)
		if len(oofo[occur[n]]) == 0 {
			delete(oofo, occur[n])
		}
		occur[n] -= 1
		if occur[n] > 0 {
			if set, ok := oofo[occur[n]]; ok {
				set[n] = struct{}{}
			} else {
				oofo[occur[n]] = make(map[int]struct{})
				oofo[occur[n]][n] = struct{}{}
			}
		}
	}
	return 3
}

func main() {
	fmt.Println(maxEqualFreq([]int{2, 2, 1, 1, 5, 3, 3, 5}))
	fmt.Println(maxEqualFreq([]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}))
	fmt.Println(maxEqualFreq([]int{1, 1, 1, 2, 2, 2}))
	fmt.Println(maxEqualFreq([]int{10, 2, 8, 9, 3, 8, 1, 5, 2, 3, 7, 6}))
	fmt.Println(maxEqualFreq([]int{1, 2, 3, 1, 2, 3, 4, 4, 4, 4, 1, 2, 3, 5, 6}))
}

package main

import "fmt"

// Given a binary array nums, you should delete one element from it. Return the size of 
// the longest non-empty subarray containing only 1's in the resulting array. 
// Return 0 if there is no such subarray. 
// Example 1: 
//   Input: nums = [1,1,0,1]
//   Output: 3
//   Explanation: After deleting the number in position 2, [1,1,1] contains 3 numbers with value of 1's. 
// Example 2: 
//   Input: nums = [0,1,1,1,0,1,1,0,1]
//   Output: 5
//   Explanation: After deleting the number in position 4, [0,1,1,1,1,1,0,1] longest 
//     subarray with value of 1's is [1,1,1,1,1]. 
// Example 3: 
//   Input: nums = [1,1,1]
//   Output: 2
//Explanation: You must delete one element. 
// Example 4: 
//   Input: nums = [1,1,0,0,1,1,1,0,1]
//   Output: 4
// Example 5: 
//   Input: nums = [0,0,0]
//   Output: 0
// Constraints: 
//   1 <= nums.length <= 10^5 
//   nums[i] is either 0 or 1. 

func longestSubarray(nums []int) int {
	prev, cur, canSplice := 0, 0, false
	max := 0
	haszero := false
    for i, v := range nums {
		if v==1 {
			if cur==0 && i>=2 && nums[i-2]==1 {  // 1s start
				canSplice = true
			}
			cur++
			if cur > max {
				max = cur
			}
		}
		if v==0 {
			haszero = true
			if cur > 0 {      // 1s terminate
				if canSplice && cur+prev > max {
					max = cur+prev
				}
				prev = cur
				cur = 0
				canSplice = false
			}
		}
	}

	// last all 1s
	if cur > 0 && canSplice && cur+prev > max {
		max = cur+prev
	}

	// return result
	if !haszero {         // handle for all 1s
		return max-1
	}
	return max
}

func main() {
	fmt.Println(longestSubarray([]int{1,1,0,1}), 3)
	fmt.Println(longestSubarray([]int{0,1,1,1,0,1,1,0,1}), 5)
	fmt.Println(longestSubarray([]int{1,1,1}), 2)
	fmt.Println(longestSubarray([]int{1,1,0,0,1,1,1,0,1}), 4)
	fmt.Println(longestSubarray([]int{0,0,0}), 0)
}
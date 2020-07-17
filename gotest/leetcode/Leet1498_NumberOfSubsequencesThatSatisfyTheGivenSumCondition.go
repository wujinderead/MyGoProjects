package main

import (
	"fmt"
	"sort"
)

// Given an array of integers nums and an integer target. Return the number of non-empty subsequences 
// of nums such that the sum of the minimum and maximum element on it is less or equal than target. 
// Since the answer may be too large, return it modulo 10^9 + 7. 
// Example 1: 
//   Input: nums = [3,5,6,7], target = 9
//   Output: 4
//   Explanation: There are 4 subsequences that satisfy the condition.
//     [3] -> Min value + max value <= target (3 + 3 <= 9)
//     [3,5] -> (3 + 5 <= 9)
//     [3,5,6] -> (3 + 6 <= 9)
//     [3,6] -> (3 + 6 <= 9)
// Example 2: 
//   Input: nums = [3,3,6,8], target = 10
//   Output: 6
//   Explanation: There are 6 subsequences that satisfy the condition. (nums can have repeated numbers).
//     [3] , [3] , [3,3], [3,6] , [3,6] , [3,3,6] 
// Example 3: 
//   Input: nums = [2,3,3,4,6,7], target = 12
//   Output: 61
//   Explanation: There are 63 non-empty subsequences, two of them don't satisfy the condition ([6,7], [7]).
//     Number of valid subsequences (63 - 2 = 61).
// Example 4: 
//   Input: nums = [5,2,4,1,7,6,8], target = 16
//   Output: 127
//   Explanation: All non-empty subset satisfy the condition (2^7 - 1) = 127 
// Constraints: 
//   1 <= nums.length <= 10^5 
//   1 <= nums[i] <= 10^6 
//   1 <= target <= 10^6 

func numSubseq(nums []int, target int) int {
	// sort the array first
	sort.Sort(sort.IntSlice(nums))

	// pow[i] = 2^i mod p
	pow := make([]int, len(nums))
	p := int(1e9+7)
	pow[0] = 1
	for i:=1; i<len(nums); i++ {
		pow[i] = (pow[i-1]*2) % p
	}

	// two pointer for this problem
	count := 0
	s, e := 0, len(nums)-1
	for s <= e {
		if nums[s]+nums[e] > target {  // e too large, reduce it
			e--
		} else {    // nums[s]+nums[e]<=target, then we choose nums[s] ∪ {all subsets for nums[s+1...e]}
			count += pow[e-s]    // we have 2^(e-s) these sets
			count = count % p
			s++     // move s to right
		}
	}
    return count
}

func main() {
	for _, v := range []struct{ints []int; target, ans int} {
		{[]int{3,5,6,7}, 9, 4},
		{[]int{3,3,6,8}, 10, 6},
		{[]int{2,3,3,4,6,7}, 12, 61},
		{[]int{2,3,3,4,6,7,9}, 10, 58},
		{[]int{2,3,4,5,6,7,9}, 10, 53},		
		{[]int{5,2,4,1,7,6,8}, 16, 127},
		{[]int{2}, 3, 0},
		{[]int{2}, 4, 1},
		{[]int{14,4,6,6,20,8,5,6,8,12,6,10,14,9,17,16,9,7,14,11,14,15,13,11,
			10,18,13,17,17,14,17,7,9,5,10,13,8,5,18,20,7,5,5,15,19,14}, 22, 272187084},	
	} {
		fmt.Println(numSubseq(v.ints, v.target), v.ans)
	}
} 
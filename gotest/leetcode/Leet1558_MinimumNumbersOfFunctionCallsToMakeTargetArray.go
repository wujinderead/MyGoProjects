package main

import "fmt"

// https://leetcode.com/problems/minimum-numbers-of-function-calls-to-make-target-array/

// func(arr, op, idx) {
//    if op==0 {
//        arr[idx]++            // increment one
//    }
//    if op==1 {
//        for i:=range arr {    // double all
//            arr[i] *= 2
//        }
//     }
// }
// Your task is to form an integer array nums from an initial array of zeros arr that is
// the same size as nums. Return the minimum number of function calls to make nums from arr.
// The answer is guaranteed to fit in a 32-bit signed integer.
// Example 1:
//   Input: nums = [1,5]
//   Output: 5
//   Explanation: Increment by 1 (second element): [0, 0] to get [0, 1] (1 operation).
//     Double all the elements: [0, 1] -> [0, 2] -> [0, 4] (2 operations).
//     Increment by 1 (both elements)  [0, 4] -> [1, 4] -> [1, 5] (2 operations).
//     Total of operations: 1 + 2 + 2 = 5.
// Example 2:
//   Input: nums = [2,2]
//   Output: 3
//   Explanation: Increment by 1 (both elements) [0, 0] -> [0, 1] -> [1, 1] (2 operations).
//     Double all the elements: [1, 1] -> [2, 2] (1 operation).
//     Total of operations: 2 + 1 = 3.
// Example 3:
//   Input: nums = [4,2,5]
//   Output: 6
//   Explanation: (initial)[0,0,0] -> [1,0,0] -> [1,0,1] -> [2,0,2] -> [2,1,2] -> [4,2,4] -> [4,2,5](nums).
// Example 4:
//   Input: nums = [3,2,2,4]
//   Output: 7
// Example 5:
//   Input: nums = [2,4,8,16]
//   Output: 8
// Constraints:
//   1 <= nums.length <= 10^5
//   0 <= nums[i] <= 10^9

func minOperations(nums []int) int {
	// find the path for individual element, the increment can't overlap, but double can.
	// e.g., [4,2,5], 4 need 1 increment and 2 doubles (0->1->2->4), 2 need 1 inc and 1 dbl,
	// 5 need 2 inc and 2 dbl. so we need total 1+1+2=4 inc and max(2,1,2)=2 dbl. the total ops is 6.
	// the inc for an element is the bit count. the double depends on the maximal number.
	ans := 0
	max := 0
	for _, v := range nums {
		ans += countBit(v) // the increment needed
		if v>max {
			max = v
		}
	}
	for max/2>0 {       // the double needed depends on the most significant bit of max number
		ans++
		max = max>>1
	}
    return ans
}

func countBit(i int) int {
	i = i - ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = (i + (i >> 4)) & 0x0f0f0f0f
	i = i + (i >> 8)
	i = i + (i >> 16)
	return i & 0x3f
}

func main() {
	for _, v := range []struct{nums []int; ans int} {
		{[]int{1,5}, 5},
		{[]int{2,2}, 3},
		{[]int{4,2,5}, 6},
		{[]int{3,2,2,4}, 7},
		{[]int{2,4,8,16}, 8},
		{[]int{0,0}, 0},
		{[]int{1,0}, 1},
	} {
		fmt.Println(minOperations(v.nums), v.ans)
	}
}
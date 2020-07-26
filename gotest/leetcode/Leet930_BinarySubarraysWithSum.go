package main

import "fmt"

// https://leetcode.com/problems/binary-subarrays-with-sum/

// In an array A of 0s and 1s, how many non-empty subarrays have sum S?
// Example 1:
//   Input: A = [1,0,1,0,1], S = 2
//   Output: 4
//   Explanation: 
//     The 4 subarrays are bolded below:
//     [1,0,1,0,1]
//     [1,0,1,0,1]
//     [1,0,1,0,1]
//     [1,0,1,0,1]
// Note:
//   A.length <= 30000
//   0 <= S <= A.length
//   A[i] is either 0 or 1.

func numSubarraysWithSum(A []int, S int) int {
	mapp := make(map[int]int)
	count := 0
	sum := 0
	mapp[0] = 1
	for _, v := range A {
		sum += v
		count += mapp[sum-S]
		mapp[sum] = mapp[sum]+1
	}
	return count
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1}, 2), 4)
	fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1}, 1), 8)
	fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1}, 3), 1)
	fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1}, 4), 0)
	fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1}, 0), 2)
	fmt.Println(numSubarraysWithSum([]int{1,1,1,1,1}, 0), 0)
}
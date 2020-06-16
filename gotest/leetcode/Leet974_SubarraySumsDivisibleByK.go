package main

import (
    "fmt"
)

// https://leetcode.com/problems/subarray-sums-divisible-by-k/

// Given an array A of integers, return the number of (contiguous, non-empty) subarrays that have a sum divisible by K.
// Example 1:
//   Input: A = [4,5,0,-2,-3,1], K = 5
//   Output: 7
//   Explanation: There are 7 subarrays with a sum divisible by K = 5:
//     [4, 5, 0, -2, -3, 1], [5], [5, 0], [5, 0, -2, -3], [0], [0, -2, -3], [-2, -3]
// Note:
//   1 <= A.length <= 30000
//   -10000 <= A[i] <= 10000
//   2 <= K <= 10000 

func subarraysDivByK(A []int, K int) int {
    cursum := 0
    mapp := make(map[int]int)
    mapp[0] = 1   // no numbers sums to 0, 1 way
    count := 0
    for i:=0; i<len(A); i++ {
    	cursum += A[i]
    	key := (cursum%K+K)%K  // avoid negative 
    	prev := mapp[key]  // how many previous sums that mod K = cursum mod K
    	count += prev
    	mapp[key] = mapp[key]+1
    }
    return count
}

func main() {
	fmt.Println(subarraysDivByK([]int{4,5,0,-2,-3,1}, 5))
	fmt.Println(subarraysDivByK([]int{-1,2,9}, 2))
}
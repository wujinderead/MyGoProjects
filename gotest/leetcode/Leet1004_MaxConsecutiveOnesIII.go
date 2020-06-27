package main

import "fmt"

// Given an array A of 0s and 1s, we may change up to K values from 0 to 1. 
// Return the length of the longest (contiguous) subarray that contains only 1s.
// Example 1: 
//   Input: A = [1,1,1,0,0,0,1,1,1,1,0], K = 2
//   Output: 6
//   Explanation: 
//     [1,1,1,0,0,1,1,1,1,1,1]
//     Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.  
// Example 2: 
//   Input: A = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
//   Output: 10
//   Explanation: 
//     [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
//     Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
// Note: 
//   1 <= A.length <= 20000 
//   0 <= K <= A.length 
//   A[i] is 0 or 1 

// the problem equals to: find the longest subarray with at most K zeros.
func longestOnes(A []int, K int) int {
	s := 0
	flip := 0
	maxlen := 0
	for i:=0; i<len(A); i++ {
		if A[i]==0 {
			if flip<K {   // can flip
				flip++
			} else {      // can't flip, move s right
				for s<len(A) && A[s]!=0 {   // skip A[s]==1
					s++
				}
				if s<len(A) {  // A[s]==0, skip A[s] so we can flip A[i]
					s++
				}
			}
		}
		maxlen = max(maxlen, i-s+1)
	}
	return maxlen
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0}, 2))
	fmt.Println(longestOnes([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 3))
	fmt.Println(longestOnes([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}, 0))
	fmt.Println(longestOnes([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1,0}, 0))
	fmt.Println(longestOnes([]int{0,0,0,0,0,0}, 0))
	fmt.Println(longestOnes([]int{0,0,0,1,0,0}, 0))
	fmt.Println(longestOnes([]int{0,0,0,0,0,0}, 1))
	fmt.Println(longestOnes([]int{0,0,0,0,0,0}, 2))
}
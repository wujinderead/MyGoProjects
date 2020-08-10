package main

import "fmt"

// https://leetcode.com/problems/bitwise-ors-of-subarrays/

// We have an array A of non-negative integers.
// For every (contiguous) subarray B = [A[i], A[i+1], ..., A[j]] (with i <= j), we take the bitwise OR ]
// of all the elements in B, obtaining a result A[i] | A[i+1] | ... | A[j]. Return the number of 
// possible results.  (Results that occur more than once are only counted once in the final answer.)
// Example 1:
//    Input: [0]
//    Output: 1
//    Explanation: 
//      There is only one possible result: 0.
// Example 2:
//   Input: [1,1,2]
//   Output: 3
//   Explanation: 
//     The possible subarrays are [1], [1], [2], [1, 1], [1, 2], [1, 1, 2].
//     These yield the results 1, 1, 2, 1, 3, 3.
//     There are 3 unique values, so the answer is 3.
// Example 3:
//   Input: [1,2,4]
//   Output: 6
//   Explanation: 
//     The possible results are 1, 2, 3, 4, 6, and 7.
// Note:
//   1 <= A.length <= 50000
//   0 <= A[i] <= 10^9

// let set[i] be the unique values of OR(arr[0...i]), OR(arr[1...i]), ..., OR(arr[i...i]),
// the size of set can be at most log(maxint). we can also know that 
// set[i+1] = {x OR arr[i+1], x ∈ set[i]} ∪ {arr[i+1]}.
func subarrayBitwiseORs(A []int) int {
	set := make(map[int]struct{}, 32)
	allset := make(map[int]struct{}, 32)
	set[A[0]] = struct{}{}
	allset[A[0]] = struct{}{}
	for i:=1; i<len(A); i++ {
		newset := make(map[int]struct{}, 32)
		for k := range set {
			newset[A[i] | k] = struct{}{}
			allset[A[i] | k] = struct{}{}
		}
		newset[A[i]] = struct{}{}
		allset[A[i]] = struct{}{}	
		set = newset	
	}
	return len(allset)
}

func main() {
	for _, v := range []struct{arr []int; ans int} {
		{[]int{0}, 1},
		{[]int{1,1,2}, 3},
		{[]int{1,2,4}, 6},
	} {
		fmt.Println(subarrayBitwiseORs(v.arr), v.ans)
	}
}

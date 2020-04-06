package main

import "fmt"

// A sequence of number is called arithmetic if it consists of at least three
// elements and if the difference between any two consecutive elements is the same.
// For example, these are arithmetic sequence: 
//   1, 3, 5, 7, 9
//   7, 7, 7, 7
//   3, -1, -5, -9
// The following sequence is not arithmetic. 1, 1, 2, 5, 7
// A zero-indexed array A consisting of N numbers is given. A slice of that array is
// any pair of integers (P, Q) such that 0 <= P < Q < N.
// A slice (P, Q) of array A is called arithmetic if the sequence: 
// A[P], A[p + 1], ..., A[Q - 1], A[Q] is arithmetic. In particular, this means 
// that P + 1 < Q.
// The function should return the number of arithmetic slices in the array A.
// Example:
//   A = [1, 2, 3, 4]
//   return: 3, for 3 arithmetic slices in A: [1, 2, 3], [2, 3, 4] and [1, 2, 3, 4]
//     itself.

func numberOfArithmeticSlices(A []int) int {
	count := 0
	i := 1
	c := 1
	for i<len(A) {
		if i<len(A) && i+1<len(A) && A[i+1]-A[i]==A[i]-A[i-1] {
			count += c
			c++
		} else {
			c = 1
		}
		i++
	}
	return count
}

// IMPROVEMENT: more concise solution:
func numberOfArithmeticSlices1(A []int) int {
	count := 0
	c := 1
	for i:=2; i<len(A); i++ {
		if A[i]-A[i-1]==A[i-1]-A[i-2] {
			count += c
			c++
		} else {
			c = 1
		}
	}
	return count
}

func main() {
    fmt.Println(numberOfArithmeticSlices([]int{1,2,3,4}))
    fmt.Println(numberOfArithmeticSlices([]int{1,1,2,3,4,6,8,10}))
    fmt.Println(numberOfArithmeticSlices([]int{1,1,2,3,4,6,8,10,11}))
    fmt.Println(numberOfArithmeticSlices([]int{1,1,2}))
    fmt.Println(numberOfArithmeticSlices([]int{1,1}))
    fmt.Println(numberOfArithmeticSlices([]int{1}))
    fmt.Println(numberOfArithmeticSlices([]int{1,2,3}))
}
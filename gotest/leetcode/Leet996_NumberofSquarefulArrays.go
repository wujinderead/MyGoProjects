package main

import "fmt"

// https://leetcode.com/problems/number-of-squareful-arrays/

// Given an array A of non-negative integers, the array is squareful if for every
// pair of adjacent elements, their sum is a perfect square.
// Return the number of permutations of A that are squareful. Two permutations A1 and
// A2 differ if and only if there is some index i such that A1[i] != A2[i].
// Example 1:
//   Input: [1,17,8]
//   Output: 2
//   Explanation: [1,8,17] and [17,8,1] are the valid permutations.
// Example 2:
//   Input: [2,2,2]
//   Output: 1
// Note:
//   1 <= A.length <= 12
//   0 <= A[i] <= 1e9

func numSquarefulPerms(A []int) int {
	if len(A) == 1 {
		if isSquare(A[0]) {
			return 1
		}
		return 0
	}
	squarefuls := 0
	next(A, 0, &squarefuls)
	return squarefuls
}

func next(A []int, curind int, squareful *int) {
	if curind == len(A) {
		fmt.Println(A)
		*squareful++
		return
	}
	for i := curind; i < len(A); i++ {
		dup := false
		for j := curind; j < i; j++ {
			if A[j] == A[i] {
				dup = true
				break
			}
		}
		if !dup {
			A[i], A[curind] = A[curind], A[i]
			// continue only if it's current and previous number sum is squareful
			if curind == 0 || isSquare(A[curind]+A[curind-1]) {
				next(A, curind+1, squareful)
			}
			A[i], A[curind] = A[curind], A[i]
		}
	}
}

func isSquare(i int) bool {
	if i < 10 && i == 0 || i == 1 || i == 4 || i == 9 {
		return true
	}
	// binary search to check if it's a square
	l, h := 3, i/3
	for l <= h {
		t := l + (h-l)/2
		if t*t == i {
			return true
		}
		if t*t < i {
			l = t + 1
		} else {
			h = t - 1
		}
	}
	return false
}

func main() {
	fmt.Println(numSquarefulPerms([]int{1, 17, 8}))
	fmt.Println(numSquarefulPerms([]int{2, 2, 2}))
	fmt.Println(numSquarefulPerms([]int{8, 3, 1, 17, 8}))
}

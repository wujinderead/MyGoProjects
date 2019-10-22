package leetcode

import "fmt"

// https://leetcode.com/problems/ugly-number-ii/

// Write a program to find the n-th ugly number.
// Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.
// Example:
//   Input: n = 10
//   Output: 12
//   Explanation: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 is the sequence of the first 10 ugly numbers.

func nthUglyNumber(n int) int {
	if n < 1 {
		return 0
	}
	ugly := make([]int, n)
	ugly[0] = 1
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		ugly[i] = min(2*ugly[i2], min(3*ugly[i3], 5*ugly[i5]))
		for ugly[i2]*2 <= ugly[i] {
			i2++
		}
		for ugly[i3]*3 <= ugly[i] {
			i3++
		}
		for ugly[i5]*5 <= ugly[i] {
			i5++
		}
	}
	return ugly[n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(i, "=", nthUglyNumber(i))
	}
}

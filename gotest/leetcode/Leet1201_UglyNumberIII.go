package main

import (
	"fmt"
)

// https://leetcode.com/problems/ugly-number-iii/

// Write a program to find the n-th ugly number.
// Ugly numbers are positive integers which are divisible by a or b or c.
// Example 1:
//   Input: n = 3, a = 2, b = 3, c = 5
//   Output: 4
//   Explanation: The ugly numbers are 2, 3, 4, 5, 6, 8, 9, 10... The 3rd is 4.
// Example 2:
//   Input: n = 4, a = 2, b = 3, c = 4
//   Output: 6
//   Explanation: The ugly numbers are 2, 3, 4, 6, 8, 9, 10, 12... The 4th is 6.
// Example 3:
//   Input: n = 5, a = 2, b = 11, c = 13
//   Output: 10
//   Explanation: The ugly numbers are 2, 4, 6, 8, 10, 11, 12, 13... The 5th is 10.
// Example 4:
//   Input: n = 1000000000, a = 2, b = 217983653, c = 336916467
//   Output: 1999999984
// Constraints:
//   1 <= n, a, b, c <= 10^9
//   1 <= a * b * c <= 10^18         (max int64 is 9.2e+18)
//   It's guaranteed that the result will be in range [1, 2 * 10^9]

func nthUglyNumber(n int, a int, b int, c int) int {
	// use venn diagram to get the number of ugly numbers less than a number
	lcmab := lcm(a, b)
	lcmbc := lcm(b, c)
	lcmac := lcm(a, c)
	lcmabc := lcm(a, lcm(b, c))
	low := 1
	high := int(2e9)
	for low < high {
		mid := (low + high) / 2
		count := mid/a + mid/b + mid/c - mid/lcmab - mid/lcmbc - mid/lcmac + mid/lcmabc
		//fmt.Println(min, mid, max, count)
		if count < n {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func gcd(a, b int) int {
	for a == 0 {
		return b
	}
	return gcd(b%a, a)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	fmt.Println(nthUglyNumber(3, 2, 3, 5))
	fmt.Println(nthUglyNumber(4, 2, 3, 4))
	fmt.Println(nthUglyNumber(5, 2, 11, 13))
	fmt.Println(nthUglyNumber(1000000000, 2, 217983653, 336916467))
}

package main

import "fmt"

// https://leetcode.com/problems/super-ugly-number/

// Write a program to find the nth super ugly number. Super ugly numbers are positive numbers
// whose all prime factors are in the given prime list primes of size k.
// Example:
//   Input: n = 12, primes = [2,7,13,19]
//   Output: 32
//   Explanation: [1,2,4,7,8,13,14,16,19,26,28,32] is the sequence of the first 12
//             super ugly numbers given primes = [2,7,13,19] of size 4.

func nthSuperUglyNumber(n int, primes []int) int {
	if n < 1 {
		return 0
	}
	ugly := make([]int, n)
	ugly[0] = 1
	ind := make([]int, len(primes))
	for i := 1; i < n; i++ {
		ugly[i] = 0x7fffffff
		for j := 0; j < len(primes); j++ {
			if primes[j]*ugly[ind[j]] < ugly[i] {
				ugly[i] = primes[j] * ugly[ind[j]]
			}
		}
		for j := 0; j < len(primes); j++ {
			for primes[j]*ugly[ind[j]] <= ugly[i] {
				ind[j]++
			}
		}
	}
	//fmt.Printf("n=%d, primes=%v, ugly=%v\n", n, primes, ugly)
	return ugly[n-1]
}

func main() {
	fmt.Println(nthSuperUglyNumber(13, []int{2, 3, 5}))
	fmt.Println(nthSuperUglyNumber(16, []int{5, 17, 23, 31}))
	fmt.Println(nthSuperUglyNumber(16, []int{2, 7, 17, 29, 53}))
}

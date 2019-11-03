package main

import "fmt"

// https://leetcode.com/problems/2-keys-keyboard/

// Initially on a notepad only one character 'A' is present.
// You can perform two operations on this notepad for each step:
// Copy All: You can copy all the characters present on the notepad (partial copy is not allowed).
// Paste: You can paste the characters which are copied last time.
// Given a number n. You have to get exactly n 'A' on the notepad by performing the minimum
// number of steps permitted. Output the minimum number of steps to get n 'A'.
// Example 1:
//   Input: 3
//   Output: 3
//   Explanation:
//     Intitally, we have one character 'A'.
//     In step 1, we use Copy All operation.
//     In step 2, we use Paste operation to get 'AA'.
//     In step 3, we use Paste operation to get 'AAA'.
// Note:
//   The n will be in the range [1, 1000].

func minSteps(n int) int {
	// it can be transformed to factorization problem
	if n == 1 {
		return 1
	}
	primes := []int{2}
	for i := 3; i <= n; i++ {
		isprime := true
		for _, j := range primes {
			if i%j == 0 {
				isprime = false
				break
			}
		}
		if isprime {
			primes = append(primes, i)
		}
	}
	fmt.Println(primes)
	factors := []int{}
	for _, j := range primes {
		for n%j == 0 {
			factors = append(factors, j)
			n = n / j
		}
	}
	fmt.Println(factors)
	sum := 0
	for _, j := range factors {
		sum += j
	}
	return sum
}

func main() {
	fmt.Println(minSteps(10))
	fmt.Println(minSteps(1))
	fmt.Println(minSteps(2))
	fmt.Println(minSteps(3))
	fmt.Println(minSteps(125))
	fmt.Println(minSteps(127))
}

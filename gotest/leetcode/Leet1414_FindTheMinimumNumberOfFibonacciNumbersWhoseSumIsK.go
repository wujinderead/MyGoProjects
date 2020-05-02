package main

import "fmt"

// https://leetcode.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/

// Given the number k, return the minimum number of Fibonacci numbers whose sum is
// equal to k, whether a Fibonacci number could be used multiple times.
// The Fibonacci numbers are defined as:
//   F1 = 1
//   F2 = 1
//   Fn = Fn-1 + Fn-2 , for n > 2.
// It is guaranteed that for the given constraints we can always find such
// fibonacci numbers that sum k.
// Example 1:
//   Input: k = 7
//   Output: 2
//   Explanation: The Fibonacci numbers are: 1, 1, 2, 3, 5, 8, 13, ...
//     For k = 7 we can use 2 + 5 = 7.
// Example 2:
//   Input: k = 10
//   Output: 2
//   Explanation: For k = 10 we can use 2 + 8 = 10.
// Example 3:
//   Input: k = 19
//   Output: 3
//   Explanation: For k = 19 we can use 1 + 5 + 13 = 19.
// Constraints:
//   1 <= k <= 10^9

func findMinFibonacciNumbers(k int) int {
	fib := []int{1, 1}
	a, b := 1, 1
	for a+b<=k {
		a, b = a+b, a
		fib = append(fib, a)
	}
	i := len(fib)
	count := 0
	for k>0 {
		for j:=i-1; j>=0; j-- {
			if fib[j]<=k {
				count++
				k = k-fib[j]
				i = j
				break
			}
		}
	}
	return count
}

func main() {
	fmt.Println(findMinFibonacciNumbers(7))
	fmt.Println(findMinFibonacciNumbers(10))
	fmt.Println(findMinFibonacciNumbers(10e9))
	fmt.Println(findMinFibonacciNumbers(17))
	fmt.Println(findMinFibonacciNumbers(100))
	for i:=0; i<=20; i++ {
		fmt.Println(i, findMinFibonacciNumbers(i))
	}
}
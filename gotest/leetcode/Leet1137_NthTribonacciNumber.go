package leetcode

import "fmt"

// https://leetcode.com/problems/n-th-tribonacci-number/

// The Tribonacci sequence Tn is defined as follows:
// T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
// Given n, return the value of Tn.
func tribonacci(n int) int {
	a, b, c := 0, 1, 1
	for i := 0; i < n; i++ {
		a, b, c = b, c, a+b+c
	}
	return a
}

func main() {
	fmt.Println(tribonacci(0))
	fmt.Println(tribonacci(1))
	fmt.Println(tribonacci(2))
	fmt.Println(tribonacci(3))
	fmt.Println(tribonacci(4))
	fmt.Println(tribonacci(5))
	fmt.Println(tribonacci(25))
}

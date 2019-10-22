package leetcode

import "fmt"

// https://leetcode.com/problems/prime-arrangements/

// Return the number of permutations of 1 to n so that prime numbers are at prime indices (1-indexed.)
// Since the answer may be large, return the answer modulo 10^9 + 7.
// Example 1:
//   Input: n = 5
//   Output: 12
// Example 2:
//   Input: n = 100
//   Output: 682289015
func numPrimeArrangements(n int) int {
	// let a be the number of prime <= n, the answer is (n-a)!a!
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41,
		43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	if n < 1 {
		return 0
	}
	a := 0
	for _, p := range primes {
		if p <= n {
			a++
		} else {
			break
		}
	}
	b := n - a
	ans := 1
	for i := 1; i <= a; i++ {
		ans = ans * i % 1000000007
	}
	for i := 1; i <= b; i++ {
		ans = ans * i % 1000000007
	}
	return ans
}

func main() {
	fmt.Println(numPrimeArrangements(0))
	fmt.Println(numPrimeArrangements(1))
	fmt.Println(numPrimeArrangements(2))
	fmt.Println(numPrimeArrangements(3))
	fmt.Println(numPrimeArrangements(10), 24*720)
	fmt.Println(numPrimeArrangements(11), 120*720)
	fmt.Println(numPrimeArrangements(12), 120*720*7)
	fmt.Println(numPrimeArrangements(100))
}

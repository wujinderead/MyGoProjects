package main

import "fmt"

func largestPalindrome(n int) int {
	prime := 1337
	if n == 1 {
		return 9
	}
	N := 1
	for i := 0; i < n; i++ {
		N *= 10
	}
	// odd e.g. n=3
	// product = (1000-a)(1000-b) = 1000(1000-a-b)+ab
	// thus '1000-a-b' is high 3 bits, 'ab' is low 3 bits, they should be palindromic
	aa, bb := make([]int, n), make([]int, n)
	for a := 1; a < N; {
		var b int
		var d int
		if a%10 == 1 {
			b = a - 2
			d = 2
		}
		if a%10 == 3 {
			b = a
			d = 4
		}
		if a%10 == 7 {
			b = a
			d = 2
		}
		if a%10 == 9 {
			b = a - 8
			d = 2
		}
		for ; b > 0; b -= 10 {
			low := (a * b) % N
			high := N - a - b + a*b/N
			if isPalindromic(low, high, n, aa, bb) {
				fmt.Println(a, b, N-a, N-b)
				ans := (N - a) % prime
				ans = (ans * (N - b)) % prime
				return ans
			}
		}
		a += d
	}
	return -1
}

func isPalindromic(a, b, n int, aa, bb []int) bool {
	for i := 0; i < n; i++ {
		aa[i] = a % 10
		bb[i] = b % 10
		a /= 10
		b /= 10
	}
	i := 0
	j := n - 1
	for i < n && j >= 0 {
		if aa[i] != bb[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(largestPalindrome(1))
	fmt.Println(largestPalindrome(2))
	fmt.Println(largestPalindrome(3))
	fmt.Println(largestPalindrome(4))
	fmt.Println(largestPalindrome(5))
	fmt.Println(largestPalindrome(6))
	fmt.Println(largestPalindrome(7))
	fmt.Println(largestPalindrome(8))
	fmt.Println(largestPalindrome(9))
	fmt.Println(largestPalindrome(10))
	fmt.Println(largestPalindrome(11))
}

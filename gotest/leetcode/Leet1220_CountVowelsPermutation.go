package leetcode

import "fmt"

// https://leetcode.com/problems/count-vowels-permutation/

// Given an integer n, your task is to count how many strings of length n can be formed under the following rules:
// Each character is a lower case vowel ('a', 'e', 'i', 'o', 'u')
// Each vowel 'a' may only be followed by an 'e'.
// Each vowel 'e' may only be followed by an 'a' or an 'i'.
// Each vowel 'i' may not be followed by another 'i'.
// Each vowel 'o' may only be followed by an 'i' or a 'u'.
// Each vowel 'u' may only be followed by an 'a'.
// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: n = 1
//   Output: 5
//   Explanation: All possible strings are: "a", "e", "i" , "o" and "u".
// Example 2:
//   Input: n = 2
//   Output: 10
//   Explanation: All possible strings are: "ae", "ea", "ei", "ia", "ie", "io", "iu", "oi", "ou" and "ua".
// Example 3:
//   Input: n = 5
//   Output: 68
// Constraints:
// 1 <= n <= 20000

func countVowelPermutation(n int) int {
	if n < 1 {
		return 0
	}
	mod := 1000000007
	a, e, i, o, u := make([]int, n), make([]int, n), make([]int, n), make([]int, n), make([]int, n)
	a[0], e[0], i[0], o[0], u[0] = 1, 1, 1, 1, 1
	for k := 1; k < n; k++ {
		a[k] = (i[k-1] + u[k-1] + e[k-1]) % mod
		e[k] = (a[k-1] + i[k-1]) % mod
		i[k] = (e[k-1] + o[k-1]) % mod
		o[k] = i[k-1]
		u[k] = (i[k-1] + o[k-1]) % mod
	}
	return (a[n-1] + e[n-1] + i[n-1] + o[n-1] + u[n-1]) % mod
}

func main() {
	fmt.Println(countVowelPermutation(0))
	fmt.Println(countVowelPermutation(1))
	fmt.Println(countVowelPermutation(2))
	fmt.Println(countVowelPermutation(3))
	fmt.Println(countVowelPermutation(4))
	fmt.Println(countVowelPermutation(5))
}

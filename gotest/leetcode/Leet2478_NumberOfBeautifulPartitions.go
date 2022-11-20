package main

import "fmt"

// https://leetcode.com/problems/number-of-beautiful-partitions/

// You are given a string s that consists of the digits '1' to '9' and two integers k and minLength.
// A partition of s is called beautiful if:
//   s is partitioned into k non-intersecting substrings.
//   Each substring has a length of at least minLength.
//   Each substring starts with a prime digit and ends with a non-prime digit.
//     Prime digits are '2', '3', '5', and '7', and the rest of the digits are non-prime.
// Return the number of beautiful partitions of s. Since the answer may be very large, return it
// modulo 10⁹ + 7.
// A substring is a contiguous sequence of characters within a string.
// Example 1:
//   Input: s = "23542185131", k = 3, minLength = 2
//   Output: 3
//   Explanation: There exists three ways to create a beautiful partition:
//     "2354 | 218 | 5131"
//     "2354 | 21851 | 31"
//     "2354218 | 51 | 31"
// Example 2:
//   Input: s = "23542185131", k = 3, minLength = 3
//   Output: 1
//   Explanation: There exists one way to create a beautiful partition: "2354 | 218 | 5131".
// Example 3:
//   Input: s = "3312958", k = 3, minLength = 1
//   Output: 1
//   Explanation: There exists one way to create a beautiful partition: "331 | 29 | 58".
// Constraints:
//   1 <= k, minLength <= s.length <= 1000
//   s consists of the digits '1' to '9'.

func beautifulPartitions(s string, k int, minLength int) int {
	if minLength == 1 {
		minLength = 2
	}
	if !isPrime(s[0]) || isPrime(s[len(s)-1]) || k*minLength > len(s) {
		return 0
	}
	cache := make(map[[2]int]int)
	return recur(cache, s, 0, k, minLength)
}

func isPrime(b byte) bool {
	return b == '2' || b == '3' || b == '5' || b == '7'
}

func recur(cache map[[2]int]int, s string, index, k, minlen int) int {
	if k == 1 {
		return 1
	}
	if v, ok := cache[[2]int{index, k}]; ok {
		return v
	}
	ans := 0
	for e := index + minlen - 1; e < len(s)-(k-1)*minlen; e++ {
		if !isPrime(s[e]) && isPrime(s[e+1]) {
			ans += recur(cache, s, e+1, k-1, minlen)
		}
	}
	cache[[2]int{index, k}] = ans % int(1e9+7)
	return ans % int(1e9+7)
}

func main() {
	for _, v := range []struct {
		s                 string
		k, minLength, ans int
	}{
		{"23542185131", 3, 2, 3},
		{"23542185131", 3, 3, 1},
		{"3312958", 3, 1, 1},
	} {
		fmt.Println(beautifulPartitions(v.s, v.k, v.minLength), v.ans)
	}
}

package main

import (
	"fmt"
	"math/big"
)

// https://leetcode.com/problems/count-the-number-of-ideal-arrays/

// You are given two integers n and maxValue, which are used to describe an ideal array.
// A 0-indexed integer array arr of length n is considered ideal if the following conditions hold:
//   Every arr[i] is a value from 1 to maxValue, for 0 <= i < n.
//   Every arr[i] is divisible by arr[i - 1], for 0 < i < n.
// Return the number of distinct ideal arrays of length n. Since the answer may be very large,
// return it modulo 10⁹ + 7.
// Example 1:
//   Input: n = 2, maxValue = 5
//   Output: 10
//   Explanation: The following are the possible ideal arrays:
//     - Arrays starting with the value 1 (5 arrays): [1,1], [1,2], [1,3], [1,4], [1,5]
//     - Arrays starting with the value 2 (2 arrays): [2,2], [2,4]
//     - Arrays starting with the value 3 (1 array): [3,3]
//     - Arrays starting with the value 4 (1 array): [4,4]
//     - Arrays starting with the value 5 (1 array): [5,5]
//     There are a total of 5 + 2 + 1 + 1 + 1 = 10 distinct ideal arrays.
// Example 2:
//   Input: n = 5, maxValue = 3
//   Output: 11
//   Explanation: The following are the possible ideal arrays:
//     Arrays starting with the value 1 (9 arrays):
//     - With no other distinct values (1 array): [1,1,1,1,1]
//     - With 2ⁿᵈ distinct value 2 (4 arrays): [1,1,1,1,2], [1,1,1,2,2], [1,1,2,2,2], [1,2,2,2,2]
//     - With 2ⁿᵈ distinct value 3 (4 arrays): [1,1,1,1,3], [1,1,1,3,3], [1,1,3,3,3], [1,3,3,3,3]
//     - Arrays starting with the value 2 (1 array): [2,2,2,2,2]
//     - Arrays starting with the value 3 (1 array): [3,3,3,3,3]
//     There are a total of 9 + 1 + 1 = 11 distinct ideal arrays.
// Constraints:
//   2 <= n <= 10⁴
//   1 <= maxValue <= 10⁴

func idealArrays(n int, maxValue int) int {
	if maxValue == 1 || n == 1 {
		return maxValue
	}
	const p = int(1e9) + 7

	// dividers[i]: the set of dividers of i
	dividers := make([][]int, maxValue+1)
	for i := range dividers {
		dividers[i] = make([]int, 1, 5)
		dividers[i][0] = 1
	}
	for i := 2; i <= maxValue/2; i++ {
		j := 2
		for i*j <= maxValue {
			dividers[i*j] = append(dividers[i*j], i)
			j++
		}
	}

	// old[i] the number of ideal array with distinct values that end at i
	old := make([]int, maxValue+1)
	for i := range old {
		old[i] = 1 // initially, old is for length 1
	}
	next := make([]int, maxValue+1)
	// count[i]: the number of ideal array with distinct values with length i
	// ideal array with distinct (also increasing) values has max length 14
	count := make([]int, min(n, 14)+1)
	count[1] = maxValue // single distinct value

	for l := 2; l <= min(n, 14); l++ {
		for i := range next {
			next[i] = 0
		}
		for i := 2; i <= maxValue; i++ {
			for _, v := range dividers[i] {
				next[i] += old[v]
			}
		}
		for i := range old { // sum
			count[l] += next[i]
		}
		old, next = next, old
	}

	// get C(n-1, 0) ... C(n-1, n-1)
	c := make([]int, n)
	c[0] = 1
	P := new(big.Int).SetInt64(int64(p))
	ii := new(big.Int)
	for i := 1; i < n; i++ {
		inv := ii.ModInverse(ii.SetInt64(int64(i)), P).Int64()
		c[i] = (((c[i-1] * (n - i)) % p) * int(inv)) % p
	}

	// find answer
	ans := count[1] // the number with only single value
	for i := 2; i <= min(n, 14); i++ {
		if count[i] == 0 {
			break
		}
		// for ideal array with distinct (also increasing) values with length i,
		// we want to extend it to length n, we have C(n-1, i-1) ways.
		// e.g., we have 3 distinct values, we want to extend to length 6,
		// we need choose 2 slots out of the 5 slots, i.e. C(5,2)
		// x x x x x x    // 6 numbers
		//  | | | | |     // 5 slots, choose 2 slots to separate distinct numbers
		// 1 2 2 2 4 4
		//  |     |       // the 2 chosen slots
		ans = (ans + count[i]*c[i-1]) % p
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		n, maxValue, ans int
	}{
		{2, 5, 10},
		{5, 3, 11},
		{4, 6, 39},
		{60, 100, 730303103},
		{184, 389, 510488787},
	} {
		fmt.Println(idealArrays(v.n, v.maxValue), v.ans)
	}
}

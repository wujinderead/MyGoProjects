package main

import "fmt"

// https://leetcode.com/problems/number-of-unique-good-subsequences/

// You are given a binary string binary. A subsequence of binary is considered good if it is
// not empty and has no leading zeros (with the exception of "0").
// Find the number of unique good subsequences of binary.
// For example, if binary = "001", then all the good subsequences are ["0", "0", "1"], so the
// unique good subsequences are "0" and "1". Note that subsequences "00", "01", and "001" are
// not good because they have leading zeros.
// Return the number of unique good subsequences of binary. Since the answer may be very large,
// return it modulo 10^9 + 7.
// A subsequence is a sequence that can be derived from another sequence by deleting some or
// no elements without changing the order of the remaining elements.
// Example 1:
//   Input: binary = "001"
//   Output: 2
//   Explanation: The good subsequences of binary are ["0", "0", "1"].
//     The unique good subsequences are "0" and "1".
// Example 2:
//   Input: binary = "11"
//   Output: 2
//   Explanation: The good subsequences of binary are ["1", "1", "11"].
//     The unique good subsequences are "1" and "11".
// Example 3:
//   Input: binary = "101"
//   Output: 5
//   Explanation: The good subsequences of binary are ["1", "0", "1", "10", "11", "101"].
//     The unique good subsequences are "0", "1", "10", "11", and "101".
// Constraints:
//   1 <= binary.length <= 10^5
//   binary consists of only '0's and '1's.

func numberOfUniqueGoodSubsequences(binary string) int {
	const p = int(1e9 + 7)
	var e0, e1, has0 int // the number of valid sequences that start with 1 and end with 0 or 1
	for _, v := range binary {
		if v == '0' {
			has0 = 1
			// append 0 to e0 and e1 sequences, make new unique sequences end with 0
			e0 = e0 + e1
			e0 %= p
		} else {
			// append 1 to e0 and e1 sequences, make new unique sequences end with 1,
			// the new 1 that can be a new sequence start
			e1 = e0 + e1 + 1
			e1 %= p
		}
	}
	return (e0 + e1 + has0) % p
}

func main() {
	for _, v := range []struct {
		b   string
		ans int
	}{
		{"001", 2},
		{"11", 2},
		{"101", 5},
	} {
		fmt.Println(numberOfUniqueGoodSubsequences(v.b), v.ans)
	}
}

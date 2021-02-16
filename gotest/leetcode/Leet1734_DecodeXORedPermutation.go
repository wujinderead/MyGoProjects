package main

import "fmt"

// https://leetcode.com/problems/decode-xored-permutation/

// There is an integer array perm that is a permutation of the first n positive
// integers, where n is always odd. It was encoded into another integer array encoded
// of length n - 1, such that encoded[i] = perm[i] XOR perm[i + 1]. For example,
// if perm = [1,3,2], then encoded = [2,1].
// Given the encoded array, return the original array perm. It is guaranteed that
// the answer exists and is unique.
// Example 1:
//   Input: encoded = [3,1]
//   Output: [1,2,3]
//   Explanation: If perm = [1,2,3], then encoded = [1 XOR 2,2 XOR 3] = [3,1]
// Example 2:
//   Input: encoded = [6,5,4,6]
//   Output: [2,4,1,5,3]
// Constraints:
//   3 <= n < 10^5
//   n is odd.
//   encoded.length == n - 1

// for example, n=5, we have encoded = [a0^a1, a1^a2, a2^a3, a3^a4]
// we can get a1^a2^a3^a4, and we know a0^...^a4=1^2^3^4^5, we can get a0.
func decode(encoded []int) []int {
	a0 := 0
	for i := 1; i < len(encoded); i += 2 {
		a0 = a0 ^ encoded[i]
	}
	for i := 1; i <= len(encoded)+1; i++ {
		a0 = a0 ^ i
	}
	ans := make([]int, len(encoded)+1)
	ans[0] = a0
	for i := 1; i < len(ans); i++ {
		ans[i] = ans[i-1] ^ encoded[i-1]
	}
	return ans
}

func main() {
	for _, v := range []struct {
		end, ans []int
	}{
		{[]int{3, 1}, []int{1, 2, 3}},
		{[]int{6, 5, 4, 6}, []int{2, 4, 1, 5, 3}},
	} {
		fmt.Println(decode(v.end), v.ans)
	}
}

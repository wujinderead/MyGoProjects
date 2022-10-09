package main

import "fmt"

// https://leetcode.com/problems/find-the-original-array-of-prefix-xor/

// You are given an integer array pref of size n. Find and return the array arr of size n that satisfies:
// pref[i] = arr[0] ^ arr[1] ^ ... ^ arr[i].
// Note that ^ denotes the bitwise-xor operation.
// It can be proven that the answer is unique.
// Example 1:
//   Input: pref = [5,2,0,3,1]
//   Output: [5,7,2,3,2]
//   Explanation: From the array [5,7,2,3,2] we have the following:
//     - pref[0] = 5.
//     - pref[1] = 5 ^ 7 = 2.
//     - pref[2] = 5 ^ 7 ^ 2 = 0.
//     - pref[3] = 5 ^ 7 ^ 2 ^ 3 = 3.
//     - pref[4] = 5 ^ 7 ^ 2 ^ 3 ^ 2 = 1.
// Example 2:
//   Input: pref = [13]
//   Output: [13]
//   Explanation: We have pref[0] = arr[0] = 13.
// Constraints:
//   1 <= pref.length <= 10⁵
//   0 <= pref[i] <= 10⁶

// p[0] = a[0]
// p[1] = p[0]^a[1]  -->  a[1] = p[0]^p[1]
// p[2] = p[1]^a[2]  -->  a[2] = p[1]^p[2]
func findArray(pref []int) []int {
	arr := make([]int, len(pref))
	arr[0] = pref[0]
	for i := 1; i < len(pref); i++ {
		arr[i] = pref[i] ^ pref[i-1]
	}
	return arr
}

func main() {
	for _, v := range []struct {
		pref, ans []int
	}{
		{[]int{5, 2, 0, 3, 1}, []int{5, 7, 2, 3, 2}},
		{[]int{13}, []int{13}},
	} {
		fmt.Println(findArray(v.pref), v.ans)
	}
}

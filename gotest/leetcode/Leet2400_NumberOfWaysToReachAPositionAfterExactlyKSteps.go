package main

import (
	"fmt"
	"math/big"
)

// https://leetcode.com/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/

// You are given two positive integers startPos and endPos. Initially, you are standing at position
// startPos on an infinite number line. With one step, you can move either one position to the left,
// or one position to the right.
// Given a positive integer k, return the number of different ways to reach the position endPos starting
// from startPos, such that you perform exactly k steps. Since the answer may be very large, return it
// modulo 10⁹ + 7.
// Two ways are considered different if the order of the steps made is not exactly the same.
// Note that the number line includes negative integers.
// Example 1:
//   Input: startPos = 1, endPos = 2, k = 3
//   Output: 3
//   Explanation: We can reach position 2 from 1 in exactly 3 steps in three ways:
//     - 1 -> 2 -> 3 -> 2.
//     - 1 -> 2 -> 1 -> 2.
//     - 1 -> 0 -> 1 -> 2.
//     It can be proven that no other way is possible, so we return 3.
// Example 2:
//   Input: startPos = 2, endPos = 5, k = 10
//   Output: 0
//   Explanation: It is impossible to reach position 5 from position 2 in exactly 10 steps.
// Constraints:
//   1 <= startPos, endPos, k <= 1000

func numberOfWays(startPos int, endPos int, k int) int {
	// x steps forward, k-x steps backward, so we have s+x-(k-x)=e, x=(e+k-s)/2
	// with condition x>=0, k-x>=0, (e+k-s)/2 == 0
	x := (endPos + k - startPos) / 2
	if (endPos+k-startPos)%2 != 0 || x < 0 || k-x < 0 {
		return 0
	}
	// x steps forward, total k steps, the number of way is C(k,x)
	if x > k-x {
		x = k - x
	}
	ans := 1
	p := int(1e9 + 7)
	pp := new(big.Int).SetInt64(int64(1e9 + 7))
	kk, fac := k, x
	for i := 0; i < x; i++ {
		ans = (ans * kk) % p
		ff := new(big.Int).SetInt64(int64(fac))
		ff.ModInverse(ff, pp)
		ans = (ans * int(ff.Int64())) % p
		kk--
		fac--
	}
	return ans
}

func main() {
	for _, v := range []struct {
		startPos, endPos, k, ans int
	}{
		{1, 2, 3, 3},
		{2, 5, 10, 0},
	} {
		fmt.Println(numberOfWays(v.startPos, v.endPos, v.k), v.ans)
	}
}

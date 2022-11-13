package main

import "fmt"

// https://leetcode.com/problems/minimum-addition-to-make-integer-beautiful/

// You are given two positive integers n and target.
// An integer is considered beautiful if the sum of its digits is less than or equal to target.
// Return the minimum non-negative integer x such that n + x is beautiful. The input will be
// generated such that it is always possible to make n beautiful.
// Example 1:
//   Input: n = 16, target = 6
//   Output: 4
//   Explanation: Initially n is 16 and its digit sum is 1 + 6 = 7. After adding 4,
//     n becomes 20 and digit sum becomes 2 + 0 = 2. It can be shown that we can not
//     make n beautiful with adding non-negative integer less than 4.
// Example 2:
//   Input: n = 467, target = 6
//   Output: 33
//   Explanation: Initially n is 467 and its digit sum is 4 + 6 + 7 = 17. After
//     adding 33, n becomes 500 and digit sum becomes 5 + 0 + 0 = 5. It can be shown that
//     we can not make n beautiful with adding non-negative integer less than 33.
// Example 3:
//   Input: n = 1, target = 1
//   Output: 0
//   Explanation: Initially n is 1 and its digit sum is 1, which is already smaller than or equal to target.
// Constraints:
//   1 <= n <= 10¹²
//   1 <= target <= 150
//   The input will be generated such that it is always possible to make n beautiful.

// try to make last digits all zero
// e.g., for n=9987, try add 10-7, 100-87, 1000-987, 10000-9987
func makeIntegerBeautiful(n int64, target int) int64 {
	add := 0
	x := 10
	for digitSum(int(n)) > target {
		nn := int(n)
		if nn%x == 0 {
			x *= 10
			continue
		}
		add = x - (nn % x)
		nn += add
		if digitSum(nn) <= target {
			break
		}
		x *= 10
	}
	return int64(add)
}

func digitSum(nn int) int {
	sum := 0
	for nn > 0 {
		sum += nn % 10
		nn /= 10
	}
	return sum
}

func main() {
	for _, v := range []struct {
		n           int64
		target, ans int
	}{
		{16, 6, 4},
		{467, 6, 33},
		{1, 1, 0},
	} {
		fmt.Println(makeIntegerBeautiful(v.n, v.target), v.ans)
	}
}

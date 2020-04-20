package main

import "fmt"

// https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/

// You have a pointer at index 0 in an array of size arrLen. At each step, you can
// move 1 position to the left, 1 position to the right in the array or stay in the
// same place (The pointer should not be placed outside the array at any time).
// Given two integers steps and arrLen, return the number of ways such that your
// pointer still at index 0 after exactly steps steps.
// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: steps = 3, arrLen = 2
//   Output: 4
//   Explanation: There are 4 differents ways to stay at index 0 after 3 steps.
//     Right, Left, Stay
//     Stay, Right, Left
//     Right, Stay, Left
//     Stay, Stay, Stay
// Example 2:
//   Input: steps = 2, arrLen = 4
//   Output: 2
//   Explanation: There are 2 different ways to stay at index 0 after 2 steps
//     Right, Left
//     Stay, Stay
// Example 3:
//   Input: steps = 4, arrLen = 2
//   Output: 8
// Constraints:
//   1 <= steps <= 500
//   1 <= arrLen <= 10^6

func numWays(steps int, arrLen int) int {
	l := steps + 1
	if l > arrLen {
		l = arrLen
	}
	old, new := make([]int, l), make([]int, l)
	old[0] = 1
	for i := 0; i < steps; i++ {
		for j := 0; j < len(old); j++ {
			new[j] = old[j]
			if j-1 >= 0 {
				new[j] += old[j-1]
			}
			if j+1 < len(old) {
				new[j] += old[j+1]
			}
			new[j] = new[j] % 1000000007
		}
		old, new = new, old
	}
	return old[0]
}

func main() {
	fmt.Println(numWays(3, 2))
	fmt.Println(numWays(2, 4))
	fmt.Println(numWays(4, 2))
	fmt.Println(numWays(1, 2))
	fmt.Println(numWays(2, 2))
	fmt.Println(numWays(100, 10))
	fmt.Println(numWays(100, 1000))
}

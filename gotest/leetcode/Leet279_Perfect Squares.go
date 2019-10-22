package leetcode

import "fmt"

// https://leetcode.com/problems/perfect-squares/

// Given a positive integer n, find the least number of perfect square numbers
// (for example, 1, 4, 9, 16, ...) which sum to n.
// Example 1:
//   Input: n = 12
//   Output: 3
//   Explanation: 12 = 4 + 4 + 4.
// Example 2:
//   Input: n = 13
//   Output: 2
//   Explanation: 13 = 4 + 9.

func numSquares(n int) int {
	sqrs := make(map[int]struct{})
	for i := 1; i*i <= n; i++ {
		sqrs[i*i] = struct{}{}
	}
	if _, ok := sqrs[n]; ok {
		return 1
	}
	ways := make([]int, n+1)
	ways[1] = 1
	for i := 2; i <= n; i++ {
		if _, ok := sqrs[i]; ok {
			ways[i] = 1
			continue
		}
		min := 0x7fffffff
		for j := 1; j <= i/2; j++ {
			if ways[j]+ways[i-j] < min {
				min = ways[j] + ways[i-j]
			}
		}
		ways[i] = min
	}
	//fmt.Println(ways)
	return ways[n]
}

func main() {
	fmt.Println(numSquares(1000))
}

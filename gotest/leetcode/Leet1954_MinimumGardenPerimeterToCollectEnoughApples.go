package main

import "fmt"

// https://leetcode.com/problems/minimum-garden-perimeter-to-collect-enough-apples/

// In a garden represented as an infinite 2D grid, there is an apple tree planted at every
// integer coordinate. The apple tree planted at an integer coordinate (i, j) has |i| + |j|
// apples growing on it.
// You will buy an axis-aligned square plot of land that is centered at (0, 0).
// Given an integer neededApples, return the minimum perimeter of a plot such that at least
// neededApples apples are inside or on the perimeter of that plot.
// The value of |x| is defined as:
//   x if x >= 0
//   -x if x < 0
// Example 1:
//   Input: neededApples = 1
//   Output: 8
//   Explanation: A square plot of side length 1 does not contain any apples.
//     However, a square plot of side length 2 has 12 apples inside (as depicted in t
//     he image above).
//     The perimeter is 2 * 4 = 8.
// Example 2:
//   Input: neededApples = 13
//   Output: 16
// Example 3:
//   Input: neededApples = 1000000000
//   Output: 5040
// Constraints:
//   1 <= neededApples <= 10^15

// if the square point is (n,n), the perimeter is 8n, and the apples in the square is 2n(n+1)(2n+1)
// e.g., n=3, we have x_coor = (-3,-2,-1,0,1,2,3) of in each row, there is 7 rows; and column is the same
// so the apples are (3+2+1+0+1+2+3)*7*2
// binary search it
func minimumPerimeter(neededApples int64) int64 {
	var l, r int64 = 1, 1000000
	for l < r {
		mid := (l + r) / 2
		apples := 2 * mid * (mid + 1) * (2*mid + 1)
		if neededApples <= apples {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return 8 * l
}

func main() {
	for _, v := range []struct {
		n, ans int64
	}{
		{1, 8},
		{11, 8},
		{12, 8},
		{13, 16},
		{59, 16},
		{60, 16},
		{61, 24},
		{1000000000, 5040},
		{1000000000000000, 503968},
	} {
		fmt.Println(minimumPerimeter(v.n), v.ans)
	}
}

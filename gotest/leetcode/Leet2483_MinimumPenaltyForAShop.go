package main

import "fmt"

// https://leetcode.com/problems/minimum-penalty-for-a-shop/

// You are given the customer visit log of a shop represented by a 0-indexed string
// customers consisting only of characters 'N' and 'Y':
//   if the iᵗʰ character is 'Y', it means that customers come at the iᵗʰ hour
//   whereas 'N' indicates that no customers come at the iᵗʰ hour.
// If the shop closes at the jᵗʰ hour (0 <= j <= n), the penalty is calculated as follows:
//   For every hour when the shop is open and no customers come, the penalty increases by 1.
//   For every hour when the shop is closed and customers come, the penalty increases by 1.
// Return the earliest hour at which the shop must be closed to incur a minimum penalty.
// Note that if a shop closes at the jᵗʰ hour, it means the shop is closed at the hour j.
// Example 1:
//   Input: customers = "YYNY"
//   Output: 2
//   Explanation:
//     - Closing the shop at the 0ᵗʰ hour incurs in 1+1+0+1 = 3 penalty.
//     - Closing the shop at the 1ˢᵗ hour incurs in 0+1+0+1 = 2 penalty.
//     - Closing the shop at the 2ⁿᵈ hour incurs in 0+0+0+1 = 1 penalty.
//     - Closing the shop at the 3ʳᵈ hour incurs in 0+0+1+1 = 2 penalty.
//     - Closing the shop at the 4ᵗʰ hour incurs in 0+0+1+0 = 1 penalty.
//     Closing the shop at 2ⁿᵈ or 4ᵗʰ hour gives a minimum penalty. Since 2 is earlier, the optimal closing time is 2.
// Example 2:
//   Input: customers = "NNNNN"
//   Output: 0
//   Explanation: It is best to close the shop at the 0ᵗʰ hour as no customers arrive.
// Example 3:
//   Input: customers = "YYYY"
//   Output: 4
//   Explanation: It is best to close the shop at the 4ᵗʰ hour as customers arrive at each hour.
// Constraints:
//   1 <= customers.length <= 10⁵
//   customers consists only of characters 'Y' and 'N'.

func bestClosingTime(customers string) int {
	var n, y int
	for i := range customers {
		if customers[i] == 'Y' {
			y++
		}
	}
	bestTime := 0
	bestPen := y
	for i := 0; i < len(customers); i++ {
		if customers[i] == 'N' {
			n++
		} else {
			y--
		}
		// the penalty is N in prefix plus Y in suffix
		if n+y < bestPen {
			bestPen = n + y
			bestTime = i + 1
		}
	}
	return bestTime
}

func main() {
	for _, v := range []struct {
		s   string
		ans int
	}{
		{"YYNY", 2},
		{"NNNNN", 0},
		{"YYYY", 4},
	} {
		fmt.Println(bestClosingTime(v.s), v.ans)
	}
}

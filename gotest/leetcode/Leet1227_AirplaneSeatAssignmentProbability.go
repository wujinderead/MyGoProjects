package leetcode

import "fmt"

// https://leetcode.com/problems/airplane-seat-assignment-probability

// n passengers board an airplane with exactly n seats. The first passenger
// has lost the ticket and picks a seat randomly. But after that,
// the rest of passengers will: Take their own seat if it is still available,
// Pick other seats randomly when they find their seat occupied.
// What is the probability that the n-th person can get his own seat?
// Example 1:
//   Input: n = 1
//   Output: 1.00000
//   Explanation: The first person can only get the first seat.
// Example 2:
//   Input: n = 2
//   Output: 0.50000
//   Explanation: The second person has a probability of 0.5 to get the second seat
//   (when first person gets the first seat).
// Constraints:
//   1 <= n <= 10^5

// f(n)= 1*(1/n) + 0*(1/n) + f(n-1)*(1/n) + f(n-2)*(n-3)/n
// which is always 1/2
func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1.0
	}
	return 0.5
}

func main() {
	fmt.Println(nthPersonGetsNthSeat(0))
	fmt.Println(nthPersonGetsNthSeat(1))
	fmt.Println(nthPersonGetsNthSeat(2))
	fmt.Println(nthPersonGetsNthSeat(3))
	fmt.Println(nthPersonGetsNthSeat(4))
	fmt.Println(nthPersonGetsNthSeat(5))
}

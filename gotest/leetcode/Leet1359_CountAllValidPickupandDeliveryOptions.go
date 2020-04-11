package main

import "fmt"

// https://leetcode.com/problems/count-all-valid-pickup-and-delivery-options/

// Given n orders, each order consist in pickup and delivery services.
// Count all valid pickup/delivery possible sequences such that delivery(i) is always after of pickup(i).
// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: n = 1
//   Output: 1
//   Explanation: Unique order (P1, D1), Delivery 1 always is after of Pickup 1.
// Example 2:
//   Input: n = 2
//   Output: 6
//   Explanation: All possible orders:
//     (P1,P2,D1,D2), (P1,P2,D2,D1), (P1,D1,P2,D2), (P2,P1,D1,D2), (P2,P1,D2,D1) and (P2,D2,P1,D1).
//     This is an invalid order (P1,D2,P2,D1) because Pickup 2 is after of Delivery 2.
// Example 3:
//   Input: n = 3
//   Output: 90
// Constraints:
//   1 <= n <= 500

// for n, it's generated from n-1 situation, e.g.:
// for a valid n=1 sequence, (P1,D1), we can insert P2 in any position,
// (p2, p1, d1), then d2 has 3 possible position (after p2, p1, d1)
// (p1, p2, d1), then d2 has 2 possible position (after p2, d1)
// (p1, d1, p2), then d2 has 1 possible position (after p2)
// we can figure out that F(n) = F(n-1) * 2n*(2n-1)/2
func countOrders(n int) int {
	f := 1
    for i:=2; i<=n; i++ {
    	mul := i*(2*i-1)
    	f = (f * mul) % 1000000007
	}
    return f
}


func main() {
	fmt.Println(countOrders(1))
	fmt.Println(countOrders(2))
	fmt.Println(countOrders(3))
	fmt.Println(countOrders(4))
}
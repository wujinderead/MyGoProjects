package main

import (
	"fmt"
)

// https://leetcode.com/problems/knight-dialer/

// A chess knight can move as indicated in the chess diagram below:           
//    *c*c*
//    c***c         1 2 3
//    **k**         4 5 6
//    c***c         7 8 9
//    *c*c*           0
// This time, we place our chess knight on any numbered key of a phone pad (indicated above), 
// and the knight makes N-1 hops. Each hop must be from one key to another numbered key.
// Each time it lands on a key (including the initial placement of the knight), 
// it presses the number of that key, pressing N digits total.
// How many distinct numbers can you dial in this manner?
// Since the answer may be large, output the answer modulo 10^9 + 7.
// Example 1:
//   Input: 1
//   Output: 10
// Example 2:
//   Input: 2
//   Output: 20
// Example 3:
//   Input: 3
//   Output: 46
// Note:
//   1 <= N <= 5000

// digit 4 and 6 have 3 next hops, digit 5 has none, other digits have 2 next hops
func knightDialer(N int) int {
	if N==1 {
		return 10
	}
	ways := [10]int{1,1,1,1,1,0,1,1,1,1}   // first hop
	newways := [10]int{}
	for i:=1; i<N; i++ {                // next n-1 hops
		newways[0] = ways[4]+ways[6]    // digit 4 and 6 can hop to digit 0
		newways[1] = ways[6]+ways[8]
		newways[2] = ways[7]+ways[9]
		newways[3] = ways[4]+ways[8]
		newways[4] = ways[3]+ways[9]+ways[0]
		newways[6] = ways[1]+ways[7]+ways[0]
		newways[7] = ways[2]+ways[6]
		newways[8] = ways[1]+ways[3]
		newways[9] = ways[2]+ways[4]
		for j := range ways {
			newways[j] = newways[j] % int(1e9+7)
		}
		ways, newways = newways, ways
	}
	total := 0
	for i := range ways {
		total += ways[i]
	}
	return total%int(1e9+7)
}

func main() {
	fmt.Println(knightDialer(1))
	fmt.Println(knightDialer(2))
	fmt.Println(knightDialer(3))
	fmt.Println(knightDialer(4))
	fmt.Println(knightDialer(100))
	fmt.Println(knightDialer(1000))
}
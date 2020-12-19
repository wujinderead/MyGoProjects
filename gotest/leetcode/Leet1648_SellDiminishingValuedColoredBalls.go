package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/sell-diminishing-valued-colored-balls/

// You have an inventory of different colored balls, and there is a customer that
// wants orders balls of any color.
// The customer weirdly values the colored balls. Each colored ball's value is the number of
// balls of that color you currently have in your inventory. For example, if you own 6 yellow
// balls, the customer would pay 6 for the first yellow ball. After the transaction, there are
// only 5 yellow balls left, so the next yellow ball is then valued at 5 (i.e., the value of
// the balls decreases as you sell more to the customer).
// You are given an integer array, inventory, where inventory[i] represents the number of balls
// of the ith color that you initially own. You are also given an integer orders, which
// represents the total number of balls that the customer wants. You can sell the balls in
// any order.
// Return the maximum total value that you can attain after selling orders colored balls.
// As the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: inventory = [2,5], orders = 4
//   Output: 14
//   Explanation: Sell the 1st color 1 time (2) and the 2nd color 3 times (5 + 4 + 3).
//     The maximum total value is 2 + 5 + 4 + 3 = 14.
// Example 2:
//   Input: inventory = [3,5], orders = 6
//   Output: 19
//   Explanation: Sell the 1st color 2 times (3 + 2) and the 2nd color 4 times (5 + 4 + 3 + 2).
//     The maximum total value is 3 + 2 + 5 + 4 + 3 + 2 = 19.
// Example 3:
//   Input: inventory = [2,8,4,10,6], orders = 20
//   Output: 110
// Example 4:
//   Input: inventory = [1000000000], orders = 1000000000
//   Output: 21
//   Explanation: Sell the 1st color 1000000000 times for a total value of 500000000500000000.
//     500000000500000000 modulo 109 + 7 = 21.
// Constraints:
//   1 <= inventory.length <= 10^5
//   1 <= inventory[i] <= 10^9
//   1 <= orders <= min(sum(inventory[i]), 10^9)

// e.g. inv=[1,3,3,5,5,5], orders=12
// firstly, order last 3 balls 2 times, got profit (5+4)*3
// then, inv=[1,3,3,3,3,3], orders=6
// order last 5 balls got 3*5 profit, and order another ball with profit 2
func maxProfit(inv []int, orders int) int {
	inv = append(inv, 0)
	sort.Sort(sort.IntSlice(inv))

	mod := int(1e9 + 7)
	ans := 0
	i := len(inv) - 1
	for i > 0 {
		for inv[i] == inv[i-1] {
			i--
		}
		if (inv[i]-inv[i-1])*(len(inv)-i) <= orders {
			orders -= (inv[i] - inv[i-1]) * (len(inv) - i)
			ans += (inv[i] + inv[i-1] + 1) * (inv[i] - inv[i-1]) * (len(inv) - i) / 2
			ans = ans % mod
			i--
		} else {
			div, res := orders/(len(inv)-i), orders%(len(inv)-i)
			ans += (inv[i]+inv[i]-div+1)*div*(len(inv)-i)/2 + (inv[i]-div)*res
			ans = ans % mod
			break
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		in      []int
		or, ans int
	}{
		{[]int{2, 5}, 4, 14},
		{[]int{3, 5}, 6, 19},
		{[]int{2, 8, 4, 10, 6}, 20, 110},
		{[]int{1000000000}, 1000000000, 21},
	} {
		fmt.Println(maxProfit(v.in, v.or), v.ans)
	}
}

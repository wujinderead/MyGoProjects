package main

import (
	"fmt"
)

// https://leetcode.com/problems/maximum-earnings-from-taxi/

// There are n points on a road you are driving your taxi on. The n points on the road are
// labeled from 1 to n in the direction you are going, and you want to drive from point 1 to
// point n to make money by picking up passengers. You cannot change the direction of the taxi.
// The passengers are represented by a 0-indexed 2D integer array rides, where rides[i] =
// [starti, endi, tipi] denotes the ith passenger requesting a ride from point starti to point endi
// who is willing to give a tipi dollar tip.
// For each passenger i you pick up, you earn endi - starti + tipi dollars. You may only drive at most
// one passenger at a time.
// Given n and rides, return the maximum number of dollars you can earn by picking up the passengers
// optimally.
// Note: You may drop off a passenger and pick up a different passenger at the same point.
// Example 1:
//   Input: n = 5, rides = [[2,5,4],[1,5,1]]
//   Output: 7
//   Explanation: We can pick up passenger 0 to earn 5 - 2 + 4 = 7 dollars.
// Example 2:
//   Input: n = 20, rides = [[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
//   Output: 20
//   Explanation: We will pick up the following passengers:
//     - Drive passenger 1 from point 3 to point 10 for a profit of 10 - 3 + 2 = 9 dollars.
//     - Drive passenger 2 from point 10 to point 12 for a profit of 12 - 10 + 3 = 5 dollars.
//     - Drive passenger 5 from point 13 to point 18 for a profit of 18 - 13 + 1 = 6 dollars.
//     We earn 9 + 5 + 6 = 20 dollars in total.
// Constraints:
//   1 <= n <= 10^5
//   1 <= rides.length <= 3 * 10^4
//   rides[i].length == 3
//   1 <= starti < endi <= n
//   1 <= tipi <= 10^5

// a typical 0-1 knapsack problem
func maxTaxiEarnings(n int, rides [][]int) int64 {
	starts := make(map[int][][2]int)
	for _, v := range rides {
		starts[v[0]] = append(starts[v[0]], [2]int{v[1], v[2]})
	}
	dp := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		for _, v := range starts[i] { // for all rides start at i
			end, tip := v[0], v[1]
			if dp[end]+(end-i+tip) > dp[i] {
				dp[i] = dp[end] + (end - i + tip)
			}
		}
	}
	return int64(dp[0])
}

func main() {
	for _, v := range []struct {
		n   int
		r   [][]int
		ans int64
	}{
		{5, [][]int{{2, 5, 4}, {1, 5, 1}}, 7},
		{20, [][]int{{1, 6, 1}, {3, 10, 2}, {10, 12, 3}, {11, 12, 2}, {12, 15, 2}, {13, 18, 1}}, 20},
		{10, [][]int{{2, 3, 6}, {8, 9, 8}, {5, 9, 7}, {8, 9, 1}, {2, 9, 2}, {9, 10, 6}, {7, 10, 10}, {6, 7, 9}, {4, 9, 7}, {2, 3, 1}}, 33},
	} {
		fmt.Println(maxTaxiEarnings(v.n, v.r), v.ans)
	}
}

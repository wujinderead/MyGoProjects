package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/count-all-possible-routes

// You are given an array of distinct positive integers locations where locations[i] represents
// the position of city i. You are also given integers start, finish and fuel representing the
// starting city, ending city, and the initial amount of fuel you have, respectively.
// At each step, if you are at city i, you can pick any city j such that j != i 
// and 0 <= j < locations.length and move to city j. Moving from city i to city j reduces the amount
// of fuel you have by |locations[i] - locations[j]|. Please notice that |x| denotes the absolute value of x.
// Notice that fuel cannot become negative at any point in time, and that you are allowed to visit
// any city more than once (including start and finish).
// Return the count of all possible routes from start to finish.
// Since the answer may be too large, return it modulo 10^9 + 7.
// Example 1:
//   Input: locations = [2,3,6,8,4], start = 1, finish = 3, fuel = 5
//   Output: 4
//   Explanation: The following are all possible routes, each uses 5 units of fuel:
//     1 -> 3
//     1 -> 2 -> 3
//     1 -> 4 -> 3
//     1 -> 4 -> 2 -> 3
// Example 2:
//   Input: locations = [4,3,1], start = 1, finish = 0, fuel = 6
//   Output: 5
//   Explanation: The following are all possible routes:
//     1 -> 0, used fuel = 1
//     1 -> 2 -> 0, used fuel = 5
//     1 -> 2 -> 1 -> 0, used fuel = 5
//     1 -> 0 -> 1 -> 0, used fuel = 3
//     1 -> 0 -> 1 -> 0 -> 1 -> 0, used fuel = 5
// Example 3:
//   Input: locations = [5,2,1], start = 0, finish = 2, fuel = 3
//   Output: 0
//   Explanation: It's impossible to get from 0 to 2 using only 3 units of fuel
//     since the shortest route needs 4 units of fuel.
// Example 4:
//   Input: locations = [2,1,5], start = 0, finish = 0, fuel = 3
//   Output: 2
//   Explanation: There are two possible routes, 0 and 0 -> 1 -> 0.
// Example 5:
//   Input: locations = [1,2,3], start = 0, finish = 2, fuel = 40
//   Output: 615088286
//   Explanation: The total number of possible routes is 2615088300.
//     Taking this number modulo 10^9 + 7 gives us 615088286.
// Constraints:
//   2 <= locations.length <= 100
//   1 <= locations[i] <= 10^9
//   All integers in locations are distinct.
//   0 <= start, finish < locations.length
//   1 <= fuel <= 200

func countRoutes(locations []int, start int, finish int, fuel int) int {
	// first check if fuel is enough
	s, f := locations[start], locations[finish]
	abs := s-f
	if abs < 0 {
		abs = -abs
	}
    if abs > fuel {
    	return 0
	}
	mod := int(1e9+7)

	// sort locations and get initial index
	sort.Sort(sort.IntSlice(locations))
	si, fi := 0, 0
	for i := range locations {
		if locations[i]==s {
			si = i
		}
		if locations[i]==f {
			fi = i
		}
	}

	// let dp(i, j) be the number of ways to loc[fi] with start point at loc[i] and fuel j.
	// then dp(i, j) is the sum of dp( x ,  j-|loc[x]-loc[i]| )
	dp := [100][201]int{}
	dp[fi][0] = 1   // 0 fuel at finish point: only one way, which is not move
	for j:=1; j<=fuel; j++ {
		for i:=0; i<len(locations); i++ {
			for x:=i-1; x>=0 && j-(locations[i]-locations[x])>=0; x-- {
				dp[i][j] += dp[x][j-(locations[i]-locations[x])]
			}
			for x:=i+1; x<len(locations) && j-(locations[x]-locations[i])>=0; x++ {
				dp[i][j] += dp[x][j-(locations[x]-locations[i])]
			}
			dp[i][j] = dp[i][j] % mod
		}
	}

	// return answer
	ans := 0
	for j:=0; j<=fuel; j++ {
		ans += dp[si][j]
	}
	return ans % mod
}

func main() {
	for _, v := range []struct{loc []int; s, f, fu, ans int} {
		{[]int{2,3,6,8,4}, 1, 3, 5, 4},
		{[]int{4,3,1}, 1, 0, 6, 5},
		{[]int{5,2,1}, 0, 2, 3, 0},
		{[]int{2,1,5}, 0, 0, 3, 2},
		{[]int{1,2,3}, 0, 2, 40, 615088286},
		{[]int{0,4}, 0, 1, 11, 1},
		{[]int{0,4}, 0, 1, 13, 2},
		{[]int{0}, 0, 0, 100, 1},
	} {
		fmt.Println(countRoutes(v.loc, v.s, v.f, v.fu), v.ans)
	}
}
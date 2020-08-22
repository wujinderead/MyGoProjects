package main

import "fmt"

// https://leetcode.com/problems/minimum-number-of-days-to-eat-n-oranges/

// There are n oranges in the kitchen and you decided to eat some of these oranges every day as follows:
// . Eat one orange.
// . If the number of remaining oranges (n) is divisible by 2 then you can eat n/2 oranges.
// . If the number of remaining oranges (n) is divisible by 3 then you can eat 2*(n/3) oranges.
// You can only choose one of the actions per day.
// Return the minimum number of days to eat n oranges.
// Example 1:
//   Input: n = 10
//   Output: 4
//   Explanation: You have 10 oranges.
//     Day 1: Eat 1 orange,  10 - 1 = 9.
//     Day 2: Eat 6 oranges, 9 - 2*(9/3) = 9 - 6 = 3. (Since 9 is divisible by 3)
//     Day 3: Eat 2 oranges, 3 - 2*(3/3) = 3 - 2 = 1.
//     Day 4: Eat the last orange  1 - 1  = 0.
//     You need at least 4 days to eat the 10 oranges.
// Example 2:
//   Input: n = 6
//   Output: 3
//   Explanation: You have 6 oranges.
//     Day 1: Eat 3 oranges, 6 - 6/2 = 6 - 3 = 3. (Since 6 is divisible by 2).
//     Day 2: Eat 2 oranges, 3 - 2*(3/3) = 3 - 2 = 1. (Since 3 is divisible by 3)
//     Day 3: Eat the last orange  1 - 1  = 0.
//     You need at least 3 days to eat the 6 oranges.
// Example 3:
//   Input: n = 1
//   Output: 1
// Example 4:
//   Input: n = 56
//   Output: 6
// Constraints: 
//   1 <= n <= 2*10^9

// dp with memorization
// for n=2*10^9, answer is 32, map has 259 elements, helper func called 515 times.
func minDays(n int) int {
	dp := make(map[int]int)
	dp[0] = 0
	dp[1] = 1
	v := helper(dp, n)
	// fmt.Println(len(dp))
	return v	
}

func helper(dp map[int]int, n int) int {
	if v, ok := dp[n]; ok {
		return v
	}
	v := 1+min(n%2+helper(dp, n/2), n%3+helper(dp, n/3))  // try the shortest path by divide 2 or 3
	dp[n] = v
	return v
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	for _, v := range [][2]int{
		{10, 4},
		{6, 3},
		{1, 1},
		{2, 2},
		{3, 2}, 
		{4, 3},
		{5, 4},
		{56, 6},
		{161, 9},
		{370, 10},
		{10000, 15},
		{10001, 16},
		{1421, 13},
		{69652, 19},
	} {
		fmt.Println(minDays(v[0]), v[1])
	}
	fmt.Println(minDays(2*int(1e9)))
}
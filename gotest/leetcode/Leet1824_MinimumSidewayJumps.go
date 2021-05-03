package main

import "fmt"

// https://leetcode.com/problems/minimum-sideway-jumps/

// There is a 3 lane road of length n that consists of n + 1 points labeled from 0 to n.
// A frog starts at point 0 in the second lane and wants to jump to point n. However,
// there could be obstacles along the way.
// You are given an array obstacles of length n + 1 where each obstacles[i] (ranging from 0 to 3)
// describes an obstacle on the lane obstacles[i] at point i. If obstacles[i] == 0, there are no
// obstacles at point i. There will be at most one obstacle in the 3 lanes at each point.
// For example, if obstacles[2] == 1, then there is an obstacle on lane 1 at point 2.
// The frog can only travel from point i to point i + 1 on the same lane if there is not an
// obstacle on the lane at point i + 1. To avoid obstacles, the frog can also perform a side
// jump to jump to another lane (even if they are not adjacent) at the same point if there
// is no obstacle on the new lane.
// For example, the frog can jump from lane 3 at point 3 to lane 1 at point 3.
// Return the minimum number of side jumps the frog needs to reach any lane at point n starting
// from lane 2 at point 0.
// Note: There will be no obstacles on points 0 and n.
// Example 1:
//   Input: obstacles = [0,1,2,3,0]
//   Output: 2
//   Explanation: The optimal solution is shown by the arrows above. There are 2 side jumps
//   (red arrows). Note that the frog can jump over obstacles only when making side jumps
//   (as shown at point 2).
// Example 2:
//   Input: obstacles = [0,1,1,3,3,0]
//   Output: 0
//   Explanation: There are no obstacles on lane 2. No side jumps are required.
// Example 3:
//   Input: obstacles = [0,2,1,0,3,0]
//   Output: 2
//   Explanation: The optimal solution is shown by the arrows above. There are 2 side jumps.
// Constraints:
//   obstacles.length == n + 1
//   1 <= n <= 5 * 10^5
//   0 <= obstacles[i] <= 3
//   obstacles[0] == obstacles[n] == 0

func minSideJumps(obstacles []int) int {
	const max = int(1e9)
	dp := make([][4]int, len(obstacles))
	dp[0][1], dp[0][3] = 1, 1 // initial position is at dp[0][2], need one move to dp[0][1] and dp[0][3]
	for i := 1; i < len(dp); i++ {
		for j := 1; j <= 3; j++ {
			if obstacles[i] == j { // position[i][j] is obstacle
				dp[i][j] = max
				continue
			}
			// position[i][j] empty
			if obstacles[i-1] == j { // position[i-1][j] is obstacle
				minv := int(1e10)
				for k := 1; k <= 3; k++ {
					// if position[i][k] and position[i-1][k] both empty
					// dp[i][j] = dp[i-1][k]+1
					if obstacles[i-1] != k && obstacles[i] != k {
						minv = min(minv, dp[i-1][k]+1)
					}
				}
				dp[i][j] = minv
			} else { // position[i-1][j] is empty
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return min(min(dp[len(dp)-1][1], dp[len(dp)-1][2]), dp[len(dp)-1][3])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		ob  []int
		ans int
	}{
		{[]int{0, 1, 1, 3, 3, 0}, 0},
		{[]int{0, 2, 1, 0, 3, 0}, 2},
	} {
		fmt.Println(minSideJumps(v.ob), v.ans)
	}
}

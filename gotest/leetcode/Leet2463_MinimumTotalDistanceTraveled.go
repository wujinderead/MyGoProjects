package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-total-distance-traveled/

// There are some robots and factories on the X-axis. You are given an integer array robot where
// robot[i] is the position of the iᵗʰ robot. You are also given a 2D integer array factory where
// factory[j] = [positionj, limitj] indicates that positionj is the position of the jᵗʰ factory and
// that the jᵗʰ factory can repair at most limitj robots.
// The positions of each robot are unique. The positions of each factory are also unique. Note that
// a robot can be in the same position as a factory initially.
// All the robots are initially broken; they keep moving in one direction. The direction could be the
// negative or the positive direction of the X-axis. When a robot reaches a factory that did not reach
// its limit, the factory repairs the robot, and it stops moving.
// At any moment, you can set the initial direction of moving for some robot.
// Your target is to minimize the total distance traveled by all the robots.
// Return the minimum total distance traveled by all the robots. The test cases are generated such that
// all the robots can be repaired.
// Note that
//   All robots move at the same speed.
//   If two robots move in the same direction, they will never collide.
//   If two robots move in opposite directions and they meet at some point, they do not collide. They cross each other.
//   If a robot passes by a factory that reached its limits, it crosses it as if it does not exist.
//   If the robot moved from a position x to a position y, the distance it moved is |y - x|.
// Example 1:
//   Input: robot = [0,4,6], factory = [[2,2],[6,2]]
//   Output: 4
//   Explanation: As shown in the figure:
//     - The first robot at position 0 moves in the positive direction. It will be
//     repaired at the first factory.
//     - The second robot at position 4 moves in the negative direction. It will be
//     repaired at the first factory.
//     - The third robot at position 6 will be repaired at the second factory. It does
//     not need to move.
//     The limit of the first factory is 2, and it fixed 2 robots.
//     The limit of the second factory is 2, and it fixed 1 robot.
//     The total distance is |2 - 0| + |2 - 4| + |6 - 6| = 4. It can be shown that we
//     cannot achieve a better total distance than 4.
// Example 2:
//   Input: robot = [1,-1], factory = [[-2,1],[2,1]]
//   Output: 2
//   Explanation: As shown in the figure:
//     - The first robot at position 1 moves in the positive direction. It will be
//     repaired at the second factory.
//     - The second robot at position -1 moves in the negative direction. It will be
//     repaired at the first factory.
//     The limit of the first factory is 1, and it fixed 1 robot.
//     The limit of the second factory is 1, and it fixed 1 robot.
//     The total distance is |2 - 1| + |(-2) - (-1)| = 2. It can be shown that we
//     cannot achieve a better total distance than 2.
// Constraints:
//   1 <= robot.length, factory.length <= 100
//   factory[j].length == 2
//   -10⁹ <= robot[i], positionj <= 10⁹
//   0 <= limitj <= robot.length
//   The input will be generated such that it is always possible to repair every robot.

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Ints(robot)
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})
	// dp(i, j): factories[0...i] to repair robots[:j]
	olds := make([]int, len(robot)+1) // dp[i-1][j]
	news := make([]int, len(robot)+1) // dp[i][j]
	for j := 1; j < len(olds); j++ {
		olds[j] = int(1e18)
		news[j] = int(1e18)
	}
	cap := 0
	for i := 0; i < len(factory); i++ { // factories[0...i]
		curcap := factory[i][1]
		curplace := factory[i][0]
		cap += curcap                                  // factories[0...i] can hold sum(factory[0...i][1]) robots
		for j := 1; j <= cap && j <= len(robot); j++ { // robots[:j]
			news[j] = olds[j] // dp[i][j] = dp[i-1][j]
			dist := 0
			for k := 1; k <= j && k <= curcap; k++ {
				dist += abs(robot[j-k] - curplace)
				// factory[i] repair last k robots, factory[i-1] repair first j-k robots
				news[j] = min(news[j], olds[j-k]+dist)
			}
		}
		olds, news = news, olds
	}
	return int64(olds[len(robot)])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	for _, v := range []struct {
		robot   []int
		factory [][]int
		ans     int64
	}{
		{[]int{0, 4, 6}, [][]int{{2, 2}, {6, 2}}, 4},
		{[]int{0, 5, 7}, [][]int{{2, 2}, {6, 2}}, 4},
		{[]int{1, -1}, [][]int{{-2, 1}, {2, 1}}, 2},
		{[]int{9, 11, 99, 101}, [][]int{{10, 1}, {7, 1}, {14, 1}, {100, 1}, {96, 1}, {103, 1}}, 6},
	} {
		fmt.Println(minimumTotalDistance(v.robot, v.factory), v.ans)
	}
}

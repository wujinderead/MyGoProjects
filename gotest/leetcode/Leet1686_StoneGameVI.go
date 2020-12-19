package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/stone-game-vi/

// Alice and Bob take turns playing a game, with Alice starting first.
// There are n stones in a pile. On each player's turn, they can remove a stone from the pile
// and receive points based on the stone's value. Alice and Bob may value the stones differently.
// You are given two integer arrays of length n, aliceValues and bobValues. Each aliceValues[i]
// and bobValues[i] represents how Alice and Bob, respectively, value the ith stone.
// The winner is the person with the most points after all the stones are chosen. If both players
// have the same amount of points, the game results in a draw. Both players will play optimally.
// Determine the result of the game, and:
//   If Alice wins, return 1.
//   If Bob wins, return -1.
//   If the game results in a draw, return 0.
// Example 1:
//   Input: aliceValues = [1,3], bobValues = [2,1]
//   Output: 1
//   Explanation:
//     If Alice takes stone 1 (0-indexed) first, Alice will receive 3 points.
//     Bob can only choose stone 0, and will only receive 2 points.
//     Alice wins.
// Example 2:
//   Input: aliceValues = [1,2], bobValues = [3,1]
//   Output: 0
//   Explanation:
//     If Alice takes stone 0, and Bob takes stone 1, they will both have 1 point.
//     Draw.
// Example 3:
//   Input: aliceValues = [2,4,3], bobValues = [1,6,7]
//   Output: -1
//   Explanation:
//     Regardless of how Alice plays, Bob will be able to have more points than Alice.
//     For example, if Alice takes stone 1, Bob can take stone 2, and Alice takes stone 0,
//     Alice will have 6 points to Bob's 7. Bob wins.
// Constraints:
//   n == aliceValues.length == bobValues.length
//   1 <= n <= 10^5
//   1 <= aliceValues[i], bobValues[i] <= 100

// if one stone score x for alice and y for bob; if alice take it, she score x,
// also prevent bob score y, so the benefit is x+y.
// e.g., alice=[7,4],   alice take 7, bob take 6, diff is 1
//         bob=[3,6],   alice take 4, bob take 3, diff is 1
// or,   alice=[7,6],   alice take 7, bob take 4, diff is 3
//         bob=[3,4],   alice take 6, bob take 3, diff is 3
// so if sum is equal, no matter which take first, the diff is fixed.
// math proof:
//   For n=2, we have [a1,a2], [b1,b2]
//   For Alice, she can choose a1, so the diff is a1-b2, or choose a2, gets a2-b1
//   if a1+b1==a2+b2, then a1-b2=a2-b1, pick anyone
//   if a1+b1>a2+b2, then a1-b2>a2-b1, pick a1
//   so sort by a+b
func stoneGameVI(aliceValues []int, bobValues []int) int {
	benefit := make([][2]int, len(aliceValues))
	for i := range benefit {
		benefit[i][0] = aliceValues[i] + bobValues[i]
		benefit[i][1] = aliceValues[i]
	}
	sort.Slice(benefit, func(i, j int) bool {
		return benefit[i][0] > benefit[j][0]
	})

	var aliScore, bobScore int
	for i := range benefit {
		if i%2 == 0 {
			aliScore += benefit[i][1]
		} else {
			bobScore += benefit[i][0] - benefit[i][1]
		}
	}

	if aliScore > bobScore {
		return 1
	} else if aliScore < bobScore {
		return -1
	}
	return 0
}

func main() {
	for _, v := range []struct {
		a, b []int
		ans  int
	}{
		{[]int{1, 3}, []int{2, 1}, 1},
		{[]int{1, 2}, []int{3, 1}, 0},
		{[]int{2, 4, 3}, []int{1, 6, 7}, -1},
	} {
		fmt.Println(stoneGameVI(v.a, v.b), v.ans)
	}
}

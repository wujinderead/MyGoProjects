package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/maximum-score-from-removing-stones/

// You are playing a solitaire game with three piles of stones of sizes a, b, and c
// respectively. Each turn you choose two different non-empty piles, take one stone
// from each, and add 1 point to your score. The game stops when there are fewer than
// two non-empty piles (meaning there are no more available moves).
// Given three integers a, b, and c, return the maximum score you can get.
// Example 1:
//   Input: a = 2, b = 4, c = 6
//   Output: 6
//   Explanation: The starting state is (2, 4, 6). One optimal set of moves is:
//     - Take from 1st and 3rd piles, state is now (1, 4, 5)
//     - Take from 1st and 3rd piles, state is now (0, 4, 4)
//     - Take from 2nd and 3rd piles, state is now (0, 3, 3)
//     - Take from 2nd and 3rd piles, state is now (0, 2, 2)
//     - Take from 2nd and 3rd piles, state is now (0, 1, 1)
//     - Take from 2nd and 3rd piles, state is now (0, 0, 0)
//     There are fewer than two non-empty piles, so the game ends. Total: 6 points.
// Example 2:
//   Input: a = 4, b = 4, c = 6
//   Output: 7
//   Explanation: The starting state is (4, 4, 6). One optimal set of moves is:
//     - Take from 1st and 2nd piles, state is now (3, 3, 6)
//     - Take from 1st and 3rd piles, state is now (2, 3, 5)
//     - Take from 1st and 3rd piles, state is now (1, 3, 4)
//     - Take from 1st and 3rd piles, state is now (0, 3, 3)
//     - Take from 2nd and 3rd piles, state is now (0, 2, 2)
//     - Take from 2nd and 3rd piles, state is now (0, 1, 1)
//     - Take from 2nd and 3rd piles, state is now (0, 0, 0)
//     There are fewer than two non-empty piles, so the game ends. Total: 7 points.
// Example 3:
//   Input: a = 1, b = 8, c = 8
//   Output: 8
//   Explanation: One optimal set of moves is to take from the 2nd and 3rd piles
//     for 8 turns until they are empty. After that, there are fewer than two
//     non-empty piles, so the game ends.
// Constraints:
//   1 <= a, b, c <= 10^5

// https://leetcode.com/problems/maximum-score-from-removing-stones/discuss/1053491/One-line-Python-O(1)
// assuming a <= b <= c, there are two cases.
// case 1 : a+b <= c , which means we can at max pair every a,b with c
// case 2 : a+b > c, which means pair every c with the pile of a,b and
// then pair the stones of a,b with each other.
func maximumScore(a int, b int, c int) int {
	ints := []int{a, b, c}
	sort.Ints(ints)
	a, b, c = ints[0], ints[1], ints[2]
	c = min(a+b, c)
	return (a + b + c) / 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		a, b, c, ans int
	}{
		{2, 4, 6, 6},
		{4, 4, 6, 7},
		{1, 8, 8, 8},
		{1, 1, 8, 2},
	} {
		fmt.Println(maximumScore(v.a, v.b, v.c), v.ans)
	}
}

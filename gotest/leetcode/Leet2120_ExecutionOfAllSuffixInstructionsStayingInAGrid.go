package main

import "fmt"

// https://leetcode.com/problems/execution-of-all-suffix-instructions-staying-in-a-grid/

// There is an n x n grid, with the top-left cell at (0, 0) and the bottom-right cell at (n - 1, n - 1).
// You are given the integer n and an integer array startPos where startPos = [startrow, startcol] indicates
// that a robot is initially at cell (startrow, startcol).
// You are also given a 0-indexed string s of length m where s[i] is the iᵗʰ instruction for the robot:
// 'L' (move left), 'R' (move right), 'U' (move up), and 'D' (move down).
// The robot can begin executing from any iᵗʰ instruction in s. It executes the instructions one by one
// towards the end of s but it stops if either of these conditions is met:
//   The next instruction will move the robot off the grid.
//   There are no more instructions left to execute.
// Return an array answer of length m where answer[i] is the number of instructions the robot can execute
// if the robot begins executing from the iᵗʰ instruction in s.
// Example 1:
//   Input: n = 3, startPos = [0,1], s = "RRDDLU"
//   Output: [1,5,4,3,1,0]
//   Explanation: Starting from startPos and beginning execution from the iᵗʰ instruction:
//     - 0ᵗʰ: "RRDDLU". Only one instruction "R" can be executed before it moves off the grid.
//     - 1ˢᵗ:  "RDDLU". All five instructions can be executed while it stays in the grid and ends at (1, 1).
//     - 2ⁿᵈ:   "DDLU". All four instructions can be executed while it stays in the grid and ends at (1, 0).
//     - 3ʳᵈ:    "DLU". All three instructions can be executed while it stays in the grid and ends at (0, 0).
//     - 4ᵗʰ:     "LU". Only one instruction "L" can be executed before it moves off the grid.
//     - 5ᵗʰ:      "U". If moving up, it would move off the grid.
// Example 2:
//   Input: n = 2, startPos = [1,1], s = "LURD"
//   Output: [4,1,0,0]
//   Explanation:
//     - 0ᵗʰ: "LURD".
//     - 1ˢᵗ:  "URD".
//     - 2ⁿᵈ:   "RD".
//     - 3ʳᵈ:    "D".
// Example 3:
//   Input: n = 1, startPos = [0,0], s = "LRUD"
//   Output: [0,0,0,0]
//   Explanation: No matter which instruction the robot begins execution from, it would move off the grid.
// Constraints:
//   m == s.length
//   1 <= n, m <= 500
//   startPos.length == 2
//   0 <= startrow, startcol < n
//   s consists of 'L', 'R', 'U', and 'D'.

func executeInstructions(n int, startPos []int, s string) []int {
	// how many step can each direction go
	max := [4]int{startPos[1], n - 1 - startPos[1], startPos[0], n - 1 - startPos[0]}
	ans := make([]int, len(s))
	pos := [4][]int{}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'L' { // found an L, add to L-list
			pos[0] = append(pos[0], i) // this L can countervail an R
			if len(pos[1]) > 0 {
				pos[1] = pos[1][:len(pos[1])-1]
			}
		} else if s[i] == 'R' {
			pos[1] = append(pos[1], i)
			if len(pos[0]) > 0 {
				pos[0] = pos[0][:len(pos[0])-1]
			}
		} else if s[i] == 'U' {
			pos[2] = append(pos[2], i)
			if len(pos[3]) > 0 {
				pos[3] = pos[3][:len(pos[3])-1]
			}
		} else {
			pos[3] = append(pos[3], i)
			if len(pos[2]) > 0 {
				pos[2] = pos[2][:len(pos[2])-1]
			}
		}
		m := len(s) - i
		for j := 0; j < 4; j++ {
			// for each direction, check how many in each list, determine how long we can go
			x := len(s) - i
			if len(pos[j]) > max[j] {
				x = pos[j][len(pos[j])-max[j]-1] - i
			}
			if x < m {
				m = x
			}
		}
		ans[i] = m
	}
	return ans
}

func main() {
	for _, v := range []struct {
		n   int
		s   []int
		ss  string
		ans []int
	}{
		{3, []int{0, 1}, "RRDDLU", []int{1, 5, 4, 3, 1, 0}},
		{2, []int{1, 1}, "LURD", []int{4, 1, 0, 0}},
		{1, []int{0, 0}, "LURD", []int{0, 0, 0, 0}},
	} {
		fmt.Println(executeInstructions(v.n, v.s, v.ss), v.ans)
	}
}

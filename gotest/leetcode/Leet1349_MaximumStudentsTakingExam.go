package main

import (
	"fmt"
)

// https://leetcode.com/problems/maximum-students-taking-exam/

// Given a m * n matrix seats that represent seats distributions in a classroom.
// If a seat is broken, it is denoted by '#' character otherwise it is denoted
// by a '.' character.
// Students can see the answers of those sitting next to the left, right, upper
// left and upper right, but he cannot see the answers of the student sitting directly
// in front or behind him. Return the maximum number of students that can take
// the exam together without any cheating being possible.
// Students must be placed in seats in good condition.
// Example 1:
//   Input: seats = [["#",".","#","#",".","#"],
//                   [".","#","#","#","#","."],
//                   ["#",".","#","#",".","#"]]
//   Output: 4
//   Explanation: Teacher can place 4 students in available seats so they don't
//   cheat on the exam.
// Example 2:
//   Input: seats = [[".","#"],
//                   ["#","#"],
//                   ["#","."],
//                   ["#","#"],
//                   [".","#"]]
//   Output: 3
//   Explanation: Place all students in available seats.
// Example 3:
//   Input: seats = [["#",".",".",".","#"],
//                   [".","#",".","#","."],
//                   [".",".","#",".","."],
//                   [".","#",".","#","."],
//                   ["#",".",".",".","#"]]
//   Output: 10
//   Explanation: Place students in available seats in column 1, 3 and 5.
// Constraints:
//   seats contains only characters '.' and'#'.
//   m == seats.length
//   n == seats[i].length
//   1 <= m <= 8
//   1 <= n <= 8

func maxStudents(seats [][]byte) int {
	// the board's state can be represented by a int64
	board := int64(0)
	for i := 0; i < len(seats); i++ {
		line := int64(0)
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == '.' {
				line |= 1 << uint(j)
			}
		}
		board |= line << uint(i*8)
	}
	// let F(board) be the max students the board can contain, then
	// for a position on board, F(board) has two candidate:
	//   1 + F(this position is occupied and associate positions are invalid)
	//   F(this position is invalid (but not occupied so won't interfere other))
	// base case: F(board)=1 when there is only one valid position.
	// we use a memorable recursive.
	mapp := make(map[int64]int)
	mapp[0] = 0
	m, n := len(seats), len(seats[0])
	return F(board, m, n, mapp)
}

func F(board int64, m, n int, mapp map[int64]int) int {
	if v, ok := mapp[board]; ok {
		return v
	}
	i := int64(0)
	for i = 0; i < 64; i++ {
		if board&(1<<uint(i)) == 1<<uint(i) { // find a valid place
			break
		}
	}
	afterboard := board ^ (1 << uint(i)) // if not occupy this position, mark it invalid
	cand := F(afterboard, m, n, mapp)
	// if we occupy this position, then we should invalid its left, right, down-left and down right.
	// (no need to check left, the left and upper of this position are all invalid)
	row := int(i / 8)
	col := int(i % 8)
	if col+1 < n && board&(1<<uint(i+1)) == 1<<uint(i+1) { // right is i+1
		afterboard = afterboard ^ (1 << uint(i+1))
	}
	if row+1 < m && col-1 >= 0 && board&(1<<uint(i+7)) == 1<<uint(i+7) { // left down is i+7
		afterboard = afterboard ^ (1 << uint(i+7))
	}
	if row+1 < m && col+1 < n && board&(1<<uint(i+9)) == 1<<uint(i+9) { // right down is i+9
		afterboard = afterboard ^ (1 << uint(i+9))
	}
	cand = max(1+F(afterboard, m, n, mapp), cand)
	mapp[board] = cand
	return cand
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxStudents([][]byte{
		{'#', '#'},
		{'#', '.'},
		{'#', '#'},
	}))
	fmt.Println(maxStudents([][]byte{
		{'#'},
	}))
	fmt.Println(maxStudents([][]byte{
		{'.', '.'},
		{'.', '.'},
		{'.', '.'},
	}))
	fmt.Println(maxStudents([][]byte{
		{'.', '.'},
		{'.', '.'},
	}))
	fmt.Println(maxStudents([][]byte{
		{'.', '.'},
	}))
	fmt.Println(maxStudents([][]byte{
		{'#', '.', '#', '#', '.', '#'},
		{'.', '#', '#', '#', '#', '.'},
		{'#', '.', '#', '#', '.', '#'},
	}))
	fmt.Println(maxStudents([][]byte{
		{'.', '#'},
		{'#', '#'},
		{'#', '.'},
		{'#', '#'},
		{'.', '#'},
	}))
	fmt.Println(maxStudents([][]byte{
		{'#', '.', '.', '.', '#'},
		{'.', '#', '.', '#', '.'},
		{'.', '.', '#', '.', '#'},
		{'.', '#', '.', '#', '.'},
		{'#', '#', '.', '.', '#'},
	}))
	fmt.Println(maxStudents([][]byte{
		{'#', '.', '.', '.', '#'},
		{'.', '#', '.', '#', '.'},
		{'.', '.', '#', '.', '.'},
		{'.', '#', '.', '#', '.'},
		{'#', '.', '.', '.', '#'},
	}))

}

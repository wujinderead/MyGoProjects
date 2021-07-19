package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/nearest-exit-from-entrance-in-maze/

// You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and
// walls (represented as '+'). You are also given the entrance of the maze, where entrance
// = [entrancerow, entrancecol] denotes the row and column of the cell you are initially
// standing at.
// In one step, you can move one cell up, down, left, or right. You cannot step into a cell with
// a wall, and you cannot step outside the maze. Your goal is to find the nearest exit from the
// entrance. An exit is defined as an empty cell that is at the border of the maze. The entrance
// does not count as an exit.
// Return the number of steps in the shortest path from the entrance to the nearest exit, or -1
// if no such path exists.
// Example 1:
//   Input: maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
//   Output: 1
//   Explanation: There are 3 exits in this maze at [1,0], [0,2], and [2,3].
//     Initially, you are at the entrance cell [1,2].
//     - You can reach [1,0] by moving 2 steps left.
//     - You can reach [0,2] by moving 1 step up.
//     It is impossible to reach [2,3] from the entrance.
//     Thus, the nearest exit is [0,2], which is 1 step away.
// Example 2:
//   Input: maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]
//   Output: 2
//   Explanation: There is 1 exit in this maze at [1,2].
//     [1,0] does not count as an exit since it is the entrance cell.
//     Initially, you are at the entrance cell [1,0].
//     - You can reach [1,2] by moving 2 steps right.
//     Thus, the nearest exit is [1,2], which is 2 steps away.
// Example 3:
//   Input: maze = [[".","+"]], entrance = [0,0]
//   Output: -1
//   Explanation: There are no exits in this maze.
// Constraints:
//   maze.length == m
//   maze[i].length == n
//   1 <= m, n <= 100
//   maze[i][j] is either '.' or '+'.
//   entrance.length == 2
//   0 <= entrancerow < m
//   0 <= entrancecol < n
//   entrance will always be an empty cell.

// bfs, mark maze[i][j]='+' if visited
func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	q := list.New()
	step := 1
	q.PushBack([2]int{entrance[0], entrance[1]})
	maze[entrance[0]][entrance[1]] = '+'
	for q.Len() > 0 {
		ll := q.Len()
		for i := 0; i < ll; i++ {
			cur := q.Remove(q.Front()).([2]int)
			for _, v := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
				ni, nj := cur[0]+v[0], cur[1]+v[1]
				if ni >= 0 && ni < m && nj >= 0 && nj < n && maze[ni][nj] == '.' {
					if ni == 0 || ni == m-1 || nj == 0 || nj == n-1 {
						return step
					}
					maze[ni][nj] = '+' // mark as visited
					q.PushBack([2]int{ni, nj})
				}
			}
		}
		step++
	}
	return -1
}

func main() {
	for _, v := range []struct {
		maze     [][]byte
		entrance []int
		ans      int
	}{
		{[][]byte{
			{'+', '+', '.', '+'},
			{'.', '.', '.', '+'},
			{'+', '+', '+', '.'},
		}, []int{1, 2}, 1},
		{[][]byte{
			{'+', '+', '+'},
			{'.', '.', '.'},
			{'+', '+', '+'},
		}, []int{1, 0}, 2},
		{[][]byte{
			{'+', '+', '+'},
			{'.', '.', '+'},
			{'+', '+', '+'},
		}, []int{1, 0}, -1},
		{[][]byte{
			{'.', '.', '.'},
			{'.', '.', '.'},
			{'.', '.', '.'},
		}, []int{1, 0}, 1},
		{[][]byte{{'.', '+'}}, []int{0, 0}, -1},
		{[][]byte{{'.', '.', '+'}}, []int{0, 0}, 1},
	} {
		fmt.Println(nearestExit(v.maze, v.entrance), v.ans)
	}
}

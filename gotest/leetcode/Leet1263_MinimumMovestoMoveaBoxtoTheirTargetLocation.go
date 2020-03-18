package main

import (
	"container/list"
	"fmt"
)

// Storekeeper is a game in which the player pushes boxes around in a warehouse
// trying to get them to target locations.
// The game is represented by a grid of size m x n, where each element is a wall,
// floor, or a box. Your task is move the box 'B' to the target position 'T' under
// the following rules:
// 1. Player is represented by character 'S' and can move up, down, left, right in
//    the grid if it is a floor (empy cell).
// 2. Floor is represented by character '.' that means free cell to walk.
// 3. Wall is represented by character '#' that means obstacle (impossible to walk
//    there).
// 4. There is only one box 'B' and one target cell 'T' in the grid.
// 5. The box can be moved to an adjacent free cell by standing next to the box and
// 6. then moving in the direction of the box. This is a push.
// 7. The player cannot walk through the box.
// Return the minimum number of pushes to move the box to the target. If there is
// no way to reach the target, return -1.
// Example 1:
//   Input: grid = [["#","#","#","#","#","#"],
//               ["#","T","#","#","#","#"],
//               ["#",".",".","B",".","#"],
//               ["#",".","#","#",".","#"],
//               ["#",".",".",".","S","#"],
//               ["#","#","#","#","#","#"]]
//   Output: 3
//   Explanation: We return only the number of times the box is pushed.
//   ######      ######      ######       ######       ######      ######
//   #T####      #T####      #T####       #T####       #T####      #B####
//   #  B # move #  BS# push # BS # push  #BS  # move  #B   # push #S   #
//   # ## #  S   # ## # left # ## # left  # ## #  S    #S## #  up  # ## #
//   #   S#      #    #      #    #  ->   #    #  ->   #    #  ->  #    #
//   ######  ->  ######  ->  ######       ######       ######      ######
// Example 2:
//   Input: grid = [["#","#","#","#","#","#"],
//               ["#","T","#","#","#","#"],
//               ["#",".",".","B",".","#"],
//               ["#","#","#","#",".","#"],
//               ["#",".",".",".","S","#"],
//               ["#","#","#","#","#","#"]]
//   Output: -1
// Example 3:
//   Input: grid = [["#","#","#","#","#","#"],
//               ["#","T",".",".","#","#"],
//               ["#",".","#","B",".","#"],
//               ["#",".",".",".",".","#"],
//               ["#",".",".",".","S","#"],
//               ["#","#","#","#","#","#"]]
//   Output: 5
//   Explanation:  push the box down, left, left, up and up.
// Example 4:
//   Input: grid = [["#","#","#","#","#","#","#"],
//               ["#","S","#",".","B","T","#"],
//               ["#","#","#","#","#","#","#"]]
//   Output: -1
// Constraints:
//   m == grid.length
//   n == grid[i].length
//   1 <= m <= 20
//   1 <= n <= 20
//   grid contains only characters '.', '#', 'S' , 'T', or 'B'.
//   There is only one character 'S', 'B' and 'T' in the grid.

func minPushBox(grid [][]byte) int {
	// use bfs to calculate all possible 'person and box pair' states
	states := make(map[[4]int]int) // minimal step for 'S and B' pair
	m, n := len(grid), len(grid[0])
	BX, BY, SX, SY, TX, TY := 0, 0, 0, 0, 0, 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] == 'S' {
				SX, SY = i, j
				grid[i][j] = '.'
			}
			if grid[i][j] == 'B' {
				BX, BY = i, j
				grid[i][j] = '.'
			}
			if grid[i][j] == 'T' {
				TX, TY = i, j
				grid[i][j] = '.'
			}
		}
	}

	// initialize queue
	queue := list.New()
	queue.PushBack([4]int{BX, BY, SX, SY})
	states[[4]int{BX, BY, SX, SY}] = 0
	minimal := 0xffffffff
	for queue.Len() > 0 { // calculate all 'B and S' pairs; when B first reaches T, step may not be minimal
		tuple := queue.Remove(queue.Front()).([4]int)
		bx, by, sx, sy := tuple[0], tuple[1], tuple[2], tuple[3]
		step := states[tuple]
		if step >= minimal { // already has minimal
			continue
		}
		if bx == TX && by == TY && step < minimal { // B reach T, got a new minimal step
			minimal = step
			continue
		}
		grid[bx][by] = 'B'
		for _, d := range [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			dx, dy := d[0], d[1]
			nx, ny := sx+dx, sy+dy // new position of s
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			v, ok := states[[4]int{bx, by, nx, ny}]
			// move S, if new SB pair not exist, or SB pair's step decreases, update SB pair's step
			if grid[nx][ny] == '.' && (!ok || step < v) {
				states[[4]int{bx, by, nx, ny}] = step
				queue.PushBack([4]int{bx, by, nx, ny})
			}
			nbx, nby := bx+dx, by+dy
			if grid[nx][ny] == 'B' && nbx >= 0 && nbx < m && nby >= 0 && nby < n && grid[nbx][nby] == '.' {
				// move SB, if new SB pair not exist, or SB pair's step decreases, update SB pair's step
				if v, ok := states[[4]int{nbx, nby, nx, ny}]; !ok || step+1 < v {
					states[[4]int{nbx, nby, nx, ny}] = step + 1
				}
				queue.PushBack([4]int{nbx, nby, nx, ny})
			}
		}
		grid[bx][by] = '.'
	}
	if minimal != 0xffffffff {
		return minimal
	}
	return -1
}

func main() {
	grid := [][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', 'B', '.', '#'},
		{'#', '.', '#', '#', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'},
	}
	fmt.Println(minPushBox(grid))

	grid = [][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', 'B', '.', '#'},
		{'#', '#', '#', '#', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'},
	}
	fmt.Println(minPushBox(grid))

	grid = [][]byte{
		{'#', '#', '#', '#', '#', '#'},
		{'#', 'T', '.', '.', '#', '#'},
		{'#', '.', '#', 'B', '.', '#'},
		{'#', '.', '.', '.', '.', '#'},
		{'#', '.', '.', '.', 'S', '#'},
		{'#', '#', '#', '#', '#', '#'},
	}
	fmt.Println(minPushBox(grid))

	grid = [][]byte{
		{'#', '#', '#', '#', '#', '#', '#'},
		{'#', 'S', '#', '.', 'B', 'T', '#'},
		{'#', '#', '#', '#', '#', '#', '#'},
	}
	fmt.Println(minPushBox(grid))

	grid = [][]byte{
		{'.', '.', '#', '.', '.', '.', '.', '#'},
		{'.', 'B', '.', '.', '.', '.', '.', '#'},
		{'.', '.', 'S', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'T', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '#', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(minPushBox(grid))

	grid = [][]byte{
		{'#', '.', '.', '#', 'T', '#', '#', '#', '#'},
		{'#', '.', '.', '#', '.', '#', '.', '.', '#'},
		{'#', '.', '.', '#', '.', '#', 'B', '.', '#'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'#', '.', '.', '.', '.', '#', '.', 'S', '#'},
		{'#', '.', '.', '#', '.', '#', '#', '#', '#'},
	}
	fmt.Println(minPushBox(grid))

	grid = [][]byte{
		{'.', '.', '#', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '#', '.', '#', 'B', '#', '.', '#', '.', '.'},
		{'.', '#', '.', '.', '.', '.', '.', '.', 'T', '.'},
		{'#', '.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.', '#'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '#', '.'},
		{'.', '.', '.', '#', '.', '.', '#', '#', '.', '.'},
		{'.', '.', '.', '.', '#', '.', '.', '#', '.', '.'},
		{'.', '#', '.', 'S', '.', '.', '.', '.', '.', '.'},
		{'#', '.', '.', '#', '.', '.', '.', '.', '.', '#'},
	}
	fmt.Println(minPushBox(grid))
}

// # . . # T # # # #
// # . . # . # . . #
// # . . # . # . . #
// # . . s b . . . #
// # . . . . # . . #
// # . . # . # # # #

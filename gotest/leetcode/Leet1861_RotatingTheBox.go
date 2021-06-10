package main

import "fmt"

// https://leetcode.com/problems/rotating-the-box/

// You are given an m x n matrix of characters box representing a side-view of a
// box. Each cell of the box is one of the following:
//   A stone '#'
//   A stationary obstacle '*'
//   Empty '.'
// The box is rotated 90 degrees clockwise, causing some of the stones to fall due to gravity.
// Each stone falls down until it lands on an obstacle, another stone, or the bottom of the box.
// Gravity does not affect the obstacles' positions, and the inertia from the box's rotation
// does not affect the stones' horizontal positions.
// It is guaranteed that each stone in box rests on an obstacle, another stone, or the bottom of the box.
// Return an n x m matrix representing the box after the rotation described above.
// Example 1:
//   Input: box = [["#",".","#"]]
//   Output: [["."],
//            ["#"],
//            ["#"]]
// Example 2:
//   Input: box = [["#",".","*","."],
//                 ["#","#","*","."]]
//   Output: [["#","."],
//            ["#","#"],
//            ["*","*"],
//            [".","."]]
// Example 3:
//   Input: box = [["#","#","*",".","*","."],
//                 ["#","#","#","*",".","."],
//                 ["#","#","#",".","#","."]]
//   Output: [[".","#","#"],
//            [".","#","#"],
//            ["#","#","*"],
//            ["#","*","."],
//            ["#",".","*"],
//            ["#",".","."]]
// Constraints:
//   m == box.length
//   n == box[i].length
//   1 <= m, n <= 500
//   box[i][j] is either '#', '*', or '.'.
//   Related Topics Array Two Pointers

// use two pointers to rearrange each line from right to left
func rotateTheBox(box [][]byte) [][]byte {
	ans := make([][]byte, len(box[0]))
	for i := range ans {
		ans[i] = make([]byte, len(box))
	}

	for _, line := range box {
		// drop stone for each line
		i, j := len(box[0])-1, len(box[0])-1
		for {
			for i >= 0 && line[i] != '.' { // find first line[i]='.' from right to left
				i--
			}
			if i <= 0 {
				break
			}
			if j >= i { // j must be left of i
				j = i - 1
			}
			for j >= 0 && line[j] == '.' { // from left of line[i]='.', find first '*' or '#'
				j--
			}
			if j < 0 {
				break
			}
			if line[j] == '#' { // find '#', e.g. '#...', it become '...#'
				line[i] = '#'
				line[j] = '.'
				i-- // the next '.' is at line[i-1]
			} else { // find '*', e.g. '*...', set i=j-1 to find '.' in left of '*'
				i = j - 1
			}
		}
	}
	// rotate
	for i := range box {
		for j := range box[0] {
			ans[j][len(box)-i-1] = box[i][j]
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		box, ans [][]byte
	}{
		{[][]byte{{'#', '.', '#'}}, [][]byte{{'.'}, {'#'}, {'#'}}},
		{[][]byte{{'#', '.', '*', '.'}, {'#', '#', '*', '.'}}, [][]byte{{'#', '.'}, {'#', '#'}, {'*', '*'}, {'.', '.'}}},
		{[][]byte{{'#', '#', '*', '.', '*', '.'}, {'#', '#', '#', '*', '.', '.'}, {'#', '#', '#', '.', '#', '.'}},
			[][]byte{{'.', '#', '#'}, {'.', '#', '#'}, {'#', '#', '*'}, {'#', '*', '.'}, {'#', '.', '*'}, {'#', '.', '.'}}},
	} {
		fmt.Println(rotateTheBox(v.box))
		fmt.Println(v.ans)
	}
}

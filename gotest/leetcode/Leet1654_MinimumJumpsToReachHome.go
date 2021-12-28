package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/minimum-jumps-to-reach-home/

// A certain bug's home is on the x-axis at position x. Help them get there from position 0.
// The bug jumps according to the following rules:
//   It can jump exactly a positions forward (to the right).
//   It can jump exactly b positions backward (to the left).
//   It cannot jump backward twice in a row.
//   It cannot jump to any forbidden positions.
// The bug may jump forward beyond its home, but it cannot jump to positions numbered with negative integers.
// Given an array of integers forbidden, where forbidden[i] means that the bug cannot jump to the position
// forbidden[i], and integers a, b, and x, return the minimum number of jumps needed for the bug to reach its home.
// If there is no possible sequence of jumps that lands the bug on position x, return -1.
// Example 1:
//   Input: forbidden = [14,4,18,1,15], a = 3, b = 15, x = 9
//   Output: 3
//   Explanation: 3 jumps forward (0 -> 3 -> 6 -> 9) will get the bug home.
// Example 2:
//   Input: forbidden = [8,3,16,6,12,20], a = 15, b = 13, x = 11
//   Output: -1
// Example 3:
//   Input: forbidden = [1,6,2,14,5,17,4], a = 16, b = 9, x = 7
//   Output: 2
//   Explanation: One jump forward (0 -> 16) then one jump backward (16 -> 7) will get the bug home.
// Constraints:
//   1 <= forbidden.length <= 1000
//   1 <= a, b, forbidden[i] <= 2000
//   0 <= x <= 2000
//   All the elements in forbidden are distinct.
//   Position x is not forbidden.

// easy problem to use bfs. what really matters is the upper bound of the search space,
// it should be max(x, max(forbidden))+a+b, according to
// https://leetcode.com/problems/minimum-jumps-to-reach-home/discuss/935419/Python-deque-BFS-O(max(x-max(forbidden))%2Ba%2Bb)
func minimumJumps(forbidden []int, a int, b int, x int) int {
	if x == 0 {
		return 0
	}
	// pos[x][0] step to pos x with right direction, pos[x][1] steps pos x with left direction
	pos := make([][2]int, 2000+a+b)
	fmap := make(map[int]struct{}, len(forbidden)) // forbidden set
	for _, v := range forbidden {
		fmap[v] = struct{}{}
	}
	queue := list.New()
	queue.PushBack([2]int{0, 0}) // pos: 0, direction: 0 (0 to right, 1 to left),
	step := 1
	for queue.Len() > 0 {
		leng := queue.Len()
		for i := 0; i < leng; i++ {
			tmp := queue.Remove(queue.Front()).([2]int)
			pre, dir := tmp[0], tmp[1]
			if pre+a < 2000+a+b {
				if _, ok := fmap[pre+a]; !ok && pos[pre+a][0] == 0 {
					if pre+a == x {
						return step
					}
					pos[pre+a][0] = step
					queue.PushBack([2]int{pre + a, 0})
				}
			}
			if pre-b > 0 && dir == 0 { // prev move must be forward
				if _, ok := fmap[pre-b]; !ok && pos[pre-b][1] == 0 {
					if pre-b == x {
						return step
					}
					pos[pre-b][1] = step
					queue.PushBack([2]int{pre - b, 1})
				}
			}
		}
		step++
	}
	return -1
}

func main() {
	for _, v := range []struct {
		f            []int
		a, b, x, ans int
	}{
		{[]int{14, 4, 18, 1, 15}, 3, 15, 9, 3},
		{[]int{8, 3, 16, 6, 12, 20}, 15, 13, 11, -1},
		{[]int{1, 6, 2, 14, 5, 17, 4}, 16, 9, 7, 2},
		{[]int{1998}, 1999, 2000, 2000, 3998},
	} {
		fmt.Println(minimumJumps(v.f, v.a, v.b, v.x), v.ans)
	}
}

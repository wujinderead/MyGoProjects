package main

import "fmt"

// https://leetcode.com/problems/building-boxes/

// You have a cubic storeroom where the width, length, and height of the room are
// all equal to n units. You are asked to place n boxes in this room where each box
// is a cube of unit side length. There are however some rules to placing the boxes:
// - You can place the boxes anywhere on the floor.
// - If box x is placed on top of the box y, then each side of the four vertical
//   sides of the box y must either be adjacent to another box or to a wall.
// Given an integer n, return the minimum possible number of boxes touching the floor.
// Example 1:
//   Input: n = 3
//   Output: 3
//   Explanation: The figure above is for the placement of the three boxes.
//     These boxes are placed in the corner of the room, where the corner is on the left side.
// Example 2:
//   Input: n = 4
//   Output: 3
//   Explanation: The figure above is for the placement of the four boxes.
//     These boxes are placed in the corner of the room, where the corner is on the left side.
// Example 3:
//   Input: n = 10
//   Output: 6
//   Explanation: The figure above is for the placement of the ten boxes.
//     These boxes are placed in the corner of the room, where the corner is on the back side.
// Constraints:
//   1 <= n <= 10^9

// let the ans be f(n)=x, we know
// f(1)=1
// f(4)=3
// f(10)=6
// f(20)=10
// f(1+(1+2)+...(1+...5)) = f(35) = 15 = 1+...+5
// f(1+(1+2)+(1+...+n)) = 1+...n
// we know that for n=10, we need 6 boxes touch floor
// we know that for n=20, we need 10 boxes touch floor
// what about 11 <= n <= 19 ?
// when floor=6, we can contain 10 boxes with diagram:
//  3 2 1
//  2 1
//  1
// add 1 box, we can contain 1 more box (7 floor can contain 11 boxes):
//  3 2 1
//  2 1
//  1
//  1
// add 1 more box, we can contain 2 more box (8 floor can contain 12-13 boxes):
//  3 2 1
//  2 1
//  2 1
//  1
// add 1 more box, we can contain 3 more box (9 floor can contain 14-16 boxes):
//  3 2 1
//  3 2 1
//  2 1
//  1
// add 1 more box, we can contain 4 more box (10 floor can contain 17-20 boxes):
//  4 3 2 1
//  3 2 1
//  2 1
//  1
func minimumBoxes(n int) int {
	var all, cur int
	i := 1
	for all < n {
		cur += i
		all += cur
		i++
	}
	if all == n {
		return cur
	}
	ans := cur - i + 1
	all = all - cur
	// the variables mean: 'ans' floor boxes can contain 'all' boxes

	i = 1
	for {
		ans++    // add 1 more floor box
		all += i // can contain i more boxes
		if all >= n {
			break
		}
		i++
	}
	return ans
}

func main() {
	for _, v := range []struct {
		n, ans int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 3},
		{5, 4},
		{6, 5},
		{7, 5},
		{8, 6},
		{9, 6},
		{10, 6},
		{11, 7},
		{12, 8},
		{13, 8},
		{14, 9},
		{15, 9},
		{16, 9},
		{17, 10},
		{18, 10},
		{19, 10},
		{20, 10},
		{1000000000, 1650467},
		{56, 21},
		{57, 22},
		{58, 23},
		{59, 23},
		{60, 24},
	} {
		//minimumBoxes(v.n)
		fmt.Println(minimumBoxes(v.n), v.ans)
	}
}

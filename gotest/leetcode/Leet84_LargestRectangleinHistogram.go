package main

import "fmt"

// https://leetcode.com/problems/largest-rectangle-in-histogram/

// Given n non-negative integers representing the histogram's bar height
// where the width of each bar is 1, find the area of largest rectangle in the histogram.
// Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
// The largest rectangle is shown in the shaded area, which has area = 10 unit.
// Example:
//   Input: [2,1,5,6,2,3]
//   Output: 10
//     2 1 5 6 2 3
//           ▇
//         ▇ ▇
//         ▇ ▇
//         ▇ ▇   ▇
//     ▇   ▇ ▇ ▇ ▇
//     ▇ ▇ ▇ ▇ ▇ ▇
//     0 1 2 3 4 5

func largestRectangleArea(heights []int) int {
	stack := make([]int, len(heights)+1)
	slen := 0
	allmax := 0
	for j := 0; j <= len(heights); j++ {
		h := 0 // add 0 as the last element to trigger final compute
		if j < len(heights) {
			h = heights[j]
		}
		fmt.Println(j, h)
		for slen > 0 && h < heights[stack[slen-1]] {
			p := stack[slen-1] // pop the index
			slen--
			wid := j
			if slen > 0 {
				wid = j - 1 - stack[slen-1] // wid = (slen==0 ? j : j-1-stack.peek)
			}
			cur := heights[p] * wid
			fmt.Println("===", cur, heights[p], wid)
			if cur > allmax {
				allmax = cur
			}
		}
		stack[slen] = j // what is pushed to the stack is the index
		slen++
	}
	return allmax
}

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println()

	//     2 1 5 5 6 2 2 3
	//             ▇
	//         ▇ ▇ ▇
	//         ▇ ▇ ▇
	//         ▇ ▇ ▇     ▇
	//     ▇   ▇ ▇ ▇ ▇ ▇ ▇
	//     ▇ ▇ ▇ ▇ ▇ ▇ ▇ ▇
	//     0 1 2 3 4 5 6 7
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 5, 6, 2, 2, 3}))
	fmt.Println()

	//     2 1 5 0 5 6 2 2 3
	//               ▇
	//         ▇   ▇ ▇
	//         ▇   ▇ ▇
	//         ▇   ▇ ▇     ▇
	//     ▇   ▇   ▇ ▇ ▇ ▇ ▇
	//     ▇ ▇ ▇   ▇ ▇ ▇ ▇ ▇
	//     0 1 2 3 4 5 6 7 8
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 0, 5, 6, 2, 2, 3}))
	fmt.Println(largestRectangleArea([]int{0, 0, 0}))
}

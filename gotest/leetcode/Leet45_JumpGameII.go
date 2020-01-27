package main

import "fmt"

// https://leetcode.com/problems/jump-game-ii/

// Given an array of non-negative integers, you are initially positioned at the first index of the array.
// Each element in the array represents your maximum jump length at that position.
// Your goal is to reach the last index in the minimum number of jumps.
// Example:
//   Input: [2,3,1,1,4]
//   Output: 2
//   Explanation:
//     The minimum number of jumps to reach the last index is 2.
//     Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Note:
//   You can assume that you can always reach the last index.

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	// use bfs to find the minimal steps
	queue := make([][2]int, len(nums)) // a circular queue
	head := 0
	end := 0
	size := 1
	queue[0] = [2]int{0, 0} // index 0, tier 0
	visited := make([]bool, len(nums))
	visited[0] = true
	for size > 0 {
		curind := queue[head][0]
		curtier := queue[head][1]
		head = (head + 1) % len(queue)
		size--
		for i := 1; i <= nums[curind]; i++ {
			if curind+i < len(nums) && !visited[curind+i] {
				visited[curind+i] = true
				if curind+i == len(nums)-1 {
					return curtier + 1
				}
				end = (end + 1) % len(queue)
				queue[end][0] = curind + i
				queue[end][1] = curtier + 1
				size++
			}
		}
	}
	// backup answer: by one step each time, you can always reach last index
	return len(nums) - 1
}

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

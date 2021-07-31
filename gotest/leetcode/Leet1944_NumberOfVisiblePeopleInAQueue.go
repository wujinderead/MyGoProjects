package main

import "fmt"

// https://leetcode.com/problems/number-of-visible-people-in-a-queue/

// There are n people standing in a queue, and they numbered from 0 to n - 1 in left to
// right order. You are given an array heights of distinct integers where heights[i]
// represents the height of the ith person.
// A person can see another person to their right in the queue if everybody in between is
// shorter than both of them. More formally, the ith person can see the jth person if i < j
// and min(heights[i], heights[j]) > max(heights[i+1], heights[i+2], ..., heights[j-1]).
// Return an array answer of length n where answer[i] is the number of people the ith person
// can see to their right in the queue.
// Example 1:
//   Input: heights = [10,6,8,5,11,9]
//   Output: [3,1,2,1,1,0]
//   Explanation:
//     Person 0 can see person 1, 2, and 4.
//     Person 1 can see person 2.
//     Person 2 can see person 3 and 4.
//     Person 3 can see person 4.
//     Person 4 can see person 5.
//     Person 5 can see no one since nobody is to the right of them.
// Example 2:
//   Input: heights = [5,1,2,3,10]
//   Output: [4,1,1,1,0]
// Constraints:
//   n == heights.length
//   1 <= n <= 10^5
//   1 <= heights[i] <= 10^5
//   All the values of heights are unique.

func canSeePersonsCount(heights []int) []int {
	ans := make([]int, len(heights))
	stack := make([]int, len(heights))
	ind := -1
	for i := len(heights) - 1; i >= 0; i-- {
		h := heights[i]
		// e.g., current 4, stack 1,2,3,4,5,6
		// pop 1,2,3,4, as current can see them
		// stack remains [5,6], current can also see 5
		// finally push 4 to stack as [4,5,6]
		for ind >= 0 && stack[ind] <= h { // pop values <= current height
			ind--
			ans[i]++ // increment ans as we can see these person with height shorter than self
		}
		if ind >= 0 { // we can also see first the higher person than current height
			ans[i]++
		}
		ind++ // push to stack
		stack[ind] = h
	}
	return ans
}

func main() {
	for _, v := range []struct {
		h, ans []int
	}{
		{[]int{10, 6, 8, 5, 11, 9}, []int{3, 1, 2, 1, 1, 0}},
		{[]int{5, 1, 2, 3, 10}, []int{4, 1, 1, 1, 0}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 1, 1, 1, 0}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 1, 1, 1, 0}},
		{[]int{1, 2}, []int{1, 0}},
		{[]int{2, 1}, []int{1, 0}},
	} {
		fmt.Println(canSeePersonsCount(v.h), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/find-the-longest-valid-obstacle-course-at-each-position/

// You want to build some obstacle courses. You are given a 0-indexed integer array obstacles
// of length n, where obstacles[i] describes the height of the ith obstacle.
// For every index i between 0 and n - 1 (inclusive), find the length of the longest obstacle
// course in obstacles such that:
//   You choose any number of obstacles between 0 and i inclusive.
//   You must include the ith obstacle in the course.
//   You must put the chosen obstacles in the same order as they appear in obstacles.
//   Every obstacle (except the first) is taller than or the same height as the obstacle immediately before it.
// Return an array ans of length n, where ans[i] is the length of the longest obstacle course
// for index i as described above.
// Example 1:
//   Input: obstacles = [1,2,3,2]
//   Output: [1,2,3,3]
//   Explanation: The longest valid obstacle course at each position is:
//     - i = 0: [1], [1] has length 1.
//     - i = 1: [1,2], [1,2] has length 2.
//     - i = 2: [1,2,3], [1,2,3] has length 3.
//     - i = 3: [1,2,3,2], [1,2,2] has length 3.
// Example 2:
//   Input: obstacles = [2,2,1]
//   Output: [1,2,1]
//   Explanation: The longest valid obstacle course at each position is:
//     - i = 0: [2], [2] has length 1.
//     - i = 1: [2,2], [2,2] has length 2.
//     - i = 2: [2,2,1], [1] has length 1.
// Example 3:
//   Input: obstacles = [3,1,5,6,4,2]
//   Output: [1,1,2,3,2,2]
//   Explanation: The longest valid obstacle course at each position is:
//     - i = 0: [3], [3] has length 1.
//     - i = 1: [3,1], [1] has length 1.
//     - i = 2: [3,1,5], [3,5] has length 2. [1,5] is also valid.
//     - i = 3: [3,1,5,6], [3,5,6] has length 3. [1,5,6] is also valid.
//     - i = 4: [3,1,5,6,4], [3,4] has length 2. [1,4] is also valid.
//     - i = 5: [3,1,5,6,4,2], [1,2] has length 2.
// Constraints:
//   n == obstacles.length
//   1 <= n <= 10^5
//   1 <= obstacles[i] <= 10^7

// the task is to find the length of the longest non-decreasing subsequence end at each index.
// for each length, we record the lowest end value
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	// lis[i] is the lowest end value of the non-decreasing subsequence with length i
	lis := []int{0}
	ans := make([]int, len(obstacles))
	for i, v := range obstacles {
		if v >= lis[len(lis)-1] {
			ans[i] = len(lis)
			lis = append(lis, v)
			continue
		}
		// find the smallest value in lis > v
		l, r := 0, len(lis)-1
		for l < r {
			mid := (l + r) / 2
			if v >= lis[mid] {
				l = mid + 1
			} else {
				r = mid
			}
			// when stop, arr[l=r] < v; arr[l-1]>=v
		}
		ans[i] = r
		lis[r] = v
	}
	return ans
}

func main() {
	for _, v := range []struct {
		ob, ans []int
	}{
		{[]int{1, 2, 3, 2}, []int{1, 2, 3, 3}},
		{[]int{2, 2, 1}, []int{1, 2, 1}},
		{[]int{3, 1, 5, 6, 4, 2}, []int{1, 1, 2, 3, 2, 2}},
		{[]int{1, 3, 5, 7, 9, 6}, []int{1, 2, 3, 4, 5, 4}},
		{[]int{2, 3, 5, 7, 9, 1}, []int{1, 2, 3, 4, 5, 1}},
	} {
		fmt.Println(longestObstacleCourseAtEachPosition(v.ob), v.ans)
	}
}

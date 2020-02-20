package main

import (
	"fmt"
)

// https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/

// You want to schedule a list of jobs in d days. Jobs are dependent (i.e To work on the i-th job,
// you have to finish all the jobs j where 0 <= j < i). You have to finish at least one task every day.
// The difficulty of a job schedule is the sum of difficulties of each day of the d days. The difficulty
// of a day is the maximum difficulty of a job done in that day. Given an array of integers jobDifficulty
// and an integer d. The difficulty of the i-th job is jobDifficulty[i].
// Return the minimum difficulty of a job schedule. If you cannot find a schedule for the jobs return -1.
// Example 1:
//   Input: jobDifficulty = [6,5,4,3,2,1], d = 2
//   Output: 7
//   Explanation:
//     First day you can finish the first 5 jobs, total difficulty = 6.
//     Second day you can finish the last job, total difficulty = 1.
//     The difficulty of the schedule = 6 + 1 = 7
// Example 2:
//   Input: jobDifficulty = [9,9,9], d = 4
//   Output: -1
//   Explanation:
//     If you finish a job per day you will still have a free day.
//     you cannot find a schedule for the given jobs.
// Example 3:
//   Input: jobDifficulty = [1,1,1], d = 3
//   Output: 3
//   Explanation: The schedule is one job per day. total difficulty will be 3.
// Example 4:
//   Input: jobDifficulty = [7,1,7,1,7,1], d = 3
//   Output: 15
// Example 5:
//   Input: jobDifficulty = [11,111,22,222,33,333,44,444], d = 6
//   Output: 843
// Constraints:
//   1 <= jobDifficulty.length <= 300
//   0 <= jobDifficulty[i] <= 1000
//   1 <= d <= 10

// we need to partition the jobs into consecutive parts, and sum the maximal value in each part
// our task is to find the minimal sum
// let dif(i, j) be the result for partition arr[0...i] to j parts, then
// dif(i, j) has these candidate values, and we want the min value:
// arr[i]+dif(i-1, j-1), max(arr[i-1...i])+dif(i-2, j-1), max(arr[i-2...i])+dif(i-3, j-1)
// thus we use a bottom-up iteration to compute dis(0, 1), dis(1, 1)... firstly
// time O(n*len(arr)²), space O(len(arr)) since update line by line
//
// for example:
//
//     arr= 5, 4, 2, 6, 1, 3
// n=1      5  5  5  6  6  6   d([5],1)  d([5,4],1)  d([5,4,3],1)...
// n=2      x  9  7 11  7  9   d([5,4],2)=d([5],1)+4, d([5,4,2],2)=min(d([5,4],1)+2, d([5],1)+max(4,2))
// n=3      x  x 11 13         d([5,4,2],3)=d([5,4],2)+2
//                             d([5,4,2,6],3)=min(d([5,4,2],2)+6, d([5,4],2)+max(2,6))
//                             d([5,4,2,6,1],3)=min(d([5,4,2,6],2)+1, d([5,4,2],2)+max(1,6), d([5,4],2)+max(2,6,1))
func minDifficulty(arr []int, n int) int {
	if len(arr) < n {
		return -1
	}
	prev := make([]int, len(arr))
	cur := make([]int, len(arr))
	// setup for n=1
	prev[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prev[i] = max(prev[i-1], arr[i])
	}
	// dp
	for d := 2; d <= n; d++ { // d, number to partition
		for i := d - 1; i < len(arr); i++ { // to calculate dif([0...i], d)
			rightmax := arr[i]
			minval := prev[i-1] + rightmax  // first candidate dif([0...i-1], d-1)+arr[i]
			for j := i - 1; j >= d-1; j-- { // dif([0...j-1], d-1)+max(arr[j], rightmax)
				rightmax = max(arr[j], rightmax)
				minval = min(minval, prev[j-1]+rightmax)
			}
			cur[i] = minval
		}
		fmt.Println(cur)
		prev, cur = cur, prev
	}
	return prev[len(arr)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minDifficulty([]int{5, 4, 2, 6, 1, 3}, 4))
	fmt.Println(minDifficulty([]int{9, 9, 9}, 4))
	fmt.Println(minDifficulty([]int{1, 1, 1}, 3))
	fmt.Println(minDifficulty([]int{7, 1, 7, 1, 7, 1}, 3))
	fmt.Println(minDifficulty([]int{11, 111, 22, 222, 33, 333, 44, 444}, 6))
}

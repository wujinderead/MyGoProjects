package main

import "fmt"

// https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/

// There is an integer array nums that consists of n unique elements, but you have
// forgotten it. However, you do remember every pair of adjacent elements in nums.
// You are given a 2D integer array adjacentPairs of size n - 1 where each
// adjacentPairs[i] = [ui, vi] indicates that the elements ui and vi are adjacent in nums.
// It is guaranteed that every adjacent pair of elements nums[i] and nums[i+1] will
// exist in adjacentPairs, either as [nums[i], nums[i+1]] or [nums[i+1], nums[i]].
// The pairs can appear in any order.
// Return the original array nums. If there are multiple solutions, return any of them.
// Example 1:
//   Input: adjacentPairs = [[2,1],[3,4],[3,2]]
//   Output: [1,2,3,4]
//   Explanation: This array has all its adjacent pairs in adjacentPairs.
//     Notice that adjacentPairs[i] may not be in left-to-right order.
// Example 2:
//   Input: adjacentPairs = [[4,-2],[1,4],[-3,1]]
//   Output: [-2,4,1,-3]
//   Explanation: There can be negative numbers.
//     Another solution is [-3,1,4,-2], which would also be accepted.
// Example 3:
//   Input: adjacentPairs = [[100000,-100000]]
//   Output: [100000,-100000]
// Constraints:
//   nums.length == n
//   adjacentPairs.length == n - 1
//   adjacentPairs[i].length == 2
//   2 <= n <= 10^5
//   -10^5 <= nums[i], ui, vi <= 10^5
//   There exists some nums that has adjacentPairs as its pairs.

func restoreArray(adjacentPairs [][]int) []int {
	const NAN = -10000000
	mapp := make(map[int][2]int)

	// store pairs to map
	for _, nums := range adjacentPairs {
		a, b := nums[0], nums[1]
		if v, ok := mapp[a]; ok {
			mapp[a] = [2]int{v[0], b}
		} else {
			mapp[a] = [2]int{b, NAN}
		}
		if v, ok := mapp[b]; ok {
			mapp[b] = [2]int{v[0], a}
		} else {
			mapp[b] = [2]int{a, NAN}
		}
	}

	var start int
	for k, v := range mapp {
		if v[1] == NAN {
			start = k
			break
		}
	}

	// restore the array
	visited := make(map[int]struct{})
	ans := make([]int, 0, len(adjacentPairs)+1)
	for i := 0; i < len(adjacentPairs)+1; i++ {
		ans = append(ans, start)
		visited[start] = struct{}{}
		v := mapp[start]
		// to next element
		if _, ok := visited[v[0]]; !ok && v[0] != NAN {
			start = v[0]
		} else if _, ok := visited[v[1]]; !ok && v[1] != NAN {
			start = v[1]
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		ap  [][]int
		ans []int
	}{
		{[][]int{{2, 1}, {3, 4}, {3, 2}}, []int{1, 2, 3, 4}},
		{[][]int{{4, -2}, {1, 4}, {-3, 1}}, []int{-2, 4, 1, -3}},
		{[][]int{{100000, -100000}}, []int{100000, -100000}},
		{[][]int{{4, -10}, {-1, 3}, {4, -3}, {-3, 3}}, []int{-10, 4, -3, 3, -1}},
	} {
		fmt.Println(restoreArray(v.ap), v.ans)
	}
}

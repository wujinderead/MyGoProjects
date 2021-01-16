package main

import "fmt"

// https://leetcode.com/problems/minimize-hamming-distance-after-swap-operations/

// You are given two integer arrays, source and target, both of length n. You are also given
// an array allowedSwaps where each allowedSwaps[i] = [ai, bi] indicates that you are allowed
// to swap the elements at index ai and index bi (0-indexed) of array source. Note that you
// can swap elements at a specific pair of indices multiple times and in any order.
// The Hamming distance of two arrays of the same length, source and target, is the number of
// positions where the elements are different. Formally, it is the number of indices i for
// 0 <= i <= n-1 where source[i] != target[i] (0-indexed).
// Return the minimum Hamming distance of source and target after performing any amount of
// swap operations on array source.
// Example 1:
//   Input: source = [1,2,3,4], target = [2,1,4,5], allowedSwaps = [[0,1],[2,3]]
//   Output: 1
//   Explanation: source can be transformed the following way:
//     - Swap indices 0 and 1: source = [2,1,3,4]
//     - Swap indices 2 and 3: source = [2,1,4,3]
//     The Hamming distance of source and target is 1 as they differ in 1 position: index 3.
// Example 2:
//   Input: source = [1,2,3,4], target = [1,3,2,4], allowedSwaps = []
//   Output: 2
//   Explanation: There are no allowed swaps.
//     The Hamming distance of source and target is 2 as they differ in 2 positions:
//     index 1 and index 2.
// Example 3:
//   Input: source = [5,1,2,4,3], target = [1,5,4,2,3], allowedSwaps = [[0,4],[4,2],[1,3],[1,4]]
//   Output: 0
// Constraints:
//   n == source.length == target.length
//   1 <= n <= 10^5
//   1 <= source[i], target[i] <= 10^5
//   0 <= allowedSwaps.length <= 10^5
//   allowedSwaps[i].length == 2
//   0 <= ai, bi <= n - 1
//   ai != bi

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	// union-find to union the allowedSwaps
	roots := make([]int, len(source))
	visited := make([]bool, len(source))
	for i := range roots {
		roots[i] = -1
	}
	for _, v := range allowedSwaps {
		ra := root(roots, v[0])
		rb := root(roots, v[1])
		visited[v[0]] = true
		visited[v[1]] = true
		if ra != rb {
			roots[ra] = rb
		}
	}
	group := make(map[int][]int)
	for i := range source {
		if visited[i] {
			r := root(roots, i)
			group[r] = append(group[r], i)
		}
	}
	// for each swappable group, count occurrence
	same := 0
	count := make(map[int]int)
	for _, v := range group {
		for _, vv := range v {
			count[target[vv]] = count[target[vv]] + 1
		}
		for _, vv := range v {
			if count[source[vv]] > 0 {
				count[source[vv]] = count[source[vv]] - 1
				same++
			}
		}
		for _, vv := range v {
			count[target[vv]] = 0
		}
	}
	// for unswappable position
	for i := range visited {
		if !visited[i] {
			if source[i] == target[i] {
				same++
			}
		}
	}
	return len(source) - same
}

func root(arr []int, i int) int {
	for arr[i] != -1 {
		x := root(arr, arr[i])
		arr[i] = x
		return x
	}
	return i
}

func main() {
	for _, v := range []struct {
		s, t []int
		a    [][]int
		ans  int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{{0, 1}, {2, 3}}, 1},
		{[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}, [][]int{}, 2},
		{[]int{5, 1, 2, 4, 3}, []int{1, 5, 4, 2, 3}, [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}}, 0},
	} {
		fmt.Println(minimumHammingDistance(v.s, v.t, v.a), v.ans)
	}
}

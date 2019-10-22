package leetcode

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/ones-and-zeroes/

// In the computer world, use restricted resource you have to
// generate maximum benefit is what we always want to pursue.
// For now, suppose you are a dominator of m 0s and n 1s respectively.
// On the other hand, there is an array with strings consisting of only 0s and 1s.
// Now your task is to find the maximum number of strings that you can form
// with given m 0s and n 1s. Each 0 and 1 can be used at most once.
// Note:
//   The given numbers of 0s and 1s will both not exceed 100
//   The size of given string array won't exceed 600.
// Example 1:
//   Input: Array = {"10", "0001", "111001", "1", "0"}, m = 5, n = 3
//   Output: 4
//   Explanation: This are totally 4 strings can be formed by the using of 5 0s and 3 1s,
//     which are “10,”0001”,”1”,”0”
// Example 2:
//   Input: Array = {"10", "0", "1"}, m = 1, n = 1
//   Output: 2
//   Explanation: You could form "10", but then you'd have nothing left. Better form "0" and "1".

// the problem can be transferred as:
// we have some resources of 2 types, and there are some items cost these 2 types of resources.
// we need to find the maximal number of items we can get.
func findMaxForm(strs []string, m int, n int) int {
	// let a(s, i, j) be the max number that can be formed by i 0s and j 1s for strs[0...i]
	// then a(s, i, j) = max( a(s-1, i, j), a(s-1, i-s0, j-s1)+1 )
	// note: it's a(s-1, i-s0, j-s1) since we can only use 's' for once.
	//       if we can use 's' multiple times, it should be a(s, i-s0, j-s1)
	// base case a(s, 0, 0) = 0
	// time O(len(strs)*m*n), space O(m*n)
	a0 := make([]int, (m+1)*(n+1))
	a1 := make([]int, (m+1)*(n+1))
	col := n + 1
	for s := range strs {
		s0 := strings.Count(strs[s], "0")
		s1 := strings.Count(strs[s], "1")
		for i := 0; i < m+1; i++ {
			for j := 0; j < n+1; j++ {
				tmp := get2d(a0, col, i, j)
				set2d(a1, col, i, j, tmp) // a(s, i, j) = a(s-1, i, j)
				if s0 <= i && s1 <= j {
					// a(s-1, i-s0, j-s1)+1 )
					set2d(a1, col, i, j, max(tmp, get2d(a0, col, i-s0, j-s1)+1))
				}
			}
		}
		print2d(a1, m+1, n+1)
		a0, a1 = a1, a0
	}
	return get2d(a0, col, m, n)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func get2d(arr []int, col, i, j int) int {
	return arr[i*col+j]
}

func set2d(arr []int, col, i, j, val int) {
	arr[i*col+j] = val
}

func print2d(arr []int, row, col int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print(arr[i*col+j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(findMaxForm([]string{"10", "1", "0"}, 1, 1))
	fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 3, 4))
}

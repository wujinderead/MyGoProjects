package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/find-array-given-subset-sums/

// You are given an integer n representing the length of an unknown array that you are trying to
// recover. You are also given an array sums containing the values of all 2^n subset sums of the
// unknown array (in no particular order).
// Return the array ans of length n representing the unknown array. If multiple answers exist,
// return any of them.
// An array sub is a subset of an array arr if sub can be obtained from arr by deleting some (possibly
// zero or all) elements of arr. The sum of the elements in sub is one possible subset sum of arr.
// The sum of an empty array is considered to be 0.
// Note: Test cases are generated such that there will always be at least one correct answer.
// Example 1:
//   Input: n = 3, sums = [-3,-2,-1,0,0,1,2,3]
//   Output: [1,2,-3]
//   Explanation: [1,2,-3] is able to achieve the given subset sums:
//     - []: sum is 0
//     - [1]: sum is 1
//     - [2]: sum is 2
//     - [1,2]: sum is 3
//     - [-3]: sum is -3
//     - [1,-3]: sum is -2
//     - [2,-3]: sum is -1
//     - [1,2,-3]: sum is 0
//     Note that any permutation of [1,2,-3] and also any permutation of [-1,-2,3] will also be accepted.
// Example 2:
//   Input: n = 2, sums = [0,0,0,0]
//   Output: [0,0]
//   Explanation: The only correct answer is [0,0].
// Example 3:
//   Input: n = 4, sums = [0,0,5,5,4,-1,4,9,9,-1,4,3,4,8,3,8]
//   Output: [0,-1,4,5]
//   Explanation: [0,-1,4,5] is able to achieve the given subset sums.
// Constraints:
//   1 <= n <= 15
//   sums.length == 2^n
//   -10^4 <= sums[i] <= 10^4

// assume the original array in non-descending order is n1,n2,...nx,0...0,p1,p2,...,px,
// then for the subset sums, min=n1+n2+...nx must be the smallest.
// for second smallest, it's either min-nx, or min+p1 (nx is smallest-absolute negative, p1 is smallest-absolute positive).
// which means for d=secondSmallestSum-smallestSum, +d or -d is a possible element, and it's the minimal-absolute element.
// the subset sums can be parted in two parts {x1, x2, ..., xn} and {x1+d, x2+d, .., xn+d}.
// so the original set is either {d} ∪ recover({x1, x2, ..., xn}), or {-d} ∪ recover({x1+d, x2+d, ..., xn+d}).
// then we can use recursion to compute.
func recoverArray(n int, sums []int) []int {
	return dfs(n, sums)
}

func dfs(n int, sums []int) []int {
	if len(sums) == 2 { // for single-element set {x}, the subset sums must be {0, x}
		if sums[0] == 0 {
			return []int{sums[1]}
		} else if sums[1] == 0 {
			return []int{sums[0]}
		} else {
			return []int{}
		}
	}
	sort.Sort(sort.IntSlice(sums))
	d := sums[1] - sums[0]
	if d == 0 {
		sub := make([]int, 0, len(sums)/2)
		for i := 0; i < len(sums); i += 2 {
			sub = append(sub, sums[i])
		}
		subans := dfs(n-1, sub)
		ans := make([]int, 1, n)
		ans[0] = 0
		return append(ans, subans...)
	}
	// part the two parts
	count := make(map[int]int)
	for _, v := range sums {
		count[v] = count[v] + 1
	}
	sub1 := make([]int, 0, len(sums)/2)
	sub2 := make([]int, 0, len(sums)/2)
	for _, v := range sums {
		if count[v] == 0 {
			continue
		}
		if count[v+d] == 0 {
			return []int{} // invalid
		}
		count[v] = count[v] - 1
		count[v+d] = count[v+d] - 1
		sub1 = append(sub1, v)
		sub2 = append(sub2, v+d)
	}
	subans := dfs(n-1, sub1) // {d} + recursion({x1, x2, ..., xn})
	if len(subans) == 0 {
		d = -d
		subans = dfs(n-1, sub2) // {-d} + recursion({x1+d, x2+d, ..., xn+d})
	}
	if len(subans) == 0 {
		return []int{}
	}
	ans := make([]int, 1, n)
	ans[0] = d
	return append(ans, subans...)
}

func main() {
	for _, v := range []struct {
		n      int
		s, ans []int
	}{
		{3, []int{-3, -2, -1, 0, 0, 1, 2, 3}, []int{1, 2, -3}},
		{2, []int{0, 0, 0, 0}, []int{0, 0}},
		{4, []int{0, 0, 5, 5, 4, -1, 4, 9, 9, -1, 4, 3, 4, 8, 3, 8}, []int{0, -1, 4, 5}},
		{1, []int{5, 0}, []int{5}},
		{3, []int{365, 44, -355, 399, 409, 764, 10, 0}, []int{-355, 365, 399}},
	} {
		fmt.Println(recoverArray(v.n, v.s), v.ans)
	}
}

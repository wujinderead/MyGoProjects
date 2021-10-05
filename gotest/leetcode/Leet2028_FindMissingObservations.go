package main

import "fmt"

// https://leetcode.com/problems/find-missing-observations/

// You have observations of n + m 6-sided dice rolls with each face numbered from 1 to 6.
// n of the observations went missing, and you only have the observations of m rolls. Fortunately,
// you have also calculated the average value of the n + m rolls.
// You are given an integer array rolls of length m where rolls[i] is the value of the ith observation.
// You are also given the two integers mean and n.
// Return an array of length n containing the missing observations such that the average value of the
// n + m rolls is exactly mean. If there are multiple valid answers, return any of them.
// If no such array exists, return an empty array.
// The average value of a set of k numbers is the sum of the numbers divided by k.
// Note that mean is an integer, so the sum of the n + m rolls should be divisible by n + m.
// Example 1:
//   Input: rolls = [3,2,4,3], mean = 4, n = 2
//   Output: [6,6]
//   Explanation: The mean of all n + m rolls is (3 + 2 + 4 + 3 + 6 + 6) / 6 = 4.
// Example 2:
//   Input: rolls = [1,5,6], mean = 3, n = 4
//   Output: [2,3,2,2]
//   Explanation: The mean of all n + m rolls is (1 + 5 + 6 + 2 + 3 + 2 + 2) / 7 = 3.
// Example 3:
//   Input: rolls = [1,2,3,4], mean = 6, n = 4
//   Output: []
//   Explanation: It is impossible for the mean to be 6 no matter what the 4 missing rolls are.
// Example 4:
//   Input: rolls = [1], mean = 3, n = 1
//   Output: [5]
//   Explanation: The mean of all n + m rolls is (1 + 5) / 2 = 3.
// Constraints:
//   m == rolls.length
//   1 <= n, m <= 10^5
//   1 <= rolls[i], mean <= 6

func missingRolls(rolls []int, mean int, n int) []int {
	need := mean * (len(rolls) + n)
	sum := 0
	for _, v := range rolls {
		sum += v
	}
	need = need - sum
	if !(need >= n && need <= 6*n) {
		return []int{}
	}
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = need / n
	}
	for i := 0; i < need%n; i++ {
		ans[i]++
	}
	return ans
}

func main() {
	for _, v := range []struct {
		r    []int
		m, n int
		ans  []int
	}{
		{[]int{3, 2, 4, 3}, 4, 2, []int{6, 6}},
		{[]int{1, 5, 6}, 3, 4, []int{2, 3, 2, 2}},
		{[]int{1, 2, 3, 4}, 6, 4, []int{}},
		{[]int{1}, 3, 1, []int{5}},
		{[]int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}, 4, 40,
			[]int{4, 4, 4, 4, 4, 4, 5, 5, 4, 4, 4, 5, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5}},
	} {
		fmt.Println(missingRolls(v.r, v.m, v.n), v.ans)
	}
}

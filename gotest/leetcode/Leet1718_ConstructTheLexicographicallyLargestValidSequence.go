package main

import "fmt"

// https://leetcode.com/problems/construct-the-lexicographically-largest-valid-sequence/

// Given an integer n, find a sequence that satisfies all of the following:
//   The integer 1 occurs once in the sequence.
//   Each integer between 2 and n occurs twice in the sequence.
//   For every integer i between 2 and n, the distance between the two occurrences of i is exactly i.
// The distance between two numbers on the sequence, a[i] and a[j], is the absolute
// difference of their indices, |j - i|.
// Return the lexicographically largest sequence. It is guaranteed that under the given constraints,
// there is always a solution.
// A sequence a is lexicographically larger than a sequence b (of the same length) if in the first
// position where a and b differ, sequence a has a number greater than the corresponding number in b.
// For example, [0,1,9,0] is lexicographically larger than [0,1,5,6] because the first position
// they differ is at the third number, and 9 is greater than 5.
// Example 1:
//   Input: n = 3
//   Output: [3,1,2,3,2]
//   Explanation: [2,3,2,1,3] is also a valid sequence,
//     but [3,1,2,3,2] is the lexicographically largest valid sequence.
// Example 2:
//   Input: n = 5
//   Output: [5,3,1,4,3,5,2,4,2]
// Constraints:
//   1 <= n <= 20

// use backtracking
func constructDistancedSequence(n int) []int {
	ans := make([]int, 2*n-1)
	use := make([]bool, n+1)
	f(ans, use, 0)
	return ans
}

func f(ans []int, use []bool, all int) bool {
	// all position are set
	if all == len(ans) {
		return true
	}
	// get the first position to populate
	var ind int
	for i := 0; i < len(ans); i++ {
		if ans[i] == 0 {
			ind = i
			break
		}
	}
	for i := len(use) - 1; i >= 1; i-- {
		if !use[i] && (i == 1 || (ind+i < len(ans) && ans[ind+i] == 0)) {
			// populate the position with a number
			use[i] = true
			if i == 1 {
				ans[ind] = 1
				all += 1
			} else {
				ans[ind] = i
				ans[ind+i] = i
				all += 2
			}
			// move to next
			if f(ans, use, all) {
				return true // if valid, then it's the largest, we can go back
			}
			// reset if in valid
			use[i] = false
			if i == 1 {
				ans[ind] = 0
				all -= 1
			} else {
				ans[ind] = 0
				ans[ind+i] = 0
				all -= 2
			}
		}
	}
	return false
}

func main() {
	for _, v := range []struct {
		n   int
		ans []int
	}{
		{1, []int{1}},
		{2, []int{2, 1, 2}},
		{3, []int{3, 1, 2, 3, 2}},
		{4, []int{4, 2, 3, 2, 4, 3, 1}},
		{5, []int{5, 3, 1, 4, 3, 5, 2, 4, 2}},
		{6, []int{6, 4, 2, 5, 2, 4, 6, 3, 5, 1, 3}},
	} {
		fmt.Println(constructDistancedSequence(v.n), v.ans)
	}
}

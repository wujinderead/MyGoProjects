package main

import "fmt"

// https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/

// You are given two integer arrays nums and multipliers of size n and m respectively,
// where n >= m. The arrays are 1-indexed.
// You begin with a score of 0. You want to perform exactly m operations. On the ith
// operation (1-indexed), you will:
//   Choose one integer x from either the start or the end of the array nums.
//   Add multipliers[i] * x to your score.
//   Remove x from the array nums.
// Return the maximum score after performing m operations.
// Example 1:
//   Input: nums = [1,2,3], multipliers = [3,2,1]
//   Output: 14
//   Explanation: An optimal solution is as follows:
//     - Choose from the end, [1,2,3], adding 3 * 3 = 9 to the score.
//     - Choose from the end, [1,2], adding 2 * 2 = 4 to the score.
//     - Choose from the end, [1], adding 1 * 1 = 1 to the score.
//     The total score is 9 + 4 + 1 = 14.
// Example 2:
//   Input: nums = [-5,-3,-3,-2,7,1], multipliers = [-10,-5,3,4,6]
//   Output: 102
//   Explanation: An optimal solution is as follows:
//     - Choose from the start, [-5,-3,-3,-2,7,1], adding -5 * -10 = 50 to the score.
//     - Choose from the start, [-3,-3,-2,7,1], adding -3 * -5 = 15 to the score.
//     - Choose from the start, [-3,-2,7,1], adding -3 * 3 = -9 to the score.
//     - Choose from the end, [-2,7,1], adding 1 * 4 = 4 to the score.
//     - Choose from the end, [-2,7], adding 7 * 6 = 42 to the score.
//     The total score is 50 + 15 - 9 + 4 + 42 = 102.
// Constraints:
//   n == nums.length
//   m == multipliers.length
//   1 <= m <= 10^3
//   m <= n <= 10^5
//   -1000 <= nums[i], multipliers[i] <= 1000

func maximumScore(nums []int, multipliers []int) int {
	dp := make(map[[2]int]int)
	n, m := len(nums), len(multipliers)
	for diff := n - m; diff < len(nums); diff++ {
		for i := 0; i+diff < len(nums); i++ {
			j := i + diff
			m := multipliers[m-1-(diff-n+m)]
			dp[[2]int{i, j}] = max(m*nums[i]+dp[[2]int{i + 1, j}], m*nums[j]+dp[[2]int{i, j - 1}])
		}
	}
	return dp[[2]int{0, len(nums) - 1}]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		num, mul []int
		ans      int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}, 14},
		{[]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}, 102},
	} {
		fmt.Println(maximumScore(v.num, v.mul), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/longest-increasing-subsequence-ii/

// You are given an integer array nums and an integer k.
// Find the longest subsequence of nums that meets the following requirements:
//   The subsequence is strictly increasing and
//   The difference between adjacent elements in the subsequence is at most k.
// Return the length of the longest subsequence that meets the requirements.
// A subsequence is an array that can be derived from another array by deleting some or no elements
// without changing the order of the remaining elements.
// Example 1:
//   Input: nums = [4,2,1,4,3,4,5,8,15], k = 3
//   Output: 5
//   Explanation:
//     The longest subsequence that meets the requirements is [1,3,4,5,8].
//     The subsequence has a length of 5, so we return 5.
//     Note that the subsequence [1,3,4,5,8,15] does not meet the requirements because
//     15 - 8 = 7 is larger than 3.
// Example 2:
//   Input: nums = [7,4,5,1,8,12,4,7], k = 5
//   Output: 4
//   Explanation:
//     The longest subsequence that meets the requirements is [4,5,8,12].
//     The subsequence has a length of 4, so we return 4.
// Example 3:
//   Input: nums = [1,5], k = 1
//   Output: 1
//   Explanation:
//     The longest subsequence that meets the requirements is [1].
//     The subsequence has a length of 1, so we return 1.
// Constraints:
//   1 <= nums.length <= 10⁵
//   1 <= nums[i], k <= 10⁵

// let dp[i][val] be the answer using nums[...i], and the last element of subsequence is val.
// then if we don't use nums[i], then dp[i][x] = dp[i-1][x] (for each valid x); if we use nums[i]=val,
// then dp[i][val] = max(dp[i-1][x])+1 for val-k <= x < val
func lengthOfLIS(nums []int, k int) int {
	N := 100000
	tree := make([]int, (1+N)*4)
	ans := 1
	for _, v := range nums {
		q := query(tree, 0, 0, N, max(0, v-k), v-1) // query max(dp[i-1][v-k...v-1])
		ori := query(tree, 0, 0, N, v, v)           // query original dp[i][v]
		if q+1 > ori {
			if q+1 > ans {
				ans = q + 1
			}
			update(tree, 0, 0, N, v, q+1) // update dp[i][v]
		}
	}
	return ans
}

func query(tree []int, ind, left, right, queryLeft, queryRight int) int {
	leftSubInd, rightSubInd, mid := 2*ind+1, 2*ind+2, left+(right-left)/2

	// remain is the same to normal query
	if queryLeft == left && queryRight == right {
		return tree[ind]
	}
	if queryRight <= mid {
		return query(tree, leftSubInd, left, mid, queryLeft, queryRight)
	} else if queryLeft > mid {
		return query(tree, rightSubInd, mid+1, right, queryLeft, queryRight)
	}
	l := query(tree, leftSubInd, left, mid, queryLeft, mid)
	r := query(tree, rightSubInd, mid+1, right, mid+1, queryRight)
	return max(l, r)
}

// update range with new value
func update(tree []int, ind, left, right, updateInd, updateVal int) {
	leftSubInd, rightSubInd, mid := 2*ind+1, 2*ind+2, left+(right-left)/2
	// if update range equal to data range, update tree node
	if left == right {
		tree[ind] = max(updateVal, tree[ind])
		return
	}
	if updateInd <= mid { // only update left part
		update(tree, leftSubInd, left, mid, updateInd, updateVal)
	} else { // only update right part
		update(tree, rightSubInd, mid+1, right, updateInd, updateVal)
	}
	// merge two sub trees
	tree[ind] = max(tree[leftSubInd], tree[rightSubInd])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct {
		nums   []int
		k, ans int
	}{
		{[]int{4, 2, 1, 4, 3, 4, 5, 8, 15}, 3, 5},
		{[]int{7, 4, 5, 1, 8, 12, 4, 7}, 5, 4},
		{[]int{1, 5}, 1, 1},
		{[]int{9}, 1, 1},
	} {
		fmt.Println(lengthOfLIS(v.nums, v.k), v.ans)
	}
}

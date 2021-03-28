package main

import "fmt"

// https://leetcode.com/problems/maximum-value-at-a-given-index-in-a-bounded-array/

// You are given three positive integers n, index and maxSum. You want to construct
// an array nums (0-indexed) that satisfies the following conditions:
//   nums.length == n
//   nums[i] is a positive integer where 0 <= i < n.
//   abs(nums[i] - nums[i+1]) <= 1 where 0 <= i < n-1.
//   The sum of all the elements of nums does not exceed maxSum.
//   nums[index] is maximized.
// Return nums[index] of the constructed array.
// Note that abs(x) equals x if x >= 0, and -x otherwise.
// Example 1:
//   Input:
//   Output: 2
//   Explanation: The arrays [1,1,2,1] and [1,2,2,1] satisfy all the conditions.
//     There are no other valid arrays with a larger value at the given index.
// Example 2:
//   Input: n = 6, index = 1,  maxSum = 10
//   Output: 3
// Constraints:
//   1 <= n <= maxSum <= 10^9
//   0 <= index < n

// e.g., n=6, index=1, maxSum=10 means for len(arr)=6 we want a[1] maximal,
// if a[1]=4, the minimal sum array is [3,4,3,2,1,0], sum=13 > 10
// if a[1]=3, the minimal sum array is [2,3,2,1,0,0], sum=8 <= 10
// so 3 is the max valid value at a[1]. use binary search to find the value.
func maxValue(n int, index int, maxSum int) int {
	l, r := 1, maxSum
	for l <= r {
		mid := l + (r-l)/2
		// get the minimal sum when a[i]=mid
		sum := 0
		if mid <= index {
			sum += (1+mid)*mid/2 + index - mid + 1
		} else {
			sum += (mid + mid - index) * (index + 1) / 2
		}
		if mid-(n-1-index) <= 1 {
			sum += (1+mid)*mid/2 + n - index - mid
		} else {
			sum += (mid + mid - (n - 1 - index)) * (n - index) / 2
		}
		sum -= mid

		// when l=r=mid, mid yes, mid+1 no; then l=mid+1, r=mid, return r=mid
		// when l=r=mid, mid-1 yes, mid no; then l=mid, r=mid-1, return r=mid-1
		if sum <= maxSum {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

func main() {
	for _, v := range []struct {
		n, ind, max, ans int
	}{
		{4, 2, 6, 2},
		{6, 1, 10, 3},
		{4, 0, 4, 1},
	} {
		fmt.Println(maxValue(v.n, v.ind, v.max), v.ans)
	}
}

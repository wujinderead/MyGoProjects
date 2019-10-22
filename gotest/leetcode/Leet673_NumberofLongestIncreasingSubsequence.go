package leetcode

import "fmt"

// https://leetcode.com/problems/number-of-longest-increasing-subsequence/

// Given an unsorted array of integers, find the number of longest increasing subsequence.
// Example 1:
//   Input: [1,3,5,4,7]
//   Output: 2
//   Explanation: The two longest increasing subsequence are [1, 3, 4, 7] and [1, 3, 5, 7].
// Example 2:
//   Input: [2,2,2,2,2]
//   Output: 5
//   Explanation: The length of longest continuous increasing subsequence is 1,
//                and there are 5 subsequences' length is 1, so output 5.

func findNumberOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	max := 1
	lis := make([]int, len(nums))
	pre := make([]int, len(nums))
	lis[0] = 1
	for i := 1; i < len(nums); i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && lis[j]+1 > lis[i] {
				lis[i] = lis[j] + 1
				if lis[i] > max {
					max = lis[i]
				}
				if pre[j] > 0 {
					pre[i] = pre[j]
				} else {
					pre[i] = 1
				}
			} else if lis[j]+1 == lis[i] {
				if pre[j] > 0 {
					pre[i] += pre[j]
				} else {
					pre[i] += 1
				}
			}
		}
	}
	//fmt.Println(nums)
	//fmt.Println(lis)
	//fmt.Println(pre)
	if max == 1 {
		return len(nums)
	}
	count := 0
	for i := range lis {
		if lis[i] == max {
			count += pre[i]
		}
	}
	return count
}

func findNumberOfLIS1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	lis := make([]int, len(nums))
	cnt := make([]int, len(nums))

	for j := 0; j < len(nums); j++ {
		cnt[j] = 1
		lis[j] = 1
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] { // i<j && nums[i]<nums[j]
				if lis[i] >= lis[j] {
					lis[j] = lis[i] + 1
					cnt[j] = cnt[i]
				} else if lis[i]+1 == lis[j] {
					cnt[j] += cnt[i]
				}
			}
		}
	}
	max := 1
	for i := range lis {
		if lis[i] > max {
			max = lis[i]
		}
	}
	count := 0
	for i := range lis {
		if lis[i] == max {
			count += cnt[i]
		}
	}
	return count
}

func main() {
	fmt.Println(findNumberOfLIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS([]int{2, 2, 2, 2, 2}))
	fmt.Println(findNumberOfLIS([]int{1, 2, 4, 3, 5, 4, 7, 2}))
	fmt.Println(findNumberOfLIS([]int{1}))
	fmt.Println(findNumberOfLIS([]int{1, 2}))
	fmt.Println(findNumberOfLIS([]int{1, 2, 2, 3, 2, 3}))
	fmt.Println(findNumberOfLIS([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}))

	fmt.Println(findNumberOfLIS1([]int{1, 3, 5, 4, 7}))
	fmt.Println(findNumberOfLIS1([]int{2, 2, 2, 2, 2}))
	fmt.Println(findNumberOfLIS1([]int{1, 2, 4, 3, 5, 4, 7, 2}))
	fmt.Println(findNumberOfLIS1([]int{1}))
	fmt.Println(findNumberOfLIS1([]int{1, 2}))
	fmt.Println(findNumberOfLIS1([]int{1, 2, 2, 3, 2, 3}))
	fmt.Println(findNumberOfLIS1([]int{1, 1, 1, 2, 2, 2, 3, 3, 3}))
}

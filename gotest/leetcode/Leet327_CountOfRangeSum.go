package main

import "fmt"

// https://leetcode.com/problems/count-of-range-sum/

// Given an integer array nums, return the number of range sums that lie in [lower, upper] inclusive.
// Range sum S(i, j) is defined as the sum of the elements in nums between indices i and j (i ≤ j), inclusive.
// Note: A naive algorithm of O(n^2) is trivial. You MUST do better than that.
// Example:
//   Input: nums = [-2,5,-1], lower = -2, upper = 2,
//   Output: 3 
//   Explanation: The three ranges are : [0,0], [2,2], [0,2] and their respective sums are: -2, -1, 2.
// Constraints:
//   0 <= nums.length <= 10^4

// get the prefix sum array, for prefix[i], if we want the range sum between [lower, upper],
// for prefix[0...i-1], we need find those in the range [prefix[i]-upper, prefix[i]-lower].
func countRangeSum(nums []int, lower int, upper int) int {
	if len(nums)==0 {
		return 0
	}

	// get prefix sum of nums
    prefix := make([]int, len(nums))
    prev := 0
    count := 0
    for i:=range nums {
    	prefix[i] = prev+nums[i]
    	if prefix[i]>=lower && prefix[i]<=upper {
    		count++
    	}
    	prev = prefix[i]
    }

    // merge sort prefix
    buf := make([]int, len(prefix))   // use buf to allocate memory only once
    mergeSort(prefix, buf, 0, len(prefix)-1, lower, upper, &count)
    return count
}

func mergeSort(data, buf []int, start, end, lower, upper int, count *int) {
	if start==end {
		return 
	}
	mid := start+(end-start)/2
	mergeSort(data, buf, start, mid, lower, upper, count)
	mergeSort(data, buf, mid+1, end, lower, upper, count)
	
	// count what we want
	left, right := data[start: mid+1], data[mid+1: end+1]   // switch to 0-indexed
	// for each number x of right, find those numbers in left that in range [x-upper, x-lower]
	l, r := 0, 0
	for j:=0; j<len(right); j++ {
		low, up := right[j]-upper, right[j]-lower  // low, up monotonously increase
		if left[0]>up {          //                 [left0    leftn-1]
			continue             //  [low  up] ->>                         
		}
		if left[len(left)-1]<low {    //  [left0    leftn-1]
			break                     //                       [low  up] ->>
		}
		for l<len(left) && left[l]<low {       // l is fisrt x that left[x]>=low
			l++
		}
		for r+1<len(left) && left[r+1]<=up {   // r is last x that left[x]<=up
			r++
		}
		*count += r-l+1
	}

	// merge 2 sorted list
	i, j := 0, 0
	for i<len(left) || j<len(right) {
		if i<len(left) && ((j<len(right) && left[i]<=right[j]) || j>=len(right)) {
			buf[i+j] = left[i]
			i++
		} else {
			buf[i+j] = right[j]
			j++
		}
	}
	copy(data[start:], buf[:len(left)+len(right)])
}

func main() {
	fmt.Println(countRangeSum([]int{-3,6,8,-5,4,-9,2,10,-1,7}, 4, 10), 19)
	fmt.Println(countRangeSum([]int{-3,6,8,-5,4,-9,2,10,-1,7}, 5, 10), 17)
	fmt.Println(countRangeSum([]int{-2,5,-1}, -2, 2), 3)
	fmt.Println(countRangeSum([]int{2}, -2, 2), 1)
	fmt.Println(countRangeSum([]int{3}, -2, 2), 0)
	fmt.Println(countRangeSum([]int{1,3}, 0, 2), 1)
}

func verify(arr []int, low, up int) int {
	prefix := make([]int, len(arr)+1)
	for i:=range arr {
		prefix[i+1] = prefix[i]+arr[i]
	}
	count := 0
	for i:=1; i<len(prefix); i++ {
		for j:=0; j<i; j++ {
			if prefix[i]-prefix[j]>=low && prefix[i]-prefix[j]<=up {
				count++
			}
		}
	}
	return count
}

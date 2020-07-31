package main

import "fmt"

// https://leetcode.com/problems/reverse-pairs/

// Given an array nums, we call (i, j) an important reverse pair if i < j and nums[i] > 2*nums[j].
// You need to return the number of important reverse pairs in the given array.
// Example1:
//   Input: [1,3,2,3,1]
//   Output: 2
// Example2:
//   Input: [2,4,3,5,1]
//   Output: 3
// Note:
//   The length of the given array will not exceed 50,000.
//   All the numbers in the input array are in the range of 32-bit integer.

// merge sort based method, O(nlogn)
func reversePairs(nums []int) int {
	if len(nums)==0 {
		return 0
	}
	count := new(int)
	buf := make([]int, len(nums))
	mergeSort(nums, buf, count, 0, len(nums)-1)
    return *count
}

func mergeSort(data, buf []int, count *int, start, end int) {
	if start==end {
		return 
	}
	mid := start + (end-start)/2
	mergeSort(data, buf, count, start, mid)
	mergeSort(data, buf, count, mid+1, end)

	// merge 2 sorted lists
	i, j := start, mid+1
	k := j-1   // k is the max value that makes data[i] > 2*data[k+1]
	bi := 0
	for i<=mid || j<=end {
		// if we need to add a left element (x) to list,
		// we check how many right elements (y1, y2, ...yk) are already in the list, 
		// these (x>2*y1), (x>2*y2), ... are reverse pairs
		// when x increases, k may also increase monotonously, so it's still linear time.
		if i<=mid && ((j<=end && data[i] <= data[j]) || j > end) {   
			for k+1<=end && data[i] > 2*data[k+1] {
				k++
			}
			buf[bi] = data[i]               
			*count += k-mid   // add count
			i++
			bi++
		} else {
			buf[bi] = data[j]
			j++
			bi++
		}
	}
	copy(data[start: end+1], buf[:end-start+1])
}

func main() {
	for _, v := range []struct{arr []int; ans int} {
		{[]int{5,2,6,1}, 3},
		{[]int{1,2}, 0},
		{[]int{2,1}, 0},
		{[]int{3,1}, 1},
		{[]int{12,5,9,3,27,18,8,13}, 6},
		{[]int{1,3,2,3,1}, 2},
		{[]int{2,4,3,5,1}, 3},
	} {
		fmt.Println(reversePairs(v.arr), v.ans)
	}
}

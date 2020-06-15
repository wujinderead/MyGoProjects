package main

import (
    "fmt"
)

// https://leetcode.com/problems/max-chunks-to-make-sorted-ii/

// This question is the same as "Max Chunks to Make Sorted" except the integers of the given array are not 
// necessarily distinct, the input array could be up to length 2000, and the elements could be up to 10**8. 
// Given an array arr of integers (not necessarily distinct), we split the array into some number of 
// "chunks" (partitions), and individually sort each chunk. After concatenating them, the result equals 
// the sorted array. What is the most number of chunks we could have made? 
// Example 1: 
//   Input: arr = [5,4,3,2,1]
//   Output: 1
//   Explanation:
//    Splitting into two or more chunks will not return the required result.
//    For example, splitting into [5, 4], [3, 2, 1] will result in [4, 5, 1, 2, 3], which isn't sorted.
// Example 2: 
//   Input: arr = [2,1,3,4,4]
//   Output: 4
//   Explanation:
//     We can split into two chunks, such as [2, 1], [3, 4, 4].
//     However, splitting into [2, 1], [3], [4], [4] is the highest number of chunks possible.
// Note: 
//   arr will have length in range [1, 2000]. 
//   arr[i] will be an integer in range [0, 10**8]. 

func maxChunksToSorted(arr []int) int {
	// we can split at a position where all left values are less than all right values.
	// i.e, the max value of left is less than the min value of right
	min := make([]int, len(arr)+1)  // min[i] is the minimal number of 
	min[len(arr)] = int(1e9)
	for i:=len(arr)-1; i>=0; i-- {
		min[i] = min[i+1]
		if arr[i]<min[i] {
			min[i] = arr[i]
		}
	}
	chunks := 0
	max := arr[0]
	for i:=range arr {
		if arr[i]>max {
			max = arr[i]
		}
		if max<=min[i+1] {   // if max(arr[:i]) <= min(arr[i+1:]), we can split after i
			chunks++
		}
	}
	return chunks
}

func main() {
	fmt.Println(maxChunksToSorted([]int{5,4,3,2,1}))
	fmt.Println(maxChunksToSorted([]int{2,1,3,4,4}))
	fmt.Println(maxChunksToSorted([]int{1}))
	fmt.Println(maxChunksToSorted([]int{5,4,3,3,5,9,8,8,7,10,12}))
}
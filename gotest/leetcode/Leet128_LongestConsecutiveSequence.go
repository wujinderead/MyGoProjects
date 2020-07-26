package main

import (
	"fmt"
)

// https://leetcode.com/problems/longest-consecutive-sequence/

// Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
// Your algorithm should run in O(n) complexity.
// Example:
//   Input: [100, 4, 200, 1, 3, 2]
//   Output: 4
//   Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

func longestConsecutive(nums []int) int {
	if len(nums)==0 {
		return 0
	}
	mapp := make(map[int]int)
	size := make(map[int]int)
	for _, v := range nums {
		if _, ok := mapp[v]; ok {
			continue    // ignore repeated
		}
		_, ok1 := mapp[v-1]
		_, ok2 := mapp[v+1]
		if ok1 && ok2 {
			r1 := getRoot(mapp, v-1)
			r2 := getRoot(mapp, v+1)
			mapp[v] = r1
			mapp[r2] = r1
			size[r1] = size[r1]+size[r2]+1			
			continue
		}
		if ok1 {
			r1 := getRoot(mapp, v-1)
			mapp[v] = r1
			size[r1] = size[r1]+1
			continue
		}
		if ok2 {
			r2 := getRoot(mapp, v+1)
			mapp[v] = r2
			size[r2] = size[r2]+1
			continue
		}

		// first occur, no neighbor
		mapp[v] = v
		size[v] = 1
	}
	max := 1
	for _, v := range size {
		if v>max {
			max = v
		}
	}
	return max
}

func getRoot(mapp map[int]int, v int) int {
	for mapp[v] != v {
		v = mapp[v]
	}
	return v
}

func longestConsecutiveNaive(nums []int) int {
	if len(nums)==0 {
		return 0
	}
	max := 1
	mapp := make(map[int]int)
	for _, v := range nums {
		if _, ok := mapp[v]; ok {
			continue    // ignore repeated
		}
		v1 := mapp[v-1]   // the length for left neighbor
		v2 := mapp[v+1]   // the length for right neighbor
		
		mapp[v] = v1+v2+1   // update length
		if v1+v2+1 > max {
			max = v1+v2+1
		}
		// update length for edge node
		mapp[v-v1] = v1+v2+1
		mapp[v+v2] = v1+v2+1
	}
	return max
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{1, 3, 0, 101, 6, 7, 100, 99, 98, 5, 2, 97, 4}))
	fmt.Println(longestConsecutive([]int{2,6,4,8,10,22,11,15}))
	fmt.Println(longestConsecutiveNaive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutiveNaive([]int{1, 3, 0, 101, 6, 7, 100, 99, 98, 5, 2, 97, 4}))
	fmt.Println(longestConsecutiveNaive([]int{2,6,4,8,10,22,11,15}))
}
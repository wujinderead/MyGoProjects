package main

import (
	"sort"
    "fmt"
)

// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/

// Given an array of integers arr and an integer k. Find the least number of unique integers 
// after removing exactly k elements. 
// Example 1: 
//   Input: arr = [5,5,4], k = 1
//   Output: 1
//   Explanation: Remove the single 4, only 5 is left.
// Example 2:
//   Input: arr = [4,3,1,1,3,3,2], k = 3
//   Output: 2
//   Explanation: Remove 4, 2 and either one of the two 1s or three 3s. 1 and 3 will be left.
// Constraints: 
//   1 <= arr.length <= 10^5 
//   1 <= arr[i] <= 10^9 
//   0 <= k <= arr.length 

func findLeastNumOfUniqueInts(arr []int, k int) int {
    count := make(map[int]int)
    for _, v := range arr {        // number, count map
    	count[v] = count[v]+1
    }
    allkind := len(count)          // kind of distinct integer
    remap := make(map[int]int)     // count, kind map
    for _, v := range count {
    	remap[v] = remap[v]+1 
    }
    counts := make([]int, 0, len(remap))   // unique count
    for k := range remap {
        counts = append(counts, k)
    }
    sort.Sort(sort.IntSlice(counts))

    // remove those integers whose count is low
    for _, v := range counts {
        n := remap[v]   // n kinds of number have count v
        if k>=n*v {
            k = k-n*v
            allkind -= n
        } else {
            allkind -= k/v
            break
        }
    }
    return allkind
}

func main() {
	fmt.Println(findLeastNumOfUniqueInts([]int{5,5,4}, 1))
	fmt.Println(findLeastNumOfUniqueInts([]int{4,3,1,1,3,3,2}, 3))
    fmt.Println(findLeastNumOfUniqueInts([]int{1,2,3,3,4,4,4,5,5,5}, 5))
}
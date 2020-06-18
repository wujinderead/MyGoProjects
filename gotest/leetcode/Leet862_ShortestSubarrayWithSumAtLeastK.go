package main

import (
    "fmt"
)

// https://leetcode.com/problems/maximum-subarray-sum-with-one-deletion/

// Return the length of the shortest, non-empty, contiguous subarray of A with sum at least K.
// If there is no non-empty subarray with sum at least K, return -1.
// Example 1:
//   Input: A = [1], K = 1
//   Output: 1
// Example 2:
//   Input: A = [1,2], K = 4
//   Output: -1
// Example 3:
//   Input: A = [2,-1,2], K = 3
//   Output: 3
// Note:
//   1 <= A.length <= 50000
//   -10 ^ 5 <= A[i] <= 10 ^ 5
//   1 <= K <= 10 ^ 9 

// like sliding window or monotonic queue.
func shortestSubarray(A []int, K int) int {
    allmin := int(1e6)

    // get prefix array
	prefix := make([]int, len(A)+1)
    for i:=0; i<len(A); i++ {
        prefix[i+1] = prefix[i]+A[i]
    }
    //fmt.Println(prefix)
    
    queue := []int{0}  // a queue
    front := 0         // queue front pointer
    for ind := range prefix {
        i := len(queue)-1
        // actually the following 2 loops are exclusive
        for i>=front && prefix[queue[i]]>prefix[ind] {    // remove larger values in rear 
            queue = queue[:i]
            i--
        }
        for front<len(queue) && prefix[ind]-prefix[queue[front]]>=K {  // remove smaller values in front
            if ind-queue[front]<allmin {
                allmin = ind-queue[front]
            }
            front++
        }
        queue = append(queue, ind)
    }

    if allmin==int(1e6) {
        return -1
    }
    return allmin
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	fmt.Println(shortestSubarray([]int{1}, 1))
	fmt.Println(shortestSubarray([]int{1,2}, 4))
	fmt.Println(shortestSubarray([]int{2,-1,2}, 3))
    fmt.Println(shortestSubarray([]int{17,85,93,-45,-21}, 150))
}
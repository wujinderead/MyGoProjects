package main

import (
    "fmt"
)

// https://leetcode.com/problems/subarrays-with-k-different-integers/

// Given an array A of positive integers, call a (contiguous, not necessarily distinct) 
// subarray of A good if the number of different integers in that subarray is exactly K. 
// (For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.) 
// Return the number of good subarrays of A. 
// Example 1: 
//   Input: A = [1,2,1,2,3], K = 2
//   Output: 7
//   Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1],
//     [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2].
// Example 2: 
//  Input: A = [1,2,1,3,4], K = 3
//  Output: 3
//  Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
// Note: 
//   1 <= A.length <= 20000 
//   1 <= A[i] <= A.length 
//   1 <= K <= A.length 

func subarraysWithKDistinct(A []int, K int) int {
    allcount := 0
    vmap := make(map[int]int)   // count for valid segment
    umap := make(map[int]int)   // count for invalid segment
    vtype, utype := 0, 0
    v, u := 0, 0
    // for each A[i]=e, we track the most left v and u that makes A[v...i] valid and A[u...i] invalid
    // thus, for each v≤j<u, A[j...i] is a valid subarray
    for _, e := range A {
    	vcount := vmap[e]
    	vmap[e] = vcount+1
    	ucount := umap[e]
    	umap[e] = ucount+1
    	if vcount==0 {   // an new element
    		vtype++
    	}
    	if ucount==0 {   // an new element
    		utype++
    	}
    	if utype==K {     // A[u...i] valid, move u to right to make it invalid (less than K)
    		for {
    			ucount := umap[A[u]]
    			umap[A[u]] = ucount-1
    			u++
    			if ucount-1==0 {   // delete A[u], if we can eliminate one type, we make A[u...i] invalid
    				utype--
    				break
    			}
    		}
    	}
    	if vtype>K {      // A[v...i] invalid (more than K), make v to right to make it valid
    		for {
    			vcount := vmap[A[v]]
    			vmap[A[v]] = vcount-1
    			v++
    			if vcount-1==0 {  // delete A[v], if we can eliminate one type, we make A[v...i] valid
    				vtype--
    				break
    			}
    		}
    	}
    	allcount += u-v   // for each v≤j<u, A[j...i] is a valid subarray
    }
    return allcount
}

func main() {
	fmt.Println(subarraysWithKDistinct([]int{1,2,1,2,3}, 2), 7)
	fmt.Println(subarraysWithKDistinct([]int{1,2,1,3,4}, 3), 3)
	fmt.Println(subarraysWithKDistinct([]int{1,1,2,1,2,2,3,3}, 1), 11)
	fmt.Println(subarraysWithKDistinct([]int{1,1,2,1,2,2,3,3}, 2), 17)
	fmt.Println(subarraysWithKDistinct([]int{20,11,37,2,22,45,31,46,6,51,27,7,21,25,20,38,42,17,12,32,18,25,25,30,11,5,38,7,
		49,18,40,12,44,11,50,15,12,40,36,36,43,42,2,11,46,20,18,13,37,37,37,23,31,36,23,21,38,31,42,32,5,24,35,29,25,11,29,
		32,27,42,43,17,45,36,10,22,45,46,50,23,20,37,24,49,28,42,18,20,11,16,3,49,2,50,51,21,47,42,40,7}, 30), 179)
}
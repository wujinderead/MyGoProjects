package main

import (
	"fmt"
	"math/rand"
)

// https://leetcode.com/problems/couples-holding-hands/
// N couples sit in 2N seats arranged in a row and want to hold hands. We want to
// know the minimum number of swaps so that every couple is sitting side by side. 
// A swap consists of choosing ANY two people, then they stand up and switch seats.
// The people and seats are represented by an integer from 0 to 2N-1, the couples
// are numbered in order, the first couple being (0, 1), the second couple being 
// (2, 3), and so on with the last couple being (2N-2, 2N-1).
// The couples' initial seating is given by row[i] being the value of the person 
// who is initially sitting in the i-th seat.
// Example 1: 
//   Input: row = [0, 2, 1, 3]
//   Output: 1
//   Explanation: We only need to swap the second (row[1]) and third (row[2]) person.
// Example 2: 
//   Input: row = [3, 2, 0, 1]
//   Output: 0
//   Explanation: All couples are already seated side by side.
// Note:
//   len(row) is even and in the range of [4, 60]. 
//   row is guaranteed to be a permutation of 0...len(row)-1. 

// union-find method:
// https://leetcode.com/problems/couples-holding-hands/discuss/117520/Java-union-find-easy-to-understand-5-ms

func minSwapsCouples(row []int) int {
	count := 0
	// we can always swap n-1 times to make all couple together, assume we have a situation
	// (i, j) (k1, kk1) (k2, kk2), where (i, j) is not couple but (i, k1), (j, k2) are couples.
	// we can swap to (i, k1) (j, kk1) or (k2, j) (i, kk2). which is better? 
	// if (j, kk1) or (i, kk2) is couple, that's better move, which we can get 2 couples by only 1 swap.
    for i:=0; i<len(row); i+=2 {
    	j:=i+1
    	if row[i]/2 != row[j]/2 {
    		for k:=i+2; k<len(row); k++ {
    			kk := (1-k%2)+k
    			if row[k]/2==row[i]/2 {  // (i, j) (k, kk) to (i, k) (j kk), if (j kk) is also a couple, that's good
    				if row[j]/2==row[kk]/2 {
    					row[j], row[k] = row[k], row[j]
    					break
    				}
    			}
    			if row[k]/2==row[j]/2 {
    				row[i], row[k] = row[k], row[i]
    			} 
    		}
    		count++
    	}
    }
    return count
}

func main() {
	fmt.Println(minSwapsCouples([]int{30,31,38,37,2,26,34,15,14,13,27,21,6,10,11,24,36,0,20,5,35,7,8,17,18,25,4,39,3,1,32,29,23,16,12,19,33,22,9,28}), 16)
	fmt.Println(minSwapsCouples([]int{11,6,9,17,12,7,13,27,22,19,16,4,8,5,28,10,24,1,23,14,29,18,15,0,26,3,25,2,21,20}), 12)
	for i:=0; i<5; i++ {
		for _, n := range []int{10,20,30,40,50,60} {
			a := rand.Perm(n)
			fmt.Println(minSwapsCouples(a))
		}
	}
}
package main

import "fmt"

// https://leetcode.com/problems/count-triplets-that-can-form-two-arrays-of-equal-xor/

// Given an array of integers arr.
// We want to select three indices i, j and k where (0 <= i < j <= k < arr.length).
// Let's define a and b as follows:
//   a = arr[i] ^ arr[i + 1] ^ ... ^ arr[j - 1]
//   b = arr[j] ^ arr[j + 1] ^ ... ^ arr[k]
// Note that ^ denotes the bitwise-xor operation.
// Return the number of triplets (i, j and k) Where a == b.
// Example 1:
//   Input: arr = [2,3,1,6,7]
//   Output: 4
//   Explanation: The triplets are (0,1,2), (0,2,2), (2,3,4) and (2,4,4)
// Example 2:
//   Input: arr = [1,1,1,1,1]
//   Output: 10
// Example 3:
//   Input: arr = [2,3]
//   Output: 0
// Example 4:
//   Input: arr = [1,3,5,7,9]
//   Output: 3
// Example 5:
//   Input: arr = [7,11,12,9,5,2,7,17,22]
//   Output: 8
// Constraints:
//   1 <= arr.length <= 300
//   1 <= arr[i] <= 10^8

// NOTE: for problems where there is some accumulation of array, try prefix array.

// O(N) solution: a==b means a^b=0, means arr[i]^...^arr[j-1]^arr[j]...^a[k]=0
// the problem equals to find (i, k) pairs make arr[i]^...^arr[k]=0
// we construct a prefix array by, prefix[0]=arr[0], prefix[k]=arr[0]^...^arr[k-1],
// then prefix[i-1]^arr[i]^...^arr[k] = prefix[i-1]^0 = prefix[k], means we need find
// pairs (i, k) that makes prefix[i]=prefix[k], then we can add k-i-1 to count.
// we can get prefix count int a map, that makes a O(N) solution.
func countTriplets(arr []int) int {
	// O(N^3)
	count := 0
    for i:=0; i<len(arr)-1; i++ {
    	aik := arr[i]
		for k:=i+1; k<len(arr); k++ {
			aik = aik^arr[k]
			aij := 0
			ajk := aik
			for j:=i+1; j<=k; j++ {
				aij = aij^arr[j-1]
				ajk = aik ^ aij
				if aij == ajk {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countTriplets([]int{2,3,1,6,7}), 4)
	fmt.Println(countTriplets([]int{1,1,1,1,1}), 10)
	fmt.Println(countTriplets([]int{2,3}), 0)
	fmt.Println(countTriplets([]int{1,3,5,7,9}), 3)
	fmt.Println(countTriplets([]int{7,11,12,9,5,2,7,17,22}), 8)
}
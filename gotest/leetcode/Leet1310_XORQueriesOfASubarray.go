package main

import "fmt"

// https://leetcode.com/problems/xor-queries-of-a-subarray/

// Given the array arr of positive integers and the array queries where queries[i] = [Li, Ri], 
// for each query i compute the XOR of elements from Li to Ri 
// (that is, arr[Li] xor arr[Li+1] xor ... xor arr[Ri] ). 
// Return an array containing the result for the given queries.
// Example 1: 
//   Input: arr = [1,3,4,8], queries = [[0,1],[1,2],[0,3],[3,3]]
//   Output: [2,7,14,8] 
//   Explanation: 
//     The binary representation of the elements in the array are:
//     1 = 0001 
//     3 = 0011 
//     4 = 0100 
//     8 = 1000 
//     The XOR values for queries are:
//     [0,1] = 1 xor 3 = 2 
//     [1,2] = 3 xor 4 = 7 
//     [0,3] = 1 xor 3 xor 4 xor 8 = 14 
//     [3,3] = 8
// Example 2: 
//   Input: arr = [4,8,2,10], queries = [[2,3],[1,3],[0,0],[0,3]]
//   Output: [8,0,4,4]
// Constraints: 
//   1 <= arr.length <= 3 * 10^4 
//   1 <= arr[i] <= 10^9 
//   1 <= queries.length <= 3 * 10^4 
//   queries[i].length == 2 
//   0 <= queries[i][0] <= queries[i][1] < arr.length 

// the subarray xor is determined by the number of 1s in each bit. 
// if odd 1s, it's 1; if even 1s, it's 0. 
func xorQueries(arr []int, queries [][]int) []int {
	prefix := make([]int, len(arr)+1)
	for i := range arr {
		prefix[i+1] = prefix[i] ^ arr[i]   // prefix[i] = xor(arr[0...i-1])
	}
	ans := make([]int, len(queries))
	for i := range queries {        // xor(arr[i..j]) = prefix[j+1] ^ prefix[i]
		ans[i] = prefix[queries[i][1]+1] ^ prefix[queries[i][0]]
	}
	return ans
}

func main() {
	fmt.Println(xorQueries([]int{1,3,4,8 }, [][]int{{0,1},{1,2},{0,3},{3,3}}), []int{2,7,14,8})
	fmt.Println(xorQueries([]int{4,8,2,10}, [][]int{{2,3},{1,3},{0,0},{0,3}}), []int{8,0,4,4})
}

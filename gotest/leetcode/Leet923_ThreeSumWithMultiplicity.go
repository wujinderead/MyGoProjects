package main

import "fmt"

// https://leetcode.com/problems/3sum-with-multiplicity/

// Given an integer array A, and an integer target, return the number of tuples
// i, j, k such that i < j < k and A[i] + A[j] + A[k] == target.
// As the answer can be very large, return it modulo 10^9 + 7.
// Example 1:
//   Input: A = [1,1,2,2,3,3,4,4,5,5], target = 8
//   Output: 20
//   Explanation:
//     Enumerating by the values (A[i], A[j], A[k]):
//       (1, 2, 5) occurs 8 times;
//       (1, 3, 4) occurs 8 times;
//       (2, 2, 4) occurs 2 times;
//       (2, 3, 3) occurs 2 times.
// Example 2:
//   Input: A = [1,1,2,2,2,2], target = 5
//   Output: 12
//   Explanation:
//     A[i] = 1, A[j] = A[k] = 2 occurs 12 times:
//     We choose one 1 from [1,1] in 2 ways,
//     and two 2s from [2,2,2,2] in 6 ways.
// Note:
//   3 <= A.length <= 3000
//   0 <= A[i] <= 100
//   0 <= target <= 300

func threeSumMulti(A []int, target int) int {
    // count the occurrence of each number
	count := make([]int, 101)
	for _, v := range A {
		count[v] += 1
	}
	allcount := 0
	for i:=0; i<len(count); i++ {
		for j:=i; j<len(count); j++ {
			k := target-i-j
			if k<0 || k>100 {
				continue
			}
			if i<j && j<k {
				allcount += count[i]*count[j]*count[k]
			}
			if i==j && j<k {
				allcount += count[i]*(count[i]-1)/2*count[k]
			} 
			if i<j && j==k {
				allcount += count[i]*count[j]*(count[j]-1)/2
			}
			if i==j && j==k {
				allcount += count[i]*(count[i]-1)*(count[i]-2)/6
			}
		}
		allcount = allcount % int(1e9+7)
	}
	return allcount
}

func main() {
	fmt.Println(threeSumMulti([]int{1,1,2,2,3,3,4,4,5,5}, 8))
	fmt.Println(threeSumMulti([]int{1,1,2,2,2,2}, 5))
	fmt.Println(threeSumMulti([]int{16,51,36,29,84,80,46,97,84,16}, 171))
}
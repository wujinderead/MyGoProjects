package main

import (
    "fmt"
)

// https://leetcode.com/problems/k-inverse-pairs-array/

// Given two integers n and k, find how many different arrays consist of numbers from 1 to n such that 
// there are exactly k inverse pairs. We define an inverse pair as following: For ith and jth element 
// in the array, if i < j and a[i] > a[j] then it's an inverse pair; Otherwise, it's not.
// Since the answer may be very large, the answer should be modulo 109 + 7.
// Example 1:
//   Input: n = 3, k = 0
//   Output: 1
//   Explanation: 
//     Only the array [1,2,3] which consists of numbers from 1 to 3 has exactly 0 inverse pair.
// Example 2:
//   Input: n = 3, k = 1
//   Output: 2
//   Explanation: 
//     The array [1,3,2] and [2,1,3] have exactly 1 inverse pair.
// Note:
//   The integer n is in the range [1, 1000] and k is in the range [0, 1000].

// dp(i,k) = dp(i-1,k)+...dp(i-1,k-i+1)
// dp(i,k+1) = dp(i-1,k+1)+...+dp(i-1,k-i+2)
// we can deduct it into dp(i,k+1) = dp(i,k)+dp(i-1,k+1)-dp(i-1,k+1-i)

func kInversePairs(n int, k int) int {
    if k>n*(n-1)/2 {  // k won't be larger than n*(n-1)/2
    	return 0
    }
    if k>n*(n-1)/4 {
    	k = n*(n-1)/2-k   // for i+j==n(n-1)/2, f(n,i)==f(n,j) 
    }
    old, new := make([]int, k+1), make([]int, k+1)  // f(n-1, k), f(n, k)
    old[0] = 1
    for i:=2; i<=n; i++ {
    	new[0] = 1
    	for j:=1; j<=k && j<=i*(i-1)/2; j++ {
    		new[j] = new[j-1]+old[j]   // deduct here to avoid a loop
    		if j>=i {
    			new[j] -= old[j-i]
    		}
    		new[j] = (new[j]+1e9+7) % (1e9+7)
     	}
     	old, new = new, old
    }
    return old[k]
}

func main() {
	for i:=1; i<=5; i++ {
		for j:=0; j<=i*(i-1)/2+1; j++ {
			fmt.Printf("n=%d, k=%d, f(n,k)=%d\n", i, j, kInversePairs(i, j))
		}
	}
	fmt.Println(kInversePairs(500, 200), 961241146)
	fmt.Println(kInversePairs(500, 999), 144363604)
}
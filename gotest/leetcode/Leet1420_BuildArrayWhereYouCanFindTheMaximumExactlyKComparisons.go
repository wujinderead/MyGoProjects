package main

import "fmt"

// https://leetcode.com/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/
// Given three integers n, m and k. Consider the following algorithm to find the 
// maximum element of an array of positive integers: 
//   max_val = -1
//   max_ind = -1
//   cost = 0
//   n = len(arr)
//   for i:=0; i<n; i++ {
//	    if max_val < arr[i] {
//	    	max_val = arr[i]
//	    	max_ind = i
//	    	cost++
//	    }
//	 }
//	 return max_ind
// You should build the array arr which has the following properties: 
//   arr has exactly n integers. 
//   1 <= arr[i] <= m where (0 <= i < n). 
//   After applying the mentioned algorithm to arr, the value search_cost is equal to k. 
// Return the number of ways to build the array arr under the mentioned conditions. 
// As the answer may grow large, the answer must be computed modulo 10^9 + 7. 
// Example 1: 
//   Input: n = 2, m = 3, k = 1
//   Output: 6
//   Explanation: The possible arrays are [1, 1], [2, 1], [2, 2], [3, 1], [3, 2] [3, 3]
// Example 2: 
//   Input: n = 5, m = 2, k = 3
//   Output: 0
//   Explanation: There are no possible arrays that satisify the mentioned conditions.
// Example 3:  
//   Input: n = 9, m = 1, k = 1
//   Output: 1
//   Explanation: The only possible array is [1, 1, 1, 1, 1, 1, 1, 1, 1]
// Example 4: 
//   Input: n = 50, m = 100, k = 25
//   Output: 34549172
//   Explanation: Don't forget to compute the answer modulo 1000000007
// Example 5: 
//   Input: n = 37, m = 17, k = 7
//   Output: 418930126
// Constraints: 
//   1 <= n <= 50 
//   1 <= m <= 100 
//   0 <= k <= n 

// consider the array start with 0 and we need to populate another n numbers.
// let dp(i, p, k) be the number of ways to populate arr[i:] with previous max(arr[0...i-1])=p and cost k.
// if we set 1<=arr[i]<=p, we have no increase, so we need dp(i+1, p, k)
// if we set arr[i] = p+1, we need dp(i+1, p+1, k-1) ...
// so dp(i, p, k) = p*dp(i+1, p, k) + dp(i+1, p+1, k-1) + dp(i+1, p+2, k-1) + ... + dp(i+1, m, k-1)
// time O(mnkm), space O(mnk)

// for this problem, arr[1] always make a increase, so we let arr[1]=1...m, and we compute dp(2, 1...m, k-1)
// the answer is the sum of dp(2, 1...m, k-1)
func numOfArrays(n int, m int, k int) int {
    if k==0 {
    	return 0
    }
    dp := make([][][]int, n+1)
    for i := range dp {
    	dp[i] = make([][]int, m+1)
    	for j:=range dp[i] {
    		dp[i][j] = make([]int, k)
    	}
    }
    
    // initialize i==n, i.e, the situation for last number 
    for p:=1; p<=m; p++ {
    	dp[n][p][0] = p                 // prev max = p, we want no increase
    	if k>1 {
    		dp[n][p][1] = m-p           // prev max = p, we want 1 increase
    	}
    }

    // dp
    for i:=n-1; i>=2; i-- {
    	for p:=1; p<=m; p++ {
    		dp[i][p][0] = (p*dp[i+1][p][0]) % (1e9+7)         // kk=0
    		for kk:=1; kk<k && kk<=m-p && kk<=n-i+1; kk++ {   // to make kk early stop
    			dp[i][p][kk] = p*dp[i+1][p][kk]
    			for j:=p+1; j<=m-kk+1; j++ {                  // to make j stop early, j+kk-1<=m, otherwise 0
    				dp[i][p][kk] += dp[i+1][j][kk-1]
    			}
    			dp[i][p][kk] %= 1e9+7
    		}
    	}
    }
    ans := 0
    for i:=1; i<=m; i++ {
    	ans += dp[2][i][k-1]
    }
    return ans%(1e9+7)
}

func main() {
	fmt.Println(numOfArrays(2,3,1), 6)
	fmt.Println(numOfArrays(2,3,2), 3)
	fmt.Println(numOfArrays(5,2,3), 0)
	fmt.Println(numOfArrays(9,1,1), 1)
	fmt.Println(numOfArrays(50,100,25), 34549172)
	fmt.Println(numOfArrays(37,17,7), 418930126)
}
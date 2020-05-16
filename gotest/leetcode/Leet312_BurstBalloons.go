package main

import "fmt"

// https://leetcode.com/problems/burst-balloons/

// Given n balloons, indexed from 0 to n-1. Each balloon is painted with a number
// on it represented by array nums. You are asked to burst all the balloons. If the 
// you burst balloon i you will get nums[left] * nums[i] * nums[right] coins. Here 
// left and right are adjacent indices of i. After the burst, the left and right 
// then becomes adjacent. Find the maximum coins you can collect by bursting the 
// balloons wisely. 
// Note: 
//   You may imagine nums[-1] = nums[n] = 1. They are not real therefore can not burst. 
//   0 ≤ n ≤ 500, 0 ≤ nums[i] ≤ 100 
// Example: 
//   Input: [3,1,5,8]
//   Output: 167 
//   Explanation: nums = [3,1,5,8] --> [3,5,8] -->   [3,8]   -->  [8]  --> []
//                coins =  3*1*5      +  3*5*8    +  1*3*8      + 1*8*1   = 167

// another genius idea:
// https://leetcode.com/problems/burst-balloons/discuss/76241/Another-way-to-think-of-this-problem-(Matrix-chain-multiplication)
// For example, given [3,5,8] and bursting 5, the number of coins you get is the number of 
// scalar multiplications you need to do to multiply two matrices A[3*5] and B[5*8]. So in this example, 
// the original problem is actually the same as given a matrix chain A[1*3]*B[3*5]*C[5*8]*D[8*1], 
// fully parenthesize it so that the total number of scalar multiplications is maximized, 
// although the orignal matrix-chain multiplication problem in the book asks to minimize it. 
// Then you can see it clearly as a classical DP problem.

// let p[i][j] be the max profit when we only burst nums[i...j]. 
// for each k (i<=k<=j), if we burst ballon nums[k] last, 
// then we can burst nums[i...k-1] and nums[k+1...j] separately, they won't interfere each other.
// finally we burst k, we got nums[i-1]*nums[k]*nums[j+1] profit. we want the k that make profit maximal.
//    
//     i-1  (i ... k-1)  k  (k+1 ...  j)  j+1
// 
// i.e., p[i][j] = max(nums[i-1]*nums[k]*nums[j+1] + p[i][k-1] + p[k+1][j]), for i<=k<=j
// if i==k, ignore p[i][k-1]; if j==k, ignore p[k+1][j]; if i==j==k, ignore both.
func maxCoins(nums []int) int {
	// make a new array num
	// num index:   0     1        2     ...     n        n+1
	// num value:   1  nums[0]  nums[1]       nums[n-1]    1         ,   n=len(nums)
	n := len(nums)
    num := make([]int, len(nums)+2)
    copy(num[1:], nums)
    num[0] = 1
    num[n+1] = 1
    
    dp := make([][]int, n+2)
    for i:= range dp {
    	dp[i] = make([]int, n+2)
    }

    for leng:=0; leng<n; leng++ {   // compute num[i][i+leng], i.e., update diagonally
    	for i:=1; i+leng<=n; i++ {
    		j := i+leng
    		for k:=i; k<=j; k++ {
    			// if i==k, ignore dp[i][k-1]; if j==k, ignore dp[k+1][j]; if i==j==k, ignore both.
    			// we add here because dp[r][c] happens to be 0 when r>c.
    			dp[i][j] = max(dp[i][j], num[i-1]*num[k]*num[j+1]+dp[i][k-1]+dp[k+1][j])
    		}
    	}
    }
    return dp[1][n]
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxCoins([]int{3,1,5,8}))
	fmt.Println(maxCoins([]int{3,8}))
	fmt.Println(maxCoins([]int{8}))
	fmt.Println(maxCoins([]int{8,3}))
	fmt.Println(maxCoins([]int{5,3,8}))
}
package main

import (
    "fmt"
)

// https://leetcode.com/problems/number-of-dice-rolls-with-target-sum/

// You have d dice, and each die has f faces numbered 1, 2, ..., f. 
// Return the number of possible ways (out of f^d total ways) modulo 10^9 + 7 to 
// roll the dice so the sum of the face up numbers equals target. 
// Example 1: 
//   Input: d = 1, f = 6, target = 3
//   Output: 1
//   Explanation: 
//     You throw one die with 6 faces.  There is only one way to get a sum of 3.
// Example 2: 
//   Input: d = 2, f = 6, target = 7
//   Output: 6
//   Explanation: 
//     You throw two dice, each with 6 faces.  There are 6 ways to get a sum of 7:
//     1+6, 2+5, 3+4, 4+3, 5+2, 6+1.
// Example 3: 
//   Input: d = 2, f = 5, target = 10
//   Output: 1
//   Explanation: 
//     You throw two dice, each with 5 faces.  There is only one way to get a sum of 10: 5+5.
// Example 4: 
//   Input: d = 1, f = 2, target = 3
//   Output: 0
//   Explanation: 
//     You throw one die with 2 faces.  There is no way to get a sum of 3.
// Example 5: 
//   Input: d = 30, f = 30, target = 500
//   Output: 222616187
//   Explanation: 
//     The answer must be returned modulo 10^9 + 7.
// Constraints: 
//   1 <= d, f <= 30 
//   1 <= target <= 1000 

func numRollsToTarget(d int, f int, target int) int {
    dp := [31][1001]int{}  // dp[i][j] means the number of ways to get sum j with i dices
    for i:=1; i<=f; i++ {
    	dp[1][i] = 1
    }
    for i:=2; i<=d; i++ {
    	for j:=1; j<=target; j++ {
    		if i>j {
    			continue
    		}
    		for k:=1; k<j && k<=f; k++ {        // to sum j with i dices
    			dp[i][j] += dp[i-1][j-k]        // i-th dice can be 1<=k<=min(j-1, f), use i-1 dices to get j-k 
    		}
    		dp[i][j] = dp[i][j] % (1e9+7)
    	}
    }
    return dp[d][target] 
}

func main() {
	fmt.Println(numRollsToTarget(1,6,3), 1)
	fmt.Println(numRollsToTarget(2,6,7), 6)
	fmt.Println(numRollsToTarget(2,5,10), 1)
	fmt.Println(numRollsToTarget(1,2,3), 0)
	fmt.Println(numRollsToTarget(30,30,500), 222616187)
}
package main

import "fmt"

// https://leetcode.com/problems/stone-game-v/

// There are several stones arranged in a row, and each stone has an associated value which is an integer 
// given in the array stoneValue.
// In each round of the game, Alice divides the row into two non-empty rows (i.e. left row and right row), 
// then Bob calculates the value of each row which is the sum of the values of all the stones in this row. 
// Bob throws away the row which has the maximum value, and Alice's score increases by the value of the 
// remaining row. If the value of the two rows are equal, Bob lets Alice decide which row will be thrown away.
// The next round starts with the remaining row.
// The game ends when there is only one stone remaining. Alice's is initially zero.
// Return the maximum score that Alice can obtain.
// Example 1:
//   Input: stoneValue = [6,2,3,4,5,5]
//   Output: 18
//   Explanation: In the first round, Alice divides the row to [6,2,3], [4,5,5]. 
//     The left row has the value 11 and the right row has value 14. 
//     Bob throws away the right row and Alice's score is now 11.
//     In the second round Alice divides the row to [6], [2,3]. 
//     This time Bob throws away the left row and Alice's score becomes 16 (11 + 5).
//     The last round Alice has only one choice to divide the row which is [2], [3]. 
//     Bob throws away the right row and Alice's score is now 18 (16 + 2). 
//     The game ends because only one stone is remaining in the row.
// Example 2:
//   Input: stoneValue = [7,7,7,7,7,7,7]
//   Output: 28
// Example 3:
//   Input: stoneValue = [4]
//   Output: 0
// Constraints:
//   1 <= stoneValue.length <= 500
//   1 <= stoneValue[i] <= 10^6

func stoneGameV(s []int) int {
	if len(s)==1 {
		return 0
	}
	prefix := make([]int, len(s)+1)
	for i:=1; i<=len(s); i++ {
		prefix[i] = prefix[i-1]+s[i-1]
	}
	// let dp(i, j) be the answer for s[i...j], then, for i<=k<=j,
	// if sum(s[i...k]) < sum(s[k+1...j]), candidate result is dp(i, k) + sum(s[i...k])
	// else, candidate result is dp(k+1, j) + sum(s[k+1...j])
	dp := [500][500]int{}
	for diff:=1; diff<len(s); diff++ {
		for i:=0; i+diff<len(s); i++ {
			j := i+diff     // calculate dp(i, j)
			for k:=i; k<j; k++ {
				left, right := prefix[k+1]-prefix[i], prefix[j+1]-prefix[k+1]
				if left <= right {    // left part smaller, use dp(left part) + left part sum
					dp[i][j] = max(dp[i][j], left + dp[i][k])					
				} 
				// use both left<=right and left>=right, we cover the equal situation
				if left >= right {   // use right part
					dp[i][j] = max(dp[i][j], right + dp[k+1][j])
				}
			}
		}
	}
	/*for i:=0; i<len(s); i++ {
		for j:=i; j<len(s); j++ {
			fmt.Println(s[i:j+1], dp[i][j])
		}
	}*/
	return dp[0][len(s)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	for _, v := range []struct{s []int; ans int} {
		{[]int{6,2,3,4,5,5}, 18},
		{[]int{7,7,7,7,7,7,7}, 28},
		{[]int{4}, 0},
		{[]int{1,1,1,1,2,5}, 7},
	} {
		fmt.Println(stoneGameV(v.s), v.ans)
	}
}

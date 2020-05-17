package main

import "fmt"

// https://leetcode.com/problems/form-largest-integer-with-digits-that-add-up-to-target/

// Given an array of integers cost and an integer target. Return the maximum integer
// you can paint under the following rules:
//   The cost of painting a digit (i+1) is given by cost[i] (0 indexed).
//   The total cost used must be equal to target.
//   Integer does not have digits 0.
// Since the answer may be too large, return it as string.
// If there is no way to paint any integer given the condition, return "0".
// Example 1:
//   Input: cost = [4,3,2,5,6,7,2,5,5], target = 9
//   Output: "7772"
//   Explanation:  The cost to paint the digit '7' is 2, and the digit '2' is 3. Then
//     cost("7772") = 2*3+ 3*1 = 9. You could also paint "997", but "7772" is the
//     largest number.
//     Digit    cost
//       1  ->   4
//       2  ->   3
//       3  ->   2
//       4  ->   5
//       5  ->   6
//       6  ->   7
//       7  ->   2
//       8  ->   5
//       9  ->   5
// Example 2:
//   Input: cost = [7,6,5,5,5,6,8,7,8], target = 12
//   Output: "85"
//   Explanation: The cost to paint the digit '8' is 7, and the digit '5' is 5.
//     Then cost("85") = 7 + 5 = 12.
// Example 3:
//   Input: cost = [2,4,6,2,4,6,4,4,4], target = 5
//   Output: "0"
//   Explanation: It's not possible to paint any integer with total cost equal to target.
// Example 4:
//   Input: cost = [6,10,15,40,40,40,40,40,40], target = 47
//   Output: "32211"
// Constraints:
//   cost.length == 9
//   1 <= cost[i] <= 5000
//   1 <= target <= 5000

// much easier logic:  https://leetcode.com/problems/form-largest-integer-with-digits-that-add-up-to-target/discuss/635267/C%2B%2BJavaPython-Strict-O(Target)
// dp := make([]int, target + 1)           // dp
// for t:=1; t<=target; t++ {
// 	   dp[t] = -10000
//     for i := 0; i < 9; i++ {
//     if t >= cost[i] {
//         dp[t] = max(dp[t], 1 + dp[t - cost[i]])
//     }
// }
// for i:=8; i>=0; i-- {         // find max string
//     for target >= cost[i] && dp[target]==dp[target-cost[i]]+1 {
//         buf = append(buf, i+1+'0')
//         target -= cost[i]
//     }
// }
func largestNumber(cost []int, target int) string {
    // get all distinct cost in c. we want the the most digits in c that sums to target.
    // denote dp(i, j) be the most digits using c[:i] sums to j. we have options:
    // if use c[i], we got dp(i, j-c[i])+1
    // if not use c[i], we got dp(i-1, j), we want the larger.
    // base case dp(i, 0)=0; dp(0, j)=j/c[0] if j%c[0]==0

    // find distinct cost
    var c, d []int    // c: distinct cost, d: the corresponding digit
    outer: for i:=8; i>=0; i-- {
    	for _, v := range c {
    		if v==cost[i] {
    			continue outer
			}
		}
		c = append(c, cost[i])
		d = append(d, i+1)
	}
	for i, j := 0, len(c)-1; i<j; i,j = i+1, j-1 {   // let c[i] sorted by d[i] ascending
		d[i], d[j] = d[j], d[i]
		c[i], c[j] = c[j], c[i]
	}
	//fmt.Println(c)
	//fmt.Println(d)

	// find the maximal numbers
    dp := make([][]int, len(c))
    for i := range dp {
    	dp[i] = make([]int, target+1)
	}

	for j:=1; j<=target; j++ {
		if j%c[0]==0 {
			dp[0][j] = j/c[0]
		}
	}

	for i:=1; i<len(c); i++ {
		for j:=1; j<=target; j++ {
			dp[i][j] = dp[i-1][j]
			if (j-c[i]==0 || (j-c[i]>0 && dp[i][j-c[i]]>0)) && dp[i][j-c[i]]+1>dp[i-1][j] {
				dp[i][j] = dp[i][j-c[i]]+1
			}
		}
	}
	if dp[len(c)-1][target]==0 {   // can't form number
		return "0"
	}

	// backtracking dp to find the digits
	//fmt.Println(dp)
	buf := make([]byte, 0, dp[len(c)-1][target])
	i := len(c)-1
	j := target
	for j>0 {
		if (j-c[i]==0 || (j-c[i]>0 && dp[i][j-c[i]]>0)) && dp[i][j-c[i]]+1 == dp[i][j] {
			buf = append(buf, byte(d[i])+'0')
			j = j-c[i]
			continue
		}
		i--
	}
	return string(buf)
}

func main() {
	fmt.Println(largestNumber([]int{4,3,2,5,6,7,2,5,5}, 9))
	fmt.Println(largestNumber([]int{7,6,5,5,5,6,8,7,8}, 12))
	fmt.Println(largestNumber([]int{2,4,6,2,4,6,4,4,4}, 5))
	fmt.Println(largestNumber([]int{6,10,15,40,40,40,40,40,40}, 47))
	fmt.Println(largestNumber([]int{2,2,2,2,2,2,2,2,2}, 8))
	fmt.Println(largestNumber([]int{2,2,2,2,2,2,2,2,2}, 9))
}
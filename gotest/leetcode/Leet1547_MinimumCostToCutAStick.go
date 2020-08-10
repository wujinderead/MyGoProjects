package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-cost-to-cut-a-stick

// Given a wooden stick of length n units. The stick is labelled from 0 to n.
// For example, a stick of length 6 is labelled as follows:
//      |---|---|---|---|---|---|
//      0   1   2   3   4   5   6
// Given an integer array cuts where cuts[i] denotes a position you should perform a cut at.
// You should perform the cuts in order, you can change the order of the cuts as you wish.
// The cost of one cut is the length of the stick to be cut, the total cost is the sum of costs of all cuts.
// When you cut a stick, it will be split into two smaller sticks (i.e. the sum of their lengths is the
// length of the stick before the cut). Please refer to the first example for a better explanation.
// Return the minimum total cost of the cuts.
// Example 1:
//   Input: n = 7, cuts = [1,3,4,5]
//   Output: 16
//      .---x---.---x---x---x---.---.              cut at 3, cost 7
//      0   1   2   3   4   5   6   7
//      .---x---.---.   .---x---x---.---.          cut at 5, cost 4
//      0   1   2   3   3   4   5   6   7
//      .---x---.---.   .---x---.  .---.---.       cut at 1 and 4, cost 3+2, total cost 16
//      0   1   2   3   3   4   5  5   6   7
//   Explanation: Using cuts order = [1, 3, 4, 5] as in the input leads to the following scenario:
//     The first cut is done to a rod of length 7 so the cost is 7.
//     The second cut is done to a rod of length 6 (i.e. the second part of the first cut),
//     the third is done to a rod of length 4 and the last cut is to a rod of length 3.
//     The total cost is 7 + 6 + 4 + 3 = 20. Rearranging the cuts to be [3, 5, 1, 4] for example will
//     lead to a scenario with total cost = 16 (as shown in the example photo 7 + 4 + 3 + 2 = 16).
// Example 2:
//   Input: n = 9, cuts = [5,6,1,4,2]
//   Output: 22
//   Explanation: If you try the given cuts ordering the cost will be 25.
//     There are much ordering with total cost <= 25, for example,
//     the order [4, 6, 5, 2, 1] has total cost = 22 which is the minimum possible.
// Constraints:
//   2 <= n <= 10^6
//   1 <= cuts.length <= min(n - 1, 100)
//   1 <= cuts[i] <= n - 1
//   All the integers in cuts array are distinct.

// add 0 and n to cuts then sort it.
// for stick between cuts[i] and cuts[j], we can increase the difference between i and j, 
// so we can get the min cost to cut stick[cuts[i]...cuts[j]], denote as dp[i][j].
// e.g., dp[i][i+1]=0, dp[i][i+2] = dp[i][i+1]+dp[i+1][i+2] + (cuts[i+2]-cuts[i]),
// dp[i][i+3] = min(dp[i][i+1]+dp[i+1][i+3], dp[i][i+2]+dp[i+2][i+3]) + (cuts[i+3]-cuts[i]),
func minCost(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Sort(sort.IntSlice(cuts))
	dp := [102][102]int{}
    for diff:=2; diff<len(cuts); diff++ {   // diff=1, no need to cut, dp[i][i+1]=0
		for i:=0; i+diff<len(cuts); i++ {
			j := i+diff
			min := 0x7ffffff
			for k:=i+1; k<j; k++ {
				if dp[i][k] + dp[k][j] < min {
					min = dp[i][k] + dp[k][j]
				}
			}
			dp[i][j] = min + cuts[j]-cuts[i]   // plus the stick length
		}
	}
	return dp[0][len(cuts)-1]
}

func main() {
	for _, v := range []struct{n int; arr []int; ans int} {
		{7, []int{1,3,4,5}, 16},
		{9, []int{5,6,1,4,2}, 22},
		{9, []int{2}, 9},
		{9, []int{2,4}, 13},
		{100000, []int{36372,70799,65123,12766,35304,77919,45836,38952,26406,15331,61885,22033,33693,3007,
			89499,69639,58935,78415,29960,67925,76232,33978,91178,43973,10567,2792,16775,17081,58346,96249,
			36341,54042,62711,97855,3404,20405,49875,33412,78792,52749,83969,8909,79070,85103,21756,7443,
			81559,18666,92279,70742,89373,87830,83380,97890,63133,60816,17784,78533,83102,78449,42040,43106,
			86522,97259,997,59157,71464,817,98278,32954,27949,80126,4638,46315,26841,42829,44907,24139,6058,
			40979,87478,87536,56204,29867,43143,90732,20853,99087,1684,46224,12628,56798,54803,89658,9302,
			32078,22643,34348,90348,23212}, 628770},
	} {
		fmt.Println(minCost(v.n, v.arr), v.ans)
	}
}
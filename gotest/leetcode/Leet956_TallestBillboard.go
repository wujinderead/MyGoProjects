package main

import "fmt"

// https://leetcode.com/problems/tallest-billboard/

// You are installing a billboard and want it to have the largest height. The billboard 
// will have two steel supports, one on each side. Each steel support must be an equal height. 
// You have a collection of rods which can be welded together. For example, if you have rods 
// of lengths 1, 2, and 3, you can weld them together to make a support of length 6. 
// Return the largest possible height of your billboard installation. 
// If you cannot support the billboard, return 0. 
// Example 1: 
//   Input: [1,2,3,6]
//   Output: 6
//   Explanation: We have two disjoint subsets {1,2,3} and {6}, which have the same sum = 6.
// Example 2: 
//   Input: [1,2,3,4,5,6]
//   Output: 10
//   Explanation: We have two disjoint subsets {2,3,5} and {4,6}, which have the same sum = 10.
// Example 3: 
//   Input: [1,2]
//   Output: 0
//   Explanation: The billboard cannot be supported, so we return 0.
// Note: 
//   0 <= rods.length <= 20 
//   1 <= rods[i] <= 1000 
//   The sum of rods is at most 5000. 

func tallestBillboard(rods []int) int {
	if len(rods)<2 {
		return 0
	}

	// think the problem as assigning coefficient (0, 1, -1) to each number and get the sum, 
	// when the sum is 0, we can get a solution by the positive sums. e.g., for [1,2,3,6], 
	// we can assign 1+2+(-1)*3=0, we got a solution 3; 1+2+3+(-1)*6=0, we got a solution 6.
	// so we track not only the sum, but also the maximal positive sum.
	// let dp(i, j) be if first i numbers can sum to j, the we check 
	// dp(i-1, j-rods[i]), dp(i-1, j+rods[i]), dp(i-1, j). to avoid negative index, we add all numbers 5000.
	oldsum, newsum := make([]bool, 10001), make([]bool, 10001)    // all sum
	oldmax, newmax := make([]int, 10001), make([]int, 10001)      // positive sum
	oldsum[0+5000] = true  // wee can always get sum 0
	for i := range rods {
		for j:=0; j<=10000; j++ {
			newsum[j] = oldsum[j]
			newmax[j] = oldmax[j]
			if j-rods[i]>=0 {
				newsum[j] = newsum[j] || oldsum[j-rods[i]]
				if oldsum[j-rods[i]] {
					newmax[j] = max(newmax[j], oldmax[j-rods[i]])
				}
			}
			if j+rods[i]<=10000 {
				newsum[j] = newsum[j] || oldsum[j+rods[i]]
				if oldsum[j+rods[i]] {
					newmax[j] = max(newmax[j], oldmax[j+rods[i]] + rods[i])
				}
			}
		}
		oldsum, newsum = newsum, oldsum
		oldmax, newmax = newmax, oldmax
	}
    return oldmax[5000]
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

// time complexity O(n*sum*sum), got TLE in leetcode.
func tallestBillboard1(rods []int) int {
	if len(rods)<2 {
		return 0
	}
	T := 0
	for i := range rods {
		T += rods[i]
	}
	T = T/2   // the max support we can get won't exceed T

	// let dp(k, m, n) be if we can use first k rods to make 2 supports with height m and n, then:
	//   if we use k-th rod in first support, we check dp(k-1, m-rods[k], n)
	//   if we use k-th rod in second support, we check dp(k-1, m, n-rods[k])
	//   if we don't use in either support, we check dp(k-1, m, n),
	// so, dp(k, m, n) = dp(k-1, m-rods[k], n) || dp(k-1, m, n-rods[k]) || dp(k-1, m, n).
	// we want the max v that makes dp(all rods, v, v) = true
	old, new := make([][]bool, T+1), make([][]bool, T+1)
	for i := range old {
		old[i] = make([]bool, T+1)
		new[i] = make([]bool, T+1)
	}
	old[0][0] = true               // dp(x, 0, 0) = true
	if rods[0] <= T {
		old[0][rods[0]] = true     // dp(0, rods[0], 0) = dp(0, 0, rods[0]) = true
		old[rods[0]][0] = true
	}

	for k:=1; k<len(rods); k++ {
		for i:=0; i<=T; i++ {         // dp(x, i, j) == dp(x, j, i), so we only compute i<=j
			for j:=0; j<=T; j++ {
				new[i][j] = old[i][j]
				if i-rods[k]>=0 {
					new[i][j] = new[i][j] || old[i-rods[k]][j]
				}
				if j-rods[k]>=0 {
					new[i][j] = new[i][j] || old[i][j-rods[k]]
				}
			}
		}
		old, new = new, old
	}
	for i:=T; i>0; i-- {
		if old[i][i] {
			return i
		}
	}
    return 0
}

func main() {
	fmt.Println(tallestBillboard([]int{1,2,3,6}), 6)
	fmt.Println(tallestBillboard([]int{1,2,3,4,5,6}), 10)
	fmt.Println(tallestBillboard([]int{1,2}), 0)
	fmt.Println(tallestBillboard([]int{3,10}), 0)
	fmt.Println(tallestBillboard([]int{142,178,178,143,133,139,117,153,144,162,160,147,136,149,163,160,130,157,159}))
	fmt.Println(tallestBillboard([]int{19,13,12,20,10,16,13,13,15,16,11,17,13,12,800,800,800,800,800,800}))

	fmt.Println(tallestBillboard1([]int{1,2,3,6}), 6)
	fmt.Println(tallestBillboard1([]int{1,2,3,4,5,6}), 10)
	fmt.Println(tallestBillboard1([]int{1,2}), 0)
	fmt.Println(tallestBillboard1([]int{3,10}), 0)
	fmt.Println(tallestBillboard1([]int{142,178,178,143,133,139,117,153,144,162,160,147,136,149,163,160,130,157,159}))
	fmt.Println(tallestBillboard1([]int{19,13,12,20,10,16,13,13,15,16,11,17,13,12,800,800,800,800,800,800}))
}
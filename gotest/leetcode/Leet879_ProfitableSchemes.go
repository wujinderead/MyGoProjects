package main

import (
    "fmt"
)

// https://leetcode.com/problems/profitable-schemes/

// There are G people in a gang, and a list of various crimes they could commit. 
// The i-th crime generates a profit[i] and requires group[i] gang members to participate. 
// If a gang member participates in one crime, that member can't participate in another crime. 
// Let's call a profitable scheme any subset of these crimes that generates at least P profit, 
// and the total number of gang members participating in that subset of crimes is at most G. 
// How many schemes can be chosen? Since the answer may be very large, return it modulo 10^9 + 7. 
// Example 1: 
//   Input: G = 5, P = 3, group = [2,2], profit = [2,3]
//   Output: 2
//   Explanation: 
//     To make a profit of at least 3, the gang could either commit crimes 0 and 1, or just crime 1.
//     In total, there are 2 schemes.
// Example 2: 
//   Input: G = 10, P = 5, group = [2,3,5], profit = [6,7,8]
//   Output: 7
//   Explanation: 
//     To make a profit of at least 5, the gang could commit any crimes, as long as they commit one.
//     There are 7 possible schemes: (0), (1), (2), (0,1), (0,2), (1,2), and (0,1,2).
// Note: 
//   1 <= G <= 100 
//   0 <= P <= 100 
//   1 <= group[i] <= 100 
//   0 <= profit[i] <= 100 
//   1 <= group.length = profit.length <= 100 

func profitableSchemes(G int, P int, group []int, profit []int) int {
    // let dp(i, j, k) be number of ways to use exact j persons with crime[:i] to get at least k profit, then, 
    // if we don't use crime[i], we have dp(i-1, j, k)
    // if we use crime[i], we have dp(i-1, j-group[i], k-profit[i])
    // finally we sum dp(all crimes, all g, P)
    old, new := [101][101]int{}, [101][101]int{}

    // initialize for profit 0
    old[0][0] = 1

    // dp
    for i:=0; i<len(group); i++ {          // for each crime
    	for k:=0; k<=P; k++ {
    		for j:=0; j<=G; j++ {
    			new[j][k] = old[j][k]
    			if j-group[i]>=0 {
    				t := k-profit[i]
    				if t<0 {
    					t = 0
    				}
    				new[j][k] += old[j-group[i]][t]
    				new[j][k] = new[j][k] % (1e9+7)
    			}
    		}
    	}
    	/*for j:=0; j<=G; j++ {
    		fmt.Println(new[j][:P+1])
    	}
    	fmt.Println()*/
    	old, new = new, old
    }
    sum := 0
    for j:=1; j<=G; j++ {
    	sum = (sum + old[j][P]) % (1e9+7)
    }
    return sum
}

func main() {
	fmt.Println(profitableSchemes(5, 3, []int{2,2}, []int{2,3}), 2)
	fmt.Println(profitableSchemes(10, 5, []int{2,3,5}, []int{6,7,8}), 7)
	fmt.Println(profitableSchemes(10, 1, 
		[]int{6,3,6,1,10,1,11,6,8,8,11,10,9,10,4,7,9,6,7,9,10,8,4,6,7,7,9,4,4,4,8,6,7,10,5,2,1,6,11,
			3,8,9,3,2,8,4,7,10,9,5,3,6,10,4,5,4,10,3,8,6,11,10,6,9,8,11,3,7,2,7,7,9,7,10,1,3,3,9,6,
			3,11,3,5,10,9,4,10,6,4,10,9,2,1,1,9,10,5,10,7,6}, 
		[]int{2,0,0,1,2,0,0,1,2,1,1,2,2,2,1,0,2,2,1,1,0,0,2,2,0,2,2,2,0,1,2,1,1,0,0,2,2,2,2,0,0,0,0,2,0,0,1,
			0,2,1,0,2,0,0,1,2,2,1,1,2,1,1,2,0,2,0,0,1,1,1,0,1,1,2,2,1,0,0,1,0,2,2,1,2,2,0,0,2,0,2,2,1,0,2,0,1,0,1,0,2},
	), 33940)
}
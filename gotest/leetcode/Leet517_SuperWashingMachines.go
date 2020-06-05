package main

import (
    "fmt"
)

// https://leetcode.com/problems/super-washing-machines/

// You have n super washing machines on a line. Initially, each washing machine has some dresses or is empty. 
// For each move, you could choose any m (1 ≤ m ≤ n) washing machines, and pass one dress of each washing 
// machine to one of its adjacent washing machines at the same time. 
// Given an integer array representing the number of dresses in each washing machine from left to right on 
// the line, you should find the minimum number of moves to make all the washing machines have the same number
// of dresses. If it is not possible to do it, return -1. 
// Example1
//   Input: [1,0,5]
//   Output: 3
//   Explanation: 
//     1st move:    1     0 <-- 5    =>    1     1     4
//     2nd move:    1 <-- 1 <-- 4    =>    2     1     3    
//     3rd move:    2     1 <-- 3    =>    2     2     2   
// Example2
//   Input: [0,3,0]
//   Output: 2
//   Explanation: 
//     1st move:    0 <-- 3     0    =>    1     2     0    
//     2nd move:    1     2 --> 0    =>    1     1     1     
// Example3
//   Input: [0,2,0]
//   Output: -1
//   Explanation: 
//     It's impossible to make all the three washing machines have the same number of dresses. 
// Note: 
//   The range of n is [1, 10000]. 
//   The range of dresses number in a super washing machine is [0, 1e5]. 

// for example, [0,0,11,5] the target is 4, subtract target is arr=[-4,-4,7,1].
// 1st -4 means we need 4 ops to get 4 from 2nd. arr=[0,-8,7,1].
// 2nd -8 means we need 8 ops to get 8 from 3rd. arr=[0,0,-1,1].
// 3rd -1 means we need 1 op to get 1 from 4th. however, this place is 11 we need 7 ops to give out 7. 
// so this cost is max(7,-1)=7. arr=[0,0,0,0]. also, the 4th cost is max(0,1)=1.
// so the total max(4,8,7,1)=4 
func findMinMoves(machines []int) int {
	sum := 0
	for _, v := range machines {
		sum += v
	}
	if sum%len(machines) != 0 {
		return -1
	}
	target := sum/len(machines)
	maxcost := 0
	cost := 0
	for _, v := range machines {
		cost += v-target   // v-target is "gain/lose"
		maxcost = max(maxcost, max(abs(cost), v-target))
	}
	return maxcost
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func abs(a int) int {
	if a<0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(findMinMoves([]int{1,0,5}))
	fmt.Println(findMinMoves([]int{0,3,0}))
	fmt.Println(findMinMoves([]int{0,2,0}))
	fmt.Println(findMinMoves([]int{0,0,11,5}))
	fmt.Println(findMinMoves([]int{0,0,11,4,0}))
	fmt.Println(findMinMoves([]int{4,0,0,4}))
}
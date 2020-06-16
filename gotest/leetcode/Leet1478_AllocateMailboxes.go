package main

import (
	"sort"
    "fmt"
)

// https://leetcode.com/problems/allocate-mailboxes/

// Given the array houses and an integer k. where houses[i] is the location of the ith house along a street, 
// your task is to allocate k mailboxes in the street. 
// Return the minimum total distance between each house and its nearest mailbox.
// The answer is guaranteed to fit in a 32-bit signed integer. 
// Example 1: 
//   Input: houses = [1,4,8,10,20], k = 3
//   Output: 5
//   Explanation: Allocate mailboxes in position 3, 9 and 20.
//     Minimum total distance from each houses to nearest mailboxes is |3-1| + |4-3| + |9-8| + |10-9| + |20-20| = 5 
// Example 2: 
//   Input: houses = [2,3,5,12,18], k = 2
//   Output: 9
//   Explanation: Allocate mailboxes in position 3 and 14.
//     Minimum total distance from each houses to nearest mailboxes is |2-3| + |3-3| + |5-3| + |12-14| + |18-14| = 9.
// Example 3: 
//   Input: houses = [7,4,6,1], k = 1
//   Output: 8
// Example 4: 
//   Input: houses = [3,6,14,10], k = 4
//   Output: 0
// Constraints: 
//   n == houses.length 
//   1 <= n <= 100 
//   1 <= houses[i] <= 10^4 
//   1 <= k <= n 
//   Array houses contain unique integers. 

func minDistance(houses []int, K int) int {
	if K==len(houses) {
		return 0
	}
	sort.Sort(sort.IntSlice(houses))
	onebox := make(map[[2]int]int)
	old, new := make([]int, len(houses)), make([]int, len(houses))

	// initialize for k=1
	for diff:=1; diff<len(houses); diff++ {
		for i:=0; i+diff<len(houses); i++ {
			j := i+diff
			// for diff=1, onebox[i+1][j-1] has no meaning, it happens to be 0
			onebox[[2]int{i,j}] = houses[j]-houses[i]+onebox[[2]int{i+1, j-1}]
			if i==0 {
				old[j] = onebox[[2]int{i,j}] // use 1 boxes in house[0...j]
			}
		}
	}

	// dp(k, j) is the total sum for house[0...j] with k boxes, then we find x that makes it minimal: 
	// dp(k, j) = dp(k-1, x) + cost[x+1][j]   // use k-1 boxes in house[0...x], use 1 box in house [x+1...j]
	for k:=2; k<=K; k++ {
		for j:=k-1; j<len(houses); j++ {   // k boxes for houses[0...j]
			new[j] = int(1e8)
			for x:=k-2; x<j; x++ {
				new[j] = min(new[j], old[x]+onebox[[2]int{x+1, j}])
			}
		}
		old, new = new, old
	}
	return old[len(houses)-1]
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minDistance([]int{1,4,8,10,20}, 3))
	fmt.Println(minDistance([]int{2,3,5,12,18}, 2))
	fmt.Println(minDistance([]int{7,4,6,1}, 1))
	fmt.Println(minDistance([]int{3,6,14,10}, 4))
}
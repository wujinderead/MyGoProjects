package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/magnetic-force-between-two-balls/

// In universe Earth C-137, Rick discovered a special form of magnetic force between two balls
// if they are put in his new invented basket. Rick has n empty baskets, the ith basket is
// at position[i], Morty has m balls and needs to distribute the balls into the baskets such that
// the minimum magnetic force between any two balls is maximum.
// Rick stated that magnetic force between two different balls at positions x and y is |x - y|.
// Given the integer array position and the integer m. Return the required force.
// Example 1:
//     |<-force=3->|<-force=3->|
//     ba          ba          ba
//     bu  bu  bu  bu          bu
//      1   2   3   4   5   6   7
//   Input: position = [1,2,3,4,7], m = 3
//   Output: 3
//   Explanation: Distributing the 3 balls into baskets 1, 4 and 7 will make the magnetic force
//     between ball pairs [3, 3, 6]. The minimum magnetic force is 3. We cannot achieve a larger
//     minimum magnetic force than 3.
// Example 2:
//   Input: position = [5,4,3,2,1,1000000000], m = 2
//   Output: 999999999
//   Explanation: We can use baskets 1 and 1000000000.
// Constraints:
//   n == position.length
//   2 <= n <= 10^5
//   1 <= position[i] <= 10^9
//   All integers in position are distinct.
//   2 <= m <= position.length

// O(nlog(max-min))
func maxDistance(position []int, m int) int {
	// sosrt the array first
	sort.Sort(sort.IntSlice(position))
	
	// binary search to find the max minimal

	// NOTE: we need to add r=max-min+1 if we use mid=(l+r)/2
	// if we use mid=(l+r)/2+1, we need not increment r, the manner will be "l=mid, r=mid-1, return l"
	l, r := 1, position[len(position)-1]-position[0]+1  
	for l<r {
		mid := l+(r-l)/2
		// count how many balls can be held if minimal gap is mid
		count, prev := 1, 0
		for i:=1; i<len(position); i++ {
			if position[i]-position[prev] >= mid {
				count++
				prev = i
			}
		}
		if count < m {
			r = mid
		} else {
			l = mid+1
		}
	}
	return l-1
}

func main() {
	for _, v := range []struct{arr []int; m, ans int} {
		{[]int{1,2,3,4,7}, 3, 3},
		{[]int{5,4,3,2,1,1000000000}, 2, 999999999},
	} {
		fmt.Println(maxDistance(v.arr, v.m), v.ans)
	}
	test()
}

// test binary search end conditions
func test() {
	for i:=1; i<=11; i++ {
		for j:=0; j<i; j++ {
			bools := make([]bool, i)
			for k:=0; k<=j; k++ {
				bools[k] = true
			}
			l, r := 0, i-1
			for l<r {
				mid := l+(r-l)/2+1
				if !bools[mid] {
					r = mid-1
				} else {
					l = mid
				}
			}
			fmt.Println(bools, l)
			l, r = 0, i   // extend the search range out of array
			for l<r {
				mid := l+(r-l)/2         // when l+1=r, mid=(l+r)/2 will fix on l, so we need l=mid+1
				if !bools[mid] {
					r = mid
				} else {
					l = mid+1   // so when arr[t t t], l and r will finally be at the dummy arr[t t t][r] 
				}
			}
			fmt.Println(bools, l-1)  // so we need return l-1
		}
	}
}
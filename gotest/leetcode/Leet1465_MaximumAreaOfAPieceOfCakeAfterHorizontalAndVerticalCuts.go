package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/

// Given a rectangular cake with height h and width w, and two arrays of integers
// horizontalCuts and verticalCuts where horizontalCuts[i] is the distance from the 
// top of the rectangular cake to the ith horizontal cut and similarly, verticalCuts[j] 
// is the distance from the left of the rectangular cake to the jth vertical cut. 
// Return the maximum area of a piece of cake after you cut at each horizontal and vertical 
// position provided in the arrays horizontalCuts and verticalCuts. Since the answer can be 
// a huge number, return this modulo 10^9 + 7.
// Example 1: 
//   Input: h = 5, w = 4, horizontalCuts = [1,2,4], verticalCuts = [1,3]
//   Output: 4 
//   Explanation: The figure above represents the given rectangular cake. Red lines
//     are the horizontal and vertical cuts. After you cut the cake, the green piece 
//     of cake has the maximum area.
// Example 2: 
//   Input: h = 5, w = 4, horizontalCuts = [3,1], verticalCuts = [1]
//   Output: 6
//   Explanation: The figure above represents the given rectangular cake. Red lines
//     are the horizontal and vertical cuts. After you cut the cake, the green and 
//     yellow pieces of cake have the maximum area.
// Example 3: 
//   Input: h = 5, w = 4, horizontalCuts = [3], verticalCuts = [3]
//   Output: 9
// Constraints: 
//   2 <= h, w <= 10^9 
//   1 <= horizontalCuts.length < min(h, 10^5) 
//   1 <= verticalCuts.length < min(w, 10^5) 
//   1 <= horizontalCuts[i] < h 
//   1 <= verticalCuts[i] < w 
//   It is guaranteed that all elements in horizontalCuts are distinct. 
//   It is guaranteed that all elements in verticalCuts are distinct. 

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
    horizontalCuts = append(horizontalCuts, h)
    verticalCuts = append(verticalCuts, w)
    sort.Sort(sort.IntSlice(horizontalCuts))
    sort.Sort(sort.IntSlice(verticalCuts))
    maxh, maxw := 0, 0
    prev := 0
    for i := range horizontalCuts {
    	if maxh<horizontalCuts[i]-prev {
    		maxh = horizontalCuts[i]-prev
    	}
    	prev = horizontalCuts[i]
    }
    prev = 0
    for i := range verticalCuts {
    	if maxw<verticalCuts[i]-prev {
    		maxw = verticalCuts[i]-prev
    	}
    	prev = verticalCuts[i]
    }
    return (maxh*maxw)%1000000007
}

func main() {
	fmt.Println(maxArea(5, 4, []int{1,2,4}, []int{1,3}), 4)
	fmt.Println(maxArea(5, 4, []int{3,1}, []int{1}), 6)
	fmt.Println(maxArea(5, 4, []int{3}, []int{3}), 9)
	fmt.Println(maxArea(1000000000, 1000000000, []int{3}, []int{3}), 9)
}
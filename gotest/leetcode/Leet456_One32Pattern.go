package main

import "fmt"

// https://leetcode.com/problems/132-pattern/

// Given a sequence of n integers a1, a2, ..., an, a 132 pattern is a subsequence
// ai, aj, ak such that i < j < k and ai < ak < aj. Design an algorithm that takes 
// a list of n numbers as input and checks whether there is a 132 pattern in the list. 
// Note: n will be less than 15,000. 
// Example 1: 
//   Input: [1, 2, 3, 4]
//   Output: False
//   Explanation: There is no 132 pattern in the sequence.
// Example 2: 
//   Input: [3, 1, 4, 2]
//   Output: True
//   Explanation: There is a 132 pattern in the sequence: [1, 4, 2].
// Example 3: 
//   Input: [-1, 3, 2, 0]
//   Output: True
//   Explanation: There are three 132 patterns in the sequence: [-1, 3, 2], [-1, 3,0] and [-1, 2, 0].

func find132pattern(nums []int) bool {
	if len(nums)==0 {
		return false
	}
    stack := [][2]int{{nums[0], nums[0]}}
    fmt.Println(stack, nums[:1])
    var min int
    for i:=1; i<len(nums); i++ {
    	t := nums[i]
    	if t<stack[len(stack)-1][0] {
    		stack = append(stack, [2]int{t, t})
    		fmt.Println(stack, nums[:i+1])
    		continue
    	}
    	min = stack[len(stack)-1][0]
    	for len(stack)>0 && t>=stack[len(stack)-1][1] {
    		stack = stack[:len(stack)-1]
    	}
    	if len(stack)>0 && t>stack[len(stack)-1][0] && t<stack[len(stack)-1][1] {
    		return true
    	}
    	stack = append(stack, [2]int{min, t})
    	fmt.Println(stack, nums[:i+1])
    }
    return false
}

func main() {
	fmt.Println(find132pattern([]int{1,2,3,4}))
	fmt.Println(find132pattern([]int{3,1,4,2}))
	fmt.Println(find132pattern([]int{-1,3,2,0}))
	// the stack and the array, the algorighm is O(n) since a element is added and removed from stack only once
	// [[90 90]]                                           [90]
	// [[90 100]]                                          [90 100]
	// [[90 100] [70 70]]                                  [90 100 70]
    // [[90 100] [70 80]]                                  [90 100 70 80]
    // [[90 100] [70 80] [65 65]]                          [90 100 70 80 65]
    // [[90 100] [70 80] [65 65] [50 50]]                  [90 100 70 80 65 50]
    // [[90 100] [70 80] [65 65] [50 60]]                  [90 100 70 80 65 50 60]
    // [[90 100] [70 80] [65 65] [50 60] [30 30]]          [90 100 70 80 65 50 60 30]
    // [[90 100] [70 80] [65 65] [50 60] [30 40]]          [90 100 70 80 65 50 60 30 40]
    // [[90 100] [70 80] [65 65] [50 60] [30 45]]          [90 100 70 80 65 50 60 30 40 45]
    // [[90 100] [70 80] [65 65] [50 60] [30 45] [20 20]]  [90 100 70 80 65 50 60 30 40 45 20]
    // [[90 100] [20 85]]                                  [90 100 70 80 65 50 60 30 40 45 20 85]
    // [[20 110]]                                          [90 100 70 80 65 50 60 30 40 45 20 85 110]
	fmt.Println(find132pattern([]int{90,100,70,80,65,50,60,30,40,45,20,85,110,105}))
}
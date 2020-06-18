package main

import (
    "fmt"
)

// https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/

// Given an array arr of positive integers, consider all binary trees such that: 
//   Each node has either 0 or 2 children; 
//   The values of arr correspond to the values of each leaf in an in-order traversal of the tree. 
//     (Recall that a node is a leaf if and only if it has 0 children.) 
//   The value of each non-leaf node is equal to the product of the largest leaf value in its 
//     left and right subtree respectively. 
// Among all possible binary trees considered, return the smallest possible sum of the values of 
// each non-leaf node. It is guaranteed this sum fits into a 32-bit integer. 
// Example 1: 
//   Input: arr = [6,2,4]
//   Output: 32
//   Explanation:
//     There are two possible trees.  The first has non-leaf node sum 36, and the sec
//     ond has non-leaf node sum 32.
//     
//         24            24
//        /  \          /  \
//       12   4        6    8
//      /  \               / \
//     6    2             2   4
// Constraints: 
//   2 <= arr.length <= 40 
//   1 <= arr[i] <= 15 
//   It is guaranteed that the answer fits into a 32-bit signed integer (ie. it is less than 2^31). 

// O(n^3) method
func mctFromLeafValues(arr []int) int {
	// key {i,j} means arr[i...j], value[0] means the sum we got from arr[i...j], value[0] means the max of arr[i...j]
    mapp := make(map[[2]int][2]int)   
    for i := range arr {
    	mapp[[2]int{i,i}] = [2]int{0, arr[i]}
    }
    for diff:=1; diff<len(arr); diff++ {
    	for i:=0; i+diff<len(arr); i++ {
    		j := i+diff
    		sum := 0x7fffffff
    		maxv := max(arr[i], mapp[[2]int{i+1,j}][1])
    		for k:=i; k<j; k++ {
    			left, right := mapp[[2]int{i,k}], mapp[[2]int{k+1,j}]
    			sum = min(sum, left[1]*right[1]+left[0]+right[0])
    		}
    		mapp[[2]int{i,j}] = [2]int{sum, maxv}
    	}
    }
    return mapp[[2]int{0, len(arr)-1}][0]
}

// O(n) solution: https://leetcode.com/problems/minimum-cost-tree-from-leaf-values/discuss/339959/One-Pass-O(N)-Time-and-Space
// for example, A=[4,3,5,2,1], stack will be:
// stack=[4]
// stack=[4,3]
// stack=[5]    sum += 3*4 + 4*5  (pop 3 got 3*4, pop 4 got 4*5)
// stack=[5,2]    
// stack=[5,2,1] 
// stack=[5]    sum += 1*2 + 2*5  (pop 1 got 1*2, pop 2 got 2*5)
func mctFromLeafValuesOn(A []int) int {
	res := 0
    stack := []int{0x7fffffff}
    for _, v := range A {
        for stack[len(stack)-1] <= v {
            mid := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res += mid * min(stack[len(stack)-1], v)
            fmt.Println(mid, min(stack[len(stack)-1], v))
        }
        stack = append(stack, v)
    }
    for len(stack)>2 {
        res += stack[len(stack)-1] * stack[len(stack)-2]
        fmt.Println(stack[len(stack)-1], stack[len(stack)-2])
        stack = stack[:len(stack)-1]

    }
    return res
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a<b {
		return a
	}
	return b
}

func main() {
	fmt.Println(mctFromLeafValues([]int{6,2,4}))
	fmt.Println(mctFromLeafValues([]int{2,4}))
	fmt.Println(mctFromLeafValues([]int{4,5,9,2,8,3,7,1,6}))
	fmt.Println(mctFromLeafValues([]int{2,1,4,3}))
	fmt.Println(mctFromLeafValues([]int{4,3,5,2,1}))
	fmt.Println(mctFromLeafValuesOn([]int{4,3,5,2,1}))
}
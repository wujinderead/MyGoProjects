package main

import (
    "fmt"
)

// https://leetcode.com/problems/adding-two-negabinary-numbers/

// Given two numbers arr1 and arr2 in base -2, return the result of adding them together.
// Each number is given in array format:  as an array of 0s and 1s, from most significant bit to least 
// significant bit. For example, arr = [1,1,0,1] represents the number (-2)^3 + (-2)^2 + (-2)^0 = -3.  
// A number arr in array format is also guaranteed to have no leading zeros: either arr == [0] or arr[0] == 1.
// Return the result of adding arr1 and arr2 in the same format: as an array of 0s and 1s with no leading zeros.
// Example 1:
//   Input: arr1 = [1,1,1,1,1], arr2 = [1,0,1]
//   Output: [1,0,0,0,0]
//   Explanation: arr1 represents 11, arr2 represents 5, the output represents 16.
// Note:
//   1 <= arr1.length <= 1000
//   1 <= arr2.length <= 1000
//   arr1 and arr2 have no leading zeros
//   arr1[i] is 0 or 1
//   arr2[i] is 0 or 1

// in -2 base, 1+1 results in 0 with -1 add to more significant bit.
// and -1 in -2 base is actually 11 in -2 base. 
func addNegabinary(arr1 []int, arr2 []int) []int {
	reverse(arr1)
	reverse(arr2)
	n := max(len(arr1), len(arr2))+2
    ret := make([]int, n)
    sig := 0
    for i:=0; i<n; i++ {
    	a1, a2 := 0, 0
    	if i<len(arr1) {
    		a1 = arr1[i]
    	}
    	if i<len(arr2) {
    		a2 = arr2[i]
    	}
    	if a1+a2+sig==0 || a1+a2+sig==1 {
    		ret[i] = a1+a2+sig
    		sig = 0
            continue
    	}
    	if a1+a2+sig==2 || a1+a2+sig==3 {
    		ret[i] = a1+a2+sig-2
    		sig = -1
            continue
    	}
    	if a1+a2+sig==-1 {
    		ret[i] = 1
    		sig = 1
    	}
    }
    reverse(ret)
    for i:=range ret {
        if ret[i] != 0 {
            return ret[i:]   // from first non-zero
        }
    }
    return []int{0}   // all 0
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func reverse(a []int) {
	for i,j:=0,len(a)-1; i<j; i,j=i+1,j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	fmt.Println(addNegabinary([]int{1,1,1,1,1}, []int{1,0,1}))
	fmt.Println(addNegabinary([]int{1}, []int{1}))
    fmt.Println(addNegabinary([]int{1}, []int{1,1}))
    fmt.Println(addNegabinary([]int{1,0,0}, []int{1,1,0,0}))
    fmt.Println(addNegabinary([]int{1,1}, []int{1,1}))
}
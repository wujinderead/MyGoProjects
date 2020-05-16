package main

import "fmt"

// https://leetcode.com/problems/find-k-closest-elements/

// Given a sorted array arr, two integers k and x, find the k closest elements to
// x in the array. The result should also be sorted in ascending order. If there is a
// tie, the smaller elements are always preferred.
// Example 1: 
//   Input: arr = [1,2,3,4,5], k = 4, x = 3
//   Output: [1,2,3,4]
// Example 2: 
//   Input: arr = [1,2,3,4,5], k = 4, x = -1
//   Output: [1,2,3,4]
// Constraints:
//   1 <= k <= arr.length
//   1 <= arr.length <= 10^4
//   Absolute value of elements in the array and x will not exceed 104

// binary search to find the position x in arr. then use two pointers.
// time O(logn+k)
func findClosestElements(arr []int, k int, x int) []int {
    if x<=arr[0] {
    	return arr[:k]
	}
	if x>=arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}
	lo, hi := 0, len(arr)-1
	for lo<hi {
		mid := (lo+hi)/2
		if arr[mid]>x {
			hi = mid
		} else if arr[mid]<x {
			lo = mid+1
		} else {
			lo = mid
			break
		}
	}
	// arr[lo]>=x
	if abs(arr[lo-1]-x) <= abs(arr[lo]-x) {
		lo = lo-1
	}
	lo, hi = lo, lo
	count:= 1
	for count<k && lo>=0 && hi<len(arr) {
		if lo==0 {
			return arr[:k]
		}
		if hi==len(arr)-1 {
			return arr[len(arr)-k:]
		}
		if abs(arr[lo-1]-x) <= abs(arr[hi+1]-x) {
			lo--
		} else {
			hi++
		}
		count++
	}
	return arr[lo: lo+k]
}

func abs(a int) int {
	if a<0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(findClosestElements([]int{1,2,3,4,5}, 4, 3))
	fmt.Println(findClosestElements([]int{1,2,3,4,5}, 4, -1))
	fmt.Println(findClosestElements([]int{1,3,6,9,10}, 2, 6))
	fmt.Println(findClosestElements([]int{1,3,6,9,10}, 2, 7))
	fmt.Println(findClosestElements([]int{1,2,3,3,6,6,7,7,9,9}, 8, 8))
	for i:=1; i<=9; i++ {
		arr := []int{1,2,3,3,6,6,7,7,9,9}
		lo, hi := 0, len(arr)-1
		for lo<hi {
			mid := (lo+hi)/2
			if arr[mid]>i {
				hi = mid
			} else if arr[mid]<i {
				lo = mid+1
			} else {
				lo = mid
				break
			}
		}
		fmt.Println(i, lo, hi, arr[lo])
	}
}
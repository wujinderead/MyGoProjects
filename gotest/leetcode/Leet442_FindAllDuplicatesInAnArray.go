package main

import "fmt"

// https://leetcode.com/problems/find-all-duplicates-in-an-array/

// Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice 
// and others appear once. Find all the elements that appear twice in this array.
// Could you do it without extra space and in O(n) runtime?
// Example:
//   Input: [4,3,2,7,8,2,3,1]
//   Output: [2,3]

// for example, arr=[4,3,2,7,8,2,3,1]; for arr[0]=4, it should be at arr[3], 
// we set the value in arr[3]=-7, we retain the information that arr[3]=7,
// and we add the information that 4 has occurred.
func findDuplicates(nums []int) []int {
	buf := []int{}
    for i := range nums {
    	if nums[abs(nums[i])-1] > 0 {
			nums[abs(nums[i])-1] = -nums[abs(nums[i])-1]
		} else {
			buf = append(buf, abs(nums[i]))
		}
    }
    return buf
}

func abs(a int) int {
	if a<0 {
		return -a
	}
	return a
}

// put integers in the right place, then we can check duplicate
func findDuplicates2Pass(nums []int) []int {
	buf := []int{}
    for i := range nums {
    	for nums[i] != i+1 {
    		v := nums[i]
    		if nums[v-1] == v {
    			break
    		} 
    		nums[i], nums[v-1] = nums[v-1], nums[i]
    	}
    }
    for i := range nums {
    	if nums[i] != i+1 {
    		buf = append(buf, nums[i])
    	}
    }
    return buf
}

func main() {
	fmt.Println(findDuplicates([]int{4,3,2,7,8,2,3,1}), []int{2,3})
	fmt.Println(findDuplicates([]int{5,4,6,7,9,3,10,9,5,6}), []int{5,6,9})
	fmt.Println(findDuplicates2Pass([]int{4,3,2,7,8,2,3,1}), []int{2,3})
	fmt.Println(findDuplicates2Pass([]int{5,4,6,7,9,3,10,9,5,6}), []int{5,6,9})
}
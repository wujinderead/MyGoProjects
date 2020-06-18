package main

import (
    "fmt"
)

// https://leetcode.com/problems/sum-of-mutated-array-closest-to-target/

// Given an integer array arr and a target value target, return the integer value such that when 
// we change all the integers larger than value in the given array to be equal to value, the sum 
// of the array gets as close as possible (in absolute difference) to target.
// In case of a tie, return the minimum such integer.
// Notice that the answer is not neccesarilly a number from arr.
// Example 1:
//   Input: arr = [4,9,3], target = 10
//   Output: 3
//   Explanation: When using 3 arr converts to [3, 3, 3] which sums 9 and that's the optimal answer.
// Example 2:
//   Input: arr = [2,3,5], target = 10
//   Output: 5
// Example 3:
//   Input: arr = [60864,25176,27249,21296,20204], target = 56803
//   Output: 11361
// Constraints:
//   1 <= arr.length <= 10^4
//   1 <= arr[i], target <= 10^5

func findBestValue(arr []int, target int) int {
	maxv := 0
	sum := 0
	for _, v := range arr {
		sum += v
		if v>maxv {
			maxv = v
		}
	}
	if target>=sum {
		return maxv   // can't be more larger
	}
	mindiff := sum-target   // base line diff
	minvalue := maxv        // base line value
    l, r := 0, maxv
    for l<=r {
    	mid := (l+r)/2

    	// get current sum
    	sum = 0
    	for _, v := range arr {
    		if v>mid {
    			sum += mid
    		} else {
    			sum += v
    		}
    	}

    	if sum > target {
    		r = mid-1
    		if sum-target < mindiff {
    			mindiff = sum-target
    			minvalue = mid
    		} else if sum-target == mindiff && mid<minvalue {  // tie
    			minvalue = mid
    		}
    	} else if sum < target {
    		l = mid+1
    		if target-sum < mindiff {
    			mindiff = target-sum
    			minvalue = mid
    		} else if target-sum == mindiff && mid<minvalue {  // tie
    			minvalue = mid
    		}
    	} else {
    		minvalue = mid  // find the value that make sum==target
    		break
    	}
    }
    return minvalue
}

func main() {
	fmt.Println(findBestValue([]int{4,9,3}, 10))
	fmt.Println(findBestValue([]int{2,3,5}, 10))
	fmt.Println(findBestValue([]int{60864,25176,27249,21296,20204}, 56803))
}
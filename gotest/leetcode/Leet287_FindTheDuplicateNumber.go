package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Given an array nums containing n + 1 integers where each integer is between 1 
// and n (inclusive), prove that at least one duplicate number must exist. Assume that 
// there is only one duplicate number, find the duplicate one. 
// Example 1: 
//   Input: [1,3,4,2,2]
//   Output: 2
// Example 2: 
//   Input: [3,1,3,4,2]
//   Output: 3 
// Note: 
//   You must not modify the array (assume the array is read only). 
//   You must use only constant, O(1) extra space. 
//   Your runtime complexity should be less than O(n2). 
//   There is only one duplicate number in the array, but it could be repeated more than once. 

// O(n) solution: if the array has no duplicate, e.g., [2,1,3], we have a map {0->2, 1->1, 2->3}; 
// and from index 0, we have a finite sequence: 0->2->3. however, if there is duplicate, e.g., [2,1,3,1], 
// the map is {0->2, (1,3)->1, 2->3}, the sequence is infinite: 0->2->3->1->1->1....
// thus it can be Transferred into "find the entrypoint of circle in linked list". 
// we use "Floyd tortoise and hare" method:
// let hare moves 2 steps and tortoise move 1 step simultaneously, the will meet in the circle.
// once they meat, keep hare still, and put tortoise to the start point, then move them both by 1 step.
// they will meat at the entrypoint.
func findDuplicate1(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	} 
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow  
}

func findDuplicate2(nums []int) int {
	// for example, for [1,3,4,2,2], there should be 2 numbers that <=2, but there are 3 numbers that <=2,
	// means some number <=2 is duplicated, so we can use binary search. the time complexity is O(nlogn).
	lo, hi := 1, len(nums)-1
	for lo<hi {
		mid := (lo+hi)/2
		count := 0
		for i:=range nums {
			if nums[i]<=mid {
				count++
			}
		}
		if count<=mid {
			lo = mid+1
		} else {
			hi = mid
		}
	}  
	return lo  
}

func main() {
	for _, findDuplicate := range []func([]int) int{findDuplicate1, findDuplicate2} {
		fmt.Println(findDuplicate([]int{1,1}))
		fmt.Println(findDuplicate([]int{2,2,2}))
		fmt.Println(findDuplicate([]int{1,3,4,2,2}))
		fmt.Println(findDuplicate([]int{3,1,3,4,2}))
		fmt.Println(findDuplicate([]int{1,2,2,2}))
		fmt.Println(findDuplicate([]int{2,2,2,3}))
		fmt.Println(findDuplicate([]int{1,3,3,2}))
		r := rand.New(rand.NewSource(time.Now().Unix()))
		a := r.Perm(20)
		for i := range a {
			if a[i]==0 {
				a[i]=17
			}
		}
		fmt.Println(a)
		fmt.Println(findDuplicate(a))
	}
}
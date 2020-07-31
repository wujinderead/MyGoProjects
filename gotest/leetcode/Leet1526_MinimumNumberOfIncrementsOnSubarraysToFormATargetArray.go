package main

import "fmt"

// https://leetcode.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/

// Given an array of positive integers target and an array initial of same size with all zeros.
// Return the minimum number of operations to form a target array from initial if you are allowed 
// to do the following operation:
//   Choose any subarray from initial and increment each value by one.
// The answer is guaranteed to fit within the range of a 32-bit signed integer.
// Example 1:
//   Input: target = [1,2,3,2,1]
//   Output: 3
//   Explanation: We need at least 3 operations to form the target array from the initial array.
//     [0,0,0,0,0] increment 1 from index 0 to 4 (inclusive).
//     [1,1,1,1,1] increment 1 from index 1 to 3 (inclusive).
//     [1,2,2,2,1] increment 1 at index 2.
//     [1,2,3,2,1] target array is formed.
// Example 2:
//   Input: target = [3,1,1,2]
//   Output: 4
//   Explanation: (initial)[0,0,0,0] -> [1,1,1,1] -> [1,1,1,2] -> [2,1,1,2] -> [3,1,1,2] (target).
// Example 3:
//   Input: target = [3,1,5,4,2]
//   Output: 7
//   Explanation: (initial)[0,0,0,0,0] -> [1,1,1,1,1] -> [2,1,1,1,1] -> [3,1,1,1,1] 
//   					-> [3,1,2,2,2] -> [3,1,3,3,2] -> [3,1,4,4,2] -> [3,1,5,4,2] (target).
// Example 4:
//   Input: target = [1,1,1,1]
//   Output: 1
// Constraints:
//   1 <= target.length <= 10^5
//   1 <= target[i] <= 10^5

// simple O(n) solution:
// https://leetcode.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/discuss/754623/Detailed-Explanation
func minNumberOperationsOn(target []int) int {
	sum := target[0]
	for i:=1; i<len(target); i++ {
		// it's easy to understand:
		// every larger number than prev need to increase by its own
		// every smaller number than prev can be covered by prev larger number
		if target[i]>target[i-1] {
			sum += target[i]-target[i-1]
		}
	}
	return sum
}

// segment tree solutions
func minNumberOperations(target []int) int {
	allsum := new(int)
	
	// build a segment tree
	tree := make([]int, 4*len(target))
	build(tree, target, 0, 0, len(target)-1)

	// main logic
	operations(tree, target, 0, 0, len(target)-1, allsum)
	return *allsum
}

func operations(tree, data []int, curmin, start, end int, allsum *int) {
	if start > end {  // skip illegal interval
		return
	}
	// get minInd that data[minInd] = min(data[start...end])
	minInd := query(tree, data, 0, 0, len(data)-1, start, end)

	// we need such operations
	*allsum += data[minInd]-curmin

	// process left and right parts of minInd, with a new curmin 
	operations(tree, data, data[minInd], start, minInd-1, allsum)
	operations(tree, data, data[minInd], minInd+1, end, allsum)
}

// tree[ind]=x where data[x] = min(data[left...right])
func build(tree, data []int, ind, left, right int) {
	if left == right {
		tree[ind] = left
		return
	}
	leftSub, rightSub, mid := ind*2+1, ind*2+2, left+(right-left)/2
	build(tree, data, leftSub, left, mid)     
	build(tree, data, rightSub, mid+1, right)

	// merge two sub trees
	if data[tree[leftSub]] <= data[tree[rightSub]] {   // prefer left index if tie
		tree[ind] = tree[leftSub]
	} else {
		tree[ind] = tree[rightSub]
	}
}

// query min value index 
func query(tree, data []int, ind, left, right, queryLeft, queryRight int) (minInd int) {
	// query range is equal to tree range
	if left==queryLeft && right==queryRight {
		return tree[ind]
	}
	leftSub, rightSub, mid := ind*2+1, ind*2+2, left+(right-left)/2

	if queryRight<=mid {        // query range in left sub tree
		return query(tree, data, leftSub, left, mid, queryLeft, queryRight)
	} else if queryLeft>mid {   // query range in right sub tree
		return query(tree, data, rightSub, mid+1, right, queryLeft, queryRight)
	}

	// else, query in both sub trees
	l := query(tree, data, leftSub, left, mid, queryLeft, mid)
	r := query(tree, data, rightSub, mid+1, right, mid+1, queryRight)
	
	// merge result for sub trees
	if data[l] <= data[r] {   // prefer left index if tie
		return l
	}
	return r
}

func main() {
	for _, v := range []struct{target []int; ans int} {
		{[]int{1,2,3,2,1}, 3},
		{[]int{3,1,1,2}, 4},
		{[]int{3,1,5,4,2}, 7},
		{[]int{1,1,1,1}, 1},
		{[]int{1,0,1,0}, 2},
		{[]int{0,0,0,0}, 0},
		{[]int{4}, 4},
		{[]int{0,5}, 5},
		{[]int{5,3,1,5,4,2,3,4,6,2,3}, 14},
	} {
		fmt.Println(minNumberOperations(v.target), minNumberOperationsOn(v.target), v.ans)
	}
}

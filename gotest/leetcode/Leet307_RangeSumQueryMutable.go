package main

import "fmt"

// https://leetcode.com/problems/range-sum-query-mutable/

// Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.
// The update(i, val) function modifies nums by updating the element at index i to val.
// Example:
//   Given nums = [1, 3, 5]
//   sumRange(0, 2) -> 9
//   update(1, 2)
//   sumRange(0, 2) -> 8
// Constraints:
//   The array is only modifiable by the update function.
//   You may assume the number of calls to update and sumRange function is distributed evenly.
//   0 <= i <= j <= nums.length - 1

type NumArray struct {
    data, tree []int
}

func Constructor(nums []int) NumArray {
	tree := make([]int, 4*len(nums))
	obj := NumArray{data: nums, tree: tree}
	if len(nums)==0 {
		return obj
	}
	buildHelper(tree, nums, 0, 0, len(nums)-1)
	return obj
}

// tree[ind] represent the range data[left...right]
func buildHelper(tree, data []int, ind, left, right int) {
	if left == right {
		tree[ind] = data[left]
		return
	}
	leftSub, rightSub, mid := ind*2+1, ind*2+2, left+(right-left)/2
	buildHelper(tree, data, leftSub, left, mid)      // tree[leftSub] represent the range data[left...mid]
	buildHelper(tree, data, rightSub, mid+1, right)  // tree[rightSub] represent the range data[mid+1...right]
	tree[ind] = tree[leftSub]+tree[rightSub]   // merge two sub trees
}

func (this *NumArray) Update(i int, val int)  {
	this.data[i] = val
	updateHelper(this.tree, 0, i, val, 0, len(this.data)-1)
}

func updateHelper(tree []int, ind, updateDataInd, val, left, right int) {
	if left == right {
		tree[ind] = val
		return
	}
	leftSub, rightSub, mid := ind*2+1, ind*2+2, left+(right-left)/2
	
	// if updateDataInd in left sub tree range (data[left...mid]), update left sub tree
	if updateDataInd <= mid {  
		updateHelper(tree, leftSub, updateDataInd, val, left, mid)
	} else {
		updateHelper(tree, rightSub, updateDataInd, val, mid+1, right)
	}
	tree[ind] = tree[leftSub]+tree[rightSub]   // update current node by merging two sub trees
}

func (this *NumArray) SumRange(i int, j int) int {
    return sumRangeHelper(this.tree, 0, 0, len(this.data)-1, i, j)
}

func sumRangeHelper(tree []int, ind, left, right, queryLeft, queryRight int) int {
	// query range is equal to tree range
	if left==queryLeft && right==queryRight {
		return tree[ind]
	}
	leftSub, rightSub, mid := ind*2+1, ind*2+2, left+(right-left)/2

	if queryRight<=mid {        // query range in left sub tree
		return sumRangeHelper(tree, leftSub, left, mid, queryLeft, queryRight)
	} else if queryLeft>mid {   // query range in right sub tree
		return sumRangeHelper(tree, rightSub, mid+1, right, queryLeft, queryRight)
	}

	// else, query in both sub trees
	l := sumRangeHelper(tree, leftSub, left, mid, queryLeft, mid)
	r := sumRangeHelper(tree, rightSub, mid+1, right, mid+1, queryRight)
	return l+r      // merge result for sub trees
}

func main() {
	obj := Constructor([]int{3})
	fmt.Println(obj.SumRange(0, 0))
	obj.Update(0, 6)
	fmt.Println(obj.SumRange(0, 0))

	obj = Constructor([]int{3, 8})
	fmt.Println(obj.SumRange(0, 1))
	fmt.Println(obj.SumRange(1, 1))
	fmt.Println(obj.SumRange(0, 0))
	obj.Update(0, 6)
	fmt.Println(obj.SumRange(0, 1))
	fmt.Println(obj.SumRange(1, 1))
	fmt.Println(obj.SumRange(0, 0))	
}

package main

import "fmt"

// https://leetcode.com/problems/count-of-smaller-numbers-after-self/

// You are given an integer array nums and you have to return a new counts array. The counts array 
// has the property where counts[i] is the number of smaller elements to the right of nums[i].
// Example:
// Input: [5,2,6,1]
// Output: [2,1,1,0] 
// Explanation:
//   To the right of 5 there are 2 smaller elements (2 and 1).
//   To the right of 2 there is only 1 smaller element (1).
//   To the right of 6 there is 1 smaller element (1).
//   To the right of 1 there is 0 smaller element.

func countSmallerMergeSort(nums []int) []int {
	if len(nums)==0 {
		return []int{}
	}
	less := make([]int, len(nums))
	data := make([][2]int, len(nums))  // nums and index pair
	for i:= range nums {
		data[i][0], data[i][1] = nums[i], i
	}
	buf := make([][2]int, len(nums))
	mergeSort(data, buf, less, 0, len(data)-1)
	return less
}

func mergeSort(data, buf [][2]int, less []int, start, end int) {
	if start==end {
		return 
	}
	mid := start + (end-start)/2
	mergeSort(data, buf, less, start, mid)
	mergeSort(data, buf, less, mid+1, end)

	// merge 2 sorted lists
	i, j := start, mid+1
	bi := 0
	for i<=mid && j<=end {
		// if we need to add a left element (x) to list,
		// we check how many right elements (y1, y2, ...) are already in the list, 
		// these (x,y1), (x,y2), ... are reverse pairs 
		if data[i][0] <= data[j][0] {       
			buf[bi] = data[i]               
			less[data[i][1]] += j-(mid+1)   // add count to original index 
			i++
			bi++
		} else {
			buf[bi] = data[j]
			j++
			bi++
		}
	}
	for i<=mid {
		buf[bi] = data[i]
		less[data[i][1]] += j-(mid+1)
		i++
		bi++
	}
	for j<=end {
		buf[bi] = data[j]
		j++
		bi++
	}
	copy(data[start: end+1], buf[:end-start+1])
}

// make a segment tree where each tree node is a sorted array of sub range.
// and we can query for some range, we can count how many numbers are less than target by binary search.
// we query n times, each query takes at most 2logn times to get the sorted array, 
// and binay search takes logn time, so the total complexity is O(n*logn*logn).
func countSmaller(nums []int) []int {
	if len(nums)==0 {
		return []int{}
	}
	ans := make([]int, len(nums))

	// build segment tree
	tree := make([][]int, len(nums)*4)
	build(tree, nums, 0, 0, len(nums)-1)
	
	// query
	for i:=len(nums)-2; i>=0; i-- {
		// get how many numbers in nums[i+1:] that are less than nums[i]
		ans[i] = query(tree, nums[i], 0, 0, len(nums)-1, i+1, len(nums)-1)
	}
	return ans 
}

// tree[ind] is the merge sort result of data[left...right]
func build(tree[][]int, data []int, ind, left, right int) {
	if left == right {
		tree[ind] = data[left: left+1]   // single element array, refer to data since we won't update
		return
	}
	leftSub, rightSub, mid := ind*2+1, ind*2+2, left+(right-left)/2
	build(tree, data, leftSub, left, mid)     
	build(tree, data, rightSub, mid+1, right)

	// merge two sub trees: merge 2 sorted lists
	tree[ind] = make([]int, right-left+1)
	i, j, llen, rlen := 0, 0, mid-left+1, right-mid
	for i < llen && j < rlen {
		if tree[leftSub][i] < tree[rightSub][j] {
			tree[ind][i+j] = tree[leftSub][i]
			i++
		} else {
			tree[ind][i+j] = tree[rightSub][j]
			j++
		}
	}
	for i < llen {
		tree[ind][i+j] = tree[leftSub][i]
		i++
	}
	for j < rlen {
		tree[ind][i+j] = tree[rightSub][j]
		j++
	}
}

// in range data[queryLeft...queryRight], get the count of x that makes data[x]<target 
func query(tree [][]int, target, ind, left, right, queryLeft, queryRight int) int {
	// query range is equal to tree range, binary search in range
	// i.e., in array tree[ind], find how many numbers that are less than target
	if left==queryLeft && right==queryRight {
		arr := tree[ind]
		if arr[len(arr)-1] < target {  // all elements < target
			return len(arr)
		}
		l, r := 0, len(arr)-1
		for l < r {   // find first index that >= target
			mid := l+(r-l)/2
			if arr[mid] < target {
				l = mid+1
			} else {
				r = mid
			}
		}
		return l
	}
	leftSub, rightSub, mid := ind*2+1, ind*2+2, left+(right-left)/2

	if queryRight <= mid {        // query range in left sub tree
		return query(tree, target, leftSub, left, mid, queryLeft, queryRight)
	} else if queryLeft > mid {   // query range in right sub tree
		return query(tree, target, rightSub, mid+1, right, queryLeft, queryRight)
	}

	// else, query in both sub trees
	l := query(tree, target, leftSub, left, mid, queryLeft, mid)
	r := query(tree, target, rightSub, mid+1, right, mid+1, queryRight)
	
	// merge result for sub trees, in this problem, sum the counts
	return l+r
}

func main() {
	countSmaller([]int{12,7,9,3,11,14,8,13})
	for _, v := range []struct{arr, ans[]int} {
		{[]int{5,2,6,1}, []int{2,1,1,0}},
		{[]int{1,2}, []int{0,0}},
		{[]int{1,1}, []int{0,0}},
		{[]int{2,1}, []int{1,0}},
		{[]int{1,3,2}, []int{0,1,0}},
		{[]int{2,1,3}, []int{1,0,0}},
		{[]int{3,2,2}, []int{2,0,0}},
		{[]int{12,7,9,3,11,14,8,13}, []int{5,1,2,0,1,2,0,0}},
	} {
		fmt.Println(countSmaller(v.arr), countSmallerMergeSort(v.arr), v.ans)
	}
}

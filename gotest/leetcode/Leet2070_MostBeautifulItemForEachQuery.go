package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/most-beautiful-item-for-each-query/

// You are given a 2D integer array items where items[i] = [pricei, beautyi] denotes the price
// and beauty of an item respectively.
// You are also given a 0-indexed integer array queries. For each queries[j], you want to determine
// the maximum beauty of an item whose price is less than or equal to queries[j]. If no such item
// exists, then the answer to this query is 0.
// Return an array answer of the same length as queries where answer[j] is the answer to the jth query.
// Example 1:
//   Input: items = [[1,2],[3,2],[2,4],[5,6],[3,5]], queries = [1,2,3,4,5,6]
//   Output: [2,4,5,5,6,6]
//   Explanation:
//     - For queries[0]=1, [1,2] is the only item which has price <= 1. Hence, the answer for this query is 2.
//     - For queries[1]=2, the items which can be considered are [1,2] and [2,4]. The maximum beauty among them is 4.
//     - For queries[2]=3 and queries[3]=4, the items which can be considered are [1,2], [3,2], [2,4], and [3,5]. The maximum beauty among them is 5.
//     - For queries[4]=5 and queries[5]=6, all items can be considered. Hence, the answer for them is the maximum beauty of all items, i.e., 6.
// Example 2:
//   Input: items = [[1,2],[1,2],[1,3],[1,4]], queries = [1]
//   Output: [4]
//   Explanation:
//     The price of every item is equal to 1, so we choose the item with the maximum beauty 4.
//     Note that multiple items can have the same price and/or beauty.
// Example 3:
//   Input: items = [[10,1000]], queries = [5]
//   Output: [0]
//   Explanation:
//     No item has a price less than or equal to 5, so no item can be chosen.
//     Hence, the answer to the query is 0.
// Constraints:
//   1 <= items.length, queries.length <= 10^5
//   items[i].length == 2
//   1 <= pricei, beautyi, queries[j] <= 10^9

// how to binary search?
// in this case, in arr=[1,2,2,2,2,2,2,3], we want the last element that <= t
// e.g., [1,2,2,2,2,2,2,3]
//        |           | |
//      t=1         t=2 t=3
// so, [1,2,2,2,2,2,2,3]
//      |     |       |
//     left  mid     right
//   t=1, arr[mid]>t, further search left part
//   t=2, arr[mid]=t, further search right part
//   t=3, arr[mid]<t, further search right part
// so, if arr[mid]>t, search left part (set right=mid or mid-1);
// if arr[mid]<=t, search right part (set left=mid, or mid+1).
// set loop condition as 'for left < right{}', then if left and right are adjacent (i.e. left+1=right),
// mid always = left, so we need left=mid+1 to avoid infinite loop, and set right=mid.
// the when left=right, arr[left] is the first element that > t, so arr[left-1] is the last element that <= t.
func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool { // sort items by price
		return items[i][0] < items[j][0]
	})
	max := make([]int, len(items)) // max[i] = the max beauty for items[0...i]
	max[0] = items[0][1]
	for i := 1; i < len(items); i++ {
		max[i] = max[i-1]
		if items[i][1] > max[i] {
			max[i] = items[i][1]
		}
	}

	ans := make([]int, len(queries))
	for i := range queries {
		t := queries[i]
		if t < items[0][0] {
			ans[i] = 0
			continue
		}
		left, right := 0, len(items)
		for left < right {
			mid := (left + right) / 2
			if t < items[mid][0] {
				right = mid
			} else {
				left = mid + 1
			}
		}
		ans[i] = max[left-1]
	}
	return ans
}

func main() {
	for _, v := range []struct {
		i      [][]int
		q, ans []int
	}{
		{[][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 5, 5, 6, 6}},
		{[][]int{{1, 2}, {1, 2}, {1, 3}, {1, 4}}, []int{1}, []int{4}},
		{[][]int{{10, 1000}}, []int{5}, []int{0}},
	} {
		fmt.Println(maximumBeauty(v.i, v.q), v.ans)
	}
}

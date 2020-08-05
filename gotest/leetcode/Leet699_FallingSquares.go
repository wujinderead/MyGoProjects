package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/falling-squares/

// On an infinite number line (x-axis), we drop given squares in the order they are given.
// The i-th square dropped (positions[i] = (left, side_length)) is a square with the left-most point 
// being positions[i][0] and sidelength positions[i][1]. The square is dropped with the bottom edge 
// parallel to the number line, and from a higher height than all currently landed squares. 
// We wait for each square to stick before dropping the next. The squares are infinitely sticky on 
// their bottom edge, and will remain fixed to any positive length surface they touch (either the number 
// line or another square). Squares dropped adjacent to each other will not stick together prematurely.
// Return a list ans of heights. Each height ans[i] represents the current highest height of any square 
// we have dropped, after dropping squares represented by positions[0], positions[1], ..., positions[i].
// Example 1:
//   Input: [[1, 2], [2, 3], [6, 1]]
//   Output: [2, 5, 5]
//   Explanation:
//     After the first drop of positions[0] = [1, 2], The maximum height of any square is 2.
//                             aaa                aaa 
//                             aaa                aaa
//                             aaa                aaa
//         aa           ->    aa          ->     aa
//        _aa____            _aa_____           _aa___a___
//        0123456            0123456            0123456
//     After the second drop of positions[1] = [2, 3], The maximum height of any square is 5. 
//     After the third drop of positions[2] = [6, 1], The maximum height of any square is still 5. 
//     Thus, we return an answer of [2, 5, 5].
// Example 2:
//   Input: [[100, 100], [200, 100]]
//   Output: [100, 100]
//   Explanation: Adjacent squares don't get stuck prematurely - only their bottom edge can stick to surfaces.
// Note:
//   1 <= positions.length <= 1000.
//   1 <= positions[i][0] <= 10^8.
//   1 <= positions[i][1] <= 10^6.

// O(N^2) solution with interval compaction.
// e.g. for interval [100,300] and [200, 400], we can deem it as [0,2] and [1,3],
// thus we can use only a 3-length int array to represent their states.
func fallingSquares(positions [][]int) []int {
	// find all distinct coordinates
	nmap := make(map[int]struct{})
	for _, v := range positions {
		nmap[v[0]] = struct{}{}
		nmap[v[0]+v[1]] = struct{}{}
	}

	// sort distinct coordinates
	nums := make([]int, 0, len(nmap))
	for k := range nmap {
		nums = append(nums, k)
	}
	sort.Sort(sort.IntSlice(nums))

	// make reverse map from real coordinates to pseudo coordiantes
	mapp := make(map[int]int)
	for i, v := range nums {
		mapp[v] = i
	}

	// main logic
	ans := make([]int, len(positions))
	height := make([]int, len(nums)-1)
	max := 0
	for i, v := range positions {
		p0 := mapp[v[0]]
		p1 := mapp[v[0]+v[1]]
		intervalmax := 0
		// get max height of this interval
		for j:=p0; j<p1; j++ {
			if height[j] > intervalmax {
				intervalmax = height[j]
			}
		}
		// set to new height
		for j:=p0; j<p1; j++ {
			height[j] = intervalmax + v[1] 
		}
		//fmt.Println(height)
		// update max height
		if intervalmax+v[1] > max {
			max = intervalmax + v[1]
		}
		ans[i] = max
	}
	return ans
}

// O(NlogN) solution with interval compaction.
// make a segment tree with range qurey and range update.
// to range update the segment tree, we need lazy propagation method:
// https://leetcode.com/articles/a-recursive-approach-to-segment-trees-range-sum-queries-lazy-propagation/.
func fallingSquaresSegmentTree(positions [][]int) []int {
	// find all distinct coordinates
	nmap := make(map[int]struct{})
	for _, v := range positions {
		nmap[v[0]] = struct{}{}
		nmap[v[0]+v[1]] = struct{}{}
	}

	// sort distinct coordinates
	nums := make([]int, 0, len(nmap))
	for k := range nmap {
		nums = append(nums, k)
	}
	sort.Sort(sort.IntSlice(nums))

	// make reverse map from real coordinates to pseudo coordiantes
	mapp := make(map[int]int)
	for i, v := range nums {
		mapp[v] = i
	}

	// main logic
	ans := make([]int, len(positions))
	// the original data is all 0 with length len(nums)-1
	// so segment tree length is 4*(len(nums)-1).
	// since data is all 0, and we want the max, so no need to build the tree.
	tree := make([]int, 4*(len(nums)-1))
	lazy := make([]int, 4*(len(nums)-1))
	max := 0
	for i, v := range positions {
		p0 := mapp[v[0]]
		p1 := mapp[v[0]+v[1]]
		// query range max
		intervalmax := query(tree, lazy, 0, 0, len(nums)-2, p0, p1-1)
		// update range with new value
		update(tree, lazy, 0, 0, len(nums)-2, p0, p1-1, intervalmax+v[1])
		//fmt.Println(p0, p1-1, intervalmax, intervalmax+v[1])
		// update max height
		if intervalmax+v[1] > max {
			max = intervalmax + v[1]
		}
		ans[i] = max
	}
	return ans
}

// query range max
func query(tree, lazy []int, ind, left, right, queryLeft, queryRight int) int {
	leftSubInd, rightSubInd, mid := 2*ind+1, 2*ind+2, left+(right-left)/2
	if lazy[ind] != 0 {         // this node is lazy
        tree[ind] = lazy[ind]   // normalize current node by removing laziness
        if (left != right) {                       // update lazy[] for children nodes
            lazy[leftSubInd] = lazy[ind]
            lazy[rightSubInd] = lazy[ind]
        }
        lazy[ind] = 0                      // current node processed. No longer lazy
    }

	// remain is the same to normal query
	if queryLeft==left && queryRight==right {
		return tree[ind]
	}
	if queryRight<=mid {
		return query(tree, lazy, leftSubInd, left, mid, queryLeft, queryRight)
	} else if queryLeft > mid {
		return query(tree, lazy, rightSubInd, mid+1, right, queryLeft, queryRight)
	}
	l := query(tree, lazy, leftSubInd, left, mid, queryLeft, mid)
	r := query(tree, lazy, rightSubInd, mid+1, right, mid+1, queryRight)
	return max(l, r)
}

// update range with new value
func update(tree, lazy []int, ind, left, right, updateLeft, updateRight, updateVal int) {
	leftSubInd, rightSubInd, mid := 2*ind+1, 2*ind+2, left+(right-left)/2
	if lazy[ind] != 0 {     // this node is lazy (since we won't update 0 in this problem)
		tree[ind] = lazy[ind]   // normalize current node by removing laziness
		if left != right {                  // update lazy[] for children nodes
            lazy[leftSubInd] = lazy[ind]
            lazy[rightSubInd] = lazy[ind]
        }
        lazy[ind] = 0              // current node processed. No longer lazy 
	}
	// if update range equal to data range, update tree node 
	if updateLeft==left && updateRight==right {
		lazy[ind] = updateVal
		tree[ind] = max(updateVal, tree[ind])
		return
	}
	if updateRight<=mid {          // only update left part
		update(tree, lazy, leftSubInd, left, mid, updateLeft, updateRight, updateVal)
	} else if updateLeft > mid {   // only update right part
		update(tree, lazy, rightSubInd, mid+1, right, updateLeft, updateRight, updateVal)
	} else {     // update both parts
		update(tree, lazy, leftSubInd, left, mid, updateLeft, mid, updateVal)
		update(tree, lazy, rightSubInd, mid+1, right, mid+1, updateRight, updateVal)
	}
	// merge two sub trees
	tree[ind] = max(tree[leftSubInd], tree[rightSubInd])
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	for _, v := range [][][]int{
		[][]int{{1, 2}, {2, 3}, {6, 1}},
		[][]int{{1, 2}, {1, 3}},
		[][]int{{100, 100}, {200, 100}},
		[][]int{{20, 5}, {24, 1}, {10, 10}, {17, 6}, {25, 20}, {19, 7}},
		[][]int{{20, 5}, {24, 1}, {10, 10}, {17, 6}, {25, 20}, {19, 6}},
	} {
		fmt.Println(fallingSquares(v))
		fmt.Println(fallingSquaresSegmentTree(v))
	}
}
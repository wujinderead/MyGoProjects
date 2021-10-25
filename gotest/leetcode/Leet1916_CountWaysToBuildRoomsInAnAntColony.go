package main

import (
	"fmt"
	"math/big"
)

// https://leetcode.com/problems/count-ways-to-build-rooms-in-an-ant-colony/

// You are an ant tasked with adding n new rooms numbered 0 to n-1 to your colony. You are given
// the expansion plan as a 0-indexed integer array of length n, prevRoom, where prevRoom[i] indicates
// that you must build room prevRoom[i] before building room i, and these two rooms must be connected
// directly. Room 0 is already built, so prevRoom[0] = -1. The expansion plan is given such that once
// all the rooms are built, every room will be reachable from room 0.
// You can only build one room at a time, and you can travel freely between rooms you have already built
// only if they are connected. You can choose to build any room as long as its previous room is already built.
// Return the number of different orders you can build all the rooms in. Since the answer may be large,
// return it modulo 10^9 + 7.
// Example 1:
//   Input: prevRoom = [-1,0,1]
//   Output: 1
//   Explanation: There is only one way to build the additional rooms: 0 → 1 → 2
// Example 2:
//   Input: prevRoom = [-1,0,0,1,2]
//   Output: 6
//   Explanation:
//     The 6 ways are:
//     0 → 1 → 3 → 2 → 4
//     0 → 2 → 4 → 1 → 3
//     0 → 1 → 2 → 3 → 4
//     0 → 1 → 2 → 4 → 3
//     0 → 2 → 1 → 3 → 4
//     0 → 2 → 1 → 4 → 3
// Constraints:
//   n == prevRoom.length
//   2 <= n <= 10^5
//   prevRoom[0] == -1
//   0 <= prevRoom[i] < n for all 1 <= i < n
//   Every room is reachable from room 0 once all the rooms are built.

// the problem equals to count the number of topological sort of a tree.
// the answer equal to n!/(a0*a1*...an) where a0, a1... is the number of nodes for each sub-tree.
// e.g. for this tree, has 5 nodes.
//        0
//       / \
//      1   2
//      |   |
//      3   4
// there is 5! permutations, but only 1/5 of them are started with 0.
// this logic applies to other sub-trees.
func waysToBuildRooms(prevRoom []int) int {
	outDegree := make([]int, len(prevRoom)) // the out degree of a node
	childNum := make([]int, len(prevRoom))  // nodes number of a sub-tree rooted at ith node
	childNum[0] = 1
	for i := 1; i < len(prevRoom); i++ {
		childNum[i] = 1
		outDegree[prevRoom[i]]++
	}
	subTreeNum := make([]int, 0, len(prevRoom))
	queue := make([]int, 0)
	for i := range outDegree { // push leaf node to queue
		if outDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 { // pop a leaf
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if childNum[cur] > 1 {
			subTreeNum = append(subTreeNum, childNum[cur]) // got its sub-tree node number
		}
		if cur == 0 { // skip root
			break
		}
		childNum[prevRoom[cur]] += childNum[cur] // add current sub-tree node number to parent
		outDegree[prevRoom[cur]]--               // decrement parent's out degree
		// push parent to queue, when outDegree=0 (all its children have been visited)
		if outDegree[prevRoom[cur]] == 0 {
			queue = append(queue, prevRoom[cur])
		}
	}

	// answer is n!/(a0*a1...) for a0, a1... in subTreeNum
	// a/b mod p = (a mod p) * (b^-1 mod p)
	mod := int(1e9 + 7)
	a, b := 1, 1
	for i := 1; i <= len(prevRoom); i++ {
		a = a * i
		a = a % mod
	}
	for _, v := range subTreeNum {
		b = b * v
		b = b % mod
	}
	p := big.NewInt(1e9 + 7)
	bb := big.NewInt(int64(b))
	binv := new(big.Int).ModInverse(bb, p)
	binvint := int(binv.Int64())
	a = (a * binvint) % mod
	return a
}

func main() {
	for _, v := range []struct {
		p   []int
		ans int
	}{
		{[]int{-1, 0, 1}, 1},
		{[]int{-1, 0, 0, 1, 2}, 6},
		{[]int{-1, 0, 0, 1, 2, 1}, 20},
		{[]int{-1, 0, 0, 0, 0, 0}, 120},
		{[]int{-1, 0, 0, 1, 1, 2, 2, 3}, 210},
	} {
		fmt.Println(waysToBuildRooms(v.p), v.ans)
	}
	p := make([]int, 0, 30000)
	p = append(p, -1)
	for i := 0; i <= 14998; i++ {
		p = append(p, i, i)
	}
	p = append(p, 14999)
	fmt.Println(waysToBuildRooms(p))
}

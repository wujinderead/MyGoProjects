package main

import (
    "fmt"
)

// https://leetcode.com/problems/kth-ancestor-of-a-tree-node/

// You are given a tree with n nodes numbered from 0 to n-1 in the form of a parent array where parent[i] 
// is the parent of node i. The root of the tree is node 0. Implement the function getKthAncestor(int node, int k) 
// to return the k-th ancestor of the given node. If there is no such ancestor, return -1. 
// The k-th ancestor of a tree node is the k-th node in the path from that node to the root. 
// Example: 
//              0
//             / \
//            1   2
//           / \ / \
//          3  4 5  6
//   Input:
//     ["TreeAncestor","getKthAncestor","getKthAncestor","getKthAncestor"]
//     [[7,[-1,0,0,1,1,2,2]],[3,1],[5,2],[6,3]]
//   Output:
//     [null,1,0,-1]
//   Explanation:
//     TreeAncestor treeAncestor = new TreeAncestor(7, [-1, 0, 0, 1, 1, 2, 2]);
//     treeAncestor.getKthAncestor(3, 1);  // returns 1 which is the parent of 3
//     treeAncestor.getKthAncestor(5, 2);  // returns 0 which is the grandparent of 5
//     treeAncestor.getKthAncestor(6, 3);  // returns -1 because there is no such ancestor
// Constraints: 
//   1 <= k <= n <= 5*10^4 
//   parent[0] == -1 indicating that 0 is the root node. 
//   0 <= parent[i] < n for all 0 < i < n 
//   0 <= node < n 
//   There will be at most 5*10^4 queries. 

type TreeAncestor struct {
    mapp map[[2]int]int
}

func Constructor(n int, parent []int) TreeAncestor {
	mapp := make(map[[2]int]int)

	// parent is the parent in one jump of each node, based on this, we can find the ancestor of 2 jumps, 4 jumps ...
	// time O(nlogn)
	for i := range parent {
		mapp[[2]int{i,1}] = parent[i]
	}
	jump:=2
	for jump<=n {
		for i:= range parent {
			p := mapp[[2]int{i, jump/2}]
			if p==-1 {
				mapp[[2]int{i, jump}] = -1
			} else {
				mapp[[2]int{i, jump}] = mapp[[2]int{p, jump/2}]
			}
		}
		jump*=2
	}
    return TreeAncestor{mapp: mapp}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	jump := 1
	for jump<=k {
		if jump&k > 0 {
			//fmt.Println(node, jump)
			node = this.mapp[[2]int{node, jump}]
			if node == -1 {
				return -1
			}
		}
		jump *= 2
	}
	return node
}

func main() {
	obj := Constructor(7,[]int{-1,3,3,0,0,4,4})
	fmt.Println(obj.mapp)
	fmt.Println(obj.GetKthAncestor(3,1))
	fmt.Println(obj.GetKthAncestor(5,2))
	fmt.Println(obj.GetKthAncestor(6,3))
	fmt.Println(obj.GetKthAncestor(0,1))
	fmt.Println(obj.GetKthAncestor(1,1))
	fmt.Println(obj.GetKthAncestor(1,2))

	obj = Constructor(10,[]int{-1,0,0,1,2,0,1,3,6,1})
	fmt.Println(obj.mapp)
	fmt.Println(obj.GetKthAncestor(8,6))
	fmt.Println(obj.GetKthAncestor(9,7))
	fmt.Println(obj.GetKthAncestor(1,1))
	fmt.Println(obj.GetKthAncestor(2,5))
	fmt.Println(obj.GetKthAncestor(4,2))
	fmt.Println(obj.GetKthAncestor(7,3))
	fmt.Println(obj.GetKthAncestor(3,7))
	fmt.Println(obj.GetKthAncestor(9,6))
	fmt.Println(obj.GetKthAncestor(3,5))
	fmt.Println(obj.GetKthAncestor(8,8))
}
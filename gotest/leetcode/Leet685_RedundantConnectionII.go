package main

import "fmt"

// https://leetcode.com/problems/redundant-connection-ii/

// In this problem, a rooted tree is a directed graph such that, there is exactly one node (the root)
// for which all other nodes are descendants of this node, plus every node has exactly one parent,
// except for the root node which has no parents.
// The given input is a directed graph that started as a rooted tree with N nodes (with distinct values
// 1, 2, ..., N), with one additional directed edge added. The added edge has two different vertices
// chosen from 1 to N, and was not an edge that already existed.
// The resulting graph is given as a 2D-array of edges. Each element of edges is a pair [u, v] that
// represents a directed edge connecting nodes u and v, where u is a parent of child v.
// Return an edge that can be removed so that the resulting graph is a rooted tree of N nodes.
// If there are multiple answers, return the answer that occurs last in the given 2D-array.
// Example 1:
//   Input: [[1,2], [1,3], [2,3]]
//   Output: [2,3]
//   Explanation: The given directed graph will be like this:
//       1
//      / \
//     v   v
//     2-->3
// Example 2:
//   Input: [[1,2], [2,3], [3,4], [4,1], [1,5]]
//   Output: [4,1]
//   Explanation: The given directed graph will be like this:
//     5 <- 1 -> 2
//          ^    |
//          |    v
//          4 <- 3
// Note: 
//   The size of the input 2D-array will be between 3 and 1000.
//   Every integer represented in the 2D-array will be between 1 and N, where N is the size of the input array.

// https://leetcode.com/problems/redundant-connection-ii/discuss/278105/topic 
// do the union-find process.
// case 1: the back edge is go back to root. we won't find any 2-parent node.
//         just return the circle's last edge.                                                 
//       A  <---╮     
//        \     |     
//       1 \    |  3  
//          B   |     
//           \  |     
//          2 \ |     
//             C   
// 
// case 2: the back edge points to its ancester. then we must encounter a 2-prarent node (B).
//         for the 2 edges (edge 1, edge 2), we need to delete the one that makes a circle.
//         so we delete edge 2 (do not use it for union-find. can't delete edge 1, because it's used).
//         then, if we can't find a circle (case 2a), the second edge (edge 2) is the answer.
//         if we can find a circle (case 2b), the first edge (edge 1) is the answer.                                    
//                   A                               A                 
//                    \                               \                     
//                   1 \                             2 \                 
//           case 2a:   B <-╮                case 2b:   B <-╮                
//                       \  |                            \  |                
//                      3 \ | 2                         3 \ | 1                
//                         C                               C   
//           
// case 3: the back edge is points to non-ancester. we must encounter a 2-parent node (D).
//         we use the same way as case 2. delete second encoutered edge (edge 4).
//         it must has no circle.                                
//        A                                      
//   1  /  \ 2
//     D   B
//     ^    \ 3
//     ╰----C  
//       4                                                                              
func findRedundantDirectedConnection(edges [][]int) []int {
	roots, parent := make([]int, len(edges)+1), make([]int, len(edges)+1)
	first, second, circle := [2]int{}, [2]int{}, [2]int{}
	hascircle := false
	for _, v := range edges {
		s, t := v[0], v[1]
		if parent[t]==0 {  // record parents
			parent[t] = s
		} else {   // find a node with 2 parents
			first = [2]int{parent[t], t}
			second = [2]int{s, t}
			continue    // when find a 2-parent node, do not add second edge to union-find
		}
		// union find
		if !hascircle {
			rs, rt := root(roots, s), root(roots, t)
			if rs == rt { // find a circle
				hascircle = true
				circle = [2]int{s, t}
			}
			roots[rs] = rt
		}
	}
	if hascircle {
		if first != [2]int{} {  // if has 2-parent node, return first
			return first[:]
		}
		return circle[:]        // else, redundant edge back to root, just return current edge
	}
	return second[:]   // no circle after remove second, return second
}

func root(roots []int, a int) int {
	for roots[a] != 0 {
		a = roots[a]
	}
	return a
}

func main() {
	for _, v := range []struct{edges [][]int; ans []int} {
		{[][]int{{1,2}, {1,3}, {2,3}}, []int{2,3}},
		{[][]int{{1,2}, {2,3}, {3,4}, {4,1}, {1,5}}, []int{4,1}},
		{[][]int{{4,2}, {1,5}, {5,2}, {5,3}, {2,4}}, []int{4,2}},
		{[][]int{{2,1}, {3,1}, {4,2}, {1,4}}, []int{2,1}},
		{[][]int{{4,1}, {1,2}, {1,3}, {4,5}, {5,6}, {6,5}}, []int{6,5}},
		{[][]int{{2,3}, {3,4}, {4,1}, {1,5}, {1,2}}, []int{1,2}},
		{[][]int{{3,1}, {1,4}, {3,5}, {1,2}, {1,5}}, []int{1,5}},
		{[][]int{{1,2}, {2,3}, {3,1}}, []int{3,1}},
		{[][]int{{3,4}, {4,1}, {1,2}, {2,3}, {5,1}}, []int{4,1}},
	} {
		fmt.Println(findRedundantDirectedConnection(v.edges), v.ans)
	}
}


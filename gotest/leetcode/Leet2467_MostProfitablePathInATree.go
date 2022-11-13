package main

import "fmt"

// https://leetcode.com/problems/most-profitable-path-in-a-tree/

// There is an undirected tree with n nodes labeled from 0 to n - 1, rooted at node 0. You are
// given a 2D integer array edges of length n - 1 where edges[i] = [ai, bi] indicates that there
// is an edge between nodes ai and bi in the tree.
// At every node i, there is a gate. You are also given an array of even integers amount, where
// amount[i] represents:
//   the price needed to open the gate at node i, if amount[i] is negative, or,
//   the cash reward obtained on opening the gate at node i, otherwise.
// The game goes on as follows:
//   Initially, Alice is at node 0 and Bob is at node bob.
//   At every second, Alice and Bob each move to an adjacent node. Alice moves
//     towards some leaf node, while Bob moves towards node 0.
//   For every node along their path, Alice and Bob either spend money to open the
//     gate at that node, or accept the reward. Note that:
// If the gate is already open, no price will be required, nor will there be any cash reward.
// If Alice and Bob reach the node simultaneously, they share the price/reward for opening the gate
// there. In other words, if the price to open the gate is c, then both Alice and Bob pay c / 2 each.
// Similarly, if the reward at the gate is c, both of them receive c / 2 each.
// If Alice reaches a leaf node, she stops moving. Similarly, if Bob reaches node 0, he stops moving.
// Note that these events are independent of each other.
// Return the maximum net income Alice can have if she travels towards the optimal leaf node.
// Example 1:
//   Input: edges = [[0,1],[1,2],[1,3],[3,4]], bob = 3, amount = [-2,4,2,-4,6]
//   Output: 6
//   Explanation:
//     The above diagram represents the given tree. The game goes as follows:
//     - Alice is initially on node 0, Bob on node 3. They open the gates of their respective nodes.
//       Alice's net income is now -2.
//     - Both Alice and Bob move to node 1.
//       Since they reach here simultaneously, they open the gate together and share the reward.
//       Alice's net income becomes -2 + (4 / 2) = 0.
//     - Alice moves on to node 3. Since Bob already opened its gate, Alice's income remains unchanged.
//       Bob moves on to node 0, and stops moving.
//     - Alice moves on to node 4 and opens the gate there. Her net income becomes 0 + 6 = 6.
//       Now, neither Alice nor Bob can make any further moves, and the game ends.
//       It is not possible for Alice to get a higher net income.
// Example 2:
//   Input: edges = [[0,1]], bob = 1, amount = [-7280,2350]
//   Output: -7280
//   Explanation:
//     Alice follows the path 0->1 whereas Bob follows the path 1->0.
//     Thus, Alice opens the gate at node 0 only. Hence, her net income is -7280.
// Constraints:
//   2 <= n <= 10⁵
//   edges.length == n - 1
//   edges[i].length == 2
//   0 <= ai, bi < n
//   ai != bi
//   edges represents a valid tree.
//   1 <= bob < n
//   amount.length == n
//   amount[i] is an even integer in the range [-10⁴, 10⁴].

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	// make graph
	n := len(edges) + 1
	graph := make([][]int, n)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], e[1])
		graph[e[1]] = append(graph[e[1]], e[0])
	}
	// first dfs to find the path from root to Bob
	var stack, bobpath []int
	visit1(graph, -1, 0, bob, &stack, &bobpath)
	// the amount of Bob's half path should be 0, the middle point amount should divide 2
	if len(bobpath)%2 == 1 {
		amount[bobpath[len(bobpath)/2]] /= 2
		for i := len(bobpath)/2 + 1; i < len(bobpath); i++ {
			amount[bobpath[i]] = 0
		}
	} else {
		for i := len(bobpath) / 2; i < len(bobpath); i++ {
			amount[bobpath[i]] = 0
		}
	}
	// second dfs to find Alice's income
	sum := new(int)
	max := -int(1e9)
	visit2(graph, -1, 0, amount, sum, &max)
	return max
}

func visit1(graph [][]int, parent, ind, bob int, stack *[]int, bobpath *[]int) {
	*stack = append(*stack, ind)
	if ind == bob {
		*bobpath = make([]int, len(*stack))
		copy(*bobpath, *stack)
	}
	for _, child := range graph[ind] {
		if child != parent {
			visit1(graph, ind, child, bob, stack, bobpath)
		}
	}
	*stack = (*stack)[:len(*stack)-1]
}

func visit2(graph [][]int, parent, ind int, amount []int, sum *int, max *int) {
	*sum += amount[ind]
	for _, child := range graph[ind] {
		if child != parent {
			visit2(graph, ind, child, amount, sum, max)
		}
	}
	if len(graph[ind]) == 1 && graph[ind][0] == parent { // found a leaf
		if *sum > *max {
			*max = *sum
		}
	}
	*sum -= amount[ind]
}

func main() {
	for _, v := range []struct {
		edges  [][]int
		bob    int
		amount []int
		ans    int
	}{
		{[][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, 3, []int{-2, 4, 2, -4, 6}, 6},
		{[][]int{{0, 1}}, 1, []int{-7280, 2350}, -7280},
		{[][]int{{0, 1}, {1, 2}, {2, 3}}, 3, []int{-5644, -6018, 1188, -8502}, -11662},
	} {
		fmt.Println(mostProfitablePath(v.edges, v.bob, v.amount), v.ans)
	}
}

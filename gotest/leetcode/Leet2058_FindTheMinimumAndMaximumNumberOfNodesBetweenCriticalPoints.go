package main

import "fmt"

// https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/

// A critical point in a linked list is defined as either a local maxima or a local minima.
// A node is a local maxima if the current node has a value strictly greater than the previous node and the next node.
// A node is a local minima if the current node has a value strictly smaller than the previous node and the next node.
// Note that a node can only be a local maxima/minima if there exists both a previous node and a next node.
// Given a linked list head, return an array of length 2 containing [minDistance, maxDistance] where minDistance is the
// minimum distance between any two distinct critical points and maxDistance is the maximum distance between any two
// distinct critical points. If there are fewer than two critical points, return [-1, -1].
// Example 1:
//Input: head = [3,1]
//Output: [-1,-1]
//Explanation: There are no critical points in [3,1].
// Example 2:
//   Input: head = [5,3,1,2,5,1,2]
//   Output: [1,3]
//   Explanation: There are three critical points:
//     - [5,3,1,2,5,1,2]: The third node is a local minima because 1 is less than 3 and 2.
//     - [5,3,1,2,5,1,2]: The fifth node is a local maxima because 5 is greater than 2 and 1.
//     - [5,3,1,2,5,1,2]: The sixth node is a local minima because 1 is less than 5 and 2.
//     The minimum distance is between the fifth and the sixth node. minDistance = 6 - 5 = 1.
//     The maximum distance is between the third and the sixth node. maxDistance = 6 - 3 = 3.
// Example 3:
//   Input: head = [1,3,2,2,3,2,2,2,7]
//   Output: [3,3]
//   Explanation: There are two critical points:
//     - [1,3,2,2,3,2,2,2,7]: The second node is a local maxima because 3 is greater than 1 and 2.
//     - [1,3,2,2,3,2,2,2,7]: The fifth node is a local maxima because 3 is greater than 2 and 2.
//     Both the minimum and maximum distances are between the second and the fifth node.
//     Thus, minDistance and maxDistance is 5 - 2 = 3.
//     Note that the last node is not considered a local maxima because it does not have a next node.
// Example 4:
//   Input: head = [2,3,3,2]
//   Output: [-1,-1]
//   Explanation: There are no critical points in [2,3,3,2].
// Constraints:
//   The number of nodes in the list is in the range [2, 10^5].
//   1 <= Node.val <= 10^5

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	criticals := make([]int, 0)
	cur := head.Next
	if cur == nil {
		return []int{-1, -1}
	}
	prev := head
	ind := 1
	for cur.Next != nil {
		next := cur.Next
		// found a critical
		if (cur.Val > prev.Val && cur.Val > next.Val) || (cur.Val < prev.Val && cur.Val < next.Val) {
			criticals = append(criticals, ind)
		}
		ind++
		prev = cur
		cur = cur.Next
	}
	if len(criticals) < 2 {
		return []int{-1, -1}
	}
	min := int(1e9)
	for i := 1; i < len(criticals); i++ {
		if criticals[i]-criticals[i-1] < min {
			min = criticals[i] - criticals[i-1]
		}
	}
	return []int{min, criticals[len(criticals)-1] - criticals[0]}
}

func main() {
	// 5,3,1,2,5,1,2
	head := &ListNode{Val: 5}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	fmt.Println(nodesBetweenCriticalPoints(head), []int{1, 3})
}

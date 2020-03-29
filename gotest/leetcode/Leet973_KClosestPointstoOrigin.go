package main

import (
	"fmt"
	"math/rand"
)

// https://leetcode.com/problems/k-closest-points-to-origin/

// We have a list of points on the plane. Find the K closest points to the origin (0, 0).
// (Here, the distance between two points on a plane is the Euclidean distance.)
// You may return the answer in any order. The answer is guaranteed to be unique
// (except for the order that it is in.)
// Example 1:
//   Input: points = [[1,3],[-2,2]], K = 1
//   Output: [[-2,2]]
//   Explanation:
//     The distance between (1, 3) and the origin is sqrt(10).
//     The distance between (-2, 2) and the origin is sqrt(8).
//     Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
//     We only want the closest K = 1 points from the origin, so the answer is just [[-2,2]].
// Example 2:
//   Input: points = [[3,3],[5,-1],[-2,4]], K = 2
//   Output: [[3,3],[-2,4]]
//   (The answer [[-2,4],[3,3]] would also be accepted.)
// Note:
//   1 <= K <= points.length <= 10000
//   -10000 < points[i][0] < 10000
//   -10000 < points[i][1] < 10000

// the top k problem
func kClosest(points [][]int, K int) [][]int {
	kClosestPoint(points, 0, len(points)-1, K-1)
	return points[:K]
}

func dist(point []int) int {
	return point[0]*point[0] + point[1]*point[1]
}

func kClosestPoint(points [][]int, start, end, K int) {
	if start >= end {
		return
	}
	if end-start < 2 {
		if dist(points[start]) > dist(points[end]) {
			points[start], points[end] = points[end], points[start]
			return
		}
	}
	pivot := partition(points, start, end)
	if pivot > K {
		kClosestPoint(points, start, pivot-1, K)
	} else if pivot < K {
		kClosestPoint(points, pivot+1, end, K)
	}
}

// partition by random pivot
func partition(points [][]int, start, end int) int {
	pivot := rand.Intn(end-start+1) + start
	points[pivot], points[end] = points[end], points[pivot]
	v := dist(points[end])
	i := start - 1
	for j := start; j < end; j++ {
		if dist(points[j]) < v {
			i++
			points[i], points[j] = points[j], points[i]
		}
	}
	points[i+1], points[end] = points[end], points[i+1]
	return i + 1
}

func main() {
	fmt.Println(kClosest([][]int{{1, 3}, {-2, 2}}, 1))
	fmt.Println(kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	fmt.Println(kClosest([][]int{{0, 1}, {1, 0}}, 2))
	fmt.Println(kClosest([][]int{{10, -2}, {2, -2}, {10, 10}, {9, 4}, {-8, 1}}, 4))
	fmt.Println(kClosest([][]int{{9, 0}, {7, 10}, {-4, -2}, {3, -9}, {9, 1}, {-5, -1}}, 5))
}

package main

import "fmt"

// https://leetcode.com/problems/car-fleet-ii/

// There are n cars traveling at different speeds in the same direction along a one-lane road. You are given
// an array cars of length n, where cars[i] = [positioni, speedi] represents:
//   positioni is the distance between the iᵗʰ car and the beginning of the road
//     in meters. It is guaranteed that positioni < positioni+1.
//   speedi is the initial speed of the iᵗʰ car in meters per second.
// For simplicity, cars can be considered as points moving along the number line. Two cars collide when they
// occupy the same position. Once a car collides with another car, they unite and form a single car fleet.
// The cars in the formed fleet will have the same position and the same speed, which is the initial speed
// of the slowest car in the fleet.
// Return an array answer, where answer[i] is the time, in seconds, at which the iᵗʰ car collides with the
// next car, or -1 if the car does not collide with the next car. Answers within 10⁻⁵ of the actual answers
// are accepted.
// Example 1:
//   Input: cars = [[1,2],[2,1],[4,3],[7,2]]
//   Output: [1.00000,-1.00000,3.00000,-1.00000]
//   Explanation: After exactly one second, the first car will collide with the second car, and form a car
//     fleet with speed 1 m/s. After exactly 3 seconds, the third car will collide with the fourth car,
//     and form a car fleet with speed 2 m/s.
// Example 2:
//   Input: cars = [[3,4],[5,4],[6,3],[9,1]]
//   Output: [2.00000,1.00000,1.50000,-1.00000]
// Constraints:
//   1 <= cars.length <= 10⁵
//   1 <= positioni, speedi <= 10⁶
//   positioni < positioni+1

func getCollisionTimes(cars [][]int) []float64 {
	ans := make([]float64, len(cars))
	var stack []int
	for i := len(cars) - 1; i >= 0; i-- { // from last car to first
		pos, speed := cars[i][0], cars[i][1]
		// pop cars that speed >= current, these cars are faster than current,
		// they collide with slower cars before current
		for len(stack) > 0 && cars[stack[len(stack)-1]][1] >= speed {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			ans[i] = -1
			stack = append(stack, i)
			continue
		}
		// for cars speed < current, will collide
		for len(stack) > 0 {
			topInd := stack[len(stack)-1] // the car on stack top
			topPos, topSpeed, topTime := cars[topInd][0], cars[topInd][1], ans[topInd]
			curTime := float64(topPos-pos) / float64(speed-topSpeed)
			if topTime == -1 || curTime <= topTime { // collide the top car, before top car collide
				ans[i] = curTime
				stack = append(stack, i)
				break
			} else { // the top car already collide, pop top car, try next top
				stack = stack[:len(stack)-1]
			}
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		cars [][]int
		ans  []float64
	}{
		{[][]int{{1, 2}, {2, 1}, {4, 3}, {7, 2}}, []float64{1.00000, -1.00000, 3.00000, -1.00000}},
		{[][]int{{3, 4}, {5, 4}, {6, 3}, {9, 1}}, []float64{2.00000, 1.00000, 1.50000, -1.00000}},
		{[][]int{{1, 10}, {2, 2}, {100, 1}}, []float64{0.125, 98, -1}},
		{[][]int{{1, 10}, {99, 2}, {100, 1}}, []float64{11, 1, -1}},
	} {
		fmt.Println(getCollisionTimes(v.cars), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/minimum-number-of-refueling-stops

// A car travels from a starting position to a destination which is target miles
// east of the starting position.
// Along the way, there are gas stations. Each station[i] represents a gas station
// that is station[i][0] miles east of the starting position, and has station[i][1]
// liters of gas.
// The car starts with an infinite tank of gas, which initially has startFuel liters
// of fuel in it. It uses 1 liter of gas per 1 mile that it drives.
// When the car reaches a gas station, it may stop and refuel, transferring all
// the gas from the station into the car.
// What is the least number of refueling stops the car must make in order to reach
// its destination? If it cannot reach the destination, return -1.
// Note that if the car reaches a gas station with 0 fuel left, the car can still
// refuel there. If the car reaches the destination with 0 fuel left, it is still
// considered to have arrived.
// Example 1:
//   Input: target = 1, startFuel = 1, stations = []
//   Output: 0
//   Explanation: We can reach the target without refueling.
// Example 2:
//   Input: target = 100, startFuel = 1, stations = [[10,100]]
//   Output: -1
//   Explanation: We can't reach the target (or even the first gas station).
// Example 3:
//   Input: target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]
//   Output: 2
//   Explanation:
//     We start with 10 liters of fuel.
//     We drive to position 10, expending 10 liters of fuel.  We refuel from 0 liters to 60 liters of gas.
//     Then, we drive from position 10 to position 60 (expending 50 liters of fuel),
//     and refuel from 10 liters to 50 liters of gas.  We then drive to and reach the target.
//     We made 2 refueling stops along the way, so we return 2.
// Note:
//   1 <= target, startFuel, stations[i][1] <= 10^9
//   0 <= stations.length <= 500
//   0 < stations[0][0] < stations[1][0] < ... < stations[stations.length-1][0] < target

// Note, O(nlogn) solution:
// When driving past a gas station, let's remember the amount of fuel it contained (use a heap).
// We don't need to decide yet whether to fuel up here or not - for example,
// there could be a bigger gas station up ahead that we would rather refuel at.
// When we run out of fuel before reaching the next station, we'll retroactively fuel up:
// greedily choosing the largest gas stations first.
// This is guaranteed to succeed because we drive the largest distance possible before
// each refueling stop, and therefore have the largest choice of gas stations to (retroactively) stop at.

// O(n²), n=len(stations)
func minRefuelStops(target int, startFuel int, stations [][]int) int {
	if startFuel >= target {
		return 0
	}
	if startFuel < target && len(stations) == 0 {
		return -1
	}
	// let f[n] be the maximal fuel we can get with n times refueling
	oldf, newf := make([]int, len(stations)+1), make([]int, len(stations)+1)
	oldf[0] = startFuel
	prevpos := 0
	for i := 0; i < len(stations); i++ {
		curpos := stations[i][0]
		curfuel := stations[i][1]
		dist := curpos - prevpos
		for j := 0; j <= i+1; j++ {
			if j < i+1 {
				oldf[j] = oldf[j] - dist
			}
			if j == 0 || oldf[j] < 0 {
				newf[j] = oldf[j]
				continue
			}
			if oldf[j-1] < 0 {
				newf[j] = oldf[j]
				continue
			}
			newf[j] = max(oldf[j], oldf[j-1]+curfuel)
		}
		prevpos = curpos
		oldf, newf = newf, oldf
	}
	for i := 0; i <= len(stations); i++ {
		if oldf[i] >= (target - stations[len(stations)-1][0]) {
			return i
		}
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minRefuelStops(1, 1, [][]int{}))
	fmt.Println(minRefuelStops(100, 1, [][]int{{10, 100}}))
	fmt.Println(minRefuelStops(100, 10, [][]int{{10, 60}, {20, 30}, {30, 30}, {60, 40}}))
}

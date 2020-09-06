package main

import "fmt"

// https://leetcode.com/problems/angle-between-hands-of-a-clock

// Given two numbers, hour and minutes. Return the smaller angle (in degrees) formed
// between the hour and the minute hand.
// Example 1:
//   Input: hour = 12, minutes = 30
//   Output: 165
// Example 2:
//   Input: hour = 3, minutes = 30
//   Output: 75
// Example 3:
//   Input: hour = 3, minutes = 15
//   Output: 7.5
// Example 4:
//   Input: hour = 4, minutes = 50
//   Output: 155
// Example 5:
//   Input: hour = 12, minutes = 0
//   Output: 0
// Constraints:
//   1 <= hour <= 12
//   0 <= minutes <= 59
//   Answers within 10^-5 of the actual value will be accepted as correct.

func angleClock(hour int, minutes int) float64 {
	h := float64(0)
	if hour<12 {
		h = 30*float64(hour)
	}
	h += float64(minutes)/2       // in one minute "hour hand" run 0.5 degree
	m := float64(minutes)*6       // in one minute "minute hand" run 6 degree
	ans := m-h
	if m-h<0 {
		ans = h-m
	}
	if ans > 180 {
		return 360-ans 
	}
	return ans
}

func main() {
	for _, v := range []struct{h, m int; ans float64} {
		{12, 30, 165},
		{3, 30, 75},
		{3, 15, 7.5},
		{4, 50, 155},
		{12, 0, 0},
	} {
		fmt.Println(angleClock(v.h, v.m), v.ans)
	}
}

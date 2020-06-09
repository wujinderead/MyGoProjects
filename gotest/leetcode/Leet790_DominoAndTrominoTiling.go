package main

import (
    "fmt"
)

// https://leetcode.com/problems/domino-and-tromino-tiling/

// We have two types of tiles: a 2x1 domino shape, and an "L" tromino shape. These shapes may be rotated. 
//  XX  <- domino
//  X  <- "L" tromino
//  XX
// Given N, how many ways are there to tile a 2 x N board? Return your answer modulo 10^9 + 7. 
// (In a tiling, every square must be covered by a tile. Two tilings are different if and only if there 
// are two 4-directionally adjacent cells on the board such that exactly one of the tilings has both squares 
// occupied by a tile.) 
// Example:
//   Input: 3
//   Output: 5
//   Explanation: 
//     The five different ways are listed below, different letters indicates different tiles:
//     XYZ XXZ XYY XXY XYY
//     XYZ YYZ XZZ XYY XXY 
// Note: 
//   N will be in range [1, 1000]. 

// r[i] is number of way to have 2xN:
// r[i] can be r[i-1]X or r[i-2]XX or l[i-1]X or      XX
//                   X          YY         XX    l[i-1]X
// l[i] is number of way to have 2xN with one tile missing
// l[i] can be r[i-2]X   or       XX
//                   XX      l[i-1]
// so we can get r[i] = 2*r[i-1]+r[i-3]
func numTilings(N int) int {
	if N<3 {
		return N
	}
    r := make([]int, N+1)
    l := make([]int, N+1)
    r[1], r[2] = 1, 2
    l[2] = 1
    for i:=3; i<=N; i++ {
    	r[i] = (r[i-1] + r[i-2] + 2*l[i-1]) % (1e9+7)
    	l[i] = (r[i-2] + l[i-1]) % (1e9+7)
    }
    return r[N]
}

func main() {
	for i:=1; i<=20; i++ {
		fmt.Println(numTilings(i))
	}
}
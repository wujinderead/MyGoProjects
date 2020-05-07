package main

import (
	"strconv"
	"fmt"
)

// https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/

// You are given an integer num. You will apply the following steps exactly two times:
//   Pick a digit x (0 <= x <= 9).
//   Pick another digit y (0 <= y <= 9). The digit y can be equal to x.
//   Replace all the occurrences of x in the decimal representation of num by y.
//   The new integer cannot have any leading zeros, also the new integer cannot be 0.
// Let a and b be the results of applying the operations to num the first and second times,
// respectively. Return the max difference between a and b.
// Example 1:
//   Input: num = 555
//   Output: 888
//   Explanation: The first time pick x = 5 and y = 9 and store the new integer in a.
//     The second time pick x = 5 and y = 1 and store the new integer in b.
//     We have now a = 999 and b = 111 and max difference = 888
// Example 2:
//   Input: num = 9
//   Output: 8
//   Explanation: The first time pick x = 9 and y = 9 and store the new integer in a.
//     The second time pick x = 9 and y = 1 and store the new integer in b.
//     We have now a = 9 and b = 1 and max difference = 8
// Example 3:
//   Input: num = 123456
//   Output: 820000
// Example 4:
//   Input: num = 10000
//   Output: 80000
// Example 5:
//   Input: num = 9288
//   Output: 8700
// Constraints:
//   1 <= num <= 10^8

func maxDiff(num int) int {
	// first turn we make the number maximal, second turn we made the number minimal.
	s := strconv.Itoa(num)
	// make maximal
	s1 := []byte(s)
	for i := range s {
		if s[i] != '9' {
			for j:=i; j<len(s); j++ {
				if s1[j] == s[i] {
					s1[j] = '9'
				}
			}
			break
		}
	}
	// make minimal
	s2 := []byte(s)
	if s[0] == '1' {
		for i := range s {
			if s[i] != '1' && s[i] != '0' {
				for j:=i; j<len(s); j++ {
					if s2[j] == s[i] {
						s2[j] = '0'
					}
				}
				break
			}
		}
	} else {
		for i:=0; i<len(s); i++ {
			if s[i] == s[0] {
				s2[i] = '1'
			}
		}
	}
	max, _ := strconv.Atoi(string(s1))
	min, _ := strconv.Atoi(string(s2))
	//fmt.Println(num, max, min)
	return max-min
}

func main() {
	fmt.Println(maxDiff(555))
	fmt.Println(maxDiff(9))
	fmt.Println(maxDiff(123456))
	fmt.Println(maxDiff(122114562))
	fmt.Println(maxDiff(10000))
	fmt.Println(maxDiff(9288))
	fmt.Println(maxDiff(1101057))
}
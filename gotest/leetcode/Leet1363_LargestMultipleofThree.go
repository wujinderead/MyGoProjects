package main

import "fmt"

// https://leetcode.com/problems/largest-multiple-of-three/

// Given an integer array of digits, return the largest multiple of three that can
// be formed by concatenating some of the given digits in any order.
// Since the answer may not fit in an integer data type, return the answer as a string.
// If there is no answer return an empty string.
// Example 1:
//   Input: digits = [8,1,9]
//   Output: "981"
// Example 2:
//   Input: digits = [8,6,7,1,0]
//   Output: "8760"
// Example 3:
//   Input: digits = [1]
//   Output: ""
// Example 4:
//   Input: digits = [0,0,0,0,0,0]
//   Output: "0"
// Constraints:
//   1 <= digits.length <= 10^4
//   0 <= digits[i] <= 9
//   The returning answer must not contain unnecessary leading zeros.

func largestMultipleOfThree(digits []int) string {
	counts := make([]int, 10)
	for _, v := range digits {
		counts[v]++ // count occurrence of each digit
	}
	sum := counts[1] + 2*counts[2] + 4*counts[4] + 5*counts[5] + 7*counts[7] + 8*counts[8]
	if sum%3 == 2 {
		switch {
		case counts[2] > 0:
			counts[2]--
		case counts[5] > 0:
			counts[5]--
		case counts[8] > 0:
			counts[8]--
		case counts[1] >= 2:
			counts[1] -= 2
		case counts[1] > 0 && counts[4] > 0:
			counts[1]--
			counts[4]--
		case counts[4] >= 2:
			counts[4] -= 2
		case counts[1] > 0 && counts[7] > 0:
			counts[1]--
			counts[7]--
		case counts[4] > 0 && counts[7] > 0:
			counts[4]--
			counts[7]--
		case counts[7] >= 2:
			counts[7] -= 2
		}

	} else if sum%3 == 1 {
		switch {
		case counts[1] > 0:
			counts[1]--
		case counts[4] > 0:
			counts[4]--
		case counts[7] > 0:
			counts[7]--
		case counts[2] >= 2:
			counts[2] -= 2
		case counts[2] > 0 && counts[5] > 0:
			counts[2]--
			counts[5]--
		case counts[5] >= 2:
			counts[5] -= 2
		case counts[2] > 0 && counts[8] > 0:
			counts[2]--
			counts[8]--
		case counts[5] > 0 && counts[8] > 0:
			counts[5]--
			counts[8]--
		case counts[8] >= 2:
			counts[8] -= 2
		}
	}
	// else sum%3==0, can use all numbers
	buf := make([]byte, 0, len(digits)/4)
	for i := 9; i > 0; i-- {
		for j := 0; j < counts[i]; j++ {
			buf = append(buf, '0'+byte(i))
		}
	}
	if len(buf) > 0 {
		for j := 0; j < counts[0]; j++ {
			buf = append(buf, '0')
		}
	} else if counts[0] > 0 {
		buf = append(buf, '0')
	}
	return string(buf)
}

func main() {
	fmt.Println(largestMultipleOfThree([]int{8, 1, 9}))
	fmt.Println(largestMultipleOfThree([]int{8, 6, 7, 1, 0}))
	fmt.Println(largestMultipleOfThree([]int{1}))
	fmt.Println(largestMultipleOfThree([]int{0, 0, 0, 0, 0, 0}))
}

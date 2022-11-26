package main

import "fmt"

// https://leetcode.com/problems/count-palindromic-subsequences/

// Given a string of digits s, return the number of palindromic subsequences of s having length 5.
// Since the answer may be very large, return it modulo 10⁹ + 7.
// Note:
//   A string is palindromic if it reads the same forward and backward.
// A subsequence is a string that can be derived from another string by deleting some or no characters
// without changing the order of the remaining characters.
// Example 1:
//   Input: s = "103301"
//   Output: 2
//   Explanation:
//     There are 6 possible subsequences of length 5: "10330","10331","10301","10301","13301","03301".
//     Two of them (both equal to "10301") are palindromic.
// Example 2:
//   Input: s = "0000000"
//   Output: 21
//   Explanation: All 21 subsequences are "00000", which is palindromic.
// Example 3:
//   Input: s = "9999900000"
//   Output: 2
//   Explanation: The only two palindromic subsequences are "99999" and "00000".
// Constraints:
//   1 <= s.length <= 10⁴
//   s consists of digits.

func countPalindromes(s string) int {
	// the reverse map, e.g., 13 -> 31, 08 -> 80, 80 -> 08
	mapp := make([]int, 100)
	for i := range mapp {
		mapp[i] = (i%10)*10 + i/10
	}
	const P = int(1e9) + 7
	leftSingle, rightSingle := [10]int{}, [10]int{}   // the count of 0, 1, ..., 9 of each part
	leftDouble, rightDouble := [100]int{}, [100]int{} // the count of 01, 02, ..., 99 of each part

	// split the string to left and right part: [0, 1, ..., l-3] l-2 [l-1]
	// left part
	for j := 0; j < len(s)-2; j++ {
		n := int(s[j] - '0')
		for i := 0; i < 10; i++ {
			leftDouble[i*10+n] += leftSingle[i]
		}
		leftSingle[n] += 1
	}

	// right part
	rightSingle[int(s[len(s)-1]-'0')] += 1
	count := 0

	// initially, left part is (0...j), j+1 is pivot, (j+2...) is right part
	// now use j as pivot, left remove j, right add j+1
	// it became left part (0...j-1), j is pivot, (j+1...) is right part
	for j := len(s) - 3; j >= 2; j-- {
		lr := int(s[j] - '0')
		ra := int(s[j+1] - '0')

		// add s[j+1] to right part
		for i := 0; i < 10; i++ {
			rightDouble[ra*10+i] += rightSingle[i]
		}
		rightSingle[ra] += 1

		// remove s[j] from left part
		leftSingle[lr] -= 1
		for i := 0; i < 10; i++ {
			leftDouble[i*10+lr] -= leftSingle[i]
		}

		// add count, leftDouble[00]*rightDouble[00] + leftDouble[01]*rightDouble[10] + ...
		for i := 0; i < 100; i++ {
			count = (count + leftDouble[i]*rightDouble[mapp[i]]) % P
		}
	}
	return count
}

func main() {
	for _, v := range []struct {
		s   string
		ans int
	}{
		{"00000", 1},
		{"0000", 0},
		{"000", 0},
		{"00", 0},
		{"0", 0},
		{"103301", 2},
		{"13033011", 10},
		{"0000000", 21},
		{"9999900000", 2},
	} {
		fmt.Println(countPalindromes(v.s), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/

// You are given a binary string s. You are allowed to perform two types of operations
// on the string in any sequence:
// Type-1: Remove the character at the start of the string s and append it to the
//   end of the string.
// Type-2: Pick any character in s and flip its value, i.e., if its value is '0'
//   it becomes '1' and vice-versa.
// Return the minimum number of type-2 operations you need to perform such that s becomes alternating.
// The string is called alternating if no two adjacent characters are equal.
// For example, the strings "010" and "1010" are alternating, while the string "0100" is not.
// Example 1:
//   Input: s = "111000"
//   Output: 2
//   Explanation: Use the first operation two times to make s = "100011".
//     Then, use the second operation on the third and sixth elements to make s = "101010".
// Example 2:
//   Input: s = "010"
//   Output: 0
//   Explanation: The string is already alternating.
// Example 3:
//   Input: s = "1110"
//   Output: 1
//   Explanation: Use the second operation on the second element to make s = "1010".
// Constraints:
//   1 <= s.length <= 10^5
//   s[i] is either '0' or '1'.

// let ss=s+s to simulate circular shift, then use sliding window to check
// whether use '010101...' or '101010...' as target alternating string
// e.g., s="1110", ss="11101110"
//         target1     10101010
//         target2     01010101
//         window1     ----
//         window2      ----
func minFlips(s string) int {
	ss := s + s
	ans, zero, one := len(s), 0, 0
	for i := 0; i < len(ss)-1; i++ {
		if i%2 == 0 {
			if ss[i] == '0' {
				zero++
			} else {
				one++
			}
		} else {
			if ss[i] == '0' {
				one++
			} else {
				zero++
			}
		}
		if i >= len(s) { // reduce former position
			j := i - len(s)
			if j%2 == 0 {
				if ss[j] == '0' {
					zero--
				} else {
					one--
				}
			} else {
				if ss[j] == '0' {
					one--
				} else {
					zero--
				}
			}
		}
		if i >= len(s)-1 { // check min
			if zero < ans {
				ans = zero
			}
			if one < ans {
				ans = one
			}
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		s   string
		ans int
	}{
		{"111000", 2},
		{"010", 0},
		{"1110", 1},
		{"01001001101", 2},
	} {
		fmt.Println(minFlips(v.s), v.ans)
	}
}

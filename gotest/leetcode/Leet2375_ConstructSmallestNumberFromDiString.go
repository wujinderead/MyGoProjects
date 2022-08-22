package main

import "fmt"

// https://leetcode.com/problems/construct-smallest-number-from-di-string/

// You are given a 0-indexed string pattern of length n consisting of the
// characters 'I' meaning increasing and 'D' meaning decreasing.
// A 0-indexed string num of length n + 1 is created using the following conditions:
//   num consists of the digits '1' to '9', where each digit is used at most once.
//   If pattern[i] == 'I', then num[i] < num[i + 1].
//   If pattern[i] == 'D', then num[i] > num[i + 1].
// Return the lexicographically smallest possible string num that meets the conditions.
// Example 1:
//   Input: pattern = "IIIDIDDD"
//   Output: "123549876"
//   Explanation:
//     At indices 0, 1, 2, and 4 we must have that num[i] < num[i+1].
//     At indices 3, 5, 6, and 7 we must have that num[i] > num[i+1].
//     Some possible values of num are "245639871", "135749862", and "123849765".
//     It can be proven that "123549876" is the smallest possible num that meets the conditions.
//     Note that "123414321" is not possible because the digit '1' is used more than once.
// Example 2:
//   Input: pattern = "DDD"
//   Output: "4321"
//   Explanation:
//     Some possible values of num are "9876", "7321", and "8742".
//     It can be proven that "4321" is the smallest possible num that meets the conditions.
// Constraints:
//   1 <= pattern.length <= 8
//   pattern consists of only the letters 'I' and 'D'.

// reverse the numbers between I
func smallestNumber(pattern string) string {
	max := byte('0')
	chars := make([]byte, 0, len(pattern)+1)
	if pattern[0] == 'I' {
		chars = append(chars, '1')
		max = '1'
	}
	i := 0
	for i < len(pattern) {
		end := i
		for end+1 < len(pattern) && pattern[end+1] == pattern[i] {
			end++
		}
		l := end - i + 1
		// found a II segment
		if pattern[i] == 'I' {
			for x := 0; x < l-1; x++ {
				max++
				chars = append(chars, max)
			}
		} else { // found a DD segment
			max = max + byte(l) + 1
			for x := 0; x < l+1; x++ {
				chars = append(chars, max-byte(x))
			}
		}
		i = end + 1
	}
	if pattern[len(pattern)-1] == 'I' {
		chars = append(chars, max+1)
	}
	return string(chars)
}

func main() {
	for _, v := range []struct {
		pattern, ans string
	}{
		{"IIIDIDDD", "123549876"},
		{"IIIIIIII", "123456789"},
		{"IIDDIIDD", "125436987"},
		{"DDIIDIDI", "321465879"},
		{"IDIDDDID", "132765498"},
		{"DIDIIDDI", "214358769"},
		{"DDD", "4321"},
		{"DDDI", "43215"},
		{"IDDDII", "1543267"},
		{"IDDDIID", "15432687"},
		{"IDDIIDII", "143257689"},
	} {
		fmt.Println(smallestNumber(v.pattern), v.ans)
	}
}

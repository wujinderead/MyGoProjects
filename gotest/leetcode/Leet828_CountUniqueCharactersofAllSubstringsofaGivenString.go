package main

import "fmt"

// https://leetcode.com/problems/count-unique-characters-of-all-substrings-of-a-given-string/

// Let's define a function countUniqueChars(s) that returns the number of unique
// characters on s, for example if s = "LEETCODE" then "L", "T","C","O","D" are the
// unique characters since they appear only once in s, therefore countUniqueChars(s) = 5.
// On this problem given a string s we need to return the sum of countUniqueChars(t)
// where t is a substring of s. Notice that some substrings can be repeated so
// on this case you have to count the repeated ones too.
// Since the answer can be very large, return the answer modulo 10 ^ 9 + 7.
// Example 1:
//   Input: s = "ABC"
//   Output: 10
//   Explanation: All possible substrings are: "A","B","C","AB","BC" and "ABC".
//     Evey substring is composed with only unique letters.
//     Sum of lengths of all substring is 1 + 1 + 1 + 2 + 2 + 3 = 10
// Example 2:
//   Input: s = "ABA"
//   Output: 8
//   Explanation: The same as example 1, except countUniqueChars("ABA") = 1.
// Example 3:
//   Input: s = "LEETCODE"
//   Output: 92
// Constraints:
//   0 <= s.length <= 10^4
//   s contain upper-case English letters only.

func uniqueLetterString(s string) int {
	// native method, for n² substrings, calculate each countUniqueChars(t), we got n³ time.
	// for each substrings start at s[i], we calculate for s[i: i], s[i: i+1]... gradually, we got n² time.
	// smarter method: for each character, like 'A', how many substrings that contain EXACTLY ONE 'A'?
	// we just need to answer this question for 26 times.
	chs := make([][]int, 26)
	for i := range s {
		chs[int(s[i]-'A')] = append(chs[int(s[i]-'A')], i)
	}
	count := 0
	for i := range chs {
		if len(chs[i]) > 0 {
			fmt.Println(string('A' + i))
		}
		for j := range chs[i] {
			prev := -1
			if j-1 >= 0 {
				prev = chs[i][j-1]
			}
			next := len(s)
			if j+1 < len(chs[i]) {
				next = chs[i][j+1]
			}
			count += (chs[i][j] - prev - 1) + (next - chs[i][j] - 1) + 1 + (chs[i][j]-prev-1)*(next-chs[i][j]-1)
			count %= 1000000007
			fmt.Println(prev, chs[i][j], next, chs[i][j]-prev-1, next-chs[i][j]-1, count)
		}
	}
	return count
}

func main() {
	fmt.Println(uniqueLetterString("ABC"))
	fmt.Println(uniqueLetterString("ABA"))
	fmt.Println(uniqueLetterString("ABAC"))
	fmt.Println(uniqueLetterString("CABACC"))
	fmt.Println(uniqueLetterString("CABACC"))
	fmt.Println(uniqueLetterString("LEETCODE"))
}

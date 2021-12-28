package main

import "fmt"

// https://leetcode.com/problems/adding-spaces-to-a-string/

// You are given a 0-indexed string s and a 0-indexed integer array spaces that describes the
// indices in the original string where spaces will be added. Each space should be inserted before
// the character at the given index.
// For example, given s = "EnjoyYourCoffee" and spaces = [5, 9], we place spaces before 'Y' and 'C',
// which are at indices 5 and 9 respectively. Thus, we obtain "Enjoy Your Coffee".
// Return the modified string after the spaces have been added.
// Example 1:
//   Input: s = "LeetcodeHelpsMeLearn", spaces = [8,13,15]
//   Output: "Leetcode Helps Me Learn"
//   Explanation:
//     The indices 8, 13, and 15 correspond to the underlined characters in "LeetcodeHelpsMeLearn".
//     We then place spaces before those characters.
// Example 2:
//   Input: s = "icodeinpython", spaces = [1,5,7,9]
//   Output: "i code in py thon"
//   Explanation:
//     The indices 1, 5, 7, and 9 correspond to the underlined characters in "icodeinpython".
//     We then place spaces before those characters.
// Example 3:
//   Input: s = "spacing", spaces = [0,1,2,3,4,5,6]
//   Output: " s p a c i n g"
//   Explanation:
//     We are also able to place spaces before the first character of the string.
// Constraints:
//   1 <= s.length <= 3 * 10⁵
//   s consists only of lowercase and uppercase English letters.
//   1 <= spaces.length <= 3 * 10⁵
//   0 <= spaces[i] <= s.length - 1
//   All the values of spaces are strictly increasing.

func addSpaces(s string, spaces []int) string {
	ans := make([]byte, 0, len(s)+len(spaces))
	si := 0
	for i := range s {
		if si < len(spaces) && i == spaces[si] {
			ans = append(ans, ' ', s[i])
			si++
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

func main() {
	for _, v := range []struct {
		s   string
		ss  []int
		ans string
	}{
		{"LeetcodeHelpsMeLearn", []int{8, 13, 15}, "Leetcode Helps Me Learn"},
		{"icodeinpython", []int{1, 5, 7, 9}, "i code in py thon"},
		{"spacing", []int{0, 1, 2, 3, 4, 5, 6}, " s p a c i n g"},
	} {
		fmt.Println(addSpaces(v.s, v.ss) == v.ans)
	}
}

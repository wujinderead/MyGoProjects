package main

import "fmt"

// https://leetcode.com/problems/replace-the-substring-for-balanced-string/

// You are given a string containing only 4 kinds of characters 'Q', 'W', 'E' and 'R'.
// A string is said to be balanced if each of its characters appears n/4 times where
// n is the length of the string.
// Return the minimum length of the substring that can be replaced with any other
// string of the same length to make the original string s balanced.
// Return 0 if the string is already balanced.
// Example 1:
//   Input: s = "QWER"
//   Output: 0
//   Explanation: s is already balanced.
// Example 2:
//   Input: s = "QQWE"
//   Output: 1
//   Explanation: We need to replace a 'Q' to 'R', so that "RQWE" (or "QRWE") is balanced.
// Example 3:
//   Input: s = "QQQW"
//   Output: 2
//   Explanation: We can replace the first "QQ" to "ER".
// Example 4:
//   Input: s = "QQQQ"
//   Output: 3
//   Explanation: We can replace the last 3 'Q' to make s = "QWER".
// Constraints:
//   1 <= s.length <= 10^5
//   s.length is a multiple of 4
//   s contains only 'Q', 'W', 'E' and 'R'.

func balancedString(s string) int {
    // part s to three parts: ABC, if A+C is valid (QWER count less than n/4),
    // we only need to alter B. if we want B minimal, we should maximize A+C.
    limit := len(s)/4
    count := [26]int{}
    left := -1
    for left+1<len(s) && count[int(s[left+1]-'A')]+1<=limit {
    	count[int(s[left+1]-'A')]++
    	left++
	}
	right := len(s)
	for count[int(s[right-1]-'A')]+1 <= limit {
		count[int(s[right-1]-'A')]++
		right--
	}
	allmax := left+1+len(s)-right
	for i:=left; i>=0; i-- {
		count[int(s[i]-'A')]--   // remove a character in left part
		for right-1>=0 && count[int(s[right-1]-'A')]+1 <= limit {
			count[int(s[right-1]-'A')]++
			right--
		}
		allmax = max(allmax, i+len(s)-right)
	}
	return len(s)-allmax
}

func max(a, b int) int {
	if a>b {
		return a
	}
	return b
}

func main() {
	fmt.Println(balancedString("QWER"))
	fmt.Println(balancedString("QQWE"))
	fmt.Println(balancedString("QQQW"))
	fmt.Println(balancedString("QQQQ"))
	fmt.Println(balancedString("QQQWQQQW"))
}
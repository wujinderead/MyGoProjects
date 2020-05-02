package main

import "fmt"

// https://leetcode.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/

// A happy string is a string that:
//   consists only of letters of the set ['a', 'b', 'c'].
//   s[i] != s[i + 1] for all values of i from 1 to s.length - 1 (string is 1-indexed).
// For example, strings "abc", "ac", "b" and "abcbabcbcb" are all happy strings 
// and strings "aa", "baa" and "ababbc" are not happy strings.
// Given two integers n and k, consider a list of all happy strings of length n 
// sorted in lexicographical order.
// Return the kth string of this list or return an empty string if there are less
// than k happy strings of length n.
// Example 1:
//   Input: n = 1, k = 3
//   Output: "c"
//   Explanation: The list ["a", "b", "c"] contains all happy strings of length 1.
//     The third string is "c".
// Example 2:
//   Input: n = 1, k = 4
//   Output: ""
//   Explanation: There are only 3 happy strings of length 1.
// Example 3:
//   Input: n = 3, k = 9
//   Output: "cab"
//   Explanation: There are 12 different happy string of length 3 ["aba", "abc", "aca",
//     "acb", "bab", "bac", "bca", "bcb", "cab", "cac", "cba", "cbc"]. You will find
//     the 9th string = "cab"
// Example 4:
//   Input: n = 2, k = 7
//   Output: ""
// Example 5:
//   Input: n = 10, k = 100
//   Output: "abacbabacb"
// Constraints:
//   1 <= n <= 10
//   1 <= k <= 100

func getHappyString(nn int, k int) string {
    // the number of all possible happy string is 3*2^(n-1)
    n := (1<<uint(nn-1))*3
    if n<k {
    	return ""
	}
	buf := make([]byte, nn)
	if k<=n/3 {
		buf[0] = 'a'
	} else if k<=2*(n/3) {
		buf[0] = 'b'
		k = k-(n/3)
	} else {
		buf[0] = 'c'
		k = k-(n/3)*2
	}
	next(buf, k, 1, n/3)
	return string(buf)
}

func next(buf []byte, k, ind, n int) {
	if ind==len(buf) {
		return
	}
	if k<=n/2 {
		if buf[ind-1]=='a' {
			buf[ind] = 'b'
		}
		if buf[ind-1]=='b' || buf[ind-1]=='c' {
			buf[ind] = 'a'
		}
	} else {
		if buf[ind-1]=='c' {
			buf[ind] = 'b'
		}
		if buf[ind-1]=='a' || buf[ind-1]=='b' {
			buf[ind] = 'c'
		}
		k = k-n/2
	}
	next(buf, k, ind+1, n/2)
}

func main() {
	fmt.Println(getHappyString(1, 3))
	fmt.Println(getHappyString(1, 4))
	fmt.Println(getHappyString(3, 9))
	fmt.Println(getHappyString(2, 7))
	fmt.Println(getHappyString(10, 100))
	fmt.Println(getHappyString(10, 1))
	fmt.Println(getHappyString(10, 512))
	fmt.Println(getHappyString(10, 513))
	fmt.Println(getHappyString(10, 1024))
	fmt.Println(getHappyString(10, 1025))
}
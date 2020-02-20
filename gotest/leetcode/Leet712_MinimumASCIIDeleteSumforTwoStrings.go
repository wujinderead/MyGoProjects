package main

import "fmt"

// https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/

// Given two strings s1, s2, find the lowest ASCII sum of deleted characters to make two strings equal.
// Example 1:
//   Input: s1 = "sea", s2 = "eat"
//   Output: 231
//   Explanation:
//     Deleting "s" from "sea" adds the ASCII value of "s" (115) to the sum.
//     Deleting "t" from "eat" adds 116 to the sum.
//     At the end, both strings are equal, and 115 + 116 = 231 is the minimum sum possible to achieve this.
// Example 2:
//   Input: s1 = "delete", s2 = "leet"
//   Output: 403
//   Explanation:
//     Deleting "dee" from "delete" to turn the string into "let",
//     adds 100[d]+101[e]+101[e] to the sum.  Deleting "e" from "leet" adds 101[e] to the sum.
//     At the end, both strings are equal to "let", and the answer is 100+101+101+101 = 403.
//     If instead we turned both strings into "lee" or "eet", we would get answers of 433 or 417,
//     which are higher.
// Note:
//   0 < s1.length, s2.length <= 1000.
//   All elements of each string will have an ASCII value in [97, 122].

// similar to the edit distance problem
// let pa[i] be the prefix of string a with length i, i.e, pa[0]="", pa[1]=a[0...0], pa[2]=a[0...1]
// let mds(i, j) be the minimal delete sum of pa[i] and pb[j], then
// if i*j==0, mds(i, j)=pa[i]+pb[j]
// if a[i-1]==b[j-1], no need to delete, mds(i, j)=mds(i-1, j-1)
// if a[i-1]!=b[j-1], mds(i, j)=min(mds(i-1, j)+a[i-1], mds(i, j-1)+b[j-1])
// for example
// a="se"  b="et"
//    i      i      i-1  i
//   se     se        s  e
//   et     e  t      et       mds(se, et)=min(mds(se, e)+t, mds(s, et)+e)
//    j   j-1  j       j
func minimumDeleteSum(a string, b string) int {
	if len(a)*len(b) == 0 {
		return len(a) + len(b)
	}
	if len(a) < len(b) {
		a, b = b, a // let b be shorter
	}
	prev, cur := make([]int, len(b)+1), make([]int, len(b)+1)
	// setup line 0
	for j := 1; j <= len(b); j++ {
		prev[j] = prev[j-1] + int(b[j-1])
	}
	// dp
	for i := 1; i <= len(a); i++ {
		cur[0] = prev[0] + int(a[i-1])
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cur[j] = prev[j-1]
			} else {
				cur[j] = min(prev[j]+int(a[i-1]), cur[j-1]+int(b[j-1]))
			}
		}
		prev, cur = cur, prev
	}
	return prev[len(b)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumDeleteSum("delete", "leet"))
	fmt.Println(minimumDeleteSum("sea", "eat"))
	fmt.Println(minimumDeleteSum("eat", "sea"))
	fmt.Println(minimumDeleteSum("ccacc", "sarc"))
	fmt.Println(minimumDeleteSum("abc", "b"))
}

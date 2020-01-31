package main

import "fmt"

// https://leetcode.com/problems/shortest-common-supersequence/

// Given two strings str1 and str2, return the shortest string that has both str1 and str2
// as subsequences. If multiple answers exist, you may return any of them. (A string S is a subsequence
// of string T if deleting some number of characters from T (possibly 0, and the characters
// are chosen anywhere from T) results in the string S.)
// Example 1:
//   Input: str1 = "abac", str2 = "cab"
//   Output: "cabac"
//   Explanation:
//     str1 = "abac" is a subsequence of "cabac" because we can delete the first "c".
//     str2 = "cab" is a subsequence of "cabac" because we can delete the last "ac".
//     The answer provided is the shortest such string that satisfies these properties.
// Note:
//   1 <= str1.length, str2.length <= 1000
//   str1 and str2 consist of lowercase English letters.

// first find the longest common subsequence, then the shortest common supersequence
// is the lcs plus those chars not in lcs. for example:
// a="delete", b="leet"
//   lcs=    LET           LEE           EET
//
//     a=   deLE Te      deLEtE         dElETe
//     b=     LEeT         LE Et        lE ET
//   scs=   deLEeTe      deLEtEt       dlElETe
//
// for either lcs, it got the scs with same length: len(a)+len(b)-lcs
func shortestCommonSupersequence(a string, b string) string {
	if len(a)*len(b) == 0 {
		return a + b
	}

	dp := make([][]int, len(a)+1)
	for i := range dp {
		dp[i] = make([]int, len(b)+1)
	}

	// dp
	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	lcslen := dp[len(a)][len(b)]
	if lcslen == 0 {
		return a + b // no lcs, a+b is the answer
	}
	scsbuf := make([]byte, len(a)+len(b)-lcslen)
	lcsbuf := make([][2]int, lcslen)

	// get lcs
	ind := lcslen - 1
	i := len(a)
	j := len(b)
	for ind >= 0 {
		if a[i-1] == b[j-1] {
			lcsbuf[ind][0] = i - 1
			lcsbuf[ind][1] = j - 1
			i--
			j--
			ind--
		} else if dp[i-1][j] == dp[i][j] {
			i--
		} else {
			j--
		}
	}

	// construct scs
	ind = 0
	// chars before lcs[0]
	for i := 0; i < lcsbuf[0][0]; i++ {
		scsbuf[ind] = a[i]
		ind++
	}
	for i := 0; i < lcsbuf[0][1]; i++ {
		scsbuf[ind] = b[i]
		ind++
	}
	// chars between lcs[0] lcs[lcslen-1]
	for i := 0; i < lcslen-1; i++ {
		scsbuf[ind] = a[lcsbuf[i][0]] // set lcs
		ind++
		for j := lcsbuf[i][0] + 1; j < lcsbuf[i+1][0]; j++ {
			scsbuf[ind] = a[j]
			ind++
		}
		for j := lcsbuf[i][1] + 1; j < lcsbuf[i+1][1]; j++ {
			scsbuf[ind] = b[j]
			ind++
		}
	}
	// lcs[lcslen-1] and chars after lcs[lcslen-1]
	scsbuf[ind] = a[lcsbuf[lcslen-1][0]]
	ind++
	for j := lcsbuf[lcslen-1][0] + 1; j < len(a); j++ {
		scsbuf[ind] = a[j]
		ind++
	}
	for j := lcsbuf[lcslen-1][1] + 1; j < len(b); j++ {
		scsbuf[ind] = b[j]
		ind++
	}
	return string(scsbuf)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(shortestCommonSupersequence("abaca", "cabdaf"))
	fmt.Println(shortestCommonSupersequence("delete", "leet"))
	fmt.Println(shortestCommonSupersequence("ccacc", "sarc"))
	fmt.Println(shortestCommonSupersequence("ab", "a"))
	fmt.Println(shortestCommonSupersequence("acd", "bce"))
	fmt.Println(shortestCommonSupersequence("acdfg", "bcefh"))

}

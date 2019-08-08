package main

import (
	"fmt"
	"math"
)

// EDIT: actually there is an O(n²) time, O(n) space solution.

// with O(n²) time and O(n) space of pre-processing,
// we can check whether a substring is palindromic in O(1) time.

// we start from the suffix of s,
// let mc[i] be the min cut for s[i...len(s)-1], then we add s[i-1],
// for mc[i-1], the suffix is [i-1, i, ..., k, k+1, ..., len(s)-1],
// we then iterate k, if s[i-1...k] is palindromic, then min cut will be 1+mc[k+1].
// this only need with O(n²) time and O(n) space

// https://leetcode.com/problems/palindrome-partitioning-ii/
// Given a string s, partition s such that every substring of the partition is a palindrome.
// Return the minimum cuts needed for a palindrome partitioning of s.
func minCut(s string) int {
	// let mc(i, j) be the minimal cut for s[i...j]
	// if s[i...j] is palindromic, no need to cut, mc(i, j)=0
	// else mc(i, j)=min( mc(i, i)+1+mc(i+1, j), min(i, i+1)+1+min(i+2, j) ...)
	// base case mc(i,i)=0; mc(i,i+1)=0 if s[i]==s[i+1] else 1
	// time O(n³), space O(n²)
	if len(s) < 2 {
		return 0
	}
	// store longest palindrome for each center
	odd := make([]int, len(s))
	even := make([]int, len(s)-1)
	for i := 0; i < len(s); i++ {
		l := 0
		for i-l >= 0 && i+l < len(s) && s[i-l] == s[i+l] {
			l++
		}
		odd[i] = 2*l - 1
	}
	for i := 0; i < len(s)-1; i++ {
		l := 0
		for i-l >= 0 && i+1+l < len(s) && s[i-l] == s[i+1+l] {
			l++
		}
		even[i] = 2 * l
	}
	fmt.Println(odd)
	fmt.Println(even)

	// initial mc
	mc := make([][]int, len(s))
	for i := range mc {
		mc[i] = make([]int, len(s))
	}
	for i := 0; i < len(s)-1; i++ {
		mc[i][i] = 0   // 0 cut to part "a"
		mc[i][i+1] = 1 // 1 cut to part "ab"
		if s[i] == s[i+1] {
			mc[i][i+1] = 0 // 0 cut to part "aa"
		}
	}
	mc[len(s)-1][len(s)-1] = 0

	// dp, update diagonally
	for diff := 2; diff < len(s); diff++ {
		for i := 0; i+diff < len(s); i++ {
			j := i + diff
			// fast path: s[i...j] is palindromic
			if isPalindromic(i, j, odd, even) {
				mc[i][j] = 0
				continue
			}
			// need iterate sub-problem
			mc[i][j] = math.MaxInt64
			for k := 0; i+k < j; k++ {
				if mc[i][i+k]+1+mc[i+k+1][j] < mc[i][j] {
					mc[i][j] = mc[i][i+k] + 1 + mc[i+k+1][j]
				}
			}
		}
	}
	return mc[0][len(s)-1]
}

func isPalindromic(i, j int, odd, even []int) bool {
	if (j-i+1)%2 == 0 { // even length
		return j-i+1 <= even[(i+j)/2]
	} else { // odd length
		return j-i+1 <= odd[(i+j)/2]
	}
}

func main() {
	s := "xabaabayz"
	odd := make([]int, len(s))
	even := make([]int, len(s)-1)
	for i := 0; i < len(s); i++ {
		l := 0
		for i-l >= 0 && i+l < len(s) && s[i-l] == s[i+l] {
			l++
		}
		odd[i] = 2*l - 1
	}
	for i := 0; i < len(s)-1; i++ {
		l := 0
		for i-l >= 0 && i+1+l < len(s) && s[i-l] == s[i+1+l] {
			l++
		}
		even[i] = 2 * l
	}
	fmt.Println(minCut("xabaabayz"))
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			fmt.Println(s[i:j+1], isPalindromic(i, j, odd, even))
		}

	}
	//fmt.Println(minCut(""))
	//fmt.Println(minCut("a"))
	//fmt.Println(minCut("ab"))
	//fmt.Println(minCut("aa"))
	//fmt.Println(minCut("aab"))
	//fmt.Println(minCut("aba"))
	//fmt.Println(minCut("aaa"))
	//fmt.Println(minCut("baa"))
}

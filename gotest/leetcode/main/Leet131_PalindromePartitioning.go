package main

import (
	"fmt"
)

// https://leetcode.com/problems/palindrome-partitioning/
// Given a string s, partition s such that every substring of the partition is a palindrome.
// Return all possible palindrome partitioning of s.

// use dp to get the number of partitions for s[i...j]; but if we want get all solutions,
// we still need to use dfs which will encounter duplicated sub-problems. we can store the
// solutions to avoid repeating calculation.
func partition(s string) [][]string {
	if s == "" {
		return [][]string{{""}}
	}
	pld := make([][]bool, len(s))
	num := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		pld[i] = make([]bool, len(s))
		num[i] = make([]int, len(s))
	}
	for i := 0; i < len(s)-1; i++ {
		pld[i][i] = true
		num[i][i] = 1
		num[i][i+1] = 1 // 1 way to part "ab"
		if s[i] == s[i+1] {
			pld[i][i+1] = true
			num[i][i+1] = 2 // 2 ways to part "aa"
		}
	}
	pld[len(s)-1][len(s)-1] = true
	num[len(s)-1][len(s)-1] = 1

	// dp
	for diff := 2; diff < len(s); diff++ {
		for i := 0; i+diff < len(s); i++ {
			j := i + diff
			// get if s[i...j] is palindromic
			pld[i][j] = s[i] == s[j] && pld[i+1][j-1]
			// if s[i...j] is palindromic, one way to part
			if pld[i][j] {
				num[i][j] = 1
			}
			for k := i; k < j; k++ {
				// we part s[i...j] as two part, if s[i...k] is palindromic,
				// only need to find the number of ways to part s[k+1...j]
				if pld[i][k] {
					num[i][j] += num[k+1][j]
				}
			}
		}
	}
	// generate partitions by step
	pps := make([][]string, num[0][len(s)-1])
	for i := range pps {
		pps[i] = make([]string, 0)
	}
	generatePartitionsHelper(s, pps, pld, num, 0, len(s), 0)
	return pps
}

func generatePartitionsHelper(s string, pps [][]string, pld [][]bool, num [][]int, start, lens, index int) {
	if start == lens {
		return
	}
	if pld[start][lens-1] {
		pps[index] = append(pps[index], s[start:lens])
		index++
	}
	for i := 0; start+i < lens-1; i++ {
		if pld[start][start+i] {
			for j := 0; j < num[start+i+1][lens-1]; j++ {
				pps[index+j] = append(pps[index+j], s[start:start+i+1])
			}
			generatePartitionsHelper(s, pps, pld, num, start+i+1, len(s), index)
			index += num[start+i+1][lens-1]
		}
	}
}

func main() {
	fmt.Println(partition("xabaabayz"))
	fmt.Println(partition(""))
	fmt.Println(partition("a"))
	fmt.Println(partition("ab"))
	fmt.Println(partition("aa"))
	fmt.Println(partition("aab"))
	fmt.Println(partition("aba"))
	fmt.Println(partition("aaa"))
	fmt.Println(partition("baa"))
}

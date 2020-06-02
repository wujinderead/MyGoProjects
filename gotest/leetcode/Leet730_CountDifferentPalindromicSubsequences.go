package main

import "fmt"

// https://leetcode.com/problems/count-different-palindromic-subsequences/

// Given a string S, find the number of different non-empty palindromic subsequences 
// in S, and return that number modulo 10^9 + 7.
// A subsequence of a string S is obtained by deleting 0 or more characters from S.
// A sequence is palindromic if it is equal to the sequence reversed.
// Two sequences A_1, A_2, ... and B_1, B_2, ... are different if there is some i for which A_i != B_i.
// Example 1:
//   Input: 
//     S = 'bccb'
//   Output: 6
//   Explanation: 
//     The 6 different non-empty palindromic subsequences are 'b', 'c', 'bb', 'cc', 'bcb', 'bccb'.
//     Note that 'bcb' is counted only once, even though it occurs twice.
// Example 2: 
//   Input: 
//   S = 'abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba'
//   Output: 104860361
//   Explanation: 
//     There are 3104860382 different non-empty palindromic subsequences, which is 104860361 modulo 10^9 + 7.
// Note:
//   The length of S will be in the range [1, 1000]. 
//   Each character S[i] will be in the set {'a', 'b', 'c', 'd'}. 

// dp diagonally, to get dp(i, j) for string S[i...j], think S[i...j] in the following form:
// i          j
// xxabbcxaxbxx
//   abbcxa              // the number of palinrome in a...a form is dp("bbcx")+dp("aa")+dp("a") 
//    bbcxaxb            // the number of palinrome in b...b form is dp("bcxaxb")+dp("bb")+dp("b")
//      c                // no palindrome in c...c mode, but dp("c") should add  
// to efficiently get the index of leftmost and rightmost 'a' in S[i...j], we use prefix and suffix arrays.
func countPalindromicSubsequences(S string) int {
    // key represent S[i...j], value means the count
    dp := make(map[[2]int]int) 

    // dp(i, i)
    for i:=0; i<len(S)-1; i++ {
    	dp[[2]int{i, i}] = 1
    	dp[[2]int{i, i+1}] = 2   // both aa or ab is 2
    }
    dp[[2]int{len(S)-1, len(S)-1}] = 1

    // to efficiently get the index of leftmost and rightmost 'a' in S[i...j], we use prefix and suffix arrays.
    left, right := make([][]int, 4), make([][]int, 4)
    for i := range left {
    	left[i] = make([]int, len(S))
    	right[i] = make([]int, len(S))
    	left[i][0] = -1
    	right[i][len(S)-1] = len(S)
    }
    left[int(S[0]-'a')][0] = 0
    right[int(S[len(S)-1]-'a')][len(S)-1] = len(S)-1

    for i:=1; i<len(S); i++ {
    	for j:=0; j<4; j++ {
    		left[j][i] = left[j][i-1]
    		if int(S[i]-'a')==j {
    			left[j][i] = i
    		}
    	}
    }

    for i:=len(S)-2; i>=0; i-- {
    	for j:=0; j<4; j++ {
    		right[j][i] = right[j][i+1]
    		if int(S[i]-'a')==j {
    			right[j][i] = i
    		}
    	}
    }

	// dp(i, i+diff)
    for diff:=2; diff<len(S); diff++ {    // update diagonally
    	for i:=0; i<len(S)-diff; i++ {
    		j := i+diff
    		count := 0
    		for k:=0; k<4; k++ {
    			l, r := right[k][i], left[k][j]   // the nearest char[k] right to i, left to j
    			if l>r {     // both out of range
    				continue
    			}
    			// both in range
    			if l==r || l+1==r {    // x or xx, add 1 or 2
    				count += r-l+1
    			} else {
    				count += dp[[2]int{l+1, r-1}] + 1 + 1   // axxxa, add dp("xxx") + ("aa") + ("a")
    			}
    		}
    		dp[[2]int{i, j}] = count % int(1e9+7)
    	}
    }
    return dp[[2]int{0, len(S)-1}]
}

func main() {
	fmt.Println(countPalindromicSubsequences("a"))
	fmt.Println(countPalindromicSubsequences("aa"))
	fmt.Println(countPalindromicSubsequences("aaa"))
	fmt.Println(countPalindromicSubsequences("ab"))
	fmt.Println(countPalindromicSubsequences("bcc"))
	fmt.Println(countPalindromicSubsequences("bccb"))
	fmt.Println(countPalindromicSubsequences("babb"))
	fmt.Println(countPalindromicSubsequences("babba"))
	fmt.Println(countPalindromicSubsequences("babbabba"))
	fmt.Println(countPalindromicSubsequences("abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba"))
}
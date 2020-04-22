package main

import (
	"fmt"
	"math/big"
)

// https://leetcode.com/problems/distinct-echo-substrings/

// Return the number of distinct non-empty substrings of text that can be written
// as the concatenation of some string with itself (i.e. it can be written as
// a + a where a is some string).
// Example 1:
//   Input: text = "abcabcabc"
//   Output: 3
//   Explanation: The 3 substrings are "abcabc", "bcabca" and "cabcab".
// Example 2:
//   Input: text = "leetcodeleetcode"
//   Output: 2
//   Explanation: The 2 substrings are "ee" and "leetcodeleetcode".
// Constraints:
//   1 <= text.length <= 2000
//   text has only lowercase English letters.

func distinctEchoSubstrings(text string) int {
    // use rabin-karp to search substrings
    set := make(map[string]struct{})
    p := 1000000007
    ainv := getPinv(26, p)
    for l:=len(text)/2; l>0; l-- {
    	// compare text[i: i+l] and text[i+l: i+2l]
    	h1, h2 := 0, 0
    	pow := ainv
		for j:=0; j<l; j++ {
			h1 = (h1*26 + int(text[j]-'a')) % p
			pow = (pow*26) % p
		}
		for j:=l; j<2*l; j++ {
			h2 = (h2*26 + int(text[j]-'a')) % p
		}
		if h1 == h2 && text[:l] == text[l: 2*l] {
			set[text[:l]] = struct{}{}
		}
		for i:=1; i+2*l<=len(text); i++ {
			h1 = (h1 - int(text[i-1]-'a')*pow) % p
			if h1 < 0 {
				h1 += p
			}
			h1 = (h1*26 + int(text[i+l-1]-'a')) % p

			h2 = (h2 - int(text[i+l-1]-'a')*pow) % p
			if h2 < 0 {
				h2 += p
			}
			h2 = (h2*26 + int(text[i+2*l-1]-'a')) % p

			if h1==h2 && text[i: i+l] == text[i+l: i+2*l] {
				set[text[i: i+l]] = struct{}{}
			}
		}
	}
    return len(set)
}

func getPinv(a, p int) int {
	aa := new(big.Int).SetInt64(int64(a))
	pp := new(big.Int).SetInt64(int64(p))
	aa.ModInverse(aa, pp)
	return int(aa.Int64())
}

func main() {
	fmt.Println(distinctEchoSubstrings("abcabcabc"))
	fmt.Println(distinctEchoSubstrings("leetcodeleetcode"))
	fmt.Println(distinctEchoSubstrings("a"))
	fmt.Println(distinctEchoSubstrings("aa"))
	fmt.Println(distinctEchoSubstrings("aaa"))
	fmt.Println(distinctEchoSubstrings("aaaa"))
	fmt.Println(distinctEchoSubstrings("aaaaa"))
	fmt.Println(distinctEchoSubstrings("aaaaaa"))
}
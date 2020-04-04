package main

import (
	"fmt"
	"math/big"
)

// A string is called a happy prefix if is a non-empty prefix which is also
// a suffix (excluding itself).
// Given a string s. Return the longest happy prefix of s .
// Return an empty string if no such prefix exists.
// Example 1:
//   Input: s = "level"
//   Output: "l"
//   Explanation: s contains 4 prefix excluding itself ("l", "le", "lev", "leve"),
//     and suffix ("l", "el", "vel", "evel"). The largest prefix which is also suffix
//     is given by "l".
// Example 2:
//   Input: s = "ababab"
//   Output: "abab"
//   Explanation: "abab" is the largest prefix which is also suffix.
//     They can overlap in the original string.
// Example 3:
//   Input: s = "leetcodeleet"
//   Output: "leet"
// Example 4:
//   Input: s = "a"
//   Output: ""
// Constraints:
//   1 <= s.length <= 10^5
//   s contains only lowercase English letters.

func longestPrefix(s string) string {
	// use rabin-karp to check if two strings are identical
	if len(s) == 1 {
		return ""
	}
	P := 211                    // prime for modular
	ainv := modInverse(26, P)   // ainv = 26^-1 mod P
	pows := make([]int, len(s)) // pows[i] = (26^i) mod P
	pows[0] = 1
	hash := int(s[0] - 'a') // e.g. hash(abcd) = a*1 + b*26 + c*26² + d*26³
	for i := 1; i < len(s); i++ {
		pows[i] = (pows[i-1] * 26) % P
		hash = (hash + int(s[i]-'a')*pows[i]) % P
	}
	prefix := hash
	suffix := hash
	for i := 0; i < len(s)-1; i++ {
		j := len(s) - 1 - i                            // suffix need to exclude s[i], prefix to exclude s[j]
		suffix = ((suffix - int(s[i]-'a')) * ainv) % P // suffix = (suffix - s[i])/26
		prefix = (prefix - int(s[j]-'a')*pows[j]) % P  // prefix = prefix - s[j]*pows[j]
		if suffix < 0 {
			suffix += P
		}
		if prefix < 0 {
			prefix += P
		}
		if prefix == suffix {
			if s[i+1:] == s[:j] {
				return s[:j]
			}
		}
	}
	return ""
}

func modInverse(a, p int) int {
	aa := new(big.Int).SetInt64(int64(a))
	pp := new(big.Int).SetInt64(int64(p))
	return int(aa.ModInverse(aa, pp).Int64())
}

func main() {
	fmt.Println(longestPrefix("abcde"))
	fmt.Println(longestPrefix("level"))
	fmt.Println(longestPrefix("ababab"))
	fmt.Println(longestPrefix("rrrrrr"))
	fmt.Println(longestPrefix("leetcodeleet"))
	fmt.Println(longestPrefix("gg"))
	fmt.Println(longestPrefix("gh"))
}

func gethash(s string) int {
	P := 211                    // prime for modular
	pows := make([]int, len(s)) // pows[i] = (26^i) mod P
	pows[0] = 1
	hash := int(s[0] - 'a') // e.g. hash(abcd) = a*1 + b*26 + c*26² + d*26³
	for i := 1; i < len(s); i++ {
		pows[i] = (pows[i-1] * 26) % P
		hash = (hash + int(s[i]-'a')*pows[i]) % P
	}
	return hash
}

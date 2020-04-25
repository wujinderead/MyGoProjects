package main

import "fmt"

// https://leetcode.com/problems/longest-chunked-palindrome-decomposition/

// Return the largest possible k such that there exists a_1, a_2, ..., a_k such that:
//   Each a_i is a non-empty string;
//   Their concatenation a_1 + a_2 + ... + a_k is equal to text;
//   For all 1 <= i <= k, a_i = a_{k+1 - i}.
// Example 1:
//   Input: text = "ghiabcdefhelloadamhelloabcdefghi"
//   Output: 7
//   Explanation: We can split the string on "(ghi)(abcdef)(hello)(adam)(hello)(abcdef)(ghi)".
// Example 2:
//   Input: text = "merchant"
//   Output: 1
//   Explanation: We can split the string on "(merchant)".
// Example 3:
//   Input: text = "antaprezatepzapreanta"
//   Output: 11
//   Explanation: We can split the string on "(a)(nt)(a)(pre)(za)(tpe)(za)(pre)(a)(nt)(a)".
// Example 4:
//   Input: text = "aaa"
//   Output: 3
//   Explanation: We can split the string on "(a)(a)(a)".
// Constraints:
//   text consists only of lowercase English characters.
//   1 <= text.length <= 1000

func longestDecomposition(text string) int {
	i := 0
	j := len(text)-1
	lh := 0
	rh := 0
	pow := 1
	p := 1000000007
	k := 0
	leng := 0
    for i<j {
		lh = (lh*26+int(text[i]-'a')) % p
		rh = (rh + int(text[j]-'a')*pow) % p
		pow = (pow * 26) % p
		leng++
		if lh == rh {
			if text[i-leng+1: i+1] == text[j: j+leng] {
				//fmt.Println(text[i-leng+1: i+1], text[j: j+leng])
				k += 2
				if i+1==j {
					k--
				}
				lh, rh, leng, pow = 0, 0, 0, 1
			}
		}
		i++
		j--
	}
	k++
	return k
}

func main() {
    fmt.Println(longestDecomposition("ghiabcdefhelloadamhelloabcdefghi"))
    fmt.Println(longestDecomposition("merchant"))
    fmt.Println(longestDecomposition("antaprezatepzapreanta"))
    fmt.Println(longestDecomposition("a"))
    fmt.Println(longestDecomposition("aa"))
    fmt.Println(longestDecomposition("aaa"))
    fmt.Println(longestDecomposition("aaaa"))
    fmt.Println(longestDecomposition("apattxyyttxapa"))
}
package main

import "fmt"

// https://leetcode.com/problems/decode-ways-ii/

// A message containing letters from A-Z is being encoded to numbers using the following mapping way:
//   'A' -> 1
//   'B' -> 2
//   ...
//   'Z' -> 26
// Beyond that, now the encoded string can also contain the character '*', which can be treated as one 
// of the numbers from 1 to 9. Given the encoded message containing digits and the character '*', 
// return the total number of ways to decode it. Also, since the answer may be very large, you 
// should return the output mod 10^9 + 7.
// Example 1:
//   Input: "*"
//   Output: 9
//   Explanation: The encoded message can be decoded to the string: "A", "B", "C", "D", "E", "F", "G", "H", "I".
// Example 2:
//   Input: "1*"
//   Output: 9 + 9 = 18
// Note:
//   The length of the input string will fit in range [1, 10^5].
//   The input string will only contain the character '*' and digits '0' - '9'.

func numDecodings(s string) int {
	if s[0]=='0' {
		return 0   // invalid
	}
	// let dp[i] be the number of decode ways for s[:i]. dp[i] depends on dp[i-1], dp[i-2].
	dp := make([]int, len(s)+1)
	dp[0] = 1         // 1 way for ""
	dp[1] = 1
	if s[0]=='*' {    // for s[0]
		dp[1] = 9     
	}
	for i:=2; i<=len(s); i++ {       // dp[i] stands for s[0...i-1]
		cur, prev := s[i-1], s[i-2] 

		// deem current char as single
		if cur=='*' {              // xxxx,1-9
			dp[i] = dp[i-1]*9
		} else if cur != '0' {     // xxxx,1-9
			dp[i] = dp[i-1]
		}

		// connect current char with prev char
		if prev=='1' || prev=='*' {   // see prev as 1
			if cur=='*' {
				dp[i] += dp[i-2]*9    // xxxx,11-19
			} else {
				dp[i] += dp[i-2]      // xxxx,10-19
			}
		}
		if prev=='2' || prev=='*' {   // see prev as 2
			if cur=='*' {
				dp[i] += dp[i-2]*6    // xxxx,21-26
			} else if cur>='0' && cur<='6' {
				dp[i] += dp[i-2]      // xxxx,20-26
			}
		}
		dp[i] = dp[i] % int(1e9+7)
	}
	return dp[len(s)]
}

func main() {
	for _, s := range []string { "*", "**", "1*", "0", "90", "2*", "*2", "***", "******", 
		"*2*", "*8*", "1812924*7", "0*1*8", "10", "20", "30"} {
		fmt.Println(s, numDecodings(s))
	}
}

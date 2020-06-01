package main

import "fmt"

// https://leetcode.com/problems/check-if-a-string-contains-all-binary-codes-of-size-k/

// Given a binary string s and an integer k. 
// Return True if all binary codes of length k is a substring of s. Otherwise, return False. 
// Example 1: 
//   Input: s = "00110110", k = 2
//   Output: true
//   Explanation: The binary codes of length 2 are "00", "01", "10" and "11". They 
//     can be all found as substrings at indicies 0, 1, 3 and 2 respectively.
// Example 2: 
//   Input: s = "00110", k = 2
//   Output: true
// Example 3: 
//   Input: s = "0110", k = 1
//   Output: true
//   Explanation: The binary codes of length 1 are "0" and "1", it is clear that both exist as a substring. 
// Example 4:  
//   Input: s = "0110", k = 2
//   Output: false
//   Explanation: The binary code "00" is of length 2 and doesn't exist in the array.
// Example 5:  
//   Input: s = "0000000001011100", k = 4
//   Output: false
// Constraints: 
//   1 <= s.length <= 5 * 10^5 
//   s consists of 0's and 1's only. 
//   1 <= k <= 20 

func hasAllCodes(s string, k int) bool {
	if len(s)<k {
		return false
	}
    n := (1<<uint(k))/8            // use bitmap to save some space
    if n==0 {
    	n = 1
    }
    mask := make([]uint64, n)
    n = 0
    for i:=0; i<k; i++ {
    	n = (n<<1)+int(s[i]-'0')
    }
    mask[n/8] = mask[n/8] | (1<<uint(n%8))
    for i:=k; i<len(s); i++ {
    	n = ((n<<1)+int(s[i]-'0')) & ((1<<uint(k))-1)
    	mask[n/8] = mask[n/8] | (1<<uint(n%8))
    }
    if 1<<uint(k) < 8 {
    	return mask[0] == (1<<uint(1<<uint(k)))-1
    }
    for i := range mask {
    	if mask[i] != 0xff {
    		return false
    	}
    }
    return true
}

func main() {
	fmt.Println(hasAllCodes("00110110", 2), true)
	fmt.Println(hasAllCodes("00110", 2), true)
	fmt.Println(hasAllCodes("0110", 1), true)
	fmt.Println(hasAllCodes("0110", 2), false)
	fmt.Println(hasAllCodes("0000000001011100", 4), false)
	fmt.Println(hasAllCodes("0101", 13), false)
}
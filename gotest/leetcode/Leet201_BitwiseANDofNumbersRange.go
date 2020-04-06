package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/bitwise-and-of-numbers-range/

// Given a range [m, n] where 0 <= m <= n <= 2147483647, return the
// bitwise AND of all numbers in this range, inclusive.
// Example 1:
//  Input: [5,7]
//  Output: 4
// Example 2:
//  Input: [0,1]
//  Output: 0 Related Topics Bit Manipulation

// if m=0xxxxx, n=1xxxxx, then 0xxxxx < 011111 < 100000 < 1xxxxx, the answer is 0;
// if m=10xxxx, n=11xxxx, then 10xxxx < 101111 < 110000 < 11xxxx, the answer is 100000;
// if m=1110xxxx, n=1111xxxx, the answer is 11100000
func rangeBitwiseAnd(m int, n int) int {
    if n==0 {
    	return 0
	}
    maxbit := 0
    for i:=0; i<32; i++ {
    	mask := 1<<uint(i)
    	if n&mask == mask {
    		maxbit = i
		}
	}
    ans := 0
    for i:=maxbit; i>=0; i-- {   // test every bit of m, n
    	mask := 1<<uint(i)
    	if m&mask==mask && n&mask==mask {  // both bits are 1
			ans |= mask      // set bit 1
		}
    	if m&mask==0 && n&mask==mask {   // m is 0 and n is 1
    		break
		}
    	// both 0, continue
	}
    return ans
}

func main() {
 	fmt.Println(rangeBitwiseAnd(0, 1))
 	fmt.Println(rangeBitwiseAnd(5, 7))
 	a, _ := strconv.ParseInt("11011011010", 2, 64)
 	b, _ := strconv.ParseInt("11011101110", 2, 64)
	fmt.Println(strconv.FormatInt(int64(rangeBitwiseAnd(int(a), int(b))), 2))
	a, _ = strconv.ParseInt("01011011010", 2, 64)
	b, _ = strconv.ParseInt("11011101110", 2, 64)
	fmt.Println(strconv.FormatInt(int64(rangeBitwiseAnd(int(a), int(b))), 2))
	a, _ = strconv.ParseInt("1101011010", 2, 64)
	b, _ = strconv.ParseInt("1110110111", 2, 64)
	fmt.Println(strconv.FormatInt(int64(rangeBitwiseAnd(int(a), int(b))), 2))
}
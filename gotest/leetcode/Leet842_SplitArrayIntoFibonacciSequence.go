package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/split-array-into-fibonacci-sequence/

// Given a string S of digits, such as S = "123456579", we can split it into a
// Fibonacci-like sequence [123, 456, 579].
// Formally, a Fibonacci-like sequence is a list F of non-negative integers such that:
//   0 <= F[i] <= 2^31 - 1, (that is, each integer fits a 32-bit signed integer type);
//   F.length >= 3;
//   and F[i] + F[i+1] = F[i+2] for all 0 <= i < F.length - 2.
// Also, note that when splitting the string into pieces, each piece must not have extra
// leading zeroes, except if the piece is the number 0 itself.
// Return any Fibonacci-like sequence split from S, or return [] if it cannot be done.
// Example 1:
//   Input: "123456579"
//   Output: [123,456,579]
// Example 2:
//   Input: "11235813"
//   Output: [1,1,2,3,5,8,13]
// Example 3:
//   Input: "112358130"
//   Output: []
//   Explanation: The task is impossible.
// Example 4:
//   Input: "0123"
//   Output: []
//   Explanation: Leading zeroes are not allowed, so "01", "2", "3" is not valid.
// Example 5:
//   Input: "1101111"
//   Output: [110, 1, 111]
//   Explanation: The output [11, 0, 11, 11] would also be accepted.
// Note:
//   1 <= S.length <= 200
//   S contains only digits.

func splitIntoFibonacci(S string) []int {
    // the first two numbers can determine a fibonacci sequence
    max := (1<<31)-1
    buf := make([]int, 0, 10)
    first := 0
	for i:=0; i<len(S)-1; i++ {
		first = 10*first+int(S[i]-'0')
		if first>max {
			break
		}
		second := 0
		buf = append(buf, first)
		for j:=i+1; j<len(S); j++ {
			second = second*10+int(S[j]-'0')
			if second>max {
				break
			}
			buf = append(buf, second)
			prev, cur := second, first+second
			st:=j+1
			//fmt.Println(first, second)
			for {
				if cur>max {
					break
				}
				str := strconv.Itoa(cur)
				if st+len(str)<=len(S) && S[st: st+len(str)]==str {
					buf = append(buf, cur)
					if st+len(str)==len(S) {
						return buf
					}
					//fmt.Println("third:", cur)
					prev, cur = cur, cur+prev
					st = st+len(str)
					continue
				}
				break
			}
			if second==0 {
				break
			}
			buf = buf[:1]
		}
		if first==0 {
			break
		}
		buf = buf[:0]
	}
	return buf[:0]
}

func main() {
	fmt.Println(splitIntoFibonacci("123456579"))
	fmt.Println(splitIntoFibonacci("11235813"))
	fmt.Println(splitIntoFibonacci("112358130"))
	fmt.Println(splitIntoFibonacci("0123"))
	fmt.Println(splitIntoFibonacci("1101111"))
	fmt.Println(splitIntoFibonacci("1101111112223"))
	fmt.Println(splitIntoFibonacci("1230123123246"))
	fmt.Println(splitIntoFibonacci("1230123124"))
	fmt.Println(splitIntoFibonacci("12"))
	fmt.Println(splitIntoFibonacci("123"))
	fmt.Println(splitIntoFibonacci("0123123"))
	fmt.Println(splitIntoFibonacci("539834657215398346785398346991079669377161950407626991734534318677529701785098211336528511"))
}
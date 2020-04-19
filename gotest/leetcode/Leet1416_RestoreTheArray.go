package main

import "fmt"

// https://leetcode.com/problems/restore-the-array/

// A program was supposed to print an array of integers. The program forgot to print
// whitespaces and the array is printed as a string of digits and all we know is that
// all integers in the array were in the range [1, k] and there are no leading
// zeros in the array.
// Given the string s and the integer k. There can be multiple ways to restore the array.
// Return the number of possible array that can be printed as a string s using the mentioned program.
// The number of ways could be very large so return it modulo 10^9 + 7
// Example 1:
//   Input: s = "1000", k = 10000
//   Output: 1
//   Explanation: The only possible array is [1000]
// Example 2:
//   Input: s = "1000", k = 10
//   Output: 0
//   Explanation: There cannot be an array that was printed this way and has all integer >= 1 and <= 10.
// Example 3:
//   Input: s = "1317", k = 2000
//   Output: 8
//   Explanation: Possible arrays are [1317],[131,7],[13,17],[1,317],[13,1,7],[1,31,7],[1,3,17],[1,3,1,7]
// Example 4:
//   Input: s = "2020", k = 30
//   Output: 1
//   Explanation: The only possible array is [20,20]. [2020] is invalid because 2020 > 30.
//     [2,020] is invalid because 020 contains leading zeros.
// Example 5:
//   Input: s = "1234567890", k = 90
//   Output: 34
// Constraints:
//   1 <= s.length <= 10^5.
//   s consists of only digits and doesn't contain leading zeros.
//   1 <= k <= 10^9.

func numberOfArrays(s string, k int) int {
	mod := 1000000007
	// for s="1317" and k=2000 as example. let f(s) be the number of ways, then:
	// for "7", we has 1 solution. F("7")=1
	// for "17", we has 1+f("7"), 17+f(""). F("17")=2
	// for "317", we has 3+f("17"), 31+f("7"), 317. F("317")=4
	// for "1317", we has 1+f("317"), 13+f("17"), 131+f("7"), 1317. F("1317")=1+1+2+4=8
	f := make([]int, len(s)+1) // f[i] is the number of ways for s[i:]
	f[len(s)] = 1
	for i := len(s) - 1; i >= 0; i-- {
		cur := 0
		for j := i; j < len(s); j++ {
			cur = cur*10 + int(s[j]-'0')
			if cur < 1 || cur > k {
				break
			}
			f[i] += f[j+1]
			f[i] %= mod // don't forget mod
		}
	}
	return f[0]
}

func main() {
	fmt.Println(numberOfArrays("1000", 10000))
	fmt.Println(numberOfArrays("1000", 10))
	fmt.Println(numberOfArrays("1317", 2000))
	fmt.Println(numberOfArrays("2020", 30))
	fmt.Println(numberOfArrays("1234567890", 90))
	fmt.Println(numberOfArrays("123456", 7))
	fmt.Println(numberOfArrays("123456", 5))
	fmt.Println(numberOfArrays("600342244431311113256628376226052681059918526204", 703))
}

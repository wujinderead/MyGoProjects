package main

import "fmt"

// https://leetcode.com/problems/minimize-xor/

// Given two positive integers num1 and num2, find the integer x such that:
//   x has the same number of set bits as num2, and
//   The value x XOR num1 is minimal.
// Note that XOR is the bitwise XOR operation.
// Return the integer x. The test cases are generated such that x is uniquely determined.
// The number of set bits of an integer is the number of 1's in its binary representation.
// Example 1:
//   Input: num1 = 3, num2 = 5
//   Output: 3
//   Explanation:
//     The binary representations of num1 and num2 are 0011 and 0101, respectively.
//     The integer 3 has the same number of set bits as num2, and the value 3 XOR 3 = 0 is minimal.
// Example 2:
//   Input: num1 = 1, num2 = 12
//   Output: 3
//   Explanation:
//     The binary representations of num1 and num2 are 0001 and 1100, respectively.
//     The integer 3 has the same number of set bits as num2, and the value 3 XOR 1 = 2 is minimal.
// Constraints:
//   1 <= num1, num2 <= 10⁹

func minimizeXor(num1 int, num2 int) int {
	var maxbit1, bit1, bit2 int
	for i := 0; i < 32; i++ {
		if (1<<i)&num1 > 0 {
			bit1++
			maxbit1 = i + 1
		}
		if (1<<i)&num2 > 0 {
			bit2++
		}
	}
	if bit2 <= bit1 {
		ans := 0
		for i := maxbit1 - 1; i >= 0 && bit2 > 0; i-- {
			if (1<<i)&num1 > 0 {
				ans = ans ^ (1<<i)&num1
				bit2--
			}
		}
		return ans
	} else if bit2 <= maxbit1 {
		ans := num1
		for i := 0; bit2-bit1 > 0 && i < 32; i++ {
			if (1<<i)&num1 == 0 {
				ans = ans ^ (1 << i)
				bit2--
			}
		}
		return ans
	}
	return (1 << bit2) - 1
}

func main() {
	for _, v := range []struct {
		n1, n2, ans int
	}{
		{0b100100111, 0b111, 0b100100100},
		{0b11, 0b101, 0b11},
		{0b100100111, 0b11111, 0b100100111},
		{0b100100111, 0b111111, 0b100101111},
		{0b100100111, 0b111111111, 0b111111111},
		{0b100100111, 0b1111111111, 0b1111111111},
	} {
		fmt.Println(minimizeXor(v.n1, v.n2), v.ans)
	}
}

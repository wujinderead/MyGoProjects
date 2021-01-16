package main

import "fmt"

// https://leetcode.com/problems/maximum-binary-string-after-change/

// You are given a binary string binary consisting of only 0's or 1's. You can apply
// each of the following operations any number of times:
//   Operation 1: If the number contains the substring "00", you can replace it with "10".
//   For example, "00010" -> "10010"
//   Operation 2: If the number contains the substring "10", you can replace it with "01".
//   For example, "00010" -> "00001"
// Return the maximum binary string you can obtain after any number of operations.
// Binary string x is greater than binary string y if x's decimal representation
// is greater than y's decimal representation.
// Example 1:
//   Input: binary = "000110"
//   Output: "111011"
//   Explanation: A valid transformation sequence can be:
//     "000110" -> "000101"
//     "000101" -> "100101"
//     "100101" -> "110101"
//     "110101" -> "110011"
//     "110011" -> "111011"
// Example 2:
//   Input: binary = "01"
//   Output: "01"
//   Explanation: "01" cannot be transformed any further.
// Constraints:
//   1 <= binary.length <= 105
//   binary consist of '0' and '1'.

// the leading ones are already good,
// for rest part, like 01010, always make operation 2 to make it as 00011
// and then make it 11011 by operation 1
// for example, binary=11101010, 3 leading ones, 3 zeros in rest part,
// first make it 11100011, then change first two 00 to 11,
// make it 11111011
func maximumBinaryString(binary string) string {
	buf := make([]byte, len(binary))
	for i := range binary {
		buf[i] = '1'
	}
	ones, zeros := 0, 0
	for i := 0; i < len(binary); i++ {
		if binary[i] == '0' {
			zeros++ // zeros is the number of 0s in rest part
		} else if zeros == 0 {
			ones++ // ones is the number of leading 1's
		}
	}
	if zeros != 0 {
		buf[ones+zeros-1] = '0'
	}
	return string(buf)
}

func main() {
	for _, v := range []struct {
		b, ans string
	}{
		{"000110", "111011"},
		{"01", "01"},
		{"0", "0"},
		{"00", "10"},
		{"11", "11"},
		{"10", "10"},
		{"110", "110"},
		{"010", "101"},
		{"0110", "1011"},
		{"0100", "1101"},
	} {
		fmt.Println(maximumBinaryString(v.b), v.ans)
	}
}

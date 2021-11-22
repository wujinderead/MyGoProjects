package main

import "fmt"

// https://leetcode.com/problems/decode-the-slanted-ciphertext/

// A string originalText is encoded using a slanted transposition cipher to a string encodedText
// with the help of a matrix having a fixed number of rows rows.
// originalText is placed first in a top-left to bottom-right manner.
// The blue cells are filled first, followed by the red cells, then the yellow cells, and so on,
// until we reach the end of originalText. The arrow indicates the order in which the cells are filled.
// All empty cells are filled with ' '. The number of columns is chosen such that the rightmost column
// will not be empty after filling in originalText.
// encodedText is then formed by appending all characters of the matrix in a row-wise fashion.
// The characters in the blue cells are appended first to encodedText, then the red cells, and so on,
// and finally the yellow cells. The arrow indicates the order in which the cells are accessed.
// For example, if originalText = "cipher" and rows = 3, then we encode it in the following manner:
// The blue arrows depict how originalText is placed in the matrix, and the red arrows denote the order
// in which encodedText is formed. In the above example, encodedText = "ch ie pr".
// Given the encoded string encodedText and number of rows rows, return the original string originalText.
// Note: originalText does not have any trailing spaces ' '. The test cases are generated such that
// there is only one possible originalText.
// Example 1:
//   Input: encodedText = "ch   ie   pr", rows = 3
//   Output: "cipher"
//   Explanation: This is the same example described in the problem description.
// Example 2:
//   Input: encodedText = "iveo    eed   l te   olc", rows = 4
//   Output: "i love leetcode"
//   Explanation: The figure above denotes the matrix that was used to encode originalText.
//     The blue arrows show how we can find originalText from encodedText.
// Example 3:
//   Input: encodedText = "coding", rows = 1
//   Output: "coding"
//   Explanation: Since there is only 1 row, both originalText and encodedText are the same.
// Example 4:
//   Input: encodedText = " b  ac", rows = 2
//   Output: " abc"
//   Explanation: originalText cannot have trailing spaces, but it may be preceded by one or more spaces.
// Constraints:
//   0 <= encodedText.length <= 10^6
//   encodedText consists of lowercase English letters and ' ' only.
//   encodedText is a valid encoding of some originalText that does not have trailing spaces.
//   1 <= rows <= 1000
//   The testcases are generated such that there is only one possible originalText.

func decodeCiphertext(encodedText string, rows int) string {
	if encodedText == "" {
		return ""
	}
	col := len(encodedText) / rows // how many column in matrix
	buf := make([]byte, 0, rows*col)
	for i := 0; i < col; i++ { // put chars diagonally to buf
		for j := 0; j < rows && i+j < col; j++ {
			// matrix[j][i+j] is the original character
			buf = append(buf, encodedText[j*col+i+j])
		}
	}
	// just strip tailing space
	k := len(buf) - 1
	for buf[k] == ' ' {
		k--
	}
	return string(buf[:k+1])
}

func main() {
	for _, v := range []struct {
		e   string
		r   int
		ans string
	}{
		{"ch   ie   pr", 3, "cipher"},
		{"iveo    eed   l te   olc", 4, "i love leetcode"},
		{"iveo    eed   l t    olc", 4, "i love leetcod"},
		{"coding", 1, "coding"},
		{" b  ac", 2, " abc"},
		{"", 5, ""},
	} {
		fmt.Println(decodeCiphertext(v.e, v.r), v.ans)
	}
}

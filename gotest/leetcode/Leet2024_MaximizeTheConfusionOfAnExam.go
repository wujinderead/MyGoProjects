package main

import "fmt"

// https://leetcode.com/problems/maximize-the-confusion-of-an-exam/

// A teacher is writing a test with n true/false questions, with 'T' denoting true and 'F' denoting false.
// He wants to confuse the students by maximizing the number of consecutive questions with the same answer
// (multiple trues or multiple falses in a row).
// You are given a string answerKey, where answerKey[i] is the original answer to the ith question. In addition,
// you are given an integer k, the maximum number of times you may perform the following operation:
//   Change the answer key for any question to 'T' or 'F' (i.e., set answerKey[i] to 'T' or 'F').
//   Return the maximum number of consecutive 'T's or 'F's in the answer key after
//    performing the operation at most k times.
// Example 1:
//   Input: answerKey = "TTFF", k = 2
//   Output: 4
//   Explanation: We can replace both the 'F's with 'T's to make answerKey = "TTTT".
//     There are four consecutive 'T's.
// Example 2:
//   Input: answerKey = "TFFT", k = 1
//   Output: 3
//   Explanation: We can replace the first 'T' with an 'F' to make answerKey = "FFFT".
//     Alternatively, we can replace the second 'T' with an 'F' to make answerKey = "TFFF".
//      In both cases, there are three consecutive 'F's.
// Example 3:
//   Input: answerKey = "TTFTTFTT", k = 1
//   Output: 5
//   Explanation: We can replace the first 'F' to make answerKey = "TTTTTFTT"
//     Alternatively, we can replace the second 'F' to make answerKey = "TTFTTTTT".
//     In both cases, there are five consecutive 'T's.
// Constraints:
//   n == answerKey.length
//   1 <= n <= 5 * 10^4
//   answerKey[i] is either 'T' or 'F'
//   1 <= k <= n

// sliding window: find longest substrings that contains k T's or F's
func maxConsecutiveAnswers(answerKey string, k int) int {
	ts := make([]int, 0)
	fs := make([]int, 0)
	ts = append(ts, -1)
	fs = append(fs, -1)
	for i := range answerKey {
		if answerKey[i] == 'T' {
			ts = append(ts, i)
		} else {
			fs = append(fs, i)
		}
	}
	ts = append(ts, len(answerKey))
	fs = append(fs, len(answerKey))
	max := 0
	if len(ts) <= k+2 || len(fs) <= k+2 {
		return len(answerKey)
	}
	for i := 0; i+k+1 < len(ts); i++ {
		if ts[i+k+1]-ts[i]-1 > max {
			max = ts[i+k+1] - ts[i] - 1
		}
	}
	for i := 0; i+k+1 < len(fs); i++ {
		if fs[i+k+1]-fs[i]-1 > max {
			max = fs[i+k+1] - fs[i] - 1
		}
	}
	return max
}

func main() {
	for _, v := range []struct {
		a      string
		k, ans int
	}{
		{"TTFF", 2, 4},
		{"TFFT", 1, 3},
		{"TTFTTFTT", 1, 5},
		{"F", 1, 1},
		{"TF", 1, 2},
	} {
		fmt.Println(maxConsecutiveAnswers(v.a, v.k), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/plates-between-candles/

// There is a long table with a line of plates and candles arranged on top of it. You are given
// a 0-indexed string s consisting of characters '*' and '|' only, where a '*' represents a plate
// and a '|' represents a candle.
// You are also given a 0-indexed 2D integer array queries where queries[i] = [lefti, righti] denotes
// the substring s[lefti...righti] (inclusive). For each query, you need to find the number of plates
// between candles that are in the substring. A plate is considered between candles if there is at
// least one candle to its left and at least one candle to its right in the substring.
// For example, s = "||**||**|*", and a query [3, 8] denotes the substring "*||**|". The number of plates
// between candles in this substring is 2, as each of the two plates has at least one candle in the
// substring to its left and right.
// Return an integer array answer where answer[i] is the answer to the ith query。
// Example 1:
//   Input: s = "**|**|***|", queries = [[2,5],[5,9]]
//   Output: [2,3]
//   Explanation:
//   - queries[0] has two plates between candles.
//   - queries[1] has three plates between candles.
// Example 2:
//   Input: s = "***|**|*****|**||**|*", queries = [[1,17],[4,5],[14,17],[5,11],[15,16]]
//   Output: [9,0,0,0,0]
//   Explanation:
//     - queries[0] has nine plates between candles.
//     - The other queries have zero plates between candles.
// Constraints:
//   3 <= s.length <= 10^5
//   s consists of '*' and '|' characters.
//   1 <= queries.length <= 105
//   queries[i].length == 2
//   0 <= lefti <= righti < s.length

func platesBetweenCandles(s string, queries [][]int) []int {
	ns := make([]int, len(s)+1)    // ns[i] = number of * in s[:i+1], so number of s in s[i...j]=ns[j+1]-ns[i]
	left := make([]int, len(s)+1)  // nearest | to the left of s[i] = left[i+1]
	right := make([]int, len(s)+1) // nearest | to the right of s[i] = right[i]
	left[0] = -1
	for i := 0; i < len(s); i++ {
		ns[i+1] = ns[i]
		left[i+1] = left[i]
		if s[i] == '*' {
			ns[i+1]++
		} else {
			left[i+1] = i
		}
	}
	right[len(s)] = -1
	for i := len(s) - 1; i >= 0; i-- {
		right[i] = right[i+1]
		if s[i] == '|' {
			right[i] = i
		}
	}
	ans := make([]int, len(queries))
	for i, v := range queries {
		s, e := v[0], v[1]
		// ******|**|****|****|*********
		//   ^   ^            ^      ^
		//   s  right[s]  left[e+1]  e
		// find nearest | right of s; nearest | left of e
		if right[s] >= left[e+1] || right[s] == -1 || left[e+1] == -1 { // skip exceptions
			continue
		}
		// find how many * in [right[s]...left[e+1]], that the answer for query
		ans[i] = ns[left[e+1]+1] - ns[right[s]]
	}
	return ans
}

func main() {
	for _, v := range []struct {
		s   string
		q   [][]int
		ans []int
	}{
		{"**|**|***|", [][]int{{2, 5}, {5, 9}}, []int{2, 3}},
		{"***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}, []int{9, 0, 0, 0, 0}},
		{"||*", [][]int{{2, 2}}, []int{0}},
		{"*****", [][]int{{2, 2}, {0, 4}, {1, 3}}, []int{0, 0, 0}},
	} {
		fmt.Println(platesBetweenCandles(v.s, v.q), v.ans)
	}
}

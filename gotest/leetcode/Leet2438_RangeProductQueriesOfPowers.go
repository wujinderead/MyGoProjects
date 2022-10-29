package main

import "fmt"

// https://leetcode.com/problems/range-product-queries-of-powers/

// Given a positive integer n, there exists a 0-indexed array called powers, composed of the minimum
// number of powers of 2 that sum to n. The array is sorted in non-decreasing order, and there is only
// one way to form the array.
// You are also given a 0-indexed 2D integer array queries, where queries[i] = [lefti, righti]. Each
// queries[i] represents a query where you have to find the product of all powers[j] with
// lefti <= j <= righti.
// Return an array answers, equal in length to queries, where answers[i] is the answer to the iᵗʰ query.
// Since the answer to the iᵗʰ query may be too large, each answers[i] should be returned modulo 10⁹ + 7.
// Example 1:
//   Input: n = 15, queries = [[0,1],[2,2],[0,3]]
//   Output: [2,4,64]
//   Explanation:
//     For n = 15, powers = [1,2,4,8]. It can be shown that powers cannot be a smaller size.
//     Answer to 1st query: powers[0] * powers[1] = 1 * 2 = 2.
//     Answer to 2nd query: powers[2] = 4.
//     Answer to 3rd query: powers[0] * powers[1] * powers[2] * powers[3] = 1 * 2 * 4 * 8 = 64.
//     Each answer modulo 10⁹ + 7 yields the same answer, so [2,4,64] is returned.
// Example 2:
//   Input: n = 2, queries = [[0,0]]
//   Output: [2]
//   Explanation:
//     For n = 2, powers = [2].
//     The answer to the only query is powers[0] = 2. The answer modulo 10⁹ + 7 is the same,
//     so [2] is returned.
// Constraints:
//   1 <= n <= 10⁹
//   1 <= queries.length <= 10⁵
//   0 <= starti <= endi < powers.length

func productQueries(n int, queries [][]int) []int {
	const P = int(1e9) + 7
	sum := 0
	cur := 0
	var powers []int // e.g., n = 13 = 2^0+2^2+2^3, powers=[0,2,3]
	for n > 0 {
		if n%2 > 0 {
			powers = append(powers, cur)
			sum += cur
		}
		cur++
		n = n / 2
	}

	// get 2^i
	two := make([]int, sum+1) // two[i]=2^i
	two[0] = 1
	for i := 1; i < len(two); i++ {
		two[i] = (two[i-1] * 2) % P
	}

	// get prefix of powers
	prefix := make([]int, len(powers)+1)
	for i := range powers {
		prefix[i+1] = prefix[i] + powers[i]
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = two[prefix[q[1]+1]-prefix[q[0]]]
	}
	return ans
}

func main() {
	for _, v := range []struct {
		n   int
		q   [][]int
		ans []int
	}{
		{15, [][]int{{0, 1}, {2, 2}, {0, 3}}, []int{2, 4, 64}},
		{2, [][]int{{0, 0}}, []int{2}},
	} {
		fmt.Println(productQueries(v.n, v.q), v.ans)
	}
}

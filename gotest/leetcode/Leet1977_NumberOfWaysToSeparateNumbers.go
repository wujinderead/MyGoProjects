package main

import "fmt"

// https://leetcode.com/problems/number-of-ways-to-separate-numbers/

// You wrote down many positive integers in a string called num. However, you realized that
// you forgot to add commas to seperate the different numbers. You remember that the list of
// integers was non-decreasing and that no integer had leading zeros.
// Return the number of possible lists of integers that you could have written down to get
// the string num. Since the answer may be large, return it modulo 10^9 + 7.
// Example 1:
//   Input: num = "327"
//   Output: 2
//   Explanation: You could have written down the numbers:
//     3, 27
//     327
// Example 2:
//   Input: num = "094"
//   Output: 0
//   Explanation: No numbers can have leading zeros and all numbers must be positive.
// Example 3:
//   Input: num = "0"
//   Output: 0
//   Explanation: No numbers can have leading zeros and all numbers must be positive.
// Example 4:
//   Input: num = "9999999999999"
//   Output: 101
// Constraints:
//   1 <= num.length <= 3500
//   num consists of digits '0' through '9'.

func numberOfCombinations(num string) int {
	const mod = int(1e9 + 7)
	if num[0] == '0' {
		return 0
	}
	ans := 0
	mapp := make(map[[2]int]int)
	// count the answer for nums[i:] with j digits (i.e. num[i-j:i]) as prefix
	for i := len(num) - 1; i > 0; i-- {
		if num[i] == '0' { // skip leading 0s
			continue
		}
		for j := 1; j <= len(num)-i && i-j >= 0; j++ {
			if num[i-j] == '0' { // skip leading 0s
				continue
			}
			start := i + j
			if !isLessOrEqual(num[i-j:i], num[i:i+j]) {
				start = i + j + 1
			}
			if start > len(num) {
				continue
			}
			count := 0
			// we have fixed num[i-j:i] and variant num[i:k],
			// we check the answer for num[k:] with num[i:k] as prefix.
			for k := start; k < len(num) && len(num)-k >= k-i; k++ {
				tmp := mapp[[2]int{k, k - i}] // here can be O(1) with prefix sum
				count += tmp
				count = count % mod
			}
			count++
			mapp[[2]int{i, j}] = count
			if i-j == 0 { // if num[i-j:i] + num[i:] is the whole num
				ans += count
				ans = ans % mod
			}
		}
	}
	return ans + 1
}

func isLessOrEqual(a, b string) bool {
	for i := range a {
		if a[i] < b[i] {
			return true
		}
		if a[i] > b[i] {
			return false
		}
	}
	return true
}

func main() {
	for _, v := range []struct {
		num string
		ans int
	}{
		{"4321", 2},
		{"22", 2},
		{"327", 2},
		{"345987", 10},
		{"3450987", 7},
		{"3123892327329732", 58},
		{"31230232007329732", 34},
		{"237", 3},
		{"094", 0},
		{"0", 0},
		{"1", 1},
		{"10", 1},
		{"12", 2},
		{"9999999999999", 101},
	} {
		fmt.Println(numberOfCombinations(v.num), v.ans)
	}
}

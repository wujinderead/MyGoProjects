package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// Given a positive integer n, find the number of non-negative integers
// less than or equal to n, whose binary representations do NOT contain consecutive ones.
// Example 1:
//   Input: 5
//   Output: 5
//   Explanation:
//     Here are the non-negative integers <= 5 with their corresponding binary representations:
//     0 : 0
//     1 : 1
//     2 : 10
//     3 : 11
//     4 : 100
//     5 : 101

func findIntegers(num int) int {
	if num < 2 {
		return num + 1
	}
	str := strconv.FormatInt(int64(num), 2)
	fib := make([]int, len(str)+1)
	fib[0], fib[1] = 1, 2
	for i := 2; i < len(fib); i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return findStr(str, fib)
}

func findStr(str string, fib []int) int {
	if len(str) == 1 {
		return int(str[0]-'0') + 1
	}
	if str[1] == '1' { // 11xxxx
		return fib[len(str)-1] + fib[len(str)-2]
	}
	for i := 1; i < len(str); i++ { // 10..01xxx
		if str[i] == '1' {
			return fib[len(str)-1] + findStr(str[i:], fib)
		}
	}
	return fib[len(str)-1] + 1 // 10..0
}

func main() {
	n := 1000000
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		str := strconv.FormatInt(int64(i), 2)
		ans[i] = ans[i-1]
		if strings.Index(str, "11") == -1 {
			ans[i] += 1
		}
		if findIntegers(i) != ans[i] {
			fmt.Println(i, ans[i], findIntegers(i))
		}
	}
}

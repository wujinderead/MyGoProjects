package main

import "fmt"

// https://leetcode.com/problems/find-unique-binary-string/

// Given an array of strings nums containing n unique binary strings each of length n,
// return a binary string of length n that does not appear in nums. If there are multiple
// answers, you may return any of them.
// Example 1:
//   Input: nums = ["01","10"]
//   Output: "11"
//   Explanation: "11" does not appear in nums. "00" would also be correct.
// Example 2:
//   Input: nums = ["00","01"]
//   Output: "11"
//   Explanation: "11" does not appear in nums. "10" would also be correct.
// Example 3:
//   Input: nums = ["111","011","001"]
//   Output: "101"
//   Explanation: "101" does not appear in nums. "000", "010", "100", and "110" would also be correct.
// Constraints:
//   n == nums.length
//   1 <= n <= 16
//   nums[i].length == n
//   nums[i] is either '0' or '1'.
//   All the strings of nums are unique.

func findDifferentBinaryString(nums []string) string {
	ns := make(map[int]struct{})
	for _, n := range nums { // convert nums to numbers
		v := 0
		for _, nn := range n {
			v = v*2 + int(nn-'0')
		}
		ns[v] = struct{}{}
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := ns[i]; !ok { // find an answer, convert it to string
			ans := make([]byte, len(nums))
			for j := range nums {
				ans[len(nums)-j-1] = byte(i%2) + '0'
				i = i / 2
			}
			return string(ans)
		}
	}
	return ""
}

func main() {
	for _, v := range []struct {
		n []string
	}{
		{[]string{"0"}},
		{[]string{"1"}},
		{[]string{"00", "01"}},
		{[]string{"111", "011", "001"}},
		{[]string{"0000", "0001", "0010", "0100"}},
	} {
		fmt.Println(findDifferentBinaryString(v.n))
	}
}

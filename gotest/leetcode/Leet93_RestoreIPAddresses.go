package main

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/restore-ip-addresses/

// Given a string containing only digits, restore it by returning all possible
// valid IP address combinations.
// Example:
//   Input: "25525511135"
//   Output: ["255.255.11.135", "255.255.111.35"]

func restoreIpAddresses(s string) []string {
	re := make([]string, 0)
	nums := make([]string, 4)
	find(s, 0, &re, nums, 0)
	return re
}

func find(s string, start int, re *[]string, nums []string, ind int) {
	if start == len(s) && ind == 4 {
		*re = append(*re, strings.Join(nums, "."))
		return
	}
	if ind == 4 || start == len(s) {
		return
	}
	cur := int(s[start] - '0')
	nums[ind] = s[start : start+1]
	find(s, start+1, re, nums, ind+1)
	if cur > 0 && start+1 < len(s) {
		cur = cur*10 + int(s[start+1]-'0')
		nums[ind] = s[start : start+2]
		find(s, start+2, re, nums, ind+1)
	}
	if cur > 0 && start+2 < len(s) && cur*10+int(s[start+2]-'0') < 256 {
		nums[ind] = s[start : start+3]
		find(s, start+3, re, nums, ind+1)
	}
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("200022"))
	fmt.Println(restoreIpAddresses("020022"))
	fmt.Println(restoreIpAddresses("000256"))
}

package main

import "fmt"

// https://leetcode.com/problems/combination-sum-iii/

// Find all possible combinations of k numbers that add up to a number n, given that only numbers 
// from 1 to 9 can be used and each combination should be a unique set of numbers.
// Note:
//   All numbers will be positive integers.
//   The solution set must not contain duplicate combinations.
// Example 1:
//   Input: k = 3, n = 7
//   Output: [[1,2,4]]
// Example 2:
//   Input: k = 3, n = 9
//   Output: [[1,2,6], [1,3,5], [2,3,4]]

func combinationSum3(k int, n int) [][]int {
	results := make([][]int, 0)
	buf := make([]int, 9)
	dfs(&results, buf, k, 0, 1, n)
	return results
}

// simple backtracking
func dfs(results *[][]int, buf []int, k, ind, start, target int) {
	if (target==0 && ind<k) || (target>0 && ind==k) {
		return
	} 
	if target==0 {
		tmp := make([]int, ind)
		copy(tmp, buf[:ind])
		*results = append(*results, tmp)
		return
	}
	for i:=start; i<=9 && target-i>=0; i++ {
		buf[ind] = i
		dfs(results, buf, k, ind+1, i+1, target-i)
	}
}

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
	fmt.Println(combinationSum3(2, 18))
	fmt.Println(combinationSum3(2, 17))
	fmt.Println(combinationSum3(2, 16))
	fmt.Println(combinationSum3(2, 2))
	fmt.Println(combinationSum3(2, 1))
}
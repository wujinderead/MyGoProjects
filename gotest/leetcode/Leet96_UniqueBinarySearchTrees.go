package leetcode

import "fmt"

// https://leetcode.com/problems/unique-binary-search-trees/

// Given n, how many structurally unique BST's  (binary search trees) that store values 1 ... n?
func numTrees(n int) int {
	num := make([]int, n+1)
	num[0] = 1
	num[1] = 1
	for i := 2; i <= n; i++ {
		// to store values 1...i, store j (1≤j≤i) as root,
		// then, left tree has j-1 values, right tree has i-j values.
		// so we have num[j-1]*num[i-j] ways to 1...i with j as root.
		for j := 1; j <= i; j++ {
			num[i] += num[j-1] * num[i-j]
		}
	}
	return num[n]
}

func main() {
	fmt.Println(numTrees(1))
	fmt.Println(numTrees(2))
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(4))
	fmt.Println(numTrees(5))
}

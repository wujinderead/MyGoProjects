package main

import "fmt"

// https://leetcode.com/problems/edit-distance/

// Given two words word1 and word2, find the minimum number of
// operations required to convert word1 to word2.
// You have the following 3 operations permitted on a word:
// Insert a character
// Delete a character
// Replace a character
// Example:
//   Input: word1 = "horse", word2 = "ros"
//   Output: 3
//   Explanation:
//     horse -> rorse (replace 'h' with 'r')
//     rorse -> rose (remove 'r')
//     rose -> ros (remove 'e')

func minDistance(a string, b string) int {
	// let pa[i] be the prefix of a with length i, i.e., pa[0]="", pa[i]=a[0...i-1]
	// let dis(i, j) be the minimal edit distance between pa[i] and pb[j], then
	// if i*j==0, (pa[i] or pb[j] is ""), dis(i,j) = i+j; if not,
	//                      dis(i-1, j)+1      (delete a[i-1])
	// dis(i, j) = min      dis(i, j-1)+1      (insert b[j-1])
	//                      dis(i-1, j-1) + (a[i-1]==b[j-1] ? 0 : 1)
	// finally, return dis(len(a), len(b))
	// time complexity O(AB), space O(min(A, B))
	if len(a)*len(b) == 0 {
		return len(a) + len(b)
	}
	if len(a) < len(b) {
		a, b = b, a // let b shorter
	}
	prev, cur := make([]int, len(b)+1), make([]int, len(b)+1)
	// line 0
	for j := 0; j <= len(b); j++ {
		prev[j] = j // dis(0, j)=j
	}
	for i := 1; i <= len(a); i++ {
		cur[0] = i // dis(i, 0)=i
		for j := 1; j <= len(b); j++ {
			cur[j] = min(prev[j]+1, cur[j-1]+1) // dis(i-1, j), dis(i, j-1)
			tmp := prev[j-1]                    // dis(i-1, j-1)
			if a[i-1] != b[j-1] {
				tmp += 1
			}
			cur[j] = min(cur[j], tmp)
		}
		prev, cur = cur, prev
	}
	return prev[len(b)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
	fmt.Println(minDistance("ab", ""))
	fmt.Println(minDistance("a", "b"))
	fmt.Println(minDistance("aa", "a"))
	fmt.Println(minDistance("ab", "b"))
	fmt.Println(minDistance("ab", "ba"))
}

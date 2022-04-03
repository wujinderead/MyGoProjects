package main

import "fmt"

// https://leetcode.com/problems/groups-of-strings/

// You are given a 0-indexed array of strings words. Each string consists of lowercase
// English letters only. No letter occurs more than once in any string of words.
// Two strings s1 and s2 are said to be connected if "the set of letters of s2" can be
// obtained from "the set of letters of s1" by any one of the following operations:
//   Adding exactly one letter to the set of the letters of s1.
//   Deleting exactly one letter from the set of the letters of s1.
//   Replacing exactly one letter from the set of the letters of s1 with any letter, including itself.
// The array words can be divided into one or more non-intersecting groups. A string belongs
// to a group if any one of the following is true:
//   It is connected to at least one other string of the group.
//   It is the only string present in the group.
// Note that the strings in words should be grouped in such a manner that a string belonging
// to a group cannot be connected to a string present in any other group. It can be proved
// that such an arrangement is always unique.
// Return an array ans of size 2 where:
//   ans[0] is the maximum number of groups words can be divided into, and
//   ans[1] is the size of the largest group.
// Example 1:
//   Input: words = ["a","b","ab","cde"]
//   Output: [2,3]
//   Explanation:
//     - words[0] can be used to obtain words[1] (by replacing 'a' with 'b'), and
//     words[2] (by adding 'b'). So words[0] is connected to words[1] and words[2].
//     - words[1] can be used to obtain words[0] (by replacing 'b' with 'a'), and
//     words[2] (by adding 'a'). So words[1] is connected to words[0] and words[2].
//     - words[2] can be used to obtain words[0] (by deleting 'b'), and words[1] (by
//     deleting 'a'). So words[2] is connected to words[0] and words[1].
//     - words[3] is not connected to any string in words.
//     Thus, words can be divided into 2 groups ["a","b","ab"] and ["cde"]. The size
//     of the largest group is 3.
// Example 2:
//   Input: words = ["a","ab","abc"]
//   Output: [1,3]
//   Explanation:
//     - words[0] is connected to words[1].
//     - words[1] is connected to words[0] and words[2].
//     - words[2] is connected to words[1].
//     Since all strings are connected to each other, they should be grouped together.
//     Thus, the size of the largest group is 3.
// Constraints:
//   1 <= words.length <= 2 * 10⁴
//   1 <= words[i].length <= 26
//   words[i] consists of lowercase English letters only.
//   No letter occurs more than once in words[i].

func groupStrings(words []string) []int {
	count := make(map[int]int)
	for _, w := range words {
		mask := 0
		for _, b := range []byte(w) {
			mask |= 1 << int(b-'a') // we are interested in the set of letters, so change a word to letter mask
		}
		count[mask] = count[mask] + 1
	}
	for k := range count {
		count[k] = dfs(count, k) // dfs to get the size of each connected sub-graph
	}
	maxv := 0
	for _, v := range count { // find max size
		if v > maxv {
			maxv = v
		}
	}
	return []int{len(count), maxv}
}

func dfs(count map[int]int, k int) int {
	sum := count[k]  // get count of current node
	delete(count, k) // delete to prevent re-visit
	for i := 0; i < 26; i++ {
		kk := k ^ (1 << i) // flip each bit (add a letter: 0 to 1, remove a letter 1 to 0)
		if _, ok := count[kk]; ok {
			sum += dfs(count, kk)
		}
		for j := i + 1; j < 26; j++ {
			if (k>>i)&1 != (k>>j)&1 { // swap two different bits means replace a letter
				kk := k ^ (1 << i) ^ (1 << j)
				if _, ok := count[kk]; ok {
					sum += dfs(count, kk)
				}
			}
		}
	}
	return sum
}

func main() {
	for _, v := range []struct {
		words []string
		ans   []int
	}{
		{[]string{"a", "b", "ab", "cde"}, []int{2, 3}},
		{[]string{"a", "ab", "abc"}, []int{1, 3}},
	} {
		fmt.Println(groupStrings(v.words), v.ans)
	}
}

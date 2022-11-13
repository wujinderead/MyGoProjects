package main

import "fmt"

// https://leetcode.com/problems/most-popular-video-creator/

// You are given two string arrays creators and ids, and an integer array views, all of length n.
// The iᵗʰ video on a platform was created by creator[i], has an id of ids[i], and has views[i] views.
// The popularity of a creator is the sum of the number of views on all of the creator's videos.
// Find the creator with the highest popularity and the id of their most viewed video.
//   If multiple creators have the highest popularity, find all of them.
//   If multiple videos have the highest view count for a creator, find the lexicographically smallest id.
// Return a 2D array of strings answer where answer[i] = [creatori, idi] means that creatori has the
// highest popularity and idi is the id of their most popular video. The answer can be returned in
// any order.
// Example 1:
//   Input: creators = ["alice","bob","alice","chris"], ids = ["one","two","three","four"], views = [5,10,5,4]
//   Output: [["alice","one"],["bob","two"]]
//   Explanation:
//     The popularity of alice is 5 + 5 = 10.
//     The popularity of bob is 10.
//     The popularity of chris is 4.
//     alice and bob are the most popular creators.
//     For bob, the video with the highest view count is "two".
//     For alice, the videos with the highest view count are "one" and "three". Since
//     "one" is lexicographically smaller than "three", it is included in the answer.
// Example 2:
//   Input: creators = ["alice","alice","alice"], ids = ["a","b","c"], views = [1,2,2]
//   Output: [["alice","b"]]
//   Explanation:
//     The videos with id "b" and "c" have the highest view count.
//     Since "b" is lexicographically smaller than "c", it is included in the answer.
// Constraints:
//   n == creators.length == ids.length == views.length
//   1 <= n <= 10⁵
//   1 <= creators[i].length, ids[i].length <= 5
//   creators[i] and ids[i] consist only of lowercase English letters.
//   0 <= views[i] <= 10⁵

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	viewmap := make(map[string]int)
	idmap := make(map[string]struct {
		string
		int
	})
	for i := range creators {
		viewmap[creators[i]] = viewmap[creators[i]] + views[i]
		if idmap[creators[i]].int == 0 || views[i] > idmap[creators[i]].int ||
			(views[i] == idmap[creators[i]].int && ids[i] < idmap[creators[i]].string) {
			idmap[creators[i]] = struct {
				string
				int
			}{string: ids[i], int: views[i]}
		}
	}
	max := 0
	for _, v := range viewmap {
		if v > max {
			max = v
		}
	}
	var ans [][]string
	for k, v := range viewmap {
		if v == max {
			ans = append(ans, []string{k, idmap[k].string})
		}
	}
	return ans
}

func main() {
	for _, v := range []struct {
		cr, id []string
		vi     []int
		ans    [][]string
	}{
		{
			[]string{"alice", "bob", "alice", "chris"},
			[]string{"one", "two", "three", "four"},
			[]int{5, 10, 5, 4},
			[][]string{{"alice", "one"}, {"bob", "two"}},
		},
		{
			[]string{"alice", "bob", "alice", "chris"},
			[]string{"xxx", "two", "three", "four"},
			[]int{5, 10, 5, 4},
			[][]string{{"alice", "three"}, {"bob", "two"}},
		},
		{
			[]string{"alice", "alice", "alice"},
			[]string{"a", "b", "b"},
			[]int{1, 2, 2},
			[][]string{{"alice", "b"}},
		},
		{
			[]string{"a"},
			[]string{"a"},
			[]int{0},
			[][]string{{"a", "a"}},
		},
	} {
		fmt.Println(mostPopularCreator(v.cr, v.id, v.vi), v.ans)
	}
}

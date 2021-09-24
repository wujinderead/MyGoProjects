package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/the-number-of-weak-characters-in-the-game/

// You are playing a game that contains multiple characters, and each of the characters has two
// main properties: attack and defense. You are given a 2D integer array properties where
// properties[i] = [attacki, defensei] represents the properties of the ith character in the game.
// A character is said to be weak if any other character has both attack and defense levels
// strictly greater than this character's attack and defense levels. More formally, a character i
// is said to be weak if there exists another character j where attackj > attacki and defensej > defensei.
// Return the number of weak characters.
// Example 1:
//   Input: properties = [[5,5],[6,3],[3,6]]
//   Output: 0
//   Explanation: No character has strictly greater attack and defense than the other.
// Example 2:
//   Input: properties = [[2,2],[3,3]]
//   Output: 1
//   Explanation: The first character is weak because the second character has a strictly greater attack and defense.
// Example 3:
//   Input: properties = [[1,5],[10,4],[4,3]]
//   Output: 1
//   Explanation: The third character is weak because the second character has a strictly greater attack and defense.
// Constraints:
//   2 <= properties.length <= 10^5
//   properties[i].length == 2
//   1 <= attacki, defensei <= 10^5

// this method, sort descending order of attack; if there are same attack, order with descending defense.
// NOTE: an even better method: sort descending attack and ascending defence, then no need perform group
func numberOfWeakCharacters(properties [][]int) int {
	sort.Sort(p(properties))
	count := 0
	i := 0
	cura := properties[0][0]
	maxb := properties[0][1]
	for i < len(properties) { // skip first group
		if properties[i][0] != cura {
			break
		}
		i++
	}
	for i < len(properties) {
		cura := properties[i][0]
		newmaxb := properties[i][1]
		for i < len(properties) && properties[i][0] == cura { // for current group
			if properties[i][1] < maxb {
				count++
			}
			i++
		}
		if newmaxb > maxb {
			maxb = newmaxb
		}
	}
	return count
}

type p [][]int

func (p p) Len() int {
	return len(p)
}

func (p p) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p p) Less(i, j int) bool {
	if p[i][0] != p[j][0] {
		return p[i][0] > p[j][0]
	}
	return p[i][1] > p[j][1]
}

func main() {
	for _, v := range []struct {
		p   [][]int
		ans int
	}{
		{[][]int{{5, 5}, {6, 3}, {3, 6}}, 0},
		{[][]int{{2, 2}, {3, 3}}, 1},
		{[][]int{{1, 5}, {10, 4}, {4, 3}}, 1},
		{[][]int{{6, 4}, {6, 1}, {5, 7}, {5, 3}, {4, 5}, {4, 3}, {3, 6}}, 4},
	} {
		fmt.Println(numberOfWeakCharacters(v.p), v.ans)
	}
}

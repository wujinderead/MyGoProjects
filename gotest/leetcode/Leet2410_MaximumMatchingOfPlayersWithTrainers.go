package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/maximum-matching-of-players-with-trainers/

// You are given a 0-indexed integer array players, where players[i] represents the ability of
// the iᵗʰ player. You are also given a 0-indexed integer array trainers, where trainers[j]
// represents the training capacity of the jᵗʰ trainer.
// The iᵗʰ player can match with the jᵗʰ trainer if the player's ability is less than or equal
// to the trainer's training capacity. Additionally, the iᵗʰ player can be matched with at most
// one trainer, and the jᵗʰ trainer can be matched with at most one player.
// Return the maximum number of matchings between players and trainers that satisfy these conditions.
// Example 1:
//   Input: players = [4,7,9], trainers = [8,2,5,8]
//   Output: 2
//   Explanation:
//     One of the ways we can form two matchings is as follows:
//     - players[0] can be matched with trainers[0] since 4 <= 8.
//     - players[1] can be matched with trainers[3] since 7 <= 8.
//     It can be proven that 2 is the maximum number of matchings that can be formed.
// Example 2:
//   Input: players = [1,1,1], trainers = [10]
//   Output: 1
//   Explanation:
//     The trainer can be matched with any of the 3 players.
//     Each player can only be matched with one trainer, so the maximum answer is 1.
// Constraints:
//   1 <= players.length, trainers.length <= 10⁵
//   1 <= players[i], trainers[j] <= 10⁹

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	var i, j, count int
	for i < len(players) && j < len(trainers) {
		if players[i] <= trainers[j] {
			i++
			j++
			count++
		} else {
			j++
		}
	}
	return count
}

func main() {
	for _, v := range []struct {
		players  []int
		trainers []int
		ans      int
	}{
		{[]int{4, 7, 9}, []int{8, 2, 5, 8}, 2},
		{[]int{1, 1, 1}, []int{10}, 1},
	} {
		fmt.Println(matchPlayersAndTrainers(v.players, v.trainers), v.ans)
	}
}

package main

import "fmt"

// https://leetcode.com/problems/watering-plants-ii/

// Alice and Bob want to water n plants in their garden. The plants are arranged
// in a row and are labeled from 0 to n - 1 from left to right where the iᵗʰ plant
// is located at x = i.
// Each plant needs a specific amount of water. Alice and Bob have a watering
// can each, initially full. They water the plants in the following way:
// Alice waters the plants in order from left to right, starting from the 0ᵗʰ
// plant. Bob waters the plants in order from right to left, starting from the (n - 1)ᵗʰ plant.
// They begin watering the plants simultaneously.
// If one does not have enough water to completely water the current plant, he/
// she refills the watering can instantaneously.
// It takes the same amount of time to water each plant regardless of how much
// water it needs.
// One cannot refill the watering can early.
// Each plant can be watered either by Alice or by Bob.
// In case both Alice and Bob reach the same plant, the one with more water
// currently in his/her watering can should water this plant. If they have the same
// amount of water, then Alice should water this plant.
// Given a 0-indexed integer array plants of n integers, where plants[i] is the
// amount of water the iᵗʰ plant needs, and two integers capacityA and capacityB
// representing the capacities of Alice's and Bob's watering cans respectively,
// return the number of times they have to refill to water all the plants.
// Example 1:
//   Input: plants = [2,2,3,3], capacityA = 5, capacityB = 5
//   Output: 1
//   Explanation:
//     - Initially, Alice and Bob have 5 units of water each in their watering cans.
//     - Alice waters plant 0, Bob waters plant 3.
//     - Alice and Bob now have 3 units and 2 units of water respectively.
//     - Alice has enough water for plant 1, so she waters it. Bob does not have
//     enough water for plant 2, so he refills his can then waters it.
//     So, the total number of times they have to refill to water all the plants is 0 + 0 + 1 + 0 = 1.
// Example 2:
//   Input: plants = [2,2,3,3], capacityA = 3, capacityB = 4
//   Output: 2
//   Explanation:
//     - Initially, Alice and Bob have 3 units and 4 units of water in their watering cans respectively.
//     - Alice waters plant 0, Bob waters plant 3.
//     - Alice and Bob now have 1 unit of water each, and need to water plants 1 and 2 respectively.
//     - Since neither of them have enough water for their current plants, they refill their cans and then water the plants.
//     So, the total number of times they have to refill to water all the plants is 0 + 1 + 1 + 0 = 2.
// Example 3:
//   Input: plants = [5], capacityA = 10, capacityB = 8
//   Output: 0
//   Explanation:
//     - There is only one plant.
//     - Alice's watering can has 10 units of water, whereas Bob's can has 8 units.
//     Since Alice has more water in her can, she waters this plant.
//     So, the total number of times they have to refill is 0.
// Example 4:
//   Input: plants = [1,2,4,4,5], capacityA = 6, capacityB = 5
//   Output: 2
//   Explanation:
//     - Initially, Alice and Bob have 6 units and 5 units of water in their watering cans respectively.
//     - Alice waters plant 0, Bob waters plant 4.
//     - Alice and Bob now have 5 units and 0 units of water respectively.
//     - Alice has enough water for plant 1, so she waters it. Bob does not have enough water for plant 3, so he refills his can then waters it.
//     - Alice and Bob now have 3 units and 1 unit of water respectively.
//     - Since Alice has more water, she waters plant 2. However, she does not have enough water to completely water this plant. Hence she refills her can then waters it.
//     So, the total number of times they have to refill to water all the plants is 0 + 0 + 1 + 1 + 0 = 2.
// Example 5:
//   Input: plants = [2,2,5,2,2], capacityA = 5, capacityB = 5
//   Output: 1
//   Explanation:
//     Both Alice and Bob will reach the middle plant with the same amount of water, so Alice will water it.
//     She will have 1 unit of water when she reaches it, so she will refill her can.
//     This is the only refill needed.
// Constraints:
//   n == plants.length
//   1 <= n <= 10⁵
//   1 <= plants[i] <= 10⁶
//   max(plants[i]) <= capacityA, capacityB <= 10⁹

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	var i, j, refill int
	pa := capacityA
	pb := capacityB
	for i, j = 0, len(plants)-1; i < j; i, j = i+1, j-1 {
		if pa < plants[i] {
			refill++
			pa = capacityA
		}
		pa -= plants[i]
		if pb < plants[j] {
			refill++
			pb = capacityB
		}
		pb -= plants[j]
	}
	if i == j {
		if pa >= pb && pa < plants[i] {
			refill++
		} else if pa < pb && pb < plants[i] {
			refill++
		}
	}
	return refill
}

func main() {
	for _, v := range []struct {
		n         []int
		a, b, ans int
	}{
		{[]int{2, 2, 3, 3}, 5, 5, 1},
		{[]int{2, 2, 3, 3}, 3, 4, 2},
		{[]int{5}, 10, 8, 0},
		{[]int{1, 2, 4, 4, 5}, 6, 5, 2},
		{[]int{2, 2, 5, 2, 2}, 5, 5, 1},
	} {
		fmt.Println(minimumRefill(v.n, v.a, v.b), v.ans)
	}
}
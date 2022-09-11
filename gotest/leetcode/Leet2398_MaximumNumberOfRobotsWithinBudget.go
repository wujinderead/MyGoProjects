package main

import (
	"container/list"
	"fmt"
)

// https://leetcode.com/problems/maximum-number-of-robots-within-budget/

// You have n robots. You are given two 0-indexed integer arrays, chargeTimes and runningCosts,
// both of length n. The iᵗʰ robot costs chargeTimes[i] units to charge and costs runningCosts[i] units
// to run. You are also given an integer budget.
// The total cost of running k chosen robots is equal to max(chargeTimes) + k * sum(runningCosts), where
// max(chargeTimes) is the largest charge cost among the k robots and sum(runningCosts) is the sum of
// running costs among the k robots.
// Return the maximum number of consecutive robots you can run such that the total cost does not exceed budget.
// Example 1:
//   Input: chargeTimes = [3,6,1,3,4], runningCosts = [2,1,3,4,5], budget = 25
//   Output: 3
//   Explanation:
//     It is possible to run all individual and consecutive pairs of robots within budget.
//     To obtain answer 3, consider the first 3 robots. The total cost will be max(3,6,1) + 3 * sum(2,1,3)
//     = 6 + 3 * 6 = 24 which is less than 25.
//     It can be shown that it is not possible to run more than 3 consecutive robots within budget, so we return 3.
// Example 2:
//   Input: chargeTimes = [11,12,19], runningCosts = [10,8,7], budget = 19
//   Output: 0
//   Explanation: No robot can be run that does not exceed the budget, so we return 0.
// Constraints:
//   chargeTimes.length == runningCosts.length == n
//   1 <= n <= 5 * 10⁴
//   1 <= chargeTimes[i], runningCosts[i] <= 10⁵
//   1 <= budget <= 10¹⁵

// binary search + sliding window max, O(nlogn)
func maximumRobots1(chargeTimes []int, runningCosts []int, budget int64) int {
	bgt := int(budget)
	queue := list.New()
	left, right := 1, len(chargeTimes)
	for left <= right {
		var (
			can bool
			sum int
			k   = (left + right) / 2
		)
		for i := 0; i < k; i++ { // get first k cost
			sum += runningCosts[i]
			if queue.Len() > 0 && i-queue.Front().Value.(int) >= k {
				queue.Remove(queue.Front())
			}
			for queue.Len() > 0 && chargeTimes[queue.Back().Value.(int)] <= chargeTimes[i] {
				queue.Remove(queue.Back())
			}
			queue.PushBack(i)
		}
		cost := chargeTimes[queue.Front().Value.(int)] + k*sum

		if cost <= bgt {
			can = true
		} else { // sliding window
			for i := k; i < len(runningCosts); i++ {
				sum += runningCosts[i] - runningCosts[i-k] // update sliding window sum
				// update sliding window cost
				if queue.Len() > 0 && i-queue.Front().Value.(int) >= k {
					queue.Remove(queue.Front())
				}
				for queue.Len() > 0 && chargeTimes[queue.Back().Value.(int)] <= chargeTimes[i] {
					queue.Remove(queue.Back())
				}
				queue.PushBack(i)
				cost = chargeTimes[queue.Front().Value.(int)] + k*sum
				if cost <= bgt {
					can = true
					break
				}
			}
		}
		if can {
			left = k + 1
		} else {
			right = k - 1
		}
		queue.Init() // empty queue
	}
	return right
}

// sliding window, O(n)
func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	sum := 0
	max := 0
	start := 0
	queue := list.New()
	for i := 0; i < len(chargeTimes); i++ {
		sum += runningCosts[i]
		// for sliding window max
		for queue.Len() > 0 && chargeTimes[queue.Back().Value.(int)] <= chargeTimes[i] {
			queue.Remove(queue.Back())
		}
		queue.PushBack(i)
		cost := chargeTimes[queue.Front().Value.(int)] + (i-start+1)*sum
		if cost <= int(budget) { // update max if can
			if i-start+1 > max {
				max = i - start + 1
			}
			continue
		}
		// cost > budget, shrink the range
		for start <= i {
			sum -= runningCosts[start]
			if start == queue.Front().Value.(int) {
				queue.Remove(queue.Front())
			}
			start++
			if queue.Len() > 0 && chargeTimes[queue.Front().Value.(int)]+(i-start+1)*sum <= int(budget) {
				break
			}
		}
	}
	return max
}

func main() {
	for _, v := range []struct {
		chargeTimes  []int
		runningCosts []int
		budget       int64
		ans          int
	}{
		{[]int{3, 6, 1, 3, 4}, []int{2, 1, 3, 4, 5}, 25, 3},
		{[]int{11, 12, 19}, []int{10, 8, 7}, 19, 0},
	} {
		fmt.Println(maximumRobots1(v.chargeTimes, v.runningCosts, v.budget), v.ans)
		fmt.Println(maximumRobots(v.chargeTimes, v.runningCosts, v.budget), v.ans)
	}
}
